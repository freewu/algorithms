package main

// 2145. Count the Hidden Sequences
// You are given a 0-indexed array of n integers differences, 
// which describes the differences between each pair of consecutive integers of a hidden sequence of length (n + 1). More formally, call the hidden sequence hidden, then we have that differences[i] = hidden[i + 1] - hidden[i].

// You are further given two integers lower and upper 
// that describe the inclusive range of values [lower, upper] that the hidden sequence can contain.

// For example, given differences = [1, -3, 4], lower = 1, upper = 6, 
// the hidden sequence is a sequence of length 4 whose elements are in between 1 and 6 (inclusive).
//     [3, 4, 1, 5] and [4, 5, 2, 6] are possible hidden sequences.
//     [5, 6, 3, 7] is not possible since it contains an element greater than 6.
//     [1, 2, 3, 4] is not possible since the differences are not correct.

// Return the number of possible hidden sequences there are. 
// If there are no possible sequences, return 0.

// Example 1:
// Input: differences = [1,-3,4], lower = 1, upper = 6
// Output: 2
// Explanation: The possible hidden sequences are:
// - [3, 4, 1, 5]
// - [4, 5, 2, 6]
// Thus, we return 2.

// Example 2:
// Input: differences = [3,-4,5,1,-2], lower = -4, upper = 5
// Output: 4
// Explanation: The possible hidden sequences are:
// - [-3, 0, -4, 1, 2, 0]
// - [-2, 1, -3, 2, 3, 1]
// - [-1, 2, -2, 3, 4, 2]
// - [0, 3, -1, 4, 5, 3]
// Thus, we return 4.

// Example 3:
// Input: differences = [4,-7,2], lower = 3, upper = 6
// Output: 0
// Explanation: There are no possible hidden sequences. Thus, we return 0.

// Constraints:
//     n == differences.length
//     1 <= n <= 10^5
//     -10^5 <= differences[i] <= 10^5
//     -10^5 <= lower <= upper <= 10^5

import "fmt"

func numberOfArrays(differences []int, lower int, upper int) int {
    mn, mx := 0, 0
    if differences[0] > 0 {
        mx = differences[0]
    } else {
        mn = differences[0]
    }
    for i := 1; i < len(differences); i++ { // 找到最大 最小值
        differences[i] += differences[i-1]
        if differences[i] > mx {
            mx = differences[i]
        }
        if differences[i] < mn {
            mn = differences[i]
        }
    }
    if lower - mn + mx > upper {
        return 0
    }
    return upper - (lower - mn + mx) + 1
}

func numberOfArrays1(differences []int, lower int, upper int) int {
    mn, mx, cur := 0, 0, 0
    for _, diff := range differences {
        cur += diff
        if cur > mx {
            mx = cur
        }
        if cur < mn {
            mn = cur
        }
    }
    if (mx - mn) > (upper - lower) {
        return 0
    }
    return (upper - lower) - (mx - mn) + 1
}

func main() {
    // Example 1:
    // Input: differences = [1,-3,4], lower = 1, upper = 6
    // Output: 2
    // Explanation: The possible hidden sequences are:
    // - [3, 4, 1, 5]
    // - [4, 5, 2, 6]
    // Thus, we return 2.
    fmt.Println(numberOfArrays([]int{1,-3,4}, 1, 6)) // 2
    // Example 2:
    // Input: differences = [3,-4,5,1,-2], lower = -4, upper = 5
    // Output: 4
    // Explanation: The possible hidden sequences are:
    // - [-3, 0, -4, 1, 2, 0]
    // - [-2, 1, -3, 2, 3, 1]
    // - [-1, 2, -2, 3, 4, 2]
    // - [0, 3, -1, 4, 5, 3]
    // Thus, we return 4.
    fmt.Println(numberOfArrays([]int{3,-4,5,1,-2}, -4, 5)) // 4
    // Example 3:
    // Input: differences = [4,-7,2], lower = 3, upper = 6
    // Output: 0
    // Explanation: There are no possible hidden sequences. Thus, we return 0.
    fmt.Println(numberOfArrays([]int{4,-7,2}, 3, 6)) // 0

    fmt.Println(numberOfArrays1([]int{1,-3,4}, 1, 6)) // 2
    fmt.Println(numberOfArrays1([]int{3,-4,5,1,-2}, -4, 5)) // 4
    fmt.Println(numberOfArrays1([]int{4,-7,2}, 3, 6)) // 0
}