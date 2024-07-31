package main

// 825. Friends Of Appropriate Ages
// There are n persons on a social media website. 
// You are given an integer array ages where ages[i] is the age of the ith person.
// A Person x will not send a friend request to a person y (x != y) if any of the following conditions is true:
//     age[y] <= 0.5 * age[x] + 7
//     age[y] > age[x]
//     age[y] > 100 && age[x] < 100

// Otherwise, x will send a friend request to y.
// Note that if x sends a request to y, y will not necessarily send a request to x. 
// Also, a person will not send a friend request to themself.

// Return the total number of friend requests made.

// Example 1:
// Input: ages = [16,16]
// Output: 2
// Explanation: 2 people friend request each other.

// Example 2:
// Input: ages = [16,17,18]
// Output: 2
// Explanation: Friend requests are made 17 -> 16, 18 -> 17.

// Example 3:
// Input: ages = [20,30,100,110,120]
// Output: 3
// Explanation: Friend requests are made 110 -> 100, 120 -> 110, 120 -> 100.

// Constraints:
//     n == ages.length
//     1 <= n <= 2 * 10^4
//     1 <= ages[i] <= 120

import "fmt"

func numFriendRequests(ages []int) int {
    people := make([]int, 121)
    for _, age := range ages {
        people[age]++
    }
    res := 0
    for i := 1; i <= 120; i++ {
        if people[i] == 0 { continue }
        limit := (i / 2) + 7
        for j := limit + 1; j < i; j++ {
            res += people[i] * people[j]
        }
        if i > limit {
            res += people[i] * (people[i] - 1)
        }
    }
    return res
}

func numFriendRequests1(ages []int) int {
    maxCnt := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, age := range ages { // 找出最大的年龄
        maxCnt = max(maxCnt, age)
    }
    presums := make([]int, maxCnt + 1)
    for _, age := range ages {
        presums[age] += 1
    }
    for i := 1; i <= maxCnt; i++ {
        presums[i] += presums[i - 1]
    }
    res := 0
    for _, age := range ages {
        if age > age / 2 + 7 { // ages[y] > 0.5 * ages[x] + 7
            res += presums[age] - presums[age / 2 + 7] - 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: ages = [16,16]
    // Output: 2
    // Explanation: 2 people friend request each other.
    fmt.Println(numFriendRequests([]int{16,16})) // 2
    // Example 2:
    // Input: ages = [16,17,18]
    // Output: 2
    // Explanation: Friend requests are made 17 -> 16, 18 -> 17.
    fmt.Println(numFriendRequests([]int{16,17,18})) // 2
    // Example 3:
    // Input: ages = [20,30,100,110,120]
    // Output: 3
    // Explanation: Friend requests are made 110 -> 100, 120 -> 110, 120 -> 100.
    fmt.Println(numFriendRequests([]int{20,30,100,110,120})) // 3

    fmt.Println(numFriendRequests1([]int{16,16})) // 2
    fmt.Println(numFriendRequests1([]int{16,17,18})) // 2
    fmt.Println(numFriendRequests1([]int{20,30,100,110,120})) // 3
}