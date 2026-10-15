package main

// 449. Serialize and Deserialize BST
// Serialization is converting a data structure or object into a sequence of bits so 
// that it can be stored in a file or memory buffer, 
// or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

// Design an algorithm to serialize and deserialize a binary search tree. 
// There is no restriction on how your serialization/deserialization algorithm should work. 
// You need to ensure that a binary search tree can be serialized to a string, 
// and this string can be deserialized to the original tree structure.

// The encoded string should be as compact as possible.

// Example 1:
// Input: root = [2,1,3]
// Output: [2,1,3]

// Example 2:
// Input: root = []
// Output: []
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     0 <= Node.val <= 10^4
//     The input tree is guaranteed to be a binary search tree.

import "fmt"
import "strings"
import "container/list"
import "strconv"
import "sort"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// == 二叉树构建 & 打印输出 start ==================================================================
// func ptr(v int) *int {
//     return &v
// }

// // buildTree 按层序数组（LeetCode 格式，nil 表示 null）创建二叉树
// func buildTree(nums []*int) *TreeNode {
//     if len(nums) == 0 || nums[0] == nil {
//         return nil
//     }
//     root := &TreeNode{Val: *nums[0]}
//     queue := []*TreeNode{root}
//     i := 1
//     for len(queue) > 0 && i < len(nums) {
//         node := queue[0]
//         queue = queue[1:]
//         if nums[i] != nil {
//             node.Left = &TreeNode{Val: *nums[i]}
//             queue = append(queue, node.Left)
//         }
//         i++
//         if i < len(nums) && nums[i] != nil {
//             node.Right = &TreeNode{Val: *nums[i]}
//             queue = append(queue, node.Right)
//         }
//         i++
//     }
//     return root
// }

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

type Codec1 struct {

}
 
func Constructor1() Codec1 {
    return Codec1{}
}

// Serializes a tree to a single string.
func (this *Codec1) serialize(root *TreeNode) string {
    res, queue := "", list.New()
    if root != nil {
        queue.PushBack(root)
    }
    for queue.Len() > 0 {
        qSize := queue.Len()
        for i := 0; i < qSize; i++ {
            current := queue.Remove(queue.Front()).(*TreeNode)
            if res != "" {
                res += ","
            }
            if current != nil {
                res += fmt.Sprint(current.Val)
                queue.PushBack(current.Left)
                queue.PushBack(current.Right)
            } else {
                res += "null"
            }
        }
    }
    return res
}

// Deserializes your encoded data to tree.
func (this *Codec1) deserialize(data string) *TreeNode {
    if data == "" {
        return nil
    }
    values := strings.Split(data, ",")
    nodes := []*TreeNode{}
    for _, value := range values {
        if value == "null" {
            nodes = append(nodes, nil)
        } else {
            val, _ := strconv.Atoi(value)
            nodes = append(nodes, &TreeNode{Val: val})
        }
    }
    i := 0
    for _, node := range nodes {
        if node == nil {
            continue
        }
        left := 2 * i + 1
        right := 2 * i + 2
        if left < len(nodes) {
            node.Left = nodes[left]
        }
        if right < len(nodes) {
            node.Right = nodes[right]
        }
        i++
    }
    return nodes[0]
}


type Codec struct {
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    arr := []string{}
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        arr = append(arr, strconv.Itoa(root.Val))
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)
    return strings.Join(arr, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
    if data == "" {
        return nil
    }
    arr := strings.Split(data, ",")
    pre := make([]int, len(arr))
    mid := make([]int, len(arr))
    for i, v := range arr {
        value, _ := strconv.Atoi(v)
        pre[i] = value
        mid[i] = value
    }
    sort.Ints(mid)
    return buildTree(pre, mid)
}

func buildTree(pre, mid []int) *TreeNode {
    if len(pre) == 0 {
        return nil
    }
    v := pre[0]
    index := 0
    for i := 0; i < len(pre); i++ {
        if mid[i] == v {
            index = i
            break
        }
    }
    return &TreeNode{
        Val:   v,
        Left:  buildTree(pre[1:index+1], mid[:index]),
        Right: buildTree(pre[index+1:], mid[index+1:]),
    }
}

 /**
  * Your Codec object will be instantiated and called as such:
  * ser := Constructor()
  * deser := Constructor()
  * tree := ser.serialize(root)
  * ans := deser.deserialize(tree)
  * return ans
  */

func main() {
    obj := Constructor()
    // Example 1:
    // Input: root = [2,1,3]
    // Output: [2,1,3]
    tree1 := &TreeNode {
        2,
        &TreeNode{1, nil, nil},
        &TreeNode{3, nil, nil},
    }
    printTree(tree1)
    e1 := obj.serialize(tree1)
    fmt.Println("e1 ", e1)
    d1 := obj.deserialize(e1)
    fmt.Println("d1 ", d1)
    fmt.Println("d1.Left ", d1.Left)
    fmt.Println("d1.Right ", d1.Right)
    // Example 2:
    // Input: root = []
    // Output: []
    tree2 := &TreeNode {}
    printTree(tree2)
    e2 := obj.serialize(tree2)
    fmt.Println("e2 ", e2)
    d2 := obj.deserialize(e2)
    fmt.Println("d2 ", d2)

    obj1 := Constructor()
    // Example 1:
    // Input: root = [2,1,3]
    // Output: [2,1,3]
    tree11 := &TreeNode {
        2,
        &TreeNode{1, nil, nil},
        &TreeNode{3, nil, nil},
    }
    e11 := obj1.serialize(tree11)
    fmt.Println("e11 ", e11)
    d11 := obj1.deserialize(e11)
    fmt.Println("d11 ", d11)
    fmt.Println("d11.Left ", d11.Left)
    fmt.Println("d11.Right ", d11.Right)
    // Example 2:
    // Input: root = []
    // Output: []
    tree12 := &TreeNode {}
    e12 := obj1.serialize(tree12)
    fmt.Println("e12 ", e12)
    d12 := obj1.deserialize(e12)
    fmt.Println("d12 ", d12)
}