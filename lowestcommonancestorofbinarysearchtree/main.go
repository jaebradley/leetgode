package lowestcommonancestorofbinarysearchtree

// TreeNode is the definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(v1, v2 *TreeNode) int {
	if v1.Val > v2.Val {
		return v1.Val
	}

	return v2.Val
}

func min(v1, v2 *TreeNode) int {
	if v1.Val < v2.Val {
		return v1.Val
	}

	return v2.Val
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return helper(root, max(p, q), min(p, q))
}

func helper(root *TreeNode, min int, max int) *TreeNode {
	if root.Val > min && root.Val > max {
		return helper(root.Left, min, max)
	}

	if root.Val < min && root.Val < max {
		return helper(root.Right, min, max)
	}

	return root
}
