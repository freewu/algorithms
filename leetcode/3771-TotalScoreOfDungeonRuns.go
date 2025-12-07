package main

// 3771. Total Score of Dungeon Runs
// You are given a positive integer hp and two positive 1-indexed integer arrays damage and requirement.

// There is a dungeon with n trap rooms numbered from 1 to n. 
// Entering room i reduces your health points by damage[i]. 
// After that reduction, if your remaining health points are at least requirement[i], you earn 1 point for that room.

// Let score(j) be the number of points you get if you start with hp health points and enter the rooms j, j + 1, ..., n in this order.

// Return the integer score(1) + score(2) + ... + score(n), the sum of scores over all starting rooms.

// Note: You cannot skip rooms. You can finish your journey even if your health points become non-positive.

// Example 1:
// Input: hp = 11, damage = [3,6,7], requirement = [4,2,5]
// Output: 3
// Explanation:
// score(1) = 2, score(2) = 1, score(3) = 0. The total score is 2 + 1 + 0 = 3.
// As an example, score(1) = 2 because you get 2 points if you start from room 1.
// You start with 11 health points.
// Enter room 1. Your health points are now 11 - 3 = 8. You get 1 point because 8 >= 4.
// Enter room 2. Your health points are now 8 - 6 = 2. You get 1 point because 2 >= 2.
// Enter room 3. Your health points are now 2 - 7 = -5. You do not get any points because -5 < 5.

// Example 2:
// Input: hp = 2, damage = [10000,1], requirement = [1,1]
// Output: 1
// Explanation:
// score(1) = 0, score(2) = 1. The total score is 0 + 1 = 1.
// score(1) = 0 because you do not get any points if you start from room 1.
// You start with 2 health points.
// Enter room 1. Your health points are now 2 - 10000 = -9998. You do not get any points because -9998 < 1.
// Enter room 2. Your health points are now -9998 - 1 = -9999. You do not get any points because -9999 < 1.
// score(2) = 1 because you get 1 point if you start from room 2.
// You start with 2 health points.
// Enter room 2. Your health points are now 2 - 1 = 1. You get 1 point because 1 >= 1.

// Constraints:
//     1 <= hp <= 10^9
//     1 <= n == damage.length == requirement.length <= 10^5
//     1 <= damage[i], requirement[i] <= 10^4

import "fmt"
import "sort"

func totalScore(hp int, damage []int, requirement []int) int64 {
    res, prefix := 0, make([]int, len(damage)+1)
    for i, req := range requirement {
        prefix[i+1] = prefix[i] + damage[i]
        low := prefix[i+1] + req - hp
        j := sort.SearchInts(prefix[:i+1], low)
        res += (i - j + 1)
    }
    return int64(res)
}

func totalScore1(hp int, damage []int, requirement []int) int64 {
    res, n := 0, len(damage)
    if n == 0 { return 0 }
    prefix := make([]int, n+1)
    for i := 1; i <= n; i++ {
        prefix[i] = prefix[i-1] + damage[i-1]
    }
    for i := 1; i <= n; i++ {
        threshold := prefix[i] + requirement[i-1] - hp
        if threshold <= 0 {
            res += i
            continue
        }
        if threshold > prefix[i-1] { continue }
        index := sort.Search(i, func(k int) bool { return prefix[k] >= threshold })
        res += (i - index)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: hp = 11, damage = [3,6,7], requirement = [4,2,5]
    // Output: 3
    // Explanation:
    // score(1) = 2, score(2) = 1, score(3) = 0. The total score is 2 + 1 + 0 = 3.
    // As an example, score(1) = 2 because you get 2 points if you start from room 1.
    // You start with 11 health points.
    // Enter room 1. Your health points are now 11 - 3 = 8. You get 1 point because 8 >= 4.
    // Enter room 2. Your health points are now 8 - 6 = 2. You get 1 point because 2 >= 2.
    // Enter room 3. Your health points are now 2 - 7 = -5. You do not get any points because -5 < 5.
    fmt.Println(totalScore(11, []int{3,6,7}, []int{4,2,5})) // 3
    // Example 2:
    // Input: hp = 2, damage = [10000,1], requirement = [1,1]
    // Output: 1
    // Explanation:
    // score(1) = 0, score(2) = 1. The total score is 0 + 1 = 1.
    // score(1) = 0 because you do not get any points if you start from room 1.
    // You start with 2 health points.
    // Enter room 1. Your health points are now 2 - 10000 = -9998. You do not get any points because -9998 < 1.
    // Enter room 2. Your health points are now -9998 - 1 = -9999. You do not get any points because -9999 < 1.
    // score(2) = 1 because you get 1 point if you start from room 2.
    // You start with 2 health points.
    // Enter room 2. Your health points are now 2 - 1 = 1. You get 1 point because 1 >= 1.
    fmt.Println(totalScore(2, []int{10000,1}, []int{1,1})) // 1

    fmt.Println(totalScore(2, []int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(totalScore(2, []int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(totalScore(2, []int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(totalScore(2, []int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 1

    fmt.Println(totalScore1(11, []int{3,6,7}, []int{4,2,5})) // 3
    fmt.Println(totalScore1(2, []int{10000,1}, []int{1,1})) // 1
    fmt.Println(totalScore1(2, []int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(totalScore1(2, []int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(totalScore1(2, []int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(totalScore1(2, []int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 1
}