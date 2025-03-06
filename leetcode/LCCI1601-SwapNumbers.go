package main

// 面试题 16.01. Swap Numbers LCCI
// Write a function to swap a number in place (that is, without temporary variables).

// Example:
// Input: numbers = [1,2]
// Output: [2,1]

// Note:
//     numbers.length == 2
//     -2147483647 <= numbers[i] <= 2147483647

import "fmt"

func swapNumbers(numbers []int) []int {
    return []int{ numbers[1], numbers[0] }
}

func main() {
    // Example:
    // Input: numbers = [1,2]
    // Output: [2,1]
    fmt.Println(swapNumbers([]int{1, 2})) // [2,1]

    fmt.Println(swapNumbers([]int{1 << 31, -1 << 31})) // [-2147483648 2147483648]
}