package main

// LCP 29. 乐团站位
// 某乐团的演出场地可视作 num * num 的二维矩阵 grid（左上角坐标为 [0,0])，每个位置站有一位成员。
// 乐团共有 9 种乐器，乐器编号为 1~9，每位成员持有 1 个乐器。

// 为保证声乐混合效果，成员站位规则为：自 grid 左上角开始顺时针螺旋形向内循环以 1，2，...，9 循环重复排列。
// 例如当 num = 5 时，站位如图所示

// <img src="https://pic.leetcode-cn.com/1616125411-WOblWH-image.png" />

// 请返回位于场地坐标 [Xpos,Ypos] 的成员所持乐器编号。

// 示例 1：
// 输入：num = 3, Xpos = 0, Ypos = 2
// 输出：3
// 解释：
// <img src="https://pic.leetcode-cn.com/1616125437-WUOwsu-image.png" />

// 示例 2：
// 输入：num = 4, Xpos = 1, Ypos = 2
// 输出：5
// 解释：
// <img src="https://pic.leetcode-cn.com/1616125453-IIDpxg-image.png" />

// 提示：
//     1 <= num <= 10^9
//     0 <= Xpos, Ypos < num

import "fmt"

func orchestraLayout(n int, x int, y int) int {
    k := 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    if  x <= y {
        k = min(x, n - y - 1)
        return  (4 * k * (n - k) + 1 + (x + y - k * 2) - 1) % 9 + 1
    }
    k = min(y, n - x - 1) + 1 
    return  (4 * k * (n - k) + 1 - (x + y - (k - 1) * 2) - 1) % 9 + 1
}

func orchestraLayout1(n int, x int, y int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    layer := min(min(n - y - 1, n - x - 1), min(x, y)) // 确定层数，层数计算错误
    a1 := (n - 2) * 4 + 4 // 对于不同大小的矩阵，首项是多少？
    // 确定是哪条边，假定坐标点所在的这一层是最外层，这样的话就好做判断了
    left, up, right, down := layer, layer, n - layer - 1, n - layer - 1 // 定义上下左右四个边界
    last := 0
    if x == up {
        last = y - left + 1
    } else if y == right {
        last = n - 2 * layer + x - up
    } else if x == down {
        last = (n - 2 * layer) * 2 - 1 + right - y
    } else if y == left {
        last = (n - 2 * layer) * 3 - 2 + down - x
    }
    sum := (a1 + a1 - (layer - 1) * 8) * (layer) / 2  + last
    res := sum % 9
    if res == 0 { return 9 }
    return res
}

func main() {
    // 示例 1：
    // 输入：num = 3, Xpos = 0, Ypos = 2
    // 输出：3
    // 解释：
    // <img src="https://pic.leetcode-cn.com/1616125437-WUOwsu-image.png" />
    fmt.Println(orchestraLayout(3, 0, 2)) // 3
    // 示例 2：
    // 输入：num = 4, Xpos = 1, Ypos = 2
    // 输出：5
    // 解释：
    // <img src="https://pic.leetcode-cn.com/1616125453-IIDpxg-image.png" />
    fmt.Println(orchestraLayout(4, 1, 2)) // 5

    fmt.Println(orchestraLayout1(3, 0, 2)) // 3
    fmt.Println(orchestraLayout1(4, 1, 2)) // 5
}