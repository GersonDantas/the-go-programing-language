package main

import (
	"fmt"

	"github.com/GersonDantas/gopl/ch2/tempconv"
)

func main() {
	fmt.Println(tempconv.CToF(tempconv.FreezingC))
	fmt.Println(tempconv.KToC(tempconv.WinterK))
	fmt.Println(tempconv.CToK(tempconv.FreezingC))
}
