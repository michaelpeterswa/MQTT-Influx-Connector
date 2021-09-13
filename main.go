package main

import (
	"encoding/json"
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func onMessageReceived(influxConn *InfluxConn) func(client MQTT.Client, message MQTT.Message) {
	return func(client MQTT.Client, message MQTT.Message) {
		fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())

		var reading BME280
		err := json.Unmarshal(message.Payload(), &reading)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(reading)
			// writeSensorData(writeAPI, reading)
			influxConn.writeSensorData(reading)
		}
	}
}

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	influxConn := initInflux()
	initMQTT(influxConn)
}
