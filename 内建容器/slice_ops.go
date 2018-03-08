package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("%v, len=%d, cap=%d\n",s,len(s),cap(s))
}

func main() {
	fmt.Println("Creating Slice")
	//创建slice
	var s []int
	// Zero value for slice is nil , s == nil
	for i :=0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)
	printSlice(s)

	s1 := []int{2,4,6,8}
	printSlice(s1) //len=4, cap=4

	//构建一个长度为16的slice
	s2 := make([]int,16)
	printSlice(s2) //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
	//构建一个长度为16,cap为32的slice
	s3 := make([]int,10,32)
	printSlice(s3) //[0 0 0 0 0 0 0 0 0 0], len=10, cap=32

	fmt.Println("Copying Slice")
	//拷贝切片，把s1的内容拷贝到s2中了，要注意，s2的cap还是16
	copy(s2,s1)
	printSlice(s2)  //[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0]

	fmt.Println("delete elements from slice")
	//删除[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0]中的8，它的下标是3
	//s2[:3] + s2[4:]，go中slice没有加法
	s2 = append(s2[:3], s2[4:]...)//s2[4:]...是下标为4之后的所有的元素
	printSlice(s2) //[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0], len=15, cap=16

	fmt.Println("Poping from front")
	//删除头部元素
	front := s2[0]
	//fmt.Println(front)
	s2 = s2[1:]
	fmt.Println(s2)

	fmt.Println("Poping from back")
	tail := s2[len(s2) - 1]

	s2 = s2[:len(s2)-1]
	fmt.Println(s2)

	fmt.Println(front,tail)

	printSlice(s2)

}
/*
len=0, cap=0
len=1, cap=1
len=2, cap=2
len=3, cap=4
len=4, cap=4
len=5, cap=8
len=6, cap=8
len=7, cap=8
len=8, cap=8
len=9, cap=16
len=10, cap=16
len=11, cap=16
len=12, cap=16
len=13, cap=16
len=14, cap=16
len=15, cap=16
len=16, cap=16
len=17, cap=32
len=18, cap=32
len=19, cap=32
len=20, cap=32
len=21, cap=32
len=22, cap=32
len=23, cap=32
len=24, cap=32
len=25, cap=32
len=26, cap=32
len=27, cap=32
len=28, cap=32
len=29, cap=32
len=30, cap=32
len=31, cap=32
len=32, cap=32
len=33, cap=64
len=34, cap=64
len=35, cap=64
len=36, cap=64
len=37, cap=64
len=38, cap=64
len=39, cap=64
len=40, cap=64
len=41, cap=64
len=42, cap=64
len=43, cap=64
len=44, cap=64
len=45, cap=64
len=46, cap=64
len=47, cap=64
len=48, cap=64
len=49, cap=64
len=50, cap=64
len=51, cap=64
len=52, cap=64
len=53, cap=64
len=54, cap=64
len=55, cap=64
len=56, cap=64
len=57, cap=64
len=58, cap=64
len=59, cap=64
len=60, cap=64
len=61, cap=64
len=62, cap=64
len=63, cap=64
len=64, cap=64
len=65, cap=128
len=66, cap=128
len=67, cap=128
len=68, cap=128
len=69, cap=128
len=70, cap=128
len=71, cap=128
len=72, cap=128
len=73, cap=128
len=74, cap=128
len=75, cap=128
len=76, cap=128
len=77, cap=128
len=78, cap=128
len=79, cap=128
len=80, cap=128
len=81, cap=128
len=82, cap=128
len=83, cap=128
len=84, cap=128
len=85, cap=128
len=86, cap=128
len=87, cap=128
len=88, cap=128
len=89, cap=128
len=90, cap=128
len=91, cap=128
len=92, cap=128
len=93, cap=128
len=94, cap=128
len=95, cap=128
len=96, cap=128
len=97, cap=128
len=98, cap=128
len=99, cap=128
**/
//[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53 55 57 59 61 63 65 67 69 71 73 75 77 79 81 83 85 87 89 91 93 95 97 99 101 103 105 107 109 111 113 115 117 119 121 123 125 127 129 131 133 135 137 139 141 143 145 147 149 151 153 155 157 159 161 163 165 167 169 171 173 175 177 179 181 183 185 187 189 191 193 195 197 199]
//len=100, cap=128