package main

// 978. Longest Turbulent Subarray
// Given an integer array arr, return the length of a maximum size turbulent subarray of arr.

// A subarray is turbulent if the comparison sign flips between each adjacent pair of elements in the subarray.

// More formally, a subarray [arr[i], arr[i + 1], ..., arr[j]] of arr is said to be turbulent if and only if:
//     For i <= k < j:
//         arr[k] > arr[k + 1] when k is odd, and
//         arr[k] < arr[k + 1] when k is even.
//     Or, for i <= k < j:
//         arr[k] > arr[k + 1] when k is even, and
//         arr[k] < arr[k + 1] when k is odd.

// Example 1:
// Input: arr = [9,4,2,10,7,8,8,1,9]
// Output: 5
// Explanation: arr[1] > arr[2] < arr[3] > arr[4] < arr[5]

// Example 2:
// Input: arr = [4,8,12,16]
// Output: 2

// Example 3:
// Input: arr = [100]
// Output: 1

// Constraints:
//     1 <= arr.length <= 4 * 10^4
//     0 <= arr[i] <= 10^9

import "fmt"

func maxTurbulenceSize(arr []int) int {
    left, prevSign, res := 0, 0, 1
    getSign := func (n int) int {
        if n == 0 { return 0 }
        if n < 0  { return -1 }
        return 1
    }
    for right := 1; right < len(arr); right++ {
        sign := getSign(arr[right] - arr[right-1])
        switch {
            case sign == 0:
                left = right
            case sign == prevSign:
                left = right - 1
            default:
                if d := right - left + 1; d > res {
                    res = d
                }
        }
        prevSign = sign
    }
    return res
}

func maxTurbulenceSize1(arr []int) int {
    if len(arr) == 0 { return 0 }
    left, right, res := 0, 0, 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for right < len(arr) - 1 {
        if left == right {
            if arr[left] == arr[left + 1] {
                left++
            }
            right++
        } else {
            if arr[right - 1] < arr[right] && arr[right] > arr[right + 1] {
                right++
            } else if arr[right - 1] > arr[right] && arr[right] < arr[right + 1] {
                right++
            } else {
                left = right
            }
        }
        res = max(res, right - left + 1)
    }
    return res 
}

func main() {
    // Example 1:
    // Input: arr = [9,4,2,10,7,8,8,1,9]
    // Output: 5
    // Explanation: arr[1] > arr[2] < arr[3] > arr[4] < arr[5]
    fmt.Println(maxTurbulenceSize([]int{9,4,2,10,7,8,8,1,9})) // 5
    // Example 2:
    // Input: arr = [4,8,12,16]
    // Output: 2
    fmt.Println(maxTurbulenceSize([]int{4,8,12,16})) // 2
    // Example 3:
    // Input: arr = [100]
    // Output: 1
    fmt.Println(maxTurbulenceSize([]int{100})) // 1

    fmt.Println(maxTurbulenceSize1([]int{9,4,2,10,7,8,8,1,9})) // 5
    fmt.Println(maxTurbulenceSize1([]int{4,8,12,16})) // 2
    fmt.Println(maxTurbulenceSize1([]int{100})) // 1
}