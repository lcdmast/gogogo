package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//匿名并发函数
	go func() {
		for i := 3; i <= 3; i-- {
			ch <- i
			time.Sleep(time.Second)
		}

		// fmt.Println("start goroutine")
		// ch <- 0
		// fmt.Println("exit ")
		//通过通道通知main 函数

	}()

	for data := range ch {
		fmt.Println(data)

		if data == 0 {
			break
		}
	}

	// fmt.Println("wait ")
	// <-ch

	// fmt.Println("done ")
}
