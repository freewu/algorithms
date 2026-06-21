package main

// 3971. Maximum Total Value
// You are given two integer arrays value and decay, and an integer m.
//     1. value[i] represents the initial value at index i.
//     2. decay[i] represents how much the value decreases after each selection of index i.

// You may select any index multiple times. 
// The total number of selections across all indices must not exceed m.

// If you select index i for the tth time, where t is 1-indexed, the value gained is value[i] - decay[i] * (t - 1).

// Return the maximum total value you can obtain. 
// Since the answer may be large, return it modulo 10^9 + 7.

// Example 1:
// Input: value = [6,5,4], decay = [2,1,1], m = 4
// Output: 19
// Explanation:
// One optimal sequence of selections is as follows:
// By selecting index 0, the value gained is 6.
// By selecting index 1, the value gained is 5.
// By selecting index 2, the value gained is 4.
// By selecting index 0 again, the value gained is 6 - 2 = 4.
// The total value is 6 + 5 + 4 + 4 = 19. No other sequence of at most 4 selections gives a higher total value.

// Example 2:
// Input: value = [7,2,2], decay = [3,2,1], m = 2
// Output: 11
// Explanation:
// One optimal sequence of selections is as follows:
// By selecting index 0, the value gained is 7.
// By selecting index 0 again, the value gained is 7 - 3 = 4.
// The total value is 7 + 4 = 11.

// Example 3:
// Input: value = [4,3], decay = [5,4], m = 5
// Output: 7
// Explanation:
// One optimal sequence of selections is as follows:
// By selecting index 0, the value gained is 4.
// By selecting index 1, the value gained is 3.
// The total value is 4 + 3 = 7.

// Constraints:
//     1 <= value.length == decay.length <= 10^5
//     1 <= value[i], decay[i] <= 10^9​​​​​​​
//     1 <= m <= 10^9

import "fmt"
import "sort"
import "slices"

func maxTotalValue(value, decay []int, m int) int {
    res := 0
    low := sort.Search(slices.Max(value), func(low int) bool {
        low++
        leftM := m
        for i, v := range value {
            if v >= low {
                leftM -= (v-low)/decay[i] + 1
                if leftM < 0 { // 提前跳出循环
                    return false
                }
            }
        }
        return true
    })
    // 计算价值严格大于 low 的价值和，以及这些价值的个数
    for i, v := range value {
        if v > low {
            dec := decay[i]
            k := (v - low - 1) / dec + 1
            m -= k
            res += (v*2 - dec * (k - 1)) * k
        }
    }
    res /= 2 // 把除以 2 提到循环外面
    res += m * low // 剩余 m 次选的价值都是 low
    return res % 1_000_000_007
}

func maxTotalValue1(value []int, decay []int, m int) int {
    target, left, right, mod := 0, 0, 0, 1_000_000_007
    for _, v := range value {
        right = max(right, v)
    }
    helper := func(x int) (int, int) {
        res, count := 0, 0
        for i, v := range value {
            if v < x {
                continue
            }
            t := (v - x) / decay[i] + 1
            count += t
            res += t*v - decay[i] * (t - 1) * t / 2
            res %= mod
        }
        return count, res
    }
    for left <= right {
        mid := (left + right) / 2
        count, _ := helper(mid)
        if count >= m {
            target = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    count, res := helper(target + 1)
    res += (m - count) * target
    res %= mod  
    return res
}

func main() {
    // Example 1:
    // Input: value = [6,5,4], decay = [2,1,1], m = 4
    // Output: 19
    // Explanation:
    // One optimal sequence of selections is as follows:
    // By selecting index 0, the value gained is 6.
    // By selecting index 1, the value gained is 5.
    // By selecting index 2, the value gained is 4.
    // By selecting index 0 again, the value gained is 6 - 2 = 4.
    // The total value is 6 + 5 + 4 + 4 = 19. No other sequence of at most 4 selections gives a higher total value.
    fmt.Println(maxTotalValue([]int{6,5,4}, []int{2,1,1}, 4)) // 19
    // Example 2:
    // Input: value = [7,2,2], decay = [3,2,1], m = 2
    // Output: 11
    // Explanation:
    // One optimal sequence of selections is as follows:
    // By selecting index 0, the value gained is 7.
    // By selecting index 0 again, the value gained is 7 - 3 = 4.
    // The total value is 7 + 4 = 11.
    fmt.Println(maxTotalValue([]int{7,2,2}, []int{3,2,1}, 2)) // 11
    // Example 3:
    // Input: value = [4,3], decay = [5,4], m = 5
    // Output: 7
    // Explanation:
    // One optimal sequence of selections is as follows:
    // By selecting index 0, the value gained is 4.
    // By selecting index 1, the value gained is 3.
    // The total value is 4 + 3 = 7.  
    fmt.Println(maxTotalValue([]int{4,3}, []int{5,4}, 5)) // 7

    fmt.Println(maxTotalValue([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 17
    fmt.Println(maxTotalValue([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 17
    fmt.Println(maxTotalValue([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 17
    fmt.Println(maxTotalValue([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 17

    fmt.Println(maxTotalValue1([]int{6,5,4}, []int{2,1,1}, 4)) // 19
    fmt.Println(maxTotalValue1([]int{7,2,2}, []int{3,2,1}, 2)) // 11
    fmt.Println(maxTotalValue1([]int{4,3}, []int{5,4}, 5)) // 7
    fmt.Println(maxTotalValue1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 17
    fmt.Println(maxTotalValue1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 17
    fmt.Println(maxTotalValue1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 17
    fmt.Println(maxTotalValue1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 17
}