package main

// 2103. Rings and Rods
// There are n rings and each ring is either red, green, or blue. 
// The rings are distributed across ten rods labeled from 0 to 9.

// You are given a string rings of length 2n that describes the n rings that are placed onto the rods. 
// Every two characters in rings forms a color-position pair that is used to describe each ring where:
//     The first character of the ith pair denotes the ith ring's color ('R', 'G', 'B').
//     The second character of the ith pair denotes the rod that the ith ring is placed on ('0' to '9').

// For example, "R3G2B1" describes n == 3 rings: 
//     a red ring placed onto the rod labeled 3, 
//     a green ring placed onto the rod labeled 2, 
//     and a blue ring placed onto the rod labeled 1.

// Return the number of rods that have all three colors of rings on them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/23/ex1final.png" />
// Input: rings = "B0B6G0R6R0R6G9"
// Output: 1
// Explanation: 
// - The rod labeled 0 holds 3 rings with all colors: red, green, and blue.
// - The rod labeled 6 holds 3 rings, but it only has red and blue.
// - The rod labeled 9 holds only a green ring.
// Thus, the number of rods with all three colors is 1.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/11/23/ex2final.png" />
// Input: rings = "B0R0G0R9R0B0G0"
// Output: 1
// Explanation: 
// - The rod labeled 0 holds 6 rings with all colors: red, green, and blue.
// - The rod labeled 9 holds only a red ring.
// Thus, the number of rods with all three colors is 1.

// Example 3:
// Input: rings = "G4"
// Output: 0
// Explanation: 
// Only one ring is given. Thus, no rods have all three colors.

// Constraints:
//     rings.length == 2 * n
//     1 <= n <= 100
//     rings[i] where i is even is either 'R', 'G', or 'B' (0-indexed).
//     rings[i] where i is odd is a digit from '0' to '9' (0-indexed).

import "fmt"

func countPoints(rings string) int {
    type rod struct {
        red, green, blue bool
    }
    res, rods := 0, []rod{}
    for i := 0; i < 10; i ++ { // The rings are distributed across ten rods labeled from 0 to 9.
        rods = append(rods, rod{ red: false, green: false, blue: false, })
    }
    for i := 0; i < len(rings); i += 2 {
        color := rings[i]
        rodNo := rings[i + 1] - '0'
        if color == 'R' { rods[rodNo].red = true }
        if color == 'G' { rods[rodNo].green = true }
        if color == 'B' { rods[rodNo].blue = true }
    }
    for _, v := range rods {
        if v.red && v.green && v.blue { res++ } // 每根杆都有三种颜色的环
    }
    return res
}

func countPoints1(rings string) int {
    res, mp := 0, make(map[byte]int)
    for i := 0; i < len(rings); i += 2 {
        if mp[rings[i+1]] == 0b111 { continue } // 本杆已满足三色
        switch rings[i] {
        case 'R': mp[rings[i+1]] |= 0b100
        case 'G': mp[rings[i+1]] |= 0b010
        case 'B': mp[rings[i+1]] |= 0b001
        }
        if mp[rings[i+1]] == 0b111 { res++ }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/11/23/ex1final.png" />
    // Input: rings = "B0B6G0R6R0R6G9"
    // Output: 1
    // Explanation: 
    // - The rod labeled 0 holds 3 rings with all colors: red, green, and blue.
    // - The rod labeled 6 holds 3 rings, but it only has red and blue.
    // - The rod labeled 9 holds only a green ring.
    // Thus, the number of rods with all three colors is 1.
    fmt.Println(countPoints("B0B6G0R6R0R6G9")) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/11/23/ex2final.png" />
    // Input: rings = "B0R0G0R9R0B0G0"
    // Output: 1
    // Explanation: 
    // - The rod labeled 0 holds 6 rings with all colors: red, green, and blue.
    // - The rod labeled 9 holds only a red ring.
    // Thus, the number of rods with all three colors is 1.
    fmt.Println(countPoints("B0R0G0R9R0B0G0")) // 1
    // Example 3:
    // Input: rings = "G4"
    // Output: 0
    // Explanation: 
    // Only one ring is given. Thus, no rods have all three colors.
    fmt.Println(countPoints("G4")) // 0

    fmt.Println(countPoints1("B0B6G0R6R0R6G9")) // 1
    fmt.Println(countPoints1("B0R0G0R9R0B0G0")) // 1
    fmt.Println(countPoints1("G4")) // 0
}