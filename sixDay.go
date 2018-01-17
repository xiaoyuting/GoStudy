package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	b := true
	if b {
		fmt.Println("yes")
	} else {

		fmt.Println("no")
	}
	s := "hello ,世界"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])
	fmt.Println("goodbye" + s[5:])
	t := s
	//s += "hhahah"
	fmt.Println(s)
	fmt.Println(t)
	//s[0] = "text"
	fmt.Println(s[:])

	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	n := 0
	for _, _ = range s {
		n++
		fmt.Println(n)
	}
	m := 0
	for range s {
		m++
		fmt.Println(m)

	}
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"

	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"

	h := fmt.Sprintf("x=%b", x) // "x=1111011"
	fmt.Println(h)
	//x, err := strconv.Atoi("123")             // x is an int
	//y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits

}
func itod(b int) bool {

	return b != 0

}
func HasPerfix(s, perfix string) bool {
	return len(s) >= len(perfix) && s[:len(perfix)] == perfix
}
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// func Contains(s, substr string) bool {
// 	for i := 0; i < len(s); i++ {
// 		if HasPrefix(s[i:], substr) {
// 			return true
// 		}
// 	}
// 	return false
// }

// 为了避免转换中不必要的内存分配，bytes包和strings同时提供了许多实用函数。下面是strings包中的六个函数：

// func Contains(s, substr string) bool
// func Count(s, sep string) int
// func Fields(s string) []string
// func HasPrefix(s, prefix string) bool
// func Index(s, sep string) int
// func Join(a []string, sep string) string
// bytes包中也对应的六个函数：

// func Contains(b, subslice []byte) bool
// func Count(s, sep []byte) int
// func Fields(s []byte) [][]byte
// func HasPrefix(s, prefix []byte) bool
// func Index(s, sep []byte) int
// func Join(s [][]byte, sep []byte) []byte
