package influx

import (
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/structs"
	gorenogymodbus "github.com/michaelpeterswa/go-renogy-modbus"
	"go.uber.org/zap"
)

type InfluxConn struct {
	client       influxdb2.Client
	Logger       *zap.Logger
	measurement  string
	bucket       string
	organization string
}

func InitInflux(settings *structs.MQTTInfluxConnectorSettings) *InfluxConn {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("unable to acquire zap logger: %s\n", err.Error())
	}

	log.Println("Creating InfluxDB v2 connection...")
	client := influxdb2.NewClient(settings.InfluxAddress, settings.InfluxToken)
	conn := &InfluxConn{client: client, Logger: logger, measurement: settings.InfluxMeasurement, bucket: settings.InfluxBucket, organization: settings.InfluxOrganization}
	return conn

}

func (i *InfluxConn) GetOrganization() string {
	return i.organization
}

func (i *InfluxConn) GetBucket() string {
	return i.bucket
}

func (conn *InfluxConn) WriteBME280SensorData(reading structs.BME280, st structs.SubTopic) {
	p := influxdb2.NewPointWithMeasurement(conn.measurement).
		AddTag("feeder", "MQTT-Influx-Connector").
		AddTag("type", st.Type).
		AddTag("location", st.Location).
		AddTag("room", st.Room).
		AddTag("name", st.Name).
		AddTag("field", st.Field).
		AddField("temperature", reading.Temperature).
		AddField("humidity", reading.Humidity).
		AddField("pressure", reading.Pressure).
		AddField("rssi", reading.RSSI).
		SetTime(time.Unix(int64(reading.Timestamp), 0))

	// support for the optional voltage reading... only if populated
	if reading.Voltage != 0 {
		p.AddField("voltage", reading.Voltage)
	}

	write := conn.client.WriteAPIBlocking(conn.GetOrganization(), conn.GetBucket())
	write.WritePoint(context.Background(), p)

}

func (conn *InfluxConn) WriteTSL2561SensorData(reading structs.TSL2561, st structs.SubTopic) {
	p := influxdb2.NewPointWithMeasurement(conn.measurement).
		AddTag("feeder", "MQTT-Influx-Connector").
		AddTag("type", st.Type).
		AddTag("location", st.Location).
		AddTag("room", st.Room).
		AddTag("name", st.Name).
		AddTag("field", st.Field).
		AddField("lux", reading.Lux).
		AddField("rssi", reading.RSSI).
		SetTime(time.Unix(int64(reading.Timestamp), 0))

	write := conn.client.WriteAPIBlocking(conn.GetOrganization(), conn.GetBucket())
	write.WritePoint(context.Background(), p)

}

func (conn *InfluxConn) WritePMSA003ISensorData(reading structs.PMSA003I, st structs.SubTopic) {
	p := influxdb2.NewPointWithMeasurement(conn.measurement).
		AddTag("feeder", "MQTT-Influx-Connector").
		AddTag("type", st.Type).
		AddTag("location", st.Location).
		AddTag("room", st.Room).
		AddTag("name", st.Name).
		AddTag("field", st.Field).
		AddField("pm10s", reading.Pm10S).
		AddField("pm25s", reading.Pm25S).
		AddField("pm100s", reading.Pm100S).
		AddField("pm10e", reading.Pm10E).
		AddField("pm25e", reading.Pm25E).
		AddField("pm100e", reading.Pm100E).
		AddField("p03um", reading.P03Um).
		AddField("p05um", reading.P05Um).
		AddField("p10um", reading.P10Um).
		AddField("p25um", reading.P25Um).
		AddField("p50um", reading.P50Um).
		AddField("p100um", reading.P100Um).
		AddField("rssi", reading.RSSI).
		SetTime(time.Unix(int64(reading.Timestamp), 0))

	write := conn.client.WriteAPIBlocking(conn.GetOrganization(), conn.GetBucket())
	write.WritePoint(context.Background(), p)

}

func (conn *InfluxConn) WriteMessageReceived() {
	p := influxdb2.NewPointWithMeasurement(conn.measurement).
		AddTag("message", "received").
		AddField("hit", 1).
		SetTime(time.Now())

	write := conn.client.WriteAPIBlocking(conn.GetOrganization(), conn.GetBucket())
	write.WritePoint(context.Background(), p)
}

