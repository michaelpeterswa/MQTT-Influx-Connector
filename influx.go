package main

import (
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxConn struct {
	client influxdb2.Client
}

func initInflux() (conn *InfluxConn) {

	log.Println("Creating InfluxDB v2 connection...")
	client := influxdb2.NewClient(settings.InfluxAddress, settings.InfluxToken)
	conn = &InfluxConn{client}
	return conn

}

func getOrganization() string {
	return settings.InfluxOrganization
}

func getBucket() string {
	return settings.InfluxBucket
}

func (conn *InfluxConn) writeBME280SensorData(reading BME280, st SubTopic) {
	p := influxdb2.NewPointWithMeasurement(settings.InfluxMeasurement).
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

	write := conn.client.WriteAPIBlocking(getOrganization(), getBucket())
	write.WritePoint(context.Background(), p)

}
