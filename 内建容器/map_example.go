package main

import "fmt"

//寻找最长不含有重复字符的子串

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


func lengthOfNonRepeatingSubStrZh(s string) int {
	lastOccurred :=make(map[rune]int)
	start := 0
	maxlenth := 0
	for i,ch := range []rune(s) {
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
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb")) //3
	fmt.Println(lengthOfNonRepeatingSubStr("aaaaa"))  //1
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))  //3
	fmt.Println(lengthOfNonRepeatingSubStr(""))  //0
	fmt.Println(lengthOfNonRepeatingSubStr("b"))  //1
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef")) //6
	fmt.Println(lengthOfNonRepeatingSubStrZh("这里是禹都一只猫")) //8
	fmt.Println(lengthOfNonRepeatingSubStrZh("一二三二一")) //3
}
