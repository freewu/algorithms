package main

// 1261. Find Elements in a Contaminated Binary Tree
// Given a binary tree with the following rules:
//         root.val == 0
//         If treeNode.val == x and treeNode.left != null, then treeNode.left.val == 2 * x + 1
//         If treeNode.val == x and treeNode.right != null, then treeNode.right.val == 2 * x + 2

// Now the binary tree is contaminated, which means all treeNode.val have been changed to -1.
// Implement the FindElements class:
//         FindElements(TreeNode* root) Initializes the object with a contaminated binary tree and recovers it.
//         bool find(int target) Returns true if the target value exists in the recovered binary tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/11/06/untitled-diagram-4-1.jpg" />
// Input
// ["FindElements","find","find"]
// [[[-1,null,-1]],[1],[2]]
// Output
// [null,false,true]
// Explanation
// FindElements findElements = new FindElements([-1,null,-1]); 
// findElements.find(1); // return False 
// findElements.find(2); // return True 

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/11/06/untitled-diagram-4.jpg" />
// Input
// ["FindElements","find","find","find"]
// [[[-1,-1,-1,-1,-1]],[1],[3],[5]]
// Output
// [null,true,true,false]
// Explanation
// FindElements findElements = new FindElements([-1,-1,-1,-1,-1]);
// findElements.find(1); // return True
// findElements.find(3); // return True
// findElements.find(5); // return False

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/11/07/untitled-diagram-4-1-1.jpg" />
// Input
// ["FindElements","find","find","find","find"]
// [[[-1,null,-1,-1,null,-1]],[2],[3],[4],[5]]
// Output
// [null,true,false,false,true]
// Explanation
// FindElements findElements = new FindElements([-1,null,-1,-1,null,-1]);
// findElements.find(2); // return True
// findElements.find(3); // return False
// findElements.find(4); // return False
// findElements.find(5); // return True
 
// Constraints:
//         TreeNode.val == -1
//         The height of the binary tree is less than or equal to 20
//         The total number of nodes is between [1, 10^4]
//         Total calls of find() is between [1, 10^4]
//         0 <= target <= 10^6

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
type FindElements struct {
    m map[int]struct{}
}

func Constructor(root *TreeNode) FindElements {
    m := make(map[int]struct{})
    var recoverTree func(*TreeNode, int)
    recoverTree = func(root *TreeNode, val int)  {
        if root == nil {
            return
        }
        // 把二叉树的值存入 map 中 用于 find
        m[val] = struct{}{}
        // 如果 treeNode.val == x 且 treeNode.left != null，那么 treeNode.left.val == 2 * x + 1
        recoverTree(root.Left, 2 * val + 1)
        // 如果 treeNode.val == x 且 treeNode.right != null，那么 treeNode.right.val == 2 * x + 2
        recoverTree(root.Right, 2 * val + 2)
    }
    // root.val == 0
    recoverTree(root, 0)
    return FindElements {
        m: m,
    }
}

func (this *FindElements) Find(target int) bool {
    if _, ok := this.m[target]; ok {
        return true
    }
    return false
}

type FindElements1 struct {
    root *TreeNode
    nodeCache map[int]struct{}
}

func Constructor1(root *TreeNode) FindElements1 {
    var recover func(*TreeNode, int)
    cache := make(map[int]struct{})
    recover = func(root *TreeNode, v int) {
        root.Val = v
        cache[v] = struct{}{}
        if root.Left != nil {
            recover(root.Left, 2*v+1)
        }
        if root.Right != nil {
            recover(root.Right, 2*v+2)
        }
    }
    recover(root, 0)
    return FindElements1{
        root: root,
        nodeCache: cache,
    }
}

func (this *FindElements1) Find(target int) bool {
    _, ok := this.nodeCache[target]
    return ok
}
 
 
/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
func main() {
    tree1 := buildTree([]*int{ptr(-1), nil, ptr(-1)})
    printTree(tree1)
    obj := Constructor(tree1);
    fmt.Println(obj.Find(1)) // false
    fmt.Println(obj.Find(2)) // true

    tree2 := buildTree([]*int{ptr(-1), ptr(-1), ptr(-1), ptr(-1), ptr(-1)})
    printTree(tree2)
    obj = Constructor(tree2);
    fmt.Println(obj.Find(1)) // true
    fmt.Println(obj.Find(3)) // true
    fmt.Println(obj.Find(5)) // false

    tree3 := buildTree([]*int{ptr(-1), nil, ptr(-1), ptr(-1), nil, ptr(-1)})
    printTree(tree3)
    obj = Constructor(tree3);
    fmt.Println(obj.Find(2)) // true
    fmt.Println(obj.Find(3)) // false
    fmt.Println(obj.Find(4)) // false
    fmt.Println(obj.Find(5)) // true


    obj1 := Constructor1(tree1);
    fmt.Println(obj1.Find(1)) // false
    fmt.Println(obj1.Find(2)) // true

    obj1 = Constructor1(tree2);
    fmt.Println(obj1.Find(1)) // true
    fmt.Println(obj1.Find(3)) // true
    fmt.Println(obj1.Find(5)) // false

    obj1 = Constructor1(tree3);
    fmt.Println(obj1.Find(2)) // true
    fmt.Println(obj1.Find(3)) // false
    fmt.Println(obj1.Find(4)) // false
    fmt.Println(obj1.Find(5)) // true
}
