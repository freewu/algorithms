package main

// 3745. Maximize Expression of Three Elements
// You are given an integer array nums.

// Choose three elements a, b, and c from nums at distinct indices such that the value of the expression a + b - c is maximized.

// Return an integer denoting the maximum possible value of this expression.

// Example 1:
// Input: nums = [1,4,2,5]
// Output: 8
// Explanation:
// We can choose a = 4, b = 5, and c = 1. The expression value is 4 + 5 - 1 = 8, which is the maximum possible.

// Example 2:
// Input: nums = [-2,0,5,-2,4]
// Output: 11
// Explanation:
// We can choose a = 5, b = 4, and c = -2. The expression value is 5 + 4 - (-2) = 11, which is the maximum possible.

// Constraints:
//     3 <= nums.length <= 100
//     -100 <= nums[i] <= 100

import "fmt"

// brute force O(n^3)
func maximizeExpressionOfThree(nums []int) int {
    res, n := -101, len(nums)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            for k := 0; k < n; k++ {
                if i == j || j == k || i == k { // 下标需 互不相同 
                    continue
                }
                res = max(res, nums[i] + nums[j] - nums[k]) // 从 nums 中选择三个元素 a、b 和 c 使表达式 a + b - c 的值最大化
            }
        }
    }
    return res
}

func maximizeExpressionOfThree1(nums []int) int {
    c := -1
    for i := range nums {
        if c == -1 || nums[i] < nums[c] {
            c = i
        }
    }
    a, b := -1, -1
    for i := range nums {
        if i == c {
            continue
        }
        if a == -1 || nums[i] >= nums[a] {
            b = a
            a = i
        } else if b == -1 || nums[i] >= nums[b] {
            b = i
        }
    }
    return nums[a] + nums[b] - nums[c]
}

func main() {
    // Example 1:
    // Input: nums = [1,4,2,5]
    // Output: 8
    // Explanation:
    // We can choose a = 4, b = 5, and c = 1. The expression value is 4 + 5 - 1 = 8, which is the maximum possible.
    fmt.Println(maximizeExpressionOfThree([]int{1,4,2,5})) // 8
    // Example 2:
    // Input: nums = [-2,0,5,-2,4]
    // Output: 11
    // Explanation:
    // We can choose a = 5, b = 4, and c = -2. The expression value is 5 + 4 - (-2) = 11, which is the maximum possible.
    fmt.Println(maximizeExpressionOfThree([]int{-2,0,5,-2,4})) // 11

    fmt.Println(maximizeExpressionOfThree([]int{1,2,3,4,5,6,7,8,9})) // 16
    fmt.Println(maximizeExpressionOfThree([]int{9,8,7,6,5,4,3,2,1})) // 16
    fmt.Println(maximizeExpressionOfThree([]int{-38,-44,-37})) // -31  

    fmt.Println(maximizeExpressionOfThree1([]int{1,4,2,5})) // 8
    fmt.Println(maximizeExpressionOfThree1([]int{-2,0,5,-2,4})) // 11
    fmt.Println(maximizeExpressionOfThree1([]int{1,2,3,4,5,6,7,8,9})) // 16
    fmt.Println(maximizeExpressionOfThree1([]int{9,8,7,6,5,4,3,2,1})) // 16
    fmt.Println(maximizeExpressionOfThree1([]int{-38,-44,-37})) // -31  
}   