package piscine

func LoafOfBread(str string) string {
	result := ""
	word := ""
	count := 0
	skip := false

	for _, r := range str {
		if r == ' ' {
			// Preserve spaces in output
			word += string(r)
			continue
		}
		if skip {
			skip = false
			continue // skip this non-space character
		}
		word += string(r)
		count++
		if count == 5 {
			result += word + " "
			word = ""
			count = 0
			skip = true // skip next non-space character
		}
	}

	if result == "" {
		return "Invalid Output\n"
	}
	return result[:len(result)-1] + "\n" // remove trailing space and add newline
}
