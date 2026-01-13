package algo

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
