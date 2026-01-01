package main

// 66. Plus One
// Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.

// You may assume the integer do not contain any leading zero, except the number 0 itself.

// The digits are stored such that the most significant digit is at the head of the list.

// Example 1:
// Input: digits = [1,2,3]
// Output: [1,2,4]
// Explanation: The array represents the integer 123.
// Incrementing by one gives 123 + 1 = 124.
// Thus, the result should be [1,2,4].

// Example 2:
// Input: digits = [4,3,2,1]
// Output: [4,3,2,2]
// Explanation: The array represents the integer 4321.
// Incrementing by one gives 4321 + 1 = 4322.
// Thus, the result should be [4,3,2,2].

// Example 3:
// Input: digits = [9]
// Output: [1,0]
// Explanation: The array represents the integer 9.
// Incrementing by one gives 9 + 1 = 10.
// Thus, the result should be [1,0].

// Constraints:
//     1 <= digits.length <= 100
//     0 <= digits[i] <= 9
//     digits does not contain any leading 0's.

import "fmt"

// 我最早的实现
func plusOne(digits []int) []int {
    var l = len(digits)
    for i := l - 1; i >= 0; i-- {
        if (digits[i] + 1) >= 10 {
            digits[i] = ((digits[i] + 1) % 10)
            // 如果是9999这样子的情况
            if 0 == i {
                var slice = make([]int, l+1)
                slice[0] = 1
                return slice
            }
        } else {
            digits[i] = digits[i] + 1
            return digits
        }
    }
    return digits
}

func plusOne1(digits []int) []int {
    for i := len(digits) - 1; i >= 0; i-- { // 从数组最后向数据前循环  低位 -> 高位
        digits[i]++
        if digits[i] != 10 { // 如果不需进位直接返回
            return digits
        }
        digits[i] = 0 // 有进位直接为 0
    }
    // 如果出现在 9...9 的请况(所有都要进一位) 能进入到这个
    digits[0] = 1 // 设置为 1  比如 9999 -> 10000
    digits = append(digits, 0) // 加个 0
    return digits
}

func plusOne2(digits []int) []int {
    n := len(digits)
    // 至少有一位不是 9
    for i := n - 1; i >= 0; i-- {
        if digits[i] != 9 {
            digits[i]++
            for j := i + 1; j < n; j++ {
                digits[j] = 0
            }
            return digits
        }
    }
    // 上述遍历完成说明全部是 9
    digits = make([]int, n+1)
    digits[0] = 1
    return digits
}

func main() {
    // Example 1:
    // Input: digits = [1,2,3]
    // Output: [1,2,4]
    // Explanation: The array represents the integer 123.
    // Incrementing by one gives 123 + 1 = 124.
    // Thus, the result should be [1,2,4].
    fmt.Printf("plusOne([]int{1, 2, 3, 4, 8, 9, 0, 3, 4}) = %v\n",plusOne([]int{1, 2, 3, 4, 8, 9, 0, 3, 4})) // [1 2 3 4 8 9 0 3 5]
    // Example 2:
    // Input: digits = [4,3,2,1]
    // Output: [4,3,2,2]
    // Explanation: The array represents the integer 4321.
    // Incrementing by one gives 4321 + 1 = 4322.
    // Thus, the result should be [4,3,2,2].
    fmt.Printf("plusOne([]int{1, 2, 3, 4, 8, 9, 0, 3, 9}) = %v\n",plusOne([]int{1, 2, 3, 4, 8, 9, 0, 3, 9})) // [1 2 3 4 8 9 0 4 0]
    // Example 3:
    // Input: digits = [9]
    // Output: [1,0]
    // Explanation: The array represents the integer 9.
    // Incrementing by one gives 9 + 1 = 10.
    // Thus, the result should be [1,0].
    fmt.Printf("plusOne([]int{9, 9}) = %v\n",plusOne([]int{9, 9})) // [1 0 0]

    fmt.Printf("plusOne([]int{1,2,3,4,5,6,7,8,9}) = %v\n",plusOne([]int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 9 0]
    fmt.Printf("plusOne([]int{9,8,7,6,5,4,3,2,1}) = %v\n",plusOne([]int{9,8,7,6,5,4,3,2,1})) // [9 8 7 6 5 4 3 2 2]

    fmt.Printf("plusOne1([]int{1, 2, 3, 4, 8, 9, 0, 3, 4}) = %v\n",plusOne1([]int{1, 2, 3, 4, 8, 9, 0, 3, 4})) // [1 2 3 4 8 9 0 3 5]
    fmt.Printf("plusOne1([]int{1, 2, 3, 4, 8, 9, 0, 3, 9}) = %v\n",plusOne1([]int{1, 2, 3, 4, 8, 9, 0, 3, 9})) // [1 2 3 4 8 9 0 4 0]
    fmt.Printf("plusOne1([]int{9, 9}) = %v\n",plusOne1([]int{9, 9})) // [1 0 0]
    fmt.Printf("plusOne1([]int{1,2,3,4,5,6,7,8,9}) = %v\n",plusOne1([]int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 9 0]
    fmt.Printf("plusOne1([]int{9,8,7,6,5,4,3,2,1}) = %v\n",plusOne1([]int{9,8,7,6,5,4,3,2,1})) // [9 8 7 6 5 4 3 2 2]

    fmt.Printf("plusOne2([]int{1, 2, 3, 4, 8, 9, 0, 3, 4}) = %v\n",plusOne2([]int{1, 2, 3, 4, 8, 9, 0, 3, 4})) // [1 2 3 4 8 9 0 3 5]
    fmt.Printf("plusOne2([]int{1, 2, 3, 4, 8, 9, 0, 3, 9}) = %v\n",plusOne2([]int{1, 2, 3, 4, 8, 9, 0, 3, 9})) // [1 2 3 4 8 9 0 4 0]
    fmt.Printf("plusOne2([]int{9, 9}) = %v\n",plusOne2([]int{9, 9})) // [1 0 0]
    fmt.Printf("plusOne2([]int{1,2,3,4,5,6,7,8,9}) = %v\n",plusOne2([]int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 9 0]
    fmt.Printf("plusOne2([]int{9,8,7,6,5,4,3,2,1}) = %v\n",plusOne2([]int{9,8,7,6,5,4,3,2,1})) // [9 8 7 6 5 4 3 2 2]
}
