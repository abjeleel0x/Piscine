package piscine

func StringToIntSlice(str string) []int {
	var omotunde []int

	for _, abduljeleel := range str {
		omotunde = append(omotunde, int(abduljeleel))
	}

	return omotunde
}
