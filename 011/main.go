package main

import (
	"fmt"
	"sort"
)

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	lcd := []string{
		"Me", "Ad", "Cc", "Bd",
		"Kd", "Lj", "Tt", "Cr",
	}

	lcd = append(lcd, "nihao")
	// 按照字母的大小排序
	sort.StringSlice(lcd).Sort()

	fmt.Println(lcd)

	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[2:2])

	lcd2 := lcd[0:4]
	lcd3 := append(lcd2, "dj")
	dump("lcd2", lcd2)
	dump("lcd3", lcd3)

	yu := make([]string, 0, 10)
	yu = append(yu, "nice1", "nice2", "nice3")
	dump("yu", yu)
}
