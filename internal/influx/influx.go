package influx

import (
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/structs"
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
