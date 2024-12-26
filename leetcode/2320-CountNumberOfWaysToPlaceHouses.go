package main

// 2320. Count Number of Ways to Place Houses
// There is a street with n * 2 plots, where there are n plots on each side of the street. 
// The plots on each side are numbered from 1 to n. On each plot, a house can be placed.

// Return the number of ways houses can be placed such that no two houses are adjacent to each other on the same side of the street. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Note that if a house is placed on the ith plot on one side of the street, a house can also be placed on the ith plot on the other side of the street.

// Example 1:
// Input: n = 1
// Output: 4
// Explanation: 
// Possible arrangements:
// 1. All plots are empty.
// 2. A house is placed on one side of the street.
// 3. A house is placed on the other side of the street.
// 4. Two houses are placed, one on each side of the street.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/05/12/arrangements.png" />
// Input: n = 2
// Output: 9
// Explanation: The 9 possible arrangements are shown in the diagram above.

// Constraints:
//     1 <= n <= 10^4

import "fmt"

func countHousePlacements(n int) int {
    const mod = 1_000_000_007
    var fib func(n int) int
    fib = func(n int) int {
        a, b := 0, 1
        for i := 0; i < n; i++ {
            a, b = b, (a + b) % mod
        }
        return a % mod
    }
    onSide := fib(n + 2)
    return (onSide * onSide) % mod
}

func countHousePlacements1(n int) int {
    const mod = 1_000_000_007
    d0, d1 := 1, 2
    for i := 1; i < n; i++ {
        d0, d1 = d1, (d0 + d1) % mod
    }
    return (d1 * d1) % mod
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: 4
    // Explanation: 
    // Possible arrangements:
    // 1. All plots are empty.
    // 2. A house is placed on one side of the street.
    // 3. A house is placed on the other side of the street.
    // 4. Two houses are placed, one on each side of the street.
    fmt.Println(countHousePlacements(1)) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/05/12/arrangements.png" />
    // Input: n = 2
    // Output: 9
    // Explanation: The 9 possible arrangements are shown in the diagram above.
    fmt.Println(countHousePlacements(2)) // 9

    fmt.Println(countHousePlacements(64)) // 6323684
    fmt.Println(countHousePlacements(999)) // 194229620
    fmt.Println(countHousePlacements(1024)) // 635550672
    fmt.Println(countHousePlacements(9999)) // 459963766
    fmt.Println(countHousePlacements(10000)) // 402613600

    fmt.Println(countHousePlacements1(1)) // 4
    fmt.Println(countHousePlacements1(2)) // 9
    fmt.Println(countHousePlacements1(64)) // 6323684
    fmt.Println(countHousePlacements1(999)) // 194229620
    fmt.Println(countHousePlacements1(1024)) // 635550672
    fmt.Println(countHousePlacements1(9999)) // 459963766
    fmt.Println(countHousePlacements1(10000)) // 402613600
}