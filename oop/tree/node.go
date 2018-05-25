//package main
package tree

import "fmt"

//go语言只支持封装，不支持继承和多态
//go语言没有class，只有struct

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() { //为treeNode结构定义方法，在函数名前面加一个(node treeNode)表示的是接收者，print是给node来接收的
	fmt.Println(node.Value)
}

/*
* 显示定义和命名方法接受者
*
*/

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil" + "node. Ignored")
		return
	}

	node.Value = value
}
//
//func (node *Node) Traverse()  {
//	if node == nil {
//		return
//	}
//	node.Left.Traverse()
//	node.Print()
//	node.Right.Traverse()
//}

func CreateNode(value int) *Node { // 使用自定义工厂函数
	return &Node{Value: value} //返回了 局部变量的地址,不需要知道分配在栈上还是堆上，编译器自动分配
}

//func main() {
//	var root treeNode
//	//fmt.Println(root) //{0 <nil> <nil>}
//	root = treeNode{value: 3}
//	root.left = &treeNode{}
//	root.right = &treeNode{5, nil, nil}
//	root.right.left = new(treeNode)
//	root.left.right = createNode(2)
//	fmt.Println(root.left.right) //&{2 <nil> <nil>}
//
//	//上面建立了一个树
//	/*
//							   3
//							／	  \
//	                       0       5
//					        \      /
//						     2    0
//	*/
//
//	nodes := []treeNode{
//		{value: 3},
//		{},
//		{6, nil, &root},
//	}
//	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc42000a060}]
//	//不论地址还是结构本身，一律使用 . 来访问
//
//	root.print() //3
//	//fmt.Println()
//	root.right.left.setValue(4) //接收者为一个指针的话，那么调用的时候直接就是一个地址
//	root.right.left.print()     //0  改为指针之后输出为：4
//	//fmt.Println()
//
//	proot := &root
//	proot.print() //3
//	proot.setValue(100)
//	proot.print() //100
//
//}

