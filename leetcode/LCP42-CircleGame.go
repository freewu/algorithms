package main

// LCP 42. 玩具套圈
// 「力扣挑战赛」场地外，小力组织了一个套玩具的游戏。
// 所有的玩具摆在平地上，toys[i] 以 [xi,yi,ri] 的形式记录了第 i 个玩具的坐标 (xi,yi) 和半径 ri。
// 小扣试玩了一下，他扔了若干个半径均为 r 的圈，circles[j] 记录了第 j 个圈的坐标 (xj,yj)。套圈的规则如下：
//     1. 若一个玩具被某个圈完整覆盖了（即玩具的任意部分均在圈内或者圈上），则该玩具被套中。
//     2. 若一个玩具被多个圈同时套中，最终仅计算为套中一个玩具

// 请帮助小扣计算，他成功套中了多少玩具。

// 注意：
//     输入数据保证任意两个玩具的圆心不会重合，但玩具之间可能存在重叠。

// 示例 1：
// 输入：toys = [[3,3,1],[3,2,1]], circles = [[4,3]], r = 2
// 输出：1
// 解释： 如图所示，仅套中一个玩具
// <img src="https://pic.leetcode-cn.com/1629194140-ydKiGF-image.png" />

// 示例 2：
// 输入：toys = [[1,3,2],[4,3,1],[7,1,2]], circles = [[1,0],[3,3]], r = 4
// 输出：2
// 解释： 如图所示，套中两个玩具
// <img src="https://pic.leetcode-cn.com/1629194157-RiOAuy-image.png" />

// 提示：
//     1 <= toys.length <= 10^4
//     0 <= toys[i][0], toys[i][1] <= 10^9
//     1 <= circles.length <= 10^4
//     0 <= circles[i][0], circles[i][1] <= 10^9
//     1 <= toys[i][2], r <= 10

import "fmt"
import "sort"

func circleGame(toys [][]int, circles [][]int, r int) int {
    res, n := 0, len(circles)
    sort.Slice(circles, func(i, j int) bool {
        return circles[i][0] < circles[j][0] || (circles[i][0] == circles[j][0] && circles[i][1] < circles[j][1])
    })
    x2arr := make([]int, n)
    for i := range x2arr {
        x2arr[i] = circles[i][0]
    }
    for _, toy := range toys {
        x1, y1, r1 := toy[0], toy[1], toy[2]
        if r1 > r { continue }
        r3 := r - r1
        x2Min, x2Max, y2Min, y2Max := x1 - r3, x1 + r3, y1 - r3, y1 + r3
        x2MinIdx := sort.SearchInts(x2arr, x2Min)
        for i := x2MinIdx; i < n && circles[i][0] <= x2Max; i++ {
            x2, y2 := circles[i][0], circles[i][1]
            if y2 >= y2Min && y2 <= y2Max && (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2) <= r3 * r3 {
                res++
                break
            }
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：toys = [[3,3,1],[3,2,1]], circles = [[4,3]], r = 2
    // 输出：1
    // 解释： 如图所示，仅套中一个玩具
    // <img src="https://pic.leetcode-cn.com/1629194140-ydKiGF-image.png" />
    fmt.Println(circleGame([][]int{{3,3,1},{3,2,1}}, [][]int{{4,3}}, 2)) // 1
    // 示例 2：
    // 输入：toys = [[1,3,2],[4,3,1],[7,1,2]], circles = [[1,0],[3,3]], r = 4
    // 输出：2
    // 解释： 如图所示，套中两个玩具
    // <img src="https://pic.leetcode-cn.com/1629194157-RiOAuy-image.png" />
    fmt.Println(circleGame([][]int{{1,3,2},{4,3,1},{7,1,2}}, [][]int{{1,0},{3,3}}, 4)) // 2
}