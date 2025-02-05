package main

// LCP 18. 早餐组合
// 小扣在秋日市集选择了一家早餐摊位，一维整型数组 staple 中记录了每种主食的价格，一维整型数组 drinks 中记录了每种饮料的价格。
// 小扣的计划选择一份主食和一款饮料，且花费不超过 x 元。请返回小扣共有多少种购买方案。

// 注意：答案需要以 1e9 + 7 (1000000007) 为底取模，如：计算初始结果为：1000000008，请返回 1

// 示例 1：
// 输入：staple = [10,20,5], drinks = [5,5,2], x = 15
// 输出：6
// 解释：小扣有 6 种购买方案，所选主食与所选饮料在数组中对应的下标分别是： 
// 第 1 种方案：staple[0] + drinks[0] = 10 + 5 = 15； 
// 第 2 种方案：staple[0] + drinks[1] = 10 + 5 = 15； 
// 第 3 种方案：staple[0] + drinks[2] = 10 + 2 = 12； 
// 第 4 种方案：staple[2] + drinks[0] = 5 + 5 = 10； 
// 第 5 种方案：staple[2] + drinks[1] = 5 + 5 = 10； 
// 第 6 种方案：staple[2] + drinks[2] = 5 + 2 = 7。

// 示例 2：
// 输入：staple = [2,1,1], drinks = [8,9,5,1], x = 9
// 输出：8
// 解释：小扣有 8 种购买方案，所选主食与所选饮料在数组中对应的下标分别是： 
// 第 1 种方案：staple[0] + drinks[2] = 2 + 5 = 7； 
// 第 2 种方案：staple[0] + drinks[3] = 2 + 1 = 3； 
// 第 3 种方案：staple[1] + drinks[0] = 1 + 8 = 9； 
// 第 4 种方案：staple[1] + drinks[2] = 1 + 5 = 6； 
// 第 5 种方案：staple[1] + drinks[3] = 1 + 1 = 2； 
// 第 6 种方案：staple[2] + drinks[0] = 1 + 8 = 9； 
// 第 7 种方案：staple[2] + drinks[2] = 1 + 5 = 6； 
// 第 8 种方案：staple[2] + drinks[3] = 1 + 1 = 2；

// 提示：
//     1 <= staple.length <= 10^5
//     1 <= drinks.length <= 10^5
//     1 <= staple[i],drinks[i] <= 10^5
//     1 <= x <= 2*10^5

import "fmt"
import "sort"

// 超出时间限制 60 / 65
func breakfastNumber(staple []int, drinks []int, x int) int {
    sort.Ints(staple)
    sort.Ints(drinks)
    res := 0
    for _, v1 := range staple {
        if v1 > x { break } // 没必要再组合了
        for _, v2 := range drinks {
            if v2 > x || v1 + v2 > x { break }
            res++
        }
    }
    return res % 1_000_000_007
}

func breakfastNumber1(staple []int, drinks []int, x int) int {
    res, mod := 0, 1_000_000_007
    sort.Ints(drinks)
    findDrink := func(drinks []int,x int) int {
        l, r := 0, len(drinks) - 1
        for l <= r {
            mid := l + (r - l) / 2
            if drinks[mid] <= x {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
        return r + 1
    } 
    for _, v := range staple {
        diff := x - v
        if diff >= drinks[0] && diff > 0 {
            res += findDrink(drinks, diff ) % mod
        }
    }
    return res % mod
}

func main() {
    // 示例 1：
    // 输入：staple = [10,20,5], drinks = [5,5,2], x = 15
    // 输出：6
    // 解释：小扣有 6 种购买方案，所选主食与所选饮料在数组中对应的下标分别是： 
    // 第 1 种方案：staple[0] + drinks[0] = 10 + 5 = 15； 
    // 第 2 种方案：staple[0] + drinks[1] = 10 + 5 = 15； 
    // 第 3 种方案：staple[0] + drinks[2] = 10 + 2 = 12； 
    // 第 4 种方案：staple[2] + drinks[0] = 5 + 5 = 10； 
    // 第 5 种方案：staple[2] + drinks[1] = 5 + 5 = 10； 
    // 第 6 种方案：staple[2] + drinks[2] = 5 + 2 = 7。
    fmt.Println(breakfastNumber([]int{10,20,5}, []int{5,5,2}, 15)) // 6
    // 示例 2：
    // 输入：staple = [2,1,1], drinks = [8,9,5,1], x = 9
    // 输出：8
    // 解释：小扣有 8 种购买方案，所选主食与所选饮料在数组中对应的下标分别是： 
    // 第 1 种方案：staple[0] + drinks[2] = 2 + 5 = 7； 
    // 第 2 种方案：staple[0] + drinks[3] = 2 + 1 = 3； 
    // 第 3 种方案：staple[1] + drinks[0] = 1 + 8 = 9； 
    // 第 4 种方案：staple[1] + drinks[2] = 1 + 5 = 6； 
    // 第 5 种方案：staple[1] + drinks[3] = 1 + 1 = 2； 
    // 第 6 种方案：staple[2] + drinks[0] = 1 + 8 = 9； 
    // 第 7 种方案：staple[2] + drinks[2] = 1 + 5 = 6； 
    // 第 8 种方案：staple[2] + drinks[3] = 1 + 1 = 2；
    fmt.Println(breakfastNumber([]int{2,1,1}, []int{8,9,5,1}, 9)) // 8

    fmt.Println(breakfastNumber1([]int{10,20,5}, []int{5,5,2}, 15)) // 6
    fmt.Println(breakfastNumber1([]int{2,1,1}, []int{8,9,5,1}, 9)) // 8
}