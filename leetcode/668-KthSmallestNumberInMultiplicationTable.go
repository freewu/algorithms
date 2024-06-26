package main

// 668. Kth Smallest Number in Multiplication Table
// Nearly everyone has used the Multiplication Table. 
// The multiplication table of size m x n is an integer matrix mat where mat[i][j] == i * j (1-indexed).
// Given three integers m, n, and k, return the kth smallest element in the m x n multiplication table.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/02/multtable1-grid.jpg" />
// Input: m = 3, n = 3, k = 5
// Output: 3
// Explanation: The 5th smallest number is 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/05/02/multtable2-grid.jpg" />
// Input: m = 2, n = 3, k = 6
// Output: 6
// Explanation: The 6th smallest number is 6.

// Constraints:
//     1 <= m, n <= 3 * 10^4
//     1 <= k <= m * n

import "fmt"
import "sort"

func findKthNumber(m int, n int, k int) int {
    if m > n {
        return findKthNumber(n, m, k)
    }
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = i + 1
    }
    res := sort.Search(n * m + 1, func(i int) bool {
        index, a := 0, 1
        for a <= m {
            t := sort.Search(n, func(j int) bool {
                return arr[j] * a > i
            })
            if t == 0 { break } 
            a++
            index += t           
        }
        return index >= k
    })
    return res
}

// 二分
func findKthNumber1(m int, n int, k int) int {
    check := func(x int) bool {
        count, r, c := 0, 1, n
        for r <= m {
            if r * c <= x {
                count += c
                if count >= k {
                    return false
                }
                r++
            } else {
                c--
            }
        }
        return true
    }
    left, right := 1, m * n + 1
    for left < right {
        mid := left + ((right - left) >> 1)
        if check(mid) {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/02/multtable1-grid.jpg" />
    // Input: m = 3, n = 3, k = 5
    // Output: 3
    // Explanation: The 5th smallest number is 3.
    fmt.Println(findKthNumber(3, 3, 5)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/05/02/multtable2-grid.jpg" />
    // Input: m = 2, n = 3, k = 6
    // Output: 6
    // Explanation: The 6th smallest number is 6.
    fmt.Println(findKthNumber(2, 3, 6)) // 6

    fmt.Println(findKthNumber1(3, 3, 5)) // 3
    fmt.Println(findKthNumber1(2, 3, 6)) // 6
}