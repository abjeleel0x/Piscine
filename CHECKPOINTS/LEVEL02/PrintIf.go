Solution-1

package piscine

func PrintIf(str string) string {
	if len(str) == 0 || len(str) >= 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}

Solution-2

package piscine

func PrintIf(str string) string {
	if str == "" {
		return "G\n"
	}
	if len(str) >= 3 {
		return "G\n"
	} else {
		return "Invalid Input\n"
	}
}
