package main

// 1458. Max Dot Product of Two Subsequences
// Given two arrays nums1 and nums2.

// Return the maximum dot product between non-empty subsequences of nums1 and nums2 with the same length.

// A subsequence of a array is a new array which is formed 
// from the original array by deleting some (can be none) of the characters 
// without disturbing the relative positions of the remaining characters. 
// (ie, [2,3,5] is a subsequence of [1,2,3,4,5] while [1,5,3] is not).

// Example 1:
// Input: nums1 = [2,1,-2,5], nums2 = [3,0,-6]
// Output: 18
// Explanation: Take subsequence [2,-2] from nums1 and subsequence [3,-6] from nums2.
// Their dot product is (2*3 + (-2)*(-6)) = 18.

// Example 2:
// Input: nums1 = [3,-2], nums2 = [2,-6,7]
// Output: 21
// Explanation: Take subsequence [3] from nums1 and subsequence [7] from nums2.
// Their dot product is (3*7) = 21.

// Example 3:
// Input: nums1 = [-1,-1], nums2 = [1,1]
// Output: -1
// Explanation: Take subsequence [-1] from nums1 and subsequence [1] from nums2.
// Their dot product is -1.

// Constraints:
//     1 <= nums1.length, nums2.length <= 500
//     -1000 <= nums1[i], nums2[i] <= 1000

import "fmt"

func maxDotProduct(nums1 []int, nums2 []int) int {
    n, m := len(nums1), len(nums2)
    memo, mx := make([][]int, n), nums1[n-1] * nums2[m-1]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n-1; i >= 0; i-- {
        memo[i] = make([]int, m)
        mx = max(mx, nums1[i] * nums2[m-1])
        memo[i][m-1] = mx
    }
    mx = nums1[n-1] * nums2[m-1]
    for j := m - 1; j >= 0; j-- {
        mx = max(mx, nums1[n-1] * nums2[j])
        memo[n-1][j] = mx
    }
    for i := n-2; i >= 0; i-- {
        for j := m-2; j >= 0; j-- {
            product := nums1[i] * nums2[j]
            mx := product
            mx = max(mx, product + memo[i+1][j+1])
            mx = max(mx, max(memo[i][j+1], memo[i+1][j]))
            memo[i][j] = mx
        }
    }
    return memo[0][0]
}

func maxDotProduct1(nums1 []int, nums2 []int) int {
    n, m, inf := len(nums1), len(nums2), 1 << 31
    dp := make([][]int, n + 1)
    for i, _ := range dp {
        dp[i] = make([]int, m + 1)
    }
    for i := 0; i <= m; i++ {
        dp[0][i] = -inf
    }
    for j:=0;j<=n;j++{
        dp[j][0] = -inf
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, x := range nums1 {
        for j, y := range nums2 {
            val := x * y
            dp[i+1][j+1] = val
            if i > 0 { dp[i+1][j+1] = max(dp[i+1][j+1],dp[i][j+1]) }
            if j > 0 { dp[i+1][j+1] = max(dp[i+1][j+1],dp[i+1][j]) }
            if i > 0 && j > 0 { dp[i+1][j+1] = max(dp[i][j]+val,dp[i+1][j+1]) }
        }
    }
    return dp[n][m]
}

func maxDotProduct2(nums1, nums2 []int) int {
    n := len(nums2)
    dp := make([]int, n + 1)
    for i := range dp {
        dp[i] = -1 << 31
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, x := range nums1 {
        prev := dp[0]
        for j, y := range nums2 {
            tmp := dp[j+1]
            dp[j+1] = max(max(prev, 0)+x*y, max(dp[j+1], dp[j]))
            prev = tmp
        }
    }
    return dp[n]
}

func main() {
    // Example 1:
    // Input: nums1 = [2,1,-2,5], nums2 = [3,0,-6]
    // Output: 18
    // Explanation: Take subsequence [2,-2] from nums1 and subsequence [3,-6] from nums2.
    // Their dot product is (2*3 + (-2)*(-6)) = 18.
    fmt.Println(maxDotProduct([]int{2,1,-2,5}, []int{3,0,-6})) // 18
    // Example 2:
    // Input: nums1 = [3,-2], nums2 = [2,-6,7]
    // Output: 21
    // Explanation: Take subsequence [3] from nums1 and subsequence [7] from nums2.
    // Their dot product is (3*7) = 21.
    fmt.Println(maxDotProduct([]int{3,-2}, []int{2,-6,7})) // 21
    // Example 3:
    // Input: nums1 = [-1,-1], nums2 = [1,1]
    // Output: -1
    // Explanation: Take subsequence [-1] from nums1 and subsequence [1] from nums2.
    // Their dot product is -1.
    fmt.Println(maxDotProduct([]int{-1,-1}, []int{1,1})) // -1

    fmt.Println(maxDotProduct([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 285
    fmt.Println(maxDotProduct([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 236
    fmt.Println(maxDotProduct([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 236
    fmt.Println(maxDotProduct([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 285

    fmt.Println(maxDotProduct1([]int{2,1,-2,5}, []int{3,0,-6})) // 18
    fmt.Println(maxDotProduct1([]int{3,-2}, []int{2,-6,7})) // 21
    fmt.Println(maxDotProduct1([]int{-1,-1}, []int{1,1})) // -1
    fmt.Println(maxDotProduct1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 285
    fmt.Println(maxDotProduct1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 236
    fmt.Println(maxDotProduct1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 236
    fmt.Println(maxDotProduct1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 285

    fmt.Println(maxDotProduct2([]int{2,1,-2,5}, []int{3,0,-6})) // 18
    fmt.Println(maxDotProduct2([]int{3,-2}, []int{2,-6,7})) // 21
    fmt.Println(maxDotProduct2([]int{-1,-1}, []int{1,1})) // -1
    fmt.Println(maxDotProduct2([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 285
    fmt.Println(maxDotProduct2([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 236
    fmt.Println(maxDotProduct2([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 236
    fmt.Println(maxDotProduct2([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 285
}