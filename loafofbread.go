package piscine

func LoafOfBread(str string) string {
	count := 0
	word := ""
	result := ""

	for _, r := range str {
		if r == ' ' {
			continue // ignore spaces when counting
		}
		if count == 5 {
			count = 0 // reset count after 5 characters
			continue  // skip the 6th character
		}
		word += string(r)
		count++
		if count == 5 {
			result += word + " "
			word = ""
		}
	}

	if result == "" {
		return "Invalid Output\n"
	}
	return result[:len(result)-1] + "\n" // remove trailing space and add newline
}
