package main

import (
	"fmt"
	"time"
)

func main() {
	var house = "Malibu point 10880, 90265"
	//对字符串取地址 ptr 类型为*string
	ptr := &house
	fmt.Printf("ptr Type : %T\n", ptr)
	//打印指针ptr地址
	fmt.Printf("address : %p\n", ptr)
	//对指针进行取值操作
	values := *ptr
	//取值类型 string
	fmt.Printf("value type: %T\n", values)
	//指针取值后就是指向变量的值
	fmt.Printf("Value: %s\n", values)

	x, y := 5, 10
	swap(&x, &y)
	fmt.Printf("%d, %d", x, y)

	go running()
	go running2()

	var input string

	fmt.Scanln(&input)
}
func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func running() {
	var i int
	for {
		i++
		fmt.Println("i", i)
		time.Sleep(time.Second)
	}
}
func running2() {
	var j int
	for {
		j++
		fmt.Println("j", j)
		time.Sleep(time.Second)
	}
}
