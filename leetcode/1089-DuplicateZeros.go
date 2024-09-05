package main

// 1089. Duplicate Zeros
// Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.

// Note that elements beyond the length of the original array are not written. 
// Do the above modifications to the input array in place and do not return anything.

// Example 1:
// Input: arr = [1,0,2,3,0,4,5,0]
// Output: [1,0,0,2,3,0,0,4]
// Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

// Example 2:
// Input: arr = [1,2,3]
// Output: [1,2,3]
// Explanation: After calling your function, the input array is modified to: [1,2,3]

// Constraints:
//     1 <= arr.length <= 10^4
//     0 <= arr[i] <= 9

import "fmt"

func duplicateZeros(arr []int)  {
    i, j, n := 0, 0, len(arr)
    back := make([]int, n)
    copy(back, arr)
    for j < n {
        if back[i] == 0 { // 遇 0 则复写
            arr[j] = 0
            j++
            if j >= n { break }
            arr[j] = 0
        } else {
            arr[j] = back[i]
        }
        j++
        i++
    }
    // return res
}

func duplicateZeros1(arr []int) {
    n, span, end := len(arr), 0, 0
    for i, v := range arr {
        if v == 0 {
            n -= 2
        } else {
            n -= 1
        }
        if n <= 0 {
            end = i
            span = len(arr) - 1 - i
            break
        }
    }
    if n == -1 {
        arr[len(arr)-1] = 0
        end--
    }
    for span > 0 {
        v := arr[end]
        if v != 0 {
            arr[end+span] = v
        } else {
            arr[end+span] = v
            span--
            arr[end+span] = 0
        }
        end--
    }
}

func main() {
    // Example 1:
    // Input: arr = [1,0,2,3,0,4,5,0]
    // Output: [1,0,0,2,3,0,0,4]
    // Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
    arr1 := []int{1,0,2,3,0,4,5,0}
    fmt.Println(arr1) // [1,0,2,3,0,4,5,0]
    duplicateZeros(arr1)
    fmt.Println(arr1) // [1,0,0,2,3,0,0,4]
    // Example 2:
    // Input: arr = [1,2,3]
    // Output: [1,2,3]
    // Explanation: After calling your function, the input array is modified to: [1,2,3]
    arr2 := []int{1,2,3}
    fmt.Println(arr2) // [1,2,3]
    duplicateZeros(arr2)
    fmt.Println(arr2) // [1,2,3]

    arr3 := []int{0,0,0,0,0,0,0}
    fmt.Println(arr3) // [0,0,0,0,0,0,0]
    duplicateZeros(arr3)
    fmt.Println(arr3) // [0,0,0,0,0,0,0]

    arr11 := []int{1,0,2,3,0,4,5,0}
    fmt.Println(arr11) // [1,0,2,3,0,4,5,0]
    duplicateZeros1(arr11)
    fmt.Println(arr11) // [1,0,0,2,3,0,0,4]
    arr12 := []int{1,2,3}
    fmt.Println(arr12) // [1,2,3]
    duplicateZeros1(arr12)
    fmt.Println(arr12) // [1,2,3]

    arr13 := []int{0,0,0,0,0,0,0}
    fmt.Println(arr13) // [0,0,0,0,0,0,0]
    duplicateZeros1(arr13)
    fmt.Println(arr13) // [0,0,0,0,0,0,0]
}