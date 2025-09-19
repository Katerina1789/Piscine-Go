package piscine

func LoafOfBread(str string) string {
	rs := []rune(str)
	var segments []string
	chunk := make([]rune, 0, 5)
	nonSpaceTotal := 0

	i := 0
	for i < len(rs) {
		r := rs[i]

		// Count only non-space characters, do not include spaces in output
		if r != ' ' {
			chunk = append(chunk, r)
			nonSpaceTotal++

			// When we reach 5 non-space chars, emit the chunk and skip next character (any)
			if len(chunk) == 5 {
				segments = append(segments, string(chunk))
				chunk = chunk[:0]
				if i+1 < len(rs) {
					i++ // skip the next character in the original stream (space or not)
				}
			}
		}
		i++
	}

	if nonSpaceTotal < 5 {
		return "Invalid Output\n"
	}
	if len(chunk) > 0 {
		segments = append(segments, string(chunk))
	}

	// Join chunks with a single space and add newline
	out := ""
	for idx, s := range segments {
		if idx > 0 {
			out += " "
		}
		out += s
	}
	return out + "\n"
}
