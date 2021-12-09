package main

import (
	"fmt"
	"math"
)

func main() {
	lcd1 := []float64{
		23.0, 28.0, 34.0, 23.0, 14.0, 23.0, 14.0, 34.0,
	}

	lcd2 := make(map[float64]int)

	for _, v := range lcd1 {
		lcd2[v]++
	}
	fmt.Println(lcd1, lcd2)

	groups := make(map[float64][]float64)

	for _, l := range lcd1 {
		g := math.Trunc(l/10) * 10
		groups[g] = append(groups[g], l)
	}
	fmt.Println(groups)

	cd := make([]float64, 0)
	cd = append(cd, 5.0)
	fmt.Println(cd[0:1])

}
