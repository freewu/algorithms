package main

// 954. Array of Doubled Pairs
// Given an integer array of even length arr, 
// return true if it is possible to reorder arr such that arr[2 * i + 1] = 2 * arr[2 * i] for every 0 <= i < len(arr) / 2, 
// or false otherwise.

// Example 1:
// Input: arr = [3,1,3,6]
// Output: false

// Example 2:
// Input: arr = [2,1,2,6]
// Output: false

// Example 3:
// Input: arr = [4,-2,2,-4]
// Output: true
// Explanation: We can take two groups, [-2,-4] and [2,4] to form [-2,-4,2,4] or [2,4,-2,-4].
 
// Constraints:
//     2 <= arr.length <= 3 * 10^4
//     arr.length is even.
//     -10^5 <= arr[i] <= 10^5

import "fmt"
import "sort"

func canReorderDoubled(arr []int) bool {
    mp := map[int]int{}
    for _, v := range arr { // 统计出现次数
        mp[v]++
    }
    sort.Slice(arr, func(i, j int) bool { return arr[i] * arr[i] < arr[j] * arr[j] })
    for _, v := range arr {
        if mp[v] != 0 { // 有出现过
            mp[v]--
            mp[v * 2]--
        }
        if mp[v * 2] < 0 { // 没有二倍数了 如 v = 2  v * 2 = 4 (2,4) (3,6) (1,2)
            return false
        }
    }
    return true
}

func canReorderDoubled1(arr []int) bool {
    cnt := make(map[int]int, len(arr))
    for _ , v := range arr { // 统计出现次数
        cnt[v] ++
    }
    if cnt[0] % 2 == 1 { // 
        return false 
    }
    ks := []int{}
    for k := range cnt {
        if k != 0 {
            ks = append(ks , k)
        }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    sort.Slice(ks , func(i , j int)bool{
        return abs(ks[i]) < abs(ks[j]) 
    })
    for _ , x := range ks {
        if cnt[2 * x] < cnt[x] {
            return false
        }
        cnt[2 * x] -= cnt[x]
    }
    return true 
}

func main() {
    // Example 1:
    // Input: arr = [3,1,3,6]
    // Output: false
    fmt.Println(canReorderDoubled([]int{3,1,3,6})) // false
    // Example 2:
    // Input: arr = [2,1,2,6]
    // Output: false
    fmt.Println(canReorderDoubled([]int{2,1,2,6})) // false
    // Example 3:
    // Input: arr = [4,-2,2,-4]
    // Output: true
    // Explanation: We can take two groups, [-2,-4] and [2,4] to form [-2,-4,2,4] or [2,4,-2,-4].
    fmt.Println(canReorderDoubled([]int{4,-2,2,-4})) // true

    fmt.Println(canReorderDoubled1([]int{3,1,3,6})) // false
    fmt.Println(canReorderDoubled1([]int{2,1,2,6})) // false
    fmt.Println(canReorderDoubled1([]int{4,-2,2,-4})) // true
}