func (conn *InfluxConn) WriteRenogyChargeControllerSensorData(reading gorenogymodbus.DynamicControllerInformation, st structs.SubTopic) {
	p := influxdb2.NewPointWithMeasurement(conn.measurement).
		AddTag("feeder", "MQTT-Influx-Connector").
		AddTag("type", st.Type).
		AddTag("location", st.Location).
		AddTag("room", st.Room).
		AddTag("name", st.Name).
		AddTag("field", st.Field).
		AddField("battery_capacity_soc", reading.BatteryCapacitySOC).
		AddField("battery_voltage", reading.BatteryVoltage).
		AddField("charging_current", reading.ChargingCurrent).
		AddField("controller_temperature", reading.ControllerTemperature).
		AddField("battery_temperature", reading.BatteryTemperature).
		AddField("street_light_load_voltage", reading.StreetLightLoadVoltage).
		AddField("street_light_load_current", reading.StreetLightLoadCurrent).
		AddField("street_light_load_power", reading.StreetLightLoadPower).
		AddField("solar_panel_voltage", reading.SolarPanelVoltage).
		AddField("solar_panel_current", reading.SolarPanelCurrent).
		AddField("charging_power", reading.ChargingPower).
		AddField("battery_minimum_voltage_current_day", reading.BatteryMinimumVoltageCurrentDay).
		AddField("battery_maximum_voltage_current_day", reading.BatteryMaximumVoltageCurrentDay).
		AddField("maximum_charging_current_current_day", reading.MaximumChargingCurrentCurrentDay).
		AddField("maximum_discharging_current_current_day", reading.MaximumDischargingCurrentCurrentDay).
		AddField("maximum_charging_power_current_day", reading.MaximumChargingPowerCurrentDay).
		AddField("maximum_discharging_power_current_day", reading.MaximumDischargingPowerCurrentDay).
		AddField("charging_amp_hours_current_day", reading.ChargingAmpHoursCurrentDay).
		AddField("discharging_amp_hours_current_day", reading.DischargingAmpHoursCurrentDay).
		AddField("power_generation_current_day", reading.PowerGenerationCurrentDay).
		AddField("power_consumption_current_day", reading.PowerConsumptionCurrentDay).
		AddField("total_operating_days", reading.TotalOperatingDays).
		AddField("total_battery_over_discharges", reading.TotalBatteryOverDischarges).
		AddField("total_battery_full_charges", reading.TotalBatteryFullCharges).
		AddField("total_charging_amp_hours", reading.TotalChargingAmpHours).
		AddField("total_discharging_amp_hours", reading.TotalDischargingAmpHours).
		AddField("cumulative_power_generation", reading.CumulativePowerGeneration).
		AddField("cumulative_power_consumption", reading.CumulativePowerConsumption).
		AddField("street_light_status", reading.StreetLightStatus).
		AddField("street_light_brightness", reading.StreetLightBrightness).
		AddField("charging_state", reading.ChargingState).
		AddField("controller_faults", reading.ControllerFaults).
		SetTime(time.Now())

	write := conn.client.WriteAPIBlocking(conn.GetOrganization(), conn.GetBucket())
	write.WritePoint(context.Background(), p)

}

// type DynamicControllerInformation struct {
// 	TotalChargingAmpHours               decimal.Decimal `json:"total_charging_amp_hours"`                // 0x118-119
// 	TotalDischargingAmpHours            decimal.Decimal `json:"total_discharging_amp_hours"`             // 0x11A-11B
// 	CumulativePowerGeneration           decimal.Decimal `json:"cumulative_power_generation"`             // 0x11C-11D
// 	CumulativePowerConsumption          decimal.Decimal `json:"cumulative_power_consumption"`            // 0x11E-11F
// 	StreetLightStatus                   bool            `json:"street_light_status"`                     // 0x120 (eight higher bits)
// 	StreetLightBrightness               int             `json:"street_light_brightness"`                 // 0x120 (eight higher bits)
// 	ChargingState                       string          `json:"charging_state"`                          // 0x120 (eight lower bits)
// 	ControllerFaults                    []string        `json:"controller_faults"`                       // 0x121-122
// }
