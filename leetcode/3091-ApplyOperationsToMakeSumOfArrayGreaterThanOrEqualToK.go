package main

// 3091. Apply Operations to Make Sum of Array Greater Than or Equal to k
// You are given a positive integer k. Initially, you have an array nums = [1].

// You can perform any of the following operations on the array any number of times (possibly zero):
//     1. Choose any element in the array and increase its value by 1.
//     2. Duplicate any element in the array and add it to the end of the array.

// Return the minimum number of operations required to make the sum of elements of the final array greater than or equal to k.

// Example 1:
// Input: k = 11
// Output: 5
// Explanation:
// We can do the following operations on the array nums = [1]:
// Increase the element by 1 three times. The resulting array is nums = [4].
// Duplicate the element two times. The resulting array is nums = [4,4,4].
// The sum of the final array is 4 + 4 + 4 = 12 which is greater than or equal to k = 11.
// The total number of operations performed is 3 + 2 = 5.

// Example 2:
// Input: k = 1
// Output: 0
// Explanation:
// The sum of the original array is already greater than or equal to 1, so no operations are needed.

// Constraints:
//     1 <= k <= 10^5

import "fmt"

func minOperations(k int) int {
    if k == 1 { return 0 }
    res, curr := 1 << 61, 1
    for curr <= k / 2 {
        // you dont have to use array just one variable and
        // increment it and by deviding you can calculate 
        // how many times you have to append
        curr++
        op := (k / curr ) + curr 
        if k % curr == 0 {
            op--
        }
        if res > op {
            res = op
        }
    }
    return res - 1
}

func minOperations1(k int) int {
    res := k - 1
    for i := 1; i < k; i++ {
        count := k / i - 1
        count += i - 1
        if k % i > 0 {
            count++
        }
        if count > res {
            break
        }
        res = count
    }
    return res
}

func main() {
    // Example 1:
    // Input: k = 11
    // Output: 5
    // Explanation:
    // We can do the following operations on the array nums = [1]:
    // Increase the element by 1 three times. The resulting array is nums = [4].
    // Duplicate the element two times. The resulting array is nums = [4,4,4].
    // The sum of the final array is 4 + 4 + 4 = 12 which is greater than or equal to k = 11.
    // The total number of operations performed is 3 + 2 = 5.
    fmt.Println(minOperations(11)) // 5
    // Example 2:
    // Input: k = 1
    // Output: 0
    // Explanation:
    // The sum of the original array is already greater than or equal to 1, so no operations are needed.
    fmt.Println(minOperations(1)) // 0

    fmt.Println(minOperations(2)) // 1
    fmt.Println(minOperations(3)) // 2
    fmt.Println(minOperations(4)) // 2
    fmt.Println(minOperations(5)) // 62
    fmt.Println(minOperations(999)) // 62
    fmt.Println(minOperations(1000)) // 62
    fmt.Println(minOperations(1024)) // 62
    fmt.Println(minOperations(99999)) // 631
    fmt.Println(minOperations(100000)) // 631

    fmt.Println(minOperations1(11)) // 5
    fmt.Println(minOperations1(1)) // 0
    fmt.Println(minOperations1(2)) // 1
    fmt.Println(minOperations1(3)) // 2
    fmt.Println(minOperations1(4)) // 2
    fmt.Println(minOperations1(5)) // 62
    fmt.Println(minOperations1(999)) // 62
    fmt.Println(minOperations1(1000)) // 62
    fmt.Println(minOperations1(1024)) // 62
    fmt.Println(minOperations1(99999)) // 631
    fmt.Println(minOperations1(100000)) // 631
}