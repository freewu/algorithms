package main

// LCP 46. 志愿者调配
// 「力扣挑战赛」有 n 个比赛场馆（场馆编号从 0 开始），场馆之间的通道分布情况记录于二维数组 edges 中，edges[i]= [x, y] 表示第 i 条通道连接场馆 x 和场馆 y(即两个场馆相邻)。
// 初始每个场馆中都有一定人数的志愿者（不同场馆人数可能不同），后续 m 天每天均会根据赛事热度进行志愿者人数调配。
// 调配方案分为如下三种：
//     1. 将编号为 idx 的场馆内的志愿者人数减半；
//     2. 将编号为 idx 的场馆相邻的场馆的志愿者人数都加上编号为 idx 的场馆的志愿者人数；
//     3. 将编号为 idx 的场馆相邻的场馆的志愿者人数都减去编号为 idx 的场馆的志愿者人数。

// 所有的调配信息记录于数组 plans 中，plans[i] = [num,idx] 表示第 i 天对编号 idx 的场馆执行了第 num 种调配方案。 
// 在比赛结束后对调配方案进行复盘时，不慎将第 0 个场馆的最终志愿者人数丢失，只保留了初始所有场馆的志愿者总人数 totalNum ，以及记录了第 1 ~ n-1 个场馆的最终志愿者人数的一维数组 finalCnt。
// 请你根据现有的信息求出初始每个场馆的志愿者人数，并按场馆编号顺序返回志愿者人数列表。

// 注意：
//     1. 测试数据保证当某场馆进行第一种调配时，该场馆的志愿者人数一定为偶数；
//     2. 测试数据保证当某场馆进行第三种调配时，该场馆的相邻场馆志愿者人数不为负数；
//     3. 测试数据保证比赛开始时每个场馆的志愿者人数都不超过 10^9；
//     4. 测试数据保证给定的场馆间的道路分布情况中不会出现自环、重边的情况。

// 示例 1：
// <img src="https://pic.leetcode-cn.com/1630061228-gnZsOz-image.png" />
// 输入：finalCnt = [1,16], totalNum = 21, edges = [[0,1],[1,2]], plans = [[2,1],[1,0],[3,0]]
// 输出：[5,7,9]
// 解释：
// <img src="https://pic.leetcode-cn.com/1630061300-WuVkeF-image.png" />

// 示例 2 ：
// 输入： finalCnt = [4,13,4,3,8], totalNum = 54, edges = [[0,3],[1,3],[4,3],[2,3],[2,5]], plans = [[1,1],[3,3],[2,5],[1,0]]
// 输出：[10,16,9,4,7,8]

// 提示：
//     2 <= n <= 5*10^4
//     1 <= edges.length <= min((n * (n - 1)) / 2, 5*10^4)
//     0 <= edges[i][0], edges[i][1] < n
//     1 <= plans.length <= 10
//     1 <= plans[i][0] <=3
//     0 <= plans[i][1] < n
//     finalCnt.length = n-1
//     0 <= finalCnt[i] < 10^9
//     0 <= totalNum < 5*10^13

import "fmt"

func volunteerDeployment(finalCnt []int, totalNum int64, edges [][]int, plans [][]int) []int {
    n, m := len(finalCnt) + 1, len(plans)
    graph := make([][]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])

    }
    tmp := make([][2]int, n)
    tmp[0] = [2]int{1, 0}
    for i, v := range finalCnt {
        tmp[i + 1][1] = v
    }
    for i := m - 1; i >= 0; i-- {
        num, index := plans[i][0], plans[i][1]
        if num == 1 {
            tmp[index][0] *= 2
            tmp[index][1] *= 2
        } else if num == 2 {
            for _, v := range graph[index] {
                tmp[v][0] -= tmp[index][0]
                tmp[v][1] -= tmp[index][1]
            }
        } else {
            for _, v := range graph[index] {
                tmp[v][0] += tmp[index][0]
                tmp[v][1] += tmp[index][1]
            }
        }
    }
    sum, sum0, sum1 := 0, 0, 0
    for _, t := range tmp {
        sum0 += t[0]
        sum1 += t[1]
    }
    res := make([]int, n)
    x := int((totalNum - int64(sum1)) / int64(sum0))
    for i := 1; i < n; i++ {
        res[i] = tmp[i][0] * x + tmp[i][1]
        sum += res[i]
    }
    res[0] = int(totalNum - int64(sum))
    return res
}

