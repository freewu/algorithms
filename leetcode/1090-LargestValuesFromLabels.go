package main

// 1090. Largest Values From Labels
// There is a set of n items. You are given two integer arrays values 
// and labels where the value and the label of the ith element are values[i] and labels[i] respectively. 
// You are also given two integers numWanted and useLimit.

// Choose a subset s of the n elements such that:
//     The size of the subset s is less than or equal to numWanted.
//     There are at most useLimit items with the same label in s.
//     The score of a subset is the sum of the values in the subset.

// Return the maximum score of a subset s.

// Example 1:
// Input: values = [5,4,3,2,1], labels = [1,1,2,2,3], numWanted = 3, useLimit = 1
// Output: 9
// Explanation: The subset chosen is the first, third, and fifth items.

// Example 2:
// Input: values = [5,4,3,2,1], labels = [1,3,3,3,2], numWanted = 3, useLimit = 2
// Output: 12
// Explanation: The subset chosen is the first, second, and third items.

// Example 3:
// Input: values = [9,8,8,7,6], labels = [0,0,0,1,1], numWanted = 3, useLimit = 1
// Output: 16
// Explanation: The subset chosen is the first and fourth items.
 
// Constraints:
//     n == values.length == labels.length
//     1 <= n <= 2 * 10^4
//     0 <= values[i], labels[i] <= 2 * 10^4
//     1 <= numWanted, useLimit <= n

import "fmt"
import "sort"

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
    res,lm := 0, map[int]int{}
    vl := make([][2]int, len(values))
    for i, l := range labels {
        lm[l] = useLimit
        vl[i] = [2]int{values[i], l}
    }
    sort.Slice(vl, func(i, j int) bool {
        return vl[i][0] > vl[j][0]
    })
    for i := 0; i < len(vl); i++ {
        v := vl[i]
        if lm[v[1]] > 0 {
            res += v[0]
            lm[v[1]]--
            numWanted--
        }
        if numWanted == 0 {
            break
        }
    }
    return res
}

func largestValsFromLabels1(values []int, labels []int, numWanted int, useLimit int) int {
    n := len(values)
    idx := make([]int, n)
    for i := 0; i < n; i++ {
        idx[i] = i
    }
    sort.Slice(idx, func(i, j int) bool {
        return values[idx[i]] > values[idx[j]]
    })
    res, choose := 0, 0
    count := make(map[int]int)
    for i := 0; i < n; i ++ {
        lable := labels[idx[i]]
        if count[lable] == useLimit {
            continue
        }
        choose++
        res += values[idx[i]]
        count[lable] +=1
        if choose == numWanted {
            break
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: values = [5,4,3,2,1], labels = [1,1,2,2,3], numWanted = 3, useLimit = 1
    // Output: 9
    // Explanation: The subset chosen is the first, third, and fifth items.
    fmt.Println(largestValsFromLabels([]int{5,4,3,2,1}, []int{1,1,2,2,3}, 3, 1)) // 9
    // Example 2:
    // Input: values = [5,4,3,2,1], labels = [1,3,3,3,2], numWanted = 3, useLimit = 2
    // Output: 12
    // Explanation: The subset chosen is the first, second, and third items.
    fmt.Println(largestValsFromLabels([]int{5,4,3,2,1}, []int{1,3,3,3,2}, 3, 2)) // 12
    // Example 3:
    // Input: values = [9,8,8,7,6], labels = [0,0,0,1,1], numWanted = 3, useLimit = 1
    // Output: 16
    // Explanation: The subset chosen is the first and fourth items.
    fmt.Println(largestValsFromLabels([]int{9,8,8,7,6}, []int{0,0,0,1,1}, 3, 1)) // 16

    fmt.Println(largestValsFromLabels1([]int{5,4,3,2,1}, []int{1,1,2,2,3}, 3, 1)) // 9
    fmt.Println(largestValsFromLabels1([]int{5,4,3,2,1}, []int{1,3,3,3,2}, 3, 2)) // 12
    fmt.Println(largestValsFromLabels1([]int{9,8,8,7,6}, []int{0,0,0,1,1}, 3, 1)) // 16
}