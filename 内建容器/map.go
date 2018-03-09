package main

import "fmt"

func main() {

	fmt.Println("Creating map")
	//定义
	//1.
	m := map[string]string {
		"name" : "ccmouse",
		"course" : "golang",
		"site" : "olei.me",
		"quality" : "notbad",
	}
	fmt.Println(m)//map[course:golang site:olei.me quality:notbad name:ccmouse]

	//2.
	m2 := make(map[string]int)//空map的定义方式 m2 == empty map
	fmt.Println(m2)//map[]

	var m3 map[string]int
	fmt.Println(m3)//map[]     m3 == nil



	fmt.Println("Traversing map")
	for k ,v := range m {
		fmt.Println(k,v)
	}
	//运行结果：
	//Traversing map
	//quality notbad
	//name ccmouse
	//course golang
	//site olei.me
	//顺序有可能不同，顺序是无序的

	fmt.Println("Getting values")
	courseName ,ok := m["course"]
	fmt.Println(courseName,ok)//golang true     表示获得了golang，是存在这个map里面的

	causeName ,ok := m["cause"]//如果拼写错误了呢？
	fmt.Println(causeName,ok)//返回的是空串  zero value，结果是:  false   表示获得了一个空串，不存在在这个map里面

	if causeName ,ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key is not exist")
	}//key is not exist


	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")

	name,ok = m["name"]
	fmt.Println(name,ok)

	if name,ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("the name's key is not exist")
	}

}

//总结
//1 创建：make(map[string]int)
//2.获取元素：m[key]
//3.key不存在的时候，获得的value类型是初始值
//4.用value,ok := m[key]来判断是否存在key
//5.delete删除一个key

//遍历
//1.使用range遍历key，或者遍历key，value对
//2.不保证遍历顺序，如需排序，需要手动对key排序
//3.使用len获取元素的个数

//map的key
//1.map使用哈希表，必须可以比较相等
//2.除了slice，map，function的内建类型都可以作为key
//3.Struct类型不包含上述字段，也可以作为key
