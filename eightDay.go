package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for _, a := range a {
		fmt.Println(a)
	}
	for i, a := range a {
		fmt.Printf("i==%d a===%d\n", i, a)
	}
	var q [3]int = [3]int{1, 2, 3}
	for _, a := range q {
		fmt.Println(a)
	}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])
	f := [...]int{1, 2, 3}
	fmt.Printf("%T\n", f)
	symbol := [...]string{USD: "$", EUR: "¢", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
