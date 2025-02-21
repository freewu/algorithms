package main

// 3458. Select K Disjoint Special Substrings
// Given a string s of length n and an integer k, determine whether it is possible to select k disjoint special substrings.

// A special substring is a substring where:
//     Any character present inside the substring should not appear outside it in the string.
//     The substring is not the entire string s.

// Note that all k substrings must be disjoint, meaning they cannot overlap.

// Return true if it is possible to select k such disjoint special substrings; otherwise, return false.

// Example 1:
// Input: s = "abcdbaefab", k = 2
// Output: true
// Explanation:
// We can select two disjoint special substrings: "cd" and "ef".
// "cd" contains the characters 'c' and 'd', which do not appear elsewhere in s.
// "ef" contains the characters 'e' and 'f', which do not appear elsewhere in s.

// Example 2:
// Input: s = "cdefdc", k = 3
// Output: false
// Explanation:
// There can be at most 2 disjoint special substrings: "e" and "f". Since k = 3, the output is false.

// Example 3:
// Input: s = "abeabe", k = 0
// Output: true

// Constraints:
//     2 <= n == s.length <= 5 * 10^4
//     0 <= k <= 26
//     s consists only of lowercase English letters.

import "fmt"
import "sort"

func maxSubstringLength(s string, k int) bool {
    n := len(s)
    if k == 0 { return true }
    first, last := make([]int, 26), make([]int, 26)
    for i := range first {
        first[i], last[i] = n, -1
    }
    // Step 1: Find first and last occurrence of each character
    for i := 0; i < n; i++ {
        index := s[i] - 'a'
        if first[index] > i {
            first[index] = i
        }
        if last[index] < i {
            last[index] = i
        }
    }
    // Step 2: Find valid segments
    segments := [][]int{}
    for i := 0; i < n; i++ {
        if i != first[s[i] - 'a'] { continue }
        mx, cur, flag := last[s[i] - 'a'], i,  true
        for cur <= mx {
            if first[s[cur] - 'a'] < i {
                flag = false
                break
            }
            if last[s[cur] - 'a'] > mx {
                mx = last[s[cur] - 'a']
            }
            cur++
        }
        if flag && !(i == 0 && mx == n - 1) {
            segments = append(segments, []int{i, mx})
        }
    }
    // Step 3: Sort the segments based on ending index
    sort.Slice(segments, func(i, j int) bool {
        return segments[i][1] < segments[j][1]
    })
    // Step 4: Count non-overlapping segments
    count, index := 0, -1
    for _, v := range segments {
        if v[0] > index {
            count++
            index = v[1]
        }
    }
    return count >= k
}

func maxSubstringLength1(s string, k int) bool {
    if k == 0 { return true }
    pos := [26][2]int{}
    for i := 0; i < 26; i++ {
        pos[i] = [2]int{ -1, -1 }
    }
    for i, v := range s {
        if pos[v - 'a'][0] == -1 {
            pos[v - 'a'][0] = i
        }
        if pos[v - 'a'][1] < i {
            pos[v - 'a'][1] = i
        }
    }
    intervals := [][2]int{}
    for left, v := range s {
        if pos[v - 'a'][0] < left { continue }
        right, flag := pos[v - 'a'][1], true
        for i := left; i < right; i++ {
            if pos[s[i] - 'a'][0] < left {
                flag = false
                break
            }
            right = max(right, pos[s[i] - 'a'][1])
        }
        if !flag { continue }
        if left == 0 && right == (len(s) -  1) { continue }
        intervals = append(intervals, [2]int{ left, right })
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    count, pre := 0, -1
    for _, p := range intervals {
        if p[0] > pre {
            count++
            pre = p[1]
        }
    }
    return count >= k
}

func main() {
    // Example 1:
    // Input: s = "abcdbaefab", k = 2
    // Output: true
    // Explanation:
    // We can select two disjoint special substrings: "cd" and "ef".
    // "cd" contains the characters 'c' and 'd', which do not appear elsewhere in s.
    // "ef" contains the characters 'e' and 'f', which do not appear elsewhere in s.
    fmt.Println(maxSubstringLength("abcdbaefab", 2)) // true
    // Example 2:
    // Input: s = "cdefdc", k = 3
    // Output: false
    // Explanation:
    // There can be at most 2 disjoint special substrings: "e" and "f". Since k = 3, the output is false.
    fmt.Println(maxSubstringLength("cdefdc", 3)) // false
    // Example 3:
    // Input: s = "abeabe", k = 0
    // Output: true
    fmt.Println(maxSubstringLength("abeabe", 0)) // true

    fmt.Println(maxSubstringLength("bluefrog", 2)) // true
    fmt.Println(maxSubstringLength("leetcode", 2)) // true

    fmt.Println(maxSubstringLength1("abcdbaefab", 2)) // true
    fmt.Println(maxSubstringLength1("cdefdc", 3)) // false
    fmt.Println(maxSubstringLength1("abeabe", 0)) // true
    fmt.Println(maxSubstringLength1("bluefrog", 2)) // true
    fmt.Println(maxSubstringLength1("leetcode", 2)) // true
}

