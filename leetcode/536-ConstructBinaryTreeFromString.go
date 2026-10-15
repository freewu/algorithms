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

// == 二叉树构建 & 打印输出 start ==================================================================
func ptr(v int) *int {
    return &v
}

// buildTree 按层序数组（LeetCode 格式，nil 表示 null）创建二叉树
func buildTree(nums []*int) *TreeNode {
    if len(nums) == 0 || nums[0] == nil {
        return nil
    }
    root := &TreeNode{Val: *nums[0]}
    queue := []*TreeNode{root}
    i := 1
    for len(queue) > 0 && i < len(nums) {
        node := queue[0]
        queue = queue[1:]
        if nums[i] != nil {
            node.Left = &TreeNode{Val: *nums[i]}
            queue = append(queue, node.Left)
        }
        i++
        if i < len(nums) && nums[i] != nil {
            node.Right = &TreeNode{Val: *nums[i]}
            queue = append(queue, node.Right)
        }
        i++
    }
    return root
}

// printTree 自顶向下打印二叉树，父子之间用 / \ 连线
// 第 r 层（根为第 1 层）有 2^(r-1) 个位置，每个节点占 block 列，
// 因此节点中心列 = (2j+1) * 2^(h-r) * block / 2（j 为该层从左往右的序号）
func printTree(root *TreeNode) {
    if root == nil {
        fmt.Println("(empty tree)")
        return
    }
    // itoa 手写整数转字符串，替代 strconv.Itoa
    itoa := func(v int) string {
        if v == 0 {
            return "0"
        }
        neg := v < 0
        u := uint(v)
        if neg {
            u = -u // 补码取绝对值，v 为最小整数时同样安全
        }
        buf := [24]byte{}
        i := len(buf)
        for u > 0 {
            i--
            buf[i] = byte('0' + u%10)
            u /= 10
        }
        if neg {
            i--
            buf[i] = '-'
        }
        return string(buf[i:])
    }
    // blankLine 生成 n 个空格的画布行（替代 strings.Repeat）
    blankLine := func(n int) []byte {
        line := make([]byte, n)
        for i := range line {
            line[i] = ' '
        }
        return line
    }
    // treeHeight 返回二叉树的层数（空树为 0，根节点为第 1 层）
    var treeHeight func(root *TreeNode) int
    treeHeight = func (root *TreeNode) int {
        if root == nil {
            return 0
        }
        l, r := treeHeight(root.Left), treeHeight(root.Right)
        if l < r {
            l = r
        }
        return l + 1
    }
    h := treeHeight(root)
    // block：每个节点占用的列宽，取最长值的宽度 + 1，并保证为偶数使中心落在整数列
    block := 2
    var scan func(node *TreeNode)
    scan = func(node *TreeNode) {
        if node == nil {
            return
        }
        if w := len(itoa(node.Val)) + 1; w > block {
            block = w
        }
        scan(node.Left)
        scan(node.Right)
    }
    scan(root)
    if block%2 != 0 {
        block++
    }

    // 画布：值占 1 行、连线占 1 行，共 2h-1 行
    width := (1 << h) * block
    rows := make([][]byte, 2*h-1)
    for i := range rows {
        rows[i] = blankLine(width)
    }
    put := func(row, col int, s string) {
        start := col - len(s)/2
        for i := 0; i < len(s); i++ {
            if p := start + i; p >= 0 && p < width {
                rows[row][p] = s[i]
            }
        }
    }
    var draw func(node *TreeNode, r, j int)
    draw = func(node *TreeNode, r, j int) {
        if node == nil {
            return
        }
        col := (2*j + 1) * (1 << (h - r)) * block / 2 // 当前节点中心列
        put(2*(r-1), col, itoa(node.Val))
        if r == h {
            return
        }
        d := (1 << (h - r - 1)) * block / 2 // 父子之间的列距
        off := (d + 1) / 2                  // 连线相对父节点中心的偏移（取行间中点）
        if node.Left != nil {
            put(2*r-1, col-off, "/")
            draw(node.Left, r+1, 2*j)
        }
        if node.Right != nil {
            put(2*r-1, col+off, "\\")
            draw(node.Right, r+1, 2*j+1)
        }
    }
    draw(root, 1, 0)

    for _, line := range rows {
        end := len(line)
        for end > 0 && line[end-1] == ' ' {
            end--
        }
        fmt.Println(string(line[:end]))
    }
}
// == 二叉树构建 & 打印输出 end ====================================================================

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