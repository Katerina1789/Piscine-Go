package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 3; i++ {
		for j := i; j <= len(deck); j++ {
			if j != i {
				fmt.Printf(", ")
			}
			fmt.Printf("%d", deck[j])
		}
		fmt.Println()
	}
}
