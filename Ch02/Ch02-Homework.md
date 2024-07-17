# HomeWork

### 1. 说明一下接口是什么, 和面向对象是什么关系?

Go语言中，接口是一组方法签名的集合，接口定义了对象的行为。一个对象只需要实现了接口定义的所有方法，即被称为实现了该接口。

```go
type Duck interface{} {
    Gua()
    Run()
}

type Rem struct {
    Name string
    age int
}

// 为 Rem 实现 Duck 接口
func(r *Rem) Gua() {
    fmt.Println("rem guagua")
}

func (r *Rem) Run() {
    fmt.Println("rem runrun")
}
```

接口是面向对象编程中的一个重要概念。接口定义了对象应该具有的行为，而实现了接口的对象则必须提供这些行为的具体实现。通过接口，可以实现对象之间的多态性，提高代码的灵活性和可复用性。

### 2. 举例说明鸭子类型

只要对象看起来像鸭子、走起来像鸭子，那么它就可以被视为鸭子
Q1中为Rem实现了Duck接口

### 3. go语言中的标准接口有哪些? 并举例说明1-2个接口的实现, 通过接口如何实现多态?

**Go语言中的常用接口**

- `io.Reader`: 用于读取数据的接口
  ```go
  type Reader interface {
      Read(p []byte) (n int, err error)
  }
  ```
- `io.Writer`: 用于写入数据的接口
  ```go
  type Writer interface {
      Write(p []byte) (n int, err error)
  }
  ```

```go
type Duck interface {
	Fly()
	Kua()
}

type Rem struct {
	Name string
	age  int
}

func (r *Rem) Fly() {
	r.age = 200
	fmt.Println("Rem can Fly!")
}

func (r *Rem) Kua() {
	fmt.Println("Rem can Kua!")
}

type Ram struct {
	Magic string
	Usage bool
}

func (r *Ram) Fly() {
	r.Usage = true
	fmt.Println("Ram use fly magic!")
}

func (r *Ram) Kua() {
	fmt.Println("Ram can Kua!")
}

func test(d Duck) {
	d.Fly()
	d.Kua()
}

func main() {
	rem := new(Rem)
	ram := new(Ram)
	test(rem)
	test(ram)
}
```

上面的代码段, `Rem`, `Ram`均实现了 `Duck`接口, 而 `test`函数则根据传入的不同实例有不同的行为, 从而实现了多态。

### 4. 函数传值和传引用有何不同？ 各举一个例子

- 传值（Pass by Value）：
  - 当以传值方式调用函数时，函数会创建原始数据的一个副本，并将副本传递给函数。
  - 任何对参数的更改都不会影响原始数据。
  - 传递大型数据结构时会消耗更多的内存，因为需要复制整个数据结构。
- 传引用（Pass by Reference）：
  - 当以传引用方式调用函数时，函数会接收原始数据的引用或地址，而不是副本。
  - 任何对参数的更改都会影响原始数据。
  - 传递引用通常更高效，因为不需要复制整个数据结构，只需要传递地址。

```go
func testValuePass(a int, b [3]string) {
	a = 10
	b[0] = "shit"
	b[1] = "shit"
	b[2] = "shit"
}

func testPointerPass(a []string) {
	a[0] = "shit"
}

func main() {
	a := 0
	b := [3]string{"rem", "ram", "fish"}
	c := []string{"rem", "ram", "fish"}
	testValuePass(a, b)
	testPointerPass(c)
	fmt.Println(a) // 0
	fmt.Println(b) // [rem ram fish]
	fmt.Println(c) // [shit ram fish]
}
```

### 5. 举例说明 函数变量

函数变量也有类型,既可以把函数变量赋给变量,也可以传递函数变量, 或者从其他函数返回函数变量。我们可以像使用普通变量一样使用函数变量,还可以把函数变量当作参数进行传递。
函数变量类型的零值是 nil,如果调用了值为 nil 的函数,则会导致宕机。我们通常使用 mil 和函数变量做判空比较
可以使用 `type Name func()`的格式来声明函数

