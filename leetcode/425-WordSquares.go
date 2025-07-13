package main

// 425. Word Squares
// Given an array of unique strings words, return all the word squares you can build from words. 
// The same word from words can be used multiple times. You can return the answer in any order.

// A sequence of strings forms a valid word square if the kth row and column read the same string, 
// where 0 <= k < max(numRows, numColumns).

// For example, the word sequence ["ball","area","lead","lady"] forms a word square 
// because each word reads the same both horizontally and vertically.
 
// Example 1:
// Input: words = ["area","lead","wall","lady","ball"]
// Output: [["ball","area","lead","lady"],["wall","area","lead","lady"]]
// Explanation:
// The output consists of two word squares. The order of output does not matter (just the order of words in each word square matters).

// Example 2:
// Input: words = ["abat","baba","atan","atal"]
// Output: [["baba","abat","baba","atal"],["baba","abat","baba","atan"]]
// Explanation:
// The output consists of two word squares. The order of output does not matter (just the order of words in each word square matters).
 
// Constraints:
//     1 <= words.length <= 1000
//     1 <= words[i].length <= 4
//     All words[i] have the same length.
//     words[i] consists of only lowercase English letters.
//     All words[i] are unique.

import "fmt"

func wordSquares(words []string) [][]string {
    res, board, mp := [][]string{}, make([]string, len(words[0])), map[string][]string{} // 行前缀哈希表
    if len(words) == 0 || len(words[0]) == 0 {
        return res
    }
    for _, i := range words { // 添加行前缀对应的字符串
        for j := 0; j < len(i); j++ {
            mp[i[:j]] = append(mp[i[:j]], i)
        }
    }
    var backtrack func(v int) 
    backtrack = func(v int) {
        if v == len(words[0]) { // 回溯结束条件
            res = append(res, append([]string(nil),board...))
            return
        }
        if _, ok := mp[board[v]]; ok { // 如果现在的行前缀没有对应的字符串，则不用继续回溯下去
            for _, i := range mp[board[v]] {
                board[v] = i
                for j := v + 1; j < len(words[0]); j++ {
                    board[j] = board[j] + string(board[v][j])
                }
                backtrack(v + 1)
                for j := v + 1; j < len(words[0]); j++ {
                    board[j] = board[j][:len(board[j])-1]
                }
                board[v] = board[v][:v]
            }
        }
    }
    backtrack(0)
    return res
}

func main() {
    // Example 1:
    // Input: words = ["area","lead","wall","lady","ball"]
    // Output: [["ball","area","lead","lady"],["wall","area","lead","lady"]]
    // Explanation:
    // The output consists of two word squares. The order of output does not matter (just the order of words in each word square matters).
    fmt.Println(wordSquares([]string{"area","lead","wall","lady","ball"})) // [["ball","area","lead","lady"],["wall","area","lead","lady"]]
    // Example 2:
    // Input: words = ["abat","baba","atan","atal"]
    // Output: [["baba","abat","baba","atal"],["baba","abat","baba","atan"]]
    // Explanation:
    // The output consists of two word squares. The order of output does not matter (just the order of words in each word square matters).
    fmt.Println(wordSquares([]string{"abat","baba","atan","atal"})) // [["baba","abat","baba","atal"],["baba","abat","baba","atan"]]

    fmt.Println(wordSquares([]string{"blue","frog","leet","code","free"})) // []
}