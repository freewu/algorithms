package main

// LCR 048. 二叉树的序列化与反序列化
// 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，
// 同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

// 请设计一个算法来实现二叉树的序列化与反序列化。
// 这里不限定你的序列 / 反序列化算法执行逻辑，只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2020/09/15/serdeser.jpg" />
// 输入：root = [1,2,3,null,null,4,5]
// 输出：[1,2,3,null,null,4,5]

// 示例 2：
// 输入：root = []
// 输出：[]

// 示例 3：
// 输入：root = [1]
// 输出：[1]

// 示例 4：
// 输入：root = [1,2]
// 输出：[1,2]
 
// 提示：
//     输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，也可以采用其他的方法解决这个问题。
//     树中结点数在范围 [0, 10^4] 内
//     -1000 <= Node.val <= 1000
import "fmt"
import "bytes"
import "strconv"
import "strings"

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

type Codec struct {

}
 
func Constructor() Codec {
    return Codec{}
}

// dfs
func (this *Codec) serialize(root *TreeNode) string {
    var buffer bytes.Buffer
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            buffer.WriteString("#,")
        } else {
            buffer.WriteString(strconv.Itoa(node.Val))
            buffer.WriteString(",")
            dfs(node.Left)
            dfs(node.Right)
        }
    }
    dfs(root)
    return buffer.String()
}

// dfs
func (this *Codec) deserialize(data string) *TreeNode {    
    tokens := strings.Split(data, ",")
    var dfs func() *TreeNode
    dfs = func() *TreeNode {
        token := tokens[0]
        tokens = tokens[1:]
        if token == "#" { return nil }
        val, _ := strconv.Atoi(token)
        return &TreeNode{val, dfs(), dfs() }
    }
    return dfs()
}
 
 /**
  * Your Codec object will be instantiated and called as such:
  * ser := Constructor();
  * deser := Constructor();
  * data := ser.serialize(root);
  * ans := deser.deserialize(data);
  */
// bfs
type Codec1 struct {
}

func Constructor1() Codec1 {
	return Codec1{}
}

// Serializes a tree to a single string.
func (this *Codec1) serialize(root *TreeNode) string {
    res := ""
    if root == nil {
        return res
    }
    queue := []*TreeNode{}
    queue = append(queue, root)
    for len(queue) > 0 {
        temp := queue[0]
        queue = queue[1:]
        if temp == nil {
            res += "#,"
        } else {
            res += strconv.Itoa(temp.Val) + ","
        }
        if temp != nil{
            queue = append(queue, temp.Left)
            queue = append(queue, temp.Right)
        }
    }
    return res
}

// Deserializes your encoded data to tree.
func (this *Codec1) deserialize(data string) *TreeNode {
    if data == "" {
        return nil
    }
    nodes := strings.Split(data, ",")
    rootVal, _ := strconv.Atoi(nodes[0])
    root, queue, index := &TreeNode{Val: rootVal}, []*TreeNode{}, 1
    queue = append(queue, root)

    for len(queue) > 0 {
        temp := queue[0]
        queue = queue[1:]
        if nodes[index] != "#" {
            val, _ := strconv.Atoi(nodes[index])
            temp.Left = &TreeNode{Val: val}
            queue = append(queue, temp.Left)
        }
        index++
        if nodes[index] != "#" {
            val, _ := strconv.Atoi(nodes[index])
            temp.Right = &TreeNode{Val: val}
            queue = append(queue, temp.Right)
        }
        index++
    }
    return root
}

func main() {
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3), nil, nil, ptr(4), ptr(5)})
    printTree(tree1)
    obj := Constructor()
    data := obj.serialize(tree1)
    fmt.Println(data)
    t := obj.deserialize(data)
    printTree(t)
    fmt.Println(t.Val)
    fmt.Println(t.Left.Val)
    fmt.Println(t.Right.Val)

    obj1 := Constructor1()
    data1 := obj1.serialize(tree1)
    fmt.Println(data1)
    t1 := obj1.deserialize(data1)
    printTree(t1)
    fmt.Println(t1.Val)
    fmt.Println(t1.Left.Val)
    fmt.Println(t1.Right.Val)
}