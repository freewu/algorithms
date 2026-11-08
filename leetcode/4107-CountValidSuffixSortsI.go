package main

// 4107. Count Valid Suffix Sorts I
// You are given two integer arrays nums and target, both of length n.

// For each index i, consider sorting the suffix nums[i..n - 1] of the original array in non-increasing order.

// Return the number of indices i for which the resulting array is equal to target.

// A subarray is a contiguous non-empty sequence of elements within an array.

// A suffix of an array is a subarray that begins at any index and extends to the end of the array.

// Example 1:
// Input: nums = [1,2,3], target = [3,2,1]
// Output: 1
// Explanation:
//     i = 0: The resulting array is [3, 2, 1], which equals target.
//     i = 1: The resulting array is [1, 3, 2].
//     i = 2: The resulting array is [1, 2, 3].
// Only i = 0 is valid, so the answer is 1.

// Example 2:
// Input: nums = [1,5,6,4,6], target = [1,5,6,6,4]
// Output: 2
// Explanation:
// Choosing either i = 2 or i = 3 results in [1, 5, 6, 6, 4]. No other index produces target. 
// Therefore, the answer is 2.

// Example 3:
// Input: nums = [7,6,8], target = [8,6,7]
// Output: 0
// Explanation:
// Sorting the suffix starting at indices 0, 1, and 2 produces [8, 7, 6], [7, 8, 6], and [7, 6, 8], respectively. 
// None equals target, so the answer is 0.

// Constraints:
//     1 <= n == nums.length == target.length <= 10^{5}
//     1 <= nums[i], target[i] <= 10^{​​​​​​​5}​​​​​​​​​​​​​​

import "fmt"
import "slices"

func countSuffix(nums []int, target []int) int {
    res, pref, suff, n, diff := 0, 0, 0, len(nums), 0   
    for i := 0; i < n; i++ {
        if nums[i] == target[i] {
            pref++
        } else {
            break
        }
    }
    suff = n - 1
    for i := n - 2; i >= 0; i-- {
        if target[i] >= target[i+1] {
            suff = i
        } else {
            break
        }
    }
    freq := make([]int, 100001)
    update := func(x, delta int) {
        if freq[x] == 0 {
            diff++
        }
        freq[x] += delta
        if freq[x] == 0 {
            diff--
        }
    }
    for j := suff; j < n; j++ {
        update(nums[j], 1)
        update(target[j], -1)
    }
    for i := suff; i <= pref && i < n; i++ {
        if diff == 0 {
            res++
        }
        update(nums[i], -1)
        update(target[i], 1)
    }
    return res
}

func countSuffix1(nums []int, target []int) int {
    res, balance, diff, n := 0, 0, 0, len(nums)
    m := max(slices.Max(nums), slices.Max(target))
    for i := range n {
        if nums[i] != target[i] {
            diff++
        }
    }
    bals := make([]int, m+1)
    for i := n-1; i >= 0; i-- {
        if i < n-1 && target[i] < target[i+1] {
            break
        }
        if nums[i] != target[i] {
            diff--
        }
        v := nums[i]
        bals[v]++
        if bals[v] == 0 {
            balance--
        } else if bals[v] == 1 {
            balance++
        }
        tar := target[i]
        bals[tar]--
        if bals[tar] == 0 {
            balance--
        } else if bals[tar] == -1 {
            balance++
        }
        if balance == 0 && diff == 0 {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3], target = [3,2,1]
    // Output: 1
    // Explanation:
    //     i = 0: The resulting array is [3, 2, 1], which equals target.
    //     i = 1: The resulting array is [1, 3, 2].
    //     i = 2: The resulting array is [1, 2, 3].
    // Only i = 0 is valid, so the answer is 1.
    fmt.Println(countSuffix([]int{1,2,3}, []int{3,2,1})) // 1
    // Example 2:
    // Input: nums = [1,5,6,4,6], target = [1,5,6,6,4]
    // Output: 2
    // Explanation:
    // Choosing either i = 2 or i = 3 results in [1, 5, 6, 6, 4]. No other index produces target. 
    // Therefore, the answer is 2.
    fmt.Println(countSuffix([]int{1,5,6,4,6}, []int{1,5,6,6,4})) // 2
    // Example 3:
    // Input: nums = [7,6,8], target = [8,6,7]
    // Output: 0
    // Explanation:
    // Sorting the suffix starting at indices 0, 1, and 2 produces [8, 7, 6], [7, 8, 6], and [7, 6, 8], respectively. 
    // None equals target, so the answer is 0.
    fmt.Println(countSuffix([]int{7,6,8}, []int{8,6,7})) // 0

    fmt.Println(countSuffix([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(countSuffix([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 1
    fmt.Println(countSuffix([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(countSuffix([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(countSuffix1([]int{1,2,3}, []int{3,2,1})) // 1
    fmt.Println(countSuffix1([]int{1,5,6,4,6}, []int{1,5,6,6,4})) // 2
    fmt.Println(countSuffix1([]int{7,6,8}, []int{8,6,7})) // 0
    fmt.Println(countSuffix1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(countSuffix1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 1
    fmt.Println(countSuffix1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(countSuffix1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 9
}