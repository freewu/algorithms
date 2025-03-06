package main

// 3397. Maximum Number of Distinct Elements After Operations
// You are given an integer array nums and an integer k.

// You are allowed to perform the following operation on each element of the array at most once:
//     1. Add an integer in the range [-k, k] to the element.

// Return the maximum possible number of distinct elements in nums after performing the operations.

// Example 1:
// Input: nums = [1,2,2,3,3,4], k = 2
// Output: 6
// Explanation:
// nums changes to [-1, 0, 1, 2, 3, 4] after performing operations on the first four elements.

// Example 2:
// Input: nums = [4,4,4,4], k = 1
// Output: 3
// Explanation:
// By adding -1 to nums[0] and 1 to nums[1], nums changes to [3, 5, 4, 4].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     0 <= k <= 10^9

import "fmt"
import "sort"

func maxDistinctElements(nums []int, k int) int {
    sort.Ints(nums)
    res, last := 0, nums[0] - k - 1
    for _, v := range nums {
        if v - k > last {
            last = v - k
        } else if v + k <= last {
            continue
        } else {
            last++
        }
        res++
    }
    return res
}

func maxDistinctElements1(nums []int, k int) int {
    n := len(nums)
    if k * 2 + 1 >= n { return n }
    sort.Ints(nums)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, pre := 0, -1 << 31 // 记录每个人左边的人的位置
    for _, v := range nums {
        v = min(max(v - k, pre + 1), v + k)
        if v > pre {
            res++
            pre = v
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,3,3,4], k = 2
    // Output: 6
    // Explanation:
    // nums changes to [-1, 0, 1, 2, 3, 4] after performing operations on the first four elements.
    fmt.Println(maxDistinctElements([]int{1,2,2,3,3,4}, 2)) // 6
    // Example 2:
    // Input: nums = [4,4,4,4], k = 1
    // Output: 3
    // Explanation:
    // By adding -1 to nums[0] and 1 to nums[1], nums changes to [3, 5, 4, 4].
    fmt.Println(maxDistinctElements([]int{4,4,4,4}, 1)) // 3
    
    fmt.Println(maxDistinctElements([]int{1,2,3,4,5,6,7,8,9}, 1)) // 9
    fmt.Println(maxDistinctElements([]int{9,8,7,6,5,4,3,2,1}, 1)) // 9

    fmt.Println(maxDistinctElements1([]int{1,2,2,3,3,4}, 2)) // 6
    fmt.Println(maxDistinctElements1([]int{4,4,4,4}, 1)) // 3
    fmt.Println(maxDistinctElements1([]int{1,2,3,4,5,6,7,8,9}, 1)) // 9
    fmt.Println(maxDistinctElements1([]int{9,8,7,6,5,4,3,2,1}, 1)) // 9
}