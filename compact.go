package piscine

func Compact(ptr *[]string) int {
	var result []string
	for _, s := range *ptr {
		if s != "" {
			result = append(result, s)
		}
	}
	*ptr = result
	return len(result)
}
