package main

// 966. Vowel Spellchecker
// Given a wordlist, we want to implement a spellchecker that converts a query word into a correct word.

// For a given query word, the spell checker handles two categories of spelling mistakes:
//     Capitalization: If the query matches a word in the wordlist (case-insensitive), 
//     then the query word is returned with the same case as the case in the wordlist.
//         Example: wordlist = ["yellow"], query = "YellOw": correct = "yellow"
//         Example: wordlist = ["Yellow"], query = "yellow": correct = "Yellow"
//         Example: wordlist = ["yellow"], query = "yellow": correct = "yellow"
//     Vowel Errors: If after replacing the vowels ('a', 'e', 'i', 'o', 'u') of the query word with any vowel individually, 
//     it matches a word in the wordlist (case-insensitive), then the query word is returned with the same case as the match in the wordlist.
//         Example: wordlist = ["YellOw"], query = "yollow": correct = "YellOw"
//         Example: wordlist = ["YellOw"], query = "yeellow": correct = "" (no match)
//         Example: wordlist = ["YellOw"], query = "yllw": correct = "" (no match)

// In addition, the spell checker operates under the following precedence rules:
//     When the query exactly matches a word in the wordlist (case-sensitive), you should return the same word back.
//     When the query matches a word up to capitlization, you should return the first such match in the wordlist.
//     When the query matches a word up to vowel errors, you should return the first such match in the wordlist.
//     If the query has no matches in the wordlist, you should return the empty string.

// Given some queries, return a list of words answer, where answer[i] is the correct word for query = queries[i].

// Example 1:
// Input: wordlist = ["KiTe","kite","hare","Hare"], queries = ["kite","Kite","KiTe","Hare","HARE","Hear","hear","keti","keet","keto"]
// Output: ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]

// Example 2:
// Input: wordlist = ["yellow"], queries = ["YellOw"]
// Output: ["yellow"]
 
// Constraints:
//     1 <= wordlist.length, queries.length <= 5000
//     1 <= wordlist[i].length, queries[i].length <= 7
//     wordlist[i] and queries[i] consist only of only English letters.

import "fmt"
import "strings"

func spellchecker(wordlist []string, queries []string) []string {
    if len(wordlist) == 0 || len(queries) == 0 {
        return []string{}
    }
    res := make([]string, len(queries))
    capital, vowel, words :=  make(map[string]string), make(map[string]string), make(map[string]bool)
    for _, w := range wordlist {
        words[w] = true
    }
    for _, w := range wordlist {
        lower := strings.ToLower(w)
        stripVowel := lower
        for _, c := range []string{"a", "e", "i", "o", "u"} { // 过滤掉元音字母
            stripVowel = strings.ReplaceAll(stripVowel, c, "#")
        }
        if _, ok := capital[lower]; !ok {
            capital[lower] = w
        }
        if _, ok := vowel[stripVowel]; !ok {
            vowel[stripVowel] = w
        }
    }
    for i := 0; i < len(queries); i++ {
        if _, ok := words[queries[i]]; ok {
            res[i] = queries[i]
            continue
        }
        lower := strings.ToLower(queries[i])
        stripVowel := lower
        for _, c := range []string{"a", "e", "i", "o", "u"} {
            stripVowel = strings.ReplaceAll(stripVowel, c, "#")
        }
        if _, ok := capital[lower]; ok {
            res[i] = capital[lower]
        } else if _, ok := vowel[stripVowel]; ok {
            res[i] = vowel[stripVowel]
        } else {
            res[i] = ""
        }
    }
    return res
}

func spellchecker1(wordlist []string, queries []string) []string {
    res, set, mp := make([]string, len(queries)), make(map[string]struct{}), make(map[string]string)
    tolower := func (w string) string {
        s := []byte{}
        for _, a := range w {
            if a >= 'A' && a <= 'Z' {
                s = append(s, byte(a - 'A' + 'a'))
            } else {
                s = append(s, byte(a))
            }
        }
        return string(s)
    }
    vowel := func(w string) string { // 去除元音字母并小写
        s := []byte{}
        for _, a := range w {
            if a == 'a' || a == 'A' || a == 'e' || a == 'E' || a == 'i' || a == 'I' || a == 'o' || a == 'O' || a == 'u' || a == 'U' {
                s = append(s , '#')
                continue
            }
            if a >= 'A' && a <= 'Z' {
                s = append(s, byte(a - 'A' + 'a'))
            } else {
                s = append(s, byte(a))
            }
        }
        return string(s)
    }
    for _, word := range wordlist {
        set[word] = struct{}{}
        w := tolower(word)
        if _, ok := mp[w]; !ok {
            mp[w] = word
        }
        v := vowel(word)
        if _, ok := mp[v]; !ok {
            mp[v] = word
        }
    }
    for i, q := range queries {
        if _, ok := set[q]; ok {
            res[i] = q
            continue
        }
        vs, ok := mp[tolower(q)]
        if ok {
            res[i] = vs
            continue
        }
        vs, ok = mp[vowel(q)]
        if ok {
            res[i] = vs
            continue
        }
        res[i] = ""
    }
    return res
}

