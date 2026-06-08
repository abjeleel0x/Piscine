Solution-1

package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	
	if !isValidCamelCase(s) {
		return s
	}
	
	result := ""
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(char)
		} else {
			result += string(char)
		}
	}
	
	return result
}

func isValidCamelCase(s string) bool {
	for i, char := range s {
		if char >= '0' && char <= '9' {
			return false
		}
		
		if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
			return false
		}
		
		if char >= 'A' && char <= 'Z' {
			if i == len(s)-1 {
				return false
			}
			
			if i > 0 {
				prevChar := rune(s[i-1])
				if prevChar >= 'A' && prevChar <= 'Z' {
					return false
				}
			}
		}
	}
	
	return true
}

Soution-2

package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	x := func(c byte) bool {
		return c >= 'a' && c <= 'z'
	}
	y := func(c byte) bool {
		return c >= 'A' && c <= 'Z'
	}

	if s[0] < 'A' || s[0] > 'Z' && s[0] < 'a' || s[0] > 'z' {
		return s
	}
	for i := 0; i < len(s); i++ {
		if !y(s[i]) && !x(s[i]) {
			return s
		}
	}

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 && y(s[i]) {
			return s
		}
		if y(s[i]) && y(s[i+1]) {
			return s
		}
	}

	var a []byte

	for i := 0; i < len(s); i++ {
		c := s[i]
		if y(c) {
			if i > 0 {
				a = append(a, '_')
			}
		}
		a = append(a, c)
	}
	return string(a)
}
