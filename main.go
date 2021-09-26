package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	yaml "gopkg.in/yaml.v2"
)

var settings MQTTInfluxConnectorSettings

func onMessageReceived(influxConn *InfluxConn) func(client MQTT.Client, message MQTT.Message) {
	return func(client MQTT.Client, message MQTT.Message) {
		fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
		st := getSubTopicFromString(message.Topic())

		switch v := st.Name; v {
		case "bme280":
			var reading BME280
			err := json.Unmarshal(message.Payload(), &reading)
			if err != nil {
				log.Println(err)
			} else {
				influxConn.writeBME280SensorData(reading, st)
			}
		default:
			fmt.Println("unknown")
		}
	}
}

func main() {

	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Printf(err.Error())
	}

	err = yaml.Unmarshal([]byte(yamlFile), &settings)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	influxConn := initInflux()
	initMQTT(influxConn)
}
