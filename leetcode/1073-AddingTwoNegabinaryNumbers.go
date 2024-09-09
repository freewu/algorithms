package main

// 1073. Adding Two Negabinary Numbers
// Given two numbers arr1 and arr2 in base -2, return the result of adding them together.

// Each number is given in array format:  as an array of 0s and 1s, from most significant bit to least significant bit.  
// For example, arr = [1,1,0,1] represents the number (-2)^3 + (-2)^2 + (-2)^0 = -3.  
// A number arr in array, format is also guaranteed to have no leading zeros: either arr == [0] or arr[0] == 1.

// Return the result of adding arr1 and arr2 in the same format: as an array of 0s and 1s with no leading zeros.

// Example 1:
// Input: arr1 = [1,1,1,1,1], arr2 = [1,0,1]
// Output: [1,0,0,0,0]
// Explanation: arr1 represents 11, arr2 represents 5, the output represents 16.

// Example 2:
// Input: arr1 = [0], arr2 = [0]
// Output: [0]

// Example 3:
// Input: arr1 = [0], arr2 = [1]
// Output: [1]

// Constraints:
//     1 <= arr1.length, arr2.length <= 1000
//     arr1[i] and arr2[i] are 0 or 1
//     arr1 and arr2 have no leading zeros

import "fmt"

func addNegabinary(arr1 []int, arr2 []int) []int {
    res, i, j, carry := []int{}, len(arr1) - 1, len(arr2) - 1, 0
    for i >= 0 || j >= 0 || carry != 0 {
        if i >= 0 && arr1[i] == 1 { // 判断是否进位
            carry += 1
        }
        if j >= 0 && arr2[j] == 1 { // 判断是否进位
            carry += 1
        }
        res = append(res, carry & 1)
        //if the current digit is 2 then, it equals to add -1 to the next digit
        carry = -(carry >> 1)
        i--
        j--
    }
    reverse := func (arr []int) []int {
        i, j := 0, len(arr) - 1
        for i < j {
            arr[i], arr[j] = arr[j], arr[i]
            i++
            j--
        }
        return arr
    }
    reverse(res)
    for len(res) > 1 && res[0] == 0 {
        res = res[1:]
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr1 = [1,1,1,1,1], arr2 = [1,0,1]
    // Output: [1,0,0,0,0]
    // Explanation: arr1 represents 11, arr2 represents 5, the output represents 16.
    fmt.Println(addNegabinary([]int{1,1,1,1,1},[]int{1,0,1})) // [1,0,0,0,0]
    // Example 2:
    // Input: arr1 = [0], arr2 = [0]
    // Output: [0]
    fmt.Println(addNegabinary([]int{0},[]int{0})) // [0]
    // Example 3:
    // Input: arr1 = [0], arr2 = [1]
    // Output: [1]
    fmt.Println(addNegabinary([]int{0},[]int{1})) // [1]
}