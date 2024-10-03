package main

// 2189. Number of Ways to Build House of Cards
// You are given an integer n representing the number of playing cards you have. 
// A house of cards meets the following conditions:
//     1. A house of cards consists of one or more rows of triangles and horizontal cards.
//     2. Triangles are created by leaning two cards against each other.
//     3. One card must be placed horizontally between all adjacent triangles in a row.
//     4. Any triangle on a row higher than the first must be placed on a horizontal card from the previous row.
//     5. Each triangle is placed in the leftmost available spot in the row.

// Return the number of distinct house of cards you can build using all n cards. 
// Two houses of cards are considered distinct if there exists a row where the two houses contain a different number of cards.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/02/27/image-20220227213243-1.png" />
// Input: n = 16
// Output: 2
// Explanation: The two valid houses of cards are shown.
// The third house of cards in the diagram is not valid because the rightmost triangle on the top row is not placed on top of a horizontal card.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/02/27/image-20220227213306-2.png" />
// Input: n = 2
// Output: 1
// Explanation: The one valid house of cards is shown.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/02/27/image-20220227213331-3.png" />
// Input: n = 4
// Output: 0
// Explanation: The three houses of cards in the diagram are not valid.
// The first house of cards needs a horizontal card placed between the two triangles.
// The second house of cards uses 5 cards.
// The third house of cards uses 2 cards.

// Constraints:
//     1 <= n <= 500

import "fmt"

func houseOfCards(n int) int {
    dp := make([]int, n + 1)
    dp[0] = 1
    for i := 1; i <= (n + 1) / 3; i++ {
        curr := 3 * i - 1
        for j := n; j >= curr; j-- {
            dp[j] += dp[j - curr]
        }
    }
    return dp[n]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/02/27/image-20220227213243-1.png" />
    // Input: n = 16
    // Output: 2
    // Explanation: The two valid houses of cards are shown.
    // The third house of cards in the diagram is not valid because the rightmost triangle on the top row is not placed on top of a horizontal card.
    fmt.Println(houseOfCards(16)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/02/27/image-20220227213306-2.png" />
    // Input: n = 2
    // Output: 1
    // Explanation: The one valid house of cards is shown.
    fmt.Println(houseOfCards(2)) // 1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/02/27/image-20220227213331-3.png" />
    // Input: n = 4
    // Output: 0
    // Explanation: The three houses of cards in the diagram are not valid.
    // The first house of cards needs a horizontal card placed between the two triangles.
    // The second house of cards uses 5 cards.
    // The third house of cards uses 2 cards.
    fmt.Println(houseOfCards(4)) // 0
}