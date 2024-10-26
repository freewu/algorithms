package main

// 1989. Maximum Number of People That Can Be Caught in Tag
// You are playing a game of tag with your friends. 
// In tag, people are divided into two teams: people who are "it", and people who are not "it". 
// The people who are "it" want to catch as many people as possible who are not "it".

// You are given a 0-indexed integer array team containing only zeros (denoting people who are not "it") 
// and ones (denoting people who are "it"), and an integer dist. 
// A person who is "it" at index i can catch any one person whose index is in the range [i - dist, i + dist] (inclusive) 
// and is not "it".

// Return the maximum number of people that the people who are "it" can catch.

// Example 1:
// Input: team = [0,1,0,1,0], dist = 3
// Output: 2
// Explanation:
// The person who is "it" at index 1 can catch people in the range [i-dist, i+dist] = [1-3, 1+3] = [-2, 4].
// They can catch the person who is not "it" at index 2.
// The person who is "it" at index 3 can catch people in the range [i-dist, i+dist] = [3-3, 3+3] = [0, 6].
// They can catch the person who is not "it" at index 0.
// The person who is not "it" at index 4 will not be caught because the people at indices 1 and 3 are already catching one person.

// Example 2:
// Input: team = [1], dist = 1
// Output: 0
// Explanation:
// There are no people who are not "it" to catch.

// Example 3:
// Input: team = [0], dist = 1
// Output: 0
// Explanation:
// There are no people who are "it" to catch people.

// Constraints:
//     1 <= team.length <= 10^5
//     0 <= team[i] <= 1
//     1 <= dist <= team.length

import "fmt"

// 双指针
func catchMaximumAmountofPeople(team []int, dist int) int {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, j := 0, 0; i < len(team) && j < len(team); i++ {
        if team[i] == 0 {
            j = max(j,i - dist)
            for ;j < len(team) && j <= i + dist; j++ {
                if team[j] == 1{
                    j++
                    res++
                    break
                }
            }
        }
    }
    return res
}

func catchMaximumAmountofPeople1(team []int, dist int) int {
    ptr0, res, n := 0, 0, len(team)
    for i := range team {
        if team[i] == 0 { continue }
        low, high := i - dist, i + dist
        if low < 0   { low = 0 }
        if high >= n { high = n - 1 }
        for ptr0 < n && (team[ptr0] != 0 || ptr0 < low) {
            ptr0++
        }
        if ptr0 <= high {
            res++
            ptr0++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: team = [0,1,0,1,0], dist = 3
    // Output: 2
    // Explanation:
    // The person who is "it" at index 1 can catch people in the range [i-dist, i+dist] = [1-3, 1+3] = [-2, 4].
    // They can catch the person who is not "it" at index 2.
    // The person who is "it" at index 3 can catch people in the range [i-dist, i+dist] = [3-3, 3+3] = [0, 6].
    // They can catch the person who is not "it" at index 0.
    // The person who is not "it" at index 4 will not be caught because the people at indices 1 and 3 are already catching one person.
    fmt.Println(catchMaximumAmountofPeople([]int{0,1,0,1,0}, 3)) // 2
    // Example 2:
    // Input: team = [1], dist = 1
    // Output: 0
    // Explanation:
    // There are no people who are not "it" to catch.
    fmt.Println(catchMaximumAmountofPeople([]int{1}, 1)) // 0
    // Example 3:
    // Input: team = [0], dist = 1
    // Output: 0
    // Explanation:
    // There are no people who are "it" to catch people.
    fmt.Println(catchMaximumAmountofPeople([]int{0}, 1)) // 0

    fmt.Println(catchMaximumAmountofPeople1([]int{0,1,0,1,0}, 3)) // 2
    fmt.Println(catchMaximumAmountofPeople1([]int{1}, 1)) // 0
    fmt.Println(catchMaximumAmountofPeople1([]int{0}, 1)) // 0
}