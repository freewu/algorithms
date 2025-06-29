package main

// 594. Longest Harmonious Subsequence
// We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.
// Given an integer array nums, return the length of its longest harmonious subsequence among all its possible subsequences.
// A subsequence of array is a sequence that can be derived from the array by deleting some or no elements without changing the order of the remaining elements.

// Example 1:
// Input: nums = [1,3,2,2,5,2,3,7]
// Output: 5
// Explanation: The longest harmonious subsequence is [3,2,2,2,3].

// Example 2:
// Input: nums = [1,2,3,4]
// Output: 2

// Example 3:
// Input: nums = [1,1,1,1]
// Output: 0

// Constraints:
//     1 <= nums.length <= 2 * 10^4
//     -10^9 <= nums[i] <= 10^9

import "fmt"

func findLHS(nums []int) int {
    if len(nums) < 2 {
        return 0
    }
    res, mp := 0, make(map[int]int)
    for _, v := range nums { // 统计出每个值出现次数
        mp[v]++
    }
    for k, c := range mp {
        if n, ok := mp[k + 1]; ok { // 存在下个数据
            if c + n > res { // 两个加起的数量大时
                res = c + n 
            }
        }
    }
    return res
}

func findLHS1(nums []int) int {
    res, mp := 0, make(map[int]int)
    for _, v := range nums {
        mp[v]++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for k, v := range mp {
        if mp[k + 1] > 0  {
            res = max(res,mp[k + 1] + v)
        }
    }
    return res
}

func findLHS2(nums []int) int {
    mp := make(map[int]int, len(nums))
    for _, v := range nums {
        mp[v]++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, c1, c2 := 0, 0, 0
    for k, v := range mp {
        c1, c2 = mp[k+1], mp[k-1]
        if c1 > 0 {
            res = max(res, v + c1)
        }
        if c2 > 0 {
            res = max(res, v + c2)
        }
    }
    return res
}


func main() {
    // Example 1:
    // Input: nums = [1,3,2,2,5,2,3,7]
    // Output: 5
    // Explanation: The longest harmonious subsequence is [3,2,2,2,3].
    fmt.Println(findLHS([]int{1,3,2,2,5,2,3,7})) // 5 [3,2,2,2,3]
    // Example 2:
    // Input: nums = [1,2,3,4]
    // Output: 2
    fmt.Println(findLHS([]int{1,2,3,4})) // 2 [1,2] | [2,3] | [3,4]
    // Example 3:
    // Input: nums = [1,1,1,1]
    // Output: 0
    fmt.Println(findLHS([]int{1,1,1,1})) // 0

    fmt.Println(findLHS([]int{1,2,3,4,5,6,7,8,9})) // 2
    fmt.Println(findLHS([]int{9,8,7,6,5,4,3,2,1})) // 2

    fmt.Println(findLHS1([]int{1,3,2,2,5,2,3,7})) // 5 [3,2,2,2,3]
    fmt.Println(findLHS1([]int{1,2,3,4})) // 2 [1,2] | [2,3] | [3,4]
    fmt.Println(findLHS1([]int{1,1,1,1})) // 0
    fmt.Println(findLHS1([]int{1,2,3,4,5,6,7,8,9})) // 2
    fmt.Println(findLHS1([]int{9,8,7,6,5,4,3,2,1})) // 2

    fmt.Println(findLHS2([]int{1,3,2,2,5,2,3,7})) // 5 [3,2,2,2,3]
    fmt.Println(findLHS2([]int{1,2,3,4})) // 2 [1,2] | [2,3] | [3,4]
    fmt.Println(findLHS2([]int{1,1,1,1})) // 0
    fmt.Println(findLHS2([]int{1,2,3,4,5,6,7,8,9})) // 2
    fmt.Println(findLHS2([]int{9,8,7,6,5,4,3,2,1})) // 2
}