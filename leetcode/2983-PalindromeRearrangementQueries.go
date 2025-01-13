package main

// 2983. Palindrome Rearrangement Queries
// You are given a 0-indexed string s having an even length n.

// You are also given a 0-indexed 2D integer array, queries, where queries[i] = [ai, bi, ci, di].

// For each query i, you are allowed to perform the following operations:
//     1. Rearrange the characters within the substring s[ai:bi], where 0 <= ai <= bi < n / 2.
//     2. Rearrange the characters within the substring s[ci:di], where n / 2 <= ci <= di < n.

// For each query, your task is to determine whether it is possible to make s a palindrome by performing the operations.

// Each query is answered independently of the others.

// Return a 0-indexed array answer, where answer[i] == true if it is possible to make s a palindrome by performing operations specified by the ith query, and false otherwise.
//     1. A substring is a contiguous sequence of characters within a string.
//     2. s[x:y] represents the substring consisting of characters from the index x to index y in s, both inclusive.

// Example 1:
// Input: s = "abcabc", queries = [[1,1,3,5],[0,2,5,5]]
// Output: [true,true]
// Explanation: In this example, there are two queries:
// In the first query:
// - a0 = 1, b0 = 1, c0 = 3, d0 = 5.
// - So, you are allowed to rearrange s[1:1] => abcabc and s[3:5] => abcabc.
// - To make s a palindrome, s[3:5] can be rearranged to become => abccba.
// - Now, s is a palindrome. So, answer[0] = true.
// In the second query:
// - a1 = 0, b1 = 2, c1 = 5, d1 = 5.
// - So, you are allowed to rearrange s[0:2] => abcabc and s[5:5] => abcabc.
// - To make s a palindrome, s[0:2] can be rearranged to become => cbaabc.
// - Now, s is a palindrome. So, answer[1] = true.

// Example 2:
// Input: s = "abbcdecbba", queries = [[0,2,7,9]]
// Output: [false]
// Explanation: In this example, there is only one query.
// a0 = 0, b0 = 2, c0 = 7, d0 = 9.
// So, you are allowed to rearrange s[0:2] => abbcdecbba and s[7:9] => abbcdecbba.
// It is not possible to make s a palindrome by rearranging these substrings because s[3:6] is not a palindrome.
// So, answer[0] = false.

// Example 3:
// Input: s = "acbcab", queries = [[1,2,4,5]]
// Output: [true]
// Explanation: In this example, there is only one query.
// a0 = 1, b0 = 2, c0 = 4, d0 = 5.
// So, you are allowed to rearrange s[1:2] => acbcab and s[4:5] => acbcab.
// To make s a palindrome s[1:2] can be rearranged to become abccab.
// Then, s[4:5] can be rearranged to become abccba.
// Now, s is a palindrome. So, answer[0] = true.

// Constraints:
//     2 <= n == s.length <= 10^5
//     1 <= queries.length <= 10^5
//     queries[i].length == 4
//     ai == queries[i][0], bi == queries[i][1]
//     ci == queries[i][2], di == queries[i][3]
//     0 <= ai <= bi < n / 2
//     n / 2 <= ci <= di < n 
//     n is even.
//     s consists of only lowercase English letters.

import "fmt"
import "slices"

func canMakePalindromeQueries(s string, queries [][]int) []bool {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    count := func(pre [][]int, i, j int) []int {
        count := make([]int, 26)
        for k := 0; k < 26; k++ {
            count[k] = pre[j+1][k] - pre[i][k]
        }
        return count
    }
    sub := func(cnt1, cnt2 []int) []int {
        cnt := make([]int, 26)
        for i := 0; i < 26; i++ {
            cnt[i] = cnt1[i] - cnt2[i]
            if cnt[i] < 0 { return []int{} }
        }
        return cnt
    }
    check := func(pre1, pre2 [][]int, diff []int, a, b, c, d int) bool {
        if diff[a] > 0 || diff[len(diff)-1]-diff[max(b, d)+1] > 0 { return false }
        if d <= b { return slices.Equal(count(pre1, a, b), count(pre2, a, b)) }
        if b < c { return diff[c]-diff[b+1] == 0 && slices.Equal(count(pre1, a, b), count(pre2, a, b)) && slices.Equal(count(pre1, c, d), count(pre2, c, d)) }
        cnt1, cnt2 := sub(count(pre1, a, b), count(pre2, a, c-1)), sub(count(pre2, c, d), count(pre1, b+1, d))
        return !slices.Equal(cnt1, []int{}) && !slices.Equal(cnt2, []int{}) && slices.Equal(cnt1, cnt2)
    }
    reverse := func(s string) string {
        arr := []byte(s)
        for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
            arr[i], arr[j] = arr[j], arr[i]
        }
        return string(arr)
    }
    n := len(s)
    m := n / 2
    t := reverse(s[m:])
    s = s[:m]
    res, pre1, pre2, diff  := []bool{}, make([][]int, m + 1), make([][]int, m + 1), make([]int, m + 1)
    pre1[0], pre2[0] = make([]int, 26), make([]int, 26)
    for i := 1; i <= m; i++ {
        pre1[i] = slices.Clone(pre1[i-1])
        pre2[i] = slices.Clone(pre2[i-1])
        pre1[i][int(s[i-1]-'a')]++
        pre2[i][int(t[i-1]-'a')]++
        diff[i] = diff[i-1]
        if s[i-1] != t[i-1] {
            diff[i]++
        }
    }
    for _, q := range queries {
        a, b, c, d := q[0], q[1],  (n - q[3] - 1), (n - q[2] - 1)
        if a <= c {
            res = append(res, check(pre1, pre2, diff, a, b, c, d))
        } else {
            res = append(res, check(pre2, pre1, diff, c, d, a, b))
        }
    }
    return res
}

