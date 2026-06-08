package piscine

func RepeatAlpha(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			// Repeat lowercase letter (char - 'a' + 1) times
			repeatCount := int(char - 'a' + 1)
			for i := 0; i < repeatCount; i++ {
				result += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			// Repeat uppercase letter (char - 'A' + 1) times
			repeatCount := int(char - 'A' + 1)
			for i := 0; i < repeatCount; i++ {
				result += string(char)
			}
		} else {
			// Non-alphabetic character → print once
			result += string(char)
		}
	}
	return result
}
