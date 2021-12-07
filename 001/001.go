package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Printf("my name is %v age %v", "lcd", "18")
	var command = "li chang dong"
	var exit = strings.Contains(command, "li")
	fmt.Println("you leave", exit)

	var room = "lake"
	switch {
	case room == "cave":
		fmt.Println("1111")
	case room == "lake":
		fmt.Println("222")
		// fallthrough
	case room == "lake3":
		fmt.Println("33")
	}
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}

	if num := rand.Intn(10); num == 0 {
		fmt.Println("11")
	} else if num == 1 {
		fmt.Println(num)
	} else {
		fmt.Println("33")
	}

	fmt.Println(rand.Intn(5))
}
