package main

// 1442. Count Triplets That Can Form Two Arrays of Equal XOR
// Given an array of integers arr.
// We want to select three indices i, j and k where (0 <= i < j <= k < arr.length).

// Let's define a and b as follows:
//     a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
//     b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]

// Note that ^ denotes the bitwise-xor operation.
// Return the number of triplets (i, j and k) Where a == b.

// Example 1:
// Input: arr = [2,3,1,6,7]
// Output: 4
// Explanation: The triplets are (0,1,2), (0,2,2), (2,3,4) and (2,4,4)

// Example 2:
// Input: arr = [1,1,1,1,1]
// Output: 10
 
// Constraints:
//     1 <= arr.length <= 300
//     1 <= arr[i] <= 10^8

import "fmt"

// 前缀和 + 暴力枚举
func countTriplets(arr []int) int {
    triplets, n := 0, len(arr)
    preXorSum := make([]int, n + 1)
    for i, x := range arr { // 预先成与前值的异或值
        preXorSum[i + 1] = preXorSum[i] ^ x 
    }
    for i := 0; i < n - 1; i++ {
        for j := i + 1; j < n; j++ {
            for k := j; k < n; k++ {
                if preXorSum[k + 1] == preXorSum[i] {
                    triplets++
                }
            }
        }
    }
    return triplets
}

// 前缀和 + 枚举优化
func countTriplets1(arr []int) int {
    triplets, n := 0, len(arr)
    preXorSum := make([]int, n + 1)
    for i, x := range arr {
        preXorSum[i + 1] = preXorSum[i] ^ x
    }
    for i := 0; i < n - 1; i++ {
        for k := i + 1; k < n; k++ {
            if preXorSum[k + 1] == preXorSum[i] { 
                triplets += k - i // i 和 k 固定时，j 取值范围为 [i + 1, k]，只要满足 preXorSum[k+1]=preXorSum[i]，j 就会产生 k−i 大小的贡献
            }
        }
    }
    return triplets
}

// 前缀和 + 哈希
func countTriplets2(arr []int) int {
    triplets, n := 0, len(arr)
    preXorSum, cnt, total := make([]int, n + 1), map[int]int{}, map[int]int{}
    for i, x := range arr {
        preXorSum[i + 1] = preXorSum[i] ^ x
    }
    for k := 1; k < n; k++ {
        cnt[preXorSum[k - 1]]++
        total[preXorSum[k - 1]] += k - 1
        m := cnt[preXorSum[k + 1]]
        s := total[preXorSum[k + 1]]
        triplets += m * k - s 
    }
    return triplets
}

// 滚动前缀和 + 哈希
func countTriplets3(arr []int) int {
    triplets, preXorSum := 0, 0
    cnt, total := map[int]int{}, map[int]int{}
    for k := 1; k < len(arr); k++ {
        cnt[preXorSum]++
        total[preXorSum] += k - 1
        preXorSum ^= arr[k - 1]

        m := cnt[preXorSum ^ arr[k]]
        s := total[preXorSum ^ arr[k]]
        triplets += m * k - s 
    }
    return triplets
}

func main() {
    // Example 1:
    // Input: arr = [2,3,1,6,7]
    // Output: 4
    // Explanation: The triplets are (0,1,2), (0,2,2), (2,3,4) and (2,4,4)
    fmt.Println(countTriplets([]int{2,3,1,6,7})) // 4
    // Example 2:
    // Input: arr = [1,1,1,1,1]
    // Output: 10
    fmt.Println(countTriplets([]int{1,1,1,1,1})) // 10

    fmt.Println(countTriplets1([]int{2,3,1,6,7})) // 4
    fmt.Println(countTriplets1([]int{1,1,1,1,1})) // 10

    fmt.Println(countTriplets2([]int{2,3,1,6,7})) // 4
    fmt.Println(countTriplets2([]int{1,1,1,1,1})) // 10

    fmt.Println(countTriplets3([]int{2,3,1,6,7})) // 4
    fmt.Println(countTriplets3([]int{1,1,1,1,1})) // 10
}