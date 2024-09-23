package main

// 1300. Sum of Mutated Array Closest to Target
// Given an integer array arr and a target value target, 
// return the integer value such that when we change all the integers larger than value in the given array to be equal to value, 
// the sum of the array gets as close as possible (in absolute difference) to target.

// In case of a tie, return the minimum such integer.

// Notice that the answer is not neccesarilly a number from arr.

// Example 1:
// Input: arr = [4,9,3], target = 10
// Output: 3
// Explanation: When using 3 arr converts to [3, 3, 3] which sums 9 and that's the optimal answer.

// Example 2:
// Input: arr = [2,3,5], target = 10
// Output: 5

// Example 3:
// Input: arr = [60864,25176,27249,21296,20204], target = 56803
// Output: 11361

// Constraints:
//     1 <= arr.length <= 10^4
//     1 <= arr[i], target <= 10^5

import "fmt"
import "sort"

func findBestValue(arr []int, target int) int {
    sort.Ints(arr)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    prefix := 0
    for i, v := range arr {
        // current: replace arr[i:] value by v
        current := prefix + (len(arr) - i) * v
        if current == target {
            return v
        }
        if current < target {
            prefix += v
        }
        if current > target {
            length := len(arr) - i
            value := (target - prefix) / length
            if abs(value * length + prefix - target) <= abs((value + 1) * length + prefix - target) {
                return value
            }
            return value + 1
        }
    }
    return arr[len(arr)-1]
}

// binary search
func findBestValue1(arr []int, target int) int {
    right := 0
    for i := range arr {
        if arr[i] > right {
            right = arr[i]
        }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    right += 1
    diff, res, left := 1 << 31, arr[0], 0
    for left < right {
        mid := (left + right) / 2
        sum := 0
        for _, val := range arr {
            if val > mid {
                sum += mid
            } else {
                sum += val
            }
        }
        temp := abs(sum - target)
        if temp < diff {
            diff = temp
            res = mid
        } else if temp == diff {
            if res > mid {
                res = mid
            }
        }
        if sum < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [4,9,3], target = 10
    // Output: 3
    // Explanation: When using 3 arr converts to [3, 3, 3] which sums 9 and that's the optimal answer.
    fmt.Println(findBestValue([]int{4,9,3}, 10)) // 3
    // Example 2:
    // Input: arr = [2,3,5], target = 10
    // Output: 5
    fmt.Println(findBestValue([]int{2,3,5}, 10)) // 5
    // Example 3:
    // Input: arr = [60864,25176,27249,21296,20204], target = 56803
    // Output: 11361
    fmt.Println(findBestValue([]int{60864,25176,27249,21296,20204}, 56803)) // 11361

    fmt.Println(findBestValue1([]int{4,9,3}, 10)) // 3
    fmt.Println(findBestValue1([]int{2,3,5}, 10)) // 5
    fmt.Println(findBestValue1([]int{60864,25176,27249,21296,20204}, 56803)) // 11361
}