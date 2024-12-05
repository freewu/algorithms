package main

// 2550. Count Collisions of Monkeys on a Polygon
// There is a regular convex polygon with n vertices. 
// The vertices are labeled from 0 to n - 1 in a clockwise direction, and each vertex has exactly one monkey. 
// The following figure shows a convex polygon of 6 vertices.
// <img src="https://assets.leetcode.com/uploads/2023/01/22/hexagon.jpg" />

// Simultaneously, each monkey moves to a neighboring vertex. 
// A collision happens if at least two monkeys reside on the same vertex after the movement or intersect on an edge.

// Return the number of ways the monkeys can move so that at least one collision happens. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 3
// Output: 6
// Explanation:
// There are 8 total possible movements.
// Two ways such that they collide at some point are:
//     Monkey 1 moves in a clockwise direction; monkey 2 moves in an anticlockwise direction; monkey 3 moves in a clockwise direction. Monkeys 1 and 2 collide.
//     Monkey 1 moves in an anticlockwise direction; monkey 2 moves in an anticlockwise direction; monkey 3 moves in a clockwise direction. Monkeys 1 and 3 collide.

// Example 2:
// Input: n = 4
// Output: 14

// Constraints:
//     3 <= n <= 10^9

import "fmt"

func monkeyMove(n int) int {
    mod := 1_000_000_007
    var pow func(k, n int) int
    pow = func(k, n int) int {
        if n == 0 { return 1 }
        if n == 1 { return k }
        half := pow(k, n / 2) % mod
        if n % 2 == 1 {
            return (half * half * k) % mod
        } else {
            return (half * half) % mod
        }
    }
    res := pow(2, n) - 2
    if res < 0 {
        return res + mod
    }
    return res
}

func monkeyMove1(n int) int {
    res, a, mod := 1, 2, 1_000_000_007
    for n > 0 {
        if (n & 1) == 1 {
            res = (res * a) % mod
        }
        a = (a * a) % mod
        n >>= 1
    }
    return (res - 2 + mod) % mod
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: 6
    // Explanation:
    // There are 8 total possible movements.
    // Two ways such that they collide at some point are:
    //     Monkey 1 moves in a clockwise direction; monkey 2 moves in an anticlockwise direction; monkey 3 moves in a clockwise direction. Monkeys 1 and 2 collide.
    //     Monkey 1 moves in an anticlockwise direction; monkey 2 moves in an anticlockwise direction; monkey 3 moves in a clockwise direction. Monkeys 1 and 3 collide.
    fmt.Println(monkeyMove(3)) // 6
    // Example 2:
    // Input: n = 4
    // Output: 14
    fmt.Println(monkeyMove(4)) // 14

    fmt.Println(monkeyMove(1024)) // 812734590
    fmt.Println(monkeyMove(999_999_999)) // 570312502
    fmt.Println(monkeyMove(1_000_000_000)) // 140624999

    fmt.Println(monkeyMove1(3)) // 6
    fmt.Println(monkeyMove1(4)) // 14
    fmt.Println(monkeyMove1(1024)) // 812734590
    fmt.Println(monkeyMove1(999_999_999)) // 570312502
    fmt.Println(monkeyMove1(1_000_000_000)) // 140624999
}