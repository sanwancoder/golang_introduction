package main

import (
	"./tempconv"
	"fmt"
)

func main() {
	fmt.Println("===============")
	fmt.Println(tempconv.FToC(100))
	fmt.Println(tempconv.CToF(100))

}