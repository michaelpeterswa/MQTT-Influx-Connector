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

	// support for the optional voltage reading... only if populated
	if reading.Voltage != 0 {
		p.AddField("voltage", reading.Voltage)
	}

	write := conn.client.WriteAPIBlocking(getOrganization(), getBucket())
	write.WritePoint(context.Background(), p)

}

func (conn *InfluxConn) writeTSL2561SensorData(reading TSL2561, st SubTopic) {
	p := influxdb2.NewPointWithMeasurement(settings.InfluxMeasurement).
		AddTag("feeder", "MQTT-Influx-Connector").
		AddTag("type", st.Type).
		AddTag("location", st.Location).
		AddTag("room", st.Room).
		AddTag("name", st.Name).
		AddTag("field", st.Field).
		AddField("lux", reading.Lux).
		AddField("rssi", reading.RSSI).
		SetTime(time.Unix(int64(reading.Timestamp), 0))

	write := conn.client.WriteAPIBlocking(getOrganization(), getBucket())
	write.WritePoint(context.Background(), p)

}

func (conn *InfluxConn) writePMSA003ISensorData(reading PMSA003I, st SubTopic) {
	p := influxdb2.NewPointWithMeasurement(settings.InfluxMeasurement).
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

	write := conn.client.WriteAPIBlocking(getOrganization(), getBucket())
	write.WritePoint(context.Background(), p)

}
