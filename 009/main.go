// 我的减重程序
package main

import "fmt"

// main是所有函数的起始程序
func main() {
	var lcd [8]string
	lcd[0] = "lichangdong"
	lcd[1] = "lichandong2"
	lcd[2] = "lichangdong3"

	eart := lcd[2]

	fmt.Println(eart, lcd[4])

	planets := [...]string{"lcd", "lcd2", "lcd3", "lcd4", "lcd5", "lcd6"}
	fmt.Println(planets)
	for i := 0; i < len(planets); i++ {
		fmt.Println(i, planets[i])
	}
	for i, v := range planets {
		fmt.Println(i, v)
	}

	lcc := planets

	lcc[2] = "lcc"
	fmt.Println(planets, lcc)
}
