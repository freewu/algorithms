package main

// 1753. Maximum Score From Removing Stones
// You are playing a solitaire game with three piles of stones of sizes a​​​​​​, b,​​​​​​ and c​​​​​​ respectively. 
// Each turn you choose two different non-empty piles, take one stone from each, and add 1 point to your score. 
// The game stops when there are fewer than two non-empty piles (meaning there are no more available moves).

// Given three integers a​​​​​, b,​​​​​ and c​​​​​, return the maximum score you can get.

// Example 1:
// Input: a = 2, b = 4, c = 6
// Output: 6
// Explanation: The starting state is (2, 4, 6). One optimal set of moves is:
// - Take from 1st and 3rd piles, state is now (1, 4, 5)
// - Take from 1st and 3rd piles, state is now (0, 4, 4)
// - Take from 2nd and 3rd piles, state is now (0, 3, 3)
// - Take from 2nd and 3rd piles, state is now (0, 2, 2)
// - Take from 2nd and 3rd piles, state is now (0, 1, 1)
// - Take from 2nd and 3rd piles, state is now (0, 0, 0)
// There are fewer than two non-empty piles, so the game ends. Total: 6 points.

// Example 2:
// Input: a = 4, b = 4, c = 6
// Output: 7
// Explanation: The starting state is (4, 4, 6). One optimal set of moves is:
// - Take from 1st and 2nd piles, state is now (3, 3, 6)
// - Take from 1st and 3rd piles, state is now (2, 3, 5)
// - Take from 1st and 3rd piles, state is now (1, 3, 4)
// - Take from 1st and 3rd piles, state is now (0, 3, 3)
// - Take from 2nd and 3rd piles, state is now (0, 2, 2)
// - Take from 2nd and 3rd piles, state is now (0, 1, 1)
// - Take from 2nd and 3rd piles, state is now (0, 0, 0)
// There are fewer than two non-empty piles, so the game ends. Total: 7 points.

// Example 3:
// Input: a = 1, b = 8, c = 8
// Output: 8
// Explanation: One optimal set of moves is to take from the 2nd and 3rd piles for 8 turns until they are empty.
// After that, there are fewer than two non-empty piles, so the game ends.

// Constraints:
//     1 <= a, b, c <= 10^5

import "fmt"

func maximumScore(a int, b int, c int) int {
    res := 0
    getZeroCount := func(nums ...int) int {
        res := 0
        for _, v := range nums {
            if v == 0 { res++ }
        }
        return res
    } 
    getMax := func(nums ...int) int {
        res := nums[0]
        for _, v := range nums {
            if v > res { res = v }
        }
        return res
    }
    getSum := func(nums ...int) int {
        res := 0
        for _, v := range nums {
            res += v
        }
        return res
    }
    for {
        if getZeroCount(a, b, c) > 1 { break }
        res++
        sum, mx := getSum(a, b, c), getMax(a, b, c)
        if sum > 2 * mx {
            if a == mx {
                b--
                c--
            } else if b == mx {
                a--
                c--
            } else {
                a--
                b--
            }
        } else {
            if a == mx {
                a--
                if b > 0 {
                    b--
                } else {
                    c--
                }
            } else if b == mx {
                b--
                if a > 0 {
                    a--
                } else {
                    c--
                }
            } else {
                c--
                if a > 0 {
                    a--
                } else {
                    b--
                }
            }
        }
    }
    return res
}

func maximumScore1(a, b, c int) int {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    sum, mx := a + b + c, max(max(a, b), c)
    if sum < mx * 2 {
        return sum - mx
    }
    return sum / 2
}

func main() {
    // Example 1:
    // Input: a = 2, b = 4, c = 6
    // Output: 6
    // Explanation: The starting state is (2, 4, 6). One optimal set of moves is:
    // - Take from 1st and 3rd piles, state is now (1, 4, 5)
    // - Take from 1st and 3rd piles, state is now (0, 4, 4)
    // - Take from 2nd and 3rd piles, state is now (0, 3, 3)
    // - Take from 2nd and 3rd piles, state is now (0, 2, 2)
    // - Take from 2nd and 3rd piles, state is now (0, 1, 1)
    // - Take from 2nd and 3rd piles, state is now (0, 0, 0)
    // There are fewer than two non-empty piles, so the game ends. Total: 6 points.
    fmt.Println(maximumScore(2, 4, 6)) // 6
    // Example 2:
    // Input: a = 4, b = 4, c = 6
    // Output: 7
    // Explanation: The starting state is (4, 4, 6). One optimal set of moves is:
    // - Take from 1st and 2nd piles, state is now (3, 3, 6)
    // - Take from 1st and 3rd piles, state is now (2, 3, 5)
    // - Take from 1st and 3rd piles, state is now (1, 3, 4)
    // - Take from 1st and 3rd piles, state is now (0, 3, 3)
    // - Take from 2nd and 3rd piles, state is now (0, 2, 2)
    // - Take from 2nd and 3rd piles, state is now (0, 1, 1)
    // - Take from 2nd and 3rd piles, state is now (0, 0, 0)
    // There are fewer than two non-empty piles, so the game ends. Total: 7 points.
    fmt.Println(maximumScore(4, 4, 6)) // 7
    // Example 3:
    // Input: a = 1, b = 8, c = 8
    // Output: 8
    // Explanation: One optimal set of moves is to take from the 2nd and 3rd piles for 8 turns until they are empty.
    // After that, there are fewer than two non-empty piles, so the game ends.
    fmt.Println(maximumScore(1, 8, 8)) // 8

    fmt.Println(maximumScore(1, 1, 1)) // 1
    fmt.Println(maximumScore(10000, 10000, 10000)) // 15000
    fmt.Println(maximumScore(1, 1, 10000)) // 2
    fmt.Println(maximumScore(1, 10000, 10000)) // 10000
    fmt.Println(maximumScore(1, 10000, 1)) // 2
    fmt.Println(maximumScore(10000, 10000, 1)) // 10000
    fmt.Println(maximumScore(10000, 1, 1)) // 2
    fmt.Println(maximumScore(10000, 1, 10000)) // 10000

    fmt.Println(maximumScore1(2, 4, 6)) // 6
    fmt.Println(maximumScore1(4, 4, 6)) // 7
    fmt.Println(maximumScore1(1, 8, 8)) // 8
    fmt.Println(maximumScore1(1, 1, 1)) // 1
    fmt.Println(maximumScore1(10000, 10000, 10000)) // 15000
    fmt.Println(maximumScore1(1, 1, 10000)) // 2
    fmt.Println(maximumScore1(1, 10000, 10000)) // 10000
    fmt.Println(maximumScore1(1, 10000, 1)) // 2
    fmt.Println(maximumScore1(10000, 10000, 1)) // 10000
    fmt.Println(maximumScore1(10000, 1, 1)) // 2
    fmt.Println(maximumScore1(10000, 1, 10000)) // 10000
}