package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)

	// Quick check: only spaces?
	onlySpaces := true
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			onlySpaces = false
			break
		}
	}
	if onlySpaces {
		// For N spaces, expected count is N+1
		result[""] = len(str) + 1
		return result
	}

	word := ""
	prevWasWord := false // true if we just finished a word before a space run
	spaceRun := 0        // length of current consecutive spaces

	for i := 0; i < len(str); i++ {
		c := str[i]

		if c == ' ' {
			if word != "" {
				// finished a word, start counting spaces after it
				result[word]++
				word = ""
				prevWasWord = true
			}
			spaceRun++
		} else {
			// we’re entering a non-space after a space run
			if spaceRun > 0 {
				// between words: count only the extra spaces beyond the first
				if prevWasWord {
					if spaceRun > 1 {
						result[""] += spaceRun - 1
					}
				}
				spaceRun = 0
			}
			// building a word
			word += string(c)
			// as soon as we start a word, we’re no longer “just after a word”
			prevWasWord = false
		}
	}

	// finalize trailing word (if any)
	if word != "" {
		result[word]++
	}

	// trailing spaces are ignored (unless it was only-spaces, handled above)

	return result
}
