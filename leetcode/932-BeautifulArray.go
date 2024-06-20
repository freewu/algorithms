package main

// 932. Beautiful Array
// An array nums of length n is beautiful if:
//     nums is a permutation of the integers in the range [1, n].
//     For every 0 <= i < j < n, there is no index k with i < k < j where 2 * nums[k] == nums[i] + nums[j].

// Given the integer n, return any beautiful array nums of length n. 
// There will be at least one valid answer for the given n.

// Example 1:
// Input: n = 4
// Output: [2,1,4,3]

// Example 2:
// Input: n = 5
// Output: [3,1,2,5,4]

// Constraints:
//     1 <= n <= 1000

import "fmt"
import "math/bits"

func beautifulArray(n int) []int {
    res := make([]int, 0, n)
    l := bits.Len16(uint16(n))
    twoPow := uint16(1 << l)
    for i := uint16(0); i < twoPow; i++ {
        num := int(bits.Reverse16(i) >> (16 - l))
        if num < n {
            res = append(res, num + 1)
        }
    }
    return res
}

// 如果A是一个漂亮数组，那么2*A+b是一个漂亮数组
// 如果A是一个奇数漂亮数组，B是一个偶数漂亮数组，那么A+B是一个漂亮数组
// 漂亮数组删除几个数后依然是漂亮数组
// 所以  N=1    ========[1]
//       N=2    ========[1,2]
//       N=3    =========[1,3,2]
//       N=4    =========[1,3,2,4]
//       N=5    =========[1,5,3,2,4]
func beautifulArray1(n int) []int {
    res := make([]int,n)
    for i := 0;i < len(res); i++ {
        res[i]= 1
    }
    var dfs func (arr []int,low int,high int) 
    dfs = func (arr []int,low int,high int) {
        if high <= low {
            return
        }
        mid := (low + high)/2
        dfs(arr,low, mid)
        dfs(arr,mid + 1, high)
        for i := low; i <= mid; i++ {
            arr[i] = 2 * arr[i] -1
        }
        for i := mid+1; i <= high; i++ {
            arr[i] = 2 * arr[i]
        }
        return
    }
    dfs(res, 0, n - 1)
    return res
}

func main() {
    // Example 1:
    // Input: n = 4
    // Output: [2,1,4,3]
    fmt.Println(beautifulArray(4)) // [2,1,4,3]
    // Example 2:
    // Input: n = 5
    // Output: [3,1,2,5,4]
    fmt.Println(beautifulArray(5)) // [3,1,2,5,4]

    fmt.Println(beautifulArray1(4)) // [2,1,4,3]
    fmt.Println(beautifulArray1(5)) // [3,1,2,5,4]
}