package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/influx"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/mqtt"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/structs"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/timescale"
	yaml "gopkg.in/yaml.v2"
)

var settings *structs.MQTTInfluxConnectorSettings

func main() {
	ctx := context.Background()

	yamlFile, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("open file error: %v\n", err)
	}

	yamlContents, err := io.ReadAll(yamlFile)
	if err != nil {
		log.Fatalf("read file error: %v\n", err)
	}

	err = yaml.Unmarshal([]byte(yamlContents), &settings)
	if err != nil {
		log.Fatalf("yaml unmarshal error: %v\n", err)
	}

	influxConn := influx.InitInflux(settings)
	timescaleConn, err := timescale.InitTimescale(ctx, settings)
	if err != nil {
		log.Fatalf("timescale init error: %v\n", err)
	}
	defer timescaleConn.Close(ctx)

	mqtt.InitMQTT(influxConn, timescaleConn, settings)
}
