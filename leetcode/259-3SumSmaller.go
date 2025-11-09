package main

// 259. 3Sum Smaller
// Given an array of n integers nums and an integer target, 
// find the number of index triplets i, j, k with 0 <= i < j < k < n that satisfy the condition nums[i] + nums[j] + nums[k] < target.

// Example 1:
// Input: nums = [-2,0,1,3], target = 2
// Output: 2
// Explanation: Because there are two triplets which sums are less than 2:
// [-2,0,1]
// [-2,0,3]

// Example 2:
// Input: nums = [], target = 0
// Output: 0

// Example 3:
// Input: nums = [0], target = 0
// Output: 0

// Constraints:
//     n == nums.length
//     0 <= n <= 3500
//     -100 <= nums[i] <= 100
//     -100 <= target <= 100

import "fmt"
import "sort"

// 双指针
func threeSumSmaller(nums []int, target int) int {
    res,n := 0, len(nums)
    if n < 3  {
        return res
    }
    sort.Ints(nums)
    for i := 0; i < n - 2; i++ {
        left, right := i + 1, n - 1
        for left < right {
            if nums[i] + nums[left] + nums[right] < target { // 三数累加小于 target
                res += right - left
                left++
            } else {
                right--
            }
        }
    }
    return res
}

func threeSumSmaller1(nums []int, target int) int {
    res, n := 0, len(nums)
    if n < 3 { return res }
    sort.Ints(nums)
    for i := 0; i < n; i++ {
        j, k := i + 1, n - 1
        for j < k {
            if nums[j] + nums[k] < target - nums[i] {
                res += k - j
                j++
            } else {
                k--
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [-2,0,1,3], target = 2
    // Output: 2
    // Explanation: Because there are two triplets which sums are less than 2:
    // [-2,0,1]
    // [-2,0,3]
    fmt.Println(threeSumSmaller([]int{-2,0,1,3}, 2)) // 2
    // Example 2:
    // Input: nums = [], target = 0
    // Output: 0
    fmt.Println(threeSumSmaller([]int{ }, 0)) // 0
    // Example 3:
    // Input: nums = [0], target = 0
    // Output: 0
    fmt.Println(threeSumSmaller([]int{0}, 0)) // 0

    fmt.Println(threeSumSmaller([]int{1,2,3,4,5,6,7,8,9}, 2)) // 0
    fmt.Println(threeSumSmaller([]int{9,8,7,6,5,4,3,2,1}, 2)) // 0

    fmt.Println(threeSumSmaller1([]int{-2,0,1,3}, 2)) // 2
    fmt.Println(threeSumSmaller1([]int{ }, 0)) // 0
    fmt.Println(threeSumSmaller1([]int{0}, 0)) // 0
    fmt.Println(threeSumSmaller1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 0
    fmt.Println(threeSumSmaller1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 0
}