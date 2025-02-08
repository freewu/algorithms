package main

// 2857. Count Pairs of Points With Distance k
// You are given a 2D integer array coordinates and an integer k, 
// where coordinates[i] = [xi, yi] are the coordinates of the ith point in a 2D plane.

// We define the distance between two points (x1, y1) and (x2, y2) as (x1 XOR x2) + (y1 XOR y2) where XOR is the bitwise XOR operation.

// Return the number of pairs (i, j) such that i < j and the distance between points i and j is equal to k.

// Example 1:
// Input: coordinates = [[1,2],[4,2],[1,3],[5,2]], k = 5
// Output: 2
// Explanation: We can choose the following pairs:
// - (0,1): Because we have (1 XOR 4) + (2 XOR 2) = 5.
// - (2,3): Because we have (1 XOR 5) + (3 XOR 2) = 5.

// Example 2:
// Input: coordinates = [[1,3],[1,3],[1,3],[1,3],[1,3]], k = 0
// Output: 10
// Explanation: Any two chosen pairs will have a distance of 0. There are 10 ways to choose two pairs.

// Constraints:
//     2 <= coordinates.length <= 50000
//     0 <= xi, yi <= 10^6
//     0 <= k <= 100

import "fmt"

// Time Limit Exceeded 1004 / 1007
func countPairs(coordinates [][]int, k int) int {
    pairs := 0
    for i := 0; i < len(coordinates); i++ {
        for j := 0; j < len(coordinates); j++ {
            if i != j && (coordinates[i][0] ^ coordinates[j][0]) + (coordinates[i][1] ^ coordinates[j][1]) == k {
                pairs++
            }
        }
    }
    return pairs / 2
}

func countPairs1(coordinates [][]int, k int) int {
    // Distance between points (x1, y1) and (x2, y2) = (x1 ^ x2) + (y1 ^ y2)
    // If x = x1 ^ x2, then x2 = x ^ z1
    // Similarly, if y = y1 ^ y2, then y2 = y ^ y1
    // Since the value of k is very small, we can iterate for all possible values of
    // x and y (based on k) and use a hash table to keep track of the values on the right
    // We'll update right and left as we process and swap left and right after each
    // outer loop iteration.
    right, left := make(map[[2]int]int), make(map[[2]int]int)
    for i := 0; i < len(coordinates); i++ {
        right[[2]int{coordinates[i][0], coordinates[i][1]}]++
    }
    res := 0
    for x := 0; x <= k; x++ {
        y := k - x
        for i := 0; i < len(coordinates); i++ {
            left[[2]int{coordinates[i][0], coordinates[i][1]}]++
            right[[2]int{coordinates[i][0], coordinates[i][1]}]--
            x2, y2 := coordinates[i][0] ^ x, coordinates[i][1] ^ y
            res += right[[2]int{x2,y2}]
        }
        right, left = left, right
    }
    return res
}

func countPairs2(coordinates [][]int, k int) int {
    res, mp := 0, make(map[[2]int]int)
    for i := 0; i < len(coordinates); i++ {
        x1, y1 := coordinates[i][0], coordinates[i][1]
        for j := 0; j <= k; j++ {
            x2, y2 := j ^ x1, (k - j) ^ y1
            res += mp[[2]int{x2, y2}]
        }
        mp[[2]int{x1, y1}]++
    }
    return res
}

func main() {
    // Example 1:
    // Input: coordinates = [[1,2],[4,2],[1,3],[5,2]], k = 5
    // Output: 2
    // Explanation: We can choose the following pairs:
    // - (0,1): Because we have (1 XOR 4) + (2 XOR 2) = 5.
    // - (2,3): Because we have (1 XOR 5) + (3 XOR 2) = 5.
    fmt.Println(countPairs([][]int{{1,2},{4,2},{1,3},{5,2}}, 5)) // 2
    // Example 2:
    // Input: coordinates = [[1,3],[1,3],[1,3],[1,3],[1,3]], k = 0
    // Output: 10
    // Explanation: Any two chosen pairs will have a distance of 0. There are 10 ways to choose two pairs.
    fmt.Println(countPairs([][]int{{1,3},{1,3},{1,3},{1,3},{1,3}}, 0)) // 10

    fmt.Println(countPairs1([][]int{{1,2},{4,2},{1,3},{5,2}}, 5)) // 2
    fmt.Println(countPairs1([][]int{{1,3},{1,3},{1,3},{1,3},{1,3}}, 0)) // 10

    fmt.Println(countPairs2([][]int{{1,2},{4,2},{1,3},{5,2}}, 5)) // 2
    fmt.Println(countPairs2([][]int{{1,3},{1,3},{1,3},{1,3},{1,3}}, 0)) // 10
}