package main

import (
	"fmt"

	"github.com/pro-grammer206/gopher_coder/grossStore"
)

const OvenTime = 40

func main() {
	// fmt.Println(bookingUpForBeauty.HasPassed("July 25, 2019 13:45:00"))
bill := map[string]int{"carrot": 12, "grapes": 3}
qty, ok := grossStore.GetItem(bill, "carrot")
fmt.Println(qty)
// Output: 12
fmt.Println(ok)
}

