package main

// LCP 20. 快速公交
// 小扣打算去秋日市集，由于游客较多，小扣的移动速度受到了人流影响：
//     1. 小扣从 x 号站点移动至 x + 1 号站点需要花费的时间为 inc；
//     2. 小扣从 x 号站点移动至 x - 1 号站点需要花费的时间为 dec。

// 现有 m 辆公交车，编号为 0 到 m-1。小扣也可以通过搭乘编号为 i 的公交车，从 x 号站点移动至 jump[i]*x 号站点，耗时仅为 cost[i]。
// 小扣可以搭乘任意编号的公交车且搭乘公交次数不限。

// 假定小扣起始站点记作 0，秋日市集站点记作 target，请返回小扣抵达秋日市集最少需要花费多少时间。
// 由于数字较大，最终答案需要对 1000000007 (1e9 + 7) 取模。

// 注意：小扣可在移动过程中到达编号大于 target 的站点。

// 示例 1：
// 输入：target = 31, inc = 5, dec = 3, jump = [6], cost = [10]
// 输出：33
// 解释： 
// 小扣步行至 1 号站点，花费时间为 5； 
// 小扣从 1 号站台搭乘 0 号公交至 6 * 1 = 6 站台，花费时间为 10； 
// 小扣从 6 号站台步行至 5 号站台，花费时间为 3； 
// 小扣从 5 号站台搭乘 0 号公交至 6 * 5 = 30 站台，花费时间为 10； 
// 小扣从 30 号站台步行至 31 号站台，花费时间为 5； 
// 最终小扣花费总时间为 33。

// 示例 2：
// 输入：target = 612, inc = 4, dec = 5, jump = [3,6,8,11,5,10,4], cost = [4,7,6,3,7,6,4]
// 输出：26
// 解释： 
// 小扣步行至 1 号站点，花费时间为 4； 
// 小扣从 1 号站台搭乘 0 号公交至 3 * 1 = 3 站台，花费时间为 4； 
// 小扣从 3 号站台搭乘 3 号公交至 11 * 3 = 33 站台，花费时间为 3； 
// 小扣从 33 号站台步行至 34 站台，花费时间为 4； 
// 小扣从 34 号站台搭乘 0 号公交至 3 * 34 = 102 站台，花费时间为 4； 
// 小扣从 102 号站台搭乘 1 号公交至 6 * 102 = 612 站台，花费时间为 7； 
// 最终小扣花费总时间为 26。

// 提示：
//     1 <= target <= 10^9
//     1 <= jump.length, cost.length <= 10
//     2 <= jump[i] <= 10^6
//     1 <= inc, dec, cost[i] <= 10^6

import "fmt"

// 正向 dfs 超出内存限制 2 / 80
func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
    memo := make(map[int]int)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(start int) int
    dfs = func(start int) int { // 返回从 start 站到达 target 站所需的最小时间
        if start == target { return 0 }
        if start > target  { return dec * (start - target) } // 坐公交会越走越远，只能步行返回
        if v, ok := memo[start]; ok { return v }
        memo[start] = inc + dfs(start + 1) // 第一步尝试步行到下一站
        if start > 1 { // 第一步尝试步行到上一站
            memo[start] = min(memo[start], dec + dfs(start - 1))
        }
        for i, v := range jump { // 第一步尝试坐公交
            memo[start] = min(memo[start], cost[i] + dfs(start * v))
        }
        return memo[start]
    }
    return (inc + dfs(1)) % 1_000_000_007 // 在第 0 站坐公交只能回到原地，是无用功，必须先步行到下一站
}

// 逆向 dfs
func busRapidTransit1(target int, inc int, dec int, jump []int, cost []int) int {
    memo := make(map[int]int)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(end int) int
    dfs = func(end int) int { // 返回从起点 0 到达 end 站点所需最小时间
        if end == 0 { return 0 }
        if end == 1 { return inc } // 从 0 站坐公交会回到原点，是无用功，肯定要步行
        if v, ok := memo[end]; ok { return v }
        res := end * inc // 先假设全靠步行
        for i, v := range jump { // 最后一步尝试坐每一辆公交
            x := end / v
            dest := x * v // 从 x 站点坐公交达到的站点
            if dest == end { // end 可以整除 v
                res = min(res, cost[i] + dfs(x))
            } else {
                res = min(res, cost[i] + dfs(x) + inc * (end - dest)) // 即 end 不能整除 v
                // 尝试从 x+1 坐公交之后步行返回的方案
                dest = (x + 1) * v
                res = min(res, cost[i] + dfs(x + 1) + dec * (dest - end))
            }
        }
        memo[end] = res
        return res
    }
    return dfs(target) % 1_000_000_007
}

func main() {
    // 示例 1：
    // 输入：target = 31, inc = 5, dec = 3, jump = [6], cost = [10]
    // 输出：33
    // 解释： 
    // 小扣步行至 1 号站点，花费时间为 5； 
    // 小扣从 1 号站台搭乘 0 号公交至 6 * 1 = 6 站台，花费时间为 10； 
    // 小扣从 6 号站台步行至 5 号站台，花费时间为 3； 
    // 小扣从 5 号站台搭乘 0 号公交至 6 * 5 = 30 站台，花费时间为 10； 
    // 小扣从 30 号站台步行至 31 号站台，花费时间为 5； 
    // 最终小扣花费总时间为 33。
    fmt.Println(busRapidTransit(31, 5, 3, []int{6}, []int{10})) // 33
    // 示例 2：
    // 输入：target = 612, inc = 4, dec = 5, jump = [3,6,8,11,5,10,4], cost = [4,7,6,3,7,6,4]
    // 输出：26
    // 解释： 
    // 小扣步行至 1 号站点，花费时间为 4； 
    // 小扣从 1 号站台搭乘 0 号公交至 3 * 1 = 3 站台，花费时间为 4； 
    // 小扣从 3 号站台搭乘 3 号公交至 11 * 3 = 33 站台，花费时间为 3； 
    // 小扣从 33 号站台步行至 34 站台，花费时间为 4； 
    // 小扣从 34 号站台搭乘 0 号公交至 3 * 34 = 102 站台，花费时间为 4； 
    // 小扣从 102 号站台搭乘 1 号公交至 6 * 102 = 612 站台，花费时间为 7； 
    // 最终小扣花费总时间为 26。
    fmt.Println(busRapidTransit(612, 4, 5, []int{3,6,8,11,5,10,4}, []int{4,7,6,3,7,6,4})) // 26

    fmt.Println(busRapidTransit1(31, 5, 3, []int{6}, []int{10})) // 33
    fmt.Println(busRapidTransit1(612, 4, 5, []int{3,6,8,11,5,10,4}, []int{4,7,6,3,7,6,4})) // 26
}