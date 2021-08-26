package chargepoint

import (
	"errors"
	"fmt"
	"github.com/lorenzodonini/ocpp-go/ocpp"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
	"github.com/xBlaz3kx/ChargePi-go/data"
	"github.com/xBlaz3kx/ChargePi-go/hardware"
	"github.com/xBlaz3kx/ChargePi-go/settings"
	"log"
)

type Connector struct {
	EvseId            int
	ConnectorId       int
	ConnectorType     string
	ConnectorStatus   core.ChargePointStatus
	ErrorCode         ocpp.ErrorCode
	relay             *hardware.Relay
	powerMeter        *hardware.PowerMeter
	PowerMeterEnabled bool
	MaxChargingTime   int
	reservationId     int
	session           *data.Session
}

// NewConnector Create a new connector object from the provided arguments. EvseId, connectorId and maxChargingTime must be greater than zero.
// At creation, it makes an empty session and defaults the status to Available.
func NewConnector(EvseId int, connectorId int, connectorType string, relay *hardware.Relay, powerMeter *hardware.PowerMeter, powerMeterEnabled bool, maxChargingTime int) (Connector, error) {
	if maxChargingTime <= 0 {
		maxChargingTime = 180
	}
	if EvseId <= 0 || connectorId <= 0 {
		return Connector{}, errors.New("invalid evse or connector id")
	}
	if relay == nil {
		return Connector{}, fmt.Errorf("relay pointer cannot be nil")
	}
	return Connector{
		EvseId:            EvseId,
		ConnectorId:       connectorId,
		ConnectorType:     connectorType,
		relay:             relay,
		powerMeter:        powerMeter,
		reservationId:     -1,
		PowerMeterEnabled: powerMeterEnabled,
		MaxChargingTime:   maxChargingTime,
		ConnectorStatus:   core.ChargePointStatusAvailable,
		session: &data.Session{
			TransactionId: "",
			TagId:         "",
			Started:       "",
			Consumption:   []types.MeterValue{},
		},
	}, nil
}

// StartCharging Start charging a connector if connector is available and session could be started.
// It turns on the relay (even if negative logic applies).
func (connector *Connector) StartCharging(transactionId string, tagId string) error {
	var hasSessionStarted = false
	if (connector.IsAvailable() || connector.IsPreparing()) && !connector.session.IsActive {
		hasSessionStarted = connector.session.StartSession(transactionId, tagId)
		if hasSessionStarted {
			connector.relay.On()
			return nil
		}
		return errors.New("transaction or tag ID invalid")
	}
	return errors.New("invalid connector status or session already active")
}

// StopCharging Stops charging the connector by turning the relay off and ends the session.
func (connector *Connector) StopCharging() error {
	if connector.IsCharging() || connector.IsPreparing() {
		connector.session.EndSession()
		connector.relay.Off()
		return nil
	}
	return errors.New("connector not charging")
}

// SamplePowerMeter Get a sample from the power meter. The measurands argument takes the list of all the types of the measurands to sample.
// It will add all the samples to the connector's Session if it is active.
func (connector *Connector) SamplePowerMeter(measurands []types.Measurand) {
	if !connector.PowerMeterEnabled || connector.powerMeter != nil {
		return
	}
	log.Println("Sampling connector", connector.ConnectorId)
	var samples []types.SampledValue
	for _, measurand := range measurands {
		var value = 0.0
		switch measurand {
		case types.MeasurandEnergyActiveExportInterval:
			value = connector.powerMeter.GetEnergy()
			break
		case types.MeasurandCurrentExport:
			value = connector.powerMeter.GetCurrent()
			break
		case types.MeasurandPowerActiveExport:
			value = connector.powerMeter.GetPower()
			break
		case types.MeasurandVoltage:
			value = connector.powerMeter.GetVoltage()
			break
		}
		if value != 0.0 {
			samples = append(samples, types.SampledValue{Value: fmt.Sprintf("%.3f", value), Measurand: measurand})
		}
	}
	connector.session.AddSampledValue(samples)
}

func (connector *Connector) IsAvailable() bool {
	return connector.ConnectorStatus == core.ChargePointStatusAvailable
}
func (connector *Connector) IsCharging() bool {
	return connector.ConnectorStatus == core.ChargePointStatusCharging
}
func (connector *Connector) IsPreparing() bool {
	return connector.ConnectorStatus == core.ChargePointStatusPreparing
}
func (connector *Connector) IsReserved() bool {
	return connector.ConnectorStatus == core.ChargePointStatusReserved
}
func (connector *Connector) IsUnavailable() bool {
	return connector.ConnectorStatus == core.ChargePointStatusUnavailable
}

func (connector *Connector) SetStatus(status core.ChargePointStatus) {
	connector.ConnectorStatus = status
	settings.UpdateConnectorStatus(connector.EvseId, connector.ConnectorId, status)
}

func (connector *Connector) ReserveConnector(reservationId int) error {
	if reservationId <= 0 {
		return fmt.Errorf("reservation id is invalid")
	}
	if !connector.IsAvailable() {
		return fmt.Errorf("connector status invalid")
	}
	connector.reservationId = reservationId
	connector.SetStatus(core.ChargePointStatusReserved)
	return nil
}
func (connector *Connector) RemoveReservation() error {
	if !connector.IsReserved() {
		return fmt.Errorf("connector status invalid")
	}
	connector.reservationId = -1
	connector.SetStatus(core.ChargePointStatusAvailable)
	return nil
}

func (connector *Connector) GetReservationId() int {
	return connector.reservationId
}

// ResumeCharging Resumes or restores the charging state after boot if a charging session was active.
func (connector *Connector) ResumeCharging(session data.Session) error {
	if connector.IsCharging() || connector.IsPreparing() {
		hasSessionStarted := connector.session.StartSession(session.TransactionId, session.TagId)
		if hasSessionStarted {
			connector.relay.On()
			connector.session.Started = session.Started
			connector.session.Consumption = append(connector.session.Consumption, session.Consumption...)
			return nil
		}
		return errors.New("cannot resume session: unable to start session")
	}
	return errors.New("cannot resume session: invalid connector status")
}
