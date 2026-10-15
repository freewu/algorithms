package main

// LCR 043. 完全二叉树插入器
// 完全二叉树是每一层（除最后一层外）都是完全填充（即，节点数达到最大，第 n 层有 2n-1 个节点）的，并且所有的节点都尽可能地集中在左侧。

// 设计一个用完全二叉树初始化的数据结构 CBTInserter，它支持以下几种操作：
//     CBTInserter(TreeNode root) 
//         使用根节点为 root 的给定树初始化该数据结构；
//     CBTInserter.insert(int v)  
//         向树中插入一个新节点，节点类型为 TreeNode，值为 v 。
//         使树保持完全二叉树的状态，并返回插入的新节点的父节点的值；
//     CBTInserter.get_root() 
//         将返回树的根节点。

// 示例 1：
// 输入：inputs = ["CBTInserter","insert","get_root"], inputs = [[[1]],[2],[]]
// 输出：[null,1,[1,2]]

// 示例 2：
// 输入：inputs = ["CBTInserter","insert","insert","get_root"], inputs = [[[1,2,3,4,5,6]],[7],[8],[]]
// 输出：[null,3,4,[1,2,3,4,5,6,7,8]]

// 提示：
//     最初给定的树是完全二叉树，且包含 1 到 1000 个节点。
//     每个测试用例最多调用 CBTInserter.insert  操作 10000 次。
//     给定节点或插入节点的每个值都在 0 到 5000 之间。

import "fmt"
import "math/bits"

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
type CBTInserter struct { 
    treeArr []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
    return CBTInserter{treeArr: bfs(root)} 
}
  
func bfs(root *TreeNode) []*TreeNode {
    queue, res := []*TreeNode{}, []*TreeNode{}
    queue = append(queue, root)
    for {
        node := queue[0]
        queue = queue[1:]
        res = append(res, node)
        if node.Left != nil {
            queue = append(queue, node.Left)   
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
        if len(queue) == 0 {
            break
        }
    }
    return res 
}

func (this *CBTInserter) Insert(v int) int {
    length := len(this.treeArr)
    insert_loc := ((float64(length) - 1) / 2.0)
    loc := int64(insert_loc)
    parent := this.treeArr[loc]
    node := &TreeNode{Val: v}
    if parent.Left == nil {
        parent.Left = node
    } else {
        parent.Right = node
    }
    this.treeArr = append(this.treeArr, node)
    return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
    return this.treeArr[0] 
}

type CBTInserter1 struct {
    root *TreeNode
    count  int
}

func Constructor1(root *TreeNode) CBTInserter1 {
    queue := []*TreeNode{root}
    count := 0
    for len(queue) > 0 { // 得到树的节点个数
        count++
        node := queue[0]
        queue = queue[1:]
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
    return CBTInserter1{root, count}
}

func (this *CBTInserter1) Insert(val int) int {
    this.count++
    child := &TreeNode{ Val: val }
    node := this.root
    for i := bits.Len(uint(this.count)) - 2; i > 0; i-- {
        if this.count >> i & 1 == 0 {
            node = node.Left
        } else {
            node = node.Right
        }
    }
    if this.count & 1 == 0 {
        node.Left = child
    } else {
        node.Right = child
    }
    return node.Val
}

func (this *CBTInserter1) Get_root() *TreeNode {
    return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */

func main() {
    // CBTInserter cBTInserter = new CBTInserter([1, 2]);
    tree1 := buildTree([]*int{ptr(1), ptr(2)})
    printTree(tree1)
    obj := Constructor(tree1)
    fmt.Println(obj)
    // cBTInserter.insert(3);  // return 1
    fmt.Println(obj.Insert(3)) // 1
    fmt.Println(obj)
    // cBTInserter.insert(4);  // return 2
    fmt.Println(obj.Insert(4)) // 2
    fmt.Println(obj)
    // cBTInserter.get_root(); // return [1, 2, 3, 4]
    res1 := obj.Get_root()
    printTree(res1)
    fmt.Println(res1) // [1, 2, 3, 4]

    tree11 := buildTree([]*int{ptr(1), ptr(2)})
    printTree(tree11)
    obj1 := Constructor1(tree11)
    fmt.Println(obj1)
    // cBTInserter.insert(3);  // return 1
    fmt.Println(obj1.Insert(3)) // 1
    fmt.Println(obj1)
    // cBTInserter.insert(4);  // return 2
    fmt.Println(obj1.Insert(4)) // 2
    fmt.Println(obj1)
    // cBTInserter.get_root(); // return [1, 2, 3, 4]
    res2 := obj1.Get_root()
    printTree(res2)
    fmt.Println(res2) // [1, 2, 3, 4]
}