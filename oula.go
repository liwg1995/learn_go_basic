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
// rune字节类型，是32位的

//类型转换
func triangle()  {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b))) //Sqrt必须是浮点型，c为int，所以再次强制转换
	fmt.Println(c)

}


func main() {
	euler()
	triangle()
}