```go
type Hello func(name string)

func hi(name string) {
	fmt.Printf("hi, %s!\n", name)
}

func shit(name string) {
	fmt.Printf("shit, %s!\n", name)
}

func main() {
	var h1 Hello = hi
	h1("rem")

	h1 = shit
	h1("rem")
}
```

### 6. 举例说明 匿名函数

```go
func main() {
	hi := func(name string) {
		fmt.Printf("hi, %s\n", name)
	}

	hi("rem")
}
```

### 7. 举例说明 闭包

闭包是包含自由变量的代码块,变量不在这个代码块或者全局上下文中定义,而是在定义代码块的环境中定义。要执行的代码块(因为自由变量包含在代码块中,所以这些自由变量及它们引用的对象没有被释放)为自由变量提供绑定的计算环境。

```go

func test(name string) (res string) {
	hi := func(name string) (newName string) {
		newName = "shit_" + name
		fmt.Printf("sorry %s, in 闭包, your name will change into %s\n", name, newName)
		return newName
	}
	res = hi(name)
	return res
}

func main() {
	name := "rem"
	newName := test(name)
	fmt.Println(newName)

}

```

### 8. 举例说明 变长函数

```go
func test(a ...int) {
	for _, value := range a {
		fmt.Println(value)
	}
}

func main() {
	test(1, 2, 4, 4)
}
```

### 9. 延长函数的调用顺序是什么？ 举例说明

`defer`的调用顺序是按照定义逆序的

```go
func main() {
	defer func() {
		fmt.Println(0)
	}()

	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(2)
	}()

	defer func() {
		fmt.Println(3)
	}()
}
```

### 10. go语言是如何做测试的？ 举例说明

go test 命令是Go语言包的测试驱动程序,这些包根据某些约定组织在一起。在一个包目录中,以*_test.go 结尾的文件就是 go test 编译的目标
以 `calc`包举例, 包下有个 `calc.go`, 该代码中存在 `Sum`函数, 相对该函数做测试

1. 在 `calc`包下新建 `calc_test.go`文件;
2. 在 `calc_test.go`中定义测试函数 `TestSum`, 测试函数输入的参数只能有 `*testing.T`类型
   ```go
   func TestSum(t *testing.T) {
       测试逻辑
   }
   ```
3. 运行测试函数
   ```bash
       go test # 运行所有的测试
       go test -v calc_test.go calc.go # 运行特定的测试文件
   ```

### 11. 如何理解 线程安全？

线程安全是指在多个线程并发执行时，程序的行为是正确的，并且不会因为多线程访问共享数据而导致数据破坏、程序崩溃或者其他不可预测的结果。
实现线程安全的方法主要如下：

1. 互斥锁
2. 原子操作
3. 无锁编程
4. 冻结变量
5. 使用线程安全的数据结构

### 12. 如何理解Go语言的并发模型？

线程是在进程之内的,可以把它理解为轻量级的进程。它可以被视为进程中代码的执行流程。
从本质上说,goroutine 是一种用户态线程,不需要操作系统进行抢占式调度
Go 语言不仅有 goroutine,还有强大的用来调度 goroutine、对接系统级线程的调度器

### 13. 缓冲通道与无缓冲通道有合不同？

无缓冲 channel 是指那些接收者没有能力保存任何值的channel。无缓冲 channel 要求执行发送操作的 goroutine 和执行接收操作的 goroutine 必须同时准备好,才能发送数据.
如果两个 goroutine 没有同时准备好,则会导致先执行发送操作或接收操作的 goroutine 处于阻塞等待状态。

```go
func main() {
	ch := make(chan string)
	go func() {
		ch <- "rem"
		fmt.Println("Send msg to channel")
	}()

	time.Sleep(1 * time.Second)
	<-ch
	fmt.Println("Recive msg from channel")
	time.Sleep(1 * time.Second)
}
```

