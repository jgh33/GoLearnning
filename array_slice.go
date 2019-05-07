package main



func  ssss()  {
	//var a [3]int
	//var b [3][2]string
	cc := [3]int{12, 78, 50}//快捷初始化+赋值
	//a := [...]int{12, 78, 50} //自动计算长度
	//Go 中的数组是值类型而不是引用类型。
	for index, value := range cc {
		print(index + value)
	}
}