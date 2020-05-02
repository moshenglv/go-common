package tree

import (
	"fmt"
	queue "go-common/structure/queue"
	"reflect"
)

type Element interface{}
type InitList []Element

type Node struct {
	Left  *Node
	Right *Node
	Val   Element
}

func CreateTree(preOrder InitList, inOder InitList) *Node {
	if len(preOrder) <= 0 {
		return nil
	}
	//将数组转换成Node
	perOrderList := make([]Node, 0)
	for _, v := range preOrder {
		perOrderList = append(perOrderList, Node{
			Val: v,
		})
	}
	inOderList := make([]Node, 0)
	for _, v := range inOder {
		inOderList = append(inOderList, Node{
			Val: v,
		})
	}

	return create(perOrderList, inOderList)
}

/**
前序遍历的第一个数肯定是root节点 然后在inorder中以找到的root节点为分割线 前面为左子树 后面为右子树
然后根据inorder中左子树和右子树 去找preorder中的 左子树和右子树  然后递归
*/
func create(preOrderList []Node, inOderList []Node) *Node {
	if len(preOrderList) == 0 {
		return nil
	}
	root := Node{Val: preOrderList[0].Val}

	leftPreTree := make([]Node, 0)
	rightPreTree := make([]Node, 0)

	leftInTree := make([]Node, 0)
	rightInTree := make([]Node, 0)

	for index, value := range inOderList {
		if reflect.DeepEqual(value.Val, root.Val) {
			leftInTree = inOderList[0:index]
			rightInTree = inOderList[index+1:]

			leftPreTree = preOrderList[1 : index+1]
			rightPreTree = preOrderList[index+1:]

			break
		}
	}

	root.Left = create(leftPreTree, leftInTree)
	root.Right = create(rightPreTree, rightInTree)
	return &root
}

/**
完全二叉树的顺序存储结构建立其二叉链式存储结构
*/
func FullTreeCreate(i int, arr []Element) *Node {
	t := &Node{nil, nil, arr[i]}
	if i < len(arr) && 2*i+1 < len(arr) {
		t.Left = FullTreeCreate(2*i+1, arr)
	}
	if i < len(arr) && 2*i+2 < len(arr) {
		t.Right = FullTreeCreate(2*i+2, arr)
	}
	return t
}

func PreOder(root *Node) {
	if root == nil {
		return
	}
	print(root)
	PreOder(root.Left)
	PreOder(root.Right)
}

func InOder(root *Node) {
	if root == nil {
		return
	}
	InOder(root.Left)
	print(root)
	InOder(root.Right)
}

func PostOder(root *Node) {
	if root == nil {
		return
	}
	PostOder(root.Left)
	PostOder(root.Right)
	print(root)
}

func LevelOder(root Node) {
	if &root == nil {
		return
	}
	queue := queue.New()
	queue.Push(root)
	for !queue.IsEmpty() {
		node := *(queue.Pop())
		n, _ := node.(Node)
		print(&n)
		if n.Left != nil {
			queue.Push(*n.Left)
		}
		if n.Right != nil {
			queue.Push(*n.Right)
		}
	}
}

func MaxDepth(root Node) int {
	if &root == nil {
		return 0
	}
	level := make([]Element, 0)
	queue := queue.New()
	queue.Push(root)

	floor, front, tail := 0, 0, queue.Size()

	for !queue.IsEmpty() {
		node := *(queue.Pop())
		n, _ := node.(Node)
		front++
		level = append(level, n.Val)
		if n.Left != nil {
			queue.Push(*n.Left)
		}
		if n.Right != nil {
			queue.Push(*n.Right)
		}

		if front == tail {
			front = 0
			tail = queue.Size()
			floor++
			fmt.Println(level)
			level = make([]Element, 0)
		}
	}
	return floor
}

func print(node *Node) {
	if node.Val == nil {
		return
	}
	fmt.Print(node.Val, " ")
}
