package main

// 4082. Minimum Operations to Stabilize a Binary Tree
// You are given the root of a binary tree. 
// Each node contains an integer value representing its current state.

// A binary tree is considered stable if for every node u:
//     1. If u.left exists, then |u.val - u.left.val| <= 1
//     2. If u.right exists, then |u.val - u.right.val| <= 1

// In one operation, you may choose any node and increase or decrease its value by 1.

// Return the minimum number of operations required to make the entire binary tree stable.

// Example 1:
// Input: root = [5,8,3]
// Output: 3
// Explanation:
//         5
//      /     \
//     8       3
// <img src="https://assets.leetcode.com/uploads/2026/02/16/screenshot-2026-02-16-at-21240pm.png" style="width: 180px; height: 153px;">
// - Let u be the root node with value 5. Decrease the left child value from 8 to 6 using 2 operations. Then |u.val - u.left.val| = |5 - 6| = 1.
// - Increase the right child value from 3 to 4 using 1 operation. Then |u.val - u.right.val| = |5 - 4| = 1.
// - Thus, the answer is 2 + 1 = 3.

// Example 2:
// Input: root = [2,-1,-4]
// Output: 5
// Explanation:
//         2
//      /     \
//     -1       -4
// <img src="https://assets.leetcode.com/uploads/2026/02/16/screenshot-2026-02-16-at-21340pm.png" style="width: 170px; height: 148px;">
// - Let u be the root node with value 2. Decrease the value of u from 2 to −2 using 4 operations.
// - Decrease the right child value from −4 to −3 using 1 operation. Then |u.val - u.right.val| = |-2 - (−3)| = 1.
// - Thus, the answer is 4 + 1 = 5.

// Example 3:
// Input: root = [0,-2,2,-3,null,1]
// Output: 3
// Explanation:
//         0
//      /     \
//     -2       2
//   /       /
// -3       1
// <img src="https://assets.leetcode.com/uploads/2026/02/16/screenshot-2026-02-16-at-21826pm.png" style="width: 180px; height: 201px;">​​​​​​​​​​​​​​
// - Let u be the root node with value 0. Increase the left child value from −2 to −1 using 1 operation, and decrease the right child value from 2 to 1 using 1 operation.
// - Now |u.val - u.left.val| = |0 - (−1)| = 1 and |u.val - u.right.val| = |0 - 1| = 1.
// - Let v be the left child of u, now having value −1. Increase its left child from −3 to −2 using 1 operation so that |v.val - v.left.val| = |−1 - (−2)| = 1.
// - The right subtree already satisfies the condition since |1 - 1| = 0.
// - Thus, the answer is 1 + 1 + 1 = 3.​​​​​​​

// Constraints:
//     - The number of nodes in the tree is in the range [1, 5 * 10^{4}].
//     - -10^{5} <= Node.val <= 10^{5}

import "fmt"

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
type Node struct {
    val   int64
    tag   int64
    left  int
    right int
    dist  int32
}

type Pool struct {
    nodes []Node
}

func NewPool() *Pool {
    return &Pool{
        nodes: []Node{
            {val: 0, tag: 0, left: 0, right: 0, dist: -1}, // index 0 是空哨兵
        },
    }
}

func (p *Pool) Alloc(v int64) int {
    u := len(p.nodes)
    p.nodes = append(p.nodes, Node{
        val:   v,
        tag:   0,
        left:  0,
        right: 0,
        dist:  0,
    })
    return u
}

// Merge 合并堆 a,b；pa/pb 附加懒标记，is_max true=大根堆
func (p *Pool) Merge(a, b int, pa, pb int64, isMax bool) int {
    if a == 0 {
        if b != 0 {
            p.nodes[b].tag += pb
        }
        return b
    }
    if b == 0 {
        p.nodes[a].tag += pa
        return a
    }
    p.nodes[a].tag += pa
    p.nodes[b].tag += pb
    ra, rb := p.nodes[a].val + p.nodes[a].tag, p.nodes[b].val + p.nodes[b].tag
    u, v := 0, 0
    if (ra > rb) == isMax {
        u, v = a, b
    } else {
        u, v = b, a
    }
    ur, ut := p.nodes[u].right, p.nodes[u].tag
    nr0 := p.Merge(ur, v, ut, 0, isMax)
    if nr0 != 0 {
        p.nodes[nr0].tag -= ut
    }
    p.nodes[u].right = nr0
    l, r := p.nodes[u].left, p.nodes[u].right   
    if p.nodes[l].dist < p.nodes[r].dist {
        p.nodes[u].left = r
        p.nodes[u].right = l
    }
    rr := p.nodes[u].right
    p.nodes[u].dist = p.nodes[rr].dist + 1
    return u
}

func (p *Pool) Push(root int, v int64, isMax bool) int {
    u := p.Alloc(v)
    return p.Merge(root, u, 0, 0, isMax)
}

func (p *Pool) Pop(root int, isMax bool) (int, int64) {
    if root == 0 {
        return 0, 0
    }
    rv := p.nodes[root].val + p.nodes[root].tag
    t, l, r := p.nodes[root].tag, p.nodes[root].left, p.nodes[root].right
    nroot := p.Merge(l, r, t, t, isMax)
    return nroot, rv
}

func (p *Pool) Peek(root int) *int64 {
    if root == 0 {
        return nil
    }
    res := p.nodes[root].val + p.nodes[root].tag
    return &res
}

func (p *Pool) Shift(root int, d int64) {
    if root != 0 {
        p.nodes[root].tag += d
    }
}

type DP struct {
    l   int // 大根堆
    r   int // 小根堆
    min int64
}

func NewDP() *DP {
    return &DP{l: 0, r: 0, min: 0}
}

