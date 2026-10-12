package main

// 4074. Maximum Product Pair With Target Sum
// You are given an integer array nums and an integer target.

// A pair of distinct indices (i, j) is valid if:
//     1. nums[i] + nums[j] == target
//     2. nums[i] > nums[j]

// Return a valid pair [i, j] whose product nums[i] * nums[j] is maximum among all valid pairs. 
// If no valid pair exists, return [-1, -1].

// If multiple valid pairs achieve the maximum product, you may return any of them.

// Example 1:
// Input: nums = [1,2,3,4], target = 5
// Output: [2,1]
// Explanation:
// There are 2 valid pairs:
//  No. | (i, j) | nums[i] | nums[j]  | Sum  | Product
//  1   | (3, 0) | 4       | 1        | 5    | 4
//  2   | (2, 1) | 3       | 2        | 5    | 6
// Both pairs sum to 5 and satisfy nums[i] > nums[j]. The second pair has the larger product, 3 * 2 = 6, so the answer is [2, 1].

// Example 2:
// Input: nums = [-3,-1,4,2], target = 1
// Output: [3,1]
// Explanation:
// There are 2 valid pairs:
//  No. | (i, j) | nums[i] | nums[j]  | Sum  | Product
//  1   | (2, 0) | 4       | -3       | 1    | -12
//  2   | (3, 1) | 2       | -1       | 1    | -2
// Since -2 > -12, the pair at indices (3, 1) is chosen and the answer is [3, 1].

// Example 3:
// Input: nums = [3,3,5], target = 6
// Output: [-1,-1]
// Explanation:
// ​​​​​​​No valid pair exists, since nums[0] and nums[1] sum to 6 but are equal.

// Constraints:
//     - 2 <= nums.length <= 100
//     - -100 <= nums[i], target <= 100​​​​​​​

import "fmt"

func maxProductPair(nums []int, target int) []int {
    a, b, mp := -1, -1, map[int]int{}
    for i, x := range nums {
        if x * 2 == target { // x = target - x
            continue
        }
        j, ok := mp[target-x]
        if ok && (a < 0 || x*(target-x) > nums[a]*(target-nums[a])) { // 也可以比谁离 target/2 更近
            a, b = i, j
        }
        mp[x] = i
    }
    if a >= 0 && nums[a] < nums[b] {    
        a, b = b, a // 保证 nums[a] > nums[b]
    }
    return []int{a, b}
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4], target = 5
    // Output: [2,1]
    // Explanation:
    // There are 2 valid pairs:
    //  No. | (i, j) | nums[i] | nums[j]  | Sum  | Product
    //  1   | (3, 0) | 4       | 1        | 5    | 4
    //  2   | (2, 1) | 3       | 2        | 5    | 6
    // Both pairs sum to 5 and satisfy nums[i] > nums[j]. The second pair has the larger product, 3 * 2 = 6, so the answer is [2, 1].
    fmt.Println(maxProductPair([]int{1,2,3,4}, 5)) // [2,1]
    // Example 2:
    // Input: nums = [-3,-1,4,2], target = 1
    // Output: [3,1]
    // Explanation:
    // There are 2 valid pairs:
    //  No. | (i, j) | nums[i] | nums[j]  | Sum  | Product
    //  1   | (2, 0) | 4       | -3       | 1    | -12
    //  2   | (3, 1) | 2       | -1       | 1    | -2
    // Since -2 > -12, the pair at indices (3, 1) is chosen and the answer is [3, 1].
    fmt.Println(maxProductPair([]int{-3,-1,4,2}, 1)) // [3,1]
    // Example 3:
    // Input: nums = [3,3,5], target = 6
    // Output: [-1,-1]
    // Explanation:
    // ​​​​​​​No valid pair exists, since nums[0] and nums[1] sum to 6 but are equal.
    fmt.Println(maxProductPair([]int{3,3,5}, 6)) // [-1,-1]

    fmt.Println(maxProductPair([]int{1,2,3,4,5,6,7,8,9}, 1)) // [-1,-1]
    fmt.Println(maxProductPair([]int{9,8,7,6,5,4,3,2,1}, 1)) // [-1,-1]
}