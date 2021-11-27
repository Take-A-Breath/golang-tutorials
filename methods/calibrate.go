// Using first-class functions, calibrate accepts a fake or real sensor as a parameter and returns a replacement function.
// Whenever the new sensor function is called, the original function is invoked, and the reading is adjusted by an offset.
package main

import "fmt"

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
}
