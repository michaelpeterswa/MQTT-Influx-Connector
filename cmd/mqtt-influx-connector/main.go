package main

import (
	"io/ioutil"
	"log"

	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/influx"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/mqtt"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/structs"
	yaml "gopkg.in/yaml.v2"
)

var settings *structs.MQTTInfluxConnectorSettings

func main() {

	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("read config error: %v\n", err)
	}

	err = yaml.Unmarshal([]byte(yamlFile), &settings)
	if err != nil {
		log.Fatalf("unmarshal error: %v\n", err)
	}

	influxConn := influx.InitInflux(settings)
	mqtt.InitMQTT(influxConn, settings)
}
