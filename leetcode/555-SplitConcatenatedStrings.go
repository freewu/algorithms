package main

// 555. Split Concatenated Strings
// You are given an array of strings strs. 
// You could concatenate these strings together into a loop, where for each string, 
// you could choose to reverse it or not. Among all the possible loops

// Return the lexicographically largest string after cutting the loop,
// which will make the looped string into a regular one.

// Specifically, to find the lexicographically largest string, you need to experience two phases:

//     1. Concatenate all the strings into a loop, where you can reverse some strings 
//        or not and connect them in the same order as given.
//     2. Cut and make one breakpoint in any place of the loop, 
//        which will make the looped string into a regular one starting from the character at the cutpoint.

// And your job is to find the lexicographically largest one among all the possible regular strings.

// Example 1:
// Input: strs = ["abc","xyz"]
// Output: "zyxcba"
// Explanation: You can get the looped string "-abcxyz-", "-abczyx-", "-cbaxyz-", "-cbazyx-", where '-' represents the looped status. 
// The answer string came from the fourth looped one, where you could cut from the middle character 'a' and get "zyxcba".

// Example 2:
// Input: strs = ["abc"]
// Output: "cba"

// Constraints:
//     1 <= strs.length <= 1000
//     1 <= strs[i].length <= 1000
//     1 <= sum(strs[i].length) <= 1000
//     strs[i] consists of lowercase English letters.

import "fmt"

func splitLoopedString(strs []string) string {
    if 0 == len(strs) {
        return ""
    }
    max := func (x, y string) string { if x > y { return x; }; return y; }
    reverseString := func(s string) string {
        bs := []byte(s)
        for i := 0; i < len(s) >> 1; i++ {
            bs[i], bs[len(s)-i-1] = bs[len(s)-i-1], bs[i]
        }
        return string(bs)
    }
    k, s, res := 0, "", "a"
    revs := make([]string, len(strs))
    for i, str := range strs {
        rev := reverseString(str)
        revs[i] = rev
        if str > rev {
            s += str
        } else {
            s += rev
        }
    }
    for i, str := range strs {
        rev, mid := revs[i], s[k+len(str):]+s[:k]
        for j := 0; j < len(str); j++ {
            if str[j] >= res[0] {
                res = max(res, str[j:] + mid + str[:j])
            }
            if rev[j] >= res[0] {
                res = max(res, rev[j:] + mid + rev[:j])
            }
        }
        k += len(str)
    }
    return res
}

func main() {
    // Example 1:
    // Input: strs = ["abc","xyz"]
    // Output: "zyxcba"
    // Explanation: You can get the looped string "-abcxyz-", "-abczyx-", "-cbaxyz-", "-cbazyx-", where '-' represents the looped status. 
    // The answer string came from the fourth looped one, where you could cut from the middle character 'a' and get "zyxcba".
    fmt.Println(splitLoopedString([]string{"abc","xyz"})) // "zyxcba"
    // Example 2:
    // Input: strs = ["abc"]
    // Output: "cba"
    fmt.Println(splitLoopedString([]string{"abc"})) // "cba"
}