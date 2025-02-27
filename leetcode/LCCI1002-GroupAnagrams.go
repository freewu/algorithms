package main

// 面试题 10.02. Group Anagrams LCCI
// Write a method to sort an array of strings so that all the anagrams are in the same group.

// Note: This problem is slightly different from the original one the book.

// Example:
// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output: [["ate","eat","tea"],["nat","tan"],["bat"]]

// Notes:
//     All inputs will be in lowercase.
//     The order of your output does not matter.

import "fmt"
import "sort"

func groupAnagrams(strs []string) [][]string {
    mp := make(map[string][]string)
    for _, v := range strs {
        arr := []byte(v)
        sort.Slice(arr,func(i, j int) bool {
            return arr[i] < arr[j]
        })
        s := string(arr)
        mp[s] = append(mp[s], v)
    }
    res := [][]string{}
    for _, v := range mp {
        res = append(res, v)
    }
    return res
}

func groupAnagrams1(strs []string) [][]string {
    anagrams := make(map[string][]string) // sorted version of the word -> list of words
    sortCharacters := func(word string) string {
        arr := []byte(word)
        sort.Slice(arr,func(i, j int) bool {
            return arr[i] < arr[j]
        })
        return string(arr)
    }
    for _, word := range strs {
        sortedWord := sortCharacters(word)
        anagrams[sortedWord] = append(anagrams[sortedWord], word)
    }
    res := make([][]string, 0)
    for _, v := range anagrams {
        res = append(res, v)
    }
    return res
}

func main() {
    // Example:
    // Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
    // Output: [["ate","eat","tea"],["nat","tan"],["bat"]]
    fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})) // [["ate","eat","tea"],["nat","tan"],["bat"]]
    
    fmt.Println(groupAnagrams([]string{"bluefrog", "leetcode"})) // [[bluefrog] [leetcode]]

    fmt.Println(groupAnagrams1([]string{"eat", "tea", "tan", "ate", "nat", "bat"})) // [["ate","eat","tea"],["nat","tan"],["bat"]]
    fmt.Println(groupAnagrams1([]string{"bluefrog", "leetcode"})) // [[bluefrog] [leetcode]]
}