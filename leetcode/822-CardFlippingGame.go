package main

// 822. Card Flipping Game
// You are given two 0-indexed integer arrays fronts and backs of length n, where the ith card has the positive integer fronts[i] printed on the front and backs[i] printed on the back. 
// Initially, each card is placed on a table such that the front number is facing up and the other is facing down. 
// You may flip over any number of cards (possibly zero).

// After flipping the cards, an integer is considered good if it is facing down on some card and not facing up on any card.
// Return the minimum possible good integer after flipping the cards. If there are no good integers, return 0.

// Example 1:
// Input: fronts = [1,2,4,4,7], backs = [1,3,4,1,3]
// Output: 2
// Explanation:
// If we flip the second card, the face up numbers are [1,3,4,4,7] and the face down are [1,2,4,1,3].
// 2 is the minimum good integer as it appears facing down but not facing up.
// It can be shown that 2 is the minimum possible good integer obtainable after flipping some cards.

// Example 2:
// Input: fronts = [1], backs = [1]
// Output: 0
// Explanation:
// There are no good integers no matter how we flip the cards, so we return 0.

// Constraints:
//     n == fronts.length == backs.length
//     1 <= n <= 1000
//     1 <= fronts[i], backs[i] <= 2000

import "fmt"

func flipgame(fronts []int, backs []int) int {
    res, set :=2001, make(map[int]struct{})
    for i := range fronts {
        if fronts[i] == backs[i] {
            set[fronts[i]] = struct{}{}
        }
    }
    for i := range fronts {
        if fronts[i] != backs[i] {
            mn, mx := fronts[i], backs[i]
            if fronts[i] > backs[i] {
                mn, mx = backs[i], fronts[i]
            }
            if _, ok := set[mn]; !ok && mn < res {
                res = mn
            } else if _, ok := set[mx]; !ok && mx < res {
                res = mx
            }
        }
    }
    if res == 2001 {
        return 0
    }
    return res
}

func flipgame1(fronts []int, backs []int) int {
    set := map[int]int{}
    for i := 0; i <len(fronts); i++ {
        if fronts[i] == backs[i] {
            set[fronts[i]] = 1
        }
    }
    mn := 99999
    for _, v := range fronts {
        if v < mn {
            if  _, ok := set[v]; !ok {
                mn = v
            }
        }
    }
    for _, v := range backs {
        if v < mn {
            if  _, ok := set[v]; !ok {
                mn = v
            }
        }
    }
    if mn > 1000 {
        return 0
    }
    return mn
}

func flipgame2(fronts []int, backs []int) int {
    ban := make([]int, 2001)
    for i := 0; i < len(fronts); i++  {
        if fronts[i] == backs[i] {
            ban[fronts[i]] = 1
            continue
        }
        if ban[fronts[i]] == 0 {
            ban[fronts[i]] = 2
        }
        if ban[backs[i]] == 0 {
            ban[backs[i]] = 2  
        } 
    }
    for i := 1; i < len(ban); i++ {
        if ban[i] == 2 {
            return i
        }
    }
    return 0
}

func main() {
    // Example 1:
    // Input: fronts = [1,2,4,4,7], backs = [1,3,4,1,3]
    // Output: 2
    // Explanation:
    // If we flip the second card, the face up numbers are [1,3,4,4,7] and the face down are [1,2,4,1,3].
    // 2 is the minimum good integer as it appears facing down but not facing up.
    // It can be shown that 2 is the minimum possible good integer obtainable after flipping some cards.
    fmt.Println(flipgame([]int{1,2,4,4,7}, []int{1,3,4,1,3})) // 2
    // Example 2:
    // Input: fronts = [1], backs = [1]
    // Output: 0
    // Explanation:
    // There are no good integers no matter how we flip the cards, so we return 0.
    fmt.Println(flipgame([]int{1}, []int{1})) // 0

    fmt.Println(flipgame1([]int{1,2,4,4,7}, []int{1,3,4,1,3})) // 2
    fmt.Println(flipgame1([]int{1}, []int{1})) // 0

    fmt.Println(flipgame2([]int{1,2,4,4,7}, []int{1,3,4,1,3})) // 2
    fmt.Println(flipgame2([]int{1}, []int{1})) // 0
}