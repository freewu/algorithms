package main

// LCP 37. 最小矩形面积
// 二维平面上有 N 条直线，形式为 y = kx + b，其中 k、b为整数 且 k > 0。
// 所有直线以 [k,b] 的形式存于二维数组 lines 中，不存在重合的两条直线。两两直线之间可能存在一个交点，最多会有 C 个交点。
// 我们用一个平行于坐标轴的矩形覆盖所有的交点，请问这个矩形最小面积是多少。
// 若直线之间无交点、仅有一个交点或所有交点均在同一条平行坐标轴的直线上，则返回0。

// 注意：返回结果是浮点数，与标准答案 绝对误差或相对误差 在 10^-4 以内的结果都被视为正确结果

// 示例 1：
// 输入：lines = [[2,3],[3,0],[4,1]]
// 输出：48.00000
// 解释：三条直线的三个交点为 (3, 9) (1, 5) 和 (-1, -3)。最小覆盖矩形左下角为 (-1, -3) 右上角为 (3,9)，面积为 48

// 示例 2：
// 输入：lines = [[1,1],[2,3]]
// 输出：0.00000
// 解释：仅有一个交点 (-2，-1)

// 限制：
//     1 <= lines.length <= 10^5 且 lines[i].length == 2
//     1 <= lines[0] <= 10000
//     -10000 <= lines[1] <= 10000
//     与标准答案绝对误差或相对误差在 10^-4 以内的结果都被视为正确结果

import "fmt"
import "math"
import "sort"

func minRecSize(lines [][]int) float64 {
    // 求交点的子函数
    getIntersect := func(line1, line2 []int) (float64, float64) {
        x := float64(line2[1]-line1[1]) / float64(line1[0]-line2[0])
        y := float64(line1[0])*x + float64(line1[1])
        return x, y
    }
    // 若不多于2条直线返回0
    n := len(lines)
    if n <= 2 {
        return 0.0
    }
    // 按照斜率k和b从小到大排序
    sort.Slice(lines, func(i, j int) bool {
        if lines[i][0] == lines[j][0] {
            return lines[i][1] < lines[j][1]
        }
        return lines[i][0] < lines[j][0]
    })
    // 右上方角点为 (maX, maY), 左下方角点为 (miX, miY)
    maX, miX, maY, miY := math.Inf(-1), math.Inf(1), math.Inf(-1), math.Inf(1)
    // lastLow/lastHigh分别表示斜率恰好小于当前直线的直线组中最下/上方的直线
    lastLow, lastHigh := lines[0], lines[0]
    // curLow/curHigh表示斜率等于当前直线的直线组中最下/上方的直线
    curLow, curHigh := lines[0], lines[0]
    for i := 1; i < n; i++ {
        // 若直线i平行于直线i-1, 更新curHigh
        if lines[i][0] == lines[i-1][0] {
            curHigh = lines[i]
        } else {
            // 若直线i与直线i-1不平行，即新的一组，更新lastLow, lastHigh, curLow, curHigh
            lastLow, lastHigh = curLow, curHigh
            curLow, curHigh = lines[i], lines[i]
        }
        // 计算直线i与lastLow, lastHigh的交点作更新
        if lines[i][0] != lastLow[0] {
            x, y := getIntersect(lines[i], lastLow)
            miX = math.Min(x, miX)
            miY = math.Min(y, miY)
        }
        if lines[i][0] != lastHigh[0] {
            x, y := getIntersect(lines[i], lastHigh)
            maX = math.Max(x, maX)
            maY = math.Max(y, maY)
        }
    }
    // 如果所有直线都平行，返回0
    if curLow[0] == lastLow[0] && curLow[1] == lastLow[1] {
        return 0.0
    }
    // 计算矩形面积
    return (maX - miX) * (maY - miY)
}

func main() {
    // 示例 1：
    // 输入：lines = [[2,3],[3,0],[4,1]]
    // 输出：48.00000
    // 解释：三条直线的三个交点为 (3, 9) (1, 5) 和 (-1, -3)。最小覆盖矩形左下角为 (-1, -3) 右上角为 (3,9)，面积为 48
    fmt.Println(minRecSize([][]int{{2,3},{3,0},{4,1}})) // 48.00000
    // 示例 2：
    // 输入：lines = [[1,1],[2,3]]
    // 输出：0.00000
    // 解释：仅有一个交点 (-2，-1)
    fmt.Println(minRecSize([][]int{{1,1},{2,3}})) // 0.00000
}