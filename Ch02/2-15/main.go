package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan interface{})
	ch2 := make(chan string)
	ch3 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 10
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "rem"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- 10086
	}()

	select {
	case data := <-ch1:
		fmt.Printf("ch1未阻塞, data: %v\n", data)
	case data := <-ch2:
		fmt.Printf("ch2 data: %s\n", data)
	case data := <-ch3:
		fmt.Printf("ch3 data: %d\n", data)
	default:
		fmt.Println("默认值")
	}

}
