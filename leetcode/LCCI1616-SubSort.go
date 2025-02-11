package main

// 面试题 16.16. Sub Sort LCCI
// Given an array of integers, write a method to find indices m and n such that if you sorted elements m through n, 
// the entire array would be sorted. Minimize n - m (that is, find the smallest such sequence).

// Return [m,n]. If there are no such m and n (e.g. the array is already sorted), return [-1, -1].

// Example:
// Input: [1,2,4,7,10,11,7,12,6,7,16,18,19]
// Output: [3,9]

// Note:
//     0 <= len(array) <= 1000000

import "fmt"

func subSort(array []int) []int {
    if len(array) == 0 { return []int{ -1, -1 } }
    mx := len(array) - 1
    li, ri, lv, rv := mx, 0, array[mx], array[0]
    for i := 1; i <= mx; i++ {
        if array[i] < rv {
            ri = i
        } else {
            rv = array[i]
        }
        j := mx - i
        if array[j] > lv {
            li = j
        } else {
            lv = array[j]
        }
    }
    if li >= ri { return []int{-1, -1} }
    return []int{ li, ri }
}

func subSort1(array []int) []int {
    mn, mx := 1 << 31, -1 << 31
    left, right, n := -1, 0, len(array)
    for i, v := range array {
        mn, mx = min(mn, array[n-i-1]), max(mx, v)
        if v != mx {
            right = i
        }
        if array[n-i-1] != mn {
            left = n - i - 1
        }
    }
    if left != -1 { return []int{ left,right } }
    return []int{ -1, -1 }
}

func main() {
    // Example:
    // Input: [1,2,4,7,10,11,7,12,6,7,16,18,19]
    // Output: [3,9]
    fmt.Println(subSort([]int{1,2,4,7,10,11,7,12,6,7,16,18,19})) // [3,9]

    fmt.Println(subSort([]int{1,2,3,4,5,6,7,8,9})) // [-1 -1]
    fmt.Println(subSort([]int{9,8,7,6,5,4,3,2,1})) // [0 8]

    fmt.Println(subSort1([]int{1,2,4,7,10,11,7,12,6,7,16,18,19})) // [3,9]
    fmt.Println(subSort1([]int{1,2,3,4,5,6,7,8,9})) // [-1 -1]
    fmt.Println(subSort1([]int{9,8,7,6,5,4,3,2,1})) // [0 8]
}