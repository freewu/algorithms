package main

// 247. Strobogrammatic Number II
// Given an integer n, return all the strobogrammatic numbers that are of length n. 
// You may return the answer in any order.
// A strobogrammatic number is a number that looks the same when rotated 180 degrees (looked at upside down).

// Example 1:
// Input: n = 2
// Output: ["11","69","88","96"]

// Example 2:
// Input: n = 1
// Output: ["0","1","8"]
 
// Constraints:
//     1 <= n <= 14

import "fmt"

func findStrobogrammatic(n int) []string {
    rotate := map[string]string{"0": "0", "1": "1", "6": "9", "8": "8", "9": "6"}
    var helper func(n int, isHead bool) []string
    helper = func(n int, isHead bool) []string {
        if n == 0 {
            return []string{""}
        }
        if n == 1 {
            return []string{"0", "1", "8"}
        }
        res := []string{}
        mids := helper(n - 2, false)
        for k, v := range rotate {
            if isHead && k == "0" { // 第一个字符需要 排除 0 开头的数字
                continue
            }
            for _, mid := range mids {
                res = append(res, k + mid + v)
            }
        }
        return res
    }
    return helper(n, true)
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: ["11","69","88","96"]
    fmt.Println(findStrobogrammatic(2)) // ["11","69","88","96"]
    // Example 2:
    // Input: n = 1
    // Output: ["0","1","8"]
    fmt.Println(findStrobogrammatic(1)) // ["0","1","8"]

    fmt.Println(findStrobogrammatic(3)) // [808 818 888 906 916 986 101 111 181 609 619 689]
    fmt.Println(findStrobogrammatic(4)) // [1001 1111 1691 1881 1961 6009 6119 6699 6889 6969 8008 8118 8698 8888 8968 9006 9116 9696 9886 9966]
}