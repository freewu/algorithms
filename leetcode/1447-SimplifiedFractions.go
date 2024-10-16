package main

// 1447. Simplified Fractions
// Given an integer n, return a list of all simplified fractions between 0 and 1 (exclusive) such that the denominator is less-than-or-equal-to n. 
// You can return the answer in any order.

// Example 1:
// Input: n = 2
// Output: ["1/2"]
// Explanation: "1/2" is the only unique fraction with a denominator less-than-or-equal-to 2.

// Example 2:
// Input: n = 3
// Output: ["1/2","1/3","2/3"]

// Example 3:
// Input: n = 4
// Output: ["1/2","1/3","1/4","2/3","3/4"]
// Explanation: "2/4" is not a simplified fraction because it can be simplified to "1/2".

// Constraints:
//     1 <= n <= 100

import "fmt"
import "strconv"

func simplifiedFractions(n int) []string {
    res, mp := []string{}, make(map[string]bool)
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    genKey := func (numerator, denominator int) string { return strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator) }
    buildForInt := func(n int) {
        for i:= 1; i<n; i++ {
            if gcd(i, n) == 1 {
                mp[genKey(i, n)] = true
            }
        }
    }
    for i:= 1; i<=n; i++ {
        buildForInt(i)
    }
    for k, _ := range mp {
        res = append(res, k)
    }
    return res
}

func simplifiedFractions1(n int) []string {
    res := []string{}
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for denominator := 2; denominator <= n; denominator++ { // 遍历所有可能的分母
        for numerator := 1; numerator < denominator; numerator++ { // 遍历所有可能的分子
            if gcd(numerator, denominator) == 1 { // 检查是否为最简分数
                res = append(res, strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator))
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: ["1/2"]
    // Explanation: "1/2" is the only unique fraction with a denominator less-than-or-equal-to 2.
    fmt.Println(simplifiedFractions(2)) // ["1/2"]
    // Example 2:
    // Input: n = 3
    // Output: ["1/2","1/3","2/3"]
    fmt.Println(simplifiedFractions(3)) // ["1/2","1/3","2/3"]
    // Example 3:
    // Input: n = 4
    // Output: ["1/2","1/3","1/4","2/3","3/4"]
    // Explanation: "2/4" is not a simplified fraction because it can be simplified to "1/2".
    fmt.Println(simplifiedFractions(4)) // ["1/2","1/3","1/4","2/3","3/4"]

    fmt.Println(simplifiedFractions1(2)) // ["1/2"]
    fmt.Println(simplifiedFractions1(3)) // ["1/2","1/3","2/3"]
    fmt.Println(simplifiedFractions1(4)) // ["1/2","1/3","1/4","2/3","3/4"]
}