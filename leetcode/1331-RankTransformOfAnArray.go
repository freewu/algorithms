package main

// 1331. Rank Transform of an Array
// Given an array of integers arr, replace each element with its rank.

// The rank represents how large the element is. The rank has the following rules:
//     1. Rank is an integer starting from 1.
//     2. The larger the element, the larger the rank. If two elements are equal, their rank must be the same.
//     3. Rank should be as small as possible.

// Example 1:
// Input: arr = [40,10,20,30]
// Output: [4,1,2,3]
// Explanation: 40 is the largest element. 10 is the smallest. 20 is the second smallest. 30 is the third smallest.

// Example 2:
// Input: arr = [100,100,100]
// Output: [1,1,1]
// Explanation: Same elements share the same rank.

// Example 3:
// Input: arr = [37,12,28,9,100,56,80,5,12]
// Output: [5,3,4,2,8,6,7,1,3]

// Constraints:
//     0 <= arr.length <= 10^5
//     -10^9 <= arr[i] <= 10^9

import "fmt"
import "sort"

func arrayRankTransform(arr []int) []int {
    n := len(arr)
    res, arr1, mp := make([]int,n), make([]int,n), make(map[int]int)
    copy(arr1, arr)
    sort.Ints(arr1)
    index := 1 // Rank is an integer starting from 1
    for _, v := range arr1 {
        if _, ok := mp[v]; !ok { // 出现相同的要跳过
            mp[v] = index
            index++
        }
    }
    for i, v := range arr {
        res[i] = mp[v]
    }
    return res
}

func arrayRankTransform1(arr []int) []int {
    res, mp := append([]int{}, arr...), make(map[int]int)
    for _, v := range arr {
        mp[v] = 0
    }
    sort.Ints(arr)
    index := 0
    for i, v := range arr {
        if(i > 0 && v == arr[i-1]) {
        } else {
            index++
        }
        mp[v]= index
    }
    index = 0
    for i, v := range res {
        res[i] = mp[v]
    }
    return res
}

func arrayRankTransform2(arr []int) []int {
    n := len(arr)
    res, t, rank := []int{}, make([]int , n), map[int]int{}
    copy(t, arr)
    sort.Ints(t)

    for _, v := range t {
        if _, ok := rank[v]; ok{ continue }
        rank[v] = len(rank) + 1
    }
    for i := 0; i < n ; i++ {
        res = append(res, rank[arr[i]])
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [40,10,20,30]
    // Output: [4,1,2,3]
    // Explanation: 40 is the largest element. 10 is the smallest. 20 is the second smallest. 30 is the third smallest.
    fmt.Println(arrayRankTransform([]int{40,10,20,30})) // [4,1,2,3]
    // Example 2:
    // Input: arr = [100,100,100]
    // Output: [1,1,1]
    // Explanation: Same elements share the same rank.
    fmt.Println(arrayRankTransform([]int{100,100,100})) // [1,1,1]
    // Example 3:
    // Input: arr = [37,12,28,9,100,56,80,5,12]
    // Output: [5,3,4,2,8,6,7,1,3]
    fmt.Println(arrayRankTransform([]int{37,12,28,9,100,56,80,5,12})) // [5,3,4,2,8,6,7,1,3]

    fmt.Println(arrayRankTransform1([]int{40,10,20,30})) // [4,1,2,3]
    fmt.Println(arrayRankTransform1([]int{100,100,100})) // [1,1,1]
    fmt.Println(arrayRankTransform1([]int{37,12,28,9,100,56,80,5,12})) // [5,3,4,2,8,6,7,1,3]

    fmt.Println(arrayRankTransform2([]int{40,10,20,30})) // [4,1,2,3]
    fmt.Println(arrayRankTransform2([]int{100,100,100})) // [1,1,1]
    fmt.Println(arrayRankTransform2([]int{37,12,28,9,100,56,80,5,12})) // [5,3,4,2,8,6,7,1,3]
}