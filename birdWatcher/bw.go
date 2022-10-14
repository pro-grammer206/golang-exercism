package birdWatcher

import "fmt"

func TotalBirdCount(birdsPerDay []int) int {
	var sum int
	for i := 0; i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	from := (week-1)*7 
	to:= from +6
	fmt.Println(birdsPerDay[from : to])
	return TotalBirdCount(birdsPerDay[from : to+1])
}
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i=i+2 {
		birdsPerDay[i] = birdsPerDay[i] + 1
	}
	return birdsPerDay
}