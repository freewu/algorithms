package main

// LCR 179. 查找总价格为目标值的两个商品
// 购物车内的商品价格按照升序记录于数组 price。请在购物车中找到两个商品的价格总和刚好是 target。若存在多种情况，返回任一结果即可。

// 示例 1：
// 输入：price = [3, 9, 12, 15], target = 18
// 输出：[3,15] 或者 [15,3]

// 示例 2：
// 输入：price = [8, 21, 27, 34, 52, 66], target = 61
// 输出：[27,34] 或者 [34,27]

// 提示：
//     1 <= price.length <= 10^5
//     1 <= price[i] <= 10^6
//     1 <= target <= 2*10^6

import "fmt"

func twoSum(price []int, target int) []int {
    r, l, sum := 0, len(price) - 1, 0
    for r < l {
        sum = price[r] + price[l]
        if sum == target {
            return []int{price[r], price[l]}
        } else if sum > target {
            l--
        } else {
            r++
        }
    }
    return []int{}
}

func main() {
    // 示例 1：
    // 输入：price = [3, 9, 12, 15], target = 18
    // 输出：[3,15] 或者 [15,3]
    fmt.Println(twoSum([]int{3, 9, 12, 15}, 18)) // [3,15]
    // 示例 2：
    // 输入：price = [8, 21, 27, 34, 52, 66], target = 61
    // 输出：[27,34] 或者 [34,27]
    fmt.Println(twoSum([]int{8, 21, 27, 34, 52, 66}, 61))  // [34,27]
}