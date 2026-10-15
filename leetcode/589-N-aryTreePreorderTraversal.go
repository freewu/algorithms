package main

// 589. N-ary Tree Preorder Traversal
// Given the root of an n-ary tree, return the preorder traversal of its nodes' values.
// Nary-Tree input serialization is represented in their level order traversal. 
// Each group of children is separated by the null value (See examples)

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png" />
// Input: root = [1,null,3,2,4,null,5,6]
// Output: [1,3,5,6,2,4]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]

// Constraints:
//         The number of nodes in the tree is in the range [0, 10^4].
//         0 <= Node.val <= 104
//         The height of the n-ary tree is less than or equal to 1000.

// Follow up: Recursive solution is trivial, could you do it iteratively?

import "fmt"

//  Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// == N 叉树构建 & 打印输出 start ==================================================================
func ptr(v int) *int {
    return &v
}

// buildNTree 按层序数组（LeetCode 格式）创建 N 叉树：
// nil 是分组符，它之后到下一个 nil 之间的值，都是队首节点的孩子
func buildNTree(nums []*int) *Node {
    if len(nums) == 0 || nums[0] == nil {
        return nil
    }
    root := &Node{Val: *nums[0]}
    queue := []*Node{root}
    i := 1
    for i < len(nums) && len(queue) > 0 {
        if nums[i] == nil { // 跳过分组符
            i++
        }
        if i >= len(nums) {
            break
        }
        parent := queue[0]
        queue = queue[1:]
        for i < len(nums) && nums[i] != nil {
            child := &Node{Val: *nums[i]}
            parent.Children = append(parent.Children, child)
            queue = append(queue, child)
            i++
        }
    }
    return root
}

// printNTree 自顶向下打印 N 叉树：根在上、孩子在下铺开，父子之间用 / | \ 连线
func printNTree(root *Node) {
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
    const gap = 4 // 兄弟子树之间的最小列间隔
    // 画布一格一个字符、用 rune 存，制表符才只占一列
    blank := func(n int) []rune {
        line := make([]rune, n)
        for i := range line {
            line[i] = ' '
        }
        return line
    }

    type canvas struct {
        lines   []string // 子树画布：值行 + 连线行 + 孩子画布
        width   int      // 画布宽度
        rootCol int      // 根节点的中心列
    }
    // render 自底向上组装：先量出每一棵子树的宽度，再横向铺开、父节点居中
    var render func(node *Node) canvas
    render = func(node *Node) canvas {
        val := []rune(itoa(node.Val))
        kids := make([]canvas, 0, len(node.Children))
        for _, child := range node.Children {
            kids = append(kids, render(child))
        }
        kidsWidth, kidsHeight := 0, 0
        for i, kid := range kids {
            if i > 0 {
                kidsWidth += gap
            }
            kidsWidth += kid.width
            if len(kid.lines) > kidsHeight {
                kidsHeight = len(kid.lines)
            }
        }
        width := len(val)
        if kidsWidth > width {
            width = kidsWidth
        }
        kidsStart := 0 // 第一个孩子画布的起始列
        if kidsWidth > 0 {
            kidsStart = (width - kidsWidth) / 2
        }
        // 每个孩子在画布里的中心列；父节点中心对齐到首尾孩子的中间
        centers := make([]int, 0, len(kids))
        col := kidsStart
        for _, kid := range kids {
            centers = append(centers, col+kid.rootCol)
            col += kid.width + gap
        }
        rootCol := width / 2 // 没有孩子时直接居中
        if len(centers) > 0 {
            rootCol = (centers[0] + centers[len(centers)-1]) / 2
        }
        row := blank(width)
        start := rootCol - len(val)/2 // 值按中心列居中
        if start < 0 {
            start = 0
        }
        if start > width-len(val) {
            start = width - len(val)
        }
        copy(row[start:], val)
        lines := []string{string(row)}
        if len(kids) == 0 {
            return canvas{lines: lines, width: width, rootCol: rootCol}
        }

        // 连线用制表符拼：父节点一根竖线 + 一条横向总线
        left, right := rootCol, rootCol // 总线的左右端
        for _, c := range centers {
            if c < left {
                left = c
            }
            if c > right {
                right = c
            }
        }
        stem := blank(width)
        stem[rootCol] = '│' // 父节点向下的竖线
        lines = append(lines, string(stem))
        if left != right { // 孩子正对父节点时总线只有一个点，一根竖线就够
            bus := blank(width)
            for c := left; c <= right; c++ {
                bus[c] = '─'
            }
            switch { // 父节点接口：向上接竖线，左右接总线
            case rootCol == left:
                bus[rootCol] = '└'
            case rootCol == right:
                bus[rootCol] = '┘'
            default:
                bus[rootCol] = '┴'
            }
            for _, c := range centers { // 孩子接口：向下接竖线
                switch {
                case c == left:
                    bus[c] = '┌'
                case c == right:
                    bus[c] = '┐'
                default:
                    bus[c] = '┬'
                }
                if c == rootCol { // 孩子正好在父节点正下方：上下左右都接上
                    switch {
                    case rootCol == left:
                        bus[c] = '├'
                    case rootCol == right:
                        bus[c] = '┤'
                    default:
                        bus[c] = '┼'
                    }
                }
            }
            lines = append(lines, string(bus))
        }

        for r := 0; r < kidsHeight; r++ { // 孩子画布：根节点行对齐，矮画布下方留空
            row := blank(width)
            col := kidsStart
            for _, kid := range kids {
                if r < len(kid.lines) {
                    copy(row[col:], []rune(kid.lines[r]))
                }
                col += kid.width + gap
            }
            lines = append(lines, string(row))
        }
        return canvas{lines: lines, width: width, rootCol: rootCol}
    }

    for _, s := range render(root).lines {
        line := []rune(s)
        end := len(line)
        for end > 0 && line[end-1] == ' ' {
            end--
        }
        fmt.Println(string(line[:end]))
    }
}
// == N 叉树构建 & 打印输出 end ====================================================================

