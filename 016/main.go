package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}
	lcd := location{123.11, 134.33}
	bytes, err := json.Marshal(lcd)
	exitOnError(err)

	fmt.Println(string(bytes))

}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
