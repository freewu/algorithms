package main

// 1165. Single-Row Keyboard
// There is a special keyboard with all keys in a single row.

// Given a string keyboard of length 26 indicating the layout of the keyboard (indexed from 0 to 25).
// Initially, your finger is at index 0. 
// To type a character, you have to move your finger to the index of the desired character. 
// The time taken to move your finger from index i to index j is |i - j|.

// You want to type a string word. Write a function to calculate how much time it takes to type it with one finger.

// Example 1:
// Input: keyboard = "abcdefghijklmnopqrstuvwxyz", word = "cba"
// Output: 4
// Explanation: The index moves from 0 to 2 to write 'c' then to 1 to write 'b' then to 0 again to write 'a'.
// Total time = 2 + 1 + 1 = 4. 

// Example 2:
// Input: keyboard = "pqrstuvwxyzabcdefghijklmno", word = "leetcode"
// Output: 73

// Constraints:
//     keyboard.length == 26
//     keyboard contains each English lowercase letter exactly once in some order.
//     1 <= word.length <= 10^4
//     word[i] is an English lowercase letter.

import "fmt"

// map
func calculateTime(keyboard string, word string) int {
    pos, count, m := 0, 0, make(map[byte]int, len(keyboard)) // 使用一个map来记录字符和索引的距离
    for i := 0; i < len(keyboard); i++ {
        m[keyboard[i]] = i
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < len(word); i++ {
        count += abs(pos - m[word[i]]) // 移动的时间
        pos = m[word[i]] // 现在在的位置
    }
    return count
}

// arr 使用数组来替代 map
func calculateTime1(keyboard string, word string) int {
    pos, count, arr := 0, 0, make([]int, 26) // 使用一个map来记录字符和索引的距离
    for i := 0; i < len(keyboard); i++ {
        arr[keyboard[i] - 'a'] = i
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < len(word); i++ {
        curr := arr[word[i] - 'a']
        count += abs(pos - curr) // 移动的时间
        pos = curr // 现在在的位置
    }
    return count
}

func main() {
    // Example 1:
    // Input: keyboard = "abcdefghijklmnopqrstuvwxyz", word = "cba"
    // Output: 4
    // Explanation: The index moves from 0 to 2 to write 'c' then to 1 to write 'b' then to 0 again to write 'a'.
    // Total time = 2 + 1 + 1 = 4. 
    fmt.Println(calculateTime("abcdefghijklmnopqrstuvwxyz","cba")) // 4
    // Example 2:
    // Input: keyboard = "pqrstuvwxyzabcdefghijklmno", word = "leetcode"
    // Output: 73
    fmt.Println(calculateTime("pqrstuvwxyzabcdefghijklmno","leetcode")) // 73

    fmt.Println(calculateTime1("abcdefghijklmnopqrstuvwxyz","cba")) // 4
    fmt.Println(calculateTime1("pqrstuvwxyzabcdefghijklmno","leetcode")) // 73
}