package main

// 面试题 17.11. Find Closest LCCI
// You have a large text file containing words. 
// Given any two different words, find the shortest distance (in terms of number of words) between them in the file. 
// If the operation will be repeated many times for the same file (but different pairs of words), can you optimize your solution?

// Example:
// Input: words = ["I","am","a","student","from","a","university","in","a","city"], word1 = "a", word2 = "student"
// Output: 1

// Note:
//     words.length <= 100000

import "fmt"

func findClosest(words []string, word1 string, word2 string) int {
    mp := make(map[int]string)
    for i, word := range words {
        mp[i] = word
    }
    arr1, arr2 := []int{}, []int{}
    for k, v := range mp { // 找到单词对应的所有索引
        if v == word1 {
            arr1 = append(arr1, k)
        }
        if v == word2 {
            arr2 = append(arr2, k)
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    res := 1 << 31
    for _, v1 := range arr1 {
        for _, v2 := range arr2 {
            res = min(res, abs(v1 - v2))
        }
    }
    return res
}

func findClosest1(words []string, word1 string, word2 string) int {
    res, i1, i2 := 1 << 31, -1, -1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i, v := range words {
        if v == word1 {
            i1 = i
        } else if v == word2 {
            i2 = i
        }
        if i1 >= 0 && i2 >= 0 {
            res = min(res, abs(i1 - i2))
        }
    }
    return res
}

func main() {
    // Example:
    // Input: words = ["I","am","a","student","from","a","university","in","a","city"], word1 = "a", word2 = "student"
    // Output: 1
    fmt.Println(findClosest([]string{"I","am","a","student","from","a","university","in","a","city"}, "a", "student")) // 1

    fmt.Println(findClosest1([]string{"I","am","a","student","from","a","university","in","a","city"}, "a", "student")) // 1
}