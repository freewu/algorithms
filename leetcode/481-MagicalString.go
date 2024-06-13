package main

// 481. Magical String
// A magical string s consists of only '1' and '2' and obeys the following rules:
//     The string s is magical because concatenating the number of contiguous occurrences of characters '1' and '2' generates the string s itself.

// The first few elements of s is s = "1221121221221121122……". 
// If we group the consecutive 1's and 2's in s, it will be "1 22 11 2 1 22 1 22 11 2 11 22 ......" 
// and the occurrences of 1's or 2's in each group are "1 2 2 1 1 2 1 2 2 1 2 2 ......". 
// You can see that the occurrence sequence is s itself.

// Given an integer n, return the number of 1's in the first n number in the magical string s.

// Example 1:
// Input: n = 6
// Output: 3
// Explanation: The first 6 elements of magical string s is "122112" and it contains three 1's, so return 3.

// Example 2:
// Input: n = 1
// Output: 1

// Constraints:
//     1 <= n <= 10^5

import "fmt"

func magicalString(n int) int {
    if n < 4 { // 122
        return 1
    }
    nums := make([]int, n + 1) // // Create slice of size 3 with storage n + 1 to prevent further reallocations
    // Set only third element, because for real we don't need others, as starting with pos == 2
    nums[2] = 2 // Next number to add. Will be changed between 1 and 2
    number2Add, pos, res, insertPos := 1, 2, 1, 3
    for insertPos < n {        
        for i := 0; i < nums[pos]; i++ {
            nums[insertPos] = number2Add
            insertPos++
        }
        if number2Add == 1 {
            res += nums[pos]
            number2Add = 2
        } else {
            number2Add = 1
        }
        pos++
    }
    if nums[n] == 1 { // If last added number(which is n + 1, counting from 0), substract excessive 1 from counter
        res--
    }    
    return res
}

func main() {
    // Example 1:
    // Input: n = 6
    // Output: 3
    // Explanation: The first 6 elements of magical string s is "122112" and it contains three 1's, so return 3.
    fmt.Println(magicalString(6)) // 3 122112
    // Example 2:
    // Input: n = 1
    // Output: 1
    fmt.Println(magicalString(1)) // 1

    fmt.Println(magicalString(1024)) // 513
}