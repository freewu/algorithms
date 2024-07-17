package main

// 1012. Numbers With Repeated Digits
// Given an integer n, return the number of positive integers in the range [1, n] that have at least one repeated digit.

// Example 1:
// Input: n = 20
// Output: 1
// Explanation: The only positive number (<= 20) with at least 1 repeated digit is 11.

// Example 2:
// Input: n = 100
// Output: 10
// Explanation: The positive numbers (<= 100) with atleast 1 repeated digit are 11, 22, 33, 44, 55, 66, 77, 88, 99, and 100.

// Example 3:
// Input: n = 1000
// Output: 262

// Constraints:
//     1 <= n <= 10^9

import "fmt"

func numDupDigitsAtMostN(n int) int {
    if n <= 10 { return 0 }
    if n <= 20 { return 1 }
    nums, t := make([]int, 0), n
    for t > 0 {
        nums = append(nums, t % 10)
        t = t/10
    }
    d, add, now, sum := 2, 9, 9, 9
    for d < len(nums) {
        now *= add
        sum += now
        d++
        add--
    }
    m := make(map[int]bool)
    for i := len(nums) - 1; i >= 0; i-- {
        now = nums[i]
        if i == len(nums) - 1 {
            now--
        } else {
            tmp := 0
            for x := 0; x < now; x++ {
                if !m[x] {
                    tmp++
                }
            }
            now = tmp
        }
        add = 10 - (len(nums) - i)
        d = i
        for d > 0 {
            now *= add
            add--
            d--
        }
        sum += now
        if m[nums[i]] {
            break
        }
        m[nums[i]] = true
    }
    if len(m) == len(nums) {
        sum++
    }
    return n - sum
}

func numDupDigitsAtMostN1(n int) int {
    var help func(m, n int) int 
    help = func(m, n int) int {
        if n == 0 { return 1 }
        return help(m, n-1) * (m - n + 1)
    }
    nonRepeating := func(n int) int {
        digits := []int{}
        for n != 0 {
            digits = append(digits, n%10)
            n /= 10
        }
        res, m, vis := 0, len(digits), make([]bool, 10)
        for i := 1; i < m; i++ {
            res += 9 * help(9, i-1)
        }
        for i := m - 1; i >= 0; i-- {
            v := digits[i]
            j := 0
            if i == m-1 {
                j = 1
            }
            for ; j < v; j++ {
                if !vis[j] {
                    res += help(10-(m-i), i)
                }
            }
            if vis[v] { break }
            vis[v] = true
            if i == 0 {
                res++
            }
        }
        return res
    }
    return n - nonRepeating(n)
}

func main() {
    // Example 1:
    // Input: n = 20
    // Output: 1
    // Explanation: The only positive number (<= 20) with at least 1 repeated digit is 11.
    fmt.Println(numDupDigitsAtMostN(20)) // 1
    // Example 2:
    // Input: n = 100
    // Output: 10
    // Explanation: The positive numbers (<= 100) with atleast 1 repeated digit are 11, 22, 33, 44, 55, 66, 77, 88, 99, and 100.
    fmt.Println(numDupDigitsAtMostN(100)) // 10
    // Example 3:
    // Input: n = 1000
    // Output: 262
    fmt.Println(numDupDigitsAtMostN(1000)) // 262

    fmt.Println(numDupDigitsAtMostN1(20)) // 1
    fmt.Println(numDupDigitsAtMostN1(100)) // 10
    fmt.Println(numDupDigitsAtMostN1(1000)) // 262
}