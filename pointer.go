package main

import "fmt"

//Go语言的指针表示方式是 "*int",没有运算方式，Go语言的传递方式是值传递
//通过指针起到了引用传递的效果


//交换两个变量的位置，不用定义中间变量
//func swap(a,b int)  {
//	b, a = a, b
//
//}
//因为a，b是值传递，所以swap不能直接用

//更改为指针方式
func swap(a,b *int)  {
	*b, *a = *a,*b
}

//非指针型的方式，非有中间值
func swap1(c, d int) (int,int)  {
	return d,c
}

func main() {
	//a,b := 3,4
	//swap(a,b)
	//fmt.Println(a,b)
	//直接调用的话为3,4
	a,b := 3,4
	swap(&a, &b)  //&取指针地址，就换过来了
	fmt.Println(a,b)

	c,d := 1,2
	c,d = swap1(c,d)
	fmt.Println(c,d)

}
