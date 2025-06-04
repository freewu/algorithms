package main

// 3403. Find the Lexicographically Largest String From the Box I
// You are given a string word, and an integer numFriends.

// Alice is organizing a game for her numFriends friends. 
// There are multiple rounds in the game, where in each round:
//     word is split into numFriends non-empty strings, such that no previous round has had the exact same split.
//     All the split words are put into a box.

// Find the lexicographically largest string from the box after all the rounds are finished.

// A string a is lexicographically smaller than a string b if in the first position where a and b differ, 
// string a has a letter that appears earlier in the alphabet than the corresponding letter in b.

// If the first min(a.length, b.length) characters do not differ, 
// then the shorter string is the lexicographically smaller one.

// Example 1:
// Input: word = "dbca", numFriends = 2
// Output: "dbc"
// Explanation: 
// All possible splits are:
// "d" and "bca".
// "db" and "ca".
// "dbc" and "a".

// Example 2:
// Input: word = "gggg", numFriends = 4
// Output: "g"
// Explanation: 
// The only possible split is: "g", "g", "g", and "g".

// Constraints:
//     1 <= word.length <= 5 * 10^3
//     word consists only of lowercase English letters.
//     1 <= numFriends <= word.length

import "fmt"
import "strings"

func answerString(word string, numFriends int) string {
    if numFriends == 1 { return word }
    n, i, j := len(word), 0, 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for j < n {
        v := 0
        for j + v < n && word[i + v] == word[j + v] {
            v++
        }
        if j + v < n && word[i + v] < word[j + v] {
            i, j = j, max(j + 1, i + v + 1)
        } else {
            j += (v + 1)
        }
    }
    return word[i:min(i + n - numFriends + 1, n)]
}

func answerString1(word string, numFriends int) string {
    if numFriends == 1 { return word }
    res, n := "", len(word)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y string) string { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        res = max(res, word[i:min(i + n - numFriends + 1, n)])
    }
    return res
}

func answerString2(word string, numFriends int) string {
    if numFriends == 1 { return word }
    res, n := "", len(word)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < n; i++ {
        s := word[i:min(n, i + n - numFriends + 1)]
        if strings.Compare(s, res) > 0 {
            res = s
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: word = "dbca", numFriends = 2
    // Output: "dbc"
    // Explanation:
    // All possible splits are:
    // "d" and "bca".
    // "db" and "ca".
    // "dbc" and "a".
    fmt.Println(answerString("dbca", 2)) // "dbc"
    // Example 2:
    // Input: word = "gggg", numFriends = 4
    // Output: "g"
    // Explanation:
    // The only possible split is: "g", "g", "g", and "g".
    fmt.Println(answerString("gggg", 4)) // "g"

    fmt.Println(answerString("bluefrog", 4)) // "uefro"
    fmt.Println(answerString("leetcode", 4)) // "tcode"

    fmt.Println(answerString1("dbca", 2)) // "dbc"
    fmt.Println(answerString1("gggg", 4)) // "g"
    fmt.Println(answerString1("bluefrog", 4)) // "uefro"
    fmt.Println(answerString1("leetcode", 4)) // "tcode"

    fmt.Println(answerString2("dbca", 2)) // "dbc"
    fmt.Println(answerString2("gggg", 4)) // "g"
    fmt.Println(answerString2("bluefrog", 4)) // "uefro"
    fmt.Println(answerString2("leetcode", 4)) // "tcode"
}