Solution-1

package piscine

func LastWord(s string) string {
	start := -1
	end := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			if end == -1 {
				end = i + 1
			}
			start = i
		} else if end != -1 {
			break
		}
	}
	if start == -1 {
		return "\n"
	}
	return s[start:end] + "\n"
}

Solution-2

package piscine

func LastWord(s string) string {
	runes := []rune(s)
	end := len(runes) - 1

	// Step 1: skip trailing spaces
	for end >= 0 && runes[end] == ' ' {
		end--
	}

	if end < 0 {
		return "\n"
	}

	// Step 2: find the start of the last word
	start := end
	for start >= 0 && runes[start] != ' ' {
		start--
	}

	// Step 3: extract the word
	word := string(runes[start+1 : end+1])
	return word + "\n"
}
