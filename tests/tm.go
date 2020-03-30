package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("I want to sleep.")
	time.Sleep(5 * time.Second)
	fmt.Println("I just sleep 5s.")
}
