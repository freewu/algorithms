package main

// 1191. K-Concatenation Maximum Sum
// Given an integer array arr and an integer k, modify the array by repeating it k times.

// For example, if arr = [1, 2] and k = 3 then the modified array will be [1, 2, 1, 2, 1, 2].

// Return the maximum sub-array sum in the modified array. 
// Note that the length of the sub-array can be 0 and its sum in that case is 0.

// As the answer can be very large, return the answer modulo 10^9 + 7.

// Example 1:
// Input: arr = [1,2], k = 3
// Output: 9

// Example 2:
// Input: arr = [1,-2,1], k = 5
// Output: 2

// Example 3:
// Input: arr = [-1,-2], k = 7
// Output: 0

// Constraints:
//     1 <= arr.length <= 10^5
//     1 <= k <= 10^5
//     -10^4 <= arr[i] <= 10^4

import "fmt"

// func kConcatenationMaxSum(arr []int, k int) int {
//     res, sum, n := 0, 0, len(arr)
//     if n == 0 { return 0 }
//     for _, v := range arr { // calculate the sum of the array
//         sum += v
//     }
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     maxSubArray := func(nums []int) int {
//         l, r, cur, n, res := 0, 0, 0, len(nums), -1 << 31
//         for l < n && r < n{
//             if cur < 0{
//                 cur -= nums[l]
//                 l++
//             } else {
//                 cur += nums[r]
//                 r++
//                 res = max(res, cur)
//             }
//         }
//         return res
//     }
//     mx1 := maxSubArray(arr)
//     //if the max sub-array is the entire arr (sum of max sub-array equals sum calculated before), the result will be sum*k 
//     if mx1 == sum || k == 1{
//         return mx1 * k
//     }
//     // find the max sub-array when two arrays are concatenated
//     arr = append(arr, arr...)
//     mx2 := maxSubArray(arr)
//     res = max(mx2 + sum * (k - 2), max(mx1, mx2))
//     return max(res, 0) % 1_000_000_007
// }

func kConcatenationMaxSum(arr []int, k int) int {
    sum, mod, list := 0, 1_000_000_007, []int{}
    for _, v := range arr {
        sum += v
        list = append(list, v)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    kadanes := func(arr []int) int {
        prev, mx := 0,0 
        for i := 0; i <len(arr); i++ {
            prev = max(prev + arr[i], arr[i])
            mx = max(mx, prev)
        }
        return mx
    }
    if k == 1 {
        return kadanes(arr) % mod 
    } else if sum < 0 { // if sum is neg applying two times kadanes
        for _, v := range arr { list = append(list, v) }
        return kadanes(list) % mod
    }
    // if sum is pos applying two times kadanes and adding the k-2 times whole sum
    for _, v := range arr { list = append(list, v) }
    return (kadanes(list) + (k - 2) * sum) % mod
}

func kConcatenationMaxSum1(arr []int, k int) int {
    res := 0
    max := func(x, y int) int { if x > y { return x; }; return y; }
    getMax := func(nums []int) (int, int) {
        res, sum, maxSum, maxTmp := 0, 0, -1 << 31, 0
        for _, v := range nums {
            sum += v
            maxTmp = max(maxTmp + v, v)
            maxSum = max(maxSum, maxTmp)
        }
        res = maxSum
        return res, sum
    }
    if k == 1 {
        res, _ = getMax(arr)
    } else if k == 2 {
        arr = append(arr, arr...)
        res, _ = getMax(arr)
    } else {
        sum := 0
        arr = append(arr, arr...)
        res, sum = getMax(arr)
        sum /= 2
        if sum > 0 {
            res += (k - 2) * sum
        }
    }
    res %= 1_000_000_007
    return max(res, 0)
}

func main() {
    // Example 1:
    // Input: arr = [1,2], k = 3
    // Output: 9
    fmt.Println(kConcatenationMaxSum([]int{1,2}, 3)) // 9
    // Example 2:
    // Input: arr = [1,-2,1], k = 5
    // Output: 2
    fmt.Println(kConcatenationMaxSum([]int{1,-2,1}, 5)) // 2
    // Example 3:
    // Input: arr = [-1,-2], k = 7
    // Output: 0
    fmt.Println(kConcatenationMaxSum([]int{-1,-2}, 7)) // 0

    fmt.Println(kConcatenationMaxSum([]int{10000,10000,10000,10000,10000,10000,10000,10000,10000,10000}, 100000)) // 999999937

    fmt.Println(kConcatenationMaxSum1([]int{1,2}, 3)) // 9
    fmt.Println(kConcatenationMaxSum1([]int{1,-2,1}, 5)) // 2
    fmt.Println(kConcatenationMaxSum1([]int{-1,-2}, 7)) // 0
    fmt.Println(kConcatenationMaxSum1([]int{10000,10000,10000,10000,10000,10000,10000,10000,10000,10000}, 100000)) // 999999937
}