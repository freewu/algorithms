package main

// 1807. Evaluate the Bracket Pairs of a String
// You are given a string s that contains some bracket pairs, with each pair containing a non-empty key.
//     For example, in the string "(name)is(age)yearsold", there are two bracket pairs that contain the keys "name" and "age".

// You know the values of a wide range of keys. 
// This is represented by a 2D string array knowledge where each knowledge[i] = [keyi, valuei] indicates that key keyi has a value of valuei.

// You are tasked to evaluate all of the bracket pairs.
// When you evaluate a bracket pair that contains some key keyi, you will:

// Replace keyi and the bracket pair with the key's corresponding valuei.
//     1. If you do not know the value of the key, you will replace keyi 
//        and the bracket pair with a question mark "?" (without the quotation marks).
//     2. Each key will appear at most once in your knowledge.
//        There will not be any nested brackets in s.

// Return the resulting string after evaluating all of the bracket pairs.

// Example 1:
// Input: s = "(name)is(age)yearsold", knowledge = [["name","bob"],["age","two"]]
// Output: "bobistwoyearsold"
// Explanation:
// The key "name" has a value of "bob", so replace "(name)" with "bob".
// The key "age" has a value of "two", so replace "(age)" with "two".

// Example 2:
// Input: s = "hi(name)", knowledge = [["a","b"]]
// Output: "hi?"
// Explanation: As you do not know the value of the key "name", replace "(name)" with "?".

// Example 3:
// Input: s = "(a)(a)(a)aaa", knowledge = [["a","yes"]]
// Output: "yesyesyesaaa"
// Explanation: The same key can appear multiple times.
// The key "a" has a value of "yes", so replace all occurrences of "(a)" with "yes".
// Notice that the "a"s not in a bracket pair are not evaluated.

// Constraints:
//     1 <= s.length <= 10^5
//     0 <= knowledge.length <= 10^5
//     knowledge[i].length == 2
//     1 <= keyi.length, valuei.length <= 10
//     s consists of lowercase English letters and round brackets '(' and ')'.
//     Every open bracket '(' in s will have a corresponding close bracket ')'.
//     The key in each bracket pair of s will be non-empty.
//     There will not be any nested bracket pairs in s.
//     keyi and valuei consist of lowercase English letters.
//     Each keyi in knowledge is unique.

import "fmt"

func evaluate(s string, knowledge [][]string) string {
    mp := make(map[string]string)
    for _, v := range knowledge {
        mp[v[0]] = v[1]
    }
    res, i, n := []byte{}, 0, len(s)
    for i < n {
        for i < n && s[i] != '(' {
            res = append(res, s[i])
            i++
        }
        if i == n { break }
        j := i + 1
        for j < n && s[j] != ')' {
            j++
        }
        if v, ok := mp[s[i+1:j]]; ok {
            for i := range v {
                res = append(res, v[i])
            }
        } else {
            res = append(res, '?')
        }
        i = j + 1
    }
    return string(res)
}

func evaluate1(s string, knowledge [][]string) string {
    mp := make(map[string]string) // Create the dictionary from knowledge pairs
    for _, v := range knowledge {
        mp[v[0]] = v[1]
    }
    res, n := []byte{}, len(s)
    for i := 0; i < n; i++ { // Iterate through the string and build the result
        if s[i] == '(' {
            j := i + 1 // Start extracting the key within brackets
            for j < n && s[j] != ')' {
                j++
            }
            key := s[i+1 : j] // Extracted key inside ( )
            if v, ok := mp[key]; ok {
                for i := range v {
                    res = append(res, v[i])
                }
            } else {
                res = append(res, '?')
            }
            i = j // Move index to the end of the closing bracket
        } else {
            res = append(res, s[i]) // Append regular characters
        }
    }
    return string(res)
}

var buf [100_000]byte

func evaluate2(s string, knowledge [][]string) string {
    mp := make(map[string]string, len(knowledge))
    for _, k := range knowledge {
        mp[k[0]] = k[1]
    }
    n, index := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            index = n
        } else if s[i] == ')' {
            val := mp[string(buf[index:n])]
            if val == "" {
                buf[index] = '?'
                n = index + 1   
            } else {
                for i := index; i < index+len(val); i++ {
                    buf[i] = val[i-index]
                }
                n = index + len(val)
            }
        } else {
            buf[n] = s[i]
            n++
        }
    }
    return string(buf[:n])
}

func main() {
    // Example 1:
    // Input: s = "(name)is(age)yearsold", knowledge = [["name","bob"],["age","two"]]
    // Output: "bobistwoyearsold"
    // Explanation:
    // The key "name" has a value of "bob", so replace "(name)" with "bob".
    // The key "age" has a value of "two", so replace "(age)" with "two".
    fmt.Println(evaluate("(name)is(age)yearsold", [][]string{{"name","bob"},{"age","two"}})) // "bobistwoyearsold"
    // Example 2:
    // Input: s = "hi(name)", knowledge = [["a","b"]]
    // Output: "hi?"
    // Explanation: As you do not know the value of the key "name", replace "(name)" with "?".
    fmt.Println(evaluate("hi(name)", [][]string{{"a","b"}})) // "hi?"
    // Example 3:
    // Input: s = "(a)(a)(a)aaa", knowledge = [["a","yes"]]
    // Output: "yesyesyesaaa"
    // Explanation: The same key can appear multiple times.
    // The key "a" has a value of "yes", so replace all occurrences of "(a)" with "yes".
    // Notice that the "a"s not in a bracket pair are not evaluated.
    fmt.Println(evaluate("(a)(a)(a)aaa", [][]string{{"a","yes"}})) // "yesyesyesaaa"

    fmt.Println(evaluate("call(bluefrog)", [][]string{{"a","b"}})) // "call?"
    fmt.Println(evaluate("use(leetcode)", [][]string{{"a","b"}})) // "use?"

    fmt.Println(evaluate1("(name)is(age)yearsold", [][]string{{"name","bob"},{"age","two"}})) // "bobistwoyearsold"
    fmt.Println(evaluate1("hi(name)", [][]string{{"a","b"}})) // "hi?"
    fmt.Println(evaluate1("(a)(a)(a)aaa", [][]string{{"a","yes"}})) // "yesyesyesaaa"
    fmt.Println(evaluate1("call(bluefrog)", [][]string{{"a","b"}})) // "call?"
    fmt.Println(evaluate1("use(leetcode)", [][]string{{"a","b"}})) // "use?"

    fmt.Println(evaluate2("(name)is(age)yearsold", [][]string{{"name","bob"},{"age","two"}})) // "bobistwoyearsold"
    fmt.Println(evaluate2("hi(name)", [][]string{{"a","b"}})) // "hi?"
    fmt.Println(evaluate2("(a)(a)(a)aaa", [][]string{{"a","yes"}})) // "yesyesyesaaa"
    fmt.Println(evaluate2("call(bluefrog)", [][]string{{"a","b"}})) // "call?"
    fmt.Println(evaluate2("use(leetcode)", [][]string{{"a","b"}})) // "use?"
}