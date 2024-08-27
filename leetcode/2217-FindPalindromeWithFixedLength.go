package main

// 2217. Find Palindrome With Fixed Length
// Given an integer array queries and a positive integer intLength, 
// return an array answer where answer[i] is either the queries[i]th smallest positive palindrome of length intLength or -1 if no such palindrome exists.

// A palindrome is a number that reads the same backwards and forwards. Palindromes cannot have leading zeros.

// Example 1:
// Input: queries = [1,2,3,4,5,90], intLength = 3
// Output: [101,111,121,131,141,999]
// Explanation:
// The first few palindromes of length 3 are:
// 101, 111, 121, 131, 141, 151, 161, 171, 181, 191, 202, ...
// The 90th palindrome of length 3 is 999.

// Example 2:
// Input: queries = [2,4,6], intLength = 4
// Output: [1111,1331,1551]
// Explanation:
// The first six palindromes of length 4 are:
// 1001, 1111, 1221, 1331, 1441, and 1551.

// Constraints:
//     1 <= queries.length <= 5 * 10^4
//     1 <= queries[i] <= 10^9
//     1 <= intLength <= 15

import "fmt"
import "math"

func kthPalindrome(queries []int, intLength int) []int64 {
    arr := make([]int64, len(queries))
    getPalindrome := func(x int, n int) int64 {
        if x > int(math.Pow10((n + 1)/2)) - int(math.Pow10((n - 1)/2)) {
            return int64(-1)
        }
        num := int((x - 1) + int(math.Pow10((n - 1)/2))) * int(math.Pow10((n) / 2))
        x = (num / (int(math.Pow10((n + 1)/2))))
        count, revx := (n) / 2, 0
        for i := 1; x > 0; i++{
            revx = revx + (x % 10) * int(math.Pow10(count - i))
            x = x / 10
        }
        return int64(num + revx)
    }
    for i, v := range queries {
        arr[i] = getPalindrome(v, intLength)
    }
    return arr
}

func kthPalindrome1(queries []int, intLength int) []int64 {
    n, base := len(queries), int(math.Pow10((intLength - 1) >> 1))
    mx := 9 * base
    res := make([]int64, n)
    for i, q := range queries {
        if q > mx {
            res[i] = -1
            continue
        }
        v := base + q - 1
        val := v
        if intLength & 1 != 0 {
            val /= 10
        }
        for ; val != 0; val /= 10 {
            v = v * 10 + val%10
        }
        res[i] = int64(v)
    }
    return res
}

func main() {
    // Example 1:
    // Input: queries = [1,2,3,4,5,90], intLength = 3
    // Output: [101,111,121,131,141,999]
    // Explanation:
    // The first few palindromes of length 3 are:
    // 101, 111, 121, 131, 141, 151, 161, 171, 181, 191, 202, ...
    // The 90th palindrome of length 3 is 999.
    fmt.Println(kthPalindrome([]int{1,2,3,4,5,90}, 3)) // [101,111,121,131,141,999]
    // Example 2:
    // Input: queries = [2,4,6], intLength = 4
    // Output: [1111,1331,1551]
    // Explanation:
    // The first six palindromes of length 4 are:
    // 1001, 1111, 1221, 1331, 1441, and 1551.
    fmt.Println(kthPalindrome([]int{2,4,6}, 4)) // [1111,1331,1551]

    fmt.Println(kthPalindrome1([]int{1,2,3,4,5,90}, 3)) // [101,111,121,131,141,999]
    fmt.Println(kthPalindrome1([]int{2,4,6}, 4)) // [1111,1331,1551]
}