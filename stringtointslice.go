package piscine

func StringToIntSlice(str string) []int {
	var list []int

	for _, ch := range str {
		list = append(list, int(ch))
	}
	return list
}
