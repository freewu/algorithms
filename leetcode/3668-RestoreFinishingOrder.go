package main

// 3668. Restore Finishing Order
// You are given an integer array order of length n and an integer array friends.
//     1. order contains every integer from 1 to n exactly once, 
//        representing the IDs of the participants of a race in their finishing order.
//     2. friends contains the IDs of your friends in the race sorted in strictly increasing order. 
//        Each ID in friends is guaranteed to appear in the order array.

// Return an array containing your friends' IDs in their finishing order.

// Example 1:
// Input: order = [3,1,2,5,4], friends = [1,3,4]
// Output: [3,1,4]
// Explanation:
// The finishing order is [3, 1, 2, 5, 4]. Therefore, the finishing order of your friends is [3, 1, 4].

// Example 2:
// Input: order = [1,4,5,3,2], friends = [2,5]
// Output: [5,2]
// Explanation:
// The finishing order is [1, 4, 5, 3, 2]. Therefore, the finishing order of your friends is [5, 2].

// Constraints:
//     1 <= n == order.length <= 100
//     order contains every integer from 1 to n exactly once
//     1 <= friends.length <= min(8, n)
//     1 <= friends[i] <= n
//     friends is strictly increasing

import "fmt"

func recoverOrder(order []int, friends []int) []int {
    set := make(map[int]bool)
    for _, v := range friends {
        set[v] = true
    }
    res := make([]int, 0)
    for _, v := range order {
        if set[v] {
            res = append(res, v)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: order = [3,1,2,5,4], friends = [1,3,4]
    // Output: [3,1,4]
    // Explanation:
    // The finishing order is [3, 1, 2, 5, 4]. Therefore, the finishing order of your friends is [3, 1, 4].
    fmt.Println(recoverOrder([]int{3,1,2,5,4}, []int{1,3,4})) // [3,1,4]
    // Example 2:
    // Input: order = [1,4,5,3,2], friends = [2,5]
    // Output: [5,2]
    // Explanation:
    // The finishing order is [1, 4, 5, 3, 2]. Therefore, the finishing order of your friends is [5, 2].
    fmt.Println(recoverOrder([]int{1,4,5,3,2}, []int{2,5})) // [5,2]

    fmt.Println(recoverOrder([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [1,2,3,4,5,6,7,8,9]
    fmt.Println(recoverOrder([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [1,2,3,4,5,6,7,8,9]
    fmt.Println(recoverOrder([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // [9 8 7 6 5 4 3 2 1]
    fmt.Println(recoverOrder([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // [9 8 7 6 5 4 3 2 1]
}