package main

// 1506. Find Root of N-Ary Tree
// You are given all the nodes of an N-ary tree as an array of Node objects, where each node has a unique value.
// Return the root of the N-ary tree.
// Custom testing:
//     An N-ary tree can be serialized as represented in its level order traversal where each group of children is separated by the null value (see examples).

// <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />
// For example, the above tree is serialized as [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14].

// The testing will be done in the following way:
//     The input data should be provided as a serialization of the tree.
//     The driver code will construct the tree from the serialized input data and put each Node object into an array in an arbitrary order.
//     The driver code will pass the array to findRoot, and your function should find and return the root Node object in the array.
//     The driver code will take the returned Node object and serialize it. If the serialized value and the input data are the same, the test passes.
    
// Example 1:
//            1
//        /   |   \ 
//       3    2    4
//      / \
//     5   6
// <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png" />
// Input: tree = [1,null,3,2,4,null,5,6]
// Output: [1,null,3,2,4,null,5,6]
// Explanation: The tree from the input data is shown above.
// The driver code creates the tree and gives findRoot the Node objects in an arbitrary order.
// For example, the passed array could be [Node(5),Node(4),Node(3),Node(6),Node(2),Node(1)] or [Node(2),Node(6),Node(1),Node(3),Node(5),Node(4)].
// The findRoot function should return the root Node(1), and the driver code will serialize it and compare with the input data.
// The input data and serialized Node(1) are the same, so the test passes.

// Example 2:
//           1
//      /   /  \   \
//     2    3   4   5
//         / \  |  / \
//        6   7 8  9  10
//            | |  |
//           11 12 13
//            |
//           14
// <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />
// Input: tree = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
 
// Constraints:
//     The total number of nodes is between [1, 5 * 10^4].
//     Each node has a unique value.

// Follow up:
//     Could you solve this problem in constant space complexity with a linear time algorithm?

import "fmt"

type Node struct {
    Val int
    Children []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func findRoot(tree []*Node) *Node {
    seen := make(map[int]bool) // 包含所有子节点的
    for _, node := range tree { // 将所有子节点添加到集合中
        for _, child := range node.Children {
            seen[child.Val] = true // 我们可以添加该值，也可以添加该节点本身。
        }
    }
    for _, node := range tree { // 查找不在子节点集中的节点
        if _, ok := seen[node.Val]; !ok {
            return node
        }
    }
    return nil
}

func findRoot1(tree []*Node) *Node {
    valueSum := 0
    for _, node := range tree {
        valueSum += node.Val // 该值作为父节点添加
        for _, child := range node.Children {
            valueSum -= child.Val // 该值将作为子节点扣除
        }
    }
    var root *Node
    for _, node := range tree {
        if node.Val == valueSum { // 根节点的值是 valueSum
            root = node
            break
        }
    }
    return root
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png" />
    // Input: tree = [1,null,3,2,4,null,5,6]
    // Output: [1,null,3,2,4,null,5,6]
    // Explanation: The tree from the input data is shown above.
    // The driver code creates the tree and gives findRoot the Node objects in an arbitrary order.
    // For example, the passed array could be [Node(5),Node(4),Node(3),Node(6),Node(2),Node(1)] or [Node(2),Node(6),Node(1),Node(3),Node(5),Node(4)].
    // The findRoot function should return the root Node(1), and the driver code will serialize it and compare with the input data.
    // The input data and serialized Node(1) are the same, so the test passes.
    tree1 := &Node{
        1,
        []*Node{
            &Node{
                3, 
                []*Node{
                    &Node{5, nil},
                    &Node{6, nil},
                }, 
            },
            &Node{2, nil },
            &Node{4, nil },
        },
    }
    fmt.Println(tree1) 
    //fmt.Println(findRoot(tree1)) // [1,null,3,2,4,null,5,6]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />
    // Input: tree = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    // Output: [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    tree2 := &Node{
        1,
        []*Node{
            &Node{2, nil },
            &Node{
                3,
                []*Node{
                    &Node{6, nil },
                    &Node{
                        7, 
                        []*Node{
                            &Node{
                                11, 
                                []*Node{
                                    &Node{14, nil },
                                },
                            },
                        },
                    },
                    
                },
            },
            &Node{
                4, 
                []*Node{
                    &Node{
                        8, 
                        []*Node{
                            &Node{12, nil },
                        },
                    },
                },
            },
            &Node{
                5,
                []*Node{
                    &Node{
                        9, 
                        []*Node{
                            &Node{13, nil },
                        },
                    },
                    &Node{10, nil },
                },
            },
        },
    }
    fmt.Println(tree2) 
    //fmt.Println(findRoot(tree2)) // [1,null,3,2,4,null,5,6]
}