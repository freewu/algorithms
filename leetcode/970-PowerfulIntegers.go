package main

// 970. Powerful Integers
// Given three integers x, y, and bound, return a list of all the powerful integers that have a value less than or equal to bound.

// An integer is powerful if it can be represented as xi + yj for some integers i >= 0 and j >= 0.

// You may return the answer in any order. In your answer, each value should occur at most once.

// Example 1:
// Input: x = 2, y = 3, bound = 10
// Output: [2,3,4,5,7,9,10]
// Explanation:
// 2 = 20 + 30
// 3 = 21 + 30
// 4 = 20 + 31
// 5 = 21 + 31
// 7 = 22 + 31
// 9 = 23 + 30
// 10 = 20 + 32

// Example 2:
// Input: x = 3, y = 5, bound = 15
// Output: [2,4,6,8,10,14]

// Constraints:
//     1 <= x, y <= 100
//     0 <= bound <= 10^6

import "fmt"

func powerfulIntegers(x int, y int, bound int) []int {
    mp := make(map[int]bool)
    for i := 1; i < bound; i *= x {
        for j := 1; i+j <= bound; j *= y {
            mp[i+j] = true
            if y == 1 { break }
        }
        if x == 1 { break }
    }
    res := []int{}
    for v, _ := range mp {
        if mp[v] {
            res = append(res, v)
        }
    }
    return res
}

func powerfulIntegers1(x int, y int, bound int) []int {
    res, mp := []int{}, make(map[int]bool)
    c, d := 1, 1
    for a := 0; a < 21; a++ {
        if c > bound || c <= 0 { break }
        d = 1
        for b := 0; b < 21; b++ {
            if c + d > bound || c + d <= 0 { break }
            mp[c+d] = true
            d = d * y
        }
        c = c * x
    }
    for k := range mp {
        res = append(res, k)
    }
    return res
}

func main() {
    // Example 1:
    // Input: x = 2, y = 3, bound = 10
    // Output: [2,3,4,5,7,9,10]
    // Explanation:
    // 2 = 20 + 30
    // 3 = 21 + 30
    // 4 = 20 + 31
    // 5 = 21 + 31
    // 7 = 22 + 31
    // 9 = 23 + 30
    // 10 = 20 + 32
    fmt.Println(powerfulIntegers(2, 3, 10)) // [2,3,4,5,7,9,10]
    // Example 2:
    // Input: x = 3, y = 5, bound = 15
    // Output: [2,4,6,8,10,14]
    fmt.Println(powerfulIntegers(3, 5, 15)) // [2,4,6,8,10,14]

    fmt.Println(powerfulIntegers1(2, 3, 10)) // [2,3,4,5,7,9,10]
    fmt.Println(powerfulIntegers1(3, 5, 15)) // [2,4,6,8,10,14]
}