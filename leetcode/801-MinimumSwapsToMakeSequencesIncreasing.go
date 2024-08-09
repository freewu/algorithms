package main

// 801. Minimum Swaps To Make Sequences Increasing
// You are given two integer arrays of the same length nums1 and nums2. 
// In one operation, you are allowed to swap nums1[i] with nums2[i].
//     For example, if nums1 = [1,2,3,8], and nums2 = [5,6,7,4], 
//     you can swap the element at i = 3 to obtain nums1 = [1,2,3,4] and nums2 = [5,6,7,8].

// Return the minimum number of needed operations to make nums1 and nums2 strictly increasing. 
// The test cases are generated so that the given input always makes it possible.

// An array arr is strictly increasing if and only if arr[0] < arr[1] < arr[2] < ... < arr[arr.length - 1].

// Example 1:
// Input: nums1 = [1,3,5,4], nums2 = [1,2,3,7]
// Output: 1
// Explanation: 
// Swap nums1[3] and nums2[3]. Then the sequences are:
// nums1 = [1, 3, 5, 7] and nums2 = [1, 2, 3, 4]
// which are both strictly increasing.

// Example 2:
// Input: nums1 = [0,3,5,8,9], nums2 = [2,1,4,6,9]
// Output: 1

// Constraints:
//     2 <= nums1.length <= 10^5
//     nums2.length == nums1.length
//     0 <= nums1[i], nums2[i] <= 2 * 10^5

import "fmt"

func minSwap(nums1 []int, nums2 []int) int {
    dp := make([][]int, len(nums1))
    for i, _ := range dp {
        dp[i] = make([]int, 2)
        dp[i][0], dp[i][1] = -1, -1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func (i, prevA, prevB, swap int) int
    dfs = func (i, prevA, prevB, swap int) int {
        if i == len(nums1) {
            return 0
        }
        if dp[i][swap] != -1 {
            return dp[i][swap]
        }
        res := 100000
        if nums1[i] > prevA && nums2[i] > prevB {
            res = dfs(i+1, nums1[i], nums2[i], 0)
        }
        if nums2[i] > prevA && nums1[i] > prevB {
            res = min(res, dfs(i+1, nums2[i], nums1[i], 1) + 1)
        }
        dp[i][swap] = res
        return res
    }
    return dfs(0, -1, -1, 0)
}

func minSwap1(nums1 []int, nums2 []int) int {
    n := len(nums1)
    dp := make([][2]int, n)
    dp[0][0], dp[0][1] = 0, 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < n; i++ {
        dp[i] = [2]int{n, n} // 答案不会超过 n，故初始化成 n 方便后面取 min
        if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
            dp[i][0] = dp[i-1][0]
            dp[i][1] = dp[i-1][1] + 1
        }
        if nums2[i-1] < nums1[i] && nums1[i-1] < nums2[i] {
            dp[i][0] = min(dp[i][0], dp[i-1][1])
            dp[i][1] = min(dp[i][1], dp[i-1][0]+1)
        }
    }
    return min(dp[n-1][0], dp[n-1][1])
}

func main() {
    // Example 1:
    // Input: nums1 = [1,3,5,4], nums2 = [1,2,3,7]
    // Output: 1
    // Explanation: 
    // Swap nums1[3] and nums2[3]. Then the sequences are:
    // nums1 = [1, 3, 5, 7] and nums2 = [1, 2, 3, 4]
    // which are both strictly increasing.
    fmt.Println(minSwap([]int{1,3,5,4}, []int{1,2,3,7})) // 1
    // Example 2:
    // Input: nums1 = [0,3,5,8,9], nums2 = [2,1,4,6,9]
    // Output: 1
    fmt.Println(minSwap([]int{0,3,5,8,9}, []int{2,1,4,6,9})) // 1

    fmt.Println(minSwap1([]int{1,3,5,4}, []int{1,2,3,7})) // 1
    fmt.Println(minSwap1([]int{0,3,5,8,9}, []int{2,1,4,6,9})) // 1
}