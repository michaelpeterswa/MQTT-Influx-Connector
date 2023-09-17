package mqtt

import (
	"crypto/tls"
	"os"
	"os/signal"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/handlers"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/helpers"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/influx"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/structs"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/timescale"
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

func InitMQTT(iConn *influx.InfluxConn, tConn *timescale.TimescaleConn, settings *structs.MQTTInfluxConnectorSettings) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	connOpts := MQTT.NewClientOptions().AddBroker(settings.MQTTAddress).SetClientID(settings.MQTTClientId).SetCleanSession(true)
	if settings.MQTTUsername != "" {
		connOpts.SetUsername(settings.MQTTUsername)
		if settings.MQTTPassword != "" {
			connOpts.SetPassword(settings.MQTTPassword)
		}
	}

	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)

	topics := topicsToMapStringByte(settings.MQTTTopics)

	connOpts.OnConnect = func(c MQTT.Client) {
		if token := c.SubscribeMultiple(topics, handlers.OnMessageReceived(iConn, tConn)); token.Wait() && token.Error() != nil {
			iConn.Logger.Fatal("couldn't subscribe multiple", zap.Error(token.Error()))
		}
	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		iConn.Logger.Fatal("couldn't connect to mqtt server", zap.Error(token.Error()))
	} else {
		iConn.Logger.Info("connected to mqtt server", zap.String("server", settings.MQTTAddress))
	}

	<-c
}
