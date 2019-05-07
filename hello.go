package main

import (
	"fmt"
	"unsafe"
)

func main() {
	println("hello go")
	var (
		name   = "naveen"
		age    = 2
	)
	fmt.Printf("type %T value %v \n", name, name)

	fmt.Println("my name is", name, ", age is", age)

	var width, height int = 100, 50 // 声明多个变量
	fmt.Println("width is", width, "height is", height)

	named, aged := "naveen", 29 // 简短声明
	fmt.Println("my name is", named, ", age is", aged)


	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) // a 的类型和大小
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小

	const aa = 55//常量
	//const bb = math.Sqrt(4) // 不允许, 编译的时候就要确定下来

	//} else {    else必须和括号在同一行

	//for i := 1; i <= 10; i++ {

	//for {           无限循环
	//}

	go running1()

	// 接受命令行输入, 不做任何事情
	var input string
	fmt.Scanln(&input)
}

func calculateBill1(price int, num int) int {
	var totalPrice = price * num // 商品总价 = 商品单价 * 数量
	return totalPrice // 返回总价
}

func calculateBill2(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps1(length, width float64)(float64, float64) { //多值返回
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectProps2(length, width float64)(area, perimeter float64) { //默认返回值
	area = length * width
	perimeter = (length + width) * 2
	return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}
