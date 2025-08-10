package main

// 869. Reordered Power of 2
// You are given an integer n. 
// We reorder the digits in any order (including the original order) such that the leading digit is not zero.

// Return true if and only if we can do this so that the resulting number is a power of two.

// Example 1:
// Input: n = 1
// Output: true

// Example 2:
// Input: n = 10
// Output: false

// Constraints:
//     1 <= n <= 10^9

import "fmt"
import "sort"
import "strconv"
import "strings"

func reorderedPowerOf2(n int) bool {
    s := strconv.Itoa(n)
    canEqual := func(a, b string) bool {
        l1, l2 := []byte(a), []byte(b)
        sort.Slice(l1, func(i, j int) bool { 
            return l1[i] < l1[j] 
        })
        sort.Slice(l2, func(i, j int) bool { 
            return l2[i] < l2[j] 
        })
        return string(l1) == string(l2)
    }
    for i := 1; i <= 1000000000; i *= 2 {
        if canEqual(s, strconv.Itoa(i)) {
            return true
        }
    }
    return false
}

func reorderedPowerOf21(n int) bool {
    SortString := func(w string) string {
        s := strings.Split(w, "")
        sort.Strings(s)
        return strings.Join(s, "")
    }
    str := SortString(strconv.Itoa(n))
    for i := 0; i <= 29; i++ {
        if SortString(strconv.Itoa(1 << i)) == str {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: true
    fmt.Println(reorderedPowerOf2(1)) // true
    // Example 2:
    // Input: n = 10
    // Output: false
    fmt.Println(reorderedPowerOf2(10)) // false

    fmt.Println(reorderedPowerOf2(1024)) // true
    fmt.Println(reorderedPowerOf2(999_999_999)) // false
    fmt.Println(reorderedPowerOf2(1_000_000_000)) // false

    fmt.Println(reorderedPowerOf21(1)) // true
    fmt.Println(reorderedPowerOf21(10)) // false
    fmt.Println(reorderedPowerOf21(1024)) // true
    fmt.Println(reorderedPowerOf21(999_999_999)) // false
    fmt.Println(reorderedPowerOf21(1_000_000_000)) // false
}