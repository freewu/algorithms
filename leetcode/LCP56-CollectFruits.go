package main

// LCP 55. 采集果实
// 欢迎各位勇者来到力扣新手村，本次训练内容为「采集果实」。

// 在新手村中，各位勇者需要采集一些果实来制作药剂。
// time[i] 表示勇者每次采集 1～limit 颗第 i 种类型的果实需要的时间（即每次最多可以采集 limit 颗果实）。

// 当前勇者需要完成「采集若干批果实」的任务， fruits[j] = [type, num] 表示第 j 批需要采集 num 颗 type 类型的果实。
// 采集规则如下：
//     1. 按 fruits 给定的顺序依次采集每一批次
//     2. 采集完当前批次的果实才能开始采集下一批次
//     3. 勇者完成当前批次的采集后将清空背包（即多余的果实将清空）

// 请计算并返回勇者完成采集任务最少需要的时间。

// 示例 1：
// 输入：time = [2,3,2], fruits = [[0,2],[1,4],[2,1]], limit = 3
// 输出：10
// 解释： 由于单次最多采集 3 颗 第 0 批需要采集 2 颗第 0 类型果实，需要采集 1 次，耗时为 2*1=2 第 1 批需要采集 4 颗第 1 类型果实，需要采集 2 次，耗时为 3*2=6 第 2 批需要采集 1 颗第 2 类型果实，需要采集 1 次，耗时为 2*1=2 返回总耗时 2+6+2=10

// 示例 2：
// 输入：time = [1], fruits = [[0,3],[0,5]], limit = 2
// 输出：5
// 解释： 由于单次最多采集 2 颗 第 0 批需要采集 3 颗第 0 类型果实，需要采集 2 次，耗时为 1*2=2 第 1 批需要采集 5 颗第 0 类型果实，需要采集 3 次，耗时为 1*3=3 需按照顺序依次采集，返回 2+3=5

// 提示：
//     1 <= time.length <= 100
//     1 <= time[i] <= 100
//     1 <= fruits.length <= 10^3
//     0 <= fruits[i][0] < time.length
//     1 <= fruits[i][1] < 10^3
//     1 <= limit <= 100

import "fmt"

func getMinimumTime(time []int, fruits [][]int, limit int) int {
    res := 0
    for _, v := range fruits {
        res += (v[1] + limit - 1) / limit * time[v[0]]
    }
    return res
}

func getMinimumTime1(time []int, fruits [][]int, limit int) int {
    res, n := 0, len(fruits)
    for i := 0; i < n; i++ {
        v := fruits[i][1] / limit
        if fruits[i][1] % limit != 0 {
            v++
        }
        res += (v * time[fruits[i][0]])
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：time = [2,3,2], fruits = [[0,2],[1,4],[2,1]], limit = 3
    // 输出：10
    // 解释： 由于单次最多采集 3 颗 第 0 批需要采集 2 颗第 0 类型果实，需要采集 1 次，耗时为 2*1=2 第 1 批需要采集 4 颗第 1 类型果实，需要采集 2 次，耗时为 3*2=6 第 2 批需要采集 1 颗第 2 类型果实，需要采集 1 次，耗时为 2*1=2 返回总耗时 2+6+2=10
    fmt.Println(getMinimumTime([]int{2,3,2}, [][]int{{0,2},{1,4},{2,1}}, 3)) // 10
    // 示例 2：
    // 输入：time = [1], fruits = [[0,3],[0,5]], limit = 2
    // 输出：5
    // 解释： 由于单次最多采集 2 颗 第 0 批需要采集 3 颗第 0 类型果实，需要采集 2 次，耗时为 1*2=2 第 1 批需要采集 5 颗第 0 类型果实，需要采集 3 次，耗时为 1*3=3 需按照顺序依次采集，返回 2+3=5
    fmt.Println(getMinimumTime([]int{1}, [][]int{{0,3},{0,5}}, 2)) // 5

    fmt.Println(getMinimumTime1([]int{2,3,2}, [][]int{{0,2},{1,4},{2,1}}, 3)) // 10
    fmt.Println(getMinimumTime1([]int{1}, [][]int{{0,3},{0,5}}, 2)) // 5
}