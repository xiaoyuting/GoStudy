package main

import (
	"fmt"
)

func main() {
	month := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	Q2 := month[4:7]

	summer := month[6:9]

	fmt.Println(Q2)
	fmt.Println(summer)

	// for _, q := range month {
	// 	fmt.Println(q)

	// }
	for _, q := range Q2 {
		for _, s := range summer {
			if s == q {
				fmt.Printf("相同的的月份=%s\n", s)
			}
		}
	}
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])

	fmt.Println(a)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Println(s[i], s[j])
		s[i], s[j] = s[j], s[i]
		fmt.Println(s[i], s[j])
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	// ...expand z to at least zlen...
	copy(z[len(x):], y)
	return z
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"
}

func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 9 8]
}
