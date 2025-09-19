package piscine

func PodiumPosition(podium [][]string) [][]string {
	var result [][]string
	i := 0
	for i < len(podium) {
		result = result[:len(result)+1] // manually extend the slice
		result[len(result)-1] = podium[len(podium)-1-i]
		i++
	}
	return result
}
