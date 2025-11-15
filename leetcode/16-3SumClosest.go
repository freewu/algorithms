package main

// 16. 3Sum Closest
// Given an array nums of n integers and an integer target,
// find three integers in nums such that the sum is closest to target.
// Return the sum of the three integers. You may assume that each input would have exactly one solution.

// Example 1:
// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

// Example 2:
// Input: nums = [0,0,0], target = 1
// Output: 0

// 解题思路:
// 	最优解 O(n^2)

import "fmt"
import "sort"

// 解法一 O(n^2)
func threeSumClosest(nums []int, target int) int {
    n, res, diff := len(nums), 0, 1 << 31
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    if n > 2 {
        sort.Ints(nums)
        for i := 0; i < n-2; i++ {
            if i > 0 && nums[i] == nums[i-1] {
                continue
            }
            for j, k := i+1, n-1; j < k; {
                sum := nums[i] + nums[j] + nums[k]
                if abs(sum-target) < diff {
                    res, diff = sum, abs(sum-target)
                }
                if sum == target {
                    return res
                } else if sum > target {
                    k--
                } else {
                    j++
                }
            }
        }
    }
    return res
}

// 解法二 暴力解法 O(n^3)
func threeSumClosest1(nums []int, target int) int {
    res, diff := 0, 1 << 31
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                if abs(nums[i]+nums[j]+nums[k]-target) < diff {
                    diff = abs(nums[i] + nums[j] + nums[k] - target)
                    res = nums[i] + nums[j] + nums[k]
                }
            }
        }
    }
    return res
}

func threeSumClosest2(nums []int, target int) int {
    sort.Ints(nums)
    res, diff := 0, 1 << 31
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := range nums {
        curr, j,  k := nums[i],i + 1, len(nums) - 1
        for k > j {
            sum := curr + nums[k] + nums[j]
            if abs(target - sum) < diff {
                res = sum
                diff = abs(target - sum)
            } else if sum == target {
                return sum
            } else if sum > target {
                k--
            } else if sum < target {
                j++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [-1,2,1,-4], target = 1
    // Output: 2
    // Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
    fmt.Println(threeSumClosest([]int{-1, 2, 1, -4},1)) // 2
    // Example 2:
    // Input: nums = [0,0,0], target = 1
    // Output: 0
    fmt.Println(threeSumClosest([]int{0,0,0},1)) // 0

    fmt.Println(threeSumClosest([]int{1,2,3,4,5,6,7,8,9},2)) // 6
    fmt.Println(threeSumClosest([]int{9,8,7,6,5,4,3,2,1},2)) // 6

    fmt.Println(threeSumClosest1([]int{-1, 2, 1, -4},1)) // 2
    fmt.Println(threeSumClosest1([]int{0,0,0},1)) // 0
    fmt.Println(threeSumClosest1([]int{1,2,3,4,5,6,7,8,9},2)) // 6
    fmt.Println(threeSumClosest1([]int{9,8,7,6,5,4,3,2,1},2)) // 6

    fmt.Println(threeSumClosest2([]int{-1, 2, 1, -4},1)) // 2
    fmt.Println(threeSumClosest2([]int{0,0,0},1)) // 0
    fmt.Println(threeSumClosest2([]int{1,2,3,4,5,6,7,8,9},2)) // 6
    fmt.Println(threeSumClosest2([]int{9,8,7,6,5,4,3,2,1},2)) // 6
}