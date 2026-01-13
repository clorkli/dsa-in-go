package algo

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func LCA(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := LCA(root.Left, p, q)
	right := LCA(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}

	return nil
}

func main() {
	// 构建测试二叉树:
	//       3
	//      / \
	//     5   1
	//    / \ / \
	//   6  2 0  8
	//     / \
	//    7   4
	
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	// 测试用例1: p = 5, q = 1
	p1 := root.Left       // 节点 5
	q1 := root.Right      // 节点 1
	lca1 := LCA(root, p1, q1)
	fmt.Printf("LCA of %d and %d is: %d\n", p1.Val, q1.Val, lca1.Val)

	// 测试用例2: p = 5, q = 4
	p2 := root.Left                      // 节点 5
	q2 := root.Left.Right.Right          // 节点 4
	lca2 := LCA(root, p2, q2)
	fmt.Printf("LCA of %d and %d is: %d\n", p2.Val, q2.Val, lca2.Val)

	// 测试用例3: p = 6, q = 4
	p3 := root.Left.Left                 // 节点 6
	q3 := root.Left.Right.Right          // 节点 4
	lca3 := LCA(root, p3, q3)
	fmt.Printf("LCA of %d and %d is: %d\n", p3.Val, q3.Val, lca3.Val)
}