package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	burger := food{}
	chips := food{}
	nuggets := food{}
	burger.preptime = 15
	chips.preptime = 10
	nuggets.preptime = 12
	if order == "burger" {
		return burger.preptime
	} else if order == "chips" {
		return chips.preptime
	} else if order == "nuggets" {
		return nuggets.preptime
	} else {
		return 404
	}
}
