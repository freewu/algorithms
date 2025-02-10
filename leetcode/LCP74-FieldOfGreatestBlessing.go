package main

// LCP 74. 最强祝福力场
// 小扣在探索丛林的过程中，无意间发现了传说中“落寞的黄金之都”。
// 而在这片建筑废墟的地带中，小扣使用探测仪监测到了存在某种带有「祝福」效果的力场。 
// 经过不断的勘测记录，小扣将所有力场的分布都记录了下来。
// forceField[i] = [x,y,side] 表示第 i 片力场将覆盖以坐标 (x,y) 为中心，边长为 side 的正方形区域。

// 若任意一点的 力场强度 等于覆盖该点的力场数量，请求出在这片地带中 力场强度 最强处的 力场强度。

// 注意：
//     力场范围的边缘同样被力场覆盖。

// 示例 1：
// 输入： forceField = [[0,0,1],[1,0,1]]
// 输出：2
// 解释：如图所示，（0.5, 0) 处力场强度最强为 2， （0.5，-0.5）处力场强度同样是 2。
// <img src="https://pic.leetcode.cn/1681805536-zGfghe-image.png" />

// 示例 2：
// 输入： forceField = [[4,4,6],[7,5,3],[1,6,2],[5,6,3]]
// 输出：3
// 解释：如下图所示， forceField[0]、forceField[1]、forceField[3] 重叠的区域力场强度最大，返回 3
// <img src="https://pic.leetcode.cn/1681805437-HQkyZS-image.png" />

// 提示：
//     1 <= forceField.length <= 100
//     forceField[i].length == 3
//     0 <= forceField[i][0], forceField[i][1] <= 10^9
//     1 <= forceField[i][2] <= 10^9

import "fmt"
import "sort"

func fieldOfGreatestBlessing(forceField [][]int) int {
    // 1. 统计所有左下和右上坐标
    xs, ys := []int{}, []int{}
    for _, f := range forceField {
        i, j, side := f[0], f[1], f[2]
        xs = append(xs, 2 * i - side, 2 * i + side)
        ys = append(ys, 2 * j - side, 2 * j + side)
    }
    // 2. 排序去重
    unique := func(arr []int) []int {
        sort.Ints(arr)
        k := 0
        for _, v := range arr[1:] {
            if arr[k] != v {
                k++
                arr[k] = v
            }
        }
        return arr[:k + 1]
    }
    xs, ys = unique(xs), unique(ys)
    // 3. 二维差分
    res, n, m := 0, len(xs), len(ys)
    diff := make([][]int, n + 2)
    for i := range diff {
        diff[i] = make([]int, m + 2)
    }
    for _, f := range forceField {
        i, j, side := f[0], f[1], f[2]
        r1 := sort.SearchInts(xs, 2 * i - side)
        r2 := sort.SearchInts(xs, 2 * i + side)
        c1 := sort.SearchInts(ys, 2 * j - side)
        c2 := sort.SearchInts(ys, 2 * j + side)
        // 将区域 r1<=r<=r2 && c1<=c<=c2 上的数都加上 x
        // 多 +1 是为了方便求后面复原
        diff[r1 + 1][c1 + 1]++
        diff[r1 + 1][c2 + 2]--
        diff[r2 + 2][c1 + 1]--
        diff[r2 + 2][c2 + 2]++
    }
    // 4. 直接在 diff 上复原，计算最大值
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            diff[i][j] += diff[i][j-1] + diff[i-1][j] - diff[i-1][j-1]
            res = max(res, diff[i][j])
        }
    }
    return res
}

func fieldOfGreatestBlessing1(forceField [][]int) int {
    xs, ys, mpx, mpy := []float64{}, []float64{}, make(map[float64]int), make(map[float64]int)
    for _, v := range forceField {
        xs = append(xs, float64(v[0]) - float64(v[2]) / 2, float64(v[0]) + float64(v[2]) / 2)
        ys = append(ys, float64(v[1]) - float64(v[2]) / 2, float64(v[1]) + float64(v[2]) / 2)
    }
    sort.Float64s(xs)
    sort.Float64s(ys)
    for i, v := range xs {
        mpx[v] = i + 1
    }
    for i, v := range ys {
        mpy[v] = i + 1
    }
    grid := make([][]int, len(xs) + 2)
    for i := 0; i <= len(xs) + 1; i++ {
        grid[i] = make([]int, len(ys) + 2)
    }
    res := 0
    for _, v := range forceField {
        x1, y1 := mpx[float64(v[0]) - float64(v[2]) / 2], mpy[float64(v[1]) - float64(v[2]) / 2]
        x2, y2 := mpx[float64(v[0]) + float64(v[2]) / 2], mpy[float64(v[1]) + float64(v[2]) / 2]
        grid[x1][y1]++
        grid[x2 + 1][y1]--
        grid[x1][y2 + 1]--
        grid[x2 + 1][y2 + 1]++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= len(xs); i++ {
        for j := 1; j <= len(ys); j++ {
            grid[i][j] += grid[i-1][j] + grid[i][j-1] - grid[i-1][j-1]
            res = max(res, grid[i][j])
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入： forceField = [[0,0,1],[1,0,1]]
    // 输出：2
    // 解释：如图所示，（0.5, 0) 处力场强度最强为 2， （0.5，-0.5）处力场强度同样是 2。
    // <img src="https://pic.leetcode.cn/1681805536-zGfghe-image.png" />
    fmt.Println(fieldOfGreatestBlessing([][]int{{0,0,1},{1,0,1}})) // 2
    // 示例 2：
    // 输入： forceField = [[4,4,6],[7,5,3],[1,6,2],[5,6,3]]
    // 输出：3
    // 解释：如下图所示， forceField[0]、forceField[1]、forceField[3] 重叠的区域力场强度最大，返回 3
    // <img src="https://pic.leetcode.cn/1681805437-HQkyZS-image.png" />
    fmt.Println(fieldOfGreatestBlessing([][]int{{4,4,6},{7,5,3},{1,6,2},{5,6,3}})) // 3

    fmt.Println(fieldOfGreatestBlessing1([][]int{{0,0,1},{1,0,1}})) // 2
    fmt.Println(fieldOfGreatestBlessing1([][]int{{4,4,6},{7,5,3},{1,6,2},{5,6,3}})) // 3
}