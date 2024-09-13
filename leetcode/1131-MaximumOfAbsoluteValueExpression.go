package main

// 1131. Maximum of Absolute Value Expression
// Given two arrays of integers with equal lengths, return the maximum value of:

// |arr1[i] - arr1[j]| + |arr2[i] - arr2[j]| + |i - j|

// where the maximum is taken over all 0 <= i, j < arr1.length.

// Example 1:
// Input: arr1 = [1,2,3,4], arr2 = [-1,4,5,6]
// Output: 13

// Example 2:
// Input: arr1 = [1,-2,-5,0,10], arr2 = [0,-2,-1,-7,-4]
// Output: 20

// Constraints:
//     2 <= arr1.length == arr2.length <= 40000
//     -10^6 <= arr1[i], arr2[i] <= 10^6

import "fmt"

// Brute force O(n^2) 超出时间限制 20 / 21 
func maxAbsValExpr(arr1 []int, arr2 []int) int {
    res := -1 << 31
    max := func (x, y int) int { if x > y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < len(arr1); i++ {
        for j := i+1; j < len(arr1); j++ {
            res = max(res, abs(i-j) + abs(arr1[i] - arr1[j]) + abs(arr2[i] - arr2[j])) // |arr1[i] - arr1[j]| + |arr2[i] - arr2[j]| + |i - j|
        }
    }
    return res
}

func maxAbsValExpr1(arr1 []int, arr2 []int) int {
    res, inf, signs := 0, 1 << 31, []int{ -1, 1 }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, s1 := range signs {
        for _, s2 := range signs {
            mx, mn := -inf, inf
            for i := 0; i < len(arr1); i++ {
                v := s1 * arr1[i] + s2 * arr2[i] + i
                mn, mx= min(v, mn), max(v, mx)
            }
            res = max(res, mx - mn)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr1 = [1,2,3,4], arr2 = [-1,4,5,6]
    // Output: 13
    fmt.Println(maxAbsValExpr([]int{1,2,3,4}, []int{-1,4,5,6})) // 13
    // Example 2:
    // Input: arr1 = [1,-2,-5,0,10], arr2 = [0,-2,-1,-7,-4]
    // Output: 20
    fmt.Println(maxAbsValExpr([]int{1,-2,-5,0,10}, []int{0,-2,-1,-7,-4})) // 20

    fmt.Println(maxAbsValExpr1([]int{1,2,3,4}, []int{-1,4,5,6})) // 13
    fmt.Println(maxAbsValExpr1([]int{1,-2,-5,0,10}, []int{0,-2,-1,-7,-4})) // 20
}