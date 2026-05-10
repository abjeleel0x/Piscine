package piscine

func JumpOver(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}
	rel := ""
	for i := 2; i < len(str); i += 3 {
		if str[i] != ' ' {
			rel += string(str[i])
		}
	}
	return rel + "\n"
}
