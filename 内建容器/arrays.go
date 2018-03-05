package main

import "fmt"

//[]内没有数字的话，是一个切片

//数组是值类型，意思就是拷贝
func printArray(arr [5]int)  {
	arr[0] = 100        // 单独定第一个数为100
	for _,v := range arr {
		fmt.Println(v)
	}

}

//指针方式
func pointArray(arr *[5]int)  {
	arr[0] = 100        // 单独定第一个数为100
	for _,v := range arr {
		fmt.Println(v)
	}

}


func main() {
	//定义方式
	var arr1 [5]int  //代表有五个int
	arr2 := [3]int{1,3,5} //:=定义得设置一个初值
	arr3 := [...]int{2,4,6,8,10}  //定义多个，自己还得数数，这个是让编译器来数"[]"里面点三个点"..."
	var grid [4][5]int  // 二维数组，四行五列，四个长度为5的二维数组

	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	//遍历方式
	for i := 0; i < len(arr3);i++{
		fmt.Println(arr3[i])
	}

	for j,v := range arr3{
		fmt.Println(j,v)
	//range关键字可以获取第几个以及第几个的值
	}
	for _,v := range arr3{ //"_"表示忽略这个参数
		fmt.Println(v)
		//只获取值，不取第几个
	}
	//为什么用range？
	//美观，明确
	fmt.Println("printArray(arr1)")
	printArray(arr1)//可以 ，单独定义第一个数为100之后，，遍历获取的第一个数为100
	fmt.Println("printArray(arr3)")
	printArray(arr3)//可以  ，单独定义第一个数为100之后，，遍历获取的第一个数为100
	//printArray(arr2)//不可以，因为arr2是3个长度的，printArray定义了5个长度
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1,arr3)//不遍历了，无指针的情况下，就直接拷贝定义的数组了。结果[0 0 0 0 0] [2 4 6 8 10]


	pointArray(&arr1)
	pointArray(&arr3)
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1,arr3)
	//有了指针类型之后，就变为引用类型了，上面pointArray函数定义了数组的一个数为100，所以引用了下来，结果[100 0 0 0 0] [100 4 6 8 10]

}

// [10]int与[20]int是不同类型
// 调用func f(arr [10]int)会  拷贝   数组
// Go语言一般不直接使用数组，一般用切片
