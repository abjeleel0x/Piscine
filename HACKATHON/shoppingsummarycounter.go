package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	AbuApon := make(map[string]int)
	var a string
	for _, AbuSixteen := range str {
		if AbuSixteen == 32 {
			AbuApon[a] += 1
			a = ""
		} else if AbuSixteen != 32 {
			a += string(byte(AbuSixteen))
		}
	}
	AbuApon[a] += 1

	return AbuApon
}
