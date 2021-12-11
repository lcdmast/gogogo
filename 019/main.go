package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) detail() {
	fmt.Println(p.title)
	fmt.Println(p.content)
	fmt.Println(p.author.fullName())
	fmt.Println(p.author.bio)
}

func main() {
	a := author{
		"li",
		"changdong",
		"Golang",
	}

	c := a.fullName()
	fmt.Println(c)

	b := post{
		"title",
		"content",
		a,
	}
	b.detail()
}
