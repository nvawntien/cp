func sumRootToLeaf(root *TreeNode) int {
    ans := 0
    
    if root.Left == nil && root.Right == nil {
        return root.Val
    }

    if root.Left != nil {
        ans += dfs(root.Left, root.Val)
    }

    if root.Right != nil {
        ans += dfs(root.Right, root.Val)
    }

    return ans
}

func dfs(root *TreeNode, sum int) int {
    cur := sum * 2 + root.Val
    ans := 0

    if root.Left == nil && root.Right == nil {
        return cur
    }

    if root.Left != nil {
        ans += dfs(root.Left, cur)
    }

    if root.Right != nil {
        ans += dfs(root.Right, cur)
    }

    return ans
}