package main

// 2602. Minimum Operations to Make All Array Elements Equal
// You are given an array nums consisting of positive integers.

// You are also given an integer array queries of size m. 
// For the ith query, you want to make all of the elements of nums equal to queries[i]. 
// You can perform the following operation on the array any number of times:
//     Increase or decrease an element of the array by 1.

// Return an array answer of size m where answer[i] is the minimum number of operations to make all elements of nums equal to queries[i].

// Note that after each query the array is reset to its original state.

// Example 1:
// Input: nums = [3,1,6,8], queries = [1,5]
// Output: [14,10]
// Explanation: For the first query we can do the following operations:
// - Decrease nums[0] 2 times, so that nums = [1,1,6,8].
// - Decrease nums[2] 5 times, so that nums = [1,1,1,8].
// - Decrease nums[3] 7 times, so that nums = [1,1,1,1].
// So the total number of operations for the first query is 2 + 5 + 7 = 14.
// For the second query we can do the following operations:
// - Increase nums[0] 2 times, so that nums = [5,1,6,8].
// - Increase nums[1] 4 times, so that nums = [5,5,6,8].
// - Decrease nums[2] 1 time, so that nums = [5,5,5,8].
// - Decrease nums[3] 3 times, so that nums = [5,5,5,5].
// So the total number of operations for the second query is 2 + 4 + 1 + 3 = 10.

// Example 2:
// Input: nums = [2,9,6,3], queries = [10]
// Output: [20]
// Explanation: We can increase each value in the array to 10. The total number of operations will be 8 + 1 + 4 + 7 = 20.

// Constraints:
//     n == nums.length
//     m == queries.length
//     1 <= n, m <= 10^5
//     1 <= nums[i], queries[i] <= 10^9

import "fmt"
import "sort"

func minOperations(nums []int, queries []int) []int64 {
    n := len(nums)
    sort.Ints(nums)
    prefix := make([]int, n)
    prefix[0] = nums[0]
    for i := 1; i < n; i++ {
        prefix[i] = prefix[i-1] + nums[i]
    }
    res := make([]int64, len(queries))
    for i, query := range queries {
        index := sort.SearchInts(nums, query)
        if index == 0 {
            res[i] = int64(prefix[n - 1] - query * n)
        } else {
            res[i] = int64(prefix[n - 1] - query * (n - index) + query * index - prefix[index - 1] << 1)
        }
    }
    return res
}

func minOperations1(nums []int, queries []int) []int64 {
    n := len(nums)
    // slices.Sort(nums)
    sort.Ints(nums)
    prefix := make([]int, n + 1)
    for i := range nums {
        prefix[i + 1] = prefix[i] + nums[i]
    }
    res := make([]int64, len(queries))
    for i, q := range queries {
        // j, _ := slices.BinarySearch(nums, q)
        j := sort.SearchInts(nums, q)
        res[i] = int64((q * j - prefix[j]) + (prefix[n] - prefix[j] - q * (n - j)))
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,1,6,8], queries = [1,5]
    // Output: [14,10]
    // Explanation: For the first query we can do the following operations:
    // - Decrease nums[0] 2 times, so that nums = [1,1,6,8].
    // - Decrease nums[2] 5 times, so that nums = [1,1,1,8].
    // - Decrease nums[3] 7 times, so that nums = [1,1,1,1].
    // So the total number of operations for the first query is 2 + 5 + 7 = 14.
    // For the second query we can do the following operations:
    // - Increase nums[0] 2 times, so that nums = [5,1,6,8].
    // - Increase nums[1] 4 times, so that nums = [5,5,6,8].
    // - Decrease nums[2] 1 time, so that nums = [5,5,5,8].
    // - Decrease nums[3] 3 times, so that nums = [5,5,5,5].
    // So the total number of operations for the second query is 2 + 4 + 1 + 3 = 10.
    fmt.Println(minOperations([]int{3,1,6,8}, []int{1,5})) // [14,10]
    // Example 2:
    // Input: nums = [2,9,6,3], queries = [10]
    // Output: [20]
    // Explanation: We can increase each value in the array to 10. The total number of operations will be 8 + 1 + 4 + 7 = 20.
    fmt.Println(minOperations([]int{2,9,6,3}, []int{10})) // [20]

    fmt.Println(minOperations1([]int{3,1,6,8}, []int{1,5})) // [14,10]
    fmt.Println(minOperations1([]int{2,9,6,3}, []int{10})) // [20]
}