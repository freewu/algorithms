package main

// 3213. Construct String with Minimum Cost
// You are given a string target, an array of strings words, and an integer array costs, both arrays of the same length.

// Imagine an empty string s.

// You can perform the following operation any number of times (including zero):
//     Choose an index i in the range [0, words.length - 1].
//     Append words[i] to s.
//     The cost of operation is costs[i].

// Return the minimum cost to make s equal to target. 
// If it's not possible, return -1.

// Example 1:
// Input: target = "abcdef", words = ["abdef","abc","d","def","ef"], costs = [100,1,1,10,5]
// Output: 7
// Explanation:
// The minimum cost can be achieved by performing the following operations:
// Select index 1 and append "abc" to s at a cost of 1, resulting in s = "abc".
// Select index 2 and append "d" to s at a cost of 1, resulting in s = "abcd".
// Select index 4 and append "ef" to s at a cost of 5, resulting in s = "abcdef".

// Example 2:
// Input: target = "aaaa", words = ["z","zz","zzz"], costs = [1,10,100]
// Output: -1
// Explanation:
// It is impossible to make s equal to target, so we return -1.

// Constraints:
//     1 <= target.length <= 5 * 10^4
//     1 <= words.length == costs.length <= 5 * 10^4
//     1 <= words[i].length <= target.length
//     The total sum of words[i].length is less than or equal to 5 * 10^4.
//     target and words[i] consist only of lowercase English letters.
//     1 <= costs[i] <= 10^4

import "fmt"
import "sort"

func minimumCost(target string, words []string, costs []int) int {
    // Init word cost map and word lengths set.
    mp, set := make(map[string]int, len(words)), make(map[int]bool)
    for i, word := range words {
        if mp[word] == 0 || mp[word] > costs[i] {
            mp[word] = costs[i]
        }
        set[len(word)] = true
    }
    // Convert set of lengths to a sorted list.
    arr := make([]int, 0, len(set))
    for i := range set {
        arr = append(arr, i)
    }
    sort.Ints(arr)
    // Fill DP, zeros mean infeasible. Can do it since costs are non-zero.
    dp := make([]int, len(target))
    cost := 0
    for i := range dp {
        if i != 0 {
            cost = dp[i - 1]
            if cost == 0 { continue }
        }
        for _, v := range arr {
            if i + v > len(target) {
                break
            }
            w := target[i:i + v]
            c, ok := mp[w]
            if !ok { continue  }
            j := i + v - 1
            if dp[j] == 0 || dp[j] > cost + c {
                dp[j] = cost + c
            }
        }
    }
    // feasibility check
    res := dp[len(dp) - 1]
    if res == 0 { return -1 }
    return res
}

func minimumCost1(target string, words []string, costs []int) int {
    n, mod, base := len(target), 1_070_777_777, int(9e8) // 9e8 - rand.Intn(1e8)
    powBase, preHash := make([]int, n + 1), make([]int, n + 1)
    powBase[0] = 1
    for i, c := range target {
        powBase[i+1] = powBase[i] * base % mod
        preHash[i+1] = (preHash[i] * base + int(c)) % mod
    }
    // target[l:r)
    subHash := func(l, r int) int { return ((preHash[r] - preHash[l] * powBase[r-l]) % mod + mod) % mod }
    minCost := make([]map[int]int, n + 1)
    lens := map[int]struct{}{}
    for i, word := range words {
        m := len(word)
        lens[m] = struct{}{}
        // word hash
        h := 0
        for _, c := range word {
            h = (h * base + int(c)) % mod
        }
        if minCost[m] == nil {
            minCost[m] = map[int]int{}
        }
        if minCost[m][h] == 0 {
            minCost[m][h] = costs[i]
        } else {
            minCost[m][h] = min(minCost[m][h], costs[i])
        }
    }
    sortedLens := make([]int, 0, len(lens))
    for l := range lens {
        sortedLens = append(sortedLens, l)
    }
    sort.Ints(sortedLens)
    dp := make([]int, n + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        dp[i] = 1 << 31
        for _, v := range sortedLens {
            if v > i { break }
            if cost, ok := minCost[v][subHash(i - v, i)]; ok {
                dp[i] = min(dp[i], dp[i - v] + cost)
            }
        }
    }
    if dp[n] >= 1 << 31 {
        return -1
    }
    return dp[n]
}

func main() {
    // Example 1:
    // Input: target = "abcdef", words = ["abdef","abc","d","def","ef"], costs = [100,1,1,10,5]
    // Output: 7
    // Explanation:
    // The minimum cost can be achieved by performing the following operations:
    // Select index 1 and append "abc" to s at a cost of 1, resulting in s = "abc".
    // Select index 2 and append "d" to s at a cost of 1, resulting in s = "abcd".
    // Select index 4 and append "ef" to s at a cost of 5, resulting in s = "abcdef".
    fmt.Println(minimumCost("abcdef", []string{"abdef","abc","d","def","ef"}, []int{100,1,1,10,5})) // 7
    // Example 2:
    // Input: target = "aaaa", words = ["z","zz","zzz"], costs = [1,10,100]
    // Output: -1
    // Explanation:
    // It is impossible to make s equal to target, so we return -1.
    fmt.Println(minimumCost("aaaa", []string{"z","zz","zzz"}, []int{1,10,100})) // -1

    fmt.Println(minimumCost("bluefrog", []string{"a","bb","fff"}, []int{1,10,100})) // -1
    fmt.Println(minimumCost("leetcode", []string{"a","bb","fff"}, []int{1,10,100})) // -1

    fmt.Println(minimumCost1("abcdef", []string{"abdef","abc","d","def","ef"}, []int{100,1,1,10,5})) // 7
    fmt.Println(minimumCost1("aaaa", []string{"z","zz","zzz"}, []int{1,10,100})) // -1
    fmt.Println(minimumCost1("bluefrog", []string{"a","bb","fff"}, []int{1,10,100})) // -1
    fmt.Println(minimumCost1("leetcode", []string{"a","bb","fff"}, []int{1,10,100})) // -1
}