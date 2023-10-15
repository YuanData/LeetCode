package lc

import "testing"

// TreeNode 代表一個二叉樹節點
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth 計算二叉樹的最大深度
func maxDepth(root *TreeNode) int {
	// 如果節點為 nil，則深度為 0
	if root == nil {
		return 0
	}

	// 遞迴計算左子樹和右子樹的深度
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// 使用內建的 max 函數返回兩者中的較大值
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected int
	}{
		{createBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}), 3},
		{createBinaryTree([]int{1, -1, 2}), 2},
		{createBinaryTree([]int{1}), 1},
		{nil, 0},
		{createBinaryTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), 4},
	}

	for _, tt := range tests {
		actual := maxDepth(tt.root)
		if actual != tt.expected {
			t.Errorf("got %d, expected %d", actual, tt.expected)
		}
	}
}

// Helper function to create a binary tree from a slice.
func createBinaryTree(data []int) *TreeNode {
	if len(data) == 0 || data[0] == -1 {
		return nil
	}
	root := &TreeNode{Val: data[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if i < len(data) && data[i] != -1 {
			node.Left = &TreeNode{Val: data[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(data) && data[i] != -1 {
			node.Right = &TreeNode{Val: data[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
