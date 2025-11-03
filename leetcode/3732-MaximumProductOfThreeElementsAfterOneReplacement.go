package main 

// 3732. Maximum Product of Three Elements After One Replacement
// You are given an integer array nums.

// You must replace exactly one element in the array with any integer value in the range [-10^5, 10^5] (inclusive).

// After performing this single replacement, determine the maximum possible product of any three elements at distinct indices from the modified array.

// Return an integer denoting the maximum product achievable.

// Example 1:
// Input: nums = [-5,7,0]
// Output: 3500000
// Explanation:
// Replacing 0 with -105 gives the array [-5, 7, -105], which has a product (-5) * 7 * (-10^5) = 3500000. The maximum product is 3500000.

// Example 2:
// Input: nums = [-4,-2,-1,-3]
// Output: 1200000
// Explanation:
// Two ways to achieve the maximum product include:
// [-4, -2, -3] → replace -2 with 105 → product = (-4) * 105 * (-3) = 1200000.
// [-4, -1, -3] → replace -1 with 105 → product = (-4) * 105 * (-3) = 1200000.
// The maximum product is 1200000.

// Example 3:
// Input: nums = [0,10,0]
// Output: 0
// Explanation:
// There is no way to replace an element with another integer and not have a 0 in the array. 
// Hence, the product of all three elements will always be 0, and the maximum product is 0.

// Constraints:
//     3 <= nums.length <= 10^5
//     -10^5 <= nums[i] <= 10^5

import "fmt"

// time: O(n), space: O(1)
func maxProduct(nums []int) int64 {
    mx1, mx2 := 0, 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, v := range nums {
        v = abs(v)
        if v > mx1 {
            mx2 = mx1
            mx1 = v
        } else if v > mx2 {
            mx2 = v
        }
    }
    return int64(mx1) * int64(mx2) * 1e5
}

func main() {
    // Example 1:
    // Input: nums = [-5,7,0]
    // Output: 3500000
    // Explanation:
    // Replacing 0 with -105 gives the array [-5, 7, -105], which has a product (-5) * 7 * (-10^5) = 3500000. The maximum product is 3500000.
    fmt.Println(maxProduct([]int{-5,7,0})) // 3500000
    // Example 2:
    // Input: nums = [-4,-2,-1,-3]
    // Output: 1200000
    // Explanation:
    // Two ways to achieve the maximum product include:
    // [-4, -2, -3] → replace -2 with 105 → product = (-4) * 105 * (-3) = 1200000.
    // [-4, -1, -3] → replace -1 with 105 → product = (-4) * 105 * (-3) = 1200000.
    // The maximum product is 1200000.
    fmt.Println(maxProduct([]int{-4,-2,-1,-3})) // 1200000
    // Example 3:
    // Input: nums = [0,10,0]
    // Output: 0
    // Explanation:
    // There is no way to replace an element with another integer and not have a 0 in the array. 
    // Hence, the product of all three elements will always be 0, and the maximum product is 0.
    fmt.Println(maxProduct([]int{0,10,0})) // 0 

    fmt.Println(maxProduct([]int{1,2,3,4,5,6,7,8,9})) // 7200000
    fmt.Println(maxProduct([]int{9,8,7,6,5,4,3,2,1})) // 0
}