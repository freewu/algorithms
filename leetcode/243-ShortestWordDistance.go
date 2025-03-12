package main

// 243. Shortest Word Distance
// Given an array of strings wordsDict and two different strings that already exist in the array word1 and word2, 
// return the shortest distance between these two words in the list.

// Example 1:
// Input: wordsDict = ["practice", "makes", "perfect", "coding", "makes"], word1 = "coding", word2 = "practice"
// Output: 3

// Example 2:
// Input: wordsDict = ["practice", "makes", "perfect", "coding", "makes"], word1 = "makes", word2 = "coding"
// Output: 1

// Constraints:
//     2 <= wordsDict.length <= 3 * 104
//     1 <= wordsDict[i].length <= 10
//     wordsDict[i] consists of lowercase English letters.
//     word1 and word2 are in wordsDict.
//     word1 != word2

import "fmt"

func shortestDistance(wordsDict []string, word1, word2 string) int {
    res, index1, index2 := len(wordsDict), -1, -1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    // 从左到右遍历数组 wordsDict，
    // 当遍历到 word1时，
    //     如果已经遍历的单词中存在 word2，为了计算最短距离，应该取最后一个已经遍历到的 word2 所在的下标，计算和当前下标的距离。
    //     同理，当遍历到 word2  时，应该取最后一个已经遍历到的 word1 所在的下标，计算和当前下标的距离。
    for i, word := range wordsDict {
        if word == word1 { // 找到单词1
            index1 = i
        } else if word == word2 { // 找到单词2
            index2 = i
        }
        // 都找到了，计算距离
        if index1 >= 0 && index2 >= 0 {
            res = min(res, abs(index1 - index2))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: wordsDict = ["practice", "makes", "perfect", "coding", "makes"], word1 = "coding", word2 = "practice"
    // Output: 3
    fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"},"coding","practice",)) // 3
    // Example 2:
    // Input: wordsDict = ["practice", "makes", "perfect", "coding", "makes"], word1 = "makes", word2 = "coding"
    // Output: 1
    fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"},"makes","coding",)) // 1
}