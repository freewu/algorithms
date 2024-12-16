package main

// 2213. Longest Substring of One Repeating Character
// You are given a 0-indexed string s. 
// You are also given a 0-indexed string queryCharacters of length k and a 0-indexed array of integer indices queryIndices of length k, 
// both of which are used to describe k queries.

// The ith query updates the character in s at index queryIndices[i] to the character queryCharacters[i].

// Return an array lengths of length k where lengths[i] is the length of the longest substring of s consisting of only one repeating character after the ith query is performed.

// Example 1:
// Input: s = "babacc", queryCharacters = "bcb", queryIndices = [1,3,3]
// Output: [3,3,4]
// Explanation: 
// - 1st query updates s = "bbbacc". The longest substring consisting of one repeating character is "bbb" with length 3.
// - 2nd query updates s = "bbbccc". 
//   The longest substring consisting of one repeating character can be "bbb" or "ccc" with length 3.
// - 3rd query updates s = "bbbbcc". The longest substring consisting of one repeating character is "bbbb" with length 4.
// Thus, we return [3,3,4].

// Example 2:
// Input: s = "abyzz", queryCharacters = "aa", queryIndices = [2,1]
// Output: [2,3]
// Explanation:
// - 1st query updates s = "abazz". The longest substring consisting of one repeating character is "zz" with length 2.
// - 2nd query updates s = "aaazz". The longest substring consisting of one repeating character is "aaa" with length 3.
// Thus, we return [2,3].

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters.
//     k == queryCharacters.length == queryIndices.length
//     1 <= k <= 10^5
//     queryCharacters consists of lowercase English letters.
//     0 <= queryIndices[i] < s.length

import "fmt"

type SegmentTree struct {
    str []byte
    mx  []int
    lmx []int
    rmx []int
}

func NewSegmentTree(s string) *SegmentTree {
    n := len(s)
    t := &SegmentTree{ str: []byte(s), mx:  make([]int, n<<2), lmx: make([]int, n<<2), rmx: make([]int, n<<2), }
    t.build(0, 0, n - 1)
    return t
}

func (t *SegmentTree) build(x, l, r int) {
    if l == r {
        t.lmx[x] = 1
        t.rmx[x] = 1
        t.mx[x] = 1
        return
    }
    m := int(uint(l + r) >> 1)
    t.build(x * 2 + 1, l, m)
    t.build(x * 2 + 2, m + 1, r)
    t.pushup(x, l, m, r)
}

func (t *SegmentTree) pushup(x, l, m, r int) {
    lch, rch := x*2+1, x*2+2
    t.lmx[x] = t.lmx[lch]
    t.rmx[x] = t.rmx[rch]
    t.mx[x] = max(t.mx[lch], t.mx[rch])
    if t.str[m] == t.str[m + 1] { // can be merged
        if t.lmx[lch] == m - l + 1 {
            t.lmx[x] += t.lmx[rch]
        }
        if t.rmx[rch] == r - m {
            t.rmx[x] += t.rmx[lch]
        }
        t.mx[x] = max(t.mx[x], t.rmx[lch] + t.lmx[rch])
    }
}

func (t *SegmentTree) update(x, l, r, pos int, val byte) {
    if l == r {
        t.str[pos] = val
        return
    }
    m := int(uint(l+r) >> 1)
    if pos <= m {
        t.update(x * 2 + 1, l, m, pos, val)
    } else {
        t.update(x * 2 + 2, m + 1, r, pos, val)
    }
    t.pushup(x, l, m, r)
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
    res, n := make([]int, len(queryCharacters)), len(s)
    t := NewSegmentTree(s)
    for i, c := range queryCharacters {
        t.update(0, 0, n-1, queryIndices[i], byte(c))
        res[i] = t.mx[0]
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "babacc", queryCharacters = "bcb", queryIndices = [1,3,3]
    // Output: [3,3,4]
    // Explanation: 
    // - 1st query updates s = "bbbacc". The longest substring consisting of one repeating character is "bbb" with length 3.
    // - 2nd query updates s = "bbbccc". 
    //   The longest substring consisting of one repeating character can be "bbb" or "ccc" with length 3.
    // - 3rd query updates s = "bbbbcc". The longest substring consisting of one repeating character is "bbbb" with length 4.
    // Thus, we return [3,3,4].
    fmt.Println(longestRepeating("babacc", "bcb", []int{1,3,3})) // [3,3,4]
    // Example 2:
    // Input: s = "abyzz", queryCharacters = "aa", queryIndices = [2,1]
    // Output: [2,3]
    // Explanation:
    // - 1st query updates s = "abazz". The longest substring consisting of one repeating character is "zz" with length 2.
    // - 2nd query updates s = "aaazz". The longest substring consisting of one repeating character is "aaa" with length 3.
    // Thus, we return [2,3].
    fmt.Println(longestRepeating("abyzz", "aa", []int{2,1})) // [2,3]
}