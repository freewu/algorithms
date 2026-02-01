package main

// 3824. Minimum K to Reduce Array Within Limit
// You are given a positive integer array nums.

// For a positive integer k, define nonPositive(nums, k) as the minimum number of operations needed to make every element of nums non-positive. 
// In one operation, you can choose an index i and reduce nums[i] by k.

// Return an integer denoting the minimum value of k such that nonPositive(nums, k) <= k2.

// Example 1:
// Input: nums = [3,7,5]
// Output: 3
// Explanation:
// When k = 3, nonPositive(nums, k) = 6 <= k2.
// Reduce nums[0] = 3 one time. nums[0] becomes 3 - 3 = 0.
// Reduce nums[1] = 7 three times. nums[1] becomes 7 - 3 - 3 - 3 = -2.
// Reduce nums[2] = 5 two times. nums[2] becomes 5 - 3 - 3 = -1.

// Example 2:
// Input: nums = [1]
// Output: 1
// Explanation:
// When k = 1, nonPositive(nums, k) = 1 <= k2.
// Reduce nums[0] = 1 one time. nums[0] becomes 1 - 1 = 0.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"
import "math"

func minimumK(nums []int) int {
    l, r := 1, 1 << 31
    for l < r {
        ops,m := 0, (l + r) / 2
        for i := 0; i < len(nums); i++ {
            if nums[i] % m == 0{
                ops += nums[i] / m
            } else{
                ops += nums[i] / m + 1
            }
        }
        if ops <= m * m {
            r = m
        } else{
            l = m + 1
        }
    }
    return l
}

func minimumK1(nums []int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    cubeRoot := math.Cbrt(float64(sum))
    res := int(math.Ceil(cubeRoot))
    for { // 使用 for 循环直到条件不满足
        val := 0
        for _, v := range nums {
            val += int(math.Ceil(float64(v) / float64(res)))
        }
        if val <= res * res { break } // 如果条件不满足，则退出循环
        res++ // 条件满足时，res自增
    }
    return res  
}

func main() {
    // Example 1:
    // Input: nums = [3,7,5]
    // Output: 3
    // Explanation:
    // When k = 3, nonPositive(nums, k) = 6 <= k2.
    // Reduce nums[0] = 3 one time. nums[0] becomes 3 - 3 = 0.
    // Reduce nums[1] = 7 three times. nums[1] becomes 7 - 3 - 3 - 3 = -2.
    // Reduce nums[2] = 5 two times. nums[2] becomes 5 - 3 - 3 = -1.
    fmt.Println(minimumK([]int{3,7,5})) // 3
    // Example 2:
    // Input: nums = [1]
    // Output: 1
    // Explanation:
    // When k = 1, nonPositive(nums, k) = 1 <= k2.
    // Reduce nums[0] = 1 one time. nums[0] becomes 1 - 1 = 0.   
    fmt.Println(minimumK([]int{1})) // 1

    fmt.Println(minimumK([]int{1,2,3,4,5,6,7,8,9})) // 4
    fmt.Println(minimumK([]int{9,8,7,6,5,4,3,2,1})) // 4

    fmt.Println(minimumK1([]int{3,7,5})) // 3
    fmt.Println(minimumK1([]int{1})) // 1
    fmt.Println(minimumK1([]int{1,2,3,4,5,6,7,8,9})) // 4
    fmt.Println(minimumK1([]int{9,8,7,6,5,4,3,2,1})) // 4
}