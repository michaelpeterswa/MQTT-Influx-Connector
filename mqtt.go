package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MqttConnection struct {
	MQTTClient MQTT.Client
}

func initMQTT(iConn *InfluxConn) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	hostname, _ := os.Hostname()

	server := flag.String("server", os.Getenv("MQTT_ADDRESS"), "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")
	topic := flag.String("topic", "sensors/home/bedroom/bme280/payload", "Topic to subscribe to")
	qos := flag.Int("qos", 0, "The QoS to subscribe to messages at")
	clientid := flag.String("clientid", hostname+strconv.Itoa(time.Now().Second()), "A clientid for the connection")
	username := flag.String("username", os.Getenv("MQTT_USERNAME"), "A username to authenticate to the MQTT server")
	password := flag.String("password", os.Getenv("MQTT_PASSWORD"), "Password to match username")
	flag.Parse()

	connOpts := MQTT.NewClientOptions().AddBroker(*server).SetClientID(*clientid).SetCleanSession(true)
	if *username != "" {
		connOpts.SetUsername(*username)
		if *password != "" {
			connOpts.SetPassword(*password)
		}
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)

	connOpts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(*topic, byte(*qos), onMessageReceived(iConn)); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", *server)
	}

	<-c
}
