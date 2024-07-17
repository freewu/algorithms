package main

// 1014. Best Sightseeing Pair
// You are given an integer array values where values[i] represents the value of the ith sightseeing spot. 
// Two sightseeing spots i and j have a distance j - i between them.

// The score of a pair (i < j) of sightseeing spots is values[i] + values[j] + i - j: 
// the sum of the values of the sightseeing spots, minus the distance between them.

// Return the maximum score of a pair of sightseeing spots.

// Example 1:
// Input: values = [8,1,5,2,6]
// Output: 11
// Explanation: i = 0, j = 2, values[i] + values[j] + i - j = 8 + 5 + 0 - 2 = 11

// Example 2:
// Input: values = [1,2]
// Output: 2

// Constraints:
//     2 <= values.length <= 5 * 10^4
//     1 <= values[i] <= 1000

import "fmt"

func maxScoreSightseeingPair(values []int) int {
    res, score := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range values {
        res = max(res, score + v)
        score = max(score, v) - 1
    }
    return res
}

func maxScoreSightseeingPair1(values []int) int {
    PreMax, res := values[0]-1, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for j := 1; j < len(values); j++ {
        res = max(res,values[j] + PreMax)
        PreMax--
        if values[j]-1 > PreMax {
            PreMax = values[j]-1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: values = [8,1,5,2,6]
    // Output: 11
    // Explanation: i = 0, j = 2, values[i] + values[j] + i - j = 8 + 5 + 0 - 2 = 11
    fmt.Println(maxScoreSightseeingPair([]int{8,1,5,2,6})) // 11
    // Example 2:
    // Input: values = [1,2]
    // Output: 2
    fmt.Println(maxScoreSightseeingPair([]int{1,2})) // 2

    fmt.Println(maxScoreSightseeingPair1([]int{8,1,5,2,6})) // 11
    fmt.Println(maxScoreSightseeingPair1([]int{1,2})) // 2
}