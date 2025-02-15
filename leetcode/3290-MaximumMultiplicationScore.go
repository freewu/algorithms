package main

// 3290. Maximum Multiplication Score
// You are given an integer array a of size 4 and another integer array b of size at least 4.

// You need to choose 4 indices i0, i1, i2, and i3 from the array b such that i0 < i1 < i2 < i3. 
// Your score will be equal to the value a[0] * b[i0] + a[1] * b[i1] + a[2] * b[i2] + a[3] * b[i3].

// Return the maximum score you can achieve.

// Example 1:
// Input: a = [3,2,5,6], b = [2,-6,4,-5,-3,2,-7]
// Output: 26
// Explanation:
// We can choose the indices 0, 1, 2, and 5. The score will be 3 * 2 + 2 * (-6) + 5 * 4 + 6 * 2 = 26.

// Example 2:
// Input: a = [-1,4,5,-2], b = [-5,-1,-3,-2,-4]
// Output: -1
// Explanation:
// We can choose the indices 0, 1, 3, and 4. The score will be (-1) * (-5) + 4 * (-1) + 5 * (-2) + (-2) * (-4) = -1.

// Constraints:
//     a.length == 4
//     4 <= b.length <= 10^5
//     -10^5 <= a[i], b[i] <= 10^5

import "fmt"

func maxScore(a []int, b []int) int64 {
    n, inf := len(b), 1 << 48
    dp1, dp2, dp3, dp4 := make([]int, n), make([]int, n), make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        dp1[i], dp2[i], dp3[i] ,dp4[i] = -inf, -inf, -inf, -inf
    }
    dp1[0] = a[0] * b[0]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        dp1[i] = max(dp1[i - 1], a[0] * b[i])
    }
    for i := 1; i < n; i++ {
        dp2[i] = max(dp2[i - 1], dp1[i - 1] + a[1] * b[i])
    }
    for i := 2; i < n; i++ {
        dp3[i] = max(dp3[i - 1], dp2[ i- 1] + a[2] * b[i])
    }
    for i := 3; i < n; i++ {
        dp4[i] = max(dp4[i - 1], dp3[i - 1] + a[3] * b[i])
    }
    return int64(dp4[n - 1])
}

func maxScore1(a []int, b []int) int64 {
    dp := make([]int64, 5)
    for i := 1; i < 5; i++ {
        dp[i] = -1 << 48
    }
    max := func (x, y int64) int64 { if x > y { return x; }; return y; }
    for i := range b {
        vb := int64(b[i])
        for i := 3; i >= 0; i-- {
            va := int64(a[i])
            dp[i+1] = max(dp[i + 1], dp[i] + va * vb)
        }
    }
    return dp[4]
}

func maxScore2(a []int, b []int) int64 {
    d0, d1, d2, d3 := -1 << 48, -1 << 48, -1 << 48, -1 << 48
    for i:= 0 ; i < len(b); i ++ {
        val := b[i]
        v3, v2, v1, v0 := d2 + a[3] * val, d1 + a[2] * val, d0 + a[1] * val, a[0] * val
        if d3 < v3 { d3 = v3 }
        if d2 < v2 { d2 = v2 }
        if d1 < v1 { d1 = v1 }
        if d0 < v0 { d0 = v0 }
    }
    return int64(d3)
}

func main() {
    // Example 1:
    // Input: a = [3,2,5,6], b = [2,-6,4,-5,-3,2,-7]
    // Output: 26
    // Explanation:
    // We can choose the indices 0, 1, 2, and 5. The score will be 3 * 2 + 2 * (-6) + 5 * 4 + 6 * 2 = 26.
    fmt.Println(maxScore([]int{3,2,5,6}, []int{2,-6,4,-5,-3,2,-7})) // 26
    // Example 2:
    // Input: a = [-1,4,5,-2], b = [-5,-1,-3,-2,-4]
    // Output: -1
    // Explanation:
    // We can choose the indices 0, 1, 3, and 4. The score will be (-1) * (-5) + 4 * (-1) + 5 * (-2) + (-2) * (-4) = -1.
    fmt.Println(maxScore([]int{-1,4,5,-2}, []int{-5,-1,-3,-2,-4})) // -1

    fmt.Println(maxScore([]int{100000,100000,100000,100000}, []int{-100000,-100000,-100000,-100000})) // -40000000000
    fmt.Println(maxScore([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 80
    fmt.Println(maxScore([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 70
    fmt.Println(maxScore([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 220
    fmt.Println(maxScore([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 230

    fmt.Println(maxScore1([]int{3,2,5,6}, []int{2,-6,4,-5,-3,2,-7})) // 26
    fmt.Println(maxScore1([]int{-1,4,5,-2}, []int{-5,-1,-3,-2,-4})) // -1
    fmt.Println(maxScore1([]int{100000,100000,100000,100000}, []int{-100000,-100000,-100000,-100000})) // -40000000000
    fmt.Println(maxScore1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 80
    fmt.Println(maxScore1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 70
    fmt.Println(maxScore1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 220
    fmt.Println(maxScore1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 230

    fmt.Println(maxScore2([]int{3,2,5,6}, []int{2,-6,4,-5,-3,2,-7})) // 26
    fmt.Println(maxScore2([]int{-1,4,5,-2}, []int{-5,-1,-3,-2,-4})) // -1
    fmt.Println(maxScore2([]int{100000,100000,100000,100000}, []int{-100000,-100000,-100000,-100000})) // -40000000000
    fmt.Println(maxScore2([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 80
    fmt.Println(maxScore2([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 70
    fmt.Println(maxScore2([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 220
    fmt.Println(maxScore2([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 230
}