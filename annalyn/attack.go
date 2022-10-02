package annalyn

func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		return true
	} else {
		return false
	}
}
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if !archerIsAwake && prisonerIsAwake {
		return true
	} else {
		return false
	}
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent && !archerIsAwake {
		return true
	} else if !petDogIsPresent && prisonerIsAwake && !knightIsAwake && !archerIsAwake {
		return true
	} else {
		return false
	}
}