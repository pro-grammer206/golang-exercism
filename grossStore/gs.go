package grossStore

func Units() map[string]int {
	units := make(map[string]int, 0)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}
func NewBill() map[string]int {
	return make(map[string]int, 0)
}
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return false
	}
	if _, ok := bill[item]; ok {
		bill[item] += units[unit]
		return true
	} else {
		bill[item] = units[unit]
		return true
	}

}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return false
	}
	if _, ok := bill[item]; ok {
		newBill := bill[item] - units[unit]
		if newBill < 0 {
			return false
		} else if newBill == 0 {
			delete(bill, item)
			return true
		} else {
			bill[item] = newBill
			return true
		}
	} else {
		return false
	}
}

func GetItem(bill map[string]int, item string) (int, bool) {
	if items, ok := bill[item]; ok {
		return items, true
	} else {
		return 0, false
	}

}
