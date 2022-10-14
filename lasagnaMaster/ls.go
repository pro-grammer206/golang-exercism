package lasagnaMaster

import (
	"fmt"
	"strconv"
)

// func RemainingOvenTime(actual int) int {
// 	return OvenTime - actual
// }
// func PreparationTime(numberOfLayers int) int {
// 	return numberOfLayers * 2
// }
// func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
// 	return PreparationTime(numberOfLayers) + actualMinutesInOven
// }
func PreparationTime(layers []string, time int)int{
	if time == 0{
		time = 2
	}
	return len(layers)*time
}
func Quantities(layers []string)(noodles int,sauce float64){

	for i:=0;i<len(layers);i++{
		if layers[i]=="sauce"{
			sauce+=0.2
		}
		if layers[i]=="noodles"{
			noodles+=50
		}
	}
	return
}
func AddSecretIngredient(friendList,myList []string){
	myList[len(myList)-1] = friendList[len(friendList)-1]
}
func ScaleRecipe(quantities []float64,numOfPortions int)[]float64{
	scaleQ:=make([]float64,len(quantities))
	scale:=float64(numOfPortions)/2
	var scaledValue string
	for i,v:= range quantities{
		val:=v*scale
		scaledValue=fmt.Sprintf("%.2f",val)
		scaleQ[i],_ = strconv.ParseFloat(scaledValue,64)
	}
	return scaleQ
}