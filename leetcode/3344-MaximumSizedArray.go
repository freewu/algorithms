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
    bitLength := func(n int64) int64 { // 返回表示该数字在二进制下的最小位数
        if n == 0 { return 0 }
        res := int64(0)
        for n != 0 {
            n >>= 1
            res++
        }
        return res
    }
    // [0, upper]中二进制第k(k>=0)位为1的数的个数. 即满足 `num & (1 << k) > 0` 的数的个数
    calc := func(upper, k int64) int64 {
        if k >= bitLength(upper) { return 0 } 
        res := upper / (1 << (k + 1)) * (1 << k)
        upper %= 1 << (k + 1)
        if upper >= 1 << k {
            res += upper - (1 << k) + 1
        }
        return res
    }
    check := func(mid int64) bool {
        pairs, sumOr := int64(mid * mid), int64(0)
        // 对于每一位 bit，计算在所有 (j OR k) 中该位为 1 的总次数
        for b := int64(0); b <= 60; b++ {
            ones := calc(mid - 1, b)
            zeros := mid - ones
            anyOnes := pairs - zeros * zeros
            sumOr += anyOnes * (1 << b)
        }
        sumI := mid * (mid - 1) / 2
        return sumI * sumOr <= s
    }
    left, right := int64(1), int64(1e16)
    for left <= right {
        mid := (left + right) / 2
        fmt.Println(check(mid), " left: ", left, " right:", right)
        if check(mid) {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return int(right)
}

// def calc(upper: int, k: int) -> int:
//     """[0, upper]中二进制第k(k>=0)位为1的数的个数.
//     即满足 `num & (1 << k) > 0` 的数的个数
//     """
//     if k >= upper.bit_length():
//         return 0
//     res = upper // (1 << (k + 1)) * (1 << k)
//     upper %= 1 << (k + 1)
//     if upper >= 1 << k:
//         res += upper - (1 << k) + 1
//     return res


// class Solution:
//     def maxSizedArray(self, s: int) -> int:
//         def check(mid: int) -> bool:
//             pairs = mid * mid
//             sumOr = 0
//             # 对于每一位 bit，计算在所有 (j OR k) 中该位为 1 的总次数
//             for b in range(60):
//                 ones = calc(mid - 1, b)
//                 zeros = mid - ones
//                 anyOnes = pairs - zeros * zeros
//                 sumOr += anyOnes * (1 << b)
//             sumI = mid * (mid - 1) // 2
//             return sumI * sumOr <= s

//         left, right = 1, int(1e16)
//         while left <= right:
//             mid = (left + right) // 2
//             if check(mid):
//                 left = mid + 1
//             else:
//                 right = mid - 1
//         return right

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

    fmt.Println(maxSizedArray(1024)) // 1
    fmt.Println(maxSizedArray(999999999)) // 1
}