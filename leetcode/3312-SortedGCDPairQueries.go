package main

// 3312. Sorted GCD Pair Queries
// You are given an integer array nums of length n and an integer array queries.

// Let gcdPairs denote an array obtained by calculating the GCD of all possible pairs (nums[i], nums[j]), where 0 <= i < j < n, and then sorting these values in ascending order.

// For each query queries[i], you need to find the element at index queries[i] in gcdPairs.

// Return an integer array answer, where answer[i] is the value at gcdPairs[queries[i]] for each query.

// The term gcd(a, b) denotes the greatest common divisor of a and b.

// Example 1:
// Input: nums = [2,3,4], queries = [0,2,2]
// Output: [1,2,2]
// Explanation:
// gcdPairs = [gcd(nums[0], nums[1]), gcd(nums[0], nums[2]), gcd(nums[1], nums[2])] = [1, 2, 1].
// After sorting in ascending order, gcdPairs = [1, 1, 2].
// So, the answer is [gcdPairs[queries[0]], gcdPairs[queries[1]], gcdPairs[queries[2]]] = [1, 2, 2].

// Example 2:
// Input: nums = [4,4,2,1], queries = [5,3,1,0]
// Output: [4,2,1,1]
// Explanation:
// gcdPairs sorted in ascending order is [1, 1, 1, 2, 2, 4].

// Example 3:
// Input: nums = [2,2], queries = [0,0]
// Output: [2,2]
// Explanation:
// gcdPairs = [2].

// Constraints:
//     2 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 5 * 10^4
//     1 <= queries.length <= 10^5
//     0 <= queries[i] < n * (n - 1) / 2

import "fmt"
import "slices"
import "sort"

func gcdValues(nums []int, queries []int64) []int {
    mx := slices.Max(nums)
    freq := make([]int, mx + 1)
    for _, v := range nums {
        freq[v]++
    }
    count := make([]int, mx + 1)
    for i := mx; i > 0; i-- {
        var c int
        for j := i; j <= mx; j += i {
            c += freq[j]
            count[i] -= count[j]
        }
        count[i] += c * (c - 1) / 2 // select two numbers from c numbers
    }
    for i := 2; i <= mx; i++ {
        count[i] += count[i - 1] // prefix sum
    }
    res:= make([]int, len(queries))
    for i, query := range queries {
        res[i] = sort.SearchInts(count, int(query) + 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,3,4], queries = [0,2,2]
    // Output: [1,2,2]
    // Explanation:
    // gcdPairs = [gcd(nums[0], nums[1]), gcd(nums[0], nums[2]), gcd(nums[1], nums[2])] = [1, 2, 1].
    // After sorting in ascending order, gcdPairs = [1, 1, 2].
    // So, the answer is [gcdPairs[queries[0]], gcdPairs[queries[1]], gcdPairs[queries[2]]] = [1, 2, 2].
    fmt.Println(gcdValues([]int{2,3,4}, []int64{0,2,2})) // [1,2,2]
    // Example 2:
    // Input: nums = [4,4,2,1], queries = [5,3,1,0]
    // Output: [4,2,1,1]
    // Explanation:
    // gcdPairs sorted in ascending order is [1, 1, 1, 2, 2, 4].
    fmt.Println(gcdValues([]int{4,4,2,1}, []int64{5,3,1,0})) // [4,2,1,1]
    // Example 3:
    // Input: nums = [2,2], queries = [0,0]
    // Output: [2,2]
    // Explanation:
    // gcdPairs = [2].
    fmt.Println(gcdValues([]int{2,2}, []int64{0,0})) // [2,2]
}