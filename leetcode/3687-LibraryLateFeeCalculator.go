package main

// 3687. Library Late Fee Calculator
// You are given an integer array daysLate where daysLate[i] indicates how many days late the ith book was returned.

// The penalty is calculated as follows:
//     If daysLate[i] == 1, penalty is 1.
//     If 2 <= daysLate[i] <= 5, penalty is 2 * daysLate[i].
//     If daysLate[i] > 5, penalty is 3 * daysLate[i].

// Return the total penalty for all books.

// Example 1:
// Input: daysLate = [5,1,7]
// Output: 32
// Explanation:
// daysLate[0] = 5: Penalty is 2 * daysLate[0] = 2 * 5 = 10.
// daysLate[1] = 1: Penalty is 1.
// daysLate[2] = 7: Penalty is 3 * daysLate[2] = 3 * 7 = 21.
// Thus, the total penalty is 10 + 1 + 21 = 32.

// Example 2:
// Input: daysLate = [1,1]
// Output: 2
// Explanation:
// daysLate[0] = 1: Penalty is 1.
// daysLate[1] = 1: Penalty is 1.
// Thus, the total penalty is 1 + 1 = 2.

// Constraints:
//     1 <= daysLate.length <= 100
//     1 <= daysLate[i] <= 100

import "fmt"

func lateFee(daysLate []int) int {
    res := 0
    for _,v := range daysLate {
        if v == 1 { //  If daysLate[i] == 1, penalty is 1.
            res++
        } else if v <= 5 { // If 2 <= daysLate[i] <= 5, penalty is 2 * daysLate[i].
            res += v * 2
        } else { //  If daysLate[i] > 5, penalty is 3 * daysLate[i].
            res += v * 3
        }
    }
    return res
}


func main() {
    // Example 1:
    // Input: daysLate = [5,1,7]
    // Output: 32
    // Explanation:
    // daysLate[0] = 5: Penalty is 2 * daysLate[0] = 2 * 5 = 10.
    // daysLate[1] = 1: Penalty is 1.
    // daysLate[2] = 7: Penalty is 3 * daysLate[2] = 3 * 7 = 21.
    // Thus, the total penalty is 10 + 1 + 21 = 32.
    fmt.Println(lateFee([]int{5,1,7})) // 32
    // Example 2:
    // Input: daysLate = [1,1]
    // Output: 2
    // Explanation:
    // daysLate[0] = 1: Penalty is 1.
    // daysLate[1] = 1: Penalty is 1.
    // Thus, the total penalty is 1 + 1 = 2.
    fmt.Println(lateFee([]int{1,1})) // 2

    fmt.Println(lateFee([]int{1,2,3,4,5,6,7,8,9})) // 119
    fmt.Println(lateFee([]int{9,8,7,6,5,4,3,2,1})) // 119
}