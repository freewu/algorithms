package main

// 1170. Compare Strings by Frequency of the Smallest Character
// Let the function f(s) be the frequency of the lexicographically smallest character in a non-empty string s. 
// For example, if s = "dcce" then f(s) = 2 because the lexicographically smallest character is 'c', which has a frequency of 2.

// You are given an array of strings words and another array of query strings queries. 
// For each query queries[i], count the number of words in words such that f(queries[i]) < f(W) for each W in words.

// Return an integer array answer, where each answer[i] is the answer to the ith query.

// Example 1:
// Input: queries = ["cbd"], words = ["zaaaz"]
// Output: [1]
// Explanation: On the first query we have f("cbd") = 1, f("zaaaz") = 3 so f("cbd") < f("zaaaz").

// Example 2:
// Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
// Output: [1,2]
// Explanation: On the first query only f("bbb") < f("aaaa"). On the second query both f("aaa") and f("aaaa") are both > f("cc").

// Constraints:
//     1 <= queries.length <= 2000
//     1 <= words.length <= 2000
//     1 <= queries[i].length, words[i].length <= 10
//     queries[i][j], words[i][j] consist of lowercase English letters.

import "fmt"
import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
    calc := func(s string) int {
        arr, mn := make([]int, 26), 25
        for _, c := range s {
            i := int(c - 'a')
            arr[i]++
            if i < mn {
                mn = i
            }
        }
        return arr[mn]
    }
    freq := make([]int, 12)
    for _, word := range words {
        freq[calc(word)]++
    }
    sum := 0
    for i := len(freq) - 1; i >= 0; i-- {
        sum += freq[i]
        freq[i] = sum
    }
    res := make([]int, len(queries))
    for i, query := range queries {
        res[i] = freq[calc(query)+1]
    }
    return res
}

func numSmallerByFrequency1(queries []string, words []string) []int {
    qcount, wcount := make([]int, len(queries)), make([]int, len(words))
    for i := 0; i < len(queries); i++ {
        count := [26]int{}
        for j := 0; j < len(queries[i]); j++ {
            count[queries[i][j] - 'a']++
        }
        for k := 0; k < 26; k++ {
            if count[k] != 0 { // 找到第一个不为0的数
                qcount[i] = count[k]
                break
            }
        }
    }
    for i := 0; i < len(words); i++ {
        count := [26]int{}
        for j := 0; j < len(words[i]); j++ {
            count[words[i][j] - 'a']++
        }
        for k := 0; k < 26; k++ {
            if count[k] != 0 { // 结束之后，找到第一个不为0的数
                wcount[i] = count[k]
                break
            }
        }
    }
    binarySearch := func(nums []int, target int) int {
        l, r := 0, len(nums) - 1 // 注意 f(-1)其实是illegal的而f(n)其实是legal的，n是数组长度
        for l <= r {
            m := l + (r - l) / 2
            if nums[m] < target {
                l = m + 1
            } else {
                r = m - 1
            }
        }
        return l
    }
    res := make([]int, len(queries))
    sort.Ints(wcount)
    for i := 0; i < len(qcount); i++  {
        // 找到第一个大于qCnt[i]的下标, 由于数组是有序的，所以剩下的部分也是有序的，可以直接做减法
        res[i] = len(wcount) - binarySearch(wcount, qcount[i] + 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: queries = ["cbd"], words = ["zaaaz"]
    // Output: [1]
    // Explanation: On the first query we have f("cbd") = 1, f("zaaaz") = 3 so f("cbd") < f("zaaaz").
    fmt.Println(numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"})) // [1]
    // Example 2:
    // Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
    // Output: [1,2]
    // Explanation: On the first query only f("bbb") < f("aaaa"). On the second query both f("aaa") and f("aaaa") are both > f("cc").
    fmt.Println(numSmallerByFrequency([]string{"bbb","cc"}, []string{"a","aa","aaa","aaaa"})) // [1,2]

    fmt.Println(numSmallerByFrequency1([]string{"cbd"}, []string{"zaaaz"})) // [1]
    fmt.Println(numSmallerByFrequency1([]string{"bbb","cc"}, []string{"a","aa","aaa","aaaa"})) // [1,2]
}