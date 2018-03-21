package main

import (
	"fmt"
	"unicode/utf8"
)

//在utf8字节中，一个汉字占三个字节

func main() {
	s := "Yes禹都一只猫!"
	fmt.Println(s)                //Yes禹都一只猫!  UTF-8
	fmt.Println(len(s))           //19
	fmt.Printf("%X\n", []byte(s)) //596573E7A6B9E983BDE4B880E58FAAE78CAB21      (打印具体的字节16进制)
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) //59 65 73 E7 A6 B9 E9 83 BD E4 B8 80 E5 8F AA E7 8C AB 21
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune.    ch是int32的，四子节的整数
		fmt.Printf("(%d %X)", i, ch) //(0 59)(1 65)(2 73)(3 79B9)(6 90FD)(9 4E00)(12 53EA)(15 732B)(18 21)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) //Rune count: 9
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch) //Y e s 禹 都 一 只 猫 !
	}
	fmt.Println()


	for i,ch := range []rune(s) { //每个rune出来是四个字节
		fmt.Printf("(%d %c) ",i,ch) //(0 Y) (1 e) (2 s) (3 禹) (4 都) (5 一) (6 只) (7 猫) (8 !)
	}
	fmt.Println()

}
/**
rune相当于go的char
＊ 使用range遍历pos，rune对。pos是不连续的，如上面程序中的i，ch
＊ 使用utf8.RuneCountInString获得字符数量
＊ len是获得字节长度
＊ 用[]byte()是获得字节
**/


/**
其他字符串的操作：
1.Fields（可以认空格），Split,Join   [是表示的是 字符串分开或者合起来]
2.Contains，Index [查找子串等]
3.ToLower，ToUpper [转为大小写]
4.Trim,TrimRight,TrimLeft
*/