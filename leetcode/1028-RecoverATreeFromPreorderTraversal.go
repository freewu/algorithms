package main

// 1028. Recover a Tree From Preorder Traversal
// We run a preorder depth-first search (DFS) on the root of a binary tree.

// At each node in this traversal, we output D dashes (where D is the depth of this node), 
// then we output the value of this node.  
// If the depth of a node is D, the depth of its immediate child is D + 1.  
// The depth of the root node is 0.

// If a node has only one child, that child is guaranteed to be the left child.

// Given the output traversal of this traversal, recover the tree and return its root.

// Example 1:
//             1
//           /   \
//          2     5
//        /   \  /  \
//       3    4 6    7
// <img src="https://assets.leetcode.com/uploads/2019/04/08/recover-a-tree-from-preorder-traversal.png" />
// Input: traversal = "1-2--3--4-5--6--7"
// Output: [1,2,5,3,4,6,7]

// Example 2:
//             1
//          /     \
//         2       5
//        /       /
//       3       6
//      /       /
//     4       7
// <img src="https://assets.leetcode.com/uploads/2019/04/11/screen-shot-2019-04-10-at-114101-pm.png" />
// Input: traversal = "1-2--3---4-5--6---7"
// Output: [1,2,5,3,null,6,null,4,null,7]

// Example 3:
//             1
//           /
//         401
//       /    \
//     349    88
//     /
//    90
// <img src="https://assets.leetcode.com/uploads/2019/04/11/screen-shot-2019-04-10-at-114955-pm.png" />
// Input: traversal = "1-401--349---90--88"
// Output: [1,401,null,349,88,90]
 
// Constraints:
//     The number of nodes in the original tree is in the range [1, 1000].
//     1 <= Node.val <= 10^9

import "fmt"
import "strconv"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(traversal string) *TreeNode {
    parent := make(map[*TreeNode]*TreeNode) // 用 map 存当前节点的父节点
    res := &TreeNode{ Val: -1, Left: nil, Right: nil, } // 定义一个虚拟节点
    parent[res] = nil
    cur, parentLevel, curLevel, p, n := res,  -1, 0, 0, len(traversal)
    for p < n {
        for ; traversal[p] == '-'; p++ {
            curLevel++
        }
        start := p
        for ; p < n && traversal[p] != '-'; p++ { }
        curVal, _ := strconv.Atoi(traversal[start:p])
        if curLevel > parentLevel {
            cur.Left = &TreeNode{ Val: curVal, Left: nil, Right: nil, }
            parentLevel = curLevel
            curLevel = 0
            parent[cur.Left] = cur
            cur = cur.Left
        } else {
            pa := cur
            for curLevel <= parentLevel {
                pa = parent[pa]
                parentLevel--
            }
            pa.Right = &TreeNode{ Val: curVal, Left: nil, Right: nil, }
            parent[pa.Right] = pa
            cur = pa.Right
            parentLevel = curLevel
        }
        curLevel = 0
    }
    return res.Left
}

// stack
func recoverFromPreorder1(traversal string) *TreeNode {
    type Item struct {
        node *TreeNode
        level int
    }
    n, d, index := len(traversal), 0, 0
    for ; index < n && traversal[index] >= '0' && traversal[index] <= '9'; index++ { // 先取第1个值
        ch := traversal[index]
        d = d * 10 + int(ch - '0')
    }
    root := &TreeNode{}
    root.Val = d
    stack := []Item{ Item{node: root, level: 0}}
    for index < n {
        d = 0
        level := 0
        for ; index < n && traversal[index] == '-'; index++ {
            level++
        }
        for ; index < n && traversal[index] != '-'; index++ {
            d = d * 10+ int(traversal[index] - '0')
        }
        tmpNode := &TreeNode{ Val : d }
        for len(stack) > 0 && stack[len(stack)-1].level != level-1 {
            stack = stack[:len(stack)-1]
        }
        parent := stack[len(stack)-1].node
        if parent.Left == nil {
            parent.Left = tmpNode
        } else {
            parent.Right = tmpNode
        }
        stack = append(stack, Item{node: tmpNode, level: level})
    }
    return root
}

func main() {
    // Example 1:
    //             1
    //           /   \
    //          2     5
    //        /   \  /  \
    //       3    4 6    7
    // <img src="https://assets.leetcode.com/uploads/2019/04/08/recover-a-tree-from-preorder-traversal.png" />
    // Input: traversal = "1-2--3--4-5--6--7"
    // Output: [1,2,5,3,4,6,7]
    fmt.Println(recoverFromPreorder("1-2--3--4-5--6--7")) // &{1 0xc000110060 0xc0001100a8}
    // Example 2:
    //             1
    //          /     \
    //         2       5
    //        /       /
    //       3       6
    //      /       /
    //     4       7
    // <img src="https://assets.leetcode.com/uploads/2019/04/11/screen-shot-2019-04-10-at-114101-pm.png" />
    // Input: traversal = "1-2--3---4-5--6---7"
    // Output: [1,2,5,3,null,6,null,4,null,7]
    fmt.Println(recoverFromPreorder("1-2--3---4-5--6---7")) // &{1 0xc000110138 0xc000110180}
    // Example 3:
    //             1
    //           /
    //         401
    //       /    \
    //     349    88
    //     /
    //    90
    // <img src="https://assets.leetcode.com/uploads/2019/04/11/screen-shot-2019-04-10-at-114955-pm.png" />
    // Input: traversal = "1-401--349---90--88"
    // Output: [1,401,null,349,88,90]
    fmt.Println(recoverFromPreorder("1-401--349---90--88")) // &{1 0xc000110210 <nil>}

    fmt.Println(recoverFromPreorder1("1-2--3--4-5--6--7")) // &{1 0xc000110060 0xc0001100a8}
    fmt.Println(recoverFromPreorder1("1-2--3---4-5--6---7")) // &{1 0xc000110138 0xc000110180}
    fmt.Println(recoverFromPreorder1("1-401--349---90--88")) // &{1 0xc000110210 <nil>}
}