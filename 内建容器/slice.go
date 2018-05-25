package main

import "fmt"

//切片,不是值类型，是一个视图

func updateSlice(s []int)  { //参数[]内没有的话，表示切片类型
	s[0] = 100
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	//s := arr[2:6]
	fmt.Println("arr[2:6] = ",arr[2:6])
	fmt.Println("arr[:6] = ",arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 = ",s1)
	s2 := arr[:]
	fmt.Println("s2 = ",s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	//输出：
	//After updateSlice(s1)
	//[100 3 4 5 6 7]
	//[0 1 100 3 4 5 6 7]

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	//输出：
	//After updateSlice(s2)
	//[100 1 100 3 4 5 6 7]
	//[100 1 100 3 4 5 6 7]

	//fmt.Println("After updateSlice(s2)")
	//updateSlice(s2)
	//fmt.Println(s2)
	//fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	//输出：
	//Reslice
	//[100 1 100 3 4 5 6 7]
	//[100 1 100 3 4]
	//[100 3 4]

	//slice拓展
	fmt.Println("Extending slice")
	arry := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println(arry)
	t1 := arry[2:6]
	fmt.Println(t1)
	fmt.Printf("t1=%v, len(t1)=%d, cap(t1)=%d\n",
		t1, len(t1), cap(t1))
	t2 := t1[3:5]
	fmt.Println(t2)
	fmt.Printf("t2=%v, len(t1)=%d, cap(t1)=%d\n",
		t1, len(t1), cap(t1))

	//输出：
	//Extending slice
	//[0 1 2 3 4 5 6 7]
	//[2 3 4 5]
	//t1=[2 3 4 5], len(t1)=4, cap(t1)=6
	//[5 6]
	//t2=[2 3 4 5], len(t1)=4, cap(t1)=6

	// 直接访问s1[4],s1[5]是访问不了的，
	// 但是，这个东西仍然存在，直接访问是访问不了的，切片却可以读出来，切片之后的6，7虽然看不见了，切片仍然可以取到，是可以扩展的
	// slice是可以向后拓展的，但不可以向前拓展
	// s[i]不可以超过len(s)，向后拓展不可以超过底层数组cop(s)
}