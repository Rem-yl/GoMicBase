package main

import (
	"fmt"
	"sync"
)

func sayHello(name string, wg *sync.WaitGroup) {
	fmt.Printf("hello, %s!\n", name)
	wg.Done() // Go的指针类型可以调用所指向对象的方法, 编译器会自动解引用指针
}

func main() {
	nameList := []string{"rem", "ram", "fisher"}
	var wg sync.WaitGroup

	for _, name := range nameList {
		wg.Add(1)
		go sayHello(name, &wg) // 这里如果不传&wg而是wg, 会有call of sayHello copies lock value: sync.WaitGroup contains sync.noCopy
	}

	wg.Wait() // goroutines没有执行完会阻塞在这里
	fmt.Println("say hello done.")
}
