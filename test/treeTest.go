package test

import (
	"fmt"
	"go-common/structure/tree"
)

func TreeTest() {

	fmt.Println("前序遍历:")
	tree.PreOder(getRoot())
	fmt.Println()

	fmt.Println("中序遍历:")
	tree.InOder(getRoot())
	fmt.Println()

	fmt.Println("后序遍历:")
	tree.PostOder(getRoot())
	fmt.Println()

	fmt.Println("层级遍历:")
	tree.LevelOder(*getRoot())
	fmt.Println()

	//根据数组创建二叉树
	fmt.Println("二叉树创建:")
	list := tree.InitList{1, 3, 4, 5, 6, 7, 8, 10, nil, nil, nil, nil, nil, nil, nil}
	node := tree.FullTreeCreate(0, list)
	tree.PostOder(node)
}

func getRoot() *tree.Node {
	preOder := tree.InitList{1, 3, 5, 10, 6, 4, 7, 8}
	inOder := tree.InitList{10, 5, 3, 6, 1, 7, 4, 8}
	root := tree.CreateTree(preOder, inOder)
	//root := CreateTest()
	return root
}

func CreateTest() *tree.Node {
	root := &tree.Node{Val: 1}
	root.Left = &tree.Node{Val: 3}
	root.Right = &tree.Node{Val: 4}

	root.Left.Left = &tree.Node{Val: 5}
	root.Left.Right = &tree.Node{Val: 6}

	root.Left.Left.Left = &tree.Node{Val: 10}

	root.Right.Left = &tree.Node{Val: 7}
	root.Right.Right = &tree.Node{Val: 8}

	return root
}
