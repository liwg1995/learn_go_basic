package main

import (
	"io/ioutil"
	"fmt"
)
//条件语句
//switch
func grade(score int) string {  //传入int型的score参数，返回一个string类型的
	g := ""
	switch {
	case score < 0 || score > 100 :
		panic(fmt.Sprintf("Wrong score: %d", score)) //panic遇到出错终端程序运行，报错
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	//default:
	//	panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}


//条件语句

func main() {
	const filename  = "abc.txt"
	//写法一：
	//contents, err := ioutil.ReadFile(filename)  //ioutil返回的是两个值，一个是[]byte，一个是error
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}

	//写法二：
	if contents, err :=ioutil.ReadFile(filename); err !=nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}


	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		//grade(101),
	)
}
