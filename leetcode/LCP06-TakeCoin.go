package main

// LCP 06. 拿硬币
// 桌上有 n 堆力扣币，每堆的数量保存在数组 coins 中。
// 我们每次可以选择任意一堆，拿走其中的一枚或者两枚，求拿完所有力扣币的最少次数。

// 示例 1：
// 输入：[4,2,1]
// 输出：4
// 解释：第一堆力扣币最少需要拿 2 次，第二堆最少需要拿 1 次，第三堆最少需要拿 1 次，总共 4 次即可拿完。

// 示例 2：
// 输入：[2,3,10]
// 输出：8

// 限制：
//     1 <= n <= 4
//     1 <= coins[i] <= 10

import "fmt"

func minCount(coins []int) int {
    res := 0
    for _, v := range coins {
        if v % 2 == 1 {
            res += (v + 1) / 2
        } else {
            res += v / 2
        }
    }
    return res
}

func minCount1(coins []int) int {
    res := 0
    for _, v := range coins {
        // 对于每堆硬币，计算需要的最少次数
        // 如果硬币数是偶数，则次数为 v / 2
        // 如果硬币数是奇数，则次数为 v / 2 + 1
        res += (v + 1) / 2
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：[4,2,1]
    // 输出：4
    // 解释：第一堆力扣币最少需要拿 2 次，第二堆最少需要拿 1 次，第三堆最少需要拿 1 次，总共 4 次即可拿完。
    fmt.Println(minCount([]int{4,2,1})) // 4
    // 示例 2：
    // 输入：[2,3,10]
    // 输出：8
    fmt.Println(minCount([]int{2,3,10})) // 8

    fmt.Println(minCount([]int{1,2,3,4,5,6,7,8,9})) // 25
    fmt.Println(minCount([]int{9,8,7,6,5,4,3,2,1})) // 25

    fmt.Println(minCount1([]int{4,2,1})) // 4
    fmt.Println(minCount1([]int{2,3,10})) // 8
    fmt.Println(minCount1([]int{1,2,3,4,5,6,7,8,9})) // 25
    fmt.Println(minCount1([]int{9,8,7,6,5,4,3,2,1})) // 25
}