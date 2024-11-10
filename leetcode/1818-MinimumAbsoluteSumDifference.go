package main

// 1818. Minimum Absolute Sum Difference
// You are given two positive integer arrays nums1 and nums2, both of length n.

// The absolute sum difference of arrays nums1 and nums2 is defined as the sum of |nums1[i] - nums2[i]| for each 0 <= i < n (0-indexed).

// You can replace at most one element of nums1 with any other element in nums1 to minimize the absolute sum difference.

// Return the minimum absolute sum difference after replacing at most one element in the array nums1. 
// Since the answer may be large, return it modulo 10^9 + 7.

// |x| is defined as:
//     x if x >= 0, or
//     -x if x < 0.

// Example 1:
// Input: nums1 = [1,7,5], nums2 = [2,3,5]
// Output: 3
// Explanation: There are two possible optimal solutions:
// - Replace the second element with the first: [1,7,5] => [1,1,5], or
// - Replace the second element with the third: [1,7,5] => [1,5,5].
// Both will yield an absolute sum difference of |1-2| + (|1-3| or |5-3|) + |5-5| = 3.

// Example 2:
// Input: nums1 = [2,4,6,8,10], nums2 = [2,4,6,8,10]
// Output: 0
// Explanation: nums1 is equal to nums2 so no replacement is needed. This will result in an 
// absolute sum difference of 0.

// Example 3:
// Input: nums1 = [1,10,4,4,2,7], nums2 = [9,3,5,1,7,4]
// Output: 20
// Explanation: Replace the first element with the second: [1,10,4,4,2,7] => [10,10,4,4,2,7].
// This yields an absolute sum difference of |10-9| + |10-3| + |4-5| + |4-1| + |2-7| + |7-4| = 20

// Constraints:
//     n == nums1.length
//     n == nums2.length
//     1 <= n <= 10^5
//     1 <= nums1[i], nums2[i] <= 10^5

import "fmt"
import "sort"

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
    res, n, mod := 0, len(nums1), 1_000_000_007
    arr := make([][2]int, n)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        arr[i] = [2]int{ nums2[i], nums1[i]}
        res = (res + abs(nums2[i] - nums1[i])) % mod
    }
    sort.Ints(nums1)
    diff := 0
    for i := 0; i < n; i++ {
        index := sort.Search(n, func(j int) bool {
            return nums1[j] > arr[i][0]
        }) - 1
        sub := abs(arr[i][0] - arr[i][1])
        if index >= 0 {
            sub = min(sub, abs(arr[i][0] - nums1[index]))
        }
        if index != n - 1 {
            sub = min(sub, abs(arr[i][0] - nums1[index + 1]))
        }
        diff = max(diff, abs(sub  - abs(arr[i][0] - arr[i][1])))
    }
    return (res + mod - diff) % mod
}

func minAbsoluteSumDiff1(nums1 []int, nums2 []int) int {
    n, nums := len(nums1), make([]int, len(nums1))
    copy(nums, nums1)
    sort.Ints(nums)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    search := func(nums []int, target int) int {
        left, right := -1, len(nums)
        for left + 1 < right {
            mid := left + (right - left) >> 1
            if nums[mid] < target {
                left = mid
            } else {
                right = mid
            }
        }
        return right
    }
    sum, mx, mod := 0, 0, 1_000_000_007
    for i := 0;i < n;i++ {
        a, b := nums1[i], nums2[i]
        if a == b { continue }
        x := abs(a - b)
        sum += x
        r := min(search(nums, b), n - 1)
        nd := abs(b - nums[r])
        if r > 0 {
            nd = min(nd, abs(b - nums[r - 1]))
        }
        if x > nd {
            mx = max(mx, x - nd)
        }
    }
    return (sum - mx) % mod
}

func main() {
    // Example 1:
    // Input: nums1 = [1,7,5], nums2 = [2,3,5]
    // Output: 3
    // Explanation: There are two possible optimal solutions:
    // - Replace the second element with the first: [1,7,5] => [1,1,5], or
    // - Replace the second element with the third: [1,7,5] => [1,5,5].
    // Both will yield an absolute sum difference of |1-2| + (|1-3| or |5-3|) + |5-5| = 3.
    fmt.Println(minAbsoluteSumDiff([]int{1,7,5}, []int{2,3,5})) // 3
    // Example 2:
    // Input: nums1 = [2,4,6,8,10], nums2 = [2,4,6,8,10]
    // Output: 0
    // Explanation: nums1 is equal to nums2 so no replacement is needed. This will result in an 
    // absolute sum difference of 0.
    fmt.Println(minAbsoluteSumDiff([]int{2,4,6,8,10}, []int{2,4,6,8,10})) // 0
    // Example 3:
    // Input: nums1 = [1,10,4,4,2,7], nums2 = [9,3,5,1,7,4]
    // Output: 20
    // Explanation: Replace the first element with the second: [1,10,4,4,2,7] => [10,10,4,4,2,7].
    // This yields an absolute sum difference of |10-9| + |10-3| + |4-5| + |4-1| + |2-7| + |7-4| = 20
    fmt.Println(minAbsoluteSumDiff([]int{1,10,4,4,2,7}, []int{9,3,5,1,7,4})) // 20

    fmt.Println(minAbsoluteSumDiff1([]int{1,7,5}, []int{2,3,5})) // 3
    fmt.Println(minAbsoluteSumDiff1([]int{2,4,6,8,10}, []int{2,4,6,8,10})) // 0
    fmt.Println(minAbsoluteSumDiff1([]int{1,10,4,4,2,7}, []int{9,3,5,1,7,4})) // 20
}