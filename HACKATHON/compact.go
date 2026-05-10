package piscine

func Compact(finger *[]string) int {
	if finger == nil {
		return 0
	}

	s := *finger
	n := 0

	for _, val := range s {
		if val != "" {
			s[n] = val
			n++
		}
	}

	*finger = s[:n]

	return n
}
