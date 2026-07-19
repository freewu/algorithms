package main

// 3995. Minimum Cost to Convert String III
// You are given two strings, source and target.

// You are also given a 2D string array rules, where rules[i] = [patterni, replacementi], 
// and an integer array costs, where costs[i] is the base cost of applying rules[i]. 
// Both arrays have the same length. 
// Additionally, patterni and replacementi have the same length.

// You may apply any rule any number of times. Each rule application works as follows:
//     1. Choose an index l such that the range of positions from l to l + patterni.length - 1 exists in the current string and none of these positions has been used in a previous rule application.
//     2. For each index j, the character patterni[j] must either be equal to the current character at position l + j, or be '*'.
//     3. Replace the characters in this range with replacementi. The replacement is used exactly as given and does not contain wildcards.
//     4. The cost of this rule application is costs[i] plus the number of '*' characters in patterni.
//     5. Once a character position has been used in a rule application, it cannot be used in any later rule application.

// Since every patterni and replacementi have the same length, character positions are preserved after every rule application.

// Return the minimum total cost required to transform source into target. If it is impossible, return -1.

// Example 1:
// Input: source = "hello", target = "world", rules = [["he","wo"],["llo","rld"]], costs = [3,4]
// Output: 7
// Explanation:
// Apply rules[0] to replace "he" with "wo" at cost 3, so the string becomes "wollo".
// Apply rules[1] to replace "llo" with "rld" at cost 4, so the string becomes "world".
// The total cost is 3 + 4 = 7.

// Example 2:
// Input: source = "cat", target = "dog", rules = [["c*t","dog"]], costs = [2]
// Output: 3
// Explanation:
// Apply rules[0] to replace "cat" with "dog". The wildcard '*' matches 'a', adding 1 to the base cost 2.
// The total cost is 2 + 1 = 3.

// Example 3:
// Input: source = "test", target = "next", rules = [["*e*t","next"]], costs = [4]
// Output: 6
// Explanation:
// Apply rules[0] to replace "test" with "next". The first wildcard matches 't' and the second wildcard matches 's', adding 2 to the base cost 4.
// The total cost is 4 + 2 = 6.

// Example 4:
// Input: source = "ab", target = "bc", rules = [["a*","bd"]], costs = [9]
// Output: -1
// Explanation:
// No sequence of rule applications can transform source into target, so the answer is -1.

// Constraints:
//     1 <= source.length, target.length <= 5000
//     source and target consist of lowercase English letters.
//     1 <= rules.length == costs.length <= 200
//     rules[i] = [patterni, replacementi]
//     1 <= patterni.length == replacementi.length <= 20
//     patterni contains at least one lowercase English letter and at most 5 '*' characters.
//     replacementi contains only lowercase English letters.
//     1 <= costs[i] <= 1000

import "fmt"

func minCost(source string, target string, rules [][]string, costs []int) int {
    type Pair struct {
        step int
        cost int
    }
    n := len(source)
    // match 判断 pattern 是否匹配text，等长，*可匹配任意字符
    match := func(pattern, text string) bool {
        if len(pattern) != len(text) {
            return false
        }
        for idx := range pattern {
            p := pattern[idx]
            t := text[idx]
            if p != '*' && p != t {
                return false
            }
        }
        return true
    }
    // jumps[i] 保存位置i所有合法跳跃(步长,成本)
    jumps := make([][]Pair, n)
    for ruleIdx, rule := range rules {
        pat, rep := rule[0], rule[1]
        baseCost := costs[ruleIdx]
        m := len(pat)
        // 计算 pat 中*的数量
        starCnt := 0
        for _, c := range pat {
            if c == '*' {
                starCnt++
            }
        }
        jumpCost := baseCost + starCnt
        // 遍历所有起始位置i
        maxI := n - m
        if maxI < 0 {
            continue
        }
        for i := 0; i <= maxI; i++ {
            sSub := source[i : i+m]
            tSub := target[i : i+m]
            if match(pat, sSub) && match(rep, tSub) {
                jumps[i] = append(jumps[i], Pair{step: m, cost: jumpCost})
            }
        }
    }
    // dp记忆化缓存
    dp := make([]int, n + 1)
    for i := range dp {
        dp[i] = -1 // -1代表未计算
    }
    var dfs func(i int) int
    dfs = func(i int) int {
        if i == n {
            return 0
        }
        if dp[i] != -1 {
            return dp[i]
        }
        res := 1 << 61
        // 选项1：单步免费跳，字符相等才能走i->i+1
        if source[i] == target[i] {
            res = dfs(i + 1)
        }
        // 选项2：所有规则跳跃
        for _, p := range jumps[i] {
            next := i + p.step
            if next > n {
                continue
            }
            cur := dfs(next) + p.cost
            if cur < res {
                res = cur
            }
        }
        dp[i] = res
        return res
    }
    res := dfs(0)
    if res == 1 << 61 { 
        return -1
    }
    return res
}

