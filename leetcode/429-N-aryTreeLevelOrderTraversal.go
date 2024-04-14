package main

// 429. N-ary Tree Level Order Traversal
// Given an n-ary tree, return the level order traversal of its nodes' values.
// Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png" />
// Input: root = [1,null,3,2,4,null,5,6]
// Output: [[1],[3,2,4],[5,6]]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: [[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
 
// Constraints:
//      The height of the n-ary tree is less than or equal to 1000
//      The total number of nodes is between [0, 10^4]

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
import "fmt"

type Node struct {
    Val      int
    Children []*Node
}

// 层序遍历 BFS
func levelOrder(root *Node) [][]int {
    var res [][]int
    var temp []int
    if root == nil {
        return res
    }
    queue := []*Node{root, nil}
    for len(queue) > 1 {
        node := queue[0]
        queue = queue[1:]
        if node == nil {
            queue = append(queue, nil)
            res = append(res, temp)
            temp = []int{}
        } else {
            temp = append(temp, node.Val)
            if len(node.Children) > 0 {
                queue = append(queue, node.Children...)
            }
        }
    }
    res = append(res, temp)
    return res
}

// bfs
func levelOrder1(root *Node) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    q := []*Node{root}
    for q != nil {
        lev := []int{}
        tmp := q
        q = nil
        for _, elem := range tmp {
            lev = append(lev, elem.Val)
            for _, child := range elem.Children {
                q = append(q, child)
            }
        }
        res = append(res, lev)
    }
    return res
}

func main() {
    tree1 := &Node{
        1,
        []*Node{
            &Node{
                3,
                []*Node{
                    &Node{5,nil},
                    &Node{6,nil},
                },
            },
            &Node{2,nil},
            &Node{4,nil},
        },
    }
    tree2 := &Node{
        1,
        []*Node{
            &Node{ 2,nil},
            &Node{ 3, 
                []*Node{ &Node{6,nil},
                    &Node{ 7, 
                        []*Node{ &Node{ 11, []*Node{ &Node{14,nil}, }, },
                        },
                    },
                },
            },
            &Node{ 4, 
                []*Node{ &Node{ 8, []*Node{ &Node{12,nil}, }, }, },
            },
            &Node{ 5,
                []*Node{
                    &Node{ 9,
                        []*Node{
                            &Node{13,nil},
                        },
                    },
                    &Node{10,nil},
                },
            },
        },
    }
    fmt.Println(levelOrder(tree1)) // [[1],[3,2,4],[5,6]]
    fmt.Println(levelOrder(tree2)) // [[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]

    fmt.Println(levelOrder1(tree1)) // [[1],[3,2,4],[5,6]]
    fmt.Println(levelOrder1(tree2)) // [[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
}
