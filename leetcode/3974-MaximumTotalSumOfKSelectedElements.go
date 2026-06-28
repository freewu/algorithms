package main

// 3974. Maximum Total Sum of K Selected Elements
// You are given an integer array nums and two integers k and mul.

// Select exactly k elements from nums. Process these elements one by one in any order you choose.

// For each selected element, independently choose one of the following:
//     1. Add the element's value to the total sum, or
//     2. Multiply the element by the current value of mul and add the result to the total sum.

// After processing each selected element, mul decreases by 1, regardless of which option was chosen. 
// The current value of mul may become 0 or negative.

// Return an integer denoting the maximum possible total sum.

// Example 1:
// Input: nums = [6,1,2,9], k = 3, mul = 2
// Output: 26
// Explanation:
// One optimal way:
// One optimal selection is nums[3] = 9, nums[0] = 6, and nums[2] = 2.
// Process nums[3] = 9 first: choose multiplication, so it contributes 9 * 2 = 18. Now, mul becomes 1.
// Process nums[0] = 6 next: choose multiplication, so it contributes 6 * 1 = 6. Now, mul becomes 0.
// Process nums[2] = 2 last: choose addition, so it contributes 2.
// The total sum is 18 + 6 + 2 = 26.

// Example 2:
// Input: nums = [3,7,5,2], k = 2, mul = 4
// Output: 43
// Explanation:
// One optimal way:
// One optimal selection is nums[1] = 7 and nums[2] = 5.
// Process nums[1] = 7 first: choose multiplication, so it contributes 7 * 4 = 28. Now, mul becomes 3.
// Process nums[2] = 5 next: choose multiplication, so it contributes 5 * 3 = 15.
// The total sum is 28 + 15 = 43.

// Example 3:
// Input: nums = [4,4], k = 1, mul = 1
// Output: 4
// Explanation:
// One optimal way:
// One optimal selection is nums[0] = 4.
// Process nums[0] = 4: choose multiplication, so it contributes 4 * 1 = 4.
// The total sum is 4.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= k <= nums.length
//     1 <= mul <= 10^5

import "fmt"
import "slices"
import "sort"

func maxSum(nums []int, k int, mul int) int64 {
    slices.SortFunc(nums, func(a, b int) int { 
        return b - a 
    })
    res := int64(0)
    for _, v := range nums[:k] {
        res += int64(v) * int64(max(mul, 1))
        mul--
    }
    return res
}

func maxSum1(nums []int, k int, mul int) int64 {
    sort.Ints(nums)
    res := 0
    for i := len(nums) - 1; i >= 0 && k > 0; i-- {
        num := nums[i]
        if num * mul > num {
            res += num * mul
            mul -= 1
        } else {
            res += num
        }
        k -= 1
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [6,1,2,9], k = 3, mul = 2
    // Output: 26
    // Explanation:
    // One optimal way:
    // One optimal selection is nums[3] = 9, nums[0] = 6, and nums[2] = 2.
    // Process nums[3] = 9 first: choose multiplication, so it contributes 9 * 2 = 18. Now, mul becomes 1.
    // Process nums[0] = 6 next: choose multiplication, so it contributes 6 * 1 = 6. Now, mul becomes 0.
    // Process nums[2] = 2 last: choose addition, so it contributes 2.
    // The total sum is 18 + 6 + 2 = 26.
    fmt.Println(maxSum([]int{6,1,2,9}, 3, 2)) // 26
    // Example 2:
    // Input: nums = [3,7,5,2], k = 2, mul = 4
    // Output: 43
    // Explanation:
    // One optimal way:
    // One optimal selection is nums[1] = 7 and nums[2] = 5.
    // Process nums[1] = 7 first: choose multiplication, so it contributes 7 * 4 = 28. Now, mul becomes 3.
    // Process nums[2] = 5 next: choose multiplication, so it contributes 5 * 3 = 15.
    // The total sum is 28 + 15 = 43.
    fmt.Println(maxSum([]int{3,7,5,2}, 2, 4)) // 43
    // Example 3:
    // Input: nums = [4,4], k = 1, mul = 1
    // Output: 4
    // Explanation:
    // One optimal way:
    // One optimal selection is nums[0] = 4.
    // Process nums[0] = 4: choose multiplication, so it contributes 4 * 1 = 4.
    // The total sum is 4.
    fmt.Println(maxSum([]int{4,4}, 1, 1)) // 4

    fmt.Println(maxSum([]int{1,2,3,4,5,6,7,8,9}, 3, 2)) // 33
    fmt.Println(maxSum([]int{9,8,7,6,5,4,3,2,1}, 3, 2)) // 33
    fmt.Println(maxSum([]int{1,2,3,4,5,6,7,8,9}, 1, 100_000)) // 900000
    fmt.Println(maxSum([]int{9,8,7,6,5,4,3,2,1}, 1, 100_000)) // 900000
    fmt.Println(maxSum([]int{1,2,3,4,5,6,7,8,9}, 1, 1)) // 9
    fmt.Println(maxSum([]int{9,8,7,6,5,4,3,2,1}, 1, 1)) // 9
    fmt.Println(maxSum([]int{1,2,3,4,5,6,7,8,9}, 9, 100_000)) // 4499880
    fmt.Println(maxSum([]int{9,8,7,6,5,4,3,2,1}, 9, 100_000)) // 4499880
    fmt.Println(maxSum([]int{1,2,3,4,5,6,7,8,9}, 9, 1)) // 45
    fmt.Println(maxSum([]int{9,8,7,6,5,4,3,2,1}, 9, 1)) // 45

    fmt.Println(maxSum1([]int{6,1,2,9}, 3, 2)) // 26
    fmt.Println(maxSum1([]int{3,7,5,2}, 2, 4)) // 43
    fmt.Println(maxSum1([]int{4,4}, 1, 1)) // 4
    fmt.Println(maxSum1([]int{1,2,3,4,5,6,7,8,9}, 3, 2)) // 33
    fmt.Println(maxSum1([]int{9,8,7,6,5,4,3,2,1}, 3, 2)) // 33
    fmt.Println(maxSum1([]int{1,2,3,4,5,6,7,8,9}, 1, 100_000)) // 900000
    fmt.Println(maxSum1([]int{9,8,7,6,5,4,3,2,1}, 1, 100_000)) // 900000
    fmt.Println(maxSum1([]int{1,2,3,4,5,6,7,8,9}, 1, 1)) // 9
    fmt.Println(maxSum1([]int{9,8,7,6,5,4,3,2,1}, 1, 1)) // 9
    fmt.Println(maxSum1([]int{1,2,3,4,5,6,7,8,9}, 9, 100_000)) // 4499880
    fmt.Println(maxSum1([]int{9,8,7,6,5,4,3,2,1}, 9, 100_000)) // 4499880
    fmt.Println(maxSum1([]int{1,2,3,4,5,6,7,8,9}, 9, 1)) // 45
    fmt.Println(maxSum1([]int{9,8,7,6,5,4,3,2,1}, 9, 1)) // 45
}
