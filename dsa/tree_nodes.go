package dsa

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func ArrayToBinaryTree(arr []int) *TreeNode {
    root := &TreeNode{}
    stack := Stack[*TreeNode]{}
    stack.Push(root)
    for ind := range arr {
        node := stack.Peek()
        if node.Val == 0 {
            node.Val = arr[ind]
            continue
        }
        if node.Left == nil {
            node.Left = &TreeNode{
                Val: arr[ind],
            }
            stack.Push(node.Left)
        } else if node.Right == nil {
            node.Right = &TreeNode{
                Val: arr[ind],
            }
            stack.Push(node.Right)
            stack.Pop()
        }
    }

    return root
}