func volunteerDeployment1(finalCnt []int, totalNum int64, edges [][]int, plans [][]int) []int {
    n := len(finalCnt)
    count := make([]int64, n)
    for i, v := range finalCnt {
        count[i] = int64(v)
    }
    dp := make([]int, n)
    n++
    x, y := 1, int64(0)
    res, graph := make([]int, n), make([][]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    for i := len(plans) - 1; i >= 0; i-- {
        num, index := plans[i][0], plans[i][1]
        if num == 1 {
            if index == 0 {
                x *= 2
                y *= 2
            } else {
                count[index - 1] *= 2
                dp[index - 1] *= 2
            }
        } else if num == 2 {
            for _, j := range graph[index] {
                if index == 0 {
                    count[j - 1] -= y
                    dp[j - 1] -= x
                } else {
                    if j==0 {
                        x -= dp[index - 1]
                        y -= count[index - 1]
                    } else {
                        count[j - 1] -= count[index - 1]
                        dp[j - 1] -= dp[index - 1]
                    }
                }
            }
        } else {
            for _, j := range graph[index] {
                if index == 0 {
                    count[j - 1] += y
                    dp[j - 1] += x
                } else {
                    if j==0 {
                        x += dp[index - 1]
                        y += count[index - 1]
                    } else {
                        count[j - 1] += count[index - 1]
                        dp[j - 1] += dp[index - 1]
                    }
                }
            }
        }
    }
    totalNum -= y
    z := x
    for i, v := range count {
        totalNum -= v
        z += dp[i]
    }
    k := totalNum / int64(z)
    res[0] = x * int(k) + int(y)
    for i := 1; i < n; i++ {
        res[i] = dp[i - 1] * int(k) + int(count[i - 1])
    }
    return res
}

func main() {
    // 示例 1：
    // <img src="https://pic.leetcode-cn.com/1630061228-gnZsOz-image.png" />
    // 输入：finalCnt = [1,16], totalNum = 21, edges = [[0,1],[1,2]], plans = [[2,1],[1,0],[3,0]]
    // 输出：[5,7,9]
    // 解释：
    // <img src="https://pic.leetcode-cn.com/1630061300-WuVkeF-image.png" />
    fmt.Println(volunteerDeployment([]int{1,16}, 21, [][]int{{0,1},{1,2}}, [][]int{{2,1},{1,0},{3,0}})) // [5,7,9]
    // 示例 2 ：
    // 输入： finalCnt = [4,13,4,3,8], totalNum = 54, edges = [[0,3],[1,3],[4,3],[2,3],[2,5]], plans = [[1,1],[3,3],[2,5],[1,0]]
    // 输出：[10,16,9,4,7,8]
    fmt.Println(volunteerDeployment([]int{4,13,4,3,8}, 54, [][]int{{0,3},{1,3},{4,3},{2,3},{2,5}}, [][]int{{1,1},{3,3},{2,5},{1,0}})) // [10,16,9,4,7,8]

    fmt.Println(volunteerDeployment1([]int{1,16}, 21, [][]int{{0,1},{1,2}}, [][]int{{2,1},{1,0},{3,0}})) // [5,7,9]
    fmt.Println(volunteerDeployment1([]int{4,13,4,3,8}, 54, [][]int{{0,3},{1,3},{4,3},{2,3},{2,5}}, [][]int{{1,1},{3,3},{2,5},{1,0}})) // [10,16,9,4,7,8]
}