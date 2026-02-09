func balanceBST(root *TreeNode) *TreeNode {
    var nodes []int
    
    var inorder func (*TreeNode)
    
    inorder = func (node *TreeNode) {
        if node == nil {
            return
        }
        inorder(node.Left)
        nodes = append(nodes, node.Val)
        inorder(node.Right)
    }

    inorder(root)

    var buildBalanceBST func(int, int) *TreeNode

    buildBalanceBST = func (left, right int) *TreeNode {
        if left > right {
            return nil
        }

        mid := (left + right) >> 1
        root := &TreeNode{Val: nodes[mid]}
        
        root.Left = buildBalanceBST(left, mid-1)
        root.Right = buildBalanceBST(mid+1, right)
        
        return root
    }

    return buildBalanceBST(0, len(nodes)-1)
}