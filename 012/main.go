package main

import (
	"fmt"
)

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	twoWorlds := terraform("New", "jay", "jj", "vae")
	fmt.Println(twoWorlds)

	planets := []string{"li", "zhao", "qian"}

	twoPlanets := terraform("xiao", planets...)

	fmt.Println(twoPlanets)
	// fmt.Print(math.Inf(-5))

	temperature := map[string]int{
		"A": 1,
		"B": 2,
	}
	temp := temperature["A"]
	fmt.Println(temp)

	lcd1 := make(map[string]int, 10)
	lcd1["li"] = 90
	lcd1["chang"] = 95

	fmt.Println(lcd1)
	delete(lcd1, "li")
	fmt.Println(lcd1)
	for _, v := range lcd1 {
		fmt.Println(v)
	}
}
