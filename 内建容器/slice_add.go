package main

import "fmt"


// 切片添加元素
func main() {

	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]
	fmt.Println(s1)
	s2 := s1[3:5]
	fmt.Println(s2)
	s3 := append(s2,10)
	fmt.Println(s3)
	s4 := append(s3,11)
	fmt.Println(s4)
	s5 := append(s4,12)
	fmt.Println(s4)
	s6 := append(s5,13)
	fmt.Println(s6)
	fmt.Println(arr)  //[0 1 2 3 4 5 6 10]
}

//1.添加元素时，如果超越了cap，系统会重新分配更大的底层数组
//2.由于值传递的关系，必须接收append的返回值
//3. s = append(s,val)
