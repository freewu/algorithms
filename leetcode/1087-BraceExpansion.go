package main

// 1087. Brace Expansion
// You are given a string s representing a list of words. Each letter in the word has one or more options.
//     If there is one option, the letter is represented as is.
//     If there is more than one option, then curly braces delimit the options. 
//         For example, "{a,b,c}" represents options ["a", "b", "c"].

// For example, if s = "a{b,c}", the first character is always 'a', but the second character can be 'b' or 'c'. 
// The original list is ["ab", "ac"].

// Return all words that can be formed in this manner, sorted in lexicographical order.

// Example 1:
// Input: s = "{a,b}c{d,e}f"
// Output: ["acdf","acef","bcdf","bcef"]

// Example 2:
// Input: s = "abcd"
// Output: ["abcd"]

// Constraints:
//     1 <= s.length <= 50
//     s consists of curly brackets '{}', commas ',', and lowercase English letters.
//     s is guaranteed to be a valid input.
//     There are no nested curly brackets.
//     All characters inside a pair of consecutive opening and ending curly brackets are different.

import "fmt"
import "strings"
import "sort"

func expand(s string) []string {
    res, n := []string{}, len(s)
    var dfs func(cur int , path string)
    dfs = func(cur int , path string){
        if cur >= n {
            res = append(res, path)
            return
        }
        if s[cur] != '{' {
            dfs(cur+1, path + string(s[cur]))
        } else {
            i := cur
            for ; s[i] != '}'; i++ {
            }
            elems := strings.Split(s[cur+1:i], ",")
            sort.Slice(elems, func(i, j int)bool { return elems[i] < elems[j]; })
            for _, elem := range elems {
                dfs(i+1, path + elem)
            }
        }
    }
    dfs(0, "")
    return res
}

func expand1(s string) []string {
    res, choices, start, end := []string{}, []byte{}, -1, len(s) - 1
    for i := 0; i < len(s); i++ {
        if s[i] == '{' {
            start = i
            for j := i + 1; j < len(s); j++ {
                if s[j] == '}' {
                    end = j
                    break
                }
            }
            break
        }
    }
    if start == -1 {
        res = append(res, s)
        return res
    }
    remains := make([]string, 0)
    if end + 1 < len(s) {
        remains = expand(s[end + 1:])
    } else {
        remains = append(remains, "")
    }
    for i := start + 1; i < end; i += 2 {
        choices = append(choices, s[i])
    }
    fix := s[:start]
    for _, ch := range choices {
        for _, remain := range remains {
            var tmp strings.Builder
            tmp.WriteString(fix)
            tmp.WriteByte(ch)
            tmp.WriteString(remain)
            res = append(res, tmp.String())
        }
    }
    mid := sort.StringSlice(res)
    sort.Sort(mid)
    return mid
}

func main() {
    // Example 1:
    // Input: s = "{a,b}c{d,e}f"
    // Output: ["acdf","acef","bcdf","bcef"]
    fmt.Println(expand("{a,b}c{d,e}f")) // ["acdf","acef","bcdf","bcef"]
    // Example 2:
    // Input: s = "abcd"
    // Output: ["abcd"]
    fmt.Println(expand("abcd")) // ["abcd"]

    fmt.Println(expand1("{a,b}c{d,e}f")) // ["acdf","acef","bcdf","bcef"]
    fmt.Println(expand1("abcd")) // ["abcd"]
}