package main

import "fmt"

var aa = 3
var ss = "kkk"
var bb = true
var (
	cc = 2
	dd = "aaa"
	ff = false
)
//可以使用var()集中多个变量
//函数外面定义变量不能使用“:==”
//函数内的变量，定义了必须要使用，包变量可以不用
//外面定义变量不是全局变量，是包内部的变量，Go没有全局变量的称呼

//无值变量
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

//有值变量
func variableInitialValue() {
	var a, b int = 3, 2
	var s string = "hello"
	fmt.Printf("%d %d %q\n", a, b, s)
}

//编译器自动判断类型
func variableTypeDeduction() {
	var a, b, c, d = 3, 2, true, "def"
	fmt.Println(a, b, c, d)
}

//更短的定义变量
func variableShorter() {
	a, b, c, d := 3, 2, true, "def"
	b = 5
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Println("hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
}
