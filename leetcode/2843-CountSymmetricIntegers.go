package main

// 2843. Count Symmetric Integers
// You are given two positive integers low and high.

// An integer x consisting of 2 * n digits is symmetric if the sum of the first n digits of x is equal to the sum of the last n digits of x. 
// Numbers with an odd number of digits are never symmetric.

// Return the number of symmetric integers in the range [low, high].

// Example 1:
// Input: low = 1, high = 100
// Output: 9
// Explanation: There are 9 symmetric integers between 1 and 100: 11, 22, 33, 44, 55, 66, 77, 88, and 99.

// Example 2:
// Input: low = 1200, high = 1230
// Output: 4
// Explanation: There are 4 symmetric integers between 1200 and 1230: 1203, 1212, 1221, and 1230.

// Constraints:
//     1 <= low <= high <= 10^4

import "fmt"
import "strconv"

func countSymmetricIntegers(low int, high int) int {
    res := 0
    symmetric := func(i int) bool {
        s := strconv.Itoa(i)
        n := len(s)
        if n % 2== 1 {return false }
        s1, s2 := 0, 0
        for i := 0; i < n / 2; i++ {
            s1 += int(s[i] - '0')
            s2 += int(s[n - i - 1] - '0')
        }
        return s1 == s2
    }
    for i := low; i <= high; i++ {
        if symmetric(i) {
           res++
        }
    }
    return res
}

func countSymmetricIntegers1(low int, high int) int {
    // 普通算法,必须偶数位 mx=10^4=1w   所以x必须为2位或者4位
    res := 0
    for i := low; i <= high; i++ {
        // trick!! 可以转换为字符串, 好处1:方便看长度n是否是偶数  好处2:方便计算左右sum, lSum=sum(range s[0:n/2]) rSum=sum(range s[n/2:n])
        if 10 <= i && i <= 99 {
            if i / 10 == i % 10 {
                res++
            }
        } else if 1000 <= i && i <= 9999 {
            if i / 1000 + i % 1000 / 100 == i % 100 / 10 + i % 10 {
                res++
            }
        }
    }
    return res
}

var s [10001]int

func init() {
    for i := 1; i <= 9; i++ {
        s[11*i]++
    }
    for x4 := 1; x4 <= 9; x4++ {
        for x3 := 0; x3 <= 9; x3++ {
            s2 := x3 + x4
            for x2 := max(0, s2-9); x2 <= min(9, s2); x2++ {
                x1 := s2 - x2
                s[x1+x2*10+x3*100+x4*1000]++
            }
        }
    }
    for i := 1; i <= 10000; i++ {
        s[i] += s[i-1]
    }
}

func countSymmetricIntegers2(low int, high int) int {
    return s[high] - s[low - 1]
}

func main() {
    // Example 1:
    // Input: low = 1, high = 100
    // Output: 9
    // Explanation: There are 9 symmetric integers between 1 and 100: 11, 22, 33, 44, 55, 66, 77, 88, and 99.
    fmt.Println(countSymmetricIntegers(1, 100)) // 9
    // Example 2:
    // Input: low = 1200, high = 1230
    // Output: 4
    // Explanation: There are 4 symmetric integers between 1200 and 1230: 1203, 1212, 1221, and 1230.
    fmt.Println(countSymmetricIntegers(1200, 1230)) // 4

    fmt.Println(countSymmetricIntegers(1, 10000)) // 624
    fmt.Println(countSymmetricIntegers(1, 1)) // 0
    fmt.Println(countSymmetricIntegers(10000, 10000)) // 0

    fmt.Println(countSymmetricIntegers1(1, 100)) // 9
    fmt.Println(countSymmetricIntegers1(1200, 1230)) // 4
    fmt.Println(countSymmetricIntegers1(1, 10000)) // 624
    fmt.Println(countSymmetricIntegers1(1, 1)) // 0
    fmt.Println(countSymmetricIntegers1(10000, 10000)) // 0

    fmt.Println(countSymmetricIntegers2(1, 100)) // 9
    fmt.Println(countSymmetricIntegers2(1200, 1230)) // 4
    fmt.Println(countSymmetricIntegers2(1, 10000)) // 624
    fmt.Println(countSymmetricIntegers2(1, 1)) // 0
    fmt.Println(countSymmetricIntegers2(10000, 10000)) // 0
}