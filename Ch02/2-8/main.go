package main

/* 2-8 写文件操作 */

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func test(filename string, p []*Person) {
	dir, _ := os.Getwd()
	newDir := dir + "/" + filename
	f, err := os.Create(newDir)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	for _, person := range p {
		line := fmt.Sprintf("Name: %s, Age: %d", person.Name, person.Age)
		fmt.Fprintln(w, line)
	}
}

func main() {
	filename := "test.txt"
	p := make([]*Person, 3)
	p[0] = &Person{"rem", 10}
	p[1] = &Person{"ram", 20}
	p[2] = &Person{"rbm", 30}
	test(filename, p)
}
