package main

// 1974. Minimum Time to Type Word Using Special Typewriter
// There is a special typewriter with lowercase English letters 'a' to 'z' arranged in a circle with a pointer. 
// A character can only be typed if the pointer is pointing to that character. 
// The pointer is initially pointing to the character 'a'.
// <img src="https://assets.leetcode.com/uploads/2021/07/31/chart.jpg" />

// Each second, you may perform one of the following operations:
//     Move the pointer one character counterclockwise or clockwise.
//     Type the character the pointer is currently on.

// Given a string word, return the minimum number of seconds to type out the characters in word.

// Example 1:
// Input: word = "abc"
// Output: 5
// Explanation: 
// The characters are printed as follows:
// - Type the character 'a' in 1 second since the pointer is initially on 'a'.
// - Move the pointer clockwise to 'b' in 1 second.
// - Type the character 'b' in 1 second.
// - Move the pointer clockwise to 'c' in 1 second.
// - Type the character 'c' in 1 second.

// Example 2:
// Input: word = "bza"
// Output: 7
// Explanation:
// The characters are printed as follows:
// - Move the pointer clockwise to 'b' in 1 second.
// - Type the character 'b' in 1 second.
// - Move the pointer counterclockwise to 'z' in 2 seconds.
// - Type the character 'z' in 1 second.
// - Move the pointer clockwise to 'a' in 1 second.
// - Type the character 'a' in 1 second.

// Example 3:
// Input: word = "zjpc"
// Output: 34
// Explanation:
// The characters are printed as follows:
// - Move the pointer counterclockwise to 'z' in 1 second.
// - Type the character 'z' in 1 second.
// - Move the pointer clockwise to 'j' in 10 seconds.
// - Type the character 'j' in 1 second.
// - Move the pointer clockwise to 'p' in 6 seconds.
// - Type the character 'p' in 1 second.
// - Move the pointer counterclockwise to 'c' in 13 seconds.
// - Type the character 'c' in 1 second.

// Constraints:
//     1 <= word.length <= 100
//     word consists of lowercase English letters.

import "fmt"

func minTimeToType(word string) int {
    res, prev := 0, 'a'
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range word {
        diff := abs(int(v) - int(prev))
        res += min(diff, (26 - diff)) + 1
        prev = v
    }
    return res
}

func minTimeToType1(word string) int {
    res, cur := 0, 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range word {
        c := int(v - 'a')
        res += min(abs(c - cur), abs(abs(c - cur) - 26)) + 1
        cur = c
    }
    return res
}

func main() {
    // Example 1:
    // Input: word = "abc"
    // Output: 5
    // Explanation: 
    // The characters are printed as follows:
    // - Type the character 'a' in 1 second since the pointer is initially on 'a'.
    // - Move the pointer clockwise to 'b' in 1 second.
    // - Type the character 'b' in 1 second.
    // - Move the pointer clockwise to 'c' in 1 second.
    // - Type the character 'c' in 1 second.
    fmt.Println(minTimeToType("abc")) // 5
    // Example 2:
    // Input: word = "bza"
    // Output: 7
    // Explanation:
    // The characters are printed as follows:
    // - Move the pointer clockwise to 'b' in 1 second.
    // - Type the character 'b' in 1 second.
    // - Move the pointer counterclockwise to 'z' in 2 seconds.
    // - Type the character 'z' in 1 second.
    // - Move the pointer clockwise to 'a' in 1 second.
    // - Type the character 'a' in 1 second.
    fmt.Println(minTimeToType("bza")) // 7
    // Example 3:
    // Input: word = "zjpc"
    // Output: 34
    // Explanation:
    // The characters are printed as follows:
    // - Move the pointer counterclockwise to 'z' in 1 second.
    // - Type the character 'z' in 1 second.
    // - Move the pointer clockwise to 'j' in 10 seconds.
    // - Type the character 'j' in 1 second.
    // - Move the pointer clockwise to 'p' in 6 seconds.
    // - Type the character 'p' in 1 second.
    // - Move the pointer counterclockwise to 'c' in 13 seconds.
    // - Type the character 'c' in 1 second.
    fmt.Println(minTimeToType("zjpc")) // 34

    fmt.Println(minTimeToType1("abc")) // 5
    fmt.Println(minTimeToType1("bza")) // 7
    fmt.Println(minTimeToType1("zjpc")) // 34
}