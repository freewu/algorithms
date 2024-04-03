package main

// 658. Find K Closest Elements
// Given a sorted integer array arr, two integers k and x, 
// return the k closest integers to x in the array. The result should also be sorted in ascending order.

// An integer a is closer to x than an integer b if:
//     |a - x| < |b - x|, or
//     |a - x| == |b - x| and a < b
    
// Example 1:

// Input: arr = [1,2,3,4,5], k = 4, x = 3
// Output: [1,2,3,4]

// Example 2:
// Input: arr = [1,2,3,4,5], k = 4, x = -1
// Output: [1,2,3,4]
 
// Constraints:
//     1 <= k <= arr.length
//     1 <= arr.length <= 10^4
//     arr is sorted in ascending order.
//     -10^4 <= arr[i], x <= 10^4

import "fmt"
import "sort"

// sort lib
func findClosestElements(arr []int, k int, x int) []int {
    return arr[sort.Search( len(arr) - k, func(i int) bool { return x - arr[i] <= arr[i + k] - x }):][:k]
}

// 二分
func findClosestElements1(arr []int, k int, x int) []int {
    low, high := 0, len(arr)-k
    for low < high {
        mid := low + (high - low) >> 1
        if x - arr[mid] > arr[mid + k] - x {
            low = mid + 1
        } else {
            high = mid
        }
    }
    return arr[low:low + k]
}

func main() {
    fmt.Println(findClosestElements([]int{1,2,3,4,5}, 4, 3)) // [1,2,3,4]
    fmt.Println(findClosestElements([]int{1,2,3,4,5}, 4, -1)) // [1,2,3,4]

    fmt.Println(findClosestElements1([]int{1,2,3,4,5}, 4, 3)) // [1,2,3,4]
    fmt.Println(findClosestElements1([]int{1,2,3,4,5}, 4, -1)) // [1,2,3,4]
}