package main

// 3272. Find the Count of Good Integers
// You are given two positive integers n and k.

// An integer x is called k-palindromic if:
//     x is a palindrome.
//     x is divisible by k.

// An integer is called good if its digits can be rearranged to form a k-palindromic integer. 
// For example, for k = 2, 2020 can be rearranged to form the k-palindromic integer 2002, whereas 1010 cannot be rearranged to form a k-palindromic integer.

// Return the count of good integers containing n digits.

// Note that any integer must not have leading zeros, neither before nor after rearrangement. 
// For example, 1010 cannot be rearranged to form 101.

// Example 1:
// Input: n = 3, k = 5
// Output: 27
// Explanation:
// Some of the good integers are:
// 551 because it can be rearranged to form 515.
// 525 because it is already k-palindromic.

// Example 2:
// Input: n = 1, k = 4
// Output: 2
// Explanation:
// The two good integers are 4 and 8.

// Example 3:
// Input: n = 5, k = 6
// Output: 2468

// Constraints:
//     1 <= n <= 10
//     1 <= k <= 9

import "fmt"
import "math"
import "strconv"
import "slices"

func countGoodIntegers(n int, k int) int64 {
    factorial := func(n int) []int64 {
        fac := make([]int64, n + 1)
        fac[0] = 1
        for i := 1; i <= n; i++ {
            fac[i] = fac[i-1] * int64(i)
        }
        return fac
    }
    reverseString := func(s string) string {
        arr := []byte(s)
        for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
            arr[i], arr[j] = arr[j], arr[i]
        }
        return string(arr)
    }
    fac := factorial(n)
    visited := make(map[string]bool)
    res, base := int64(0), int(math.Pow(10, float64((n-1)/2)))
    for i := base; i < base * 10; i++ {
        s := strconv.Itoa(i)
        rev := reverseString(s)
        s += rev[n % 2:]
        v, _ := strconv.ParseInt(s, 10, 64)
        if v % int64(k) != 0 { continue }
        bs := []byte(s)
        slices.Sort(bs)
        t := string(bs)
        if visited[t] { continue }
        visited[t] = true
        count := make([]int, 10)
        for _, c := range t {
            count[c - '0']++
        }
        val := (int64(n) - int64(count[0])) * fac[n-1]
        for _, v := range count {
            val /= fac[v]
        }
        res += val
    }
    return res
}

func countGoodIntegers1(n int, k int) int64 {
    mp := map[[10]int]bool{}
    put := func(x int) {
        count := [10]int{}
        for ; x > 0; x /= 10 {
            count[x % 10]++
        }
        mp[count] = true
    }
    helper := func(x int) {
        res, power, i := 0, 1, x
        if n&1 == 1 {
            i /= 10
        }
        for ; i > 0; i /= 10 {
            res = res * 10 + i % 10
            power *= 10
        }
        x = x * power + res
        if x % k == 0 {
            put(x)
        }
    }
    res, m := 0, (n + 1) / 2
    mx := int(math.Pow10(m))
    for i := int(math.Pow10(m - 1)); i < mx; i++ {
        helper(i)
    }
    arr := [11]int{}
    arr[0] = 1
    for i := 1; i <= 10; i++ {
        arr[i] = arr[i - 1] * i
    }
    for count := range mp {
        sum := 0
        for _, v := range count {
            sum += v
        }
        val := (sum - count[0]) * arr[sum - 1]
        for _, v := range count {
            val /= arr[v]
        }
        res += val
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: n = 3, k = 5
    // Output: 27
    // Explanation:
    // Some of the good integers are:
    // 551 because it can be rearranged to form 515.
    // 525 because it is already k-palindromic.
    fmt.Println(countGoodIntegers(3, 5)) // 27
    // Example 2:
    // Input: n = 1, k = 4
    // Output: 2
    // Explanation:
    // The two good integers are 4 and 8.
    fmt.Println(countGoodIntegers(1, 4)) // 2
    // Example 3:
    // Input: n = 5, k = 6
    // Output: 2468
    fmt.Println(countGoodIntegers(5, 6)) // 22468

    fmt.Println(countGoodIntegers(1, 1)) // 9
    fmt.Println(countGoodIntegers(1, 9)) // 1
    fmt.Println(countGoodIntegers(10, 1)) // 41457024
    fmt.Println(countGoodIntegers(9, 9)) // 4623119
    fmt.Println(countGoodIntegers(10, 9)) // 4610368

    fmt.Println(countGoodIntegers1(3, 5)) // 27
    fmt.Println(countGoodIntegers1(1, 4)) // 2
    fmt.Println(countGoodIntegers1(5, 6)) // 22468
    fmt.Println(countGoodIntegers1(1, 1)) // 9
    fmt.Println(countGoodIntegers1(1, 9)) // 1
    fmt.Println(countGoodIntegers1(10, 1)) // 41457024
    fmt.Println(countGoodIntegers1(9, 9)) // 4623119
    fmt.Println(countGoodIntegers1(10, 9)) // 4610368
}