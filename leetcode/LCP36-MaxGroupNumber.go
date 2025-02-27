package main

// LCP 36. 最多牌组数
// 麻将的游戏规则中，共有两种方式凑成「一组牌」：
//     顺子：三张牌面数字连续的麻将，例如 [4,5,6]
//     刻子：三张牌面数字相同的麻将，例如 [10,10,10]

// 给定若干数字作为麻将牌的数值（记作一维数组 tiles），请返回所给 tiles 最多可组成的牌组数。

// 注意：凑成牌组时，每张牌仅能使用一次。

// 示例 1：
// 输入：tiles = [2,2,2,3,4]
// 输出：1
// 解释：最多可以组合出 [2,2,2] 或者 [2,3,4] 其中一组牌。

// 示例 2：
// 输入：tiles = [2,2,2,3,4,1,3]
// 输出：2
// 解释：最多可以组合出 [1,2,3] 与 [2,3,4] 两组牌。

// 提示：
//     1 <= tiles.length <= 10^5
//     1 <= tiles[i] <= 10^9

import "fmt"
import "sort"

func maxGroupNumber(tiles []int) int {
    // 统计每个牌面的数量
    count := make(map[int]int)
    for _, tile := range tiles {
        count[tile]++
    }
    // 将牌面按点数从小到大排序
    keys := make([]int, 0, len(count))
    for k := range count {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    // dp[x][y] 表示在预留 x 张 [tile-2] 和 y 张 [tile-1] 的前提下，[tile] 之前的牌能组成的牌组数
    dp := make([][]int, 5)
    for i := range dp {
        dp[i] = make([]int, 5)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    dp[0][0] = 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    prevTile := 0 // 前一张牌的点数
    for _, tile := range keys {
        cnt := count[tile]
        // 如果上一张牌和这张牌没法连起来
        if prevTile != tile - 1 {
            dp00 := dp[0][0]
            dp = make([][]int, 5)
            for i := range dp {
                dp[i] = make([]int, 5)
                for j := range dp[i] {
                    dp[i][j] = -1
                }
            }
            dp[0][0] = dp00
        }
        // 新的 dp 数组
        newDp := make([][]int, 5)
        for i := range newDp {
            newDp[i] = make([]int, 5)
            for j := range newDp[i] {
                newDp[i][j] = -1
            }
        }
        for cnt2 := 0; cnt2 < 5; cnt2++ { // [tile-2] 的牌数
            for cnt1 := 0; cnt1 < 5; cnt1++ { // [tile-1] 的牌数
                if dp[cnt2][cnt1] < 0 { continue }
                // 顺子数量不能超过 [tile-2]、[tile-1]、[tile] 的牌数
                for shunzi := 0; shunzi <= min(cnt2, min(cnt1, cnt)); shunzi++ {
                    new2 := cnt1 - shunzi // 对于下一个点数 newTile = tile + 1 而言，[newTile-2] 就是当前的 [tile-1]
                    // 预留的数量不超过 4 张，也不超过 [tile] 的牌数减去顺子数量
                    for new1 := 0; new1 <= min(4, cnt-shunzi); new1++ {
                        // 新的牌组数等于以下三者相加：
                        // 1. dp 数组保存的，留下 cnt2 张 [tile-2] 和 cnt1 张 [tile-1] 的前提下，tile-1 之前的牌面能凑出来的牌组数
                        // 2. 顺子数量
                        // 3. [tile] 组成的刻子数量 = ([tile] - 顺子数量 - 留下备用的牌) / 3
                        newScore := dp[cnt2][cnt1] + shunzi + (cnt-shunzi-new1)/3
                        newDp[new2][new1] = max(newDp[new2][new1], newScore)
                    }
                }
            }
        }
        // 将 newDp 数组赋值给 dp 数组
        dp = newDp
        // 将当前 tile 记录到上一个 tile 中
        prevTile = tile
    }
    // 找到并返回 dp 的最大值
    res := 0
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            res = max(res, dp[i][j])
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：tiles = [2,2,2,3,4]
    // 输出：1
    // 解释：最多可以组合出 [2,2,2] 或者 [2,3,4] 其中一组牌。
    fmt.Println(maxGroupNumber([]int{2,2,2,3,4})) // 1
    // 示例 2：
    // 输入：tiles = [2,2,2,3,4,1,3]
    // 输出：2
    // 解释：最多可以组合出 [1,2,3] 与 [2,3,4] 两组牌。
    fmt.Println(maxGroupNumber([]int{2,2,2,3,4,1,3})) // 2

    fmt.Println(maxGroupNumber([]int{10,10,10})) // 1
    fmt.Println(maxGroupNumber([]int{1,2,2,2,3,4,5})) // 2
    fmt.Println(maxGroupNumber([]int{1,2,3,4,5,6,7,8,9})) // 3
    fmt.Println(maxGroupNumber([]int{9,8,7,6,5,4,3,2,1})) // 3
}