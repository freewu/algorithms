package main

// LCP 15. 游乐园的迷宫
// 小王来到了游乐园，她玩的第一个项目是模拟推销员。
// 有一个二维平面地图，其中散布着 N 个推销点，编号 0 到 N-1，不存在三点共线的情况。
// 每两点之间有一条直线相连。游戏没有规定起点和终点，但限定了每次转角的方向。
// 首先，小王需要先选择两个点分别作为起点和终点，然后从起点开始访问剩余 N-2 个点恰好一次并回到终点。
// 访问的顺序需要满足一串给定的长度为 N-2 由 L 和 R 组成的字符串 direction，表示从起点出发之后在每个顶点上转角的方向。
// 根据这个提示，小王希望你能够帮她找到一个可行的遍历顺序，输出顺序下标（若有多个方案，输出任意一种）。
// 可以证明这样的遍历顺序一定是存在的。

// <img src="https://pic.leetcode-cn.com/595b60797d4a461287864a8cd05bba1d3b8760104ff83f43b902fd68477be9c3-Screenshot%202020-03-20%20at%2017.04.58.png" / >

// （上图：A->B->C 右转； 下图：D->E->F 左转）

// 示例 1：
// <img src="https://pic.leetcode-cn.com/c01c1efc423b916267c2a3a170266c925c368d62afa047c267cc1020970e55d9-%E5%9B%BE%E7%89%87.gif" />
// 输入：points = [[1,1],[1,4],[3,2],[2,1]], direction = "LL"
// 输出：[0,2,1,3]
// 解释：[0,2,1,3] 是符合"LL"的方案之一。在 [0,2,1,3] 方案中，0->2->1 是左转方向， 2->1->3 也是左转方向图片.gif

// 示例 2：
// 输入：points = [[1,3],[2,4],[3,3],[2,1]], direction = "LR"
// 输出：[0,3,1,2]
// 解释：[0,3,1,2] 是符合"LR"的方案之一。在 [0,3,1,2] 方案中，0->3->1 是左转方向， 3->1->2 是右转方向

// 限制：
//     3 <= points.length <= 1000 且 points[i].length == 2
//     1 <= points[i][0],points[i][1] <= 10000
//     direction.length == points.length - 2
//     direction 只包含 "L","R"

import "fmt"

func visitOrder(points [][]int, direction string) []int {
    res, cur := []int{}, 0 // 找最外围的一点作为起点，这里找最上边的点
    for i, v := range points {
        if v[1] > points[cur][1] {
            cur = i
        }
    }
    // 点 a 指向点 b 的向量
    sub := func (a, b []int) []int { return []int{b[0] - a[0], b[1] - a[1]} }
    // 向量 a 和向量 b 的叉乘
    cross := func (a, b []int) int { return a[0] * b[1] - a[1] * b[0] }
    used := make([]bool, len(points))
    used[cur] = true
    res = append(res, cur)
    for _, d := range direction {
        next := -1
        for i, p := range points {
            if used[i] { continue }
            if next == -1 {
                next = i
                continue
            }
            x, y := sub(points[cur], points[next]), sub(points[cur], p)
            cs := cross(x, y)
            if d == 'L' && cs < 0 || d == 'R' && cs > 0 {
                next = i
            }
        }
        cur = next
        used[cur] = true
        res = append(res, cur)
    } 
    for i := range points {
        if !used[i] {
            res = append(res, i)
            break
        }
    }
    return res
}

func main() {
    // 示例 1：
    // <img src="https://pic.leetcode-cn.com/c01c1efc423b916267c2a3a170266c925c368d62afa047c267cc1020970e55d9-%E5%9B%BE%E7%89%87.gif" />
    // 输入：points = [[1,1],[1,4],[3,2],[2,1]], direction = "LL"
    // 输出：[0,2,1,3]
    // 解释：[0,2,1,3] 是符合"LL"的方案之一。在 [0,2,1,3] 方案中，0->2->1 是左转方向， 2->1->3 也是左转方向图片.gif
    fmt.Println(visitOrder([][]int{{1,1},{1,4},{3,2},{2,1}}, "LL")) // [0,2,1,3]
    // 示例 2：
    // 输入：points = [[1,3],[2,4],[3,3],[2,1]], direction = "LR"
    // 输出：[0,3,1,2]
    // 解释：[0,3,1,2] 是符合"LR"的方案之一。在 [0,3,1,2] 方案中，0->3->1 是左转方向， 3->1->2 是右转方向
    fmt.Println(visitOrder([][]int{{1,3},{2,4},{3,3},{2,1}}, "LR")) // [0,3,1,2]
}