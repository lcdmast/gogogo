package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern")
	var commond = "walk outside"
	var exit = strings.Contains(commond, "outside")
	fmt.Println("You leave the cave:", exit)

	fmt.Println("The year is 2100, should you leap?")

	var year = 2021
	var leap = (year%400 == 0) || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground")
	}

	fmt.Println("Tghereis acavern entrance here and apath to the east.")
	var command = "go inside"
	switch command {
	case "go east":
		fmt.Println("You head further up the mountain.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors' .")
	default:
		fmt.Println("Didnt quite get that.")
	}

	var room = "lake"
	switch {
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}

	switch room {
	case "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case "underwater":
		fmt.Println("The water is freeing cold.")
	}

	// var count = 10
	// for count > 0 {
	// 	fmt.Println(count)
	// 	time.Sleep(time.Second)
	// 	count--
	// }
	// fmt.Println("LeftOff")

	// var degrees = 0
	// for {
	// 	fmt.Println(degrees)
	// 	degrees++
	// 	if degrees >= 360 {
	// 		degrees = 0
	// 		if rand.Intn(2) == 0 {
	// 			break
	// 		}
	// 	}
	// }

	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("LeftOff!")
	} else {
		fmt.Println("Launch failed.")
	}

	var lcd = 12

	for {
		randNum := rand.Intn(100)
		fmt.Printf("随机数%d\n", randNum)
		if randNum == lcd {
			fmt.Println("相等")
			break
		}
		if randNum > lcd {
			fmt.Println("大于lcd")
		}
		if randNum < lcd {
			fmt.Println("小于lcd")
		}
	}
}
