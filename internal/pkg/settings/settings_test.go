package settings

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kkyr/fig"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"github.com/xBlaz3kx/ChargePi-go/internal/chargepoint/components/hardware/evcc"
	"github.com/xBlaz3kx/ChargePi-go/internal/pkg/models/settings"
	"os/exec"
	"testing"
	"time"
)

const (
	fileName = "evse-1.json"
)

type SettingsManagerTestSuite struct {
	suite.Suite
	evse       settings.EVSE
	session    settings.Session
	evcc       settings.EVCC
	powerMeter settings.PowerMeter
}

func (s *SettingsManagerTestSuite) SetupTest() {
	s.session = settings.Session{
		IsActive:      false,
		TransactionId: "",
		TagId:         "",
		Started:       "",
		Consumption:   nil,
	}

	s.evcc = settings.EVCC{
		Type: evcc.Relay,
	}

	s.powerMeter = settings.PowerMeter{
		Enabled: false,
		Type:    "",
	}

	s.evse = settings.EVSE{
		EvseId:     1,
		Status:     "Available",
		Session:    s.session,
		EVCC:       s.evcc,
		PowerMeter: s.powerMeter,
	}

	evse1 := viper.New()
	evse1.SetConfigFile(fileName)

	marshal, err := json.Marshal(s.evse)
	s.Require().NoError(err)

	err = evse1.ReadConfig(bytes.NewBuffer(marshal))
	s.Require().NoError(err)

	EVSESettings.Store(fmt.Sprintf("evse%d", 1), evse1)
}

func (s *SettingsManagerTestSuite) TestUpdateSessionInfo() {
	var (
		evseFromFile settings.EVSE
		newSession   = settings.Session{
			IsActive:      true,
			TransactionId: "Transaction1234",
			TagId:         "Tag1234",
			Started:       "",
			Consumption:   nil,
		}
	)

	UpdateSession(s.evse.EvseId, &newSession)

	err := fig.Load(&evseFromFile, fig.File(fileName))
	s.Assert().FileExists("./" + fileName)
	s.Assert().NoError(err)
	s.Assert().EqualValues(newSession, evseFromFile.Session)

	// Clean up
	cmd := exec.Command("rm", fileName)
	_ = cmd.Run()
}

func (s *SettingsManagerTestSuite) TestUpdateConnectorStatus() {
	if testing.Short() {
		return
	}

	cmd := exec.Command("touch", fileName)
	err := cmd.Run()
	s.Require().NoError(err)

	var evseFromFile settings.EVSE

	UpdateEVSEStatus(s.evse.EvseId, core.ChargePointStatusCharging)

	time.Sleep(time.Second)

	err = fig.Load(&evseFromFile, fig.File(fileName))
	s.Assert().FileExists("./" + fileName)
	s.Assert().NoError(err)
	s.Assert().EqualValues(core.ChargePointStatusCharging, evseFromFile.Status)

	// Clean up
	cmd = exec.Command("rm", fileName)
	_ = cmd.Run()
}

func TestSettingsManager(t *testing.T) {
	suite.Run(t, new(SettingsManagerTestSuite))
}