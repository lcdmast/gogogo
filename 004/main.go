package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	fmt.Println(lightSpeed, secondsPerDay)
	fmt.Printf("%T\n", lightSpeed)
	fmt.Printf("%T", secondsPerDay)
	distance := new(big.Int)
	distance.SetString("24000000000000000000000000000000", 10)

	fmt.Println(distance)
	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)
	fmt.Println(seconds)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println(days)
	const num = 24000000
	fmt.Println(num)

	fmt.Println(`\n 你好`)
	fmt.Println("Nihao \nhello ")

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang rune = 33
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)

	fmt.Printf("%c %c %c %c", pi, alpha, omega, bang)
}
