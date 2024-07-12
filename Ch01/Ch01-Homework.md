# HomeWork

### 1. Go语言的变量命名规范

> 变量是对一块内存的命名, 程序可以通过定义一个变量来申请一块内存, 之后通过引用变量来使用这块内存

**命名规范**
名称的开头是一个字母(Unicode中的字符)或者下划线, 后面跟任意数量的字符(不能有特殊字符)

- 不能使用关键字;
- 不能使用内置的常量、类型和函数

### 2. 如何获取变量的地址? 如何取到地址中的值?

使用 `&`可以获取变量的地址, 对变量地址使用 `*`可以获取地址中的值

> **使用new(T)可以创建一个T类型的变量, 返回地址*T**

```go
t := new(int)  // var t *int
fmt.Println(t)	// 0x1400000e118
fmt.Println(*t)	// 0
```

```go
var value int = 10
pValue := &value
fmt.Printf("value :%d\n", *pValue)
```

### 3. 变量的生命周期是什么? 作用域是什么
- 生命周期是指程序运行时被程序其他部分所引用的起止时间
- 变量的生命周期指在程序执行过程中变量存在的时间段。包级别变量的生命周期就是整个程序的执行时间，称之为全局变量
- 局部变量的生命周期：每次执行声明语句时创建一个新的变量, 该变量一直存活到它不可访问为止, 这时它占用的存储空间将被回收。
- 堆是用来存放进程执行中被动态分配的内存段的, 它的大小不固定
- 栈是用来存放程序暂时创建的局部变量的, 即在函数大括号{}中定义的局部变量

作用域是声明在程序中出现的位置, 通常使用{}。当编译到一个名字的引用时, 将从最内层到全局中寻找其声明。
作用域的最大好处就是控制程序的访问权限

### 4. 创建变量有哪几种方式

```go
var v1 int 
var v2 int = 10
v3 := 10
pV := new(int)
*pV = 10
```

### 5. Go语言简单的数据类型都有哪些

``` text
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
complex64, complex128
bool, string
byte, rune
```

1. 初始化数组有哪几种方式

```go
var arr [5]int
arr := [5]int
arr := [5]int{1,2,3,4,5}
arr := [5]int{1:2, 3:5}
arr := new([5]int)
```

### 6. 遍历数组

```go
arr := [5]int{1,2,3,4,5}
for _, value := range arr {
    fmt.Println(value)
}
```

### 7. 如何复制切片

```go
arr := [5]int{1,2,3,4,5}
slice := arr[1:3]
s1 := make([]int, len(slice))
copy(s1, slice)
fmt.Println(s1)	// [2, 3]
```

### 8. 切片的增删查改

```go
s1 := make([]int, 0)
tmp := [3]int{1, 2, 3}
s1 = append(s1, tmp[:]...)
fmt.Println(s1)                // [1,2,3]
s1 = append(s1[:1], s1[2:]...) // [1,3]
fmt.Println(s1)
```

### 9. 修复下面代码的问题

```go
s := []string{"ram", "rbm", "rem"}
s2 := make([]*string, len(s))
for i, v := range s{
    s2[i] = &v
}
```

上面的代码没有问题, s2中存的是s对应string变量的地址, 可以使用 `*`来获取地址对应的值

```go
for i, p := range s{
    fmt.Printf("s[%d] = %s\n", i, *p)
}

```

### 10. 分别写一个if, switch和枚举的例子

```go
url := "http://www.baidu.com"

if resp, err := http.Get(url); err != nil {
		fmt.Println(err)
	} else {
		statusCode := resp.StatusCode
		switch {
		case statusCode == 200:
			fmt.Println("ping ok")
		case statusCode == 404:
			fmt.Println("404 not found")
		default:
			fmt.Printf("statusCode = %d\n", statusCode)
		}
	}

const (
	sunday = iota
	monday
)	// cannot use iota outside const

```

### 11. map有什么特点?
    1. map存储键值对;
    2. map所有键的类型和值的类型分别都是相同的;
    3. map不会存储键, 但是会存储他们的哈希值


### 12. 什么样的类型可以做map的key
    键的类型不能是函数、切片和map, 建议不要使用接口来做key


### 13. map的增删查改

```go
m1 := make(map[string]int, 5)
fmt.Println(m1)
m1["a"] = 0
m1["b"] = 2
fmt.Println(m1)

delete(m1, "a")
fmt.Printf("m1['b'] = %d\n", m1["b"])
m1["b"] = 10
fmt.Println(m1)
```

### 14. 函数的定义
函数的定义如下: 
```go
func 函数名(参数列表)(返回列表){
	函数体
}
```

### 15. 函数传参, 传值还是传引用?
实参是按值传递的, 但是如果调用方提供的实参包含引用类型(切片、指针、map、函数或者通道等), 那么函数修改形参时有可能间接的修改实参


### 16. 如何定义函数的多返回值?

```go
func test(a int, b string) (int, string ,error) {
	return 0, "", nil
}
```

### 17. 举例说明函数变量、匿名函数、闭包、变长函数?
**函数变量**
```go
type Test func(name string) // 函数声明

// 函数定义
func helloName(name string) {
	fmt.Printf("hello, %s!\n", name)
}

func main() {
	var hello Test = helloName
	hello("rem")	// hello, rem!
}
```

下面的示例只是给出了函数声明的具体用法
```go
type HelloT func(name string) // 函数声明

func helloName(name string) {
	fmt.Printf("hello, %s!\n", name)
}

func shitName(name string) {
	fmt.Printf("holy shit, %s!\n", name)
}

func main() {
	name := "rem"
	// name := "ram"
	var hi HelloT
	if name == "rem" {
		hi = helloName
	} else {
		hi = shitName
	}

	hi(name)
}
```

**匿名函数**
```go
do := func(name string) {
	fmt.Printf("hello, %s!\n", name)
}
do("rem")
```

**闭包**
闭包是包含自由变量的代码块,变量不在这个代码块或者全局上下文中定义,而是在定义代码块的环境中定义。要执行的代码块(因为自由变量包含在代码块中,所以这些自由变量及它们引用的对象没有被释放)为自由变量提供绑定的计算环境。
```go
func helloName(name string) string {
	do := func(name string) string {
		newName := fmt.Sprintf("shit_%s", name)
		return fmt.Sprintf("sorry, %s, 在闭包里你的名字变成%s\n", name, newName)
	}
	res := do(name)
	return res
}

func main() {
	name := "rem"
	res := helloName(name)
	fmt.Println(res)
}
```

**变长函数**
```go
func test(nums ...int) {
	for _, v := range nums {
		fmt.Println(v)
	}
}

func main() {
	test(1, 2, 3, 4)
}

```

### 18. 说一下面向对象设计的好处?
1. 模块化和封装;
2. 重用性;
3. 可维护性;
4. 灵活性和可扩展性;
5. 自然的建模方式

### 19. 方法的定义
> 方法可以看做是某种特定类型的函数。方法的声明和普通函数的声明类似, 只是在函数名称前面多了一个参数。这个参数是一个类型, 可以把这个方法绑定在对应的类型上。

```go
type Person struct {
	name string
	age  int
}

func (p Person) helloName() {
	fmt.Printf("hello, %s!\n", p.name)
}

func main() {
	p := Person{name: "rem", age: 200}
	p.helloName()
}

```
### 20. 指针接受者和值接收者有什么不同?
> 在Go语言中,值传递时会复制一个变量,在遇到下面两种情况时,应该使用指针类型作为方法的接收者。
> - 在调用方法时,需要更新变量。
> - 类型的成员很多,占用内存很大,这样的开销会让内存使用率迅速增大