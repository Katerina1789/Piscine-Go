package piscine

import "os"

func IsNegative(nb int) {
	if nb >= 0 {
		os.Stdout.Write([]byte("F\n"))
	} else {
		os.Stdout.Write([]byte("T\n"))
	}
}
