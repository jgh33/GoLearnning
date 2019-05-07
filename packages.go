package main


var Name = "sss" //大写可以被导出，相当于别的包可以用到它，不然找不到
var Age = 12

func init() {  //所有包都有一个或多个init函数，包的初始化顺序为
//首先初始化包级别的变量
//调用init函数。先初始化被导入的包，只会初始化一次。

}