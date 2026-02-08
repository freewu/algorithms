package main

// 3836. Maximum Score Using Exactly K Pairs
// You are given two integer arrays nums1 and nums2 of lengths n and m respectively, and an integer k.

// You must choose exactly k pairs of indices (i1, j1), (i2, j2), ..., (ik, jk) such that:
//     1. 0 <= i1 < i2 < ... < ik < n
//     2. 0 <= j1 < j2 < ... < jk < m

// For each chosen pair (i, j), you gain a score of nums1[i] * nums2[j].

// The total score is the sum of the products of all selected pairs.

// Return an integer representing the maximum achievable total score.

// Example 1:
// Input: nums1 = [1,3,2], nums2 = [4,5,1], k = 2
// Output: 22
// Explanation:
// One optimal choice of index pairs is:
// (i1, j1) = (1, 0) which scores 3 * 4 = 12
// (i2, j2) = (2, 1) which scores 2 * 5 = 10
// This gives a total score of 12 + 10 = 22.

// Example 2:
// Input: nums1 = [-2,0,5], nums2 = [-3,4,-1,2], k = 2
// Output: 26
// Explanation:
// One optimal choice of index pairs is:
// (i1, j1) = (0, 0) which scores -2 * -3 = 6
// (i2, j2) = (2, 1) which scores 5 * 4 = 20
// The total score is 6 + 20 = 26.

// Example 3:
// Input: nums1 = [-3,-2], nums2 = [1,2], k = 2
// Output: -7
// Explanation:
// The optimal choice of index pairs is:
// (i1, j1) = (0, 0) which scores -3 * 1 = -3
// (i2, j2) = (1, 1) which scores -2 * 2 = -4
// The total score is -3 + (-4) = -7.

// Constraints:
//     1 <= n == nums1.length <= 100
//     1 <= m == nums2.length <= 100
//     -10^6 <= nums1[i], nums2[i] <= 10^6
//     1 <= k <= min(n, m)

import "fmt"

func maxScore(nums1, nums2 []int, k int) int64 {
    n, m, inf := len(nums1), len(nums2), 1 << 61
    memo := make([][][]int, k+1)
    for i := range memo {
        memo[i] = make([][]int, n)
        for j := range memo[i] {
            memo[i][j] = make([]int, m)
            for p := range memo[i][j] {
                memo[i][j][p] = -inf
            }
        }
    }
    var dfs func(int, int, int) int
    dfs = func(k, i, j int) int {
        if k == 0 { return 0 }// 选完
        if i+1 < k || j + 1 < k { // 剩余元素不足 k 个
            return -inf // 下面计算 max 不会取到 -inf
        }
        p := &memo[k][i][j]
        if *p != -inf { return *p } // 之前计算过
        res1 := dfs(k, i-1, j)                         // 不选 nums1[i]
        res2 := dfs(k, i, j-1)                         // 不选 nums2[j]
        res3 := dfs(k-1, i-1, j-1) + nums1[i]*nums2[j] // 选 nums1[i] 和 nums2[j]
        res := max(res1, res2, res3)
        *p = res // 记忆化
        return res
    }
    return int64(dfs(k, n-1, m-1))
}

func maxScore1(nums1 []int, nums2 []int, k int) int64 {
    n, m, inf := len(nums1), len(nums2), 1 << 61
    prefix, arr := make([][]int64, n),  make([][]int64, n)    
    for i := range prefix {
        prefix[i], arr[i] = make([]int64, m), make([]int64, m)
    }
    for h := 0; h < k; h ++ {
        for i := 0; i < n; i++ {
            for j := 0; j < m; j++ {
                val := int64(-inf)
                if i > 0 {
                    val = max(val, arr[i-1][j])
                }
                if j > 0 {
                    val = max(val, arr[i][j-1])
                }
                if h == 0 {
                    val = max(val, int64(nums1[i] * nums2[j]))
                } else if i > 0 && j > 0 && prefix[i-1][j-1] != int64(-inf) {
                    val = max(val, int64(nums1[i] * nums2[j]) + prefix[i-1][j-1])
                }
                arr[i][j] = val
                // fmt.Printf("h = %d i: %d, j: %d, val: %d\n", h, i, j, val)
            }
        }
        prefix, arr = arr, prefix
    }
    return prefix[n-1][m-1]
}

func main() {
    // Example 1:
    // Input: nums1 = [1,3,2], nums2 = [4,5,1], k = 2
    // Output: 22
    // Explanation:
    // One optimal choice of index pairs is:
    // (i1, j1) = (1, 0) which scores 3 * 4 = 12
    // (i2, j2) = (2, 1) which scores 2 * 5 = 10
    // This gives a total score of 12 + 10 = 22.
    fmt.Println(maxScore([]int{1,3,2}, []int{4,5,1}, 2)) // 22
    // Example 2:
    // Input: nums1 = [-2,0,5], nums2 = [-3,4,-1,2], k = 2
    // Output: 26
    // Explanation:
    // One optimal choice of index pairs is:
    // (i1, j1) = (0, 0) which scores -2 * -3 = 6
    // (i2, j2) = (2, 1) which scores 5 * 4 = 20
    // The total score is 6 + 20 = 26.
    fmt.Println(maxScore([]int{-2,0,5}, []int{-3,4,-1,2}, 2)) // 26 
    // Example 3:
    // Input: nums1 = [-3,-2], nums2 = [1,2], k = 2
    // Output: -7
    // Explanation:
    // The optimal choice of index pairs is:
    // (i1, j1) = (0, 0) which scores -3 * 1 = -3
    // (i2, j2) = (1, 1) which scores -2 * 2 = -4
    // The total score is -3 + (-4) = -7.
    fmt.Println(maxScore([]int{-3,-2}, []int{1,2}, 2)) // -7

    fmt.Println(maxScore([]int{-891015}, []int{805820}, 1)) // -717997707300
    fmt.Println(maxScore([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 145
    fmt.Println(maxScore([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 144
    fmt.Println(maxScore([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 144
    fmt.Println(maxScore([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 145

    fmt.Println(maxScore1([]int{1,3,2}, []int{4,5,1}, 2)) // 22
    fmt.Println(maxScore1([]int{-2,0,5}, []int{-3,4,-1,2}, 2)) // 26 
    fmt.Println(maxScore1([]int{-3,-2}, []int{1,2}, 2)) // -7
    fmt.Println(maxScore1([]int{-891015}, []int{805820}, 1)) // -717997707300
    fmt.Println(maxScore1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 145
    fmt.Println(maxScore1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 144
    fmt.Println(maxScore1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 144
    fmt.Println(maxScore1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 145
}