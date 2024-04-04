package main

// 1609. Even Odd Tree
// A binary tree is named Even-Odd if it meets the following conditions:
// 	    The root of the binary tree is at level index 0, its children are at level index 1, their children are at level index 2, etc.
// 	    For every even-indexed level, all nodes at the level have odd integer values in strictly increasing order (from left to right).
// 	    For every odd-indexed level, all nodes at the level have even integer values in strictly decreasing order (from left to right).

// Given the root of a binary tree, return true if the binary tree is Even-Odd, otherwise return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/15/sample_1_1966.png" />
// Input: root = [1,10,4,3,null,7,9,12,8,6,null,null,2]
// Output: true
// Explanation: The node values on each level are:
// Level 0: [1]
// Level 1: [10,4]
// Level 2: [3,7,9]
// Level 3: [12,8,6,2]
// Since levels 0 and 2 are all odd and increasing and levels 1 and 3 are all even and decreasing, the tree is Even-Odd.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/09/15/sample_2_1966.png" />
// Input: root = [5,4,2,3,3,7]
// Output: false
// Explanation: The node values on each level are:
// Level 0: [5]
// Level 1: [4,2]
// Level 2: [3,3,7]
// Node values in level 2 must be in strictly increasing order, so the tree is not Even-Odd.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/09/22/sample_1_333_1966.png" />
// Input: root = [5,9,1,3,5,7]
// Output: false
// Explanation: Node values in the level 1 should be even integers.
 
// Constraints:
//         The number of nodes in the tree is in the range [1, 10^5].
//         1 <= Node.val <= 10^6

import "fmt"

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
func isEvenOddTree(root *TreeNode) bool {
    level := 0
    queue := []*TreeNode{root}
    
    odd := func (nums []int) bool {
        cur := nums[0]
        if cur%2 != 0 {
            return false
        }
        for _, num := range nums[1:] {
            if num >= cur || num%2 != 0 {
                return false
            }
            cur = num
        }
        return true
    }
    even := func (nums []int) bool {
        cur := nums[0]
        if cur%2 == 0 {
            return false
        }
        for _, num := range nums[1:] {
            if num <= cur || num%2 == 0 {
                return false
            }
            cur = num
        }
        return true
    }
    for len(queue) != 0 {
        length := len(queue)
        var nums []int
        for i := 0; i < length; i++ {
            node := queue[i]
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            nums = append(nums, node.Val)
        }
        if level % 2 == 0 {
            if !even(nums) {
                return false
            }
        } else {
            if !odd(nums) {
                return false
            }
        }
        queue = queue[length:]
        level++
    }
    return true
}

// best solution
func isEvenOddTree1(root *TreeNode) bool {
    q := []*TreeNode{root}
    level := -1
    for len(q) > 0 {
        level++
        size := len(q)
        pre := -1
        for i := 0; i < size; i++ {
            node := q[0]
            q = q[1:]
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
            if level % 2 == 0 {
                if node.Val % 2 == 0 || node.Val <= pre {
                    return false
                }
            } else {
                if node.Val % 2 == 1 || i > 0 && node.Val >= pre {
                    return false
                }
            }
            pre = node.Val
        }
    }
    return true
}

func main() {
    fmt.Println(isEvenOddTree(
        &TreeNode {
            1,
            &TreeNode {
                10,
                &TreeNode{
                    3, 
                    &TreeNode{12, nil, nil},
                    &TreeNode{8, nil, nil},
                },
                nil,
            },
            &TreeNode {
                4,
                &TreeNode{
                    7, 
                    &TreeNode{6, nil, nil},
                    nil,
                },
                &TreeNode{
                    9, 
                    nil, 
                    &TreeNode{2, nil, nil},
                },
            },
        },
    )) // true

    fmt.Println(isEvenOddTree(
        &TreeNode {
            5,
            &TreeNode {
                4,
                &TreeNode{3, nil, nil},
                &TreeNode{3, nil, nil},
            },
            &TreeNode {
                2,
                &TreeNode{7, nil, nil},
                nil,
            },
        },
    )) // false

    fmt.Println(isEvenOddTree(
        &TreeNode {
            5,
            &TreeNode {
                9,
                &TreeNode{3, nil, nil},
                &TreeNode{5, nil, nil},
            },
            &TreeNode {
                1,
                &TreeNode{7, nil, nil},
                nil,
            },
        },
    )) // false

    fmt.Println(isEvenOddTree1(
        &TreeNode {
            1,
            &TreeNode {
                10,
                &TreeNode{
                    3, 
                    &TreeNode{12, nil, nil},
                    &TreeNode{8, nil, nil},
                },
                nil,
            },
            &TreeNode {
                4,
                &TreeNode{
                    7, 
                    &TreeNode{6, nil, nil},
                    nil,
                },
                &TreeNode{
                    9, 
                    nil, 
                    &TreeNode{2, nil, nil},
                },
            },
        },
    )) // true

    fmt.Println(isEvenOddTree1(
        &TreeNode {
            5,
            &TreeNode {
                4,
                &TreeNode{3, nil, nil},
                &TreeNode{3, nil, nil},
            },
            &TreeNode {
                2,
                &TreeNode{7, nil, nil},
                nil,
            },
        },
    )) // false

    fmt.Println(isEvenOddTree1(
        &TreeNode {
            5,
            &TreeNode {
                9,
                &TreeNode{3, nil, nil},
                &TreeNode{5, nil, nil},
            },
            &TreeNode {
                1,
                &TreeNode{7, nil, nil},
                nil,
            },
        },
    )) // false
}