缓冲 channel 有一个元素队列,队列的最大长度是在创建的时候通过 make 函数的第二个参数指定的.
缓冲 channel 上的发送操作是在队列尾部插入元素,接收操作是从队列头部移除一个元素。如果 channel 满了,则发送操作会阻塞所在的那个goroutine 上,直到另一个 goroutine 对此 channel 进行接收操作,留出相应的空间。如果 channel 是空的,则执行接收操作的 goroutine 会被阻塞, 直到另一个 goroutine 在此 channel 上发送数据。

```go
func main() {
	ch := make(chan string, 3)
	go func() {
		ch <- "rem"
		fmt.Println("Send rem to channel")
		ch <- "ram"
		fmt.Println("Send ram to channel")
		ch <- "fish"
		fmt.Println("Send fish to channel")
		ch <- "boy"
		fmt.Println("Send boy to channel")
		close(ch)
	}()

	time.Sleep(3 * time.Second)

	for res := range ch {
		fmt.Printf("Recived %s from channel\n", res)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("all done.")

}
```

运行上面的示例代码, 可以很明显的看到, goroutine在发送了三个名字之后会被阻塞, 直到主程序开始读取channel才会立即发送第四个名字

### 14. 单向通道优势是什么？

当一个 channel 作为一个参数时,我们可以让它同时发送和接收消息,这会使得调用者产生困惑。遵循最小化原则,我们可以只执行一种操作,要么发送 chan<- string,要么接收<-chan string,避免误用

### 15. 关闭通道，会造成哪些影响？

channel 会设置一个标志位来提示当前发送操作已经完毕,这个 channel 后面没有值了。
当一个 channel 关闭后,若再次发送数据,则会引发 panic。
当使用 for range 语句循环 channel 的值时,如果 channel 的值为 nil,那么这条 for 语句会永远地阻塞在 for 关键字那一行上。

### 16. 什么场景使用select?

select 是用来监听与 channel 有关的IO操作的,当VO操作发生时,会触发相应的动作。
select 的用法和switch 的用法几乎相同,都由一系列的 case 语句和一个默认分支组成。每个 case 语句指定一次通信。select 会一直等待,直到有通信通知有 case 可以执行止。

### 17. 举例说明 mutex和rwmutex

竞态是指多个 goroutine 在交错执行程序时无法得到正确的结果
sync.RWMutex中的Lock和UnLock方法分别对写锁进行锁定和解锁, RLock和RUnlock方法分别对读锁进行锁定和解锁

### 18. 举例说明 条件变量

条件变量(sync.Cond)是用于协调想要访问共享资源的那些goroutine 的。当共享资源的状态发生变化时,可以用它来通知被互斥锁阻塞的 goroutine
条件变量提供了三个方法:等待通知(wait)、单一通知 (signal)和广播通知 (broadcast)。
当利用条件变量等待通知时,需要在它基于的互斥锁保护下进行,并且需要在对应的互斥锁解锁之后再做通知或广播操作。

- [ ] 这里的概念还不是很清楚, 等后面遇到具体项目的时候再来盘一盘

### 19. 举例说明 WaitGroup

程序中也是存在的,多个协程各自执行任务,但是在某一时刻,需要其他协程的结果,要么其他协程已执行完,可以直接拿到结果;要么其他协程还没有执行完,需要等待。
sync 包的 WaitGroup 类型是并发安全的,它有三个指针方法:Add、Done 和 Wait。

```go
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
```

### 20. 举例说明 context.Context

Context 类型代表上下文的值,并且是并发安全的,可以被传播给多个 goroutine

- [ ] 等后续理解

### 21. 说说你对GO语言错误处理的理解？

Go不使用异常处理机制，而是依靠显式的错误返回值，让调用者决定如何处理错误。这种方式鼓励检查和处理每一个可能的错误，从而提高代码的健壮性和可读性。

### 22. go语言如何做依赖管理？

go语言使用 `go mod`来进行依赖包的管理

### 23. go mod的常用命令

```text
go mod 
- download 
- edit	# 编辑go.mod
- graph
- init
- tidy
- vendor  # 将依赖复制到vendor下
- verify
- Why
```
