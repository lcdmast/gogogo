package main

import "fmt"

type Worker interface {
	Work()
}
type Person struct {
	name string
}

type Other struct {
	price int
}

func (o Other) Work() {
	fmt.Println(o.price, " money")
}

func (p Person) Work() {
	fmt.Println(p.name, " name is working ")
}

func describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func main() {
	p := Person{
		"lichangdong",
	}
	var w Worker
	w = p
	describe(w)
	w.Work()
	p2 := Other{
		1200,
	}
	w = p2
	describe(w)
	w.Work()
}
