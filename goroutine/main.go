package main

import (
	"fmt"
	"time"
)

func sleepgopher(id int) {

	time.Sleep(3 * time.Second)
	fmt.Println("...snore...", id)

}

func main() {
	for i := 0; i < 5; i++ {
		go sleepgopher(i)
	}
	time.Sleep(4 * time.Second)
}
