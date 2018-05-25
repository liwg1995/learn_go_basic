package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode)  print() {
	fmt.Print(node.value)
}//定义方法

func (node *treeNode) setValue(value int)  {
	if node == nil{
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.value = value
}

//遍历树
func (node *treeNode) traverse()  {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
} //使用自定义的工厂函数，返回了局部变量的地址

func main() {
	var root treeNode
	fmt.Println(root) //{0 <nil> <nil>}
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	fmt.Println(root.left.right) //&{2 <nil> <nil>}

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc42000a060}]

	//root.print()
	root.right.left.setValue(4)
	root.traverse()
	

	
	//root.right.left.print()//0  node变为指针的时候，setValue之后获得的就是4了
	//fmt.Println()
	//root.print()
	//root.setValue(100)
	//fmt.Println()
	//root.print()//100
	//
	////pRoot := &root
	////fmt.Println()
	////pRoot.print()//100
	////pRoot.setValue(200)
	////fmt.Println()
	////pRoot.print()//200
	//var pRoot *treeNode
	//pRoot.setValue(200)
	////pRoot.print()
	//pRoot = &root
	//pRoot.setValue(300)
	//pRoot.print()

	

}

/**
小记

（值接受者是go特有的）
值/指针接受者均可接收值/指针
要改变值的话，需要用指针接受者
结构过大，也考虑指针接受者
不改变值的话，用值接受者（即非指针）
一致性：建议如有指针接受者，最好都是指针接受者

*/