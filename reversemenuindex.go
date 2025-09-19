package piscine

func ReverseMenuIndex(menu []string) []string {
	order := make([]string, len(menu)) // adds to var order the slices of the menu and the length
	for i := 0; i < len(menu); i++ {   // forloop where i gets added one until it's smaller than the length of the menu
		order[i] = menu[len(menu)-1-i] //
	}
	return order
}
