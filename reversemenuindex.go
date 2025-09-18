package piscine

func ReverseMenuIndex(menu []string) []string {
	order := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		order[i] = menu[len(menu)-1-i]
	}
	return order
}
