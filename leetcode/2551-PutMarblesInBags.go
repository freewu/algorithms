package main

// 2551. Put Marbles in Bags
// You have k bags. You are given a 0-indexed integer array weights where weights[i] is the weight of the ith marble. 
// You are also given the integer k.

// Divide the marbles into the k bags according to the following rules:
//     1. No bag is empty.
//     2. If the ith marble and jth marble are in a bag, 
//        then all marbles with an index between the ith and jth indices should also be in that same bag.
//     3. If a bag consists of all the marbles with an index from i to j inclusively, 
//        then the cost of the bag is weights[i] + weights[j].

// The score after distributing the marbles is the sum of the costs of all the k bags.

// Return the difference between the maximum and minimum scores among marble distributions.

// Example 1:
// Input: weights = [1,3,5,1], k = 2
// Output: 4
// Explanation: 
// The distribution [1],[3,5,1] results in the minimal score of (1+1) + (3+1) = 6. 
// The distribution [1,3],[5,1], results in the maximal score of (1+3) + (5+1) = 10. 
// Thus, we return their difference 10 - 6 = 4.

// Example 2:
// Input: weights = [1, 3], k = 2
// Output: 0
// Explanation: The only distribution possible is [1],[3]. 
// Since both the maximal and minimal score are the same, we return 0.

// Constraints:
//     1 <= k <= weights.length <= 10^5
//     1 <= weights[i] <= 10^9

import "fmt"
import "sort"

func putMarbles(weights []int, k int) int64 {
    if k == 1 { return int64(0) }
    mx, mn, n := 0, 0, len(weights)
    arr := make([]int, n - 1)
    for i := 0; i < n - 1; i++ {
        arr[i] = weights[i] + weights[i+1]
    }
    sort.Ints(arr)
    for i := 0; i < k - 1; i++ {
        mn += arr[i] 
    }
    for i := n - 2; i > n - k - 1; i-- {
        mx += arr[i] 
    }
    return int64(mx - mn)
}

func putMarbles1(weights []int, k int) int64 {
    mx, mn, n := 0, 0, len(weights)
    for i := 0; i < n - 1; i++ {
        weights[i] += weights[i+1]
    }
    weights = weights[:n-1]
    sort.Ints(weights)
    for i := 0; i < k-1; i++ {
        mn += weights[i]
        mx += weights[n - i - 2]
    }
    return int64(mx - mn)
}

func putMarbles2(weights []int, k int) int64 {
    res, n := 0, len(weights)
    arr := make([]int, n - 1)
    for i, v := range weights[:n-1] {
        arr[i] = v + weights[i+1]
    }
    sort.Ints(arr)
    for i := 0; i < k - 1; i++ {
        res += (arr[n-2-i] - arr[i])
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: weights = [1,3,5,1], k = 2
    // Output: 4
    // Explanation: 
    // The distribution [1],[3,5,1] results in the minimal score of (1+1) + (3+1) = 6. 
    // The distribution [1,3],[5,1], results in the maximal score of (1+3) + (5+1) = 10. 
    // Thus, we return their difference 10 - 6 = 4.
    fmt.Println(putMarbles([]int{1,3,5,1}, 2)) // 4
    // Example 2:
    // Input: weights = [1, 3], k = 2
    // Output: 0
    // Explanation: The only distribution possible is [1],[3]. 
    // Since both the maximal and minimal score are the same, we return 0.
    fmt.Println(putMarbles([]int{1,3}, 2)) // 0

    fmt.Println(putMarbles([]int{1,2,3,4,5,6,7,8,9}, 2)) // 14
    fmt.Println(putMarbles([]int{9,8,7,6,5,4,3,2,1}, 2)) // 14

    fmt.Println(putMarbles1([]int{1,3,5,1}, 2)) // 4
    fmt.Println(putMarbles1([]int{1,3}, 2)) // 0
    fmt.Println(putMarbles1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 14
    fmt.Println(putMarbles1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 14

    fmt.Println(putMarbles2([]int{1,3,5,1}, 2)) // 4
    fmt.Println(putMarbles2([]int{1,3}, 2)) // 0
    fmt.Println(putMarbles2([]int{1,2,3,4,5,6,7,8,9}, 2)) // 14
    fmt.Println(putMarbles2([]int{9,8,7,6,5,4,3,2,1}, 2)) // 14
}