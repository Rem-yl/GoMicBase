package checkout

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

type Checkout func(a int, b int) int

// GetTotal返回值是个 func(int) int
func GetTotal(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

type GenerateRandom func() int

// RandomSum 返回 GenerateRandom 这个函数
func RandomSum() GenerateRandom {
	a, b := rand.Intn(10), rand.Intn(20)
	return func() int {
		a, b = b, a+b
		return a
	}
}

func (g GenerateRandom) Read(p []byte) (n int, err error) {
	next := g() // GenerateRandom是 func() int 函数的别名
	if next > 10 {
		fmt.Printf("next > 10: %d\n", next)
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func PrintResult(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
