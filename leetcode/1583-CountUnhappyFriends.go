package main

// 1583. Count Unhappy Friends
// You are given a list of preferences for n friends, where n is always even.

// For each person i, preferences[i] contains a list of friends sorted in the order of preference. 
// In other words, a friend earlier in the list is more preferred than a friend later in the list. 
// Friends in each list are denoted by integers from 0 to n-1.

// All the friends are divided into pairs. 
// The pairings are given in a list pairs, where pairs[i] = [xi, yi] denotes xi is paired with yi and yi is paired with xi.

// However, this pairing may cause some of the friends to be unhappy. 
// A friend x is unhappy if x is paired with y and there exists a friend u who is paired with v but:
//     x prefers u over y, and
//     u prefers x over v.

// Return the number of unhappy friends.

// Example 1:
// Input: n = 4, preferences = [[1, 2, 3], [3, 2, 0], [3, 1, 0], [1, 2, 0]], pairs = [[0, 1], [2, 3]]
// Output: 2
// Explanation:
// Friend 1 is unhappy because:
// - 1 is paired with 0 but prefers 3 over 0, and
// - 3 prefers 1 over 2.
// Friend 3 is unhappy because:
// - 3 is paired with 2 but prefers 1 over 2, and
// - 1 prefers 3 over 0.
// Friends 0 and 2 are happy.

// Example 2:
// Input: n = 2, preferences = [[1], [0]], pairs = [[1, 0]]
// Output: 0
// Explanation: Both friends 0 and 1 are happy.

// Example 3:
// Input: n = 4, preferences = [[1, 3, 2], [2, 3, 0], [1, 3, 0], [0, 2, 1]], pairs = [[1, 3], [0, 2]]
// Output: 4

// Constraints:
//     2 <= n <= 500
//     n is even.
//     preferences.length == n
//     preferences[i].length == n - 1
//     0 <= preferences[i][j] <= n - 1
//     preferences[i] does not contain i.
//     All values in preferences[i] are unique.
//     pairs.length == n/2
//     pairs[i].length == 2
//     xi != yi
//     0 <= xi, yi <= n - 1
//     Each person is contained in exactly one pair.

import "fmt"

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {   
    res, mp := 0, make(map[int]int) // Let's make a pair map for easier access of pairs
    for _, v := range pairs {
        mp[v[0]], mp[v[1]] = v[1], v[0]
    }
    better := make([]map[int]bool, n) // Let's compute the better candidates for each person than their current pair
    for i := 0; i < n; i++ {
        better[i] = make(map[int]bool)
        for j := 0; j < len(preferences[i]) && preferences[i][j] != mp[i]; j++ {
            better[i][preferences[i][j]] = true
        }
    }
    // Now we'll go through everyone and find if any of them are unhappy
    for i := 0; i < n; i++ {
        for b := range better[i] {
            // i prefers b over his current pair and b prefers i over his current pair
            if better[b][i] {
                res++
                break
            }     
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, preferences = [[1, 2, 3], [3, 2, 0], [3, 1, 0], [1, 2, 0]], pairs = [[0, 1], [2, 3]]
    // Output: 2
    // Explanation:
    // Friend 1 is unhappy because:
    // - 1 is paired with 0 but prefers 3 over 0, and
    // - 3 prefers 1 over 2.
    // Friend 3 is unhappy because:
    // - 3 is paired with 2 but prefers 1 over 2, and
    // - 1 prefers 3 over 0.
    // Friends 0 and 2 are happy.
    fmt.Println(unhappyFriends(4, [][]int{{1, 2, 3}, {3, 2, 0}, {3, 1, 0}, {1, 2, 0}}, [][]int{{0, 1}, {2, 3}})) // 2
    // Example 2:
    // Input: n = 2, preferences = [[1], [0]], pairs = [[1, 0]]
    // Output: 0
    // Explanation: Both friends 0 and 1 are happy.
    fmt.Println(unhappyFriends(2, [][]int{{1}, {0}}, [][]int{{1, 0}})) // 0
    // Example 3:
    // Input: n = 4, preferences = [[1, 3, 2], [2, 3, 0], [1, 3, 0], [0, 2, 1]], pairs = [[1, 3], [0, 2]]
    // Output: 4
    fmt.Println(unhappyFriends(4, [][]int{{1, 3, 2}, {2, 3, 0}, {1, 3, 0}, {0, 2, 1}}, [][]int{{1, 3}, {0, 2}})) // 4
}