package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}
type myTreeNode struct {
	node *Node
}

//方法
func (nodes *Node) print() {
	fmt.Println(nodes.value)
}

//函数
func createNode(value int) *Node {
	return &Node{value: value}
}
func (myNode *myTreeNode) postO() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.left}
	left.postO()
	right := myTreeNode{myNode.node.right}
	right.postO()
	myNode.node.print()
}
func main() {
	var root Node
	root = Node{value: 3}
	root.left = &Node{}
	root.right = &Node{5, nil, nil}
	root.left.right = createNode(2)
	root.right.left = &Node{}
	root.right.left.value = 2
	myR := myTreeNode{&root}
	myR.postO()
}
