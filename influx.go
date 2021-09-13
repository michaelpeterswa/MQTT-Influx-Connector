package main

import (
	"context"
	"log"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxConn struct {
	client influxdb2.Client
}

func initInflux() (conn *InfluxConn) {

	log.Println("Creating InfluxDB v2 connection...")
	client := influxdb2.NewClient(os.Getenv("INFLUX_ADDRESS"), os.Getenv("INFLUX_TOKEN"))
	conn = &InfluxConn{client}
	return conn

}

func getOrganization() string {
	return os.Getenv("INFLUX_ORGANIZATION")
}

func getBucket() string {
	return os.Getenv("INFLUX_BUCKET")
}

func (conn *InfluxConn) writeSensorData(reading BME280) {
	p := influxdb2.NewPointWithMeasurement("bme280").
		AddTag("feeder", "MQTT-Influx-Connector").
		AddTag("loc", reading.Loc).
		AddTag("name", reading.Name).
		AddField("temperature", reading.Temperature).
		AddField("humidity", reading.Humidity).
		AddField("pressure", reading.Pressure).
		AddField("rssi", reading.RSSI).
		SetTime(time.Unix(int64(reading.Timestamp), 0))

	write := conn.client.WriteAPIBlocking(getOrganization(), getBucket())
	write.WritePoint(context.Background(), p)

}
