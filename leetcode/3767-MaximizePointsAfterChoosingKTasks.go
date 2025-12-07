package main

// 3767. Maximize Points After Choosing K Tasks
// You are given two integer arrays, technique1 and technique2, each of length n, where n represents the number of tasks to complete.
//     1. If the ith task is completed using technique 1, you earn technique1[i] points.
//     2. If it is completed using technique 2, you earn technique2[i] points.

// You are also given an integer k, representing the minimum number of tasks that must be completed using technique 1.

// You must complete at least k tasks using technique 1 (they do not need to be the first k tasks).

// The remaining tasks may be completed using either technique.

// Return an integer denoting the maximum total points you can earn.

// Example 1:
// Input: technique1 = [5,2,10], technique2 = [10,3,8], k = 2
// Output: 22
// Explanation:
// We must complete at least k = 2 tasks using technique1.
// Choosing technique1[1] and technique1[2] (completed using technique 1), and technique2[0] (completed using technique 2), yields the maximum points: 2 + 10 + 10 = 22.

// Example 2:
// Input: technique1 = [10,20,30], technique2 = [5,15,25], k = 2
// Output: 60
// Explanation:
// We must complete at least k = 2 tasks using technique1.
// Choosing all tasks using technique 1 yields the maximum points: 10 + 20 + 30 = 60.

// Example 3:
// Input: technique1 = [1,2,3], technique2 = [4,5,6], k = 0
// Output: 15
// Explanation:
// Since k = 0, we are not required to choose any task using technique1.
// Choosing all tasks using technique 2 yields the maximum points: 4 + 5 + 6 = 15.

// Constraints:
//     1 <= n == technique1.length == technique2.length <= 10^5
//     1 <= technique1[i], technique2​​​​​​​[i] <= 10​​​​​​​^5
//     0 <= k <= n

import "fmt"
import "slices" 
import "sort"

func maxPoints(technique1 []int, technique2 []int, k int) int64 {
    res, n := 0, len(technique1)
    arr := technique1[:0]
    for i, v := range technique1 {
        res += v
        diff := technique2[i] - v
        if diff > 0 {
            arr = append(arr, diff)
        }
    }
    slices.SortFunc(arr, func(a, b int) int { 
        return b - a 
    })
    for _, v := range arr[:min(n-k, len(arr))] {
        res += (v)
    }
    return int64(res)
}

func maxPoints1(technique1 []int, technique2 []int, k int) int64 {
    res := 0
    for _, v := range technique1 {
        res += v
    }
    n := len(technique1)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = technique2[i] - technique1[i]
    }
    sort.Ints(arr)
    for len(arr) > k && arr[len(arr)-1] > 0 {
        res += arr[len(arr) - 1]
        arr = arr[:len(arr) - 1]
    }
    return int64(res)
}

func maxPoints2(technique1 []int, technique2 []int, k int) int64 {
    res, n, diff := 0, len(technique1), make([]int,0)
    for i := 0; i < n; i++ {
        a, b := technique1[i], technique2[i]
        if a < b {
            diff = append(diff, technique2[i]-technique1[i])
        }
        res += max(a, b)
    }
    m := len(diff)
    if k <= n - m {
        return int64(res)
    }
    k -= n-m
    sort.Ints(diff)
    for i := 0; i < k; i++ {
        res -= diff[i]
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: technique1 = [5,2,10], technique2 = [10,3,8], k = 2
    // Output: 22
    // Explanation:
    // We must complete at least k = 2 tasks using technique1.
    // Choosing technique1[1] and technique1[2] (completed using technique 1), and technique2[0] (completed using technique 2), yields the maximum points: 2 + 10 + 10 = 22.
    fmt.Println(maxPoints([]int{5,2,10}, []int{10,3,8}, 2)) // 22
    // Example 2:
    // Input: technique1 = [10,20,30], technique2 = [5,15,25], k = 2
    // Output: 60
    // Explanation:
    // We must complete at least k = 2 tasks using technique1.
    // Choosing all tasks using technique 1 yields the maximum points: 10 + 20 + 30 = 60.
    fmt.Println(maxPoints([]int{10,20,30}, []int{5,15,25}, 2)) // 60
    // Example 3:
    // Input: technique1 = [1,2,3], technique2 = [4,5,6], k = 0
    // Output: 15
    // Explanation:
    // Since k = 0, we are not required to choose any task using technique1.
    // Choosing all tasks using technique 2 yields the maximum points: 4 + 5 + 6 = 15.
    fmt.Println(maxPoints([]int{1,2,3}, []int{4,5,6}, 0)) // 15

    fmt.Println(maxPoints([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 65
    fmt.Println(maxPoints([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 45
    fmt.Println(maxPoints([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 45
    fmt.Println(maxPoints([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 65

    fmt.Println(maxPoints1([]int{5,2,10}, []int{10,3,8}, 2)) // 22
    fmt.Println(maxPoints1([]int{10,20,30}, []int{5,15,25}, 2)) // 60
    fmt.Println(maxPoints1([]int{1,2,3}, []int{4,5,6}, 0)) // 15
    fmt.Println(maxPoints1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 65
    fmt.Println(maxPoints1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 45
    fmt.Println(maxPoints1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 45
    fmt.Println(maxPoints1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 65

    fmt.Println(maxPoints2([]int{5,2,10}, []int{10,3,8}, 2)) // 22
    fmt.Println(maxPoints2([]int{10,20,30}, []int{5,15,25}, 2)) // 60
    fmt.Println(maxPoints2([]int{1,2,3}, []int{4,5,6}, 0)) // 15
    fmt.Println(maxPoints2([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 65
    fmt.Println(maxPoints2([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 45
    fmt.Println(maxPoints2([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 45
    fmt.Println(maxPoints2([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 65
}