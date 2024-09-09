package main

// 1053. Previous Permutation With One Swap
// Given an array of positive integers arr (not necessarily distinct), 
// return the lexicographically largest permutation that is smaller than arr, 
// that can be made with exactly one swap. If it cannot be done, then return the same array.

// Note that a swap exchanges the positions of two numbers arr[i] and arr[j]

// Example 1:
// Input: arr = [3,2,1]
// Output: [3,1,2]
// Explanation: Swapping 2 and 1.

// Example 2:
// Input: arr = [1,1,5]
// Output: [1,1,5]
// Explanation: This is already the smallest permutation.

// Example 3:
// Input: arr = [1,9,4,6,7]
// Output: [1,7,4,6,9]
// Explanation: Swapping 9 and 7.

// Constraints:
//     1 <= arr.length <= 10^4
//     1 <= arr[i] <= 10^4

import "fmt"

func prevPermOpt1(arr []int) []int {
    index := -1
    for i := len(arr)-2; i >= 0; i-- {
        if arr[i] > arr[i+1] {
            index = i
            break
        }
    }
    if index == -1 {
        return arr
    }
    currMax := index+1
    for i := index + 1; i < len(arr); i++ {
        if arr[i] < arr[index] && arr[i] > arr[currMax] {
            currMax = i
        }
    }
    arr[index], arr[currMax] = arr[currMax], arr[index] // 交换两数字 arr[i] 和 arr[j] 的位置
    return arr
}

func prevPermOpt11(arr []int) []int {
    left, n := -1, len(arr)
    for i := n - 2; i >= 0; i-- {
        if arr[i] > arr[i+1] {
            left = i
            break
        }
    }
    if left == -1 {
        return arr
    }
    mx, right := 0, n
    for i := left + 1; i < n; i++ {
        if arr[i] < arr[left] && mx < arr[i] {
            mx = arr[i]
            right = i
        }
    }
    if right == n {
        return arr
    }
    arr[left], arr[right] = arr[right], arr[left]
    return arr
}

func main() {
    // Example 1:
    // Input: arr = [3,2,1]
    // Output: [3,1,2]
    // Explanation: Swapping 2 and 1.
    fmt.Println(prevPermOpt1([]int{3,2,1})) // [3,1,2]
    // Example 2:
    // Input: arr = [1,1,5]
    // Output: [1,1,5]
    // Explanation: This is already the smallest permutation.
    fmt.Println(prevPermOpt1([]int{1,1,5})) // [1,1,5]
    // Example 3:
    // Input: arr = [1,9,4,6,7]
    // Output: [1,7,4,6,9]
    // Explanation: Swapping 9 and 7.
    fmt.Println(prevPermOpt1([]int{1,9,4,6,7})) // [1,7,4,6,9]

    fmt.Println(prevPermOpt11([]int{3,2,1})) // [3,1,2]
    fmt.Println(prevPermOpt11([]int{1,1,5})) // [1,1,5]
    fmt.Println(prevPermOpt11([]int{1,9,4,6,7})) // [1,7,4,6,9]
}