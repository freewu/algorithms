package main

// 2901. Longest Unequal Adjacent Groups Subsequence II
// You are given a string array words, and an array groups, both arrays having length n.

// The hamming distance between two strings of equal length is the number of positions at which the corresponding characters are different.

// You need to select the longest subsequence from an array of indices [0, 1, ..., n - 1], 
// such that for the subsequence denoted as [i0, i1, ..., ik-1] having length k, the following holds:
//     1. For adjacent indices in the subsequence, their corresponding groups are unequal, 
//        i.e., groups[ij] != groups[ij+1], for each j where 0 < j + 1 < k.
//     2. words[ij] and words[ij+1] are equal in length, 
//        and the hamming distance between them is 1, where 0 < j + 1 < k, for all indices in the subsequence.

// Return a string array containing the words corresponding to the indices (in order) in the selected subsequence. 
// If there are multiple answers, return any of them.

// Note: strings in words may be unequal in length.

// Example 1:
// Input: words = ["bab","dab","cab"], groups = [1,2,2]
// Output: ["bab","cab"]
// Explanation: A subsequence that can be selected is [0,2].
//     groups[0] != groups[2]
//     words[0].length == words[2].length, and the hamming distance between them is 1.
// So, a valid answer is [words[0],words[2]] = ["bab","cab"].
// Another subsequence that can be selected is [0,1].
//     groups[0] != groups[1]
//     words[0].length == words[1].length, and the hamming distance between them is 1.
// So, another valid answer is [words[0],words[1]] = ["bab","dab"].
// It can be shown that the length of the longest subsequence of indices that satisfies the conditions is 2.

// Example 2:
// Input: words = ["a","b","c","d"], groups = [1,2,3,4]
// Output: ["a","b","c","d"]
// Explanation: We can select the subsequence [0,1,2,3].
// It satisfies both conditions.
// Hence, the answer is [words[0],words[1],words[2],words[3]] = ["a","b","c","d"].
// It has the longest length among all subsequences of indices that satisfy the conditions.
// Hence, it is the only answer.

// Constraints:
//     1 <= n == words.length == groups.length <= 1000
//     1 <= words[i].length <= 10
//     1 <= groups[i] <= n
//     words consists of distinct strings.
//     words[i] consists of lowercase English letters.

import "fmt"

func getWordsInLongestSubsequence(words []string, groups []int) []string {
    res, n := []string{}, len(words)
    length, hash := make([]int, n), make([]int, n)
    mx, index := 1, 0 // took mx as 1 because atleast one word can be taken to form a subsequence of length 1
    for i := 0; i < n; i++ {
        length[i] = 1 // initialise with 1 cause min length of sub sequence that can end with i is atleast one at begining
        hash[i] = i   // at begining initialise with its own index
    }
    isOne := func(s, t string) bool {
        if len(s) != len(t) { return false }
        count := 0
        for i :=0 ; i < len(s); i++ {
            if s[i] != t[i]{
                count++
                if count > 1 { return false }
            }
        }
        return count == 1
    }
    // if i am standing at 'i' th index i am checking for how many indexes 
    // I can join my word at 'i' it is posssible to join if satisfies the given condition 
    for i := 1; i < n; i++ {
        for j := i - 1; j >= 0; j-- {
            if isOne(words[i],words[j]) && groups[i] != groups[j] {
                if length[i] < 1 + length[j] {
                    hash[i] = j // keep track where the subsequence is joining
                    length[i] = 1 + length[j] // increase the length if word is joined
                }
            }
        }
    }
    for i := 0; i < n; i++ { // find the max length index and start going back
        if length[i] > mx {
            mx, index = length[i], i
        }
    }
    for index != hash[index] { // if hash[i]!=i then subsequence starts from there
        res = append(res, words[index])
        index = hash[index]
    }
    res = append(res, words[index]) // end
    for i, j := 0, len(res) - 1; i < j; i, j = i + 1, j - 1 { // reverse
        res[i], res[j] = res[j], res[i]
    }
    return res
}

func getWordsInLongestSubsequence1(words []string, groups []int) []string {
    type Pair struct{ mx, i int }
    n := len(words)
    mp, from := map[int]Pair{}, make([]int, n)
    count, index := 0, 0
    for i := n - 1; i >= 0; i-- {
        word, g := words[i], groups[i]
        // 计算 word 的哈希值
        hash := 0
        for _, ch := range word {
            hash = hash << 5 | int(ch & 31)
        }
        // 计算方法一中的 f[i]
        f := 0
        for j := range word {
            h := hash | 31 << (j * 5) // 用记号笔把 w[k] 涂黑（置为 11111）
            t := mp[h]
            if t.mx > f && g != groups[t.i] {
                f = t.mx
                from[i] = t.i
            }
        }
        f++
        if f > count {
            count, index = f, i
        }
        // 用 f 更新 mp[h]
        for j := range word {
            h := hash | 31 << (j * 5)
            if f > mp[h].mx {
                mp[h] = Pair{f, i}
            }
        }
    }
    res, i := make([]string, count), index
    for k := range res {
        res[k] = words[i]
        i = from[i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["bab","dab","cab"], groups = [1,2,2]
    // Output: ["bab","cab"]
    // Explanation: A subsequence that can be selected is [0,2].
    //     groups[0] != groups[2]
    //     words[0].length == words[2].length, and the hamming distance between them is 1.
    // So, a valid answer is [words[0],words[2]] = ["bab","cab"].
    // Another subsequence that can be selected is [0,1].
    //     groups[0] != groups[1]
    //     words[0].length == words[1].length, and the hamming distance between them is 1.
    // So, another valid answer is [words[0],words[1]] = [].
    // It can be shown that the length of the longest subsequence of indices that satisfies the conditions is 2.
    fmt.Println(getWordsInLongestSubsequence([]string{"bab","dab","cab"}, []int{1,2,2})) // ["bab","dab"]
    // Example 2:
    // Input: words = ["a","b","c","d"], groups = [1,2,3,4]
    // Output: ["a","b","c","d"]
    // Explanation: We can select the subsequence [0,1,2,3].
    // It satisfies both conditions.
    // Hence, the answer is [words[0],words[1],words[2],words[3]] = ["a","b","c","d"].
    // It has the longest length among all subsequences of indices that satisfy the conditions.
    // Hence, it is the only answer.
    fmt.Println(getWordsInLongestSubsequence([]string{"a","b","c","d"}, []int{1,2,3,4})) // ["a","b","c","d"]

    fmt.Println(getWordsInLongestSubsequence([]string{"b","l","u","e","f","r","o","g"}, []int{1,2,3,4,1,2,3,4})) // [b l u e f r o g]
    fmt.Println(getWordsInLongestSubsequence([]string{"l","e","e","t","c","o","d","e"}, []int{1,2,3,4,1,2,3,4})) // [l e t c o d e]

    fmt.Println(getWordsInLongestSubsequence1([]string{"bab","dab","cab"}, []int{1,2,2})) // ["bab","dab"]
    fmt.Println(getWordsInLongestSubsequence1([]string{"a","b","c","d"}, []int{1,2,3,4})) // ["a","b","c","d"]
    fmt.Println(getWordsInLongestSubsequence1([]string{"b","l","u","e","f","r","o","g"}, []int{1,2,3,4,1,2,3,4})) // [b l u e f r o g]
    fmt.Println(getWordsInLongestSubsequence1([]string{"l","e","e","t","c","o","d","e"}, []int{1,2,3,4,1,2,3,4})) // [l e t c o d e]
}