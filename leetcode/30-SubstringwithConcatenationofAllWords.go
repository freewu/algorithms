package main

// 30. Substring with Concatenation of All Words
// You are given a string s and an array of strings words. All the strings of words are of the same length.
// A concatenated substring in s is a substring that contains all the strings of any permutation of words concatenated.
//     For example, if words = ["ab","cd","ef"], 
//     then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. 
//     "acdbef" is not a concatenated substring because it is not the concatenation of any permutation of words.

// Return the starting indices of all the concatenated substrings in s. You can return the answer in any order.

// Example 1:
// Input: s = "barfoothefoobarman", words = ["foo","bar"]
// Output: [0,9]
// Explanation: Since words.length == 2 and words[i].length == 3, the concatenated substring has to be of length 6.
// The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
// The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.
// The output order does not matter. Returning [9,0] is fine too.

// Example 2:
// Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
// Output: []
// Explanation: Since words.length == 4 and words[i].length == 4, the concatenated substring has to be of length 16.
// There is no substring of length 16 in s that is equal to the concatenation of any permutation of words.
// We return an empty array.

// Example 3:
// Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
// Output: [6,9,12]
// Explanation: Since words.length == 3 and words[i].length == 3, the concatenated substring has to be of length 9.
// The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"] which is a permutation of words.
// The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"] which is a permutation of words.
// The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"] which is a permutation of words.

// Constraints:
//     1 <= s.length <= 10^4
//     1 <= words.length <= 5000
//     1 <= words[i].length <= 30
//     s and words[i] consist of lowercase English letters.

import "fmt"

func findSubstring(s string, words []string) []int {
    if len(words) == 0 {
        return []int{}
    }
    res := []int{}
    counter := map[string]int{}
    for _, w := range words {
        counter[w]++
    }
    checkWords := func(s map[string]int) bool {
        flag := true
        for _, v := range s {
            if v > 0 {
                flag = false
                break
            }
        }
        return flag
    }
    copyMap := func (s map[string]int) map[string]int {
        c := map[string]int{}
        for k, v := range s {
            c[k] = v
        }
        return c
    }
    length, totalLen, tmpCounter := len(words[0]), len(words[0])*len(words), copyMap(counter)
    for i, start := 0, 0; i < len(s)-length+1 && start < len(s)-length+1; i++ {
        //fmt.Printf("sub = %v i = %v lenght = %v start = %v tmpCounter = %v totalLen = %v\n", s[i:i+length], i, length, start, tmpCounter, totalLen)
        if tmpCounter[s[i:i+length]] > 0 {
            tmpCounter[s[i:i+length]]--
            //fmt.Printf("******sub = %v i = %v lenght = %v start = %v tmpCounter = %v totalLen = %v\n", s[i:i+length], i, length, start, tmpCounter, totalLen)
            if checkWords(tmpCounter) && (i+length-start == totalLen) {
                res = append(res, start)
                continue
            }
            i = i + length - 1
        } else {
            start++
            i = start - 1
            tmpCounter = copyMap(counter)
        }
    }
    return res
}

// best solution
func findSubstring1(s string, words []string) []int {
    m, cm := map[string]int{}, map[string]int{}
    for _, word := range words {
        if _, ok := m[word]; !ok {
            m[word] = 0
        }
        m[word] += 1
    }
    wl := len(words[0])
    str, tmp, res := "", "",[]int{}
    clearMap := func (m map[string]int) { for k := range m { delete(m, k); }; }
    for i := 0; i < wl; i++ {
        count := 0
        start := i
        for end := i; end + wl <= len(s); end += wl {
            str = s[end:end + wl]
            if _, ok := m[str]; ok {
                if _, ok := cm[str]; ok {
                    cm[str] += 1
                } else {
                    cm[str] = 1
                }
                if m[str] >= cm[str] { count++ }

                for cm[str] > m[str] {
                    tmp = s[start:end + wl]
                    cm[tmp] -= 1
                    start += wl
                    if m[tmp] > cm[tmp] { count-- }
                }
                if count == len(words) {
                    res = append(res, start)
                    tmp = s[start:end + wl]
                    cm[tmp] -= 1
                    count --
                    start += wl
                }
            } else {
                clearMap(cm)
                count = 0
                start = end + wl
            }
        }
        clearMap(cm)
    }
    return res
}

