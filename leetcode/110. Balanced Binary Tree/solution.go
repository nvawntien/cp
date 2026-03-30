/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return checkHeight(root) != -1
}

func checkHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftHeight := checkHeight(root.Left)
    
    if leftHeight == -1 {
        return -1
    }

    rightHeight := checkHeight(root.Right)

    if rightHeight == -1 {
        return -1
    }

    if abs(leftHeight - rightHeight) > 1 {
        return -1
    }

    return max(leftHeight, rightHeight) + 1
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}