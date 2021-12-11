package main

import "fmt"

type Person struct {
	age int
}

func (p *Person) bir() {
	p.age++
}

func main() {
	var nobody *Person
	fmt.Println(nobody)
	if nobody != nil {
		nobody.bir()
	}
}
