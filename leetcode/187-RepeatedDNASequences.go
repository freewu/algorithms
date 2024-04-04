package main

// 187. Repeated DNA Sequences 
// The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.
//     For example, "ACGAATTCCG" is a DNA sequence.

// When studying DNA, it is useful to identify repeated sequences within the DNA.
// Given a string s that represents a DNA sequence, 
// return all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule. 
// You may return the answer in any order.

// Example 1:
// Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
// Output: ["AAAAACCCCC","CCCCCAAAAA"]

// Example 2:
// Input: s = "AAAAAAAAAAAAA"
// Output: ["AAAAAAAAAA"]

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either 'A', 'C', 'G', or 'T'.

import "fmt"

func findRepeatedDnaSequences(s string) []string {
    res := []string{}
    m := map[string]int{}
    for i := 0; i <= len(s) - 10; i++ {
        // 从 0 到 len(s) - 10 每次取10个字符做 key
        tmp := s[i : i + 10]
        m[tmp]++
        if m[tmp] == 2 { // 存在重复加入到结果集中
            res = append(res, tmp)
            m[tmp]++
        }
    }
    return res
}

func main() {
    fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")) // ["AAAAACCCCC","CCCCCAAAAA"]
    fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA")) // ["AAAAAAAAAA"]
}