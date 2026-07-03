package main

import (
	"fmt"
	"os"
	"strings"
)

/*
*
Reads in a list of files given as command-line arguments (or reads from stdin if no arguments are given)

	and reads through each of the files in turn, counting the number of duplicate lines and print those which
	are duplicates to stdout
*/
func main() {
	counts := make(map[string]int)
	filesByLine := make(map[string][]string)
	// files := os.Args[1:]
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if !contains(filesByLine[line], filename) {
				filesByLine[line] = append(filesByLine[line], filename)
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(filesByLine[line], ", "))
		}
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
