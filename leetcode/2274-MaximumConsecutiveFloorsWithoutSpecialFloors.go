package main

// 2274. Maximum Consecutive Floors Without Special Floors
// Alice manages a company and has rented some floors of a building as office space. 
// Alice has decided some of these floors should be special floors, used for relaxation only.

// You are given two integers bottom and top, which denote that Alice has rented all the floors from bottom to top (inclusive). 
// You are also given the integer array special, where special[i] denotes a special floor that Alice has designated for relaxation.

// Return the maximum number of consecutive floors without a special floor.

// Example 1:
// Input: bottom = 2, top = 9, special = [4,6]
// Output: 3
// Explanation: The following are the ranges (inclusive) of consecutive floors without a special floor:
// - (2, 3) with a total amount of 2 floors.
// - (5, 5) with a total amount of 1 floor.
// - (7, 9) with a total amount of 3 floors.
// Therefore, we return the maximum number which is 3 floors.

// Example 2:
// Input: bottom = 6, top = 8, special = [7,6,8]
// Output: 0
// Explanation: Every floor rented is a special floor, so we return 0.

// Constraints:
//     1 <= special.length <= 10^5
//     1 <= bottom <= special[i] <= top <= 10^9
//     All the values of special are unique.

import "fmt"
import "sort"

func maxConsecutive(bottom int, top int, special []int) int {
    sort.Ints(special)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := max(special[0] - bottom, top - special[len(special)-1])
    for i, v := range special[1:] {
        res = max(res, v - special[i] - 1)
    }
    return res
}

func maxConsecutive1(bottom int, top int, special []int) int {
    special = append(special, bottom - 1)
    special = append(special, top + 1)
    sort.Ints(special)
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(special) - 1; i++ {
        res = max(res, special[i + 1] - special[i] - 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: bottom = 2, top = 9, special = [4,6]
    // Output: 3
    // Explanation: The following are the ranges (inclusive) of consecutive floors without a special floor:
    // - (2, 3) with a total amount of 2 floors.
    // - (5, 5) with a total amount of 1 floor.
    // - (7, 9) with a total amount of 3 floors.
    // Therefore, we return the maximum number which is 3 floors.
    fmt.Println(maxConsecutive(2, 9, []int{4,6})) // 3
    // Example 2:
    // Input: bottom = 6, top = 8, special = [7,6,8]
    // Output: 0
    // Explanation: Every floor rented is a special floor, so we return 0.
    fmt.Println(maxConsecutive(6, 8, []int{7,6,8})) // 0

    fmt.Println(maxConsecutive1(2, 9, []int{4,6})) // 3
    fmt.Println(maxConsecutive1(6, 8, []int{7,6,8})) // 0
}