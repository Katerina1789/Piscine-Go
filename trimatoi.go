package piscine

func TrimAtoi(s string) int {
	sign := 1
	number := 0
	foundDigit := true

	for i := 0; i < len(s); i++ {
		if s[i] == '-' && foundDigit {
			sign = -1
		}
		if s[i] >= '0' && s[i] <= '9' {
			foundDigit = false
			number = number*10 + int(s[i]-'0')
		}
	}
	return number * sign
}
