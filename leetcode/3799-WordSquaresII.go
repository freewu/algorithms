package main

// 3799. Word Squares II
// You are given a string array words, consisting of distinct 4-letter strings, each containing lowercase English letters.

// A word square consists of 4 distinct words: top, left, right and bottom, arranged as follows:
//     1. top forms the top row.
//     2. bottom forms the bottom row.
//     3. left forms the left column (top to bottom).
//     4. right forms the right column (top to bottom).

// It must satisfy:
//     1. top[0] == left[0], top[3] == right[0]
//     2. bottom[0] == left[3], bottom[3] == right[3]

// Return all valid distinct word squares, sorted in ascending lexicographic order by the 4-tuple (top, left, right, bottom)​​​​​​​.

// Example 1:
// Input: words = ["able","area","echo","also"]
// Output: [["able","area","echo","also"],["area","able","also","echo"]]
// Explanation:
// There are exactly two valid 4-word squares that satisfy all corner constraints:
// "able" (top), "area" (left), "echo" (right), "also" (bottom)
// top[0] == left[0] == 'a'
// top[3] == right[0] == 'e'
// bottom[0] == left[3] == 'a'
// bottom[3] == right[3] == 'o'
// "area" (top), "able" (left), "also" (right), "echo" (bottom)
// All corner constraints are satisfied.
// Thus, the answer is [["able","area","echo","also"],["area","able","also","echo"]].

// Example 2:
// Input: words = ["code","cafe","eden","edge"]
// Output: []
// Explanation:
// No combination of four words satisfies all four corner constraints. Thus, the answer is empty array [].

// Constraints:
//     4 <= words.length <= 15
//     words[i].length == 4
//     words[i] consists of only lowercase English letters.
//     All words[i] are distinct.

import "fmt"
import "slices"

// 排列型回溯
func wordSquares(words []string) [][]string {
    slices.Sort(words) // 保证答案有序
    path, onPath := make([]int, 4), make([]bool, len(words))
    var res [][]string
    var dfs func(i int)
    dfs = func(i int) {
        if i == 4 {
            top, left, right, bottom  := words[path[0]], words[path[1]], words[path[2]], words[path[3]]
            if top[0] == left[0] && top[3] == right[0] && bottom[0] == left[3] && bottom[3] == right[3] {
                res = append(res, []string{top, left, right, bottom})
            }
            return
        }
        for j, on := range onPath {
            if !on {
                path[i] = j      // 从没有选的下标中选一个
                onPath[j] = true // 已选上
                dfs(i + 1)
                onPath[j] = false // 恢复现场
            }
        }
    }
    dfs(0)
    return res
}

func wordSquares1(words []string) [][]string {
    slices.Sort(words) // 保证答案有序
    res := [][]string{}
    var findSquare func([]string)
    findSquare = func(wordCheck []string) {
        if len(wordCheck) == 4 {
            result := make([]string, 4)
            copy(result, wordCheck)
            res = append(res, result)
            return
        }
        mp := make(map[string]bool)
        for _, word := range wordCheck {
            mp[word] = true
        }
        for _, word := range words {
            if mp[word] { continue }
            i := len(wordCheck) + 1
            if i == 2 {
                top := wordCheck[0]
                if word[0] == top[0] {
                    findSquare(append(wordCheck, word))
                }
            }
            if i == 3 {
                top := wordCheck[0]
                if word[0] == top[3] {
                    findSquare(append(wordCheck, word))
                }
            }
            if i == 4 {
                left := wordCheck[1]
                right := wordCheck[2]
                if word[3] == right[3] && word[0] == left[3] {
                    findSquare(append(wordCheck, word))
                }
            }
        }
    }
    for _, word := range words {
        wordCheck := []string{word}
        findSquare(wordCheck)
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["able","area","echo","also"]
    // Output: [["able","area","echo","also"],["area","able","also","echo"]]
    // Explanation:
    // There are exactly two valid 4-word squares that satisfy all corner constraints:
    // "able" (top), "area" (left), "echo" (right), "also" (bottom)
    // top[0] == left[0] == 'a'
    // top[3] == right[0] == 'e'
    // bottom[0] == left[3] == 'a'
    // bottom[3] == right[3] == 'o'
    // "area" (top), "able" (left), "also" (right), "echo" (bottom)
    // All corner constraints are satisfied.
    // Thus, the answer is [["able","area","echo","also"],["area","able","also","echo"]].
    fmt.Println(wordSquares([]string{"able","area","echo","also"})) // [["able","area","echo","also"],["area","able","also","echo"]]    
    // Example 2:
    // Input: words = ["code","cafe","eden","edge"]
    // Output: []
    // Explanation:
    // No combination of four words satisfies all four corner constraints. Thus, the answer is empty array [].
    fmt.Println(wordSquares([]string{"code","cafe","eden","edge"})) // []

    fmt.Println(wordSquares([]string{"blue","frog","leet","code"})) // []

    fmt.Println(wordSquares1([]string{"able","area","echo","also"})) // [["able","area","echo","also"],["area","able","also","echo"]]    
    fmt.Println(wordSquares1([]string{"code","cafe","eden","edge"})) // []
    fmt.Println(wordSquares1([]string{"blue","frog","leet","code"})) // []
}