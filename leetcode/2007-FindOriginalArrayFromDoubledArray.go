package main

// 2007. Find Original Array From Doubled Array
// An integer array original is transformed into a doubled array changed by appending twice the value of every element in original,
// and then randomly shuffling the resulting array.

// Given an array changed, return original if changed is a doubled array. 
// If changed is not a doubled array, return an empty array. The elements in original may be returned in any order.

// Example 1:
// Input: changed = [1,3,4,2,6,8]
// Output: [1,3,4]
// Explanation: One possible original array could be [1,3,4]:
// - Twice the value of 1 is 1 * 2 = 2.
// - Twice the value of 3 is 3 * 2 = 6.
// - Twice the value of 4 is 4 * 2 = 8.
// Other original arrays could be [4,3,1] or [3,1,4].

// Example 2:
// Input: changed = [6,3,0,1]
// Output: []
// Explanation: changed is not a doubled array.

// Example 3:
// Input: changed = [1]
// Output: []
// Explanation: changed is not a doubled array.
 
// Constraints:
//     1 <= changed.length <= 10^5
//     0 <= changed[i] <= 10^5

import "fmt"
import "sort"

// O(nlogn)
func findOriginalArray(changed []int) []int {
    sort.Ints(changed) // 从小到大排列
    memo, res := [200001]int{}, []int{}
    for _, v := range changed {
        if memo[v] == 0 { // 如果存在  如果 v = 2 到 memo[4] 不会进来 这里在最外面 memo[4]-- 掉了
            memo[v * 2]++ // 双倍的数 [v * 2] +1 如果存在会在后面减掉 memo[4]++  
            res = append(res, v)
            continue
        }
        memo[v]--
    }
    for _, v := range memo {
        if v != 0 { // 如果 change 是 双倍 数组，那么请你返回 original数组，否则请返回空数组。
            return []int{}
        }
    }
    return res
}

// O(n)
func findOriginalArray1(changed []int) []int {
    memo, res := [100001]int{}, []int{}
    for _, v := range changed {
        memo[v]++
    }
    if memo[0] % 2 == 1 { // 处理 0 * 2 的情况
        return []int{}
    }
    for i := 0; i < memo[0]/2; i++ {
        res = append(res, 0)
    }
    for i := 1 ; i < len(memo); i++ {
        if memo[i] == 0 {
            continue
        }
        if memo[i] > 50000 || memo[i] > memo[i*2] {
            return []int{}
        }
        memo[i*2] -= memo[i]
        for j := 0; j < memo[i]; j++ {
            res = append(res, i)
        }
    }
    return res
}

func findOriginalArray2(changed []int) []int {
    n := len(changed)
    if (n & 1) != 0 { // 奇数直接返回
        return nil
    }
    cnt := [100001]int{}
    for _, x := range changed {
        cnt[x]++
    }
    res := make([]int, 0, n/2)
    for x := range cnt {
        y := x << 1
        for cnt[x] > 0 {
            res = append(res, x)
            if y >= len(cnt) || cnt[y] <= 0 {
                return nil
            }
            cnt[x]--
            cnt[y]--
        }
    }
    return res
}

func main() {

    // Explanation: One possible original array could be [1,3,4]:
    // - Twice the value of 1 is 1 * 2 = 2.
    // - Twice the value of 3 is 3 * 2 = 6.
    // - Twice the value of 4 is 4 * 2 = 8.
    // Other original arrays could be [4,3,1] or [3,1,4].
    fmt.Println(findOriginalArray([]int{1,3,4,2,6,8})) // [1,3,4]
    // Explanation: changed is not a doubled array.
    fmt.Println(findOriginalArray([]int{6,3,0,1})) // []
    // Explanation: changed is not a doubled array.
    fmt.Println(findOriginalArray([]int{1})) // []

    fmt.Println(findOriginalArray1([]int{1,3,4,2,6,8})) // [1,3,4]
    fmt.Println(findOriginalArray1([]int{6,3,0,1})) // []
    fmt.Println(findOriginalArray1([]int{1})) // []

    fmt.Println(findOriginalArray2([]int{1,3,4,2,6,8})) // [1,3,4]
    fmt.Println(findOriginalArray2([]int{6,3,0,1})) // []
    fmt.Println(findOriginalArray2([]int{1})) // []
}