package main

import (
	"fmt"
)

func main() {
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.3f\n", third)

	red, green, blue := 0, 141, 213
	fmt.Printf("%x %x %x", red, green, blue)

	fmt.Printf("color: #%02x%02x%02x;", red, green, blue)
}
