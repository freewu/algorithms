package main

// 3388. Count Beautiful Splits in an Array
// You are given an array nums.

// A split of an array nums is beautiful if:
//     1. The array nums is split into three subarrays: nums1, nums2, and nums3, such that nums can be formed by concatenating nums1, nums2, and nums3 in that order.
//     2. The subarray nums1 is a prefix of nums2 OR nums2 is a prefix of nums3.

// Return the number of ways you can make this split.

// Example 1:
// Input: nums = [1,1,2,1]
// Output: 2
// Explanation:
// The beautiful splits are:
// A split with nums1 = [1], nums2 = [1,2], nums3 = [1].
// A split with nums1 = [1], nums2 = [1], nums3 = [2,1].

// Example 2:
// Input: nums = [1,2,3,4]
// Output: 0
// Explanation:
// There are 0 beautiful splits.

// Constraints:
//     1 <= nums.length <= 5000
//     0 <= nums[i] <= 50

import "fmt"

// 超出时间限制 582 / 583 
func beautifulSplits(nums []int) int {
    res, n := 0, len(nums)
    if n < 3 { return 0 }
    check := func(arr []int, l, start int) bool {
        if start + 2 * l > len(arr) { return false }
        for i := start; i < start + l; i++ {
            if arr[i] != arr[l + i] {
                return false
            }
        }
        return true
    }
    for i := 1; i < n - 1; i++ {
        for j := i + 1; j < n; j++ {
            if j >= 2 * i && check(nums, i, 0) {
                res += n - 2 * i
                break
            }
            if n - j >= j - i && check(nums, j - i, i) {
                res++
                continue
            }
        }
    }
    return res
}

func beautifulSplits1(nums []int) int {
    res, n := 0, len(nums)
    lcp := make([][]int, n+1)
    for i := range lcp {
        lcp[i] = make([]int, n+1)
    }
    for i := n - 1; i >= 0; i-- {
        for j := n - 1; j > i; j-- {
            if nums[i] == nums[j] {
                lcp[i][j] = lcp[i + 1][j + 1] + 1
            }
        }
    }
    for i := 1; i < n - 1; i++ {
        for j := i + 1; j < n; j++ {
            if (i <= j-i && lcp[0][i] >= i) || (j - i <= n - j && lcp[i][j] >= j - i) {
                res++
            }
        }
    }
    return res
}

func beautifulSplits2(nums []int) int {
    res, n := 0, len(nums)
    calc := func(nums []int) []int {
        n, left, right := len(nums), 0, 0
        z := make([]int, n)
        for i := 1; i < n; i++ {
            if i >= right {
                // 不在zbox中，直接查找新的zbox
                j, k := 0, i
                for ; k < n && nums[j] == nums[k]; k, j = k+1, j+1 { }
                left, right = i, k
                z[i] = j
            } else {
                // i在box中
                if z[i-left] < right-i {
                    // z[i-l]对应的数组没有超出zbox
                    z[i] = z[i-left]
                } else {
                    // 正好活超出zbox，重新计算zbox
                    j, k := right - i, right
                    for ; k < n && nums[j] == nums[k]; k, j = k + 1, j + 1 {
                    }
                    left, right = i, k
                    z[i] = j
                }
            }
        }
        return z
    }
    z := calc(nums)
    for i := 1; i < n-1; i++ {
        z0 := calc(nums[i:])
        for j := i + 1; j < n; j++ {
            if i <= j - i && z[i] >= i {
                res++
            } else if j - i <= n - j {
                if z0[j - i] >= j - i {
                    res++
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,2,1]
    // Output: 2
    // Explanation:
    // The beautiful splits are:
    // A split with nums1 = [1], nums2 = [1,2], nums3 = [1].
    // A split with nums1 = [1], nums2 = [1], nums3 = [2,1].
    fmt.Println(beautifulSplits([]int{1,1,2,1})) // 2
    // Example 2:
    // Input: nums = [1,2,3,4]
    // Output: 0
    // Explanation:
    // There are 0 beautiful splits.
    fmt.Println(beautifulSplits([]int{1,2,3,4})) // 0

    fmt.Println(beautifulSplits([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(beautifulSplits([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(beautifulSplits1([]int{1,1,2,1})) // 2
    fmt.Println(beautifulSplits1([]int{1,2,3,4})) // 0
    fmt.Println(beautifulSplits1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(beautifulSplits1([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(beautifulSplits2([]int{1,1,2,1})) // 2
    fmt.Println(beautifulSplits2([]int{1,2,3,4})) // 0
    fmt.Println(beautifulSplits2([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(beautifulSplits2([]int{9,8,7,6,5,4,3,2,1})) // 0
}