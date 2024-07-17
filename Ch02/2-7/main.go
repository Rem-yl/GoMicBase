/* 2-7 函数式编程 */
package main

import (
	"GoMicBase/Ch02/2-7/checkout"
	"fmt"
)

func hello(name string) {
	fmt.Printf("hello, %s!\n", name)
}

func testHello() {
	hello("rem")
	helloRam := hello // 函数是一等公民
	helloRam("ram")

	helloFish := func() {
		fmt.Printf("Fish say hello!\n")
	} // 匿名函数
	helloFish()
}

func testCheckout() {
	var ck checkout.Checkout = func(a int, b int) int {
		return a + b
	}

	fmt.Println(ck(10, 20))

	totalFunc := checkout.GetTotal(12) // GetTotal 返回的是 func(int) int函数
	sum := totalFunc(68)
	fmt.Println(sum)

	totalFunc = checkout.GetTotal(sum) // 在调用一次这个
	sum = totalFunc(50)
	fmt.Println(sum)
}

func testRandomSum() {
	r := checkout.RandomSum() // var r GenerateRandom
	checkout.PrintResult(r)   // GenerateRandom 实现了 Reader接口
}

func main() {
	testHello()
	fmt.Println("------ Test Checkout -------")
	testCheckout()
	fmt.Println("------ Test RandomSum -------")
	testRandomSum()
}