func (dp *DP) Erode(pool *Pool) {
    pool.Shift(dp.l, -1)
    pool.Shift(dp.r, 1)
}

func (dp *DP) Fix(pool *Pool) {
    for {
        xPtr, yPtr := pool.Peek(dp.l), pool.Peek(dp.r)
        if xPtr == nil || yPtr == nil {
            break
        }
        x := *xPtr
        y := *yPtr
        if x <= y {
            break
        }
        nl, _ := pool.Pop(dp.l, true)
        nr, _ := pool.Pop(dp.r, false)
        dp.l, dp.r = nl, nr
        dp.min += x - y
        dp.l = pool.Push(dp.l, y, true)
        dp.r = pool.Push(dp.r, x, false)
    }
}

func (dp *DP) AddAbs(a int64, pool *Pool) {
    dp.l = pool.Push(dp.l, a, true)
    dp.r = pool.Push(dp.r, a, false)
    dp.Fix(pool)
}

func (dp *DP) Merge(other *DP, pool *Pool) {
    dp.min += other.min
    dp.l = pool.Merge(dp.l, other.l, 0, 0, true)
    dp.r = pool.Merge(dp.r, other.r, 0, 0, false)
    dp.Fix(pool)
}

func stabilizeBinaryTree(root *TreeNode) int64 {
    if root == nil {
        return 0
    }
    var lid, rid []int
    var vals []int64
    var ord []int
    lid = append(lid, -1)
    rid = append(rid, -1)
    vals = append(vals, int64(root.Val))
    type StackItem struct {
        node *TreeNode
        id   int
    }
    stack := []StackItem{{node: root, id: 0}}
    next := 1
    for len(stack) > 0 {
        top := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        node, id := top.node, top.id
        ord = append(ord, id)
        lc, rc := node.Left, node.Right
        if lc != nil {
            cid := next
            next++
            lid = append(lid, -1)
            rid = append(rid, -1)
            vals = append(vals, int64(lc.Val))
            lid[id] = cid
            stack = append(stack, StackItem{node: lc, id: cid})
        }
        if rc != nil {
            cid := next
            next++
            lid = append(lid, -1)
            rid = append(rid, -1)
            vals = append(vals, int64(rc.Val))
            rid[id] = cid
            stack = append(stack, StackItem{node: rc, id: cid})
        }
    }
    m := next
    pool := NewPool()
    slots := make([]*DP, m)
    // 逆序后序遍历
    for i := len(ord) - 1; i >= 0; i-- {
        id := ord[i]
        base := NewDP()
        for _, cid := range []int{lid[id], rid[id]} {
            if cid == -1 {
                continue
            }
            childDP := slots[cid]
            slots[cid] = nil // take，转移所有权，和rust take()一致
            childDP.Erode(pool)
            base.Merge(childDP, pool)
        }
        base.AddAbs(vals[id], pool)
        slots[id] = base
    }
    res := slots[0].min
    return res
}

func main() {
    // Example 1:
    // Input: root = [5,8,3]
    // Output: 3
    // Explanation:
    //         5
    //      /     \
    //     8       3
    // <img src="https://assets.leetcode.com/uploads/2026/02/16/screenshot-2026-02-16-at-21240pm.png" style="width: 180px; height: 153px;">
    // - Let u be the root node with value 5. Decrease the left child value from 8 to 6 using 2 operations. Then |u.val - u.left.val| = |5 - 6| = 1.
    // - Increase the right child value from 3 to 4 using 1 operation. Then |u.val - u.right.val| = |5 - 4| = 1.
    // - Thus, the answer is 2 + 1 = 3.
    tree1 := buildTree([]*int{ptr(5), ptr(8), ptr(3)})
    printTree(tree1)
    fmt.Println(stabilizeBinaryTree(tree1)) // 3
    // Example 2:
    // Input: root = [2,-1,-4]
    // Output: 5
    // Explanation:
    //         2
    //      /     \
    //     -1       -4
    // <img src="https://assets.leetcode.com/uploads/2026/02/16/screenshot-2026-02-16-at-21340pm.png" style="width: 170px; height: 148px;">
    // - Let u be the root node with value 2. Decrease the value of u from 2 to −2 using 4 operations.
    // - Decrease the right child value from −4 to −3 using 1 operation. Then |u.val - u.right.val| = |-2 - (−3)| = 1.
    // - Thus, the answer is 4 + 1 = 5.
    tree2 := buildTree([]*int{ptr(2), ptr(-1), ptr(-4)})
    printTree(tree2)
    fmt.Println(stabilizeBinaryTree(tree2)) // 5
    // Example 3:
    // Input: root = [0,-2,2,-3,null,1]
    // Output: 3
    // Explanation:
    //         0
    //      /     \
    //     -2       2
    //   /       /
    // -3       1
    // <img src="https://assets.leetcode.com/uploads/2026/02/16/screenshot-2026-02-16-at-21826pm.png" style="width: 180px; height: 201px;">​​​​​​​​​​​​​​
    // - Let u be the root node with value 0. Increase the left child value from −2 to −1 using 1 operation, and decrease the right child value from 2 to 1 using 1 operation.
    // - Now |u.val - u.left.val| = |0 - (−1)| = 1 and |u.val - u.right.val| = |0 - 1| = 1.
    // - Let v be the left child of u, now having value −1. Increase its left child from −3 to −2 using 1 operation so that |v.val - v.left.val| = |−1 - (−2)| = 1.
    // - The right subtree already satisfies the condition since |1 - 1| = 0.
    // - Thus, the answer is 1 + 1 + 1 = 3.​​​​​​​
    tree3 := buildTree([]*int{ptr(0), ptr(-2), ptr(2), ptr(-3), nil, ptr(1)})
    printTree(tree3)
    fmt.Println(stabilizeBinaryTree(tree3)) // 3
}