package main



import (
	"math/cmplx"
	"fmt"
	"math"
)

func euler()  {
	c := 3 + 4i   // 类型为复数类型complex：实部虚部都是float，complex64的是float32，complex128是float64
	fmt.Println(cmplx.Abs(c))

	//欧拉公式:  e^iπ + 1 = 0
	//1：
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1) //Pow，括号意思以e为底数，iπ次方
	//2.
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1) //Exp即为e为底数的函数

}
// 没有char，只有rune字节类型，是32位的

//类型转换
func triangle()  {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b))) //Sqrt必须是浮点型，c定义为int，所以再次强制转换
	fmt.Println(c)

}


//常量定义  ：也可以定义在包内部，即函数的外面
func consts()  {
	const filename string = "abc.txt"
	const a, b = 3, 4  //没有定义a,b的数据类型，可以各种类型的使用
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename, c)
	
}

//枚举类型的常量
func enums()  {
	const(
		cpp = iota    //iota是自增值
		//java
		_     //    _表示占位
		python
		golang
		javascript
	)

	// b, kb, mb, gb, tb, pb
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb

	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}


func main() {
	euler()
	triangle()
	consts()
	enums()
}
