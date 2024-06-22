package main

// 564. Find the Closest Palindrome
// Given a string n representing an integer, 
// return the closest integer (not including itself), which is a palindrome. 
// If there is a tie, return the smaller one.

// The closest is defined as the absolute difference minimized between two integers.
 
// Example 1:
// Input: n = "123"
// Output: "121"

// Example 2:
// Input: n = "1"
// Output: "0"
// Explanation: 0 and 2 are the closest palindromes but we return the smallest which is 0.

// Constraints:
//     1 <= n.length <= 18
//     n consists of only digits.
//     n does not have leading zeros.
//     n is representing an integer in the range [1, 10^18 - 1].

import "fmt"
import "strconv"
import "math"

func nearestPalindromic(n string) string {
    runes := []rune(n)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    mirror := func(arr []rune) string {
        res := make([]rune, len(runes))
        copy(res, runes)
        copy(res, arr)
        m := (len(res) - 1) / 2
        for i := len(res) -1; i > m; i-- {
            index := len(res)-1-i
            res[i] = res[index]
        }
        return string(res)
    }
    m := len(runes) / 2 + len(n) % 2
    cand, left := []int{}, runes[:m]
    cand = append(cand, int(math.Pow10(len(runes)-1)-1))
    for _, d := range []int{-1, 0, 1} {
        num, _ := strconv.Atoi(string(left))
        num += d
        l := strconv.Itoa(num)
        c := mirror([]rune(l))
        v, _ := strconv.Atoi(c)
        cand = append(cand, v)
    }
    cand = append(cand, int(math.Pow10(len(runes))+1))
    res, diff := 0, math.MaxInt
    orig, _ := strconv.Atoi(n)
    for _, v := range cand {
        if abs(v - orig) < diff && v != orig {
            res = v
            diff = abs(v - orig)
        }
    }
    return strconv.Itoa(res)
}

func nearestPalindromic1(n string) string {
    m := len(n)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    candidates := []int{int(math.Pow10(m-1)) - 1, int(math.Pow10(m)) + 1}
    selfPrefix, _ := strconv.Atoi(n[:(m+1)/2])
    for _, x := range []int{selfPrefix - 1, selfPrefix, selfPrefix + 1} {
        y := x
        if m & 1 == 1 {
            y /= 10
        }
        for ; y > 0; y /= 10 {
            x = x*10 + y%10
        }
        candidates = append(candidates, x)
    }
    res := -1
    selfNumber, _ := strconv.Atoi(n)
    for _, candidate := range candidates {
        if candidate != selfNumber {
            if res == -1 ||
               abs(candidate - selfNumber) < abs(res - selfNumber) ||
               abs(candidate - selfNumber) == abs(res - selfNumber) && 
               candidate < res {
                res = candidate
            }
        }
    }
    return strconv.Itoa(res)
}

func main() {
    // Example 1:
    // Input: n = "123"
    // Output: "121"
    fmt.Println(nearestPalindromic("123")) // "121"
    // Example 2:
    // Input: n = "1"
    // Output: "0"
    // Explanation: 0 and 2 are the closest palindromes but we return the smallest which is 0.
    fmt.Println(nearestPalindromic("1")) // "0"

    fmt.Println(nearestPalindromic1("123")) // "121"
    fmt.Println(nearestPalindromic1("1")) // "0"
}
