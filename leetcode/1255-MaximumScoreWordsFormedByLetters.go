package main

// 1255. Maximum Score Words Formed by Letters
// Given a list of words, list of  single letters (might be repeating) and score of every character.
// Return the maximum score of any valid set of words formed by using the given letters (words[i] cannot be used two or more times).
// It is not necessary to use all characters in letters and each letter can only be used once. 
// Score of letters 'a', 'b', 'c', ... ,'z' is given by score[0], score[1], ... , score[25] respectively.

// Example 1:
// Input: words = ["dog","cat","dad","good"], letters = ["a","a","c","d","d","d","g","o","o"], score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
// Output: 23
// Explanation:
// Score  a=1, c=9, d=5, g=3, o=2
// Given letters, we can form the words "dad" (5+1+5) and "good" (3+2+2+5) with a score of 23.
// Words "dad" and "dog" only get a score of 21.

// Example 2:
// Input: words = ["xxxz","ax","bx","cx"], letters = ["z","a","b","c","x","x","x"], score = [4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10]
// Output: 27
// Explanation:
// Score  a=4, b=4, c=4, x=5, z=10
// Given letters, we can form the words "ax" (4+5), "bx" (4+5) and "cx" (4+5) with a score of 27.
// Word "xxxz" only get a score of 25.

// Example 3:
// Input: words = ["leetcode"], letters = ["l","e","t","c","o","d"], score = [0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0]
// Output: 0
// Explanation:
// Letter "e" can only be used once.
 
// Constraints:
//     1 <= words.length <= 14
//     1 <= words[i].length <= 15
//     1 <= letters.length <= 100
//     letters[i].length == 1
//     score.length == 26
//     0 <= score[i] <= 10
//     words[i], letters[i] contains only lower case English letters.

import "fmt"
import "sort"

func maxScoreWords(words []string, letters []byte, score []int) int {
    type Word struct {
        word  string
        score int
    }
    getScore := func (scores []int, word string) int {
        score := 0
        for _, r := range word {
            score += scores[r - 'a']
        }
        return score
    }
    res, weight, usedLetter := 0, make([]*Word, len(words)), make([]bool, len(letters))
    for i, word := range words { // 把每个单词的得分缓存起来
        weight[i] = &Word{word, getScore(score, word)}
    }
    sort.Slice(letters, func(i, j int) bool { // letters排序
        return letters[i] < letters[j]
    })
    mark := func (letters []byte, usedLetter []bool, bytes []byte) {
        for i, j := 0, 0; i < len(bytes); j++ {
            if !usedLetter[j] && bytes[i] == letters[j] {
                usedLetter[j] = true
                i++
            }
        }
    }
    unmark := func (letters []byte, usedLetter []bool, bytes []byte) {
        for i, j := 0, 0; i < len(bytes); j++ {
            if usedLetter[j] && bytes[i] == letters[j] {
                usedLetter[j] = false
                i++
            }
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(int, int)
    dfs = func(index, score int) {
        if index == len(weight) {
            return
        }
        for i := index; i < len(weight); i++ {
            bytes := []byte(weight[i].word)
            sort.Slice(bytes, func(i, j int) bool { // bytes也排序，可以在线性时间内在 letters 中找到全部的 bytes 字符
                return bytes[i] < bytes[j]
            })
            j := 0
            for k := 0; j < len(bytes) && k < len(letters); k++ {
                if usedLetter[k] {
                    continue
                }
                p, q := bytes[j], letters[k]
                if p == q {
                    j++
                } else if p < q {
                    break
                }
            }
            if j == len(bytes) { // 可以构造出当前的单词
                mark(letters, usedLetter, bytes)
                res = max(res, score + weight[i].score)
                dfs(i+1, score + weight[i].score)
                unmark(letters, usedLetter, bytes)
            }
        }
    }
    dfs(0, 0)
    return res
}

func maxScoreWords1(words []string, letters []byte, score []int) int {
    res, left := 0, [26]int{}
    for _, c := range letters {
        left[c - 'a']++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(int, int)
    dfs = func(i, total int) {
        if i < 0 { 
            res = max(res, total)
            return
        }
        dfs(i-1, total) // 不选
        for j, c := range words[i] { // 选
            c -= 'a'
            if left[c] == 0 { // 剩余字母不足
                for _, c := range words[i][:j] { // 撤销
                    left[c-'a']++
                }
                return
            }
            left[c]-- // 减少剩余字母
            total += score[c] // 累加得分
        }
        dfs(i-1, total)
        for _, c := range words[i] {// 恢复现场
            left[c-'a']++
        }
    }
    dfs(len(words)-1, 0)
    return res
}

func main() {
    // Example 1:
    // Input: words = ["dog","cat","dad","good"], letters = ["a","a","c","d","d","d","g","o","o"], score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
    // Output: 23
    // Explanation:
    // Score  a=1, c=9, d=5, g=3, o=2
    // Given letters, we can form the words "dad" (5+1+5) and "good" (3+2+2+5) with a score of 23.
    // Words "dad" and "dog" only get a score of 21.
    fmt.Println(maxScoreWords([]string{"dog","cat","dad","good"},[]byte{'a','a','c','d','d','d','g','o','o'},[]int{1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0})) // 23
    // Example 2:
    // Input: words = ["xxxz","ax","bx","cx"], letters = ["z","a","b","c","x","x","x"], score = [4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10]
    // Output: 27
    // Explanation:
    // Score  a=4, b=4, c=4, x=5, z=10
    // Given letters, we can form the words "ax" (4+5), "bx" (4+5) and "cx" (4+5) with a score of 27.
    // Word "xxxz" only get a score of 25.
    fmt.Println(maxScoreWords([]string{"xxxz","ax","bx","cx"},[]byte{'z','a','b','c','x','x','x'},[]int{4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10})) // 27
    // Example 3:
    // Input: words = ["leetcode"], letters = ["l","e","t","c","o","d"], score = [0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0]
    // Output: 0
    // Explanation:
    // Letter "e" can only be used once.
    fmt.Println(maxScoreWords([]string{"leetcode"},[]byte{'l','e','t','c','o','d'},[]int{0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0})) // 0

    fmt.Println(maxScoreWords1([]string{"dog","cat","dad","good"},[]byte{'a','a','c','d','d','d','g','o','o'},[]int{1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0})) // 23
    fmt.Println(maxScoreWords1([]string{"xxxz","ax","bx","cx"},[]byte{'z','a','b','c','x','x','x'},[]int{4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10})) // 27
    fmt.Println(maxScoreWords1([]string{"leetcode"},[]byte{'l','e','t','c','o','d'},[]int{0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0})) // 0

}