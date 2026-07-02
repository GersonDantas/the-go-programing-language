package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("os.Args[0]:", os.Args[0])

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("tempo de execucao:", time.Since(start))
}
