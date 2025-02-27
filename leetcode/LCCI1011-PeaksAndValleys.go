package main

// 面试题 10.11. Peaks and Valleys LCCI
// In an array of integers, a "peak" is an element which is greater than or equal to the adjacent integers and a "valley" is an element which is less than or equal to the adjacent inte­gers. 
// For example, in the array {5, 8, 4, 2, 3, 4, 6}, {8, 6} are peaks and {5, 2} are valleys. 
// Given an array of integers, sort the array into an alternating sequence of peaks and valleys.

// Example:
// Input: [5, 3, 1, 2, 3]
// Output: [5, 1, 3, 2, 3]

// Note:
//     nums.length <= 10000

import "fmt"
import "sort"

func wiggleSort(nums []int)  {
    sort.Ints(nums)
    for i := 0; i < len(nums) - 1; i = i + 2 {
        nums[i], nums[i + 1] = nums[i + 1], nums[i]
    }
}

func main() {
    // Example:
    // Input: [5, 3, 1, 2, 3]
    // Output: [5, 1, 3, 2, 3]
    arr1 := []int{5, 3, 1, 2, 3}
    fmt.Println("arr1: ", arr1) // [5 3 1 2 3]
    wiggleSort(arr1)
    fmt.Println("arr1: ", arr1) // [2 1 3 3 5]

    arr2 := []int{1,2,3,4,5,6,7,8,9}
    fmt.Println("arr2: ", arr2) // [1 2 3 4 5 6 7 8 9]
    wiggleSort(arr2)
    fmt.Println("arr2: ", arr2) // [2 1 4 3 6 5 8 7 9]

    arr3 := []int{9,8,7,6,5,4,3,2,1}
    fmt.Println("arr3: ", arr3) // [9 8 7 6 5 4 3 2 1]
    wiggleSort(arr3)
    fmt.Println("arr3: ", arr3) // [2 1 4 3 6 5 8 7 9]
}