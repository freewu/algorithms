package main

// 717. 1-bit and 2-bit Characters
// We have two special characters:
//     The first character can be represented by one bit 0.
//     The second character can be represented by two bits (10 or 11).

// Given a binary array bits that ends with 0, return true if the last character must be a one-bit character.

// Example 1:
// Input: bits = [1,0,0]
// Output: true
// Explanation: The only way to decode it is two-bit character and one-bit character.
// So the last character is one-bit character.

// Example 2:
// Input: bits = [1,1,1,0]
// Output: false
// Explanation: The only way to decode it is two-bit character and two-bit character.
// So the last character is not one-bit character.

// Constraints:
//     1 <= bits.length <= 1000
//     bits[i] is either 0 or 1.

import "fmt"

func isOneBitCharacter(bits []int) bool {
    i, n := 0, len(bits)
    for i < n - 1 {
        if bits[i] == 1 {// 10, 11 占两个
            i += 2
        } else { // 0 // 占1个
            i++
        }
    }
    return i == n - 1
}

func main() {
    // Example 1:
    // Input: bits = [1,0,0]
    // Output: true
    // Explanation: The only way to decode it is two-bit character and one-bit character.
    // So the last character is one-bit character. (1, 0), (0)
    fmt.Println(isOneBitCharacter([]int{1,0,0}))
    // Example 2:
    // Input: bits = [1,1,1,0]
    // Output: false
    // Explanation: The only way to decode it is two-bit character and two-bit character.
    // So the last character is not one-bit character. (1, 1), (1, 0)
    fmt.Println(isOneBitCharacter([]int{1,1,1,0})) // false

    fmt.Println(isOneBitCharacter([]int{0,0,0,0,0,0,0,0,0,0})) // true
    fmt.Println(isOneBitCharacter([]int{1,1,1,1,1,1,1,1,1,1})) // false
    fmt.Println(isOneBitCharacter([]int{0,0,0,0,0,1,1,1,1,1})) // true
    fmt.Println(isOneBitCharacter([]int{1,1,1,1,1,0,0,0,0,0})) // true
    fmt.Println(isOneBitCharacter([]int{0,1,0,1,0,1,0,1,0,1})) // true
    fmt.Println(isOneBitCharacter([]int{1,0,1,0,1,0,1,0,1,0})) // false
}