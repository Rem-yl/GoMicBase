package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("hello, %s!\n", name)
	time.Sleep(2 * time.Second)
}

func main() {
	var wg sync.WaitGroup
	now := time.Now()

	wg.Add(1)
	go hello("rem", &wg)

	wg.Add(1)
	go hello("ram", &wg)

	wg.Add(1)
	go hello("fish", &wg)

	wg.Wait()
	waitTime := time.Since(now)

	fmt.Println(waitTime)
}
