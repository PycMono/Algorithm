package tree

import "fmt"

func Test() {
	root := NewBinaryTreeNode(10)
	rightNode := NewBinaryTreeNode(20)
	leftNode := NewBinaryTreeNode(30)
	root.SetLeftChild(leftNode)
	root.SetRightChild(rightNode)

	root.PreOrder()
}

// 二叉树节点信息
type BinaryTreeNode struct {
	// 节点数据
	data interface{}

	// 父节点
	parent *BinaryTreeNode

	// 左节点
	leftNode *BinaryTreeNode

	// 右节点
	rightNode *BinaryTreeNode
}

// 设置左节点孩子
func (b *BinaryTreeNode) SetLeftChild(l *BinaryTreeNode) {
	if l == nil {
		return
	}

	// 如果当前节点包含了子节点，退出去
	if b.isLeftChild() {
		fmt.Println(fmt.Sprintf("SetLeftChild当前节点包含了左子节点b.leftNode.data=%d", b.leftNode.data))
		return
	}

	b.leftNode = l
}

// 设置右节点
func (b *BinaryTreeNode) SetRightChild(r *BinaryTreeNode) {
	if r == nil {
		return
	}

	// 如果当前节点包含了子节点，退出去
	if b.isRightChild() {
		fmt.Println(fmt.Sprintf("SetRightChild当前节点包含了左子节点b.rightNode.data=%d", b.rightNode.data))
		return
	}

	b.rightNode = r
}

// 前序遍历
func (b *BinaryTreeNode) PreOrder() {
	// 如果节点为nil，退出
	if b == nil {
		return
	}

	fmt.Println(fmt.Sprintf("PreOrder,value=%d", b.data))
	b.leftNode.PreOrder()
	b.rightNode.PreOrder()
}

// 中序遍历
func (b *BinaryTreeNode) InOrder() {
	// 如果节点为nil，退出
	if b == nil {
		return
	}

	b.leftNode.PreOrder()
	fmt.Println(fmt.Sprintf("PreOrder,value=%d", b.data))
	b.rightNode.PreOrder()
}

// 后序遍历
func (b *BinaryTreeNode) PosOrder() {
	b.leftNode.PreOrder()
	b.rightNode.PreOrder()
	fmt.Println(fmt.Sprintf("PreOrder,value=%d", b.data))
}

// 获取右节点
func (b *BinaryTreeNode) GetLeftNode() *BinaryTreeNode {
	return b.leftNode
}

// 获取右节点
func (b *BinaryTreeNode) GetRightNode() *BinaryTreeNode {
	return b.rightNode
}

// 判断是否存在左节点
func (b *BinaryTreeNode) isLeftChild() bool {
	return b.leftNode != nil
}

// 判断是否存在右节点
func (b *BinaryTreeNode) isRightChild() bool {
	return b.rightNode != nil
}

// 断开关系连接
func (b *BinaryTreeNode) breakDownRelationship(isLeft bool) {
	if b.parent == nil {
		return
	}

	b.parent.rightNode = nil
	b.parent.leftNode = nil
	b.parent = nil
}

// 创建新的节点
// 参数：
// data：数据
// 返回值：
// BinaryTreeNode对象
func NewBinaryTreeNode(data interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{
		data: data,
	}
}