func spellchecker2(wordlist []string, queries []string) []string {
    // isVowel is a helper to check if a rune is a vowel.
    isVowel := func(r rune) bool {
        return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
    }
    devowel := func(word string) string {
        lower := strings.ToLower(word)
        var sb strings.Builder
        sb.Grow(len(lower))
        for _, r := range lower {
            if isVowel(r) {
                sb.WriteRune('*')
            } else {
                sb.WriteRune(r)
            }
        }
        return sb.String()
    }
    // exact match
    exact := map[string]struct{}{}
    for _, w := range wordlist {
        exact[w] = struct{}{}
    }
    // case insensitive match
    ci := map[string]string{}
    for _, w := range wordlist {
        k := strings.ToLower(w)
        if _, ok := ci[k]; !ok {
            ci[k] = w
        }
    }
    // devoweld
    d := map[string]string{}
    for _, w := range wordlist {
        k := devowel(w)
        if _, ok := d[k]; !ok {
            d[k] = w
        }
    }
    r := make([]string, len(queries))
    for i, q := range queries {
        if _, ok := exact[q]; ok {
            r[i] = q
            continue
        }
        if orig, ok := ci[strings.ToLower(q)]; ok {
            r[i] = orig
            continue
        }
        if orig, ok := d[devowel(q)]; ok {
            r[i] = orig
            continue
        }
        r[i] = ""
    }
    return r
}


func spellchecker3(wordlist []string, queries []string) []string {
    n := len(wordlist)
    m1, m2, m3 := make(map[string]string, n),make(map[string]string, n), make(map[string]string, n)
    f := func(str string) (b, c string) {
        t2 := []byte(str)
        for j := range len(t2) {
            if t2[j] >= 'a' {
                t2[j] -= 'a' - 'A'
            }
        }
        b = string(t2)
        var j2 int
        for j := range len(t2) {
            switch t2[j] {
            case 'E', 'I', 'O', 'U':
                t2[j2] = 'A'
            default:
                t2[j2] = t2[j]
            }
            j2++
        }
        c = string(t2)
        return
    }
    for i := range n {
        t := wordlist[n-1-i]
        m1[t] = t
        b, c := f(t)
        m2[b] = t
        m3[c] = t
    }
    res := make([]string, len(queries))
    for i, t := range queries {
        if v, ok := m1[t]; ok {
            res[i] = v
        } else {
            b, c := f(t)
            if v, ok := m2[b]; ok {
                res[i] = v
            } else if v, ok := m3[c]; ok {
                res[i] = v
            } else {
                res[i] = v
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: wordlist = ["KiTe","kite","hare","Hare"], queries = ["kite","Kite","KiTe","Hare","HARE","Hear","hear","keti","keet","keto"]
    // Output: ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]
    fmt.Println(spellchecker([]string{"KiTe","kite","hare","Hare"},[]string{"kite","Kite","KiTe","Hare","HARE","Hear","hear","keti","keet","keto"})) // ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]
    // Example 2:
    // Input: wordlist = ["yellow"], queries = ["YellOw"]
    // Output: ["yellow"]
    fmt.Println(spellchecker([]string{"yellow"},[]string{"YellOw"})) // ["yellow"]

    fmt.Println(spellchecker([]string{"bluefrog"},[]string{"leetcode"})) // []

    fmt.Println(spellchecker1([]string{"KiTe","kite","hare","Hare"},[]string{"kite","Kite","KiTe","Hare","HARE","Hear","hear","keti","keet","keto"})) // ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]
    fmt.Println(spellchecker1([]string{"yellow"},[]string{"YellOw"})) // ["yellow"]
    fmt.Println(spellchecker1([]string{"bluefrog"},[]string{"leetcode"})) // []

    fmt.Println(spellchecker2([]string{"KiTe","kite","hare","Hare"},[]string{"kite","Kite","KiTe","Hare","HARE","Hear","hear","keti","keet","keto"})) // ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]
    fmt.Println(spellchecker2([]string{"yellow"},[]string{"YellOw"})) // ["yellow"]
    fmt.Println(spellchecker2([]string{"bluefrog"},[]string{"leetcode"})) // []

    fmt.Println(spellchecker3([]string{"KiTe","kite","hare","Hare"},[]string{"kite","Kite","KiTe","Hare","HARE","Hear","hear","keti","keet","keto"})) // ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]
    fmt.Println(spellchecker3([]string{"yellow"},[]string{"YellOw"})) // ["yellow"]
    fmt.Println(spellchecker3([]string{"bluefrog"},[]string{"leetcode"})) // []
}