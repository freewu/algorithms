package main

// 1534. Count Good Triplets
// Given an array of integers arr, and three integers a, b and c. You need to find the number of good triplets.
// A triplet (arr[i], arr[j], arr[k]) is good if the following conditions are true:
//     0 <= i < j < k < arr.length
//     |arr[i] - arr[j]| <= a
//     |arr[j] - arr[k]| <= b
//     |arr[i] - arr[k]| <= c
//     Where |x| denotes the absolute value of x.

// Return the number of good triplets.

// Example 1:
// Input: arr = [3,0,1,1,9,7], a = 7, b = 2, c = 3
// Output: 4
// Explanation: There are 4 good triplets: [(3,0,1), (3,0,1), (3,1,1), (0,1,1)].

// Example 2:
// Input: arr = [1,1,2,2,3], a = 0, b = 0, c = 1
// Output: 0
// Explanation: No triplet satisfies all conditions.
 
// Constraints:
//     3 <= arr.length <= 100
//     0 <= arr[i] <= 1000
//     0 <= a, b, c <= 1000

import "fmt"

// 暴力解法
func countGoodTriplets(arr []int, a int, b int, c int) int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    res := 0
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            for k := j + 1; k < len(arr); k++ {
                //     |arr[i] - arr[j]| <= a
                //     |arr[j] - arr[k]| <= b
                //     |arr[i] - arr[k]| <= c
                if (abs(arr[i] - arr[j]) <= a) && (abs(arr[j] - arr[k]) <= b) && (abs(arr[i] - arr[k]) <= c)  {
                    res++
                }
            }
        }
    }
    return res
}

// best solution
func countGoodTriplets1(arr []int, a int, b int, c int) int {
    max := func(x, y int) int { if x > y { return x; }; return y; }
    min:= func(x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    m := arr[0]
    for _, v := range arr {
        if v > m {
            m = v
        }
    }
    sum := make([]int, m+1)
    res := 0
    for j, v := range arr {
        for k:=j+1; k<len(arr); k++{
            if abs(arr[j] - arr[k]) > b {
                continue
            }
            r := min(m,min(a+arr[j], c+arr[k]))
            l := max(0,max(arr[j]-a, arr[k]-c))
            if r < l {
                continue
            }
            if l <= 0 {
                res += sum[r]
            }else {
                res += sum[r] - sum[l - 1]
            }
        }
        for i := v; i <= m; i++ {
            sum[i] += 1
        }
    }
    return res
}

func main() {
    // There are 4 good triplets: [(3,0,1), (3,0,1), (3,1,1), (0,1,1)].
    fmt.Println(countGoodTriplets([]int{3,0,1,1,9,7}, 7, 2, 3)) // 4
    fmt.Println(countGoodTriplets([]int{1,1,2,2,3}, 0, 0, 1)) // 0

    fmt.Println(countGoodTriplets1([]int{3,0,1,1,9,7}, 7, 2, 3)) // 4
    fmt.Println(countGoodTriplets1([]int{1,1,2,2,3}, 0, 0, 1)) // 0
}