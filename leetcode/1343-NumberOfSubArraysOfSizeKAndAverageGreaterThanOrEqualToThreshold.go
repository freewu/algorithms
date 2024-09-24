package main

// 1343. Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold
// Given an array of integers arr and two integers k and threshold, 
// return the number of sub-arrays of size k and average greater than or equal to threshold.

// Example 1:
// Input: arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
// Output: 3
// Explanation: Sub-arrays [2,5,5],[5,5,5] and [5,5,8] have averages 4, 5 and 6 respectively. All other sub-arrays of size 3 have averages less than 4 (the threshold).

// Example 2:
// Input: arr = [11,13,17,23,29,31,7,5,2,3], k = 3, threshold = 5
// Output: 6
// Explanation: The first 6 sub-arrays of size 3 have averages greater than 5. Note that averages are not integers.

// Constraints:
//     1 <= arr.length <= 10^5
//     1 <= arr[i] <= 10^4
//     1 <= k <= arr.length
//     0 <= threshold <= 10^4

import "fmt"

func numOfSubarrays(arr []int, k int, threshold int) int {
    res, sum, subarr := 0, 0, arr[:k]
    for _, v := range subarr {
        sum += v
    }
    if sum / k >= threshold {
        res++
    }
    for i := 1; i <= len(arr)-k; i++ {
        sum -= arr[i - 1]
        sum += arr[i + k - 1]
        if float64(sum) / float64(k) >= float64(threshold) {
            res++
        }
    }
    return res
}

func numOfSubarrays1(arr []int, k int, threshold int) int {
    res, i, n, sum := 0, 0, len(arr), 0
    for i < n {
        if i < k-1 {
            sum += arr[i]
            i++
            continue
        }
        sum += arr[i] // 加入右侧值
        if sum / k >= threshold { // 	计算现在区间的平均值
            res++
        }
        sum -= arr[i-k+1] // 	踢出左侧值
        i++
    }
    return res
}

func numOfSubarrays2(arr []int, k int, threshold int) int {
    threshold *= k // to avoid dividing by k every round
    meetsThreshold, begin, sum := 0, 0, 0
    for end := range arr {
        sum += arr[end]
        if end - begin + 1 == k {
            if sum >= threshold {
                meetsThreshold++
            }
            sum -= arr[begin]
            begin++
        }
    }
    return meetsThreshold
}

func main() {
    // Example 1:
    // Input: arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
    // Output: 3
    // Explanation: Sub-arrays [2,5,5],[5,5,5] and [5,5,8] have averages 4, 5 and 6 respectively. All other sub-arrays of size 3 have averages less than 4 (the threshold).
    fmt.Println(numOfSubarrays([]int{2,2,2,2,5,5,5,8}, 3, 4)) // 3
    // Example 2:
    // Input: arr = [11,13,17,23,29,31,7,5,2,3], k = 3, threshold = 5
    // Output: 6
    // Explanation: The first 6 sub-arrays of size 3 have averages greater than 5. Note that averages are not integers.
    fmt.Println(numOfSubarrays([]int{11,13,17,23,29,31,7,5,2,3}, 3, 5)) // 6

    fmt.Println(numOfSubarrays1([]int{2,2,2,2,5,5,5,8}, 3, 4)) // 3
    fmt.Println(numOfSubarrays1([]int{11,13,17,23,29,31,7,5,2,3}, 3, 5)) // 6

    fmt.Println(numOfSubarrays2([]int{2,2,2,2,5,5,5,8}, 3, 4)) // 3
    fmt.Println(numOfSubarrays2([]int{11,13,17,23,29,31,7,5,2,3}, 3, 5)) // 6
}