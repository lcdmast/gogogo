package main

import (
	"fmt"
	"strings"
)

func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	//声明切片
	lcds := [...]string{
		"lcd1",
		"lcd2",
		"lcd3",
		"lcd4",
		"lcd5",
		"lcd6",
		"lcd7",
	}

	lcds1 := lcds[0:4]
	lcds2 := lcds[3:4]
	lcds3 := lcds[:]
	fmt.Println(lcds1, lcds2, lcds3)

	lcdd := "lichangdong"
	lcdd1 := lcdd[:4]
	fmt.Println(lcdd1)

	planets := []string{"li   ", "chang   ", "dong"}
	hyperspace(planets)
	fmt.Println(strings.Join(planets, ""))

	fmt.Printf("[%q]", strings.Trim("!!Achtung! Achtung! !!!", " !"))
}
