package chargepoint

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/reservation"
	log "github.com/sirupsen/logrus"
	"github.com/xBlaz3kx/ChargePi-go/internal/api"
	"github.com/xBlaz3kx/ChargePi-go/internal/chargepoint/util"
	v16 "github.com/xBlaz3kx/ChargePi-go/internal/chargepoint/v16"
	"github.com/xBlaz3kx/ChargePi-go/internal/components/auth"
	connectorManager "github.com/xBlaz3kx/ChargePi-go/internal/components/connector-manager"
	s "github.com/xBlaz3kx/ChargePi-go/internal/components/settings"
	"github.com/xBlaz3kx/ChargePi-go/internal/models/charge-point"
	"github.com/xBlaz3kx/ChargePi-go/internal/models/settings"
	"github.com/xBlaz3kx/ChargePi-go/internal/pkg/grpc"
	"github.com/xBlaz3kx/ChargePi-go/pkg/cache"
	"github.com/xBlaz3kx/ChargePi-go/pkg/logging"
	"github.com/xBlaz3kx/ChargePi-go/pkg/scheduler"
	"github.com/xBlaz3kx/ocppManager-go/configuration"
	"os"
	"os/signal"
	"syscall"
)

func CreateChargePoint(
	ctx context.Context,
	protocolVersion settings.ProtocolVersion,
	logger *log.Logger,
	manager connectorManager.Manager,
	sch *gocron.Scheduler,
	authCache *auth.Cache,
	hardware settings.Hardware,
) chargePoint.ChargePoint {
	switch protocolVersion {
	case settings.OCPP16:
		// Create the client
		return v16.NewChargePoint(
			manager,
			sch,
			authCache,
			v16.WithDisplayFromSettings(ctx, hardware.Lcd),
			v16.WithReaderFromSettings(ctx, hardware.TagReader),
			v16.WithLogger(logger),
		)
	case settings.OCPP201:
		logger.Fatal("Version 2.0.1 is not supported yet.")
		return nil
	default:
		logger.WithField("protocolVersion", protocolVersion).Fatal("Protocol version not supported")
		return nil
	}
}

func Run(isDebug bool, settingsFilePath, configurationFilePath, connectorsFolderPath, authFilePath string, exposeApi bool, apiAddress string, apiPort int) {
	var (
		// ChargePoint components
		handler   chargePoint.ChargePoint
		config    *settings.Settings
		mem       = cache.GetCache()
		authCache = auth.NewAuthCache(authFilePath)
		logger    = log.StandardLogger()
		manager   = connectorManager.GetManager()
		sch       = scheduler.GetScheduler()
		// Api
		apiReceiveChannel = make(chan api.Message, 5)
		apiSendChannel    = make(chan api.Message, 5)
		// Execution
		ctx, cancel = context.WithCancel(context.Background())
		quitChannel = make(chan os.Signal, 1)
	)
	defer cancel()
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	// Read settings file and cache it
	config = s.GetSettings(mem, settingsFilePath)
	go authCache.LoadAuthFile(authFilePath)

	// Create the logger
	logging.Setup(logger, config.ChargePoint.Logging, isDebug)

	var (
		chargePointInfo = config.ChargePoint.Info
		hardware        = config.ChargePoint.Hardware
		serverUrl       = util.CreateConnectionUrl(config.ChargePoint)
		protocolVersion = settings.ProtocolVersion(chargePointInfo.ProtocolVersion)
	)

	// Setup OCPP configuration manager
	s.SetupOcppConfigurationManager(
		configurationFilePath,
		configuration.ProtocolVersion(config.ChargePoint.Info.ProtocolVersion),
		core.ProfileName,
		reservation.ProfileName)

	handler = CreateChargePoint(ctx, protocolVersion, logger, manager, sch, authCache, hardware)

	// Initialize the client
	handler.Init(config)

	// Add connectors from file
	connectors := s.GetConnectors(mem, connectorsFolderPath)
	handler.AddConnectors(connectors)

	// Finally, connect to the central system
	handler.Connect(ctx, serverUrl)

	if exposeApi {
		// Expose the API endpoints
		go func() {
			address := fmt.Sprintf("%s:%d", apiAddress, apiPort)
			grpc.CreateAndRunGrpcServer(address, apiSendChannel, apiReceiveChannel)
		}()
	}

Loop:
	for {
		select {
		// Capture the terminate signal
		case <-quitChannel:
			handler.CleanUp(core.ReasonLocal)
			break Loop
		case <-ctx.Done():
			handler.CleanUp(core.ReasonPowerLoss)
			break Loop
		}
	}
}