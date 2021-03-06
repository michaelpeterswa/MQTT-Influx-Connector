package mqtt

import (
	"crypto/tls"
	"flag"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/handlers"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/helpers"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/influx"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/structs"
	"go.uber.org/zap"
)

type MqttConnection struct {
	MQTTClient MQTT.Client
}

func topicsToMapStringByte(t []structs.Topic) map[string]byte {
	m := make(map[string]byte)
	for _, topic := range t {
		m[helpers.BuildTopic(topic.Topic)] = topic.QoS
	}
	return m
}

func InitMQTT(iConn *influx.InfluxConn, settings *structs.MQTTInfluxConnectorSettings) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	hostname, _ := os.Hostname()

	server := flag.String("server", settings.MQTTAddress, "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")
	clientid := flag.String("clientid", hostname+strconv.Itoa(time.Now().Second()), "A clientid for the connection")
	username := flag.String("username", settings.MQTTPassword, "A username to authenticate to the MQTT server")
	password := flag.String("password", settings.MQTTPassword, "Password to match username")
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

	topics := topicsToMapStringByte(settings.MQTTTopics)

	connOpts.OnConnect = func(c MQTT.Client) {
		if token := c.SubscribeMultiple(topics, handlers.OnMessageReceived(iConn)); token.Wait() && token.Error() != nil {
			iConn.Logger.Fatal("couldn't subscribe multiple", zap.Error(token.Error()))
		}
	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		iConn.Logger.Fatal("couldn't connect to mqtt server", zap.Error(token.Error()))
	} else {
		iConn.Logger.Info("connected to mqtt server", zap.String("server", *server))
	}

	<-c
}
