package main

import "fmt"

func GoName() chan string {
	ch := make(chan string)
	go func(ch chan string) {
		ch <- "rem"
		ch <- "ram"
		close(ch) // 记得关闭通道, 不关闭会导致死锁
	}(ch)

	return ch
}
func main() {
	ch := GoName()

	for name := range ch {
		fmt.Println(name)
	}
}
