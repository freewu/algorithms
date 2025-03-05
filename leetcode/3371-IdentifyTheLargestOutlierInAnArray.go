package main

// 3371. Identify the Largest Outlier in an Array
// You are given an integer array nums. This array contains n elements, where exactly n - 2 elements are special numbers.
// One of the remaining two elements is the sum of these special numbers, and the other is an outlier.

// An outlier is defined as a number that is neither one of the original special numbers nor the element representing the sum of those numbers.

// Note that special numbers, the sum element, and the outlier must have distinct indices, but may share the same value.

// Return the largest potential outlier in nums.

// Example 1:
// Input: nums = [2,3,5,10]
// Output: 10
// Explanation:
// The special numbers could be 2 and 3, thus making their sum 5 and the outlier 10.

// Example 2:
// Input: nums = [-2,-1,-3,-6,4]
// Output: 4
// Explanation:
// The special numbers could be -2, -1, and -3, thus making their sum -6 and the outlier 4.

// Example 3:
// Input: nums = [1,1,1,1,1,5,5]
// Output: 5
// Explanation:
// The special numbers could be 1, 1, 1, 1, and 1, thus making their sum 5 and the other 5 as the outlier.

// Constraints:
//     3 <= nums.length <= 10^5
//     -1000 <= nums[i] <= 1000
//     The input is generated such that at least one potential outlier exists in nums.

import "fmt"

func getLargestOutlier(nums []int) int {
    res, sum := -1 << 31, 0
    mp := make(map[int]int)
    for _, v := range nums {
        sum += v
        mp[v * 2]++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        k := sum - v
        if mp[k] >= 2 || mp[k] == 1 && k != v * 2 { // 既不是原始特殊数字之一，也不是所有特殊数字的和
            res = max(res, v)
        }
    }
    if res == -1 << 31 { return -1 }
    return res
}

func getLargestOutlier1(nums []int) int {
    res, sum := -1 << 31, 0
    mp := make(map[int]int)
    for _, v := range nums {
        sum += v
        mp[v]++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for v := range mp {
        k := sum - 2 * v
        if count, ok := mp[k]; ok {
            if k != v || count > 1 { // 既不是原始特殊数字之一，也不是所有特殊数字的和
                res = max(res, k)
            }
        }
    }
    if res == -1 << 31 { return -1 }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,3,5,10]
    // Output: 10
    // Explanation:
    // The special numbers could be 2 and 3, thus making their sum 5 and the outlier 10.
    fmt.Println(getLargestOutlier([]int{2,3,5,10})) // 10
    // Example 2:
    // Input: nums = [-2,-1,-3,-6,4]
    // Output: 4
    // Explanation:
    // The special numbers could be -2, -1, and -3, thus making their sum -6 and the outlier 4.
    fmt.Println(getLargestOutlier([]int{-2,-1,-3,-6,4})) // 4
    // Example 3:
    // Input: nums = [1,1,1,1,1,5,5]
    // Output: 5
    // Explanation:
    // The special numbers could be 1, 1, 1, 1, and 1, thus making their sum 5 and the other 5 as the outlier.
    fmt.Println(getLargestOutlier([]int{1,1,1,1,1,5,5})) // 5

    fmt.Println(getLargestOutlier([]int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(getLargestOutlier([]int{9,8,7,6,5,4,3,2,1})) // -1

    fmt.Println(getLargestOutlier1([]int{2,3,5,10})) // 10
    fmt.Println(getLargestOutlier1([]int{-2,-1,-3,-6,4})) // 4
    fmt.Println(getLargestOutlier1([]int{1,1,1,1,1,5,5})) // 5
    fmt.Println(getLargestOutlier1([]int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(getLargestOutlier1([]int{9,8,7,6,5,4,3,2,1})) // -1
}