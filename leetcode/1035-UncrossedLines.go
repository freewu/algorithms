package main

// 1035. Uncrossed Lines
// You are given two integer arrays nums1 and nums2. 
// We write the integers of nums1 and nums2 (in the order they are given) on two separate horizontal lines.
// We may draw connecting lines: a straight line connecting two numbers nums1[i] and nums2[j] such that:
//     nums1[i] == nums2[j], and
//     the line we draw does not intersect any other connecting (non-horizontal) line.

// Note that a connecting line cannot intersect even at the endpoints (i.e., each number can only belong to one connecting line).
// Return the maximum number of connecting lines we can draw in this way.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/04/26/142.png" />
// Input: nums1 = [1,4,2], nums2 = [1,2,4]
// Output: 2
// Explanation: We can draw 2 uncrossed lines as in the diagram.
// We cannot draw 3 uncrossed lines, because the line from nums1[1] = 4 to nums2[2] = 4 will intersect the line from nums1[2]=2 to nums2[1]=2.

// Example 2:
// Input: nums1 = [2,5,1,2,5], nums2 = [10,5,2,1,5,2]
// Output: 3

// Example 3:
// Input: nums1 = [1,3,7,1,7,5], nums2 = [1,9,2,5,1]
// Output: 2
 
// Constraints:
//     1 <= nums1.length, nums2.length <= 500
//     1 <= nums1[i], nums2[j] <= 2000

import "fmt"

// dp
func maxUncrossedLines(nums1 []int, nums2 []int) int {
    l1, l2 := len(nums1), len(nums2)
    dp := make([][]int, l1 + 1)
    for i := l1; i >= 0; i-- {
        dp[i] = make([]int, l2 + 1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := l1 - 1; i >= 0; i-- {
        for j := l2- 1; j >= 0; j-- {
            if nums1[i] == nums2[j] { // 连线
                dp[i][j] = 1 + dp[i+1][j+1]
            } else {
                dp[i][j] = max(dp[i+1][j], dp[i][j + 1])
            }
        }
    }
    return dp[0][0]
}

func maxUncrossedLines1(nums1 []int, nums2 []int) int {
    l1, l2 := len(nums1), len(nums2)
    res, dp := 0, make([][]int, l1)
    for i := 0; i < l1; i++ {
        dp[i]= make([]int, l2)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < l1; i++ {
        for j := 0; j <l2; j++ {
            if nums1[i] == nums2[j] && (i == 0 || j == 0) {// 相交且有一个是起始点
                dp[i][j]= 1
            } else if nums1[i] == nums2[j] {
                dp[i][j] = 1+ dp[i-1][j-1]
            } else {
                x, y := 0, 0
                if i > 0 {
                    x = dp[i-1][j]
                }
                if j > 0{
                    y = dp[i][j-1]
                }
                dp[i][j] = max(x, y)
            }
            res = max(res, dp[i][j])
        }
    }
    return res
}

func maxUncrossedLines2(nums1, nums2 []int) int {
    m, n := len(nums1), len(nums2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range nums1 {
        for j, w := range nums2 {
            if v == w {
                dp[i+1][j+1] = dp[i][j] + 1
            } else {
                dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
            }
        }
    }
    return dp[m][n]
}

func main() {
    // Explanation: We can draw 2 uncrossed lines as in the diagram.
    // We cannot draw 3 uncrossed lines, because the line from nums1[1] = 4 to nums2[2] = 4 will intersect the line from nums1[2]=2 to nums2[1]=2.
    fmt.Println(maxUncrossedLines([]int{1,4,2},[]int{1,2,4})) // 2
    fmt.Println(maxUncrossedLines([]int{2,5,1,2,5},[]int{10,5,2,1,5,2})) // 3
    fmt.Println(maxUncrossedLines([]int{1,3,7,1,7,5},[]int{1,9,2,5,1})) // 2

    fmt.Println(maxUncrossedLines1([]int{1,4,2},[]int{1,2,4})) // 2
    fmt.Println(maxUncrossedLines1([]int{2,5,1,2,5},[]int{10,5,2,1,5,2})) // 3
    fmt.Println(maxUncrossedLines1([]int{1,3,7,1,7,5},[]int{1,9,2,5,1})) // 2

    fmt.Println(maxUncrossedLines2([]int{1,4,2},[]int{1,2,4})) // 2
    fmt.Println(maxUncrossedLines2([]int{2,5,1,2,5},[]int{10,5,2,1,5,2})) // 3
    fmt.Println(maxUncrossedLines2([]int{1,3,7,1,7,5},[]int{1,9,2,5,1})) // 2
}