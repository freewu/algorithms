package main

// 119. Pascal’s Triangle II
// Given an integer rowIndex, return the rowIndexth row of the Pascal’s triangle.
// Notice that the row index starts from 0.
// In Pascal’s triangle, each number is the sum of the two numbers directly above it.

// Follow up:
// Could you optimize your algorithm to use only O(k) extra space?

// Example 1:
// Input: rowIndex = 3
// Output: [1,3,3,1]

// Example 2:
// Input: rowIndex = 0
// Output: [1]

// Example 3:
// Input: rowIndex = 1
// Output: [1,1]

// Constraints:
//     0 <= rowIndex <= 33

// 解题思路:
// 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
// 要求我们只能使用 O(k) 的空间
// row[i] = row[i-1] * (rowLen - i  + 1) / m

// [
//         [1],
//     [1,1],
//     [1,2,1],
//     [1,3,3,1],
//     [1,4,6,4,1]
// ]

// Input: 0
// Output: [1]

// Input: 1
// Output: [1,1]

// Input: 2
// Output: [1,2,1]

// Input: 3
// Output: [1,3,3,1]

// Input: 4
// Output: [1,4,6,4,1]

import "fmt"

func getRow(rowIndex int) []int {
    res := make([]int, rowIndex + 1) // 只能使用 O(k) 的空间
    res[0] = 1 // 第一个是 1
    for i := 1; i <= rowIndex; i++ {
        res[i] = res[i-1] * (rowIndex - i + 1) / i // 杨辉三角，每个数字是 (a+b)^n 二项式展开的系数
    }
    return res
}

func main() {
    // Example 1:
    // Input: rowIndex = 3
    // Output: [1,3,3,1]
    fmt.Printf("getRow(3) = %v\n",getRow(3)) // [1 3 3 1]
    // Example 2:
    // Input: rowIndex = 0
    // Output: [1]
    fmt.Printf("getRow(0) = %v\n",getRow(0)) // [1]
    // Example 3:
    // Input: rowIndex = 1
    // Output: [1,1]
    fmt.Printf("getRow(1) = %v\n",getRow(1)) // [1,1]

    fmt.Printf("getRow(0) = %v\n",getRow(0)) // [1]
    fmt.Printf("getRow(1) = %v\n",getRow(1)) // [1,1]
    fmt.Printf("getRow(2) = %v\n",getRow(2)) // [1 2 1]
    fmt.Printf("getRow(3) = %v\n",getRow(3)) // [1 3 3 1]
    fmt.Printf("getRow(4) = %v\n",getRow(4)) // [1 4 6 4 1]
    fmt.Printf("getRow(5) = %v\n",getRow(5)) // [1 5 10 10 5 1]
    fmt.Printf("getRow(6) = %v\n",getRow(6)) // [1 6 15 20 15 6 1]
    fmt.Printf("getRow(7) = %v\n",getRow(7)) // [1 7 21 35 35 21 7 1]
    fmt.Printf("getRow(8) = %v\n",getRow(8)) // [1 8 28 56 70 56 28 8 1]

    fmt.Printf("getRow(33) = %v\n",getRow(33)) // [1 33 528 5456 40920 237336 1107568 4272048 13884156 38567100 92561040 193536720 354817320 573166440 818809200 1037158320 1166803110 1166803110 1037158320 818809200 573166440 354817320 193536720 92561040 38567100 13884156 4272048 1107568 237336 40920 5456 528 33 1]
}
