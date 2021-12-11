package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Type = %T ,Value=%v\n", i, i)
}

func main() {
	s := "Hello word"
	describe(s)
	i := 5
	describe(i)

	strt := struct {
		name string
	}{
		name: "Naveen R",
	}

	describe(strt)
}
