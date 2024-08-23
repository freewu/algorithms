package main

// 1013. Partition Array Into Three Parts With Equal Sum
// Given an array of integers arr, return true if we can partition the array into three non-empty parts with equal sums.

// Formally, we can partition the array if we can find indexes i + 1 < j with 
// (arr[0] + arr[1] + ... + arr[i] == arr[i + 1] + arr[i + 2] + ... + arr[j - 1] == arr[j] + arr[j + 1] + ... + arr[arr.length - 1])

// Example 1:
// Input: arr = [0,2,1,-6,6,-7,9,1,2,0,1]
// Output: true
// Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1

// Example 2:
// Input: arr = [0,2,1,-6,6,7,9,-1,2,0,1]
// Output: false

// Example 3:
// Input: arr = [3,3,6,5,-2,2,5,1,-9,4]
// Output: true
// Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4

// Constraints:
//     3 <= arr.length <= 5 * 10^4
//     -10^4 <= arr[i] <= 10^4

import "fmt"

func canThreePartsEqualSum(arr []int) bool {
    sum := 0
    for _, v := range arr { 
        sum += v 
    }
    if sum % 3 != 0 { // 不能被 3整除
        return false
    }
    target, part, count := sum / 3, 0, 0
    for k, v := range arr {
        part += v
        if part == target {
            part = 0
            count++
            if count == 2 && k != len(arr) - 1 {
                return true
            }
        }
    }
    return false
}

func canThreePartsEqualSum1(arr []int) bool {
    n := len(arr)
    if n < 3 {
        return false
    }
    sum := 0
    for _, v := range arr {
        sum += v
    }
    if sum % 3 != 0 {
        return false
    }
    l, r := 0, n - 1
    left, right := arr[0], arr[n-1]
    for l+1 < r {
        if left == sum / 3 && right == sum / 3 {
            return true
        }
        if left != sum / 3 {
            l++
            left += arr[l]
        }
        if right != sum / 3 {
            r--
            right += arr[r]
        }
    }
    return false
}

func canThreePartsEqualSum2(arr []int) bool {
    for i := 1; i < len(arr); i++ {
        arr[i] += arr[i-1]
    }
    preLastIndex, lastIndex := len(arr)-2, len(arr)-1
    if arr[lastIndex] % 3 > 0 {
        return false
    }
    oneThird, twoThirds := arr[lastIndex] / 3, arr[lastIndex] / 3 * 2
    for i := 0; i < preLastIndex; i++ {
        if arr[i] != oneThird {
            continue
        }
        for j := i + 1; j < lastIndex; j++ {
            if arr[j] == twoThirds && arr[lastIndex]-arr[j] == oneThird {
                return true
            }
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: arr = [0,2,1,-6,6,-7,9,1,2,0,1]
    // Output: true
    // Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
    fmt.Println(canThreePartsEqualSum([]int{0,2,1,-6,6,-7,9,1,2,0,1})) // true
    // Example 2:
    // Input: arr = [0,2,1,-6,6,7,9,-1,2,0,1]
    // Output: false
    fmt.Println(canThreePartsEqualSum([]int{0,2,1,-6,6,7,9,-1,2,0,1})) // false
    // Example 3:
    // Input: arr = [3,3,6,5,-2,2,5,1,-9,4]
    // Output: true
    // Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
    fmt.Println(canThreePartsEqualSum([]int{3,3,6,5,-2,2,5,1,-9,4})) // true

    fmt.Println(canThreePartsEqualSum1([]int{0,2,1,-6,6,-7,9,1,2,0,1})) // true
    fmt.Println(canThreePartsEqualSum1([]int{0,2,1,-6,6,7,9,-1,2,0,1})) // false
    fmt.Println(canThreePartsEqualSum1([]int{3,3,6,5,-2,2,5,1,-9,4})) // true

    fmt.Println(canThreePartsEqualSum2([]int{0,2,1,-6,6,-7,9,1,2,0,1})) // true
    fmt.Println(canThreePartsEqualSum2([]int{0,2,1,-6,6,7,9,-1,2,0,1})) // false
    fmt.Println(canThreePartsEqualSum2([]int{3,3,6,5,-2,2,5,1,-9,4})) // true
}