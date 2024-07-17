package main

import (
	"fmt"
	"time"
)

func Hello(name string) {
	fmt.Printf("hello, %s!\n", name)
}

func main() {
	go Hello("rem")

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("i is %d\n", i)
		}()
	}
	time.Sleep(1 * time.Second)
}
