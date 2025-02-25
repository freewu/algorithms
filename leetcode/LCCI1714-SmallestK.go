package main

// 面试题 17.14. Smallest K LCCI
// Design an algorithm to find the smallest K numbers in an array.

// Example:
// Input:  arr = [1,3,5,7,2,4,6,8], k = 4
// Output:  [1,2,3,4]

// Note:
//     0 <= len(arr) <= 100000
//     0 <= k <= min(100000, len(arr))

import "fmt"
import "sort"

func smallestK(arr []int, k int) []int {
    sort.Ints(arr)
    return arr[:k]
}

func smallestK1(arr []int, k int) []int {
    if arr == nil || len(arr) == 0 || k <= 0 || k > len(arr) { return nil }
    comparator := func(a, b int, f func(a, b int) bool) bool { return f(a, b) }
    var findsmallestk func(arr []int, start, end int, k int)
    findsmallestk = func(arr []int, start, end int, k int) {
        if start >= end { return }
        pivotIndex := start + (end - start) / 2
        left, right, pivot := start, end, arr[pivotIndex]
        for left <= right {
            for left <= right && comparator(arr[left], pivot, func(a, b int) bool {
                return a < b
            }) {
                left++
            }
            for left <= right && comparator(arr[right], pivot, func(a, b int) bool {
                return a > b
            }) {
                right--
            }
            if left <= right {
                arr[left], arr[right] = arr[right], arr[left]
                left++
                right--
            }
        }
        if right > k {
            findsmallestk(arr, start, right, k)
        }
        if left <= k {
            findsmallestk(arr, left, end, k)
        }
        if right == k {
            return
        }
    }
    findsmallestk(arr, 0, len(arr) - 1, k - 1)
    return arr[:k]
}

func main() {
    // Example:
    // Input:  arr = [1,3,5,7,2,4,6,8], k = 4
    // Output:  [1,2,3,4]
    fmt.Println(smallestK([]int{1,3,5,7,2,4,6,8}, 4)) // [1,2,3,4]

    fmt.Println(smallestK([]int{1,2,3,4,5,6,7,8,9}, 4)) // [1,2,3,4]
    fmt.Println(smallestK([]int{9,8,7,6,5,4,3,2,1}, 4)) // [1,2,3,4]
    fmt.Println(smallestK([]int{9,8,7,6,5,4,3,2,1,1,2}, 4)) // [1,1,2,2]

    fmt.Println(smallestK1([]int{1,3,5,7,2,4,6,8}, 4)) // [1,2,3,4]
    fmt.Println(smallestK1([]int{1,2,3,4,5,6,7,8,9}, 4)) // [1,2,3,4]
    fmt.Println(smallestK1([]int{9,8,7,6,5,4,3,2,1}, 4)) // [1,2,3,4]
    fmt.Println(smallestK1([]int{9,8,7,6,5,4,3,2,1,1,2}, 4)) // [1,1,2,2]
}