package main

import (
	"fmt"

	"github.com/pro-grammer206/gopher_coder/needForSpeed"
)

const OvenTime = 40

func main() {
	speed := 5
	batteryDrain := 2
	car:=needForSpeed.NewCar(speed,batteryDrain)
	 distance := 100
 track := needForSpeed.NewTrack(distance)

	fmt.Println(needForSpeed.CanFinish(car, track))

}

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}