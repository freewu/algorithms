package main

// 3927. Minimize Array Sum Using Divisible Replacements
// You are given an integer array nums.

// You can perform the following operation any number of times:
//     1. Choose two indices a and b such that nums[a] % nums[b] == 0.
//     2. Replace nums[a] with nums[b].

// Return the minimum possible sum of the array after performing any number of operations.
 
// Example 1:
// Input: nums = [3,6,2]
// Output: 7
// Explanation:
// Choose a = 1, b = 2, where nums[a] = 6 and nums[b] = 2. Since 6 % 2 == 0, replace nums[1] with nums[2].
// The array becomes [3, 2, 2].
// No further operation reduces the sum. Thus, the final sum is 3 + 2 + 2 = 7.

// Example 2:
// Input: nums = [4,2,8,3]
// Output: 9
// Explanation:
// Choose a = 0, b = 1, where nums[a] = 4 and nums[b] = 2. Since 4 % 2 == 0, replace nums[0] with nums[1].
// Choose a = 2, b = 1, where nums[a] = 8 and nums[b] = 2. Since 8 % 2 == 0, replace nums[2] with nums[1].
// The array becomes [2, 2, 2, 3].
// No further operation reduces the sum. Thus, the final sum is 2 + 2 + 2 + 3 = 9.

// Example 3:
// Input: nums = [7,5,9]
// Output: 21
// Explanation:
// There is no pair (a, b) such that nums[a] % nums[b] == 0.
// Hence, no operation can be performed. The sum remains 7 + 5 + 9 = 21.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10​​​​^​​​5

import "fmt"
import "slices"

const MX = 100_001
var divisors [MX][]int

func init() {
    for i := 1; i < MX; i++ {
        for j := i; j < MX; j += i { // 枚举 i 的倍数 j
            divisors[j] = append(divisors[j], i) // i 是 j 的因子
        }
    }
}

func minArraySum(nums []int) int64 {
    res, mp := 0, map[int]int{}
    for _, v := range nums {
        mp[v]++
    }
    for i, c := range mp { // 遍历 mp 而不是 nums，这样重复元素只会计算一次
        for _, d := range divisors[i] { // 从小到大枚举 i 的因子 d
            if mp[d] > 0 {
                res += (d * c) // 把 i 变成 d 是最优的
                break
            }
        }
    }
    return int64(res)
}

func minArraySum1(nums []int) int64 {
    res, mx := 0, slices.Max(nums)
    best, present := make([]int, 100_001), make([]bool, mx+1)
    for _, x := range nums {
        present[x] = true
    }
    for i := range mx + 1 {
        best[i] = i
    }
    for d := 1; d <= mx; d++ {
        if !present[d] {
            continue
        }
        for m := d; m <= mx; m += d {
            if present[m] && d < best[m] {
                best[m] = d
            }
        }
    }
    for _, v := range nums {
        res += best[v]
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [3,6,2]
    // Output: 7
    // Explanation:
    // Choose a = 1, b = 2, where nums[a] = 6 and nums[b] = 2. Since 6 % 2 == 0, replace nums[1] with nums[2].
    // The array becomes [3, 2, 2].
    // No further operation reduces the sum. Thus, the final sum is 3 + 2 + 2 = 7.
    fmt.Println(minArraySum([]int{3,6,2})) // 7
    // Example 2:
    // Input: nums = [4,2,8,3]
    // Output: 9
    // Explanation:
    // Choose a = 0, b = 1, where nums[a] = 4 and nums[b] = 2. Since 4 % 2 == 0, replace nums[0] with nums[1].
    // Choose a = 2, b = 1, where nums[a] = 8 and nums[b] = 2. Since 8 % 2 == 0, replace nums[2] with nums[1].
    // The array becomes [2, 2, 2, 3].
    // No further operation reduces the sum. Thus, the final sum is 2 + 2 + 2 + 3 = 9.
    fmt.Println(minArraySum([]int{4,2,8,3})) // 9
    // Example 3:
    // Input: nums = [7,5,9]
    // Output: 21
    // Explanation:
    // There is no pair (a, b) such that nums[a] % nums[b] == 0.
    // Hence, no operation can be performed. The sum remains 7 + 5 + 9 = 21.
    fmt.Println(minArraySum([]int{7,5,9})) // 21

    fmt.Println(minArraySum([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(minArraySum([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(minArraySum1([]int{3,6,2})) // 7
    fmt.Println(minArraySum1([]int{4,2,8,3})) // 9
    fmt.Println(minArraySum1([]int{7,5,9})) // 21
    fmt.Println(minArraySum1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(minArraySum1([]int{9,8,7,6,5,4,3,2,1})) // 9
}