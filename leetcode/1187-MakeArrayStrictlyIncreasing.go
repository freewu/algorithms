package main

// 1187. Make Array Strictly Increasing
// Given two integer arrays arr1 and arr2, 
// return the minimum number of operations (possibly zero) needed to make arr1 strictly increasing.

// In one operation, you can choose two indices 0 <= i < arr1.length and 0 <= j < arr2.length 
// and do the assignment arr1[i] = arr2[j].

// If there is no way to make arr1 strictly increasing, return -1.

// Example 1:
// Input: arr1 = [1,5,3,6,7], arr2 = [1,3,2,4]
// Output: 1
// Explanation: Replace 5 with 2, then arr1 = [1, 2, 3, 6, 7].

// Example 2:
// Input: arr1 = [1,5,3,6,7], arr2 = [4,3,1]
// Output: 2
// Explanation: Replace 5 with 3 and then replace 3 with 4. arr1 = [1, 3, 4, 6, 7].

// Example 3:
// Input: arr1 = [1,5,3,6,7], arr2 = [1,6,3,3]
// Output: -1
// Explanation: You can't make arr1 strictly increasing.
 
// Constraints:
//     1 <= arr1.length, arr2.length <= 2000
//     0 <= arr1[i], arr2[i] <= 10^9

import "fmt"
import "sort"

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
    sort.Ints(arr2)
    m, inf := 0, 1 << 32
    for _, x := range arr2 {
        if m == 0 || x != arr2[ m - 1] {
            arr2[m] = x
            m++
        }
    }
    // fmt.Printf("m = %v, arr2 = %v \r\n", m, arr2)
    arr2 = arr2[:m]
    arr1 = append([]int{-inf}, arr1...)
    arr1 = append(arr1, inf)
    n := len(arr1)
    dp := make([]int, n)
    for i := range dp {
        dp[i] = inf
    }
    dp[0] = 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < n; i++ {
        if arr1[i-1] < arr1[i] {
            dp[i] = dp[i-1]
        }
        j := sort.SearchInts(arr2, arr1[i])
        for k := 1; k <= min(i-1, j); k++ {
            if arr1[i-k-1] < arr2[j-k] {
                dp[i] = min(dp[i], dp[i-k-1] + k)
            }
        }
    }
    if dp[n-1] >= inf {
        return -1
    }
    return dp[n-1]
}

func main() {
    // Example 1:
    // Input: arr1 = [1,5,3,6,7], arr2 = [1,3,2,4]
    // Output: 1
    // Explanation: Replace 5 with 2, then arr1 = [1, 2, 3, 6, 7].
    fmt.Println(makeArrayIncreasing([]int{1,5,3,6,7},[]int{1,3,2,4})) // 1 [1, 2, 3, 6, 7]
    // Example 2:
    // Input: arr1 = [1,5,3,6,7], arr2 = [4,3,1]
    // Output: 2
    // Explanation: Replace 5 with 3 and then replace 3 with 4. arr1 = [1, 3, 4, 6, 7].
    fmt.Println(makeArrayIncreasing([]int{1,5,3,6,7},[]int{4,3,1})) // 2 [1, 3, 4, 6, 7]
    // Example 3:
    // Input: arr1 = [1,5,3,6,7], arr2 = [1,6,3,3]
    // Output: -1
    // Explanation: You can't make arr1 strictly increasing.
    fmt.Println(makeArrayIncreasing([]int{1,5,3,6,7},[]int{1,6,3,3})) // -1
}