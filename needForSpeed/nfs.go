package needForSpeed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}
type Track struct {
	distance int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}
func Drive(car Car) Car {
	if car.batteryDrain > car.battery {
		return car
	} else {
		return Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
			battery:      car.battery - car.batteryDrain,
			distance:     car.distance + car.speed,
		}
	}
}
func CanFinish(car Car, track Track) bool {
	times := car.battery / car.batteryDrain
	maxDistance := car.speed * times
	if maxDistance >= track.distance {
		return true
	} else {
		return false
	}

}