package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	slash := 0
	for i := 0; i < len(name); i++ {
		if name[i] == '/' {
			slash = i + 1
		}
	}
	for _, ch := range name[slash:] {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
