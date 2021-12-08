package main

import (
	"fmt"
)

func main() {
	type celsius float64

	const degress = 20

	var temp celsius = degress
	temp += 10
	fmt.Println(temp)
}
