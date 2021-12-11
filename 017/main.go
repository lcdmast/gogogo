package main

import "fmt"

type coordiante struct {
	d, m, s float64
	h       rune
}

func (c coordiante) decimal() float64 {
	sign := 1.0

	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {
	lat := coordiante{4, 35, 22.2, 'W'}
	long := coordiante{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal())
}
