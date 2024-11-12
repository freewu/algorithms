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

// 对于n，整个数组的和，提取公因式就是 0 * (任意或的和) + 1 * (任意或的和) + ... + (n - 1) * (任意或的和)
func maxSizedArray(s int64) int {
    low, high := 1, 200000
    getResult := func(n int) int64 {
        sum, subone := int64(0), int64(n - 1)
        for i := 0; i <= 20; i++ { // 每一位遍历, 这里取的20，可以更小，但懒得算具体多小了
            bit, zeroCount := int64(1 << i), int64(0) // 0 到 n-1 中，此位是 0 的数字个数
            mod := subone % bit // 完整周期之外的
            round := (subone - mod) / bit // 完整周期数
            half := (round >> 1) // 有一半的周期是0
            zeroCount += half * bit
            //fmt.Println(round & 1)
            if (round & 1) == 1 {
                zeroCount += bit // 如果是奇数个完整周期，最后一个周期，对应的位也是0
            } else {
                zeroCount += (mod + 1) // 如果是偶数个完整周期， 多出来的，对应位都是0
            }
            count := int64(n * n) - (zeroCount * zeroCount) // 所有两两或操作的结果中，此位1的数量
            sum += (bit * count) // 累加到或操作的和里去
        }
        return int64((n - 1) * n / 2) * sum // 整个数组的和
    }
    for low <= high {
        mid := (low + high) / 2
        v := getResult(mid)
        fmt.Println("get result: ",v, " low: ", low, " high: ", high)
        if  v <= s {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return high
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

    fmt.Println(maxSizedArray(1024)) // 1
    fmt.Println(maxSizedArray(999999999)) // 1
}