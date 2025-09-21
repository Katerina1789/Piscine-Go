package piscine // Defines the package name; "main" means this is an executable program

import (
	"fmt" // Provides functions for formatted I/O (printing to the terminal)
	"io"  // Provides functions to read/write data streams (we use it to read stdin)
	"os"  // Provides OS-level functions (we use it to access stdin and handle errors)
)

func main() {
	// Step 1: Read all input from standard input (stdin)
	// This captures whatever is piped into the program, e.g. ./quadA 3 3 | go run .
	data, err := io.ReadAll(os.Stdin) // Reads all bytes from stdin into 'data'
	if err != nil {                   // If an error occurred while reading...
		return // Exit silently (no output)
	}

	// Step 2: Remove exactly one trailing newline if present
	// This prevents an extra empty line from being processed as part of the input
	if len(data) > 0 && data[len(data)-1] == '\n' { // Check if last byte is '\n'
		data = data[:len(data)-1] // Slice off the last byte
	}

	// Step 3: Split the input into lines
	// Converts the raw bytes into a slice of strings, one string per line
	lines := splitLines(data)

	// Step 4: Validate the input
	// If there are no lines, or the lines have different widths, it's not a valid quad
	if len(lines) == 0 || !sameWidth(lines) {
		fmt.Println("Not a quad function") // Print error message
		return                             // Exit program
	}

	// Step 5: Get the dimensions of the quad
	width := len(lines[0]) // Width = number of characters in the first line
	height := len(lines)   // Height = number of lines

	// Step 6: Check which quads match the input
	// We generate each quad with the same width/height and compare it to the input
	matches := []string{}                   // Slice to store names of matching quads
	if match(lines, quadA(width, height)) { // Compare with generated quadA
		matches = append(matches, "quadA") // If match, add to results
	}
	if match(lines, quadB(width, height)) {
		matches = append(matches, "quadB")
	}
	if match(lines, quadC(width, height)) {
		matches = append(matches, "quadC")
	}
	if match(lines, quadD(width, height)) {
		matches = append(matches, "quadD")
	}
	if match(lines, quadE(width, height)) {
		matches = append(matches, "quadE")
	}

	// Step 7: Print the result
	if len(matches) == 0 { // If no matches found
		fmt.Println("Not a quad function")
	} else {
		// Print all matches in alphabetical order (already in order because of the if sequence)
		for i, name := range matches { // Loop through matches
			if i > 0 {
				fmt.Print(" || ") // Print separator between multiple matches
			}
			// Print in format: [quadX] [width] [height]
			fmt.Printf("[%s] [%d] [%d]", name, width, height)
		}
		fmt.Println() // Always end with a newline
	}
}

// splitLines takes raw byte data and returns a slice of strings, one per line
func splitLines(data []byte) []string {
	var lines []string       // Slice to store lines
	current := ""            // Current line being built
	for _, b := range data { // Loop through each byte
		if b == '\n' { // If newline character found
			lines = append(lines, current) // Add current line to slice
			current = ""                   // Reset current line
		} else {
			current += string(b) // Add character to current line
		}
	}
	if current != "" { // Add last line if not empty
		lines = append(lines, current)
	}
	return lines
}

// sameWidth checks if all lines have the same number of characters
func sameWidth(lines []string) bool {
	w := len(lines[0]) // Expected width from first line
	for _, line := range lines {
		if len(line) != w { // If any line has different width
			return false
		}
	}
	return true
}

// match compares two slices of strings line-by-line for exact equality
func match(a, b []string) bool {
	if len(a) != len(b) { // If different number of lines
		return false
	}
	for i := range a {
		if a[i] != b[i] { // If any line differs
			return false
		}
	}
	return true
}

// quadA generates the ASCII art for quadA with given width (w) and height (h)
func quadA(w, h int) []string {
	var out []string         // Slice to store lines
	for i := 0; i < h; i++ { // Loop through rows
		line := ""               // Build one line at a time
		for j := 0; j < w; j++ { // Loop through columns
			if (i == 0 || i == h-1) && (j == 0 || j == w-1) {
				line += "o" // Corners
			} else if i == 0 || i == h-1 {
				line += "-" // Top/bottom edges
			} else if j == 0 || j == w-1 {
				line += "|" // Left/right edges
			} else {
				line += " " // Inside space
			}
		}
		out = append(out, line) // Add line to output
	}
	return out
}

// quadB generates the ASCII art for quadB
func quadB(w, h int) []string {
	var out []string
	for i := 0; i < h; i++ {
		line := ""
		for j := 0; j < w; j++ {
			if i == 0 && j == 0 {
				line += "/"
			} else if i == 0 && j == w-1 {
				line += "\\"
			} else if i == h-1 && j == 0 {
				line += "\\"
			} else if i == h-1 && j == w-1 {
				line += "/"
			} else if i == 0 || i == h-1 || j == 0 || j == w-1 {
				line += "*"
			} else {
				line += " "
			}
		}
		out = append(out, line)
	}
	return out
}

// quadC generates the ASCII art for quadC
func quadC(w, h int) []string {
	var out []string
	for i := 0; i < h; i++ {
		line := ""
		for j := 0; j < w; j++ {
			if i == 0 && (j == 0 || j == w-1) {
				line += "A"
			} else if i == h-1 && (j == 0 || j == w-1) {
				line += "C"
			} else if i == 0 || i == h-1 || j == 0 || j == w-1 {
				line += "B"
			} else {
				line += " "
			}
		}
		out = append(out, line)
	}
	return out
}

// quadD generates the ASCII art for quadD
func quadD(w, h int) []string {
	var out []string
	for i := 0; i < h; i++ {
		line := ""
		for j := 0; j < w; j++ {
			if (i == 0 || i == h-1) && j == 0 {
				line += "A"
			} else if (i == 0 || i == h-1) && j == w-1 {
				line += "C"
			} else if i == 0 || i == h-1 || j == 0 || j == w-1 {
				line += "B"
			} else {
				line += " "
			}
		}
		out = append(out, line)
	}
	return out
}

// quadE generates the ASCII art for quadE
func quadE(w, h int) []string {
	var out []string
	for i := 0; i < h; i++ {
		line := ""
		for j := 0; j < w; j++ {
			if i == 0 && j == 0 {
				line += "A"
			} else if i == h-1 && j == w-1 {
				line += "A"
			} else if i == 0 && j == w-1 {
				line += "C"
			} else if i == h-1 && j == 0 {
				line += "C"
			} else if i == 0 || i == h-1 || j == 0 || j == w-1 {
				line += "B"
			} else {
				line += " "
			}
		}
		out = append(out, line)
	}
	return out
}
