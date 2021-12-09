package main

import "fmt"

func main() {
	mapSlice := make([]map[string]int, 10)
	for i, v := range mapSlice {
		fmt.Println(i, v)
	}

	var lcd1 int16 = 10
	var lcd2 = float64(lcd1)
	fmt.Println(lcd2)
	fmt.Printf("%T %T\n", lcd2, lcd1)
	// fmt.Println(mapSlice)

	//map不会被复制
	lcdMap := map[string]string{
		"lcd":  "l",
		"lcd2": "c",
		"lcd3": "d",
	}

	lcdMap2 := lcdMap

	lcdMap["lcd"] = "L"
	delete(lcdMap, "lcd")
	fmt.Println(lcdMap2)
}
