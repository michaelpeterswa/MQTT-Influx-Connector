package main

type BME280 struct {
	Name        string  `json:"name"`
	Loc         string  `json:"loc"`
	Timestamp   int     `json:"time"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Pressure    float32 `json:"pressure"`
	RSSI        int     `json:"RSSI"`
}
