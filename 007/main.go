package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K \n", k)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func main() {
	// sensor := fakeSensor
	// fmt.Println(sensor())

	// sensor = realSensor
	// fmt.Println(sensor())

	// fmt.Println(rand.Intn(151))
	measureTemperature(3, fakeSensor)
}
