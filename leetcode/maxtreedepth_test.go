package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Rigth *TreeNode
}

func maxDepth(root *TreeNode) int {
	depth := 0
	return depth
}

func figureDepth(curNode *TreeNode, d int) int {
	if curNode == nil {
		return d
	}
	if curNode.Left == nil {
		return figureDepth(curNode.Rigth)
	}
}
