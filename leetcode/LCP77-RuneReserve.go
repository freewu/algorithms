package main

// LCP 77. 符文储备
// 远征队在出发前需要携带一些「符文」，作为后续的冒险储备。runes[i] 表示第 i 枚符文的魔力值。

// 他们将从中选取若干符文进行携带，并对这些符文进行重新排列，以确保任意相邻的两块符文之间的魔力值相差不超过 1。

// 请返回他们能够携带的符文 最大数量。

// 示例 1：
// 输入：runes = [1,3,5,4,1,7]
// 输出：3
// 解释：最佳的选择方案为[3,5,4] 将其排列为 [3,4,5] 后，任意相邻的两块符文魔力值均不超过 1，携带数量为 3 其他满足条件的方案为 [1,1] 和 [7]，数量均小于 3。 因此返回可携带的最大数量 3。

// 示例 2：
// 输入：runes = [1,1,3,3,2,4]
// 输出：6
// 解释：排列为 [1,1,2,3,3,4]，可携带所有的符文

// 提示：
//     1 <= runes.length <= 10^4
//     0 <= runes[i] <= 10^4

import "fmt"
import "sort"

func runeReserve(runes []int) int {
    sort.Ints(runes)
    res, count := 1, 1
    for i, n := 1, len(runes); i < n; i++ {
        if runes[i] - runes[i-1] > 1 {
            count = 1 // 重新统计
            continue
        }
        count++
        if count > res {
            res = count
        }
    }
    return res
}

func runeReserve1(runes []int) int {
    sort.Ints(runes)
    res, count := 1, 1
    for i, n := 1, len(runes); i < n; i++ {
        if runes[i] - runes[i - 1] > 1 {
            count = 1 // 重新统计
        } else if count++; count > res {
            res = count
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：runes = [1,3,5,4,1,7]
    // 输出：3
    // 解释：最佳的选择方案为[3,5,4] 将其排列为 [3,4,5] 后，任意相邻的两块符文魔力值均不超过 1，携带数量为 3 其他满足条件的方案为 [1,1] 和 [7]，数量均小于 3。 因此返回可携带的最大数量 3。
    fmt.Println(runeReserve([]int{1,3,5,4,1,7})) // 3
    // 示例 2：
    // 输入：runes = [1,1,3,3,2,4]
    // 输出：6
    // 解释：排列为 [1,1,2,3,3,4]，可携带所有的符文
    fmt.Println(runeReserve([]int{1,1,3,3,2,4})) // 6

    fmt.Println(runeReserve([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(runeReserve([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(runeReserve1([]int{1,3,5,4,1,7})) // 3
    fmt.Println(runeReserve1([]int{1,1,3,3,2,4})) // 6
    fmt.Println(runeReserve1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(runeReserve1([]int{9,8,7,6,5,4,3,2,1})) // 9
}