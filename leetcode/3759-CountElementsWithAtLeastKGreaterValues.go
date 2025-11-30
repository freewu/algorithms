package main

// 3759. Count Elements With at Least K Greater Values
// You are given an integer array nums of length n and an integer k.

// An element in nums is said to be qualified if there exist at least k elements in the array that are strictly greater than it.

// Return an integer denoting the total number of qualified elements in nums.

// Example 1:
// Input: nums = [3,1,2], k = 1
// Output: 2
// Explanation:
// The elements 1 and 2 each have at least k = 1 element greater than themselves.
// ​​​​​​​No element is greater than 3. Therefore, the answer is 2.

// Example 2:
// Input: nums = [5,5,5], k = 2
// Output: 0
// Explanation:
// Since all elements are equal to 5, no element is greater than the other. Therefore, the answer is 0.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     0 <= k < n

import "fmt"
import "sort"

func countElements(nums []int, k int) int {
    n := len(nums)
    if k == 0 {
        return n
    }
    sort.Ints(nums)
    return sort.SearchInts(nums, nums[n-k]) // 小于第 k 大的元素个数
}

func countElements1(nums []int, k int) int {
    freq := make(map[int]int)
    for _, v := range nums {
        freq[v]++
    }
    keys := make([]int, 0, len(freq))
    for v := range freq {
        keys = append(keys, v)
    }
    sort.Ints(keys)
    res, total := 0, len(nums)
    for _, key := range keys {
        v := freq[key]
        total -= v
        if total >= k { 
            res += v
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,1,2], k = 1
    // Output: 2
    // Explanation:
    // The elements 1 and 2 each have at least k = 1 element greater than themselves.
    // ​​​​​​​No element is greater than 3. Therefore, the answer is 2.
    fmt.Println(countElements([]int{3,1,2}, 1)) // 2
    // Example 2:
    // Input: nums = [5,5,5], k = 2
    // Output: 0
    // Explanation:
    // Since all elements are equal to 5, no element is greater than the other. Therefore, the answer is 0.
    fmt.Println(countElements([]int{5,5,5}, 2)) // 0

    fmt.Println(countElements([]int{1,2,3,4,5,6,7,8,9}, 2)) // 7
    fmt.Println(countElements([]int{9,8,7,6,5,4,3,2,1}, 2)) // 7

    fmt.Println(countElements1([]int{3,1,2}, 1)) // 2
    fmt.Println(countElements1([]int{5,5,5}, 2)) // 0
    fmt.Println(countElements1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 7
    fmt.Println(countElements1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 7
}