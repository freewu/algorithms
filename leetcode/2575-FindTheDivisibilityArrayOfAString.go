package main

// 2575. Find the Divisibility Array of a String
// You are given a 0-indexed string word of length n consisting of digits, and a positive integer m.
// The divisibility array div of word is an integer array of length n such that:
//         div[i] = 1 if the numeric value of word[0,...,i] is divisible by m, or
//         div[i] = 0 otherwise.

// Return the divisibility array of word.

// Example 1:
// Input: word = "998244353", m = 3
// Output: [1,1,0,0,0,1,1,0,0]
// Explanation: There are only 4 prefixes that are divisible by 3: "9", "99", "998244", and "9982443".

// Example 2:
// Input: word = "1010", m = 10
// Output: [0,1,0,1]
// Explanation: There are only 2 prefixes that are divisible by 10: "10", and "1010".

// Constraints:
//         1 <= n <= 10^5
//         word.length == n
//         word consists of digits from 0 to 9
//         1 <= m <= 10^9

import "fmt"
import "strconv"

// 大数有问题
func divisibilityArray(word string, m int) []int {
    l := len(word)
    res := make([]int,l)
    for i := 0; i < l; i++ {
        n,_ := strconv.Atoi(word[:i + 1])
        if n % m == 0 {
            res[i] = 1
        }
        //fmt.Println(word[:i])
    }
    return res
}

// 模运算
// 一个整数可表示为 a × 10 + b:
//      (a × 10 + b) mod m = (a mod m × 10 + b) mod m
// 所以我们可以按照上面的递推式，根据当前表示整数的余数，算出包含下一位字符所表示的整数的余数。
// 当余数为零时即为可整除数组，否则不是。最后返回结果即可
func divisibilityArray1(word string, m int) []int {
    res := make([]int, 0)
    cur := 0
    for _, c := range word {
        // 一个整数可表示为 a × 10 + b 点睛之笔
        cur = (cur * 10 + int(c - '0')) % m
        if cur == 0 {
            res = append(res, 1)
        } else {
            res = append(res, 0)
        }
    }
    return res
}

func divisibilityArray2(word string, m int) []int {
    l := len(word)
    res := make([]int,l)
    cur := 0
    for i := 0; i < l; i++ {
        // 一个整数可表示为 a × 10 + b 点睛之笔
        cur = (cur * 10 + int(word[i] - '0')) % m
        // fmt.Println("cur: ", cur," i: ",i," word[i]: ",word[i])
        if cur == 0 { 
            res[i] = 1
        }
    }
    return res
}

func main() {
    // 仅有 4 个前缀可以被 3 整除："9"、"99"、"998244" 和 "9982443"
    fmt.Println(divisibilityArray("998244353",3)) // [1,1,0,0,0,1,1,0,0]
    // There are only 2 prefixes that are divisible by 10: "10", and "1010".
    fmt.Println(divisibilityArray("1010",10)) // [0,1,0,1]

    // 仅有 4 个前缀可以被 3 整除："9"、"99"、"998244" 和 "9982443"
    fmt.Println(divisibilityArray1("998244353",3)) // [1,1,0,0,0,1,1,0,0]
    // There are only 2 prefixes that are divisible by 10: "10", and "1010".
    fmt.Println(divisibilityArray1("1010",10)) // [0,1,0,1]

    // 仅有 4 个前缀可以被 3 整除："9"、"99"、"998244" 和 "9982443"
    fmt.Println(divisibilityArray2("998244353",3)) // [1,1,0,0,0,1,1,0,0]
    // There are only 2 prefixes that are divisible by 10: "10", and "1010".
    fmt.Println(divisibilityArray2("1010",10)) // [0,1,0,1]
}