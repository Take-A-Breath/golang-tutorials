package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightspeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)

	seconds := new(big.Int)
	seconds.Div(distance, lightspeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	fmt.Println("That is", days, "days of travel at light speed.")
}
