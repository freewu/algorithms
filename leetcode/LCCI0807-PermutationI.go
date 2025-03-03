package main

// 面试题 08.07. Permutation I LCCI
// Write a method to compute all permutations of a string of unique characters.

// Example1:
// Input: S = "qwe"
// Output: ["qwe", "qew", "wqe", "weq", "ewq", "eqw"]

// Example2:
// Input: S = "ab"
// Output: ["ab", "ba"]

// Note:
//     All charaters are English letters.
//     1 <= S.length <= 9

import "fmt"
import "strings"

func permutation(S string) []string {
    res, n := []string{}, len(S)
    if n < 1 || n > 9 { return nil }
    var dfs func(root string)
    dfs = func(root string) {
        // 触底条件
        s1 := root
        if len(root) == len(S) {
            res = append(res, s1)
            return
        }
        // 选择
        for i := 0; i < len(S); i++ {
            if !strings.Contains(root, S[i:i+1]) {
                root = root + S[i:i+1]
                // 递归
                dfs(root)
                // 撤销
                root = root[:len(root) - 1]
            }

        }
    }
    dfs("")
    return res
}

func permutation1(S string) []string {
    res, arr := []string{}, []byte(S)
    n := len(arr)
    visited := make([]bool, n)
    var dfs func(i int)
    dfs = func(i int) {
        if i >= n {
            res = append(res, string(arr))
            return
        }
        for j := range S {
            if !visited[j] {
                visited[j] = true
                arr[i] = S[j]
                dfs(i + 1)
                visited[j] = false
            }
        }
    }
    dfs(0)
    return res
}

func main() {
    // Example1:
    // Input: S = "qwe"
    // Output: ["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
    fmt.Println(permutation("qwe")) // ["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
    // Example2:
    // Input: S = "ab"
    // Output: ["ab", "ba"]
    fmt.Println(permutation("ab")) // ["ab", "ba"]

    fmt.Println(permutation("abc")) // [abc acb bac bca cab cba]
    //fmt.Println(permutation("bluefrog")) // 
    //fmt.Println(permutation("leetcode")) // 

    fmt.Println(permutation1("qwe")) // ["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
    fmt.Println(permutation1("ab")) // ["ab", "ba"]
    fmt.Println(permutation1("abc")) // [abc acb bac bca cab cba]
    //fmt.Println(permutation("bluefrog")) // 
    //fmt.Println(permutation("leetcode")) // 
}