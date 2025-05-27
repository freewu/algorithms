package main

// 3557. Find Maximum Number of Non Intersecting Substrings
// You are given a string word.

// Return the maximum number of non-intersecting substrings of word that are at least four characters long and start and end with the same letter.

// Example 1:
// Input: word = "abcdeafdef"
// Output: 2
// Explanation:
// The two substrings are "abcdea" and "fdef".

// Example 2:
// Input: word = "bcdaaaab"
// Output: 1
// Explanation:
// The only substring is "aaaa". Note that we cannot also choose "bcdaaaab" since it intersects with the other substring.

// Constraints:
//     1 <= word.length <= 2 * 10^5
//     word consists only of lowercase English letters.

import "fmt"
import "sort"

func maxSubstrings(word string) int {
    arr := [26][]int{}
    for i, v:= range word {
        arr[int(v - 'a')] = append(arr[int(v - 'a')], i)
    }
    dp := make([]int, len(word))
    for i := 0; i < len(dp); i++ {
        dp[i] = -1
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(index int) int 
    dfs = func(index int) (res int) {
        if index == len(word) { return 0 }
        if dp[index] > -1     { return dp[index] }
        res = dfs(index+1)
        w := int(word[index] - 'a')
        i := sort.SearchInts(arr[w], index)
        for x := i; x < len(arr[w]); x++ {
            if arr[w][x] - arr[w][i] >= 3 {
                res = max(res, dfs(arr[w][x] + 1) + 1)
                break
            }
        }
        dp[index] = res
        return res
    }
    return dfs(0)
}

func maxSubstrings1(word string) int {
    res, pos := 0, [26]int{}
    for i := range pos {
        pos[i] = -1
    }
    for i, v := range word {
        v -= 'a'
        if pos[v] < 0 {
            pos[v] = i
        } else if i - pos[v] > 2 {
            res++
            for j := range pos { // 开始找下一个子串
                pos[j] = -1
            }
        }
    }
    return res
}

func maxSubstrings2(word string) int {
    pos := [26]int{}
    res, last := 0, 1
    for i, v := range word {
        li := pos[v - 'a']
        if last > li {
            pos[v - 'a'] = i + 1
            continue
        }
        if i + 1- li >= 3 {
            last = i + 1
            res++
            continue
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: word = "abcdeafdef"
    // Output: 2
    // Explanation:
    // The two substrings are "abcdea" and "fdef".
    fmt.Println(maxSubstrings("abcdeafdef")) // 2
    // Example 2:
    // Input: word = "bcdaaaab"
    // Output: 1
    // Explanation:
    // The only substring is "aaaa". Note that we cannot also choose "bcdaaaab" since it intersects with the other substring.
    fmt.Println(maxSubstrings("bcdaaaab")) // 1

    fmt.Println(maxSubstrings("bluefrog")) // 0
    fmt.Println(maxSubstrings("leetcode")) // 1

    fmt.Println(maxSubstrings1("abcdeafdef")) // 2
    fmt.Println(maxSubstrings1("bcdaaaab")) // 1
    fmt.Println(maxSubstrings1("bluefrog")) // 0
    fmt.Println(maxSubstrings1("leetcode")) // 1
    
    fmt.Println(maxSubstrings2("abcdeafdef")) // 2
    fmt.Println(maxSubstrings2("bcdaaaab")) // 1
    fmt.Println(maxSubstrings2("bluefrog")) // 0
    fmt.Println(maxSubstrings2("leetcode")) // 1
}