package main

// 2333. Minimum Sum of Squared Difference
// You are given two positive 0-indexed integer arrays nums1 and nums2, both of length n.

// The sum of squared difference of arrays nums1 and nums2 is defined as the sum of (nums1[i] - nums2[i])2 for each 0 <= i < n.

// You are also given two positive integers k1 and k2. 
// You can modify any of the elements of nums1 by +1 or -1 at most k1 times. 
// Similarly, you can modify any of the elements of nums2 by +1 or -1 at most k2 times.

// Return the minimum sum of squared difference after modifying array nums1 at most k1 times and modifying array nums2 at most k2 times.

// Note: You are allowed to modify the array elements to become negative integers.

// Example 1:
// Input: nums1 = [1,2,3,4], nums2 = [2,10,20,19], k1 = 0, k2 = 0
// Output: 579
// Explanation: The elements in nums1 and nums2 cannot be modified because k1 = 0 and k2 = 0. 
// The sum of square difference will be: (1 - 2)2 + (2 - 10)2 + (3 - 20)2 + (4 - 19)2 = 579.

// Example 2:
// Input: nums1 = [1,4,10,12], nums2 = [5,8,6,9], k1 = 1, k2 = 1
// Output: 43
// Explanation: One way to obtain the minimum sum of square difference is: 
// - Increase nums1[0] once.
// - Increase nums2[2] once.
// The minimum of the sum of square difference will be: 
// (2 - 5)2 + (4 - 8)2 + (10 - 7)2 + (12 - 9)2 = 43.
// Note that, there are other ways to obtain the minimum of the sum of square difference, but there is no way to obtain a sum smaller than 43.

// Constraints:
//     n == nums1.length == nums2.length
//     1 <= n <= 10^5
//     0 <= nums1[i], nums2[i] <= 10^5
//     0 <= k1, k2 <= 10^9

import "fmt"
import "sort"

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
    res, mx, k, n := 0, 0,  k1 + k2, len(nums1)
    freq := make(map[int]int)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        diff := abs(nums1[i] - nums2[i])
        freq[diff]++
        mx = max(mx, diff)
    }
    for mx > 0 && k > 0 {
        if v, exists := freq[mx]; exists {
            mn := min(k, v)
            k -= mn
            freq[mx] -= mn
            freq[mx - 1] += mn
        }
        mx--
    }
    for diff, count := range freq {
        res += (diff * diff * count)
    }
    return int64(res)
}

func minSumSquareDiff1(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    nums := make([]int, len(nums1))
    for i, a := range nums1 {
        b := nums2[i]
        nums[i] = abs(a - b)
    }
    sort.Ints(nums)
    res, i, k := 0, 0, k1 + k2
    for i = len(nums) - 2; k > 0 && i >= 0; i-- {
        reduction := (nums[i+1] - nums[i]) * (len(nums)-1 - i)
        if reduction <= k {
            k -= reduction
        } else {
            break
        }
    }
    // Run reduction
    target := nums[i+1] - k / (len(nums) - 1 - i)
    if target < 0 {
        target = 0
    }
    rest := k % (len(nums)-1 - i)
    for j := i+1; j < len(nums); j++ {
        nums[j] = target
        if rest > 0 && nums[j] > 0 {
            nums[j]--
            rest--
        }
    }
    for _, v := range nums {
        res += (v * v)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums1 = [1,2,3,4], nums2 = [2,10,20,19], k1 = 0, k2 = 0
    // Output: 579
    // Explanation: The elements in nums1 and nums2 cannot be modified because k1 = 0 and k2 = 0. 
    // The sum of square difference will be: (1 - 2)2 + (2 - 10)2 + (3 - 20)2 + (4 - 19)2 = 579.
    fmt.Println(minSumSquareDiff([]int{1,2,3,4}, []int{2,10,20,19}, 0, 0)) // 579
    // Example 2:
    // Input: nums1 = [1,4,10,12], nums2 = [5,8,6,9], k1 = 1, k2 = 1
    // Output: 43
    // Explanation: One way to obtain the minimum sum of square difference is: 
    // - Increase nums1[0] once.
    // - Increase nums2[2] once.
    // The minimum of the sum of square difference will be: 
    // (2 - 5)2 + (4 - 8)2 + (10 - 7)2 + (12 - 9)2 = 43.
    // Note that, there are other ways to obtain the minimum of the sum of square difference, but there is no way to obtain a sum smaller than 43.
    fmt.Println(minSumSquareDiff([]int{1,4,10,12}, []int{5,8,6,9}, 1, 1)) // 43

    fmt.Println(minSumSquareDiff1([]int{1,2,3,4}, []int{2,10,20,19}, 0, 0)) // 579
    fmt.Println(minSumSquareDiff1([]int{1,4,10,12}, []int{5,8,6,9}, 1, 1)) // 43
}