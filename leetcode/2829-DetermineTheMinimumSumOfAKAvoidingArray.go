package main

// 2829. Determine the Minimum Sum of a k-avoiding Array
// You are given two integers, n and k.

// An array of distinct positive integers is called a k-avoiding array if there does not exist any pair of distinct elements that sum to k.

// Return the minimum possible sum of a k-avoiding array of length n.

// Example 1:
// Input: n = 5, k = 4
// Output: 18
// Explanation: Consider the k-avoiding array [1,2,4,5,6], which has a sum of 18.
// It can be proven that there is no k-avoiding array with a sum less than 18.

// Example 2:
// Input: n = 2, k = 6
// Output: 3
// Explanation: We can construct the array [1,2], which has a sum of 3.
// It can be proven that there is no k-avoiding array with a sum less than 3.

// Constraints:
//     1 <= n, k <= 50

import "fmt"

func minimumSum(n int, k int) int {
    mid := k / 2
    if n <= mid { // if n <= middle, grab all from 1 to n
        return (n + 1) * n / 2
    }
    // if n > mid 
    // first grab all from 1 to mid
    // cuz now we got: [1, mid]
    // and no pair could have sum k, so no [mid + 1, k-1]
    // so the next move's range: [k, ifinity)
    // but we only need n - mid more elements
    // so the next part is [k, k + (n - mid)-1]
    return (mid + 1) * mid / 2 + (2 * k + n - mid - 1) * (n - mid) / 2
}

func minimumSum1(n int, k int) int {
    count, sum := 0, 0
    for i := 1; i <= k/2; i++ {
        sum += i
        count++
        if count == n {
            return sum
        }
    }
    for i := k;; i++ {
        sum += i
        count++
        if count == n {
            break
        }
    }
    return sum
}

func main() {
    // Example 1:
    // Input: n = 5, k = 4
    // Output: 18
    // Explanation: Consider the k-avoiding array [1,2,4,5,6], which has a sum of 18.
    // It can be proven that there is no k-avoiding array with a sum less than 18.
    fmt.Println(minimumSum(5, 4)) // 18
    // Example 2:
    // Input: n = 2, k = 6
    // Output: 3
    // Explanation: We can construct the array [1,2], which has a sum of 3.
    // It can be proven that there is no k-avoiding array with a sum less than 3.
    fmt.Println(minimumSum(2, 6)) // 3

    fmt.Println(minimumSum(1, 1)) // 1
    fmt.Println(minimumSum(50, 50)) // 1875
    fmt.Println(minimumSum(1, 50)) // 1
    fmt.Println(minimumSum(50, 1)) // 1275

    fmt.Println(minimumSum1(5, 4)) // 18
    fmt.Println(minimumSum1(2, 6)) // 3
    fmt.Println(minimumSum1(1, 1)) // 1
    fmt.Println(minimumSum1(50, 50)) // 1875
    fmt.Println(minimumSum1(1, 50)) // 1
    fmt.Println(minimumSum1(50, 1)) // 1275
}