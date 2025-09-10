package piscine

func TrimAtoi(s string) int {
	sign := 1
	number := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			sign = -1
			return sign
		}
		if s[i] >= '0' && s[i] <= '9' {
			number = number*10 + int(s[i]-'0')
		}
	}
	return number
}
