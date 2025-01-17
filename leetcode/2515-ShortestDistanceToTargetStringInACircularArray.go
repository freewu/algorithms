package main

// 2515. Shortest Distance to Target String in a Circular Array
// You are given a 0-indexed circular string array words and a string target. 
// A circular array means that the array's end connects to the array's beginning.
//     Formally, the next element of words[i] is words[(i + 1) % n] 
//     and the previous element of words[i] is words[(i - 1 + n) % n], where n is the length of words.

// Starting from startIndex, you can move to either the next word or the previous word with 1 step at a time.

// Return the shortest distance needed to reach the string target. 
// If the string target does not exist in words, return -1.

// Example 1:
// Input: words = ["hello","i","am","leetcode","hello"], target = "hello", startIndex = 1
// Output: 1
// Explanation: We start from index 1 and can reach "hello" by
// - moving 3 units to the right to reach index 4.
// - moving 2 units to the left to reach index 4.
// - moving 4 units to the right to reach index 0.
// - moving 1 unit to the left to reach index 0.
// The shortest distance to reach "hello" is 1.

// Example 2:
// Input: words = ["a","b","leetcode"], target = "leetcode", startIndex = 0
// Output: 1
// Explanation: We start from index 0 and can reach "leetcode" by
// - moving 2 units to the right to reach index 3.
// - moving 1 unit to the left to reach index 3.
// The shortest distance to reach "leetcode" is 1.

// Example 3:
// Input: words = ["i","eat","leetcode"], target = "ate", startIndex = 0
// Output: -1
// Explanation: Since "ate" does not exist in words, we return -1.

// Constraints:
//     1 <= words.length <= 100
//     1 <= words[i].length <= 100
//     words[i] and target consist of only lowercase English letters.
//     0 <= startIndex < words.length

import "fmt"

func closetTarget(words []string, target string, startIndex int) int {
    if words[startIndex] == target { return 0 } // base case
    n, distance, left, right := len(words), 1, startIndex - 1, startIndex + 1
    // check if the target exists while incrementing the pointers
    // exit if we completer a circle and the target isn't found
    for {
        if left < 0 { left = n - 1 }
        if right > (n - 1) { right = 0 }
        if words[left] == target || words[right] == target { return distance }
        if left == right { return -1 }
        left--
        right++
        distance++
    }
    return -1
}

func closetTarget1(words []string, target string, startIndex int) int {
    for i, n := 0, len(words); i <= n / 2; i++ {
        if words[(startIndex + i) % n] == target {
            return i
        }
        if words[(startIndex - i + n) % n] == target {
            return i
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: words = ["hello","i","am","leetcode","hello"], target = "hello", startIndex = 1
    // Output: 1
    // Explanation: We start from index 1 and can reach "hello" by
    // - moving 3 units to the right to reach index 4.
    // - moving 2 units to the left to reach index 4.
    // - moving 4 units to the right to reach index 0.
    // - moving 1 unit to the left to reach index 0.
    // The shortest distance to reach "hello" is 1.
    fmt.Println(closetTarget([]string{"hello","i","am","leetcode","hello"}, "hello", 1)) // 1
    // Example 2:
    // Input: words = ["a","b","leetcode"], target = "leetcode", startIndex = 0
    // Output: 1
    // Explanation: We start from index 0 and can reach "leetcode" by
    // - moving 2 units to the right to reach index 3.
    // - moving 1 unit to the left to reach index 3.
    // The shortest distance to reach "leetcode" is 1.
    fmt.Println(closetTarget([]string{"a","b","leetcode"}, "leetcode", 0)) // 1
    // Example 3:
    // Input: words = ["i","eat","leetcode"], target = "ate", startIndex = 0
    // Output: -1
    // Explanation: Since "ate" does not exist in words, we return -1.
    fmt.Println(closetTarget([]string{"i","eat","leetcode"}, "ate", 0)) // -1

    fmt.Println(closetTarget1([]string{"hello","i","am","leetcode","hello"}, "hello", 1)) // 1
    fmt.Println(closetTarget1([]string{"a","b","leetcode"}, "leetcode", 0)) // 1
    fmt.Println(closetTarget1([]string{"i","eat","leetcode"}, "ate", 0)) // -1
}