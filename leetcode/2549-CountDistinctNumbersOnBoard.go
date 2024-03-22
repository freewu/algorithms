package main

// 2549. Count Distinct Numbers on Board
// You are given a positive integer n, that is initially placed on a board. Every day, for 10^9 days, you perform the following procedure:
//     For each number x present on the board, find all numbers 1 <= i <= n such that x % i == 1.
//     Then, place those numbers on the board.

// Return the number of distinct integers present on the board after 10^9 days have elapsed.

// Note:
//     Once a number is placed on the board, it will remain on it until the end.
//     % stands for the modulo operation. For example, 14 % 3 is 2.
    

// Example 1:
// Input: n = 5
// Output: 4
// Explanation: Initially, 5 is present on the board. 
// The next day, 2 and 4 will be added since 5 % 2 == 1 and 5 % 4 == 1. 
// After that day, 3 will be added to the board because 4 % 3 == 1. 
// At the end of a billion days, the distinct numbers on the board will be 2, 3, 4, and 5. 

// Example 2:
// Input: n = 3
// Output: 2
// Explanation: 
// Since 3 % 2 == 1, 2 will be added to the board. 
// After a billion days, the only two distinct numbers on the board are 2 and 3. 
 
// Constraints:
//     1 <= n <= 100

import "fmt"

func distinctIntegers(n int) int {
    if n == 1  {
        return 1
    }
    return n - 1
}

func distinctIntegers1(n int) int {
    set := make(map[int]bool)
    queue := []int{n}
    for len(queue) > 0 {
        x := queue[0]
        if !set[x] {
            set[queue[0]] = true
            for i := 1; i <= n; i++ {
                if x % i == 1 {
                    queue = append(queue, i)
                }
            }
        }
        queue = queue[1:]
    }
    return len(set)
}

func main() {
    // Explanation: Initially, 5 is present on the board. 
    // The next day, 2 and 4 will be added since 5 % 2 == 1 and 5 % 4 == 1. 
    // After that day, 3 will be added to the board because 4 % 3 == 1. 
    // At the end of a billion days, the distinct numbers on the board will be 2, 3, 4, and 5. 
    fmt.Println(distinctIntegers(5)) // 4
    // Since 3 % 2 == 1, 2 will be added to the board. 
    // After a billion days, the only two distinct numbers on the board are 2 and 3. 
    fmt.Println(distinctIntegers(3)) // 2

    fmt.Println(distinctIntegers1(5)) // 4
    fmt.Println(distinctIntegers1(3)) // 2
}