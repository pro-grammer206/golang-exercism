package carsAssembler

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsperHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(carsperHour / 60)
}

func CalculateCost(numberOfCars int) uint {
	costOfTenCars := 95000
	costOfOneCar := 10000
	return uint(numberOfCars/10)*uint(costOfTenCars) + uint(numberOfCars%10)*uint(costOfOneCar)

}