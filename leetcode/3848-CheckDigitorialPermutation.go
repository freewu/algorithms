package main

// 3848. Check Digitorial Permutation
// You are given an integer n.

// A number is called digitorial if the sum of the factorials of its digits is equal to the number itself.

// Determine whether any permutation of n (including the original order) forms a digitorial number.

// Return true if such a permutation exists, otherwise return false.

// Note:
//     1. The factorial of a non-negative integer x, denoted as x!, 
//        is the product of all positive integers less than or equal to x, and 0! = 1.
//     2. A permutation is a rearrangement of all the digits of a number that does not start with zero. 
//        Any arrangement starting with zero is invalid.
 
// Example 1:
// Input: n = 145
// Output: true
// Explanation:
// The number 145 itself is digitorial since 1! + 4! + 5! = 1 + 24 + 120 = 145. Thus, the answer is true.

// Example 2:
// Input: n = 10
// Output: false
// Explanation:​​​​​​​
// 10 is not digitorial since 1! + 0! = 2 is not equal to 10, and the permutation "01" is invalid because it starts with zero.

// Constraints:
//     1 <= n <= 10^9   

import "fmt"
import "sort"

var fac = [10]int{ 1 }

func init() {
    for i := 1; i < len(fac); i++ { // 预处理阶乘
        fac[i] = fac[i-1] * i
    }
}

func isDigitorialPermutation(n int) bool {
    sum, count := 0, [10]int{}
    for ; n > 0; n /= 10 {
        d := n % 10
        sum += fac[d]
        count[d]++
    }
    for ; sum > 0; sum /= 10 {
        count[sum % 10]--
    }
    return count == [10]int{} // count[i] == 0
}

func isDigitorialPermutation1(n int) bool {
    arr1, arr2 := []int{}, []int{}
    for n > 0 {
        arr1 = append(arr1,n % 10)
        n /= 10
    }
    sort.Ints(arr1)
    sum := 0
    for _,v := range arr1 {
        t := 1
        for i := 1; i <= v; i++ {
            t *= i
        }
        sum += t
    }
    for sum > 0 {
        arr2 = append(arr2,sum % 10)
        sum /= 10
    }
    sort.Ints(arr2)
    if len(arr1) != len(arr2) { return false }
    for i := range arr1 {
        if arr1[i] != arr2[i] { 
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: n = 145
    // Output: true
    // Explanation:
    // The number 145 itself is digitorial since 1! + 4! + 5! = 1 + 24 + 120 = 145. Thus, the answer is true.
    fmt.Println(isDigitorialPermutation(145)) // true
    // Example 2:
    // Input: n = 10
    // Output: false
    // Explanation:​​​​​​​
    // 10 is not digitorial since 1! + 0! = 2 is not equal to 10, and the permutation "01" is invalid because it starts with zero.
    fmt.Println(isDigitorialPermutation(10)) // false

    fmt.Println(isDigitorialPermutation(1)) // true
    fmt.Println(isDigitorialPermutation(8)) // false
    fmt.Println(isDigitorialPermutation(99)) // false
    fmt.Println(isDigitorialPermutation(100)) // false
    fmt.Println(isDigitorialPermutation(1024)) // false
    fmt.Println(isDigitorialPermutation(999_999_999)) // false
    fmt.Println(isDigitorialPermutation(1_000_000_000)) // false

    fmt.Println(isDigitorialPermutation1(145)) // true
    fmt.Println(isDigitorialPermutation1(10)) // false
    fmt.Println(isDigitorialPermutation1(1)) // true
    fmt.Println(isDigitorialPermutation1(8)) // false
    fmt.Println(isDigitorialPermutation1(99)) // false
    fmt.Println(isDigitorialPermutation1(100)) // false
    fmt.Println(isDigitorialPermutation1(1024)) // false
    fmt.Println(isDigitorialPermutation1(999_999_999)) // false
    fmt.Println(isDigitorialPermutation1(1_000_000_000)) // false
}