package piscine

func PodiumPosition(podium [][]string) [][]string {
	var result [][]string
	i := 0
	for i < len(podium) {
		result = append(result, podium[len(podium)-1-i])
		i++
	}
	return result
}
