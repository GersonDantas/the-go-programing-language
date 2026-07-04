package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
*
reads lines in from standard output (until Ctrl+D) and stores the lines as
keys in a map where the values are the number of times this line has been inputted.
At the end it prints the count of all lines which are duplicates
*/
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%x\t%s\n", n, line)
		}
	}
}
