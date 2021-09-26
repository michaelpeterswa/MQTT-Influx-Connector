package main

type MQTTInfluxConnectorSettings struct {
	InfluxAddress      string `yaml:"influx-address"`
	InfluxToken        string `yaml:"influx-token"`
	InfluxOrganization string `yaml:"influx-organization"`
	InfluxBucket       string `yaml:"influx-bucket"`
	InfluxMeasurement  string `yaml:"influx-measurement"`

	MQTTAddress  string `yaml:"mqtt-address"`
	MQTTClientId string `yaml:"mqtt-client-id"`
	MQTTUsername string `yaml:"mqtt-username"`
	MQTTPassword string `yaml:"mqtt-password"`

	MQTTTopics []Topic `yaml:"mqtt-topics"`
}

type Topic struct {
	Topic SubTopic `yaml:"topic"`
	QoS   byte     `yaml:"qos"`
}

type SubTopic struct {
	Type     string `yaml:"type"`
	Location string `yaml:"location"`
	Room     string `yaml:"room"`
	Name     string `yaml:"name"`
	Field    string `yaml:"field"`
}

type BME280 struct {
	Name        string  `json:"name"`
	Loc         string  `json:"loc"`
	Timestamp   int     `json:"time"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Pressure    float32 `json:"pressure"`
	RSSI        int     `json:"RSSI"`
}
