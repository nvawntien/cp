/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func createBinaryTree(descriptions [][]int) *TreeNode {
    nodes := make(map[int]*TreeNode)
    orphaned := make(map[int]bool)

    for _, des := range descriptions {
        parent, child, isLeft := des[0], des[1], des[2]
        
        if _, exists := nodes[parent]; !exists {
            nodes[parent] = &TreeNode {
                Val: parent,
            }
        }

        if _, exists := nodes[child]; !exists {
            nodes[child] = &TreeNode {
                Val: child,
            }
        }

        if isLeft == 1 {
            nodes[parent].Left = nodes[child]
        } else {
            nodes[parent].Right = nodes[child]
        }

        orphaned[child] = true
    }

    for _, des := range descriptions {
        if !orphaned[des[0]] {
            return nodes[des[0]]
        }
    }

    return nil
}