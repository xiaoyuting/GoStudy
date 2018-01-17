package main

import (
	"fmt"
)

//const boilingF = 212.0

func main() {
	// var f = boilingF
	// var c = (f - 32) * 5 / 9

	const freezingF, bolingF = 32.0, 212.0
	fmt.Println("bolingF point=%g For %g", fToc(freezingF), fToc(bolingF))
	fmt.Print("11111")
}
func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
