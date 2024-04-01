package main

// 936. Stamping The Sequence
// You are given two strings stamp and target. Initially, there is a string s of length target.length with all s[i] == '?'.
// In one turn, you can place stamp over s and replace every letter in the s with the corresponding letter from stamp.
// For example, if stamp = "abc" and target = "abcba", then s is "?????" initially. In one turn you can:
//     place stamp at index 0 of s to obtain "abc??",
//     place stamp at index 1 of s to obtain "?abc?", or
//     place stamp at index 2 of s to obtain "??abc".

// Note that stamp must be fully contained in the boundaries of s in order to stamp (i.e., you cannot place stamp at index 3 of s).
// We want to convert s to target using at most 10 * target.length turns.

// Return an array of the index of the left-most letter being stamped at each turn. 
// If we cannot obtain target from s within 10 * target.length turns, return an empty array.

// Example 1:
// Input: stamp = "abc", target = "ababc"
// Output: [0,2]
// Explanation: Initially s = "?????".
// - Place stamp at index 0 to get "abc??".
// - Place stamp at index 2 to get "ababc".
// [1,0,2] would also be accepted as an answer, as well as some other answers.

// Example 2:
// Input: stamp = "abca", target = "aabcaca"
// Output: [3,0,1]
// Explanation: Initially s = "???????".
// - Place stamp at index 3 to get "???abca".
// - Place stamp at index 0 to get "abcabca".
// - Place stamp at index 1 to get "aabcaca".
 
// Constraints:
//     1 <= stamp.length <= target.length <= 1000
//     stamp and target consist of lowercase English letters.

import "fmt"
import "slices"

func movesToStamp(stamp string, target string) []int {
    letters, toErase := []byte(target), len(target)
    contains := func(offset int) bool {
        empty := true
        for i := 0; i < len(stamp); i++ {
            if l := letters[offset+i]; l != '.' && l != stamp[i] {
                return false
            } else if l != '.' {
                empty = false
            }
        }
        return !empty
    }
    erase := func(offset int) {
        for i := 0; i < len(stamp); i++ {
            if letters[offset+i] != '.' {
                letters[offset+i] = '.'
                toErase--
            }
        }
    }
    var result []int
    for toErase != 0 {
        erased := false
        for i := 0; i < len(letters)-len(stamp)+1; i++ {
            if contains(i) {
                erase(i)
                result = append(result, i)
                erased = true
                i += len(stamp)
                if toErase == 0 {
                    break
                }
            }
        }
        if !erased {
            return []int{}
        }
    }
    for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
        result[l], result[r] = result[r], result[l]
    }
    return result
}

func movesToStamp1(stamp string, target string) []int {
    n := len(stamp)
    s := []byte(target)
    var res []int
    var check = func(start int) {
        for i := 0; i < n; i++ {
            j := i + start
            if stamp[i] != s[j] && s[j] != '#' { return; }
        }
        flag := false
        for i := 0; i < n; i++ {
            j := i + start
            if s[j] != '#' {  flag = true; }
            s[j] = '#'
        }
        if flag {  res = append(res, start); }
    }
    for i := 0; i+n-1 < len(s); i++ { check(i);  }
    for i := len(s)-n; i >= 0; i-- {  check(i); }
    for i := range s{
        if s[i] != '#' { return []int{}; }
    }
    slices.Reverse(res)
    return res
}

func main() {
    // Explanation: Initially s = "?????".
    // - Place stamp at index 0 to get "abc??".
    // - Place stamp at index 2 to get "ababc".
    // [1,0,2] would also be accepted as an answer, as well as some other answers.
    fmt.Println(movesToStamp("abc","ababc")) // [0,2]

    // Explanation: Initially s = "???????".
    // - Place stamp at index 3 to get "???abca".
    // - Place stamp at index 0 to get "abcabca".
    // - Place stamp at index 1 to get "aabcaca".
    fmt.Println(movesToStamp("abca","aabcaca")) // [3,0,1]

    fmt.Println(movesToStamp1("abc","ababc")) // [0,2]
    fmt.Println(movesToStamp1("abca","aabcaca")) // [3,0,1]
}