// 迭代
func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, r.Val)
		tmp := []*Node{}
        // 将父节点的所有孩子节点逆序入栈，逆序的目的是为了让前序节点永远在栈顶。依次循环直到栈里所有元素都出栈
		for _, v := range r.Children {
			tmp = append([]*Node{v}, tmp...) // 逆序存点
		}
		stack = append(stack, tmp...)
	}
	return res
}

// 递归
func preorder1(root *Node) []int {
	res := []int{}
	preorderdfs(root, &res)
	return res
}

func preorderdfs(root *Node, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		for i := 0; i < len(root.Children); i++ {
			preorderdfs(root.Children[i], res)
		}
	}
}

// best solution
func preorder2(root *Node) []int {
    ret := []int{}
    if root != nil {
        stack := []*Node{root}
        for len(stack) != 0 {
            cur := stack[len(stack) - 1]
            stack = stack[:len(stack)-1]
            ret = append(ret, cur.Val)
            for i := len(cur.Children) - 1;i >= 0;i-= 1 {
                stack = append(stack, cur.Children[i])
            }
        }
    }
    return ret
}

// best solution
func preorder3(root *Node) []int {
    res := []int{}
    if root == nil {
        return res
    }
    res = append(res, root.Val)
    for _, child := range root.Children {
        res = append(res, preorder3(child)...)
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png" />
    // Input: root = [1,null,3,2,4,null,5,6]
    // Output: [1,3,5,6,2,4]
    tree1 := buildNTree([]*int{ptr(1), nil, ptr(3), ptr(2), ptr(4), nil, ptr(5), ptr(6)})
    printNTree(tree1)
    fmt.Println(preorder(tree1))  // [1,3,5,6,2,4]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />
    // Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    // Output: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]
    tree2 := buildNTree([]*int{ptr(1), nil, ptr(2), ptr(3), ptr(4), ptr(5), nil, nil, ptr(6), ptr(7), nil, ptr(8), nil, ptr(9), ptr(10), nil, nil, ptr(11), nil, ptr(12), nil, ptr(13), nil, nil, ptr(14)})
    printNTree(tree2)
    fmt.Println(preorder(tree2))  // [1,2,3,6,7,11,14,4,8,12,5,9,13,10]

    fmt.Println(preorder1(tree1)) // [1,3,5,6,2,4]
    fmt.Println(preorder1(tree2)) // [1,2,3,6,7,11,14,4,8,12,5,9,13,10]

    fmt.Println(preorder2(tree1)) // [1,3,5,6,2,4]
    fmt.Println(preorder2(tree2)) // [1,2,3,6,7,11,14,4,8,12,5,9,13,10]

    fmt.Println(preorder3(tree1)) // [1,3,5,6,2,4]
    fmt.Println(preorder3(tree2)) // [1,2,3,6,7,11,14,4,8,12,5,9,13,10]
}