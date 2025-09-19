package piscine

func PodiumPosition(podium [][]string) [][]string {
	i := 0
	for i < len(podium)/2 {
		podium[i], podium[len(podium)-1-i] = podium[len(podium)-1-i], podium[i]
		i++
	}
	return podium
}
