package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	items := []string{}
	buf := make([]byte, 0, len(str))

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			items = append(items, string(buf))
			buf = buf[:0]
		} else {
			buf = append(buf, str[i])
		}
	}
	items = append(items, string(buf))

	m := make(map[string]int)
	for _, w := range items {
		m[w]++
	}
	return m
}
