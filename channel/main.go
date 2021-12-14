package main

import (
	"fmt"
)

// func greet(ch chan int) {
// 	data := <-ch

// 	fmt.Println("hello ", data)
// }

func main() {
	// fmt.Println("start")
	// ch := make(chan int)
	// go greet(ch)
	// ch <- 1
	// fmt.Println("ll")
	// n := runtime.NumCPU()
	// fmt.Println(n)
	// fmt.Println(runtime.GOROOT())
	// fmt.Println(runtime.GOOS)

	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Println("goroutine...")

	// 	}
	// }()
	// for i := 0; i < 4; i++ {
	// 	runtime.Gosched()
	// 	fmt.Println("main")
	// }
	// var ch1 chan bool
	ch1 := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine中，i", i)
		}
		ch1 <- "lichangdong"
	}()
	data := <-ch1 // 从ch1通道中读取数据
	fmt.Println("data-->", data)

}
