package main

// 3356. Zero Array Transformation II
// You are given an integer array nums of length n and a 2D array queries where queries[i] = [li, ri, vali].

// Each queries[i] represents the following action on nums:
//     1. Decrement the value at each index in the range [li, ri] in nums by at most vali.
//     2. The amount by which each value is decremented can be chosen independently for each index.

// A Zero Array is an array with all its elements equal to 0.

// Return the minimum possible non-negative value of k, 
// such that after processing the first k queries in sequence, nums becomes a Zero Array. 
// If no such k exists, return -1.

// Example 1:
// Input: nums = [2,0,2], queries = [[0,2,1],[0,2,1],[1,1,3]]
// Output: 2
// Explanation:
//     For i = 0 (l = 0, r = 2, val = 1):
//         Decrement values at indices [0, 1, 2] by [1, 0, 1] respectively.
//         The array will become [1, 0, 1].
//     For i = 1 (l = 0, r = 2, val = 1):
//         Decrement values at indices [0, 1, 2] by [1, 0, 1] respectively.
//         The array will become [0, 0, 0], which is a Zero Array. Therefore, the minimum value of k is 2.

// Example 2:
// Input: nums = [4,3,2,1], queries = [[1,3,2],[0,2,1]]
// Output: -1
// Explanation:
//     For i = 0 (l = 1, r = 3, val = 2):
//         Decrement values at indices [1, 2, 3] by [2, 2, 1] respectively.
//         The array will become [4, 1, 0, 0].
//     For i = 1 (l = 0, r = 2, val = 1):
//         Decrement values at indices [0, 1, 2] by [1, 1, 0] respectively.
//         The array will become [3, 0, 0, 0], which is not a Zero Array.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 5 * 10^5
//     1 <= queries.length <= 10^5
//     queries[i].length == 3
//     0 <= li <= ri < nums.length
//     1 <= vali <= 5

import "fmt"

func minZeroArray(nums []int, queries [][]int) int {
    n := len(queries)
    res, left, right, allZero := -1, 0, n - 1, true
    for _, v := range nums {
        if v != 0 {
            allZero = false
            break
        }
    }
    if allZero { return 0 }
    isPossible := func(index int) bool {
        pos, prefix := make([]int,len(nums) + 1), 0
        for i := 0; i <= index; i++ {
            start, end, val := queries[i][0], queries[i][1], queries[i][2]
            pos[start] += val
            pos[end + 1] -= val
        }
        for i, v := range nums {
            prefix += pos[i]
            if prefix < v { return false }
        }
        return true
    }
    for left <= right {
        mid := left + (right - left) / 2
        if isPossible(mid) {
            res, right = mid, mid - 1
        } else {
            left = mid + 1
        }
    }
    if res != -1 { res++ }
    return res
}

func minZeroArray1(nums []int, queries [][]int) int {
    res, i, n, m  := 0,0, len(nums), len(queries)
    diff := make([]int, n + 1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i < n {
        for nums[i] + diff[i] > 0 && res < m {
            start, end, val := queries[res][0], queries[res][1], queries[res][2]
            start = max(start, i)
            res++
            if start <= end {
                diff[start] -= val
                diff[end + 1] += val
            }
        }
        if nums[i] + diff[i] > 0 { return -1 }
        i++
        if i < n {
            diff[i] += diff[i-1]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,0,2], queries = [[0,2,1],[0,2,1],[1,1,3]]
    // Output: 2
    // Explanation:
    //     For i = 0 (l = 0, r = 2, val = 1):
    //         Decrement values at indices [0, 1, 2] by [1, 0, 1] respectively.
    //         The array will become [1, 0, 1].
    //     For i = 1 (l = 0, r = 2, val = 1):
    //         Decrement values at indices [0, 1, 2] by [1, 0, 1] respectively.
    //         The array will become [0, 0, 0], which is a Zero Array. Therefore, the minimum value of k is 2.
    fmt.Println(minZeroArray([]int{2,0,2}, [][]int{{0,2,1},{0,2,1},{1,1,3}})) // 2
    // Example 2:
    // Input: nums = [4,3,2,1], queries = [[1,3,2],[0,2,1]]
    // Output: -1
    // Explanation:
    //     For i = 0 (l = 1, r = 3, val = 2):
    //         Decrement values at indices [1, 2, 3] by [2, 2, 1] respectively.
    //         The array will become [4, 1, 0, 0].
    //     For i = 1 (l = 0, r = 2, val = 1):
    //         Decrement values at indices [0, 1, 2] by [1, 1, 0] respectively.
    //         The array will become [3, 0, 0, 0], which is not a Zero Array.
    fmt.Println(minZeroArray([]int{4,3,2,1}, [][]int{{1,3,2},{0,2,1}})) // -1
    
    fmt.Println(minZeroArray1([]int{2,0,2}, [][]int{{0,2,1},{0,2,1},{1,1,3}})) // 2
    fmt.Println(minZeroArray1([]int{4,3,2,1}, [][]int{{1,3,2},{0,2,1}})) // -1
}