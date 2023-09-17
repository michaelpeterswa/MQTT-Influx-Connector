package handlers

import (
	"context"
	"encoding/json"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/helpers"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/influx"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/structs"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/timescale"
	gorenogymodbus "github.com/michaelpeterswa/go-renogy-modbus"
	"go.uber.org/zap"
)

func OnMessageReceived(influxConn *influx.InfluxConn, timescaleConn *timescale.TimescaleConn) func(client MQTT.Client, message MQTT.Message) {
	return func(client MQTT.Client, message MQTT.Message) {
		st := helpers.GetSubTopicFromString(message.Topic())

		influxConn.WriteMessageReceived()

		switch v := st.Name; v {
		case "bme280":
			var reading structs.BME280
			err := json.Unmarshal(message.Payload(), &reading)
			if err != nil {
				influxConn.Logger.Error("failed to unmarshal payload", zap.String("sensor", st.Name))
			} else {
				influxConn.WriteBME280SensorData(reading, st)
				err := timescaleConn.WriteBME280SensorData(context.TODO(), reading, st)
				if err != nil {
					influxConn.Logger.Error("failed to write to timescale", zap.Error(err))
				}
			}
		case "tsl2561":
			var reading structs.TSL2561
			err := json.Unmarshal(message.Payload(), &reading)
			if err != nil {
				influxConn.Logger.Error("failed to unmarshal payload", zap.String("sensor", st.Name))
			} else {
				influxConn.WriteTSL2561SensorData(reading, st)
			}
		case "tsl2591":
			var reading structs.TSL2561
			err := json.Unmarshal(message.Payload(), &reading)
			if err != nil {
				influxConn.Logger.Error("failed to unmarshal payload", zap.String("sensor", st.Name))
			} else {
				influxConn.WriteTSL2561SensorData(reading, st)
			}
		case "pmsa003i":
			var reading structs.PMSA003I
			err := json.Unmarshal(message.Payload(), &reading)
			if err != nil {
				influxConn.Logger.Error("failed to unmarshal payload", zap.String("sensor", st.Name))
			} else {
				influxConn.WritePMSA003ISensorData(reading, st)
			}
		case "renogy-charge-controller":
			var reading gorenogymodbus.DynamicControllerInformation
			err := json.Unmarshal(message.Payload(), &reading)
			if err != nil {
				influxConn.Logger.Error("failed to unmarshal payload", zap.String("sensor", st.Name))
			} else {
				influxConn.WriteRenogyChargeControllerSensorData(reading, st)
			}
		default:
			influxConn.Logger.Warn("unknown sensor type", zap.String("sensor", st.Name))
		}
	}
}
