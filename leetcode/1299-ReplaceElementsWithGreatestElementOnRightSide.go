package main

// 1299. Replace Elements with Greatest Element on Right Side
// Given an array arr, 
// replace every element in that array with the greatest element among the elements to its right, 
// and replace the last element with -1.

// After doing so, return the array.

// Example 1:
// Input: arr = [17,18,5,4,6,1]
// Output: [18,6,6,6,1,-1]
// Explanation: 
// - index 0 --> the greatest element to the right of index 0 is index 1 (18).
// - index 1 --> the greatest element to the right of index 1 is index 4 (6).
// - index 2 --> the greatest element to the right of index 2 is index 4 (6).
// - index 3 --> the greatest element to the right of index 3 is index 4 (6).
// - index 4 --> the greatest element to the right of index 4 is index 5 (1).
// - index 5 --> there are no elements to the right of index 5, so we put -1.

// Example 2:
// Input: arr = [400]
// Output: [-1]
// Explanation: There are no elements to the right of index 0.

// Constraints:
//     1 <= arr.length <= 10^4
//     1 <= arr[i] <= 10^5

import "fmt"

func replaceElements(arr []int) []int {
    res := make([]int,len(arr))
    res[len(arr)-1] = -1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := len(arr) - 2; i >= 0; i-- { // 倒数第二开始
        res[i] = max(arr[i+1], res[i+1])
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [17,18,5,4,6,1]
    // Output: [18,6,6,6,1,-1]
    // Explanation: 
    // - index 0 --> the greatest element to the right of index 0 is index 1 (18).
    // - index 1 --> the greatest element to the right of index 1 is index 4 (6).
    // - index 2 --> the greatest element to the right of index 2 is index 4 (6).
    // - index 3 --> the greatest element to the right of index 3 is index 4 (6).
    // - index 4 --> the greatest element to the right of index 4 is index 5 (1).
    // - index 5 --> there are no elements to the right of index 5, so we put -1.
    fmt.Println(replaceElements([]int{17,18,5,4,6,1})) // [18,6,6,6,1,-1]
    // Example 2:
    // Input: arr = [400]
    // Output: [-1]
    // Explanation: There are no elements to the right of index 0.
    fmt.Println(replaceElements([]int{400})) // [-1]
}