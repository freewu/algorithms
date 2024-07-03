package main

// 1509. Minimum Difference Between Largest and Smallest Value in Three Moves
// You are given an integer array nums.
// In one move, you can choose one element of nums and change it to any value.
// Return the minimum difference between the largest and smallest value of nums after performing at most three moves.

// Example 1:
// Input: nums = [5,3,2,4]
// Output: 0
// Explanation: We can make at most 3 moves.
// In the first move, change 2 to 3. nums becomes [5,3,3,4].
// In the second move, change 4 to 3. nums becomes [5,3,3,3].
// In the third move, change 5 to 3. nums becomes [3,3,3,3].
// After performing 3 moves, the difference between the minimum and maximum is 3 - 3 = 0.

// Example 2:
// Input: nums = [1,5,0,10,14]
// Output: 1
// Explanation: We can make at most 3 moves.
// In the first move, change 5 to 0. nums becomes [1,0,0,10,14].
// In the second move, change 10 to 0. nums becomes [1,0,0,0,14].
// In the third move, change 14 to 1. nums becomes [1,0,0,0,1].
// After performing 3 moves, the difference between the minimum and maximum is 1 - 0 = 1.
// It can be shown that there is no way to make the difference 0 in 3 moves.

// Example 3:
// Input: nums = [3,100,20]
// Output: 0
// Explanation: We can make at most 3 moves.
// In the first move, change 100 to 7. nums becomes [3,7,20].
// In the second move, change 20 to 7. nums becomes [3,7,7].
// In the third move, change 3 to 7. nums becomes [7,7,7].
// After performing 3 moves, the difference between the minimum and maximum is 7 - 7 = 0.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"
import "sort"

func minDifference(nums []int) int {
    n := len(nums)
    if n <= 3 {
        return 0
    }
    sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
    min := func (x, y int) int { if x < y { return x; }; return y; }
    maxDiff := nums[n-1] - nums[0] // 找出最大差值
    for i, j := 3, n - 1; i >= 0; i, j = i - 1, j - 1 {
        maxDiff = min(maxDiff, nums[j] - nums[i])
    }
    return maxDiff
}

func minDifference1(nums []int) int {
    res, k, picked := 1 << 32 - 1, 3, make([]bool, len(nums))
    getMaxKNums := func (nums []int, picked []bool, k int) []int {
        res := make([]int, 0, k)
        for i := 0; i < k; i++ {
            index, mx := -1, -1 << 32 -1
            for idx, v := range nums {
                if v > mx && !picked[idx] {
                    index = idx
                    mx = v
                }
            }
            if index != -1 {
                res = append(res, mx)
                picked[index] = true
            }
        }
        return res
    }
    getMinKNums := func(nums []int, picked []bool, k int) []int {
        res := make([]int, 0, k)
        for i := 0; i < k; i++ {
            index, mn := -1, 1 << 32 - 1
            for idx, v := range nums {
                if v < mn && !picked[idx] {
                    index = idx
                    mn = v
                }
            }
            if index != -1 {
                res = append(res, mn)
                picked[index] = true
            }
        }
        return res
    }
    mxs, mns := getMaxKNums(nums, picked, k+1), getMinKNums(nums, picked, k+1)
    if len(mxs)+len(mns) < k + 2 {
        return 0
    }
    for i := k; i >= 0; i-- {
        mns = append(mns, mxs[i])
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i <= k; i++ {
        res = min(res, mns[len(mns) - 1 - i] - mns[k-i])
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,3,2,4]
    // Output: 0
    // Explanation: We can make at most 3 moves.
    // In the first move, change 2 to 3. nums becomes [5,3,3,4].
    // In the second move, change 4 to 3. nums becomes [5,3,3,3].
    // In the third move, change 5 to 3. nums becomes [3,3,3,3].
    // After performing 3 moves, the difference between the minimum and maximum is 3 - 3 = 0.
    fmt.Println(minDifference([]int{5,3,2,4})) // 0
    // Example 2:
    // Input: nums = [1,5,0,10,14]
    // Output: 1
    // Explanation: We can make at most 3 moves.
    // In the first move, change 5 to 0. nums becomes [1,0,0,10,14].
    // In the second move, change 10 to 0. nums becomes [1,0,0,0,14].
    // In the third move, change 14 to 1. nums becomes [1,0,0,0,1].
    // After performing 3 moves, the difference between the minimum and maximum is 1 - 0 = 1.
    // It can be shown that there is no way to make the difference 0 in 3 moves.
    fmt.Println(minDifference([]int{1,5,0,10,14})) // 1
    // Example 3:
    // Input: nums = [3,100,20]
    // Output: 0
    // Explanation: We can make at most 3 moves.
    // In the first move, change 100 to 7. nums becomes [3,7,20].
    // In the second move, change 20 to 7. nums becomes [3,7,7].
    // In the third move, change 3 to 7. nums becomes [7,7,7].
    // After performing 3 moves, the difference between the minimum and maximum is 7 - 7 = 0.
    fmt.Println(minDifference([]int{3,100,20})) // 0

    fmt.Println(minDifference1([]int{5,3,2,4})) // 0
    fmt.Println(minDifference1([]int{1,5,0,10,14})) // 1
    fmt.Println(minDifference1([]int{3,100,20})) // 0
}