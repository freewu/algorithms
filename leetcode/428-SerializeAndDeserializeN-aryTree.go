package main

// 428. Serialize and Deserialize N-ary Tree
// Serialization is the process of converting a data structure 
// or object into a sequence of bits so that it can be stored in a file or memory buffer, 
// or transmitted across a network connection link to be reconstructed later in the same 
// or another computer environment.

// Design an algorithm to serialize and deserialize an N-ary tree. 
// An N-ary tree is a rooted tree in which each node has no more than N children. 
// There is no restriction on how your serialization/deserialization algorithm should work. 
// You just need to ensure that an N-ary tree can be serialized to a string 
// and this string can be deserialized to the original tree structure.

// For example, you may serialize the following 3-ary tree
// <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png" />

// as [1 [3[5 6] 2 4]]. Note that this is just an example, you do not necessarily need to follow this format.
// Or you can follow LeetCode's level order traversal serialization format, where each group of children is separated by the null value.

// For example, the above tree may be serialized as [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14].
// <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />

// You do not necessarily need to follow the above-suggested formats, 
// there are many more different formats that work so please be creative and come up with different approaches yourself.

// Example 1:
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]

// Example 2:
// Input: root = [1,null,3,2,4,null,5,6]
// Output: [1,null,3,2,4,null,5,6]

// Example 3:
// Input: root = []
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     0 <= Node.val <= 10^4
//     The height of the n-ary tree is less than or equal to 1000
//     Do not use class member/global/static variables to store states. Your encode and decode algorithms should be stateless.

import "fmt"
import "strconv"
import "strings"

type Node struct {
    Val int
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

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Codec1 struct {

}

func Constructor1() *Codec1 {
    return &Codec1{}
}

func (this *Codec1) serialize(root *Node) string {
    if root == nil {
        return ""
    }
    res, q := make([]string, 0), make([]*Node, 0)
    q = append(q, root)
    for ;len(q) > 0; {
        nxtLevel := make([]*Node, 0)
        for ; len(q) > 0; {
            cn := q[0]
            q = q[1:]
            res = append(res, strconv.Itoa(cn.Val))
            if cn.Children == nil || len(cn.Children) == 0 {
                res = append(res, "0")
            } else {
                res = append(res, strconv.Itoa(len(cn.Children)))
            }

            for _, cd := range cn.Children {
                nxtLevel = append(nxtLevel, cd)
            }
        }
        q = nxtLevel
    }
    return strings.Join(res, ",")
}

func (this *Codec1) deserialize(data string) *Node {
    if len(data) == 0 {
        return nil
    }
    res := strings.Split(data, ",")
    i2n := make(map[int]*Node)
    fi, ci := 0, 2
    for ; fi < len(res); fi += 2{
        if _, e := i2n[fi]; !e {
            nv, _ := strconv.Atoi(res[fi])
            i2n[fi] = &Node{Val: nv}
            i2n[fi].Children = make([]*Node, 0)
        }
        nc, _ := strconv.Atoi(res[fi + 1])
        bdr := ci + nc * 2
        for ; ci < bdr; ci += 2 {
            cv, _ := strconv.Atoi(res[ci])
            cn := &Node{Val: cv}
            cn.Children = make([]*Node, 0)
            i2n[ci] = cn
            i2n[fi].Children = append(i2n[fi].Children, cn)
        }
    }
    return i2n[0]
}

type Codec struct {
}

func Constructor() *Codec {
    return &Codec{}
}

func (this *Codec) serialize(root *Node) string {
    if root == nil {
        return "[]"
    }
    res := strconv.Itoa(root.Val) + ","
    if root.Children == nil {
        return res
    }
    res += "["
    for _, n := range root.Children {
        res += this.serialize(n)
    }
    return res + "]"
}

func (this *Codec) deserialize(data string) *Node {
    if data == "[]" {
        return nil
    }
    parent := &Node{}
    sum, stack, res := 0, []*Node{}, parent
    var node *Node
    for _, v := range data {
        switch true {
        case '0' <= v && v <= '9':
            sum = sum * 10 + int(v - '0')
        case v == ',':
            node = &Node{Val: sum}
            parent.Children = append(parent.Children, node)
            sum = 0
        case v == '[':
            stack = append(stack, parent)
            parent = node
            sum = 0
        case v == ']':
            top := len(stack) - 1
            parent = stack[top]
            stack = stack[:top]
        }
    }
    return res.Children[0]
}

/**
  * Your Codec object will be instantiated and called as such:
  * obj := Constructor();
  * data := obj.serialize(root);
  * ans := obj.deserialize(data);
  */

func main() {
    obj := Constructor()
    // Example 1:
    // Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    // Output: [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    tree1 := buildNTree([]*int{ptr(1), nil, ptr(2), ptr(3), ptr(4), ptr(5), nil, nil, ptr(6), ptr(7), nil, ptr(8), nil, ptr(9), ptr(10), nil, nil, ptr(11), nil, ptr(12), nil, ptr(13), nil, nil, ptr(14)})
    printNTree(tree1)
    data1 := obj.serialize(tree1)
    fmt.Println(data1)
    t1 := obj.deserialize(data1)
    fmt.Println(t1.Val)
    fmt.Println(len(t1.Children))
    // Example 2:
    // Input: root = [1,null,3,2,4,null,5,6]
    // Output: [1,null,3,2,4,null,5,6]
    tree2 := buildNTree([]*int{ptr(1), nil, ptr(3), ptr(2), ptr(4), nil, ptr(5), ptr(6)})
    printNTree(tree2)
    data2 := obj.serialize(tree2)
    fmt.Println(data2)
    t2 := obj.deserialize(data2)
    fmt.Println(t2.Val)
    fmt.Println(len(t2.Children))
    // Example 3:
    // Input: root = []
    // Output: []
    data3 := obj.serialize(nil)
    fmt.Println(data3)
    t3 := obj.deserialize(data3)
    fmt.Println(t3)

    obj1 := Constructor1()
    data11 := obj1.serialize(tree1)
    fmt.Println(data11)
    t11 := obj1.deserialize(data11)
    fmt.Println(t11.Val)
    fmt.Println(len(t11.Children))

    data12 := obj1.serialize(tree2)
    fmt.Println(data12)
    t12 := obj1.deserialize(data12)
    fmt.Println(t12.Val)
    fmt.Println(len(t12.Children))
    // Example 3:
    // Input: root = []
    // Output: []
    data13 := obj1.serialize(nil)
    fmt.Println(data13)
    t13 := obj1.deserialize(data13)
    fmt.Println(t13)
}