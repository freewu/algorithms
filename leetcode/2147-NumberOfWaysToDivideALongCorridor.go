package main

// 2147. Number of Ways to Divide a Long Corridor
// Along a long library corridor, there is a line of seats and decorative plants. 
// You are given a 0-indexed string corridor of length n consisting of letters 'S' and 'P' where each 'S' represents a seat and each 'P' represents a plant.

// One room divider has already been installed to the left of index 0, and another to the right of index n - 1. 
// Additional room dividers can be installed. 
// For each position between indices i - 1 and i (1 <= i <= n - 1), at most one divider can be installed.

// Divide the corridor into non-overlapping sections, where each section has exactly two seats with any number of plants. 
// There may be multiple ways to perform the division. 
// Two ways are different if there is a position with a room divider installed in the first way but not in the second way.

// Return the number of ways to divide the corridor. 
// Since the answer may be very large, return it modulo 10^9 + 7. If there is no way, return 0.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/04/1.png" />
// Input: corridor = "SSPPSPS"
// Output: 3
// Explanation: There are 3 different ways to divide the corridor.
// The black bars in the above image indicate the two room dividers already installed.
// Note that in each of the ways, each section has exactly two seats.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/12/04/2.png" />
// Input: corridor = "PPSPSP"
// Output: 1
// Explanation: There is only 1 way to divide the corridor, by not installing any additional dividers.
// Installing any would create some section that does not have exactly two seats.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/12/12/3.png" />
// Input: corridor = "S"
// Output: 0
// Explanation: There is no way to divide the corridor because there will always be a section that does not have exactly two seats.

// Constraints:
//     n == corridor.length
//     1 <= n <= 10^5
//     corridor[i] is either 'S' or 'P'.

import "fmt"

func numberOfWays(corridor string) int {
    count, pos, res := 0, -1, 1 // Count of seats, Index of previous seat
    for i, v := range corridor {
        if v == 'S' {
            if pos > -1 && count % 2 == 0 { // If we have passed an even number of seats - it is an odd seat
                res = (res * (i - pos)) % 1_000_000_007
            }
            count++
            pos = i
        }
    }
    if count < 2 || count % 2 == 1 { return 0 } // if it is odd count of seats (or less than 2) - result is 0
    return res
}

func numberOfWays1(corridor string) int {
    res, count, last := 1, 0, 0
    for i, ch := range corridor {
        if ch == 'S' {
            count++
            if count >= 3 && count % 2 > 0 { // 对于第 3,5,7,... 个座位，可以在其到其左侧最近座位之间的任意空隙放置屏风
                res = res * (i - last) % 1_000_000_007
            }
            last = i // 记录上一个座位的位置
        }
    }
    if count == 0 || count%2 > 0 { // 座位个数不能为 0 或奇数
        return 0
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/04/1.png" />
    // Input: corridor = "SSPPSPS"
    // Output: 3
    // Explanation: There are 3 different ways to divide the corridor.
    // The black bars in the above image indicate the two room dividers already installed.
    // Note that in each of the ways, each section has exactly two seats.
    fmt.Println(numberOfWays("SSPPSPS")) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/12/04/2.png" />
    // Input: corridor = "PPSPSP"
    // Output: 1
    // Explanation: There is only 1 way to divide the corridor, by not installing any additional dividers.
    // Installing any would create some section that does not have exactly two seats.
    fmt.Println(numberOfWays("PPSPSP")) // 1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/12/12/3.png" />
    // Input: corridor = "S"
    // Output: 0
    // Explanation: There is no way to divide the corridor because there will always be a section that does not have exactly two seats.
    fmt.Println(numberOfWays("S")) // 0

    fmt.Println(numberOfWays("SSSSSSSSSS")) // 0
    fmt.Println(numberOfWays("PPPPPPPPPP")) // 0
    fmt.Println(numberOfWays("SSSSSPPPPP")) // 0
    fmt.Println(numberOfWays("PPPPPSSSSS")) // 0
    fmt.Println(numberOfWays("SPSPSPSPSP")) // 0
    fmt.Println(numberOfWays("PSPSPSPSPS")) // 0

    fmt.Println(numberOfWays1("SSPPSPS")) // 3
    fmt.Println(numberOfWays1("PPSPSP")) // 1
    fmt.Println(numberOfWays1("S")) // 0
    fmt.Println(numberOfWays1("SSSSSSSSSS")) // 0
    fmt.Println(numberOfWays1("PPPPPPPPPP")) // 0
    fmt.Println(numberOfWays1("SSSSSPPPPP")) // 0
    fmt.Println(numberOfWays1("PPPPPSSSSS")) // 0
    fmt.Println(numberOfWays1("SPSPSPSPSP")) // 0
    fmt.Println(numberOfWays1("PSPSPSPSPS")) // 0
}