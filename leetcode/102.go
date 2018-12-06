package main

import (
    "fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}

    if root == nil {
        return res
    }

    queue := make([]*TreeNode, 0)
    queue = append(queue, root)

    for len(queue) > 0 {
        node_count := len(queue)
        curr_depth := []int{}

        for node_count > 0 {
            curr := queue[0]
            curr_depth = append(curr_depth, curr.Val)

            if curr.Left != nil {
                queue = append(queue, curr.Left)
            }

            if curr.Right != nil {
                queue = append(queue, curr.Right)
            }

            queue = queue[1:]
            node_count -= 1
        }

        res = append(res, curr_depth)
    }

    return res;
}

func main() {
    tree := TreeNode {
        3,
        &TreeNode {
            9,
            &TreeNode {
                16,
                nil,
                nil,
            },
            nil,
        },
        &TreeNode {
            20,
            nil,
            &TreeNode {
                8,
                nil,
                nil,
            },
        },
    }
    fmt.Printf("%v", levelOrder(&tree))
}