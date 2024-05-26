package main

// 390. Elimination Game
// You have a list arr of all integers in the range [1, n] sorted in a strictly increasing order. 
// Apply the following algorithm on arr:
//     Starting from left to right, remove the first number and every other number afterward until you reach the end of the list.
//     Repeat the previous step again, but this time from right to left, remove the rightmost number and every other number from the remaining numbers.
//     Keep repeating the steps again, alternating left to right and right to left, until a single number remains.

// Given the integer n, return the last number that remains in arr.

// Example 1:
// Input: n = 9
// Output: 6
// Explanation:
// arr = [1, 2, 3, 4, 5, 6, 7, 8, 9]
// arr = [2, 4, 6, 8]
// arr = [2, 6]
// arr = [6]

// Example 2:
// Input: n = 1
// Output: 1
 
// Constraints:
//     1 <= n <= 10^9

import "fmt"

func lastRemaining(n int) int {
    res := 1
    for isLeftToRight, step := true, 1; n > 1; n, isLeftToRight, step = n / 2, !isLeftToRight, step * 2 {
        if isLeftToRight || n%2 != 0 {
            res += step
        }
    }
    return res
}

// 模拟题。按照题意
//      第一轮从左往右删除数字，第二轮从右往左删除数字。题目要求最后剩下的数字，
//      模拟过程中不需要真的删除元素。只需要标记起始元素，该轮步长和方向即可。
//      最后总元素只剩下一个即为所求
func lastRemaining1(n int) int {
    start, dir, step := 1, true, 1
    for n > 1 {
        if dir { // 正向 
            start += step
        } else { // 反向
            if n % 2 == 1 {
                start += step
            }
        }
        dir = !dir
        n >>= 1
        step <<= 1
    }
    return start
}

func main() {
    // Example 1:
    // Input: n = 9
    // Output: 6
    // Explanation:
    // arr = [1, 2, 3, 4, 5, 6, 7, 8, 9]
    // arr = [2, 4, 6, 8]
    // arr = [2, 6]
    // arr = [6]
    fmt.Println(lastRemaining(9)) // 6
    // Example 2:
    // Input: n = 1
    // Output: 1
    fmt.Println(lastRemaining(1)) // 1

    fmt.Println(lastRemaining1(9)) // 6
    fmt.Println(lastRemaining1(1)) // 1
}