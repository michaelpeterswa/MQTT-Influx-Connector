package main

type BME280 struct {
	name        string
	loc         string
	timestamp   int
	payload     string
	temperature float32
	humidity    float32
	pressure    float32
	rssi        int
}
