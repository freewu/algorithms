package main

// LCP 71. 集水器
// 字符串数组 shape 描述了一个二维平面中的矩阵形式的集水器，shape[i][j] 表示集水器的第 i 行 j 列为：
//     'l'表示向左倾斜的隔板（即从左上到右下）；
//     'r'表示向右倾斜的隔板（即从左下到右上）；
//     '.' 表示此位置没有隔板
// <img src="https://pic.leetcode-cn.com/1664424667-wMnPja-image.png" />

// 已知当隔板构成存储容器可以存水，每个方格代表的蓄水量为 2。
// 集水器初始浸泡在水中，除内部密闭空间外，所有位置均被水填满。 
// 现将其从水中竖直向上取出，请返回集水器最终的蓄水量。

// 注意：
//     隔板具有良好的透气性，因此空气可以穿过隔板，但水无法穿过

// 示例 1：
// 输入： shape = ["....rl","l.lr.r",".l..r.","..lr.."]
// 输出：18
// 解释：如下图所示，由于空气会穿过隔板，因此红框区域没有水
// <img src="https://pic.leetcode-cn.com/1664436239-eyYxeP-image.png" />

// 示例 2：
// 输入： shape = [".rlrlrlrl","ll..rl..r",".llrrllrr","..lr..lr."] 
// 输出：18
// 解释：如图所示。由于红框右侧未闭合，因此多余的水会从该处流走。
// <img src="https://pic.leetcode-cn.com/1664436082-SibVMv-image.png" />

// 示例 3：
// 输入： shape = ["rlrr","llrl","llr."] 输出：6
// 解释：如图所示。
// <img src="https://pic.leetcode-cn.com/1664424855-dwpUHO-image.png" />

// 示例 4：
// 输入： shape = ["...rl...","..r..l..",".r.rl.l.","r.r..l.l","l.l..rl.",".l.lr.r.","..l..r..","...lr..."]
// 输出：30
// 解释：如下图所示。由于中间为内部密闭空间，无法蓄水。
// <img src="https://pic.leetcode-cn.com/1664424894-mClEXh-image.png" />

// 提示：
//     1 <= shape.length <= 50
//     1 <= shape[i].length <= 50
//     shape[i][j] 仅为 'l'、'r' 或 '.'

import "fmt"

func reservoir(shape []string) int {
    n, m := len(shape), len(shape[0])
    // 每个格子分成四个区域（上下左右），标上序号，方便用并查集连通
    // 假设左右下还有一圈格子，直接连到超级汇点 0
    u, d, l, r := make([][]int, n + 1), make([][]int, n + 1), make([][]int, n + 1), make([][]int, n + 1)
    for i := range u {
        u[i], d[i], l[i], r[i] = make([]int, m + 2), make([]int, m + 2), make([]int, m + 2), make([]int, m + 2)
    }
    c := 1
    for i := 0; i < n; i++ {
        for j := 1; j <= m; j++ { // 假设格子的列号从 1 开始，这样方便表示左右边界
            u[i][j] = c; c++
            d[i][j] = c; c++
            l[i][j] = c; c++
            r[i][j] = c; c++
        }
    }
    // 并查集模板
    fa := make([]int, c)
    for i := range fa {
        fa[i] = i
    }
    var find func(int) int
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }
    merge := func(x, y int) { fa[find(x)] = find(y) }
    ok := make([]bool, c) // 能否容纳水
    // 倒着判断每一行，寻找可能有水的区域
    for i := n - 1; i >= 0; i-- {
        for j := 0; j <= m; j++ {
            merge(r[i][j], l[i][j+1]) // 连通左右
        }
        for j := 1; j <= m; j++ {
            merge(d[i][j], u[i+1][j]) // 连通下
            // 根据格子的类型连接格子内部四个区域
            switch shape[i][j-1] {
            case '.':
                merge(l[i][j], u[i][j])
                merge(l[i][j], d[i][j])
                merge(l[i][j], r[i][j])
            case 'l':
                merge(l[i][j], d[i][j])
                merge(r[i][j], u[i][j])
            default:
                merge(l[i][j], u[i][j])
                merge(r[i][j], d[i][j])
            }
        }
        for j := 1; j <= m; j++ {
            // 在没有连接第 i-1 行的情况下，无法到达左右下边界 => 能容纳水
            ok[l[i][j]] = find(l[i][j]) != find(0)
            ok[r[i][j]] = find(r[i][j]) != find(0)
            ok[u[i][j]] = find(u[i][j]) != find(0)
            ok[d[i][j]] = find(d[i][j]) != find(0)
        }
    }
    // 第一行连上超级汇点，方便后面统一判断是否在闭合区域里面
    for j := 1; j <= m; j++ {
        merge(u[0][j], 0)
    }
    res := 0
    for i, b := range ok {
        if b && find(i) == find(0) { // 能容纳水，且不在闭合区域里面
            res++
        }
    }
    return res / 2
}

func main() {
    // 示例 1：
    // 输入： shape = ["....rl","l.lr.r",".l..r.","..lr.."]
    // 输出：18
    // 解释：如下图所示，由于空气会穿过隔板，因此红框区域没有水
    // <img src="https://pic.leetcode-cn.com/1664436239-eyYxeP-image.png" />
    fmt.Println(reservoir([]string{"....rl","l.lr.r",".l..r.","..lr.."})) // 18
    // 示例 2：
    // 输入： shape = [".rlrlrlrl","ll..rl..r",".llrrllrr","..lr..lr."] 
    // 输出：18
    // 解释：如图所示。由于红框右侧未闭合，因此多余的水会从该处流走。
    // <img src="https://pic.leetcode-cn.com/1664436082-SibVMv-image.png" />
    fmt.Println(reservoir([]string{".rlrlrlrl","ll..rl..r",".llrrllrr","..lr..lr."})) // 18
    // 示例 3：
    // 输入： shape = ["rlrr","llrl","llr."] 输出：6
    // 解释：如图所示。
    // <img src="https://pic.leetcode-cn.com/1664424855-dwpUHO-image.png" />
    fmt.Println(reservoir([]string{"rlrr","llrl","llr."})) // 6
    // 示例 4：
    // 输入： shape = ["...rl...","..r..l..",".r.rl.l.","r.r..l.l","l.l..rl.",".l.lr.r.","..l..r..","...lr..."]
    // 输出：30
    // 解释：如下图所示。由于中间为内部密闭空间，无法蓄水。
    // <img src="https://pic.leetcode-cn.com/1664424894-mClEXh-image.png" />
    fmt.Println(reservoir([]string{"...rl...","..r..l..",".r.rl.l.","r.r..l.l","l.l..rl.",".l.lr.r.","..l..r..","...lr..."})) // 30
}