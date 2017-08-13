package main

import (
	"fmt"
	"time"
)

func main() {
	hello()
}

func hello() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("hello")
	}
}