func findSubstring2(s string, words []string) []int {
    res := make([]int, 0) // Initialize the answer as a slice
    cnt, dif := make(map[string]int), 0 // Create the hashmap, also count the different elements in words
        for _, curString := range words {
            cnt[curString]++;
            if cnt[curString] == 1 {
                dif++
            }
        }
    sz, totSz := len(words[0]), len(words[0]) * len(words) // Initialize some length for sliding window
    for st := 0; st < sz; st++ { // Offset
        for k := range cnt { // Reset the map
            delete(cnt, k)
        }
        for _, curString := range words {
            cnt[curString]++;
        }        
        
        cntOk := 0 // Number of key in cnt has the value 0
        for i := st; i + sz - 1 < len(s); i += sz {
            if i >= totSz { // Remove the left part
                cur := s[i-totSz:i-totSz+sz]
                if _, ok := cnt[cur]; ok {
                    cnt[cur]++
                    if cnt[cur] == 1 {
                        cntOk--
                    }
                }
            }
            cur := s[i:i+sz]; // Insert the right part
            if _, ok := cnt[cur]; ok {
                cnt[cur]--
                if cnt[cur] == 0 {
                    cntOk++
                }
            }
            if cntOk == dif { // If everything is good
                res = append(res, i - totSz + sz) // Then add the current range into the answer
            }
        }
    }
    return res
}

func main() {
    // Explanation: Since words.length == 2 and words[i].length == 3, the concatenated substring has to be of length 6.
    // The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
    // The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.
    // The output order does not matter. Returning [9,0] is fine too.
    fmt.Printf("findSubstring(\"barfoothefoobarman\", []string{\"foo\", \"bar\"}) = %v\n", findSubstring("barfoothefoobarman", []string{"foo", "bar"})) // [0,9]

    // Explanation: Since words.length == 4 and words[i].length == 4, the concatenated substring has to be of length 16.
    // There is no substring of length 16 in s that is equal to the concatenation of any permutation of words.
    // We return an empty array.
    fmt.Printf("findSubstring(\"wordgoodgoodgoodbestword\", []string{\"word\", \"good\", \"best\", \"word\"}) = %v\n", findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})) // []
        
    // Explanation: Since words.length == 3 and words[i].length == 3, the concatenated substring has to be of length 9.
    // The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"] which is a permutation of words.
    // The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"] which is a permutation of words.
    // The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"] which is a permutation of words.
    fmt.Printf("findSubstring(\"barfoofoobarthefoobarman\", []string{\"bar\", \"foo\",\"the\"}) = %v\n", findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})) // [6,9,12]

    fmt.Printf("findSubstring1(\"barfoothefoobarman\", []string{\"foo\", \"bar\"}) = %v\n", findSubstring1("barfoothefoobarman", []string{"foo", "bar"})) // [0,9]
    fmt.Printf("findSubstring1(\"wordgoodgoodgoodbestword\", []string{\"word\", \"good\", \"best\", \"word\"}) = %v\n", findSubstring1("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})) // []
    fmt.Printf("findSubstring1(\"barfoofoobarthefoobarman\", []string{\"bar\", \"foo\",\"the\"}) = %v\n", findSubstring1("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})) // [6,9,12]

    fmt.Printf("findSubstring2(\"barfoothefoobarman\", []string{\"foo\", \"bar\"}) = %v\n", findSubstring2("barfoothefoobarman", []string{"foo", "bar"})) // [0,9]
    fmt.Printf("findSubstring2(\"wordgoodgoodgoodbestword\", []string{\"word\", \"good\", \"best\", \"word\"}) = %v\n", findSubstring2("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})) // []
    fmt.Printf("findSubstring2(\"barfoofoobarthefoobarman\", []string{\"bar\", \"foo\",\"the\"}) = %v\n", findSubstring2("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})) // [6,9,12]
}
