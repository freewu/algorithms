package main

// 1523. Count Odd Numbers in an Interval Range
// Given two non-negative integers low and high. 
// Return the count of odd numbers between low and high (inclusive).

// Example 1:
// Input: low = 3, high = 7
// Output: 3
// Explanation: The odd numbers between 3 and 7 are [3,5,7].

// Example 2:
// Input: low = 8, high = 10
// Output: 1
// Explanation: The odd numbers between 8 and 10 are [9].
 
// Constraints:
//     0 <= low <= high <= 10^9

import "fmt"

// 判断奇数个数，其实就和 (high-low) / 2 有关，无非就是相差1的关系，只需要个例验证就可以推出一般结论
// [1,2,3]        [1,3]    2  ((3 - 1) / 2) + 1 
// [1,2,3,4,5]    [1,3,5]  3  ((5 - 1) / 2) + 1 
// [1,2,3,4,5,6]  [1,3,5]  3  ((6 - 1) / 2) + 1 
// [3,4,5,6,7]    [3,5,7]  3  ((7 - 3) / 2) + 1 
// [3,4,5,6,7,8]  [3,5,7]  3  ((8 - 3) / 2) + 1 

// [2,3,4]         [3]      1  (4 - 2) / 2
// [4,5,6,7,8]     [5,7]    2  (8 - 4) / 2
// [2,3,4,5,6,7,8] [3,5,7]  3  (8 - 2) / 2

func countOdds(low int, high int) int {
    if low % 2 == 0 && high % 2 == 0 {
        return (high - low) / 2
    } else {
        return (high - low) / 2 + 1
    }
}

func countOdds1(low int, high int) int {
    if low % 2 == 0 {
        low++
    }
    if high % 2 == 0 {
        high--
    }
    return (high - low) / 2 + 1
}

func main() {
    // Explanation: The odd numbers between 3 and 7 are [3,5,7].
    fmt.Println(countOdds(3,7)) // 3
    // Explanation: The odd numbers between 8 and 10 are [9].
    fmt.Println(countOdds(8,10)) // 1

    fmt.Println(countOdds(0,0)) // 0
    fmt.Println(countOdds(0,1_000_000_000)) // 500000000
    fmt.Println(countOdds(1024,1_000_000_000)) // 499999488
    fmt.Println(countOdds(1_000_000_000,1_000_000_000)) // 0

    fmt.Println(countOdds1(3,7)) // 3
    fmt.Println(countOdds1(8,10)) // 1
    fmt.Println(countOdds1(0,0)) // 0
    fmt.Println(countOdds1(0,1_000_000_000)) // 500000000
    fmt.Println(countOdds1(1024,1_000_000_000)) // 499999488
    fmt.Println(countOdds1(1_000_000_000,1_000_000_000)) // 0
}