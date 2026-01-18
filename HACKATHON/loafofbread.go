package piscine

func LoafOfBread(str string) string {
	rs := []rune(str)
	var segments []string
	chunk := make([]rune, 0, 5)
	nonSpaceTotal := 0

	i := 0
	for i < len(rs) {
		r := rs[i]

		if r != ' ' {
			chunk = append(chunk, r)
			nonSpaceTotal++

			if len(chunk) == 5 {
				segments = append(segments, string(chunk))
				chunk = chunk[:0]
				if i+1 < len(rs) {
					i++ // skip next character
				}
			}
		}
		i++
	}

	if nonSpaceTotal == 0 {
		return "\n"
	}
	if nonSpaceTotal < 5 {
		return "Invalid Output\n"
	}
	if len(chunk) > 0 {
		segments = append(segments, string(chunk))
	}

	out := ""
	for idx, s := range segments {
		if idx > 0 {
			out += " "
		}
		out += s
	}
	return out + "\n"
}
