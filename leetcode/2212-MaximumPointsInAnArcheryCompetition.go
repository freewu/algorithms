package main

// 2212. Maximum Points in an Archery Competition
// Alice and Bob are opponents in an archery competition. 
// The competition has set the following rules:
//     1. Alice first shoots numArrows arrows and then Bob shoots numArrows arrows.
//     2. The points are then calculated as follows:
//         2.1 The target has integer scoring sections ranging from 0 to 11 inclusive.
//         2.2 For each section of the target with score k (in between 0 to 11), 
//             say Alice and Bob have shot ak and bk arrows on that section respectively. 
//             If ak >= bk, then Alice takes k points. If ak < bk, then Bob takes k points.
//         2.3 However, if ak == bk == 0, then nobody takes k points.

// For example, if Alice and Bob both shot 2 arrows on the section with score 11, then Alice takes 11 points. 
// On the other hand, if Alice shot 0 arrows on the section with score 11 and Bob shot 2 arrows on that same section, then Bob takes 11 points.

// You are given the integer numArrows and an integer array aliceArrows of size 12, which represents the number of arrows Alice shot on each scoring section from 0 to 11. 
// Now, Bob wants to maximize the total number of points he can obtain.

// Return the array bobArrows which represents the number of arrows Bob shot on each scoring section from 0 to 11. 
// The sum of the values in bobArrows should equal numArrows.

// If there are multiple ways for Bob to earn the maximum total points, return any one of them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/02/24/ex1.jpg" />
// Input: numArrows = 9, aliceArrows = [1,1,0,1,0,0,2,1,0,1,2,0]
// Output: [0,0,0,0,1,1,0,0,1,2,3,1]
// Explanation: The table above shows how the competition is scored. 
// Bob earns a total point of 4 + 5 + 8 + 9 + 10 + 11 = 47.
// It can be shown that Bob cannot obtain a score higher than 47 points.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/02/24/ex2new.jpg" />
// Input: numArrows = 3, aliceArrows = [0,0,1,0,0,0,0,0,0,0,0,2]
// Output: [0,0,0,0,0,0,0,0,1,1,1,0]
// Explanation: The table above shows how the competition is scored.
// Bob earns a total point of 8 + 9 + 10 = 27.
// It can be shown that Bob cannot obtain a score higher than 27 points.

// Constraints:
//     1 <= numArrows <= 10^5
//     aliceArrows.length == bobArrows.length == 12
//     0 <= aliceArrows[i], bobArrows[i] <= numArrows
//     sum(aliceArrows[i]) == numArrows

import "fmt"

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
    n, state, mx := len(aliceArrows), 0, -1
    for mask := 1; mask < 1 << n; mask++ {
        count, points := 0, 0
        for i, alice := range aliceArrows {
            if (mask >> i) & 1 == 1 {
                count += alice + 1
                points += i
            }
        }
        if count <= numArrows && mx < points {
            state, mx = mask, points
        }
    }
    res := make([]int, n)
    for i, alice := range aliceArrows {
        if (state >> i) & 1 == 1 {
            res[i] = alice + 1
            numArrows -= res[i]
        }
    }
    res[0] += numArrows
    return res
}

func maximumBobPoints1(numArrows int, aliceArrows []int) []int {
    res, bestScore := make([]int, len(aliceArrows)), 0
    var backtracking func(i int, remainingArrows int, score int, bobArrows []int)
    backtracking = func(i, remainingArrows, score int, bobArrows []int) {
        if i == len(aliceArrows) {
            if score > bestScore {
                bestScore = score
                res = append([]int{}, bobArrows...)
            }
            return
        }
        backtracking(i + 1, remainingArrows, score, bobArrows)
        if aliceArrows[i] + 1 <= remainingArrows {
            bobArrows[i] = aliceArrows[i] + 1
            backtracking(i + 1, remainingArrows - bobArrows[i], score + i, bobArrows)
            bobArrows[i] = 0
        }
    }
    sum := func(nums []int) int {
        res := 0
        for _, v := range nums {
            res += v
        }
        return res
    }
    backtracking(0, numArrows, 0, make([]int, len(aliceArrows)))
    res[0] += numArrows - sum(res)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/02/24/ex1.jpg" />
    // Input: numArrows = 9, aliceArrows = [1,1,0,1,0,0,2,1,0,1,2,0]
    // Output: [0,0,0,0,1,1,0,0,1,2,3,1]
    // Explanation: The table above shows how the competition is scored. 
    // Bob earns a total point of 4 + 5 + 8 + 9 + 10 + 11 = 47.
    // It can be shown that Bob cannot obtain a score higher than 47 points.
    fmt.Println(maximumBobPoints(9, []int{1,1,0,1,0,0,2,1,0,1,2,0})) // [0,0,0,0,1,1,0,0,1,2,3,1]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/02/24/ex2new.jpg" />
    // Input: numArrows = 3, aliceArrows = [0,0,1,0,0,0,0,0,0,0,0,2]
    // Output: [0,0,0,0,0,0,0,0,1,1,1,0]
    // Explanation: The table above shows how the competition is scored.
    // Bob earns a total point of 8 + 9 + 10 = 27.
    // It can be shown that Bob cannot obtain a score higher than 27 points.
    fmt.Println(maximumBobPoints(3, []int{0,0,1,0,0,0,0,0,0,0,0,2})) // [0,0,0,0,0,0,0,0,1,1,1,0]

    fmt.Println(maximumBobPoints1(9, []int{1,1,0,1,0,0,2,1,0,1,2,0})) // [0,0,0,0,1,1,0,0,1,2,3,1]
    fmt.Println(maximumBobPoints1(3, []int{0,0,1,0,0,0,0,0,0,0,0,2})) // [0,0,0,0,0,0,0,0,1,1,1,0]
}