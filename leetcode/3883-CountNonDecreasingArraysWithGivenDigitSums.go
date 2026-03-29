package main

// 3883. Count Non Decreasing Arrays With Given Digit Sums
// You are given an integer array digitSum of length n.

// An array arr of length n is considered valid if:
//     1. 0 <= arr[i] <= 5000
//     2. it is non-decreasing.
//     3. the sum of the digits of arr[i] equals digitSum[i].

// Return an integer denoting the number of distinct valid arrays. 
// Since the answer may be large, return it modulo 10^9 + 7.

// An array is said to be non-decreasing if each element is greater than or equal to the previous element, if it exists.

// Example 1:
// Input: digitSum = [25,1]
// Output: 6
// Explanation:
// Numbers whose sum of digits is 25 are 799, 889, 898, 979, 988, and 997.
// The only number whose sum of digits is 1 that can appear after these values while keeping the array non-decreasing is 1000.
// Thus, the valid arrays are [799, 1000], [889, 1000], [898, 1000], [979, 1000], [988, 1000], and [997, 1000].
// Hence, the answer is 6.

// Example 2:
// Input: digitSum = [1]
// Output: 4
// Explanation:
// The valid arrays are [1], [10], [100], and [1000].
// Thus, the answer is 4.

// Example 3:
// Input: digitSum = [2,49,23]
// Output: 0
// Explanation:
// There is no integer in the range [0, 5000] whose sum of digits is 49. Thus, the answer is 0.

// Constraints:
//     1 <= digitSum.length <= 1000
//     0 <= digitSum[i] <= 50

import "fmt"

const MX = 5001
const MOD = 1_000_000_007
const maxDigitSum = 31 // 4999 的数位和最大
var sums [maxDigitSum + 1][]int

func init() {
    sum := [MX]int{}
    for x := range sum {
        // 去掉 x 的个位，问题变成 x/10 的数位和，即 digSum[x/10]
        sum[x] = sum[x/10] + x%10
        sums[sum[x]] = append(sums[sum[x]], x)
    }
}

func countArrays(digitSum []int) int {
    res, pre := 0, 0
    f := [MX]int{1} // f[x] 表示以 x 结尾的有效数组的个数

    for _, v := range digitSum {
        if v > maxDigitSum { return 0 }
        a := sums[pre]
        j, m, sum := 0, len(a), 0
        for _, x := range sums[v] {
            // 有效数组的前一个数只要 <= x 就行
            for ; j < m && a[j] <= x; j++ {
                sum += f[a[j]]
            }
            // sum 现在就是以 x 结尾的有效数组的个数
            f[x] = sum % MOD
        }
        pre = v // 记录上一个数位和
    }
    for _, v := range sums[pre] {
        res += f[v]
    }
    return res % MOD
}

func main() {
    // Example 1:
    // Input: digitSum = [25,1]
    // Output: 6
    // Explanation:
    // Numbers whose sum of digits is 25 are 799, 889, 898, 979, 988, and 997.
    // The only number whose sum of digits is 1 that can appear after these values while keeping the array non-decreasing is 1000.
    // Thus, the valid arrays are [799, 1000], [889, 1000], [898, 1000], [979, 1000], [988, 1000], and [997, 1000].
    // Hence, the answer is 6.
    fmt.Println(countArrays([]int{25,1})) // 6
    // Example 2:
    // Input: digitSum = [1]
    // Output: 4
    // Explanation:
    // The valid arrays are [1], [10], [100], and [1000].
    // Thus, the answer is 4.
    fmt.Println(countArrays([]int{1})) // 4
    // Example 3:
    // Input: digitSum = [2,49,23]
    // Output: 0
    // Explanation:
    // There is no integer in the range [0, 5000] whose sum of digits is 49. Thus, the answer is 0.
    fmt.Println(countArrays([]int{2,49,23})) // 0

    fmt.Println(countArrays([]int{1,2,3,4,5,6,7,8,9})) // 407295959
    fmt.Println(countArrays([]int{9,8,7,6,5,4,3,2,1})) // 53
}