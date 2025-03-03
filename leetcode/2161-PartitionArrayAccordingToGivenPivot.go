package main

// 2161. Partition Array According to Given Pivot
// You are given a 0-indexed integer array nums and an integer pivot. 
// Rearrange nums such that the following conditions are satisfied:
//     1. Every element less than pivot appears before every element greater than pivot.
//     2. Every element equal to pivot appears in between the elements less than and greater than pivot.
//     3、 The relative order of the elements less than pivot and the elements greater than pivot is maintained.
//         More formally, consider every pi, pj where pi is the new position of the ith element and pj is the new position of the jth element. 
//         For elements less than pivot, if i < j and nums[i] < pivot and nums[j] < pivot, 
//         then pi < pj. Similarly for elements greater than pivot, if i < j and nums[i] > pivot and nums[j] > pivot, then pi < pj.

// Return nums after the rearrangement.

// Example 1:
// Input: nums = [9,12,5,10,14,3,10], pivot = 10
// Output: [9,5,3,10,10,12,14]
// Explanation: 
// The elements 9, 5, and 3 are less than the pivot so they are on the left side of the array.
// The elements 12 and 14 are greater than the pivot so they are on the right side of the array.
// The relative ordering of the elements less than and greater than pivot is also maintained. [9, 5, 3] and [12, 14] are the respective orderings.

// Example 2:
// Input: nums = [-3,4,3,2], pivot = 2
// Output: [-3,2,4,3]
// Explanation: 
// The element -3 is less than the pivot so it is on the left side of the array.
// The elements 4 and 3 are greater than the pivot so they are on the right side of the array.
// The relative ordering of the elements less than and greater than pivot is also maintained. [-3] and [4, 3] are the respective orderings.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^6 <= nums[i] <= 10^6
//     pivot equals to an element of nums.

import "fmt"

func pivotArray(nums []int, pivot int) []int {
    n, start, middle := len(nums), 0, 0
    res := make([]int, n)
    for i := 0; i < n; i++ {
        if nums[i] < pivot {
            res[start] = nums[i]
            start++
        }
        if nums[i] == pivot {
            middle++
        }
    }
    end := start
    for i := 0; i < n; i++ {
        if nums[i] > pivot {
           res[start + middle] = nums[i]
           middle++
        }
        if nums[i] == pivot {
            res[end] = pivot
            end++
        }
    }
    return res
}

func pivotArray1(nums []int, pivot int) []int {
    n := len(nums)
    res := make([]int, n)
    l, r := 0, n - 1
    for i := 0; i < n; i++ {
        if nums[i] < pivot {
            res[l] = nums[i]
            l++
        } else if nums[i] > pivot {
            res[r] = nums[i]
            r--
        }
    }
    reverse := func (arr []int, l, r int) {
        for l < r {
            arr[l], arr[r] = arr[r], arr[l]
            l++
            r--
        }
    }
    reverse(res,  r + 1, n - 1)
    for l <= r {
        res[l] = pivot
        l++
    }
    return res
}

func pivotArray2(nums []int, pivot int) []int {
    count := 0
    for _, v := range nums {
        if v == pivot {
            count++
        }
    }
    i, res := 0, make([]int, len(nums))
    for _, v := range nums {
        if v < pivot {
            res[i] = v
            i++
        }
    }
    for ; count > 0; count-- {
        res[i] = pivot
        i++
    }
    for _, v := range nums {
        if v > pivot {
            res[i] = v
            i++
        }
    }
    return res
}

func pivotArray3(nums []int, pivot int) []int {
    l, r, n := 0, len(nums) - 1, len(nums)
    res := make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = pivot
        if nums[i] < pivot {
            res[l] = nums[i]
            l++
        }
    }
    for i := 0; i < n; i++ {
        if nums[n - 1 - i] > pivot {
            res[r] = nums[n - 1 - i]
            r--
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [9,12,5,10,14,3,10], pivot = 10
    // Output: [9,5,3,10,10,12,14]
    // Explanation: 
    // The elements 9, 5, and 3 are less than the pivot so they are on the left side of the array.
    // The elements 12 and 14 are greater than the pivot so they are on the right side of the array.
    // The relative ordering of the elements less than and greater than pivot is also maintained. [9, 5, 3] and [12, 14] are the respective orderings.
    fmt.Println(pivotArray([]int{9,12,5,10,14,3,10}, 10)) // [9,5,3,10,10,12,14]
    // Example 2:
    // Input: nums = [-3,4,3,2], pivot = 2
    // Output: [-3,2,4,3]
    // Explanation: 
    // The element -3 is less than the pivot so it is on the left side of the array.
    // The elements 4 and 3 are greater than the pivot so they are on the right side of the array.
    // The relative ordering of the elements less than and greater than pivot is also maintained. [-3] and [4, 3] are the respective orderings.
    fmt.Println(pivotArray([]int{-3,4,3,2}, 2)) // [-3,2,4,3]

    fmt.Println(pivotArray([]int{1,2,3,4,5,6,7,8,9}, 2)) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(pivotArray([]int{9,8,7,6,5,4,3,2,1}, 2)) // [1 2 9 8 7 6 5 4 3]

    fmt.Println(pivotArray1([]int{9,12,5,10,14,3,10}, 10)) // [9,5,3,10,10,12,14]
    fmt.Println(pivotArray1([]int{-3,4,3,2}, 2)) // [-3,2,4,3]
    fmt.Println(pivotArray1([]int{1,2,3,4,5,6,7,8,9}, 2)) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(pivotArray1([]int{9,8,7,6,5,4,3,2,1}, 2)) // [1 2 9 8 7 6 5 4 3]

    fmt.Println(pivotArray2([]int{9,12,5,10,14,3,10}, 10)) // [9,5,3,10,10,12,14]
    fmt.Println(pivotArray2([]int{-3,4,3,2}, 2)) // [-3,2,4,3]
    fmt.Println(pivotArray2([]int{1,2,3,4,5,6,7,8,9}, 2)) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(pivotArray2([]int{9,8,7,6,5,4,3,2,1}, 2)) // [1 2 9 8 7 6 5 4 3]

    fmt.Println(pivotArray3([]int{9,12,5,10,14,3,10}, 10)) // [9,5,3,10,10,12,14]
    fmt.Println(pivotArray3([]int{-3,4,3,2}, 2)) // [-3,2,4,3]
    fmt.Println(pivotArray3([]int{1,2,3,4,5,6,7,8,9}, 2)) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(pivotArray3([]int{9,8,7,6,5,4,3,2,1}, 2)) // [1 2 9 8 7 6 5 4 3]
}