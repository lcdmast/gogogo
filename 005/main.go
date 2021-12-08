package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	message := "uv vagreangvbany fcnpr fgbgvba"

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	s := "abc"

	for k, i := range s {
		fmt.Printf("%d %d\n", k, i)
	}
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	for i, v := range a {
		fmt.Println(i, v)
	}

	var bh float64 = 32768
	var h = int16(bh)
	fmt.Println(h)

	if math.MaxInt16 > h && math.MinInt16 < h {
		fmt.Println(h)
	}

	v := 42
	if v >= 0 && v <= math.MaxUint8 {
		v8 := uint(v)
		fmt.Println("converted ", v8)
	}

	fmt.Println(string(65))

	countdown := 10
	fmt.Println("Launch in minus " + strconv.Itoa(countdown) + " sec0nds ")

	countdown2, err := strconv.Atoi("A")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(countdown2)

	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(celsius)

}
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
