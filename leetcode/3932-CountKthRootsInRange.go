package main

// 3932. Count K-th Roots in a Range
// You are given three integers l, r, and k.

// An integer y is said to be a perfect kth power if there exists an integer x such that y = xk.

// Return the number of integers y in the range [l, r] (inclusive) that are perfect kth powers.

// Example 1:
// Input: l = 1, r = 9, k = 3
// Output: 2
// Explanation:
// The perfect cubes in the range [1, 9] are:
// 1 = 13
// 8 = 23
// Hence, the answer is 2.

// Example 2:
// Input: l = 8, r = 30, k = 2
// Output: 3
// Explanation:
// The perfect squares in the range [8, 30] are:
// 9 = 32
// 16 = 42
// 25 = 52
// Hence, the answer is 3.

// Constraints:
//     0 <= l <= r <= 10^9
//     1 <= k <= 30

import "fmt"
import "math"

func countKthRoots(l int, r int, k int) int {
    if k == 1  {
        return r - l + 1
    }
    check := func(val int, k int,l int, r int) bool {
        p := int(math.Pow(float64(val), float64(k)))
        if p >= l && p <= r {
            return true
        }
        return false
    }
    res := 0
    for i := 1; int(math.Pow(float64(i),float64(k))) <= r; i++ {
        if check(i, k, l, r) {
            res++
        }
    }
    if l == 0 {
        res++
    }
    return res
}

func countKthRoots1(l int, r int, k int) int {
    count := func(limit int) int {
        if limit < 0 { return 0 }
        if limit == 0 { return 1 }
        if k == 1 { return limit + 1  }
        res, low, high := 0, 0, limit
        if high > 31622 {
            high = 31622
        }
        for low <= high {
            mid := low + (high - low) / 2
            v := 1
            flag := false
            for i := 0; i < k; i++ {
                if mid > 0 && v > limit/  mid {
                    flag = true
                    break
                }
                v *= mid
            }
            if !flag && v <= limit {
                res, low = mid, mid + 1
            } else {
                high = mid - 1
            }
        }
        return res + 1
    }
    return count(r) - count(l - 1)
}

func main() {
    // Example 1:
    // Input: l = 1, r = 9, k = 3
    // Output: 2
    // Explanation:
    // The perfect cubes in the range [1, 9] are:
    // 1 = 13
    // 8 = 23
    // Hence, the answer is 2.
    fmt.Println(countKthRoots(1, 9, 3)) // 2
    // 2:
    // Input: l = 8, r = 30, k = 2
    // Output: 3
    // Explanation:
    // The perfect squares in the range [8, 30] are:
    // 9 = 32
    // 16 = 42
    // 25 = 52
    // Hence, the answer is 3.
    fmt.Println(countKthRoots(8, 30, 2)) // 3

    fmt.Println(countKthRoots(0, 0, 1)) // 1
    fmt.Println(countKthRoots(0, 0, 30)) // 1
    fmt.Println(countKthRoots(0, 1_000_000_000, 1)) // 1000000001
    fmt.Println(countKthRoots(0, 1_000_000_000_000, 30)) // 3
    fmt.Println(countKthRoots(1_000_000_000, 1_000_000_000, 1)) // 1
    fmt.Println(countKthRoots(1_000_000_000, 1_000_000_000_000, 30)) // 1

    fmt.Println(countKthRoots1(1, 9, 3)) // 2
    fmt.Println(countKthRoots1(8, 30, 2)) // 3
    fmt.Println(countKthRoots1(0, 0, 1)) // 1
    fmt.Println(countKthRoots1(0, 0, 30)) // 1
    fmt.Println(countKthRoots1(0, 1_000_000_000, 1)) // 1000000001
    fmt.Println(countKthRoots1(0, 1_000_000_000_000, 30)) // 3
    fmt.Println(countKthRoots1(1_000_000_000, 1_000_000_000, 1)) // 1
    fmt.Println(countKthRoots1(1_000_000_000, 1_000_000_000_000, 30)) // 1
}