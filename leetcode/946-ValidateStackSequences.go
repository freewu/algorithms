package main

// 946. Validate Stack Sequences
// Given two integer arrays pushed and popped each with distinct values, 
// return true if this could have been the result of a sequence of push 
// and pop operations on an initially empty stack, or false otherwise.

// Example 1:
// Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
// Output: true
// Explanation: We might do the following sequence:
// push(1), push(2), push(3), push(4),
// pop() -> 4,
// push(5),
// pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

// Example 2:
// Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
// Output: false
// Explanation: 1 cannot be popped before 2.

// Constraints:
//     1 <= pushed.length <= 1000
//     0 <= pushed[i] <= 1000
//     All the elements of pushed are unique.
//     popped.length == pushed.length
//     popped is a permutation of pushed.

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
    stack, x := []int{}, 0
    for _, val := range pushed {
        stack = append(stack, val)
        for len(stack) > 0 && stack[ len(stack) - 1] == popped[x] {
            x++
            stack = stack[:len(stack)-1] // pop operation
        }
    }
    return len(stack) == 0
}

func main() {
    // Example 1:
    // Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
    // Output: true
    // Explanation: We might do the following sequence:
    // push(1), push(2), push(3), push(4),
    // pop() -> 4,
    // push(5),
    // pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
    fmt.Println(validateStackSequences([]int{1,2,3,4,5}, []int{4,5,3,2,1})) // true
    // Example 2:
    // Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
    // Output: false
    // Explanation: 1 cannot be popped before 2.
    fmt.Println(validateStackSequences([]int{1,2,3,4,5}, []int{4,3,5,1,2})) // false
}