func minCost1(source string, target string, rules [][]string, costs []int) int {
    n, inf := len(source), int64(1 << 61)
    dp := make([]int64, n + 1)
    for i := range dp {
        dp[i] = inf
    }
    dp[0] = 0
    for i := 0; i < n; i++ {
        if dp[i] == inf {
            continue
        }
        if source[i] == target[i] && dp[i] < dp[i+1] {
            dp[i+1] = dp[i]
        }
        for j, rule := range rules {
            p, r := rule[0], rule[1]
            m := len(p)
            if i + m > n {
                continue
            }
            cost := costs[j]
            ok := true
            for k := 0; k < m && ok; k++ {
                if r[k] != target[i+k] {
                    ok = false
                }
                if p[k] == '*' {
                    cost++
                } else if p[k] != source[i+k] {
                    ok = false
                }
            }
            if ok {
                v := dp[i] + int64(cost)
                if v < dp[i + m] {
                    dp[i + m] = v
                }
            }
        }
    }
    if dp[n] == inf {
        return -1
    }
    return int(dp[n])
}

func main() {
    // Example 1:
    // Input: source = "hello", target = "world", rules = [["he","wo"],["llo","rld"]], costs = [3,4]
    // Output: 7
    // Explanation:
    // Apply rules[0] to replace "he" with "wo" at cost 3, so the string becomes "wollo".
    // Apply rules[1] to replace "llo" with "rld" at cost 4, so the string becomes "world".
    // The total cost is 3 + 4 = 7.
    fmt.Println(minCost("hello", "world", [][]string{{"he", "wo"}, {"llo", "rld"}}, []int{3, 4})) // 7
    // Example 2:
    // Input: source = "cat", target = "dog", rules = [["c*t","dog"]], costs = [2]
    // Output: 3
    // Explanation:
    // Apply rules[0] to replace "cat" with "dog". The wildcard '*' matches 'a', adding 1 to the base cost 2.
    // The total cost is 2 + 1 = 3.
    fmt.Println(minCost("cat", "dog", [][]string{{"c*t", "dog"}}, []int{2})) // 3
    // Example 3:
    // Input: source = "test", target = "next", rules = [["*e*t","next"]], costs = [4]
    // Output: 6
    // Explanation:
    // Apply rules[0] to replace "test" with "next". The first wildcard matches 't' and the second wildcard matches 's', adding 2 to the base cost 4.
    // The total cost is 4 + 2 = 6.
    fmt.Println(minCost("test", "next", [][]string{{"*e*t", "next"}}, []int{4})) // 6
    // Example 4:
    // Input: source = "ab", target = "bc", rules = [["a*","bd"]], costs = [9]
    // Output: -1
    // Explanation:
    // No sequence of rule applications can transform source into target, so the answer is -1.
    fmt.Println(minCost("ab", "bc", [][]string{{"a*", "bd"}}, []int{9})) // -1

    fmt.Println(minCost1("hello", "world", [][]string{{"he", "wo"}, {"llo", "rld"}}, []int{3, 4})) // 7
    fmt.Println(minCost1("cat", "dog", [][]string{{"c*t", "dog"}}, []int{2})) // 3
    fmt.Println(minCost1("test", "next", [][]string{{"*e*t", "next"}}, []int{4})) // 6
    fmt.Println(minCost1("ab", "bc", [][]string{{"a*", "bd"}}, []int{9})) // -1
}