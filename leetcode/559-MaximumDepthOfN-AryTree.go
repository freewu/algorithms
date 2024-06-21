package main

// 559. Maximum Depth of N-ary Tree
// Given a n-ary tree, find its maximum depth.
// The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
// Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png" />
// Input: root = [1,null,3,2,4,null,5,6]
// Output: 3

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: 5
 
// Constraints:
//     The total number of nodes is in the range [0, 10^4].
//     The depth of the n-ary tree is less than or equal to 1000.

import "fmt"

type Node struct {
    Val      int
    Children []*Node
}

// dfs
func maxDepth(root *Node) int {
    if root ==  nil {
        return 0
    } 
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(root *Node, level int) int 
    dfs = func(root *Node, level int) int {
        res := level + 1
        if root == nil || len(root.Children) == 0 {
            return res
        }
        for _, v := range root.Children {
            res = max(res,dfs(v,level + 1))
        }
        return res
    }
    return dfs(root,0)
}

// bfs
func maxDepth1(root *Node) int {
    if root == nil {
        return 0
    }
    depth, queue := 0, make([]*Node,0)
    queue = append(queue,root) // 根节点入队
    for len(queue) != 0 {
        num := len(queue)  // 记录当前层个数
        // 层序遍历，下一层节点进入队列
        for i := 0; i< num; i++ {
            for j:= 0; j < len(queue[i].Children); j++ {
                queue = append(queue, queue[i].Children[j])
            }
        }
        queue = queue[num:len(queue)] // 去掉当前层，深度+1
        depth++
    }
    return depth
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png" />
    // Input: root = [1,null,3,2,4,null,5,6]
    // Output: 3
    tree1 := &Node{
        1,
        []*Node{
            &Node{ 3, []*Node{ &Node{5,nil}, &Node{6,nil}, }, },
            &Node{ 2, nil},
            &Node{ 4, nil},
        },
    }
    fmt.Println(maxDepth(tree1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />
    // Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    // Output: 5
    tree2 := &Node{
        1,
        []*Node {
            &Node{2, nil },
            &Node{3, []*Node{ &Node{6, nil }, &Node{7, []*Node{ &Node{ 11, []*Node{ &Node{14, nil }} }, } }, }  },
            &Node{4, []*Node{ &Node{8, []*Node{ &Node{12, nil }}  }, } },
            &Node{5, []*Node{ &Node{9, []*Node{ &Node{13, nil }}  }, &Node{10, nil } } },
        },
    }
    fmt.Println(maxDepth(tree2)) // 5

    fmt.Println(maxDepth1(tree1)) // 3
    fmt.Println(maxDepth1(tree2)) // 5
}