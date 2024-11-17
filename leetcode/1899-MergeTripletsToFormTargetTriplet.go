package main

// 1899. Merge Triplets to Form Target Triplet
// A triplet is an array of three integers. 
// You are given a 2D integer array triplets, where triplets[i] = [ai, bi, ci] describes the ith triplet. 
// You are also given an integer array target = [x, y, z] that describes the triplet you want to obtain.

// To obtain target, you may apply the following operation on triplets any number of times (possibly zero):
//     Choose two indices (0-indexed) i and j (i != j) and update triplets[j] to become [max(ai, aj), max(bi, bj), max(ci, cj)].
//         For example, if triplets[i] = [2, 5, 3] and triplets[j] = [1, 7, 5], triplets[j] will be updated to [max(2, 1), max(5, 7), max(3, 5)] = [2, 7, 5].

// Return true if it is possible to obtain the target triplet [x, y, z] as an element of triplets, or false otherwise.

// Example 1:
// Input: triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]
// Output: true
// Explanation: Perform the following operations:
// - Choose the first and last triplets [[2,5,3],[1,8,4],[1,7,5]]. Update the last triplet to be [max(2,1), max(5,7), max(3,5)] = [2,7,5]. triplets = [[2,5,3],[1,8,4],[2,7,5]]
// The target triplet [2,7,5] is now an element of triplets.

// Example 2:
// Input: triplets = [[3,4,5],[4,5,6]], target = [3,2,5]
// Output: false
// Explanation: It is impossible to have [3,2,5] as an element because there is no 2 in any of the triplets.

// Example 3:
// Input: triplets = [[2,5,3],[2,3,4],[1,2,5],[5,2,3]], target = [5,5,5]
// Output: true
// Explanation: Perform the following operations:
// - Choose the first and third triplets [[2,5,3],[2,3,4],[1,2,5],[5,2,3]]. Update the third triplet to be [max(2,1), max(5,2), max(3,5)] = [2,5,5]. triplets = [[2,5,3],[2,3,4],[2,5,5],[5,2,3]].
// - Choose the third and fourth triplets [[2,5,3],[2,3,4],[2,5,5],[5,2,3]]. Update the fourth triplet to be [max(2,5), max(5,2), max(5,3)] = [5,5,5]. triplets = [[2,5,3],[2,3,4],[2,5,5],[5,5,5]].
// The target triplet [5,5,5] is now an element of triplets.

// Constraints:
//     1 <= triplets.length <= 10^5
//     triplets[i].length == target.length == 3
//     1 <= ai, bi, ci, x, y, z <= 1000

import "fmt"

func mergeTriplets(triplets [][]int, target []int) bool {
    res := 0
    for _, v := range triplets {
        if v[0] > target[0] || v[1] > target[1] || v[2] > target[2] { continue }
        for i := 0; i < len(target); i++ {
            if v[i] == target[i] {
                res = res | (1 << i)
            }
        }
    }
    return res == 7 // 111
}

func mergeTriplets1(triplets [][]int, target []int) bool {
    candidate := [][]int{}
    for _ , v := range triplets {
        if v[0] >  target[0] || v[1] >  target[1] || v[2] >  target[2] { continue }
        if v[0] == target[0] || v[1] == target[1] || v[2] == target[2] {
            candidate = append(candidate , v) 
        }
    }
    if len(candidate) == 0 { return false }
    for i := 0 ; i < 3 ; i ++{
        mx := 0 
        for _ , v := range candidate {
            if v[i] > mx {  mx = v[i] }
        }
        if target[i] != mx { return false }
    }
    return true 
}

func main() {
    // Example 1:
    // Input: triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]
    // Output: true
    // Explanation: Perform the following operations:
    // - Choose the first and last triplets [[2,5,3],[1,8,4],[1,7,5]]. Update the last triplet to be [max(2,1), max(5,7), max(3,5)] = [2,7,5]. triplets = [[2,5,3],[1,8,4],[2,7,5]]
    // The target triplet [2,7,5] is now an element of triplets.
    fmt.Println(mergeTriplets([][]int{{2,5,3},{1,8,4},{1,7,5}}, []int{2,7,5})) // true
    // Example 2:
    // Input: triplets = [[3,4,5],[4,5,6]], target = [3,2,5]
    // Output: false
    // Explanation: It is impossible to have [3,2,5] as an element because there is no 2 in any of the triplets.
    fmt.Println(mergeTriplets([][]int{{3,4,5},{4,5,6}}, []int{3,2,5})) // false
    // Example 3:
    // Input: triplets = [[2,5,3],[2,3,4],[1,2,5],[5,2,3]], target = [5,5,5]
    // Output: true
    // Explanation: Perform the following operations:
    // - Choose the first and third triplets [[2,5,3],[2,3,4],[1,2,5],[5,2,3]]. Update the third triplet to be [max(2,1), max(5,2), max(3,5)] = [2,5,5]. triplets = [[2,5,3],[2,3,4],[2,5,5],[5,2,3]].
    // - Choose the third and fourth triplets [[2,5,3],[2,3,4],[2,5,5],[5,2,3]]. Update the fourth triplet to be [max(2,5), max(5,2), max(5,3)] = [5,5,5]. triplets = [[2,5,3],[2,3,4],[2,5,5],[5,5,5]].
    // The target triplet [5,5,5] is now an element of triplets.
    fmt.Println(mergeTriplets([][]int{{2,5,3},{2,3,4},{1,2,5},{5,2,3}}, []int{5,5,5})) // true

    fmt.Println(mergeTriplets1([][]int{{2,5,3},{1,8,4},{1,7,5}}, []int{2,7,5})) // true
    fmt.Println(mergeTriplets1([][]int{{3,4,5},{4,5,6}}, []int{3,2,5})) // false
    fmt.Println(mergeTriplets1([][]int{{2,5,3},{2,3,4},{1,2,5},{5,2,3}}, []int{5,5,5})) // true
}