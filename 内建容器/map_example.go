package main

import "fmt"

//寻找最长不含有重复字符的字串

//比如：abcabcbb  -->  adc   ;   aaaaa --> a;    pwwkew  -->  wke
//分析：
/*
对于每一个字母：
1.lastOccurred[x]不存在，或者<start -> 无需操作
2.lastOccurred[x] >= start -> 更新start
3.更新lastOccurred[x],更新maxLength
*/

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred :=make(map[byte]int)
	start := 0
	maxlenth := 0
	for i,ch := range []byte(s) {
		if lastI,ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxlenth {
			maxlenth = i - start + 1
		}
		lastOccurred[ch]  = i

	}
	return maxlenth

}
func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("aaaaa"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("这里是禹都一只猫"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
	//3
	//1
	//3
	//0
	//1
	//6
	//17
	//5      因为做了byte的转换s，所以这些ASII码就不行了，所以中文现实的是这些，，
}
