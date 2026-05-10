package piscine

func Unmatch(a []int) int {
	var b int
	for _, i := range a {
		b = 0
		for _, j := range a {
			if j == i {
				b++
			}
		}
		if b%2 != 0 {
			return i
		}
	}
	return -1
}
