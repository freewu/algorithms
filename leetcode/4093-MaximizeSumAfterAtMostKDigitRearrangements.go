package main 

// 4093. Maximize Sum After At Most K Digit Rearrangements
// You are given an integer array nums of length n and an integer k.

// In one operation, you may choose any element of nums and rearrange its decimal digits in any order to form a new integer.

// Return the maximum possible sum of the array after performing at most k operations.

// Example 1:
// Input: nums = [123,45], k = 1
// Output: 366
// Explanation:
// - We choose nums[0] = 123 and rearrange its digits to form 321.
// - The array becomes [321, 45]. Thus, the total sum is 321 + 45 = 366.

// Example 2:
// Input: nums = [102,90,14], k = 2
// Output: 341
// Explanation:
// - We choose nums[0] = 102 and rearrange its digits to form 210.
// - We choose nums[2] = 14 and rearrange its digits to form 41.
// - The array becomes [210, 90, 41]. Thus, the total sum is 210 + 90 + 41 = 341.

// Constraints:
//     - 1 <= n == nums.length <= 10^{5}
//     - 1 <= nums[i] <= 10^{9}
//     - 1 <= k <= n

import "fmt"
import "strconv"
import "slices"

func maximizeSum(nums []int, k int) int64 {
    res := int64(0)
    for i, x := range nums {
        res += int64(x)
        // 把 x 从大到小排列，得到 y
        s := []byte(strconv.Itoa(x))
        slices.SortFunc(s, func(a, b byte) int { return int(b) - int(a) })
        y, _ := strconv.Atoi(string(s))
        // 把修改后的增量保存到 nums[i] 中
        nums[i] = y - x
    }
    slices.Sort(nums)
    // 选择增量最大的 k 个数
    for _, x := range nums[len(nums)-k:] {
        res += int64(x)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [123,45], k = 1
    // Output: 366
    // Explanation:
    // - We choose nums[0] = 123 and rearrange its digits to form 321.
    // - The array becomes [321, 45]. Thus, the total sum is 321 + 45 = 366.
    fmt.Println(maximizeSum([]int{123,45}, 1)) // 366
    // Example 2:
    // Input: nums = [102,90,14], k = 2
    // Output: 341
    // Explanation:
    // - We choose nums[0] = 102 and rearrange its digits to form 210.
    // - We choose nums[2] = 14 and rearrange its digits to form 41.
    // - The array becomes [210, 90, 41]. Thus, the total sum is 210 + 90 + 41 = 341.
    fmt.Println(maximizeSum([]int{102,90,14}, 2)) // 341

    fmt.Println(maximizeSum([]int{1,2,3,4,5,6,7,8,9}, 2)) // 45
    fmt.Println(maximizeSum([]int{9,8,7,6,5,4,3,2,1}, 2)) // 45
}