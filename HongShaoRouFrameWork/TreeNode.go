package HongShaoRouFrameWork

import "fmt"

// 二叉树定义
//结点
type Node struct {
	Value       int
	Left, Right *Node
}

//打印结点值
func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

//设置结点值
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil " +
			"node.Ignored.")
		return
	}
	node.Value = value
}

//前序遍历
func (node *Node) preOrder() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.preOrder()
	node.Right.preOrder()
}

//创建结点
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
