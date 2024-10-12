package main

// 2856. Minimum Array Length After Pair Removals
// Given an integer array num sorted in non-decreasing order.

// You can perform the following operation any number of times:
//     1. Choose two indices, i and j, where nums[i] < nums[j].
//     2. Then, remove the elements at indices i and j from nums. 
//        The remaining elements retain their original order, and the array is re-indexed.

// Return the minimum length of nums after applying the operation zero or more times.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: 0
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/18/tcase1.gif" />

// Example 2:
// Input: nums = [1,1,2,2,3,3]
// Output: 0
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/19/tcase2.gif" />

// Example 3:
// Input: nums = [1000000000,1000000000]
// Output: 2
// Explanation:
// Since both numbers are equal, they cannot be removed.

// Example 4:
// Input: nums = [2,3,4,4,4]
// Output: 1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/19/tcase3.gif" />

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     nums is sorted in non-decreasing order.

import "fmt"
import "sort"

func minLengthAfterRemovals(nums []int) int {
    mp := make(map[int]int)
    for _, v := range nums {
        mp[v] += 1
    }
    mx := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range mp {
        mx = max(mx, v)
    }
    if mx <= len(nums) / 2 {
        if len(nums) % 2 != 0 {
            return 1
        } else {
            return 0
        }
    }
    return mx * 2 - len(nums)
}

func minLengthAfterRemovals1(nums []int) int {
    n := len(nums)
    m := sort.SearchInts(nums ,nums[n/2] + 1) - sort.SearchInts(nums, nums[n/2])
    if m * 2 > n {
        return m * 2 - n
    }
    return n & 1
}

func minLengthAfterRemovals2(nums []int) int {
    n := len(nums)
    if n < 2 { return n }
    res, count, mx := []int{}, 1, 0
    for i := 1; i < n; i ++ {
        if nums[i] != nums[i - 1] {
            res = append(res, count)
            if count > mx { mx = count }
            count = 1
        } else {
            count ++
        }
    }
    res = append(res, count)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    if count > mx { mx = count }
    if len(res) == 1 {
        return n
    } else if len(res) == 2 {
        return abs(res[0] - res[1])
    }
    if n - mx >= mx { return n % 2 } 
    return mx - (n - mx)
}


func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: 0
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/18/tcase1.gif" />
    fmt.Println(minLengthAfterRemovals([]int{1,2,3,4})) // 0
    // Example 2:
    // Input: nums = [1,1,2,2,3,3]
    // Output: 0
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/19/tcase2.gif" />
    fmt.Println(minLengthAfterRemovals([]int{1,1,2,2,3,3})) // 0
    // Example 3:
    // Input: nums = [1000000000,1000000000]
    // Output: 2
    // Explanation:
    // Since both numbers are equal, they cannot be removed.
    fmt.Println(minLengthAfterRemovals([]int{1000000000,1000000000})) // 2
    // Example 4:
    // Input: nums = [2,3,4,4,4]
    // Output: 1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/19/tcase3.gif" />
    fmt.Println(minLengthAfterRemovals([]int{2,3,4,4,4})) // 1

    fmt.Println(minLengthAfterRemovals1([]int{1,2,3,4})) // 0
    fmt.Println(minLengthAfterRemovals1([]int{1,1,2,2,3,3})) // 0
    fmt.Println(minLengthAfterRemovals1([]int{1000000000,1000000000})) // 2
    fmt.Println(minLengthAfterRemovals1([]int{2,3,4,4,4})) // 1

    fmt.Println(minLengthAfterRemovals2([]int{1,2,3,4})) // 0
    fmt.Println(minLengthAfterRemovals2([]int{1,1,2,2,3,3})) // 0
    fmt.Println(minLengthAfterRemovals2([]int{1000000000,1000000000})) // 2
    fmt.Println(minLengthAfterRemovals2([]int{2,3,4,4,4})) // 1
}