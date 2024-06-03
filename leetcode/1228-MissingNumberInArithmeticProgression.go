package main

// 1228. Missing Number In Arithmetic Progression
// In some array arr, the values were in arithmetic progression: 
//     the values arr[i + 1] - arr[i] are all equal for every 0 <= i < arr.length - 1.

// A value from arr was removed that was not the first or last value in the array.
// Given arr, return the removed value.

// Example 1:
// Input: arr = [5,7,11,13]
// Output: 9
// Explanation: The previous array was [5,7,9,11,13].

// Example 2:
// Input: arr = [15,13,12]
// Output: 14

// Explanation: The previous array was [15,14,13,12].
 
// Constraints:
// 3 <= arr.length <= 1000
// 0 <= arr[i] <= 10^5
// The given array is guaranteed to be a valid array.

import "fmt"
import "sort"

func missingNumber(arr []int) int {
    n := len(arr)
    diff := (arr[n-1] - arr[0]) / n
    if diff == 0 {
        return arr[0]
    }
    left, right := 0, n - 1
    for left < right - 1 {
        mid := left + (right - left) / 2
        if arr[mid] == arr[left] + diff * (mid - left) {
            left = mid
        } else {
            right = mid
        }
    }
    return arr[left] + diff
}

func missingNumber1(arr []int) int {
    n := len(arr)
    diff := (arr[n-1] - arr[0]) / n
    if diff == 0 {
        return arr[0]
    }
    i := sort.Search(n, func(p int) bool {
        return arr[p] != arr[0] + p * diff
    })
    return arr[i] - diff
}

func main() {
    // Example 1:
    // Input: arr = [5,7,11,13]
    // Output: 9
    // Explanation: The previous array was [5,7,9,11,13].
    fmt.Println(missingNumber([]int{5,7,11,13})) // 9
    // Example 2:
    // Input: arr = [15,13,12]
    // Output: 14
    fmt.Println(missingNumber([]int{15,13,12})) // 14

    fmt.Println(missingNumber1([]int{5,7,11,13})) // 9
    fmt.Println(missingNumber1([]int{15,13,12})) // 14
}