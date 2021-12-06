package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("my name is %v age %v", "lcd", "18")
	var command = "li chang dong"
	var exit = strings.Contains(command, "li")
	fmt.Println("you leave", exit)
}
