package piscine

func PodiumPosition(podium [][]string) [][]string {
	nbr := len(podium)
	result := make([][]string, nbr)
	for i := 0; i < nbr; i++ {
		result[i] = podium[nbr-1-i]
	}
	return result
}
