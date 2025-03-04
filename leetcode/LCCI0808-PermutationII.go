package main

// 面试题 08.08. Permutation II LCCI
// Write a method to compute all permutations of a string whose characters are not necessarily unique. 
// The list of permutations should not have duplicates.

// Example1:
// Input: S = "qqe"
// Output: ["eqq","qeq","qqe"]

// Example2:
// Input: S = "ab"
// Output: ["ab", "ba"]

// Note:
//     All characters are English letters.
//     1 <= S.length <= 9

import "fmt"
import "sort"
import "strings"

func permutation(S string) []string {
    slice := []byte(S)
    sort.Slice(slice, func(i, j int) bool {
        return slice[i] < slice[j]
    })
    var helper func(str string, index int) []string
    helper = func(str string, index int) []string {
        if len(str[index:]) <= 1 { return []string{str} }
        origin, visited := []byte(str), []byte{}
        res := helper(str, index + 1)
        for i := index + 1; i < len(str); i++ {
            if str[i] == str[index] { continue }
            if strings.Contains(string(visited), string([]byte{str[i]})) {
                continue
            }
            visited = append(visited, str[i])
            origin[index], origin[i] = origin[i], origin[index]
            res = append(res, helper(string(origin), index + 1)...)
        }
        return res
    }
    return helper(string(slice), 0)
}

func permutation1(s string) []string {
    n, arr := len(s), []byte(s)
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    res, path, visited := make([]string, 0, n * (n - 1)),  make([]byte, n), make([]bool, n)
    var helper func(index int)
    helper = func(index int) {
        if index == n {
            res = append(res, string(path))
            return
        }
        for i := 0; i < n; i++ {
            if visited[i] || (i > 0 && arr[i] == arr[i-1] && !visited[i - 1]) { continue }
            visited[i], path[index] = true, arr[i]
            helper(index + 1)
            visited[i] = false
        }
    }
    helper(0)
    return res
}


func main() {
    // Example1:
    // Input: S = "qqe"
    // Output: ["eqq","qeq","qqe"]
    fmt.Println(permutation("qqe")) // ["eqq","qeq","qqe"]
    // Example2:
    // Input: S = "ab"
    // Output: ["ab", "ba"]
    fmt.Println(permutation("ab")) // ["ab", "ba"]

    fmt.Println(permutation("abc")) // [abc acb bac bca cab cba]
    //fmt.Println(permutation("bluefrog")) // 
    //fmt.Println(permutation("leetcode")) // 

    fmt.Println(permutation1("qqe")) // ["eqq","qeq","qqe"]
    fmt.Println(permutation1("ab")) // ["ab", "ba"]
    fmt.Println(permutation1("abc")) // [abc acb bac bca cab cba]
    // //fmt.Println(permutation("bluefrog")) // 
    // //fmt.Println(permutation("leetcode")) // 
}