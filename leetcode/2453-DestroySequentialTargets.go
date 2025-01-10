package main

// 2453. Destroy Sequential Targets
// You are given a 0-indexed array nums consisting of positive integers, representing targets on a number line.
// You are also given an integer space.

// You have a machine which can destroy targets. 
// Seeding the machine with some nums[i] allows it to destroy all targets with values that can be represented as nums[i] + c * space, where c is any non-negative integer. 
// You want to destroy the maximum number of targets in nums.

// Return the minimum value of nums[i] you can seed the machine with to destroy the maximum number of targets.

// Example 1:
// Input: nums = [3,7,8,1,1,5], space = 2
// Output: 1
// Explanation: If we seed the machine with nums[3], then we destroy all targets equal to 1,3,5,7,9,... 
// In this case, we would destroy 5 total targets (all except for nums[2]). 
// It is impossible to destroy more than 5 targets, so we return nums[3].

// Example 2:
// Input: nums = [1,3,5,2,4,6], space = 2
// Output: 1
// Explanation: Seeding the machine with nums[0], or nums[3] destroys 3 targets. 
// It is not possible to destroy more than 3 targets.
// Since nums[0] is the minimal integer that can destroy 3 targets, we return 1.

// Example 3:
// Input: nums = [6,2,5], space = 100
// Output: 2
// Explanation: Whatever initial seed we select, we can only destroy 1 target. The minimal seed is nums[1].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= space <= 10^9

import "fmt"

func destroyTargets(nums []int, space int) int {
    mp := make(map[int][]int)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range nums {
        if arr, ok := mp[v % space]; !ok {
            mp[v % space] = []int{ v, 1 }
        } else {
            arr[0] = min(arr[0], v)
            arr[1]++
        }
    }
    targets, res := 0, 0
    for _, v := range mp {
        if v[1] > targets {
            targets, res = v[1], v[0]
        } else if v[1] == targets && res > v[0] {
            res = v[0]
        }
    }
    return res
}

func destroyTargets1(nums []int, space int) int {
    mx, index := 1, 1 << 31
    mp := make(map[int]int)
    for _, v := range nums {
        mp[v % space]++
    }
    for _, v := range nums {
        t := mp[v % space]
        if t > mx {
            mx, index = t, v
        } else if t == mx && index > v {
            index = v
        }
    }
    return index
}

func main() {
    // Example 1:
    // Input: nums = [3,7,8,1,1,5], space = 2
    // Output: 1
    // Explanation: If we seed the machine with nums[3], then we destroy all targets equal to 1,3,5,7,9,... 
    // In this case, we would destroy 5 total targets (all except for nums[2]). 
    // It is impossible to destroy more than 5 targets, so we return nums[3].
    fmt.Println(destroyTargets([]int{3,7,8,1,1,5}, 2)) // 1
    // Example 2:
    // Input: nums = [1,3,5,2,4,6], space = 2
    // Output: 1
    // Explanation: Seeding the machine with nums[0], or nums[3] destroys 3 targets. 
    // It is not possible to destroy more than 3 targets.
    // Since nums[0] is the minimal integer that can destroy 3 targets, we return 1.
    fmt.Println(destroyTargets([]int{1,3,5,2,4,6}, 2)) // 1
    // Example 3:
    // Input: nums = [6,2,5], space = 100
    // Output: 2
    // Explanation: Whatever initial seed we select, we can only destroy 1 target. The minimal seed is nums[1].
    fmt.Println(destroyTargets([]int{6,2,5}, 100)) // 2

    fmt.Println(destroyTargets([]int{1,2,3,4,5,6,7,8,9}, 2)) // 1
    fmt.Println(destroyTargets([]int{9,8,7,6,5,4,3,2,1}, 2)) // 1

    fmt.Println(destroyTargets1([]int{3,7,8,1,1,5}, 2)) // 1
    fmt.Println(destroyTargets1([]int{1,3,5,2,4,6}, 2)) // 1
    fmt.Println(destroyTargets1([]int{6,2,5}, 100)) // 2
    fmt.Println(destroyTargets1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 1
    fmt.Println(destroyTargets1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 1
}