func canMakePalindromeQueries1(s string, queries [][]int) []bool {
    // 分成左右两半，右半反转
    n := len(s) / 2
    t := []byte(s[n:])
    slices.Reverse(t)
    s = s[:n]
    // 预处理三种前缀和
    sumS := make([][26]int, n+1)
    for i, b := range s {
        sumS[i+1] = sumS[i]
        sumS[i+1][b-'a']++
    }
    sumT := make([][26]int, n+1)
    for i, b := range t {
        sumT[i+1] = sumT[i]
        sumT[i+1][b-'a']++
    }
    sumNe := make([]int, n+1)
    for i := range s {
        sumNe[i+1] = sumNe[i]
        if s[i] != t[i] {
            sumNe[i+1]++
        }
    }
    // 计算子串中各个字符的出现次数，闭区间 [l,r]
    count := func(sum [][26]int, l, r int) []int {
        res := sum[r+1]
        for i, s := range sum[l][:] {
            res[i] -= s
        }
        return res[:]
    }
    subtract := func(s1, s2 []int) []int {
        for i, s := range s2 {
            s1[i] -= s
            if s1[i] < 0 {
                return nil
            }
        }
        return s1
    }
    check := func(l1, r1, l2, r2 int, sumS, sumT [][26]int) bool {
        if sumNe[l1] > 0 || // [0,l1-1] 有 s[i] != t[i]
            sumNe[n]-sumNe[max(r1, r2)+1] > 0 { // [max(r1,r2)+1,n-1] 有 s[i] != t[i]
            return false
        }
        if r2 <= r1 { // 区间包含
            return slices.Equal(count(sumS, l1, r1), count(sumT, l1, r1))
        }
        if r1 < l2 { // 区间不相交
            return sumNe[l2]-sumNe[r1+1] == 0 && // [r1+1,l2-1] 都满足 s[i] == t[i]
                slices.Equal(count(sumS, l1, r1), count(sumT, l1, r1)) &&
                slices.Equal(count(sumS, l2, r2), count(sumT, l2, r2))
        }
        // 区间相交但不包含
        s1 := subtract(count(sumS, l1, r1), count(sumT, l1, l2-1))
        s2 := subtract(count(sumT, l2, r2), count(sumS, r1+1, r2))
        return s1 != nil && s2 != nil && slices.Equal(s1, s2)
    }
    res := make([]bool, len(queries))
    for i, q := range queries {
        l1, r1, l2, r2 := q[0], q[1], n*2-1-q[3], n*2-1-q[2]
        if l1 <= l2 {
            res[i] = check(l1, r1, l2, r2, sumS, sumT)
        } else {
            res[i] = check(l2, r2, l1, r1, sumT, sumS)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abcabc", queries = [[1,1,3,5],[0,2,5,5]]
    // Output: [true,true]
    // Explanation: In this example, there are two queries:
    // In the first query:
    // - a0 = 1, b0 = 1, c0 = 3, d0 = 5.
    // - So, you are allowed to rearrange s[1:1] => abcabc and s[3:5] => abcabc.
    // - To make s a palindrome, s[3:5] can be rearranged to become => abccba.
    // - Now, s is a palindrome. So, answer[0] = true.
    // In the second query:
    // - a1 = 0, b1 = 2, c1 = 5, d1 = 5.
    // - So, you are allowed to rearrange s[0:2] => abcabc and s[5:5] => abcabc.
    // - To make s a palindrome, s[0:2] can be rearranged to become => cbaabc.
    // - Now, s is a palindrome. So, answer[1] = true.
    fmt.Println(canMakePalindromeQueries("abcabc", [][]int{{1,1,3,5},{0,2,5,5}})) // [true,true]
    // Example 2:
    // Input: s = "abbcdecbba", queries = [[0,2,7,9]]
    // Output: [false]
    // Explanation: In this example, there is only one query.
    // a0 = 0, b0 = 2, c0 = 7, d0 = 9.
    // So, you are allowed to rearrange s[0:2] => abbcdecbba and s[7:9] => abbcdecbba.
    // It is not possible to make s a palindrome by rearranging these substrings because s[3:6] is not a palindrome.
    // So, answer[0] = false.
    fmt.Println(canMakePalindromeQueries("abbcdecbba", [][]int{{0,2,7,9}})) // [false]
    // Example 3:
    // Input: s = "acbcab", queries = [[1,2,4,5]]
    // Output: [true]
    // Explanation: In this example, there is only one query.
    // a0 = 1, b0 = 2, c0 = 4, d0 = 5.
    // So, you are allowed to rearrange s[1:2] => acbcab and s[4:5] => acbcab.
    // To make s a palindrome s[1:2] can be rearranged to become abccab.
    // Then, s[4:5] can be rearranged to become abccba.
    // Now, s is a palindrome. So, answer[0] = true.
    fmt.Println(canMakePalindromeQueries("acbcab", [][]int{{1,2,4,5}})) // [true]

    fmt.Println(canMakePalindromeQueries1("abcabc", [][]int{{1,1,3,5},{0,2,5,5}})) // [true,true]
    fmt.Println(canMakePalindromeQueries1("abbcdecbba", [][]int{{0,2,7,9}})) // [false]
    fmt.Println(canMakePalindromeQueries1("acbcab", [][]int{{1,2,4,5}})) // [true]
}