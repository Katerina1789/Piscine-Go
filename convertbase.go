package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {

	decimal := 0
	for i := 0; i < len(nbr); i++ {
		for j := 0; j < len(baseFrom); j++ {
			if nbr[i] == baseFrom[j] {
				decimal = decimal*len(baseFrom) + j
				break
			}
		}
	}

	if decimal == 0 {
		return string(baseTo[0])
	}

	result := ""
	for decimal > 0 {
		remainder := decimal % len(baseTo)
		result = string(baseTo[remainder]) + result
		decimal = decimal / len(baseTo)
	}

	return result
}
