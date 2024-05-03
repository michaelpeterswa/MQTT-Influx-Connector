package timescale

import (
	"context"
	"fmt"
	"time"

	_ "embed"

	"github.com/jackc/pgx/v5"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/structs"
	gorenogymodbus "github.com/michaelpeterswa/go-renogy-modbus"
)

//go:embed queries/write_bme280.pgsql
var writeBME280SQL string

//go:embed queries/write_pmsa003i.pgsql
var writePMSA003ISQL string

//go:embed queries/write_tsl2561.pgsql
var writeTSL2561SQL string

//go:embed queries/write_veml7700.pgsql
var writeVEML7700SQL string

//go:embed queries/write_renogychargecontroller.pgsql
var writeRenogyChargeControllerSQL string

type TimescaleConn struct {
	Conn *pgx.Conn
}

func InitTimescale(ctx context.Context, settings *structs.MQTTInfluxConnectorSettings) (*TimescaleConn, error) {
	conn, err := pgx.Connect(ctx, settings.TimescaleAddress)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return &TimescaleConn{
		Conn: conn,
	}, nil
}

func (t *TimescaleConn) Close(ctx context.Context) {
	t.Conn.Close(ctx)
}

func (t *TimescaleConn) WriteBME280SensorData(ctx context.Context, reading structs.BME280, st structs.SubTopic) error {
	_, err := t.Conn.Exec(
		ctx,
		writeBME280SQL,
		time.Unix(int64(reading.Timestamp), 0),
		st.Type,
		st.Location,
		st.Room,
		st.Name,
		st.Field,
		reading.Temperature,
		reading.Humidity,
		reading.Pressure,
		reading.RSSI,
		reading.Voltage,
	)
	if err != nil {
		return fmt.Errorf("unable to write bme280 data to timescale: %w", err)
	}

	return nil
}

func (t *TimescaleConn) WriteTSL2561SensorData(ctx context.Context, reading structs.TSL2561, st structs.SubTopic) error {
	_, err := t.Conn.Exec(
		ctx,
		writeTSL2561SQL,
		time.Unix(int64(reading.Timestamp), 0),
		st.Type,
		st.Location,
		st.Room,
		st.Name,
		st.Field,
		reading.Lux,
		reading.RSSI,
	)
	if err != nil {
		return fmt.Errorf("unable to write tsl2561 data to timescale: %w", err)
	}

	return nil
}

func (t *TimescaleConn) WriteVEML7700SensorData(ctx context.Context, reading structs.VEML7700, st structs.SubTopic) error {
	_, err := t.Conn.Exec(
		ctx,
		writeVEML7700SQL,
		time.Unix(int64(reading.Timestamp), 0),
		st.Type,
		st.Location,
		st.Room,
		st.Name,
		st.Field,
		reading.Lux,
		reading.RSSI,
	)
	if err != nil {
		return fmt.Errorf("unable to write tsl2561 data to timescale: %w", err)
	}

	return nil
}

func (t *TimescaleConn) WritePMSA003ISensorData(ctx context.Context, reading structs.PMSA003I, st structs.SubTopic) error {
	_, err := t.Conn.Exec(
		ctx,
		writePMSA003ISQL,
		time.Unix(int64(reading.Timestamp), 0),
		st.Type,
		st.Location,
		st.Room,
		st.Name,
		st.Field,
		reading.Pm10S,
		reading.Pm25S,
		reading.Pm100S,
		reading.Pm10E,
		reading.Pm25E,
		reading.Pm100E,
		reading.P03Um,
		reading.P05Um,
		reading.P10Um,
		reading.P25Um,
		reading.P50Um,
		reading.P100Um,
		reading.RSSI,
	)
	if err != nil {
		return fmt.Errorf("unable to write pmsa003i data to timescale: %w", err)
	}

	return nil
}

func (t *TimescaleConn) WriteRenogyChargeControllerSensorData(ctx context.Context, reading gorenogymodbus.DynamicControllerInformation, st structs.SubTopic) error {
	_, err := t.Conn.Exec(
		ctx,
		writeRenogyChargeControllerSQL,
		time.Now(),
		st.Type,
		st.Location,
		st.Room,
		st.Name,
		st.Field,
		reading.BatteryCapacitySOC,
		reading.BatteryVoltage,
		reading.ChargingCurrent,
		reading.ControllerTemperature,
		reading.BatteryTemperature,
		reading.StreetLightLoadVoltage,
		reading.StreetLightLoadCurrent,
		reading.StreetLightLoadPower,
		reading.SolarPanelVoltage,
		reading.SolarPanelCurrent,
		reading.ChargingPower,
		reading.BatteryMinimumVoltageCurrentDay,
		reading.BatteryMaximumVoltageCurrentDay,
		reading.MaximumChargingCurrentCurrentDay,
		reading.MaximumDischargingCurrentCurrentDay,
		reading.MaximumChargingPowerCurrentDay,
		reading.MaximumDischargingPowerCurrentDay,
		reading.ChargingAmpHoursCurrentDay,
		reading.DischargingAmpHoursCurrentDay,
		reading.PowerGenerationCurrentDay,
		reading.PowerConsumptionCurrentDay,
		reading.TotalOperatingDays,
		reading.TotalBatteryOverDischarges,
		reading.TotalBatteryFullCharges,
		reading.TotalChargingAmpHours,
		reading.TotalDischargingAmpHours,
		reading.CumulativePowerGeneration,
		reading.CumulativePowerConsumption,
		reading.StreetLightStatus,
		reading.StreetLightBrightness,
		reading.ChargingState,
		reading.ControllerFaults,
		nil,
	)
	if err != nil {
		return fmt.Errorf("unable to write renogychargecontroller data to timescale: %w", err)
	}

	return nil
}
