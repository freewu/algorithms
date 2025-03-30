package main

// 3502. Minimum Cost to Reach Every Position
// You are given an integer array cost of size n. 
// You are currently at position n (at the end of the line) in a line of n + 1 people (numbered from 0 to n).

// You wish to move forward in the line, but each person in front of you charges a specific amount to swap places. 
// The cost to swap with person i is given by cost[i].

// You are allowed to swap places with people as follows:
//     1. If they are in front of you, you must pay them cost[i] to swap with them.
//     2. If they are behind you, they can swap with you for free.

// Return an array answer of size n, where answer[i] is the minimum total cost to reach each position i in the line.

// Example 1:
// Input: cost = [5,3,4,1,3,2]
// Output: [5,3,3,1,1,1]
// Explanation:
// We can get to each position in the following way:
// i = 0. We can swap with person 0 for a cost of 5.
// i = 1. We can swap with person 1 for a cost of 3.
// i = 2. We can swap with person 1 for a cost of 3, then swap with person 2 for free.
// i = 3. We can swap with person 3 for a cost of 1.
// i = 4. We can swap with person 3 for a cost of 1, then swap with person 4 for free.
// i = 5. We can swap with person 3 for a cost of 1, then swap with person 5 for free.

// Example 2:
// Input: cost = [1,2,4,6,7]
// Output: [1,1,1,1,1]
// Explanation:
// We can swap with person 0 for a cost of 1, then we will be able to reach any position i for free.

// Constraints:
//     1 <= n == cost.length <= 100
//     1 <= cost[i] <= 100

import "fmt"

func minCosts(cost []int) []int {
    n := len(cost)
    res := make([]int, n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < n; i++ {
        res[i] = cost[i]
        if i > 0 {
            res[i] = min(res[i], res[i - 1])
        }
    }
    return res
}

func minCosts1(cost []int) []int {
    mn := 1 << 31
    for i := range cost {
        if cost[i] < mn {
            mn = cost[i]
        }
        cost[i] = mn
    }
    return cost
}

func main() {
    // Example 1:
    // Input: cost = [5,3,4,1,3,2]
    // Output: [5,3,3,1,1,1]
    // Explanation:
    // We can get to each position in the following way:
    // i = 0. We can swap with person 0 for a cost of 5.
    // i = 1. We can swap with person 1 for a cost of 3.
    // i = 2. We can swap with person 1 for a cost of 3, then swap with person 2 for free.
    // i = 3. We can swap with person 3 for a cost of 1.
    // i = 4. We can swap with person 3 for a cost of 1, then swap with person 4 for free.
    // i = 5. We can swap with person 3 for a cost of 1, then swap with person 5 for free.
    fmt.Println(minCosts([]int{5,3,4,1,3,2})) // [5,3,3,1,1,1]
    // Example 2:
    // Input: cost = [1,2,4,6,7]
    // Output: [1,1,1,1,1]
    // Explanation:
    // We can swap with person 0 for a cost of 1, then we will be able to reach any position i for free.
    fmt.Println(minCosts([]int{1,2,4,6,7})) // [1,1,1,1,1]

    fmt.Println(minCosts([]int{1,2,3,4,5,6,7,8,9})) // [1 1 1 1 1 1 1 1 1]
    fmt.Println(minCosts([]int{9,8,7,6,5,4,3,2,1})) // [9 8 7 6 5 4 3 2 1]

    fmt.Println(minCosts1([]int{5,3,4,1,3,2})) // [5,3,3,1,1,1]
    fmt.Println(minCosts1([]int{1,2,4,6,7})) // [1,1,1,1,1]
    fmt.Println(minCosts1([]int{1,2,3,4,5,6,7,8,9})) // [1 1 1 1 1 1 1 1 1]
    fmt.Println(minCosts1([]int{9,8,7,6,5,4,3,2,1})) // [9 8 7 6 5 4 3 2 1]
}