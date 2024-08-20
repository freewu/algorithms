package main

// 989. Add to Array-Form of Integer
// The array-form of an integer num is an array representing its digits in left to right order.
//     For example, for num = 1321, the array form is [1,3,2,1].

// Given num, the array-form of an integer, and an integer k, return the array-form of the integer num + k.

// Example 1:
// Input: num = [1,2,0,0], k = 34
// Output: [1,2,3,4]
// Explanation: 1200 + 34 = 1234

// Example 2:
// Input: num = [2,7,4], k = 181
// Output: [4,5,5]
// Explanation: 274 + 181 = 455

// Example 3:
// Input: num = [2,1,5], k = 806
// Output: [1,0,2,1]
// Explanation: 215 + 806 = 1021

// Constraints:
//     1 <= num.length <= 10^4
//     0 <= num[i] <= 9
//     num does not contain any leading zeros except for the zero itself.
//     1 <= k <= 10^4

import "fmt"

func addToArrayForm(num []int, k int) []int {
    num[len(num) - 1] += k
    index := len(num) - 1
    for index >= 0 {
        if num[index] < 10 {
            return num
        }
        if index == 0 {
            num = append([]int{num[index] / 10}, num...)
            num[1] = num[1] % 10
        } else {
            num[index - 1] += num[index] / 10
            num[index] = num[index] % 10
            index--
        }
    }
    return num
}

func addToArrayForm1(num []int, k int) []int {
    up, n := 0, len(num)
    for i := n - 1; i >= 0; i-- {
        t := num[i] + k % 10 + up
        num[i], k, up = t % 10, k / 10, t / 10
    }

    for k > 0 || up > 0 {
        t := k % 10 + up
        num, k, up = append([]int{ t % 10 }, num...), k / 10, t / 10
    }
    return num
}

func main() {
    // Example 1:
    // Input: num = [1,2,0,0], k = 34
    // Output: [1,2,3,4]
    // Explanation: 1200 + 34 = 1234
    fmt.Println(addToArrayForm([]int{1,2,0,0}, 34 )) // [1,2,3,4]
    // Example 2:
    // Input: num = [2,7,4], k = 181
    // Output: [4,5,5]
    // Explanation: 274 + 181 = 455
    fmt.Println(addToArrayForm([]int{2,7,4}, 181 )) // [4,5,5]
    // Example 3:
    // Input: num = [2,1,5], k = 806
    // Output: [1,0,2,1]
    // Explanation: 215 + 806 = 1021
    fmt.Println(addToArrayForm([]int{2,1,5}, 806 )) // [1,0,2,1]

    fmt.Println(addToArrayForm1([]int{1,2,0,0}, 34 )) // [1,2,3,4]
    fmt.Println(addToArrayForm1([]int{2,7,4}, 181 )) // [4,5,5]
    fmt.Println(addToArrayForm1([]int{2,1,5}, 806 )) // [1,0,2,1]
}