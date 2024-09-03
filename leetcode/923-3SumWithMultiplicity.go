package main

// 923. 3Sum With Multiplicity
// Given an integer array arr, and an integer target, 
// return the number of tuples i, j, k such that i < j < k and arr[i] + arr[j] + arr[k] == target.

// As the answer can be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: arr = [1,1,2,2,3,3,4,4,5,5], target = 8
// Output: 20
// Explanation: 
// Enumerating by the values (arr[i], arr[j], arr[k]):
// (1, 2, 5) occurs 8 times;
// (1, 3, 4) occurs 8 times;
// (2, 2, 4) occurs 2 times;
// (2, 3, 3) occurs 2 times.

// Example 2:
// Input: arr = [1,1,2,2,2,2], target = 5
// Output: 12
// Explanation: 
// arr[i] = 1, arr[j] = arr[k] = 2 occurs 12 times:
// We choose one 1 from [1,1] in 2 ways,
// and two 2s from [2,2,2,2] in 6 ways.

// Example 3:
// Input: arr = [2,1,3], target = 6
// Output: 1
// Explanation: (1, 2, 3) occured one time in the array so we return 1.

// Constraints:
//     3 <= arr.length <= 3000
//     0 <= arr[i] <= 100
//     0 <= target <= 300

import "fmt"

func threeSumMulti(arr []int, target int) int {
    occurrences, sums, res := make([]int, 101), make([]int, 201), 0
    for _, n := range arr {
        if i := target - n; i >= 0 && i < len(sums) {
            res += sums[i]
        }
        for k, v := range occurrences {
            sums[k + n] += v
        }
        occurrences[n]++
    }
    return res % 1_000_000_007
}

func threeSumMulti1(arr []int, target int) int {
    res, n, limit := 0, len(arr), target
    mp := make(map[int]int)
    for i := 0; i < n; i++ {
        mp[arr[i]]++
    }
    if target > 100 { limit = 100 }
    for i := 0; i <= limit; i++ {
        if _, ok := mp[i]; !ok { continue }
        for j := i; j <= limit; j++ {
            k := target - i - j
            if k < 0 || k < j || k > 100 { continue }
            if  mp[i] == 0 || mp[j] == 0 || mp[k] == 0 { continue }
            if i == j && j == k {
                res += mp[i] * (mp[i] - 1) * (mp[i] - 2) / 6 //Cm3 num!/((num-count)!*count!)
            } else if i == j {
                res += mp[i] * (mp[i] - 1) * mp[k] / 2 
            } else if k == j{
                res += mp[j] * (mp[j] - 1) * mp[i]/2 
            } else {
                res += mp[i] * mp[j] * mp[k]
            }
        }
    }
    return res % 1_000_000_007
}

func main() {
    // Example 1:
    // Input: arr = [1,1,2,2,3,3,4,4,5,5], target = 8
    // Output: 20
    // Explanation: 
    // Enumerating by the values (arr[i], arr[j], arr[k]):
    // (1, 2, 5) occurs 8 times;
    // (1, 3, 4) occurs 8 times;
    // (2, 2, 4) occurs 2 times;
    // (2, 3, 3) occurs 2 times.
    fmt.Println(threeSumMulti([]int{1,1,2,2,3,3,4,4,5,5}, 8)) // 20
    // Example 2: 
    // Input: arr = [1,1,2,2,2,2], target = 5
    // Output: 12
    // Explanation: 
    // arr[i] = 1, arr[j] = arr[k] = 2 occurs 12 times:
    // We choose one 1 from [1,1] in 2 ways,
    // and two 2s from [2,2,2,2] in 6 ways.
    fmt.Println(threeSumMulti([]int{1,1,2,2,2,2}, 5)) // 12
    // Example 3:
    // Input: arr = [2,1,3], target = 6
    // Output: 1
    // Explanation: (1, 2, 3) occured one time in the array so we return 1.
    fmt.Println(threeSumMulti([]int{2,1,3}, 6)) // 1

    fmt.Println(threeSumMulti1([]int{1,1,2,2,3,3,4,4,5,5}, 8)) // 20
    fmt.Println(threeSumMulti1([]int{1,1,2,2,2,2}, 5)) // 12
    fmt.Println(threeSumMulti1([]int{2,1,3}, 6)) // 1
}