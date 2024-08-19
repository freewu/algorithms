package main

// 1307. Verbal Arithmetic Puzzle
// Given an equation, represented by words on the left side and the result on the right side.

// You need to check if the equation is solvable under the following rules:
//     Each character is decoded as one digit (0 - 9).
//     No two characters can map to the same digit.
//     Each words[i] and result are decoded as one number without leading zeros.
//     Sum of numbers on the left side (words) will equal to the number on the right side (result).

// Return true if the equation is solvable, otherwise return false.

// Example 1:
// Input: words = ["SEND","MORE"], result = "MONEY"
// Output: true
// Explanation: Map 'S'-> 9, 'E'->5, 'N'->6, 'D'->7, 'M'->1, 'O'->0, 'R'->8, 'Y'->'2'
// Such that: "SEND" + "MORE" = "MONEY" ,  9567 + 1085 = 10652

// Example 2:
// Input: words = ["SIX","SEVEN","SEVEN"], result = "TWENTY"
// Output: true
// Explanation: Map 'S'-> 6, 'I'->5, 'X'->0, 'E'->8, 'V'->7, 'N'->2, 'T'->1, 'W'->'3', 'Y'->4
// Such that: "SIX" + "SEVEN" + "SEVEN" = "TWENTY" ,  650 + 68782 + 68782 = 138214

// Example 3:
// Input: words = ["LEET","CODE"], result = "POINT"
// Output: false
// Explanation: There is no possible mapping to satisfy the equation, so we return false.
// Note that two different characters cannot map to the same digit.
 
// Constraints:
//     2 <= words.length <= 5
//     1 <= words[i].length, result.length <= 7
//     words[i], result contain only uppercase English letters.
//     The number of different characters used in the expression is at most 10.

import "fmt"
import "sort"

func isSolvable(words []string, result string) bool {
    // charMap maps each unique char to sum of base values
    // firstChars keeps tracks if a char can't be 0 (ie. used as first char in a word)
    charMap, firstChars := make(map[byte]int), make(map[byte]bool)
    for _, word := range words {
        base := 1
        for i := len(word) - 1; i >= 0; i-- {
            charMap[word[i]] += base
            base *= 10
        }
        if len(word) > 1 { // special case: if word is length first char can be 0
            firstChars[word[0]] = true
        }
    }
    base := 1
    for i := len(result) - 1; i >= 0; i-- {
        charMap[result[i]] -= base
        base *= 10
    }
    if len(result) > 1 {
        firstChars[result[0]] = true
    }
    
    // turn into map into array: each index represents a character
    // separate "zeroable" chars in the array so it is easy to check during DFS
    nonZeroChars, zeroChars := make([]int, 0), make([]int, 0)
    sort.Ints(nonZeroChars)
    sort.Ints(zeroChars)
    
    for k, v := range charMap {
        if firstChars[k] == true {
            nonZeroChars = append(nonZeroChars, v)
        } else {
            zeroChars = append(zeroChars, v)
        }
    }
    chars := append(nonZeroChars, zeroChars...)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    // keep a sums suffix array for pruning
    sums := make([]int, 1)
    for i := len(chars) - 1; i >= 0; i-- {
        sums = append([]int{abs(chars[i])*9 + sums[0]}, sums...)
    }

    // used is a bitmask to quickly tell if we have used an array index of not
    var dfs func(chars []int, zeroInd int, ind int, used int, curSum int, sums []int) bool
    dfs = func(chars []int, zeroInd int, ind int, used int, curSum int, sums []int) bool {
        // whenever we reach end, check if curSum is zero (ie. if answer to problem is true)
        if ind == len(chars) {
            return curSum == 0
        // if at any point, our current sum exceeds the maximum value we could generate
        // with the characters that are unused, return false
        } else if abs(curSum) > sums[ind] {
            return false
        }
        for i := 0; i < 10; i++ {
            // if used or we are not allowed to use zero, continue
            if ((used >> i) & 1) == 1 || (i == 0 && ind < zeroInd) {
                continue
            }
            // assign char using bit mask, add to curSum
            if dfs(chars, zeroInd, ind + 1, used | (1 << i), curSum + chars[ind]*i, sums) {
                return true
            }
        }
        return false
    }
    
    // since we separated the nonZeroChars, we can just pass in the index at which
    // "zeroable" characters start
    return dfs(chars, len(nonZeroChars), 0, 0, 0, sums)
}

func main() {
    // Example 1:
    // Input: words = ["SEND","MORE"], result = "MONEY"
    // Output: true
    // Explanation: Map 'S'-> 9, 'E'->5, 'N'->6, 'D'->7, 'M'->1, 'O'->0, 'R'->8, 'Y'->'2'
    // Such that: "SEND" + "MORE" = "MONEY" ,  9567 + 1085 = 10652
    fmt.Println(isSolvable([]string{"SEND","MORE"}, "MONEY")) // true
    // Example 2:
    // Input: words = ["SIX","SEVEN","SEVEN"], result = "TWENTY"
    // Output: true
    // Explanation: Map 'S'-> 6, 'I'->5, 'X'->0, 'E'->8, 'V'->7, 'N'->2, 'T'->1, 'W'->'3', 'Y'->4
    // Such that: "SIX" + "SEVEN" + "SEVEN" = "TWENTY" ,  650 + 68782 + 68782 = 138214
    fmt.Println(isSolvable([]string{"SIX","SEVEN","SEVEN"}, "TWENTY")) // true
    // Example 3:
    // Input: words = ["LEET","CODE"], result = "POINT"
    // Output: false
    // Explanation: There is no possible mapping to satisfy the equation, so we return false.
    // Note that two different characters cannot map to the same digit.
    fmt.Println(isSolvable([]string{"LEET","CODE"}, "POINT")) // false
}