package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)


func eval(a, b int, op string) (int, error)  {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b) //有两个值，但是只是想要一个值，那么另外一个一个下划线就可以了"_"

		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s" , op)
	}
	
}


// 返回多个值
// 求带余数的除法
func div(a, b int)  (q, r int) {
	return a / b ,a % b

}

//参数op是个函数，这个函数有两个参数是int型，返回结果为int型，apply接下来还有a，b两个int的参数,apply最后返回也是一个int
// 然后将apply传入的参数传给op这个函数
func apply(op func(int, int) int,a, b int) int {
	//获取函数的名字
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling function %s with args" + "(%d, %d)\n", opName, a, b)
	return op(a, b)

}//结果：Calling function main.pow with args(3, 4)  81


func pow(a,b int) int {
	return int(math.Pow(float64(a), float64(b)))
	
}


//可变参数列表   ...int表示随便多个int的参数
func sum(number ...int) int  {
	s := 0
	for i := range(number) {
		s += number[i]
	}
	return s
}


func main() {
	if result, err := eval(3,4,"x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}


	fmt.Println(eval(3,4,"*"))
	//fmt.Println(div(13, 3))
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3,4))

	fmt.Println(apply(
		func(a int, b int) int {
		return int(math.Pow(
			float64(a), float64(b)))

	},3,4))  //结果：Calling function main.main.func1 with args(3, 4)  81



	fmt.Println(sum(1, 2, 3, 4, 5)) //结果：15

}
