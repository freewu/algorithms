package main

// 1178. Number of Valid Words for Each Puzzle
// With respect to a given puzzle string, a word is valid if both the following conditions are satisfied:
//     word contains the first letter of puzzle.
//     For each letter in word, that letter is in puzzle.
//         For example, if the puzzle is "abcdefg", then valid words are "faced", "cabbage", and "baggage", while
//         invalid words are "beefed" (does not include 'a') and "based" (includes 's' which is not in the puzzle).

// Return an array answer, where answer[i] is the number of words in the given word list words that is valid with respect to the puzzle puzzles[i].

// Example 1:
// Input: words = ["aaaa","asas","able","ability","actt","actor","access"], puzzles = ["aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"]
// Output: [1,1,3,2,4,0]
// Explanation: 
// 1 valid word for "aboveyz" : "aaaa" 
// 1 valid word for "abrodyz" : "aaaa"
// 3 valid words for "abslute" : "aaaa", "asas", "able"
// 2 valid words for "absoryz" : "aaaa", "asas"
// 4 valid words for "actresz" : "aaaa", "asas", "actt", "access"
// There are no valid words for "gaswxyz" cause none of the words in the list contains letter 'g'.

// Example 2:
// Input: words = ["apple","pleas","please"], puzzles = ["aelwxyz","aelpxyz","aelpsxy","saelpxy","xaelpsy"]
// Output: [0,1,3,2,0]

// Constraints:
//     1 <= words.length <= 10^5
//     4 <= words[i].length <= 50
//     1 <= puzzles.length <= 10^4
//     puzzles[i].length == 7
//     words[i] and puzzles[i] consist of lowercase English letters.
//     Each puzzles[i] does not contain repeated characters.

import "fmt"

func findNumOfValidWords(words []string, puzzles []string) []int {
    freq := make(map[int]int)
    bits := func (s string) int {
        res := 0
        for _, c := range s {
            // set the bit in the mask
            res |= 1 << (c - 'a')
        }
        return res
    }
    for _, word := range words { // Count how many words are represented by a specific bitmask
        freq[bits(word)]++
    }
    res := make([]int, len(puzzles))
    for i, puzzle := range puzzles {
        mask := bits(puzzle) // get mask for puzzle
        num := 0 // number of words matched in a puzzle
        // set first char bit
        fb := 1 << (puzzle[0] - 'a')
        // Iterate subsets of current puzzle's character set
        // cur = (cur - 1) & mask runs through (0..(1<<N)) & mask
        for cur := mask;cur != 0;cur = ((cur-1) & mask) {
            if (cur & fb) == 0 { continue }
            num += freq[cur]
        }
        res[i] = num
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["aaaa","asas","able","ability","actt","actor","access"], puzzles = ["aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"]
    // Output: [1,1,3,2,4,0]
    // Explanation: 
    // 1 valid word for "aboveyz" : "aaaa" 
    // 1 valid word for "abrodyz" : "aaaa"
    // 3 valid words for "abslute" : "aaaa", "asas", "able"
    // 2 valid words for "absoryz" : "aaaa", "asas"
    // 4 valid words for "actresz" : "aaaa", "asas", "actt", "access"
    // There are no valid words for "gaswxyz" cause none of the words in the list contains letter 'g'.
    fmt.Println(findNumOfValidWords([]string{"aaaa","asas","able","ability","actt","actor","access"}, []string{"aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"})) // [1,1,3,2,4,0]
    // Example 2:
    // Input: words = ["apple","pleas","please"], puzzles = ["aelwxyz","aelpxyz","aelpsxy","saelpxy","xaelpsy"]
    // Output: [0,1,3,2,0]
    fmt.Println(findNumOfValidWords([]string{"apple","pleas","please"}, []string{"aelwxyz","aelpxyz","aelpsxy","saelpxy","xaelpsy"})) // [0,1,3,2,0]
}