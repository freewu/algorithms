package main

// 1177. Can Make Palindrome from Substring
// You are given a string s and array queries where queries[i] = [lefti, righti, ki]. 
// We may rearrange the substring s[lefti...righti] for each query and then choose up to ki of them to replace with any lowercase English letter.

// If the substring is possible to be a palindrome string after the operations above, the result of the query is true. 
// Otherwise, the result is false.

// Return a boolean array answer where answer[i] is the result of the ith query queries[i].

// Note that each letter is counted individually for replacement, so if, for example s[lefti...righti] = "aaa", 
// and ki = 2, we can only replace two of the letters. 
// Also, note that no query modifies the initial string s.

// Example :
// Input: s = "abcda", queries = [[3,3,0],[1,2,0],[0,3,1],[0,3,2],[0,4,1]]
// Output: [true,false,false,true,true]
// Explanation:
// queries[0]: substring = "d", is palidrome.
// queries[1]: substring = "bc", is not palidrome.
// queries[2]: substring = "abcd", is not palidrome after replacing only 1 character.
// queries[3]: substring = "abcd", could be changed to "abba" which is palidrome. Also this can be changed to "baab" first rearrange it "bacd" then replace "cd" with "ab".
// queries[4]: substring = "abcda", could be changed to "abcba" which is palidrome.

// Example 2:
// Input: s = "lyb", queries = [[0,1,0],[2,2,1]]
// Output: [false,true]

// Constraints:
//     1 <= s.length, queries.length <= 10^5
//     0 <= lefti <= righti < s.length
//     0 <= ki <= s.length
//     s consists of lowercase English letters.

import "fmt"
import "math/bits"

func canMakePaliQueries(s string, queries [][]int) []bool {
    res, c := []bool{}, [][]int{}
    c = append(c, make([]int, 26))
    for i := 0; i < len(s); i ++{
        if i > 0 {
            c = append(c, make([]int, 26))
            for j := 0; j < 26; j ++{
                c[i][j] = c[i - 1][j]
            }
        }
        c[i][s[i] - 'a'] ++
    }
    for i := 0; i < len(queries); i ++{
        count := 0
        for j := 0; j < 26; j ++{
            t := c[queries[i][1]][j]
            if queries[i][0] != 0 {
                t -= c[queries[i][0] - 1][j]
            }
            if t % 2 != 0 {
                count++
            }
        }
        res = append(res, ((count / 2) <= queries[i][2]))
    }    
    return res
}

func canMakePaliQueries1(s string, queries [][]int) []bool {
    prefix := make([]uint32, len(s) + 1)
    for i, c := range s {
        bit := uint32(1) << (c - 'a')
        prefix[i+1] = prefix[i] ^ bit
    }
    res := make([]bool, len(queries))
    for i, query := range queries {
        left, right, k := query[0], query[1], query[2]
        res[i] = (bits.OnesCount32(prefix[right + 1] ^ prefix[left]) / 2) <= k
    }
    return res
}

func main() {
    // Example :
    // Input: s = "abcda", queries = [[3,3,0],[1,2,0],[0,3,1],[0,3,2],[0,4,1]]
    // Output: [true,false,false,true,true]
    // Explanation:
    // queries[0]: substring = "d", is palidrome.
    // queries[1]: substring = "bc", is not palidrome.
    // queries[2]: substring = "abcd", is not palidrome after replacing only 1 character.
    // queries[3]: substring = "abcd", could be changed to "abba" which is palidrome. Also this can be changed to "baab" first rearrange it "bacd" then replace "cd" with "ab".
    // queries[4]: substring = "abcda", could be changed to "abcba" which is palidrome.
    fmt.Println(canMakePaliQueries("abcda",[][]int{{3,3,0},{1,2,0},{0,3,1},{0,3,2},{0,4,1}})) // [true,false,false,true,true]
    // Example 2:
    // Input: s = "lyb", queries = [[0,1,0],[2,2,1]]
    // Output: [false,true]
    fmt.Println(canMakePaliQueries("lyb",[][]int{{0,1,0},{2,2,1}})) // [false,true]

    fmt.Println(canMakePaliQueries1("abcda",[][]int{{3,3,0},{1,2,0},{0,3,1},{0,3,2},{0,4,1}})) // [true,false,false,true,true]
    fmt.Println(canMakePaliQueries1("lyb",[][]int{{0,1,0},{2,2,1}})) // [false,true]
}