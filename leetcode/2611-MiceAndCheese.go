package main

// 2611. Mice and Cheese
// There are two mice and n different types of cheese, each type of cheese should be eaten by exactly one mouse.

// A point of the cheese with index i (0-indexed) is:
//     reward1[i] if the first mouse eats it.
//     reward2[i] if the second mouse eats it.

// You are given a positive integer array reward1, a positive integer array reward2, and a non-negative integer k.

// Return the maximum points the mice can achieve if the first mouse eats exactly k types of cheese.

// Example 1:
// Input: reward1 = [1,1,3,4], reward2 = [4,4,1,1], k = 2
// Output: 15
// Explanation: In this example, the first mouse eats the 2nd (0-indexed) and the 3rd types of cheese, and the second mouse eats the 0th and the 1st types of cheese.
// The total points are 4 + 4 + 3 + 4 = 15.
// It can be proven that 15 is the maximum total points that the mice can achieve.

// Example 2:
// Input: reward1 = [1,1], reward2 = [1,1], k = 2
// Output: 2
// Explanation: In this example, the first mouse eats the 0th (0-indexed) and 1st types of cheese, and the second mouse does not eat any cheese.
// The total points are 1 + 1 = 2.
// It can be proven that 2 is the maximum total points that the mice can achieve.

// Constraints:
//     1 <= n == reward1.length == reward2.length <= 10^5
//     1 <= reward1[i], reward2[i] <= 1000
//     0 <= k <= n

import "fmt"
import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
    reward := make([][]int, len(reward1))
    for i := 0; i < len(reward1); i++ {
        reward[i] = []int{ reward1[i] - reward2[i], reward1[i], reward2[i] }
    }
    sort.Slice(reward, func(i, j int) bool {
        return reward[i][0] > reward[j][0]
    })
    res := 0
    for _, v := range reward[:k] {
        res += v[1]
    }
    for _, v := range reward[k:] {
        res += v[2]
    }
    return res
}

func miceAndCheese1(reward1, reward2 []int, k int) int {
    res, n := 0, len(reward1)
    for i, _ := range reward1 {
        reward1[i] -= reward2[i]
        res += reward2[i]
    }
    sort.Ints(reward1)
    for _, v := range reward1[n - k : n] {
        res += v
    }
    return res
}

func main() {
    // Example 1:
    // Input: reward1 = [1,1,3,4], reward2 = [4,4,1,1], k = 2
    // Output: 15
    // Explanation: In this example, the first mouse eats the 2nd (0-indexed) and the 3rd types of cheese, and the second mouse eats the 0th and the 1st types of cheese.
    // The total points are 4 + 4 + 3 + 4 = 15.
    // It can be proven that 15 is the maximum total points that the mice can achieve.
    fmt.Println(miceAndCheese([]int{1,1,3,4}, []int{4,4,1,1}, 2)) // 15
    // Example 2:
    // Input: reward1 = [1,1], reward2 = [1,1], k = 2
    // Output: 2
    // Explanation: In this example, the first mouse eats the 0th (0-indexed) and 1st types of cheese, and the second mouse does not eat any cheese.
    // The total points are 1 + 1 = 2.
    // It can be proven that 2 is the maximum total points that the mice can achieve.
    fmt.Println(miceAndCheese([]int{1,1}, []int{1,1}, 2)) // 2

    fmt.Println(miceAndCheese([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 59
    fmt.Println(miceAndCheese([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 59

    fmt.Println(miceAndCheese1([]int{1,1,3,4}, []int{4,4,1,1}, 2)) // 15
    fmt.Println(miceAndCheese1([]int{1,1}, []int{1,1}, 2)) // 2
    fmt.Println(miceAndCheese1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 59
    fmt.Println(miceAndCheese1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 59
}