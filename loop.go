package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

//for 循环,for与if后面的条件没有括号，没有while，if里面可以定义变量

//十进制转二进制，省略初始条件，相当于while
func convertToBin(n int) string  {
	result := ""
	for ; n > 0; n /= 2 {
		lsb :=n % 2
		result = strconv.Itoa(lsb) + result  //strconv.Itoa()函数为转字符串的函数
	}
	return result
	
}

// 一行一行读取文件并显示出来，省略初始条件以及递增条件，相当于while
func printFile(finename string)  {
	file, err := os.Open(finename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	
}


//for后面什么都不加，会陷入死循环，一直打印出abc
func forever()  {
	for {
		fmt.Println("abc")
	}
	
}

func main() {
	fmt.Println(
		convertToBin(5),   //101
		convertToBin(13),  //1101
		convertToBin(189),
		convertToBin(0),

	)

	printFile("abc.txt")
}