package main

// 面试题 17.22. Word Transformer LCCI
// Given two words of equal length that are in a dictionary, write a method to transform one word into another word by changing only one letter at a time. 
// The new word you get in each step must be in the dictionary.

// Write code to return a possible transforming sequence. 
// If there is more than one sequence, return any of them.

// Example 1:
// Input:
// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]
// Output:
// ["hit","hot","dot","lot","log","cog"]

// Example 2:
// Input:
// beginWord = "hit"
// endWord = "cog"
// wordList = ["hot","dot","dog","lot","log"]
// Output: []
// Explanation: endWord "cog" is not in the dictionary, so there's no possible transforming sequence.

import "fmt"

func findLadders(beginWord string, endWord string, wordList []string) []string {
    type WordNode struct {
        word string
        path []string
    }
    wordList = append(wordList, beginWord)
    n := len(beginWord)
    mp := make(map[string][]string)
    // 预处理单词列表，构建通用状态映射
    for _, word := range wordList {
        for i := 0; i < n; i++ {
            w := word[:i] + "*" + word[i+1:]
            mp[w] = append(mp[w], word)
        }
    }
    // 使用队列进行广度优先搜索
    queue := []WordNode{{ word: beginWord, path: []string{ beginWord }}}
    visited := make(map[string]bool)
    visited[beginWord] = true
    for len(queue) > 0 {
        currentNode := queue[0]
        queue = queue[1:]
        currentWord := currentNode.word
        currentPath := currentNode.path
        for i := 0; i < n; i++ {
            intermediateWord := currentWord[:i] + "*" + currentWord[i+1:]
            for _, word := range mp[intermediateWord] {
                if word == endWord {
                    return append(currentPath, endWord)
                }
                if !visited[word] {
                    visited[word] = true
                    newPath := make([]string, len(currentPath))
                    copy(newPath, currentPath)
                    newPath = append(newPath, word)
                    queue = append(queue, WordNode{word: word, path: newPath})
                }
            }
        }
    }
    return []string{} // 如果没有找到路径，返回空切片
}

func main() {
    // Example 1:
    // Input:
    // beginWord = "hit",
    // endWord = "cog",
    // wordList = ["hot","dot","dog","lot","log","cog"]
    // Output:
    // ["hit","hot","dot","lot","log","cog"]
    fmt.Println(findLadders("hit", "cog", []string{"hot","dot","dog","lot","log","cog"})) // ["hit","hot","dot","lot","log","cog"]
    // Example 2:
    // Input:
    // beginWord = "hit"
    // endWord = "cog"
    // wordList = ["hot","dot","dog","lot","log"]
    // Output: []
    // Explanation: endWord "cog" is not in the dictionary, so there's no possible transforming sequence.
    fmt.Println(findLadders("hit", "cog", []string{"hot","dot","dog","lot","log"})) // []
}