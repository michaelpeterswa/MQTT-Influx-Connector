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
	Voltage     float32 `json:"voltage,omitempty"`
}

type TSL2561 struct {
	Name      string  `json:"name"`
	Loc       string  `json:"loc"`
	Timestamp int     `json:"time"`
	Lux       float32 `json:"lux"`
	RSSI      int     `json:"RSSI"`
}

type PMSA003I struct {
	Name      string `json:"name"`
	Loc       string `json:"loc"`
	Timestamp int    `json:"time"`
	Pm10S     int    `json:"pm10s"`
	Pm25S     int    `json:"pm25s"`
	Pm100S    int    `json:"pm100s"`
	Pm10E     int    `json:"pm10e"`
	Pm25E     int    `json:"pm25e"`
	Pm100E    int    `json:"pm100e"`
	P03Um     int    `json:"p03um"`
	P05Um     int    `json:"p05um"`
	P10Um     int    `json:"p10um"`
	P25Um     int    `json:"p25um"`
	P50Um     int    `json:"p50um"`
	P100Um    int    `json:"p100um"`
	RSSI      int    `json:"RSSI"`
}
