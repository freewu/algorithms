package main

// 536. Construct Binary Tree from String
// You need to construct a binary tree from a string consisting of parenthesis and integers.

// The whole input represents a binary tree. 
// It contains an integer followed by zero, one or two pairs of parenthesis. 
// The integer represents the root's value and a pair of parenthesis contains a child binary tree with the same structure.

// You always start to construct the left child node of the parent first if it exists.

// Example 1:
// <img src="" />
// Input: s = "4(2(3)(1))(6(5))"
// Output: [4,2,6,3,1,5]

// Example 2:
// Input: s = "4(2(3)(1))(6(5)(7))"
// Output: [4,2,6,3,1,5,7]

// Example 3:
// Input: s = "-4(2(3)(1))(6(5)(7))"
// Output: [-4,2,6,3,1,5,7]

// Constraints:
//     0 <= s.length <= 3 * 10^4
//     s consists of digits, '(', ')', and '-' only.

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
// dfs
func str2tree(s string) *TreeNode {
    var dfs func(s string, l,r int) *TreeNode
    dfs = func(s string, l,r int) *TreeNode{
        if r < l {
            return nil
        }
        node, flag := &TreeNode{}, 1
        for l <= r {
            if s[l]=='-' { // 处理 - 号
                flag = -1
            } else if s[l] != '(' {
                node.Val = node.Val * 10 + int(s[l]-'0')
            } else {
                left, p := 1, l
                for left > 0 {
                    if s[l+1] == '('{
                        left++
                    }else if s[l+1]== ')'{
                        left--
                    }
                    l++
                }
                node.Left = dfs(s, p+1, l-1)
                node.Right = dfs(s, l+2, r-1)
                break
            }
            l++
        }
        node.Val *= flag
        return node
    }
    return dfs(s, 0, len(s)-1)
}

// stack bfs
func str2tree1(s string) *TreeNode {
    if len(s) == 0 {
        return nil
    }
    stack := make([]*TreeNode, 0)
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            continue
        }
        if s[i]== ')' { // 需要出栈了
            stack = stack[:len(stack)-1]
            continue
        }
        j, sign, num := i, 1, 0
        if s[j] == '-' {
            sign = -1
            j++
        }
        for ; j <len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
            num = num * 10 + int(s[j]-'0')
        }
        i, num = j-1, num * sign
        node:= &TreeNode { Val: num, }
        if len(stack) > 0 {
            tmp := stack[len(stack)-1]
            if tmp.Left == nil {
                tmp.Left = node
            } else {
                tmp.Right = node
            }
        }
        stack = append(stack, node)
    }
    return stack[len(stack)-1]
}

func main() {
    // Example 1:
    // <img src="" />
    // Input: s = "4(2(3)(1))(6(5))"
    // Output: [4,2,6,3,1,5]
    fmt.Println(str2tree("4(2(3)(1))(6(5))"))
    // Example 2:
    // Input: s = "4(2(3)(1))(6(5)(7))"
    // Output: [4,2,6,3,1,5,7]
    fmt.Println(str2tree("4(2(3)(1))(6(5)(7))"))
    // Example 3:
    // Input: s = "-4(2(3)(1))(6(5)(7))"
    // Output: [-4,2,6,3,1,5,7]
    fmt.Println(str2tree("-4(2(3)(1))(6(5)(7))"))

    fmt.Println(str2tree1("4(2(3)(1))(6(5))"))
    fmt.Println(str2tree1("4(2(3)(1))(6(5)(7))"))
    fmt.Println(str2tree1("-4(2(3)(1))(6(5)(7))"))
}