package main

// 面试题 05.04. Closed Number LCCI
// Given a positive integer, print the next smallest and the next largest number that have the same number of 1 bits in their binary representation.

// Example1:
// Input: num = 2 (0b10)
// Output: [4, 1] ([0b100, 0b1])

// Example2:
// Input: num = 1
// Output: [2, -1]

// Note:
//     1 <= num <= 2147483647
//     If there is no next smallest or next largest number, output -1.

import "fmt"

func findClosedNumbers(num int) []int {
    res, count := []int{-1, -1 }, num & 1
    for i := 1; i < 31; i++ {
        if res[0] == -1 && (num & (1 << i)) == 0 && (num & (1 << (i - 1))) != 0 {
            res[0] = num ^ (1 << i) ^ (1 << (i - 1))
            if count - 1 > 0 {
                res[0] = res[0] &^ ((1 << (i - 1)) - 1)
                res[0] |= ((1 << (count - 1)) - 1)
            }
        }
        if res[1] == -1 && (num&(1 << i)) != 0 && (num & (1 <<(i - 1))) == 0 {
            res[1] = num ^ (1 << i) ^ (1 << (i - 1))
            if count > 0 {
                res[1] = res[1] &^ ((1 << (i - 1)) - 1)
                res[1] |= ((1 << (count)) - 1) << (i - 1 - count)
            }
        }
        if num & (1 << i) != 0 {
            count++
        }
    }
    return res
}

func main() {
    // Example1:
    // Input: num = 2 (0b10)
    // Output: [4, 1] ([0b100, 0b1])
    fmt.Println(findClosedNumbers(2)) // [4, 1]
    // Example2:
    // Input: num = 1
    // Output: [2, -1]
    fmt.Println(findClosedNumbers(1)) // [2, -1]

    fmt.Println(findClosedNumbers(8)) // [16 4]
    fmt.Println(findClosedNumbers(999)) // [1003 990]
    fmt.Println(findClosedNumbers(1024)) // [2048 512]
    fmt.Println(findClosedNumbers(999_999_9999)) // [10000000511 9999998972]
    fmt.Println(findClosedNumbers(1_000_000_000)) // [1000000512 999999744]
    fmt.Println(findClosedNumbers((1 << 31) - 1)) // [-1, -1]
    fmt.Println(findClosedNumbers(1 << 31)) // [-1, -1]
}