package main

// 3344. Maximum Sized Array
// Given a positive integer s, let A be a 3D array of dimensions n × n × n, where each element A[i][j][k] is defined as:
//     A[i][j][k] = i * (j OR k), where 0 <= i, j, k < n.

// Return the maximum possible value of n such that the sum of all elements in array A does not exceed s.

// Example 1:
// Input: s = 10
// Output: 2
// Explanation:
// Elements of the array A for n = 2:
// A[0][0][0] = 0 * (0 OR 0) = 0
// A[0][0][1] = 0 * (0 OR 1) = 0
// A[0][1][0] = 0 * (1 OR 0) = 0
// A[0][1][1] = 0 * (1 OR 1) = 0
// A[1][0][0] = 1 * (0 OR 0) = 0
// A[1][0][1] = 1 * (0 OR 1) = 1
// A[1][1][0] = 1 * (1 OR 0) = 1
// A[1][1][1] = 1 * (1 OR 1) = 1
// The total sum of the elements in array A is 3, which does not exceed 10, so the maximum possible value of n is 2.

// Example 2:
// Input: s = 0
// Output: 1
// Explanation:
// Elements of the array A for n = 1:
// A[0][0][0] = 0 * (0 OR 0) = 0
// The total sum of the elements in array A is 0, which does not exceed 0, so the maximum possible value of n is 1.

// Constraints:
//     0 <= s <= 10^15

import "fmt"

func maxSizedArray(s int64) int {
    left, right := 0, 1200
    calc := func(n int) int64 {
        if n == 0 { return 0 }
        res := make([]int, 32)
        //  计算 0～n 每个数二进制表示时，１的个数
        for j := 0; j < n + 1; j++ {
            for i := 0; i < 32 ; i++ {
                if (1 << i & j) != 0 {
                    res[i]++
                }
            }
        }
        // 计算sum([j|k for j in range(n+1) for k in range(n+1)])
        sum := 0
        for i := 0; i < 32 ; i++ {
            c0, c1 := n + 1 - res[i], res[i]
            res[i] = c0 * res[i] + c1 * (n + 1)
            sum += (res[i] * 1 << i)
        }
        return int64((n * n + n) * sum / 2)
    }
    for left <= right {
        mid := (left + right) / 2
        if calc(mid) > s {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return right + 1
}

func maxSizedArray1(s int64) int {
    mx := 1330
    facts := make([]int64, mx)
    init := func () {
        facts[0] = 0
        for i := 1; i < mx; i++ {
            facts[i] = facts[i-1] + int64(i)
            for j := 0; j < i; j++ {
                facts[i] += 2 * int64(i|j)
            }
        }
    }
    init()
    left, right := 1, mx
    for left < right {
        mid := (left + right + 1) >> 1
        if facts[mid - 1] * int64(mid - 1) * int64(mid) / 2 <= s {
            left = mid
        } else {
            right = mid - 1
        }
    }
    return left
}

func main() {
    // Example 1:
    // Input: s = 10
    // Output: 2
    // Explanation:
    // Elements of the array A for n = 2:
    // A[0][0][0] = 0 * (0 OR 0) = 0
    // A[0][0][1] = 0 * (0 OR 1) = 0
    // A[0][1][0] = 0 * (1 OR 0) = 0
    // A[0][1][1] = 0 * (1 OR 1) = 0
    // A[1][0][0] = 1 * (0 OR 0) = 0
    // A[1][0][1] = 1 * (0 OR 1) = 1
    // A[1][1][0] = 1 * (1 OR 0) = 1
    // A[1][1][1] = 1 * (1 OR 1) = 1
    // The total sum of the elements in array A is 3, which does not exceed 10, so the maximum possible value of n is 2.
    fmt.Println(maxSizedArray(10)) // 2
    // Example 2:
    // Input: s = 0
    // Output: 1
    // Explanation:
    // Elements of the array A for n = 1:
    // A[0][0][0] = 0 * (0 OR 0) = 0
    // The total sum of the elements in array A is 0, which does not exceed 0, so the maximum possible value of n is 1.
    fmt.Println(maxSizedArray(0)) // 1

    fmt.Println(maxSizedArray(1024)) // 5
    fmt.Println(maxSizedArray(999999999)) // 75

    fmt.Println(maxSizedArray1(10)) // 2
    fmt.Println(maxSizedArray1(0)) // 1
    fmt.Println(maxSizedArray1(1024)) // 5
    fmt.Println(maxSizedArray1(999999999)) // 75
}