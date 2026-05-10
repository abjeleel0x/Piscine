package piscine

func ShoppingListSort(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}

	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if len(slice[j]) > len(slice[j+1]) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}
