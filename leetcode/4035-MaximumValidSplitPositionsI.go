package main

// 4035. Maximum Valid Split Positions I
// You are given an integer array nums.

// You may remove at most one element from nums. 
// Let arr be the array of remaining elements in their original order, and let m be its length.

// A split position i of arr is valid if:
//     1. 0 <= i < m - 1, and
//     2. gcd(arr[0..i]) == gcd(arr[i + 1..m - 1]).

// An array of length 1 has no valid split positions.

// The score of arr is the number of valid split positions in it.

// Return the maximum possible score of arr.

// Here, gcd(a) denotes the greatest common divisor of all elements in the array a.

// Example 1:
// Input: nums = [10,30,15,10]
// Output: 2
// Explanation:
// One optimal solution is to remove nums[2] = 15. Then arr = [10, 30, 10].
// The split positions are:
// Split Position i | gcd(arr[0..i])    | gcd(arr[i + 1..m - 1])
// 0                | 10                | 10
// 1                | 10                | 10
// All split positions are valid. Thus, the answer is 2.

// Example 2:
// Input: nums = [2,10,14]
// Output: 1
// Explanation:
// One optimal solution is to not remove any element. Then arr = [2, 10, 14].
// The split positions are:
// Split Position i | gcd(arr[0..i])   | gcd(arr[i + 1..m - 1])
// 0                | 2                | 2
// 1                | 2                | 14
// Only the split position at index 0 is valid. Thus, the answer is 1.

// Example 3:
// Input: nums = [2,4]
// Output: 0
// Explanation:
// The only remaining array that has a split position is arr = [2, 4].
// The split positions are:
// Split Position i | gcd(arr[0..i])   | gcd(arr[i + 1..m - 1])
// 0                | 2                | 4       
// There are no valid split positions. Thus, the answer is 0.

// Constraints:
//     2 <= nums.length <= 1000
//     1 <= nums[i] <= 10^9​​​​​​​

import "fmt"

func maxValidSplits(nums []int) int {
    res, n := 0, len(nums)
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    // i = -1：不删除任何元素；i >=0：删除原数组下标为i的元素
    for i := -1; i < n; i++ {
        // 构造删除i之后得到的新数组arr
        arr := []int{}
        for j := 0; j < n; j++ {
            if i == j {
                continue // 跳过要删除的下标
            }
            arr = append(arr, nums[j])
        }
        m := len(arr)
        if m < 2 {
            continue // 数组长度小于2，不存在分割点，直接跳过
        }
        // prefix[j]：arr[0...j] 的最大公约数  suffix[j]：arr[j...m-1] 的最大公约数
        prefix, suffix := make([]int, m), make([]int, m)
        prefix[0] = arr[0]
        for j := 1; j < m; j++ {
            prefix[j] = gcd(prefix[j-1], arr[j])
        }
        suffix[m-1] = arr[m-1]
        for j := m - 2; j >= 0; j-- {
            suffix[j] = gcd(suffix[j+1], arr[j])
        }
        score := 0
        // j 为分割点，分割为 [0..j] 和 [j+1..m-1]
        for j := 0; j < m-1; j++ {
            // 左边整体gcd等于右边整体gcd，该分割点有效
            if prefix[j] == suffix[j+1] {
                score++
            }
        }
        // 更新全局最大有效分割数量
        res = max(res, score)
    }
    return res
}

func maxValidSplits1(nums []int) int {
    // 题目意思是删除某个元素之后剩下的最大公约数作为划分，前后要都为这个值咯。
    // 假设删除一个数(a,x,b)，则左右的最大公约数可以计算出来，也知道：
    // 1.如果a==b，则至少为1个，再寻找如果某个地方，从左往右如果等于当前a，则一定是对应的值，是可以可以划分的点。
    // 1.1.后续多少个点则也是。
    // 2.如果a!=b，则说明x很重要，影响到前后。
    // 过程：
    // 1.不移除时，计算前缀和后缀最大公约数，找到前缀和后缀相同的地方计数+1
    // 2.移除之后，
    // 不移除
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    calc := func(nums []int) int {
        res := 0
        p, r := make([]int, len(nums)), make([]int, len(nums))
        p[0] = nums[0]
        r[len(nums)-1] = nums[len(nums)-1]
        for i := 1; i < len(nums); i++ {
            p[i] = gcd(nums[i], p[i-1])
        }
        for i := len(nums) - 2; i >= 0; i-- {
            r[i] = gcd(nums[i], r[i+1])
        }
        for i := 0; i < len(nums)-1; i++ {
            if p[i] == r[i+1] {
                res++
            }
        }
        return res
    }
    res := calc(nums)
    for i := range nums {
        // 移除第i个元素
        if i == 0 {
            // 找最大公约数相同的地方，如果已知后面最大公约数，
            res = max(res, calc(nums[1:]))
        } else if i == len(nums)-1 {
            res = max(res, calc(nums[:len(nums)-1]))
        } else {
            t := make([]int, 0, len(nums))
            t = append(t, nums[:i]...)
            t = append(t, nums[i+1:]...)
            res = max(res, calc(t))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [10,30,15,10]
    // Output: 2
    // Explanation:
    // One optimal solution is to remove nums[2] = 15. Then arr = [10, 30, 10].
    // The split positions are:
    // Split Position i | gcd(arr[0..i])    | gcd(arr[i + 1..m - 1])
    // 0                | 10                | 10
    // 1                | 10                | 10
    // All split positions are valid. Thus, the answer is 2.
    fmt.Println(maxValidSplits([]int{10,30,15,10})) // 2
    // Example 2:
    // Input: nums = [2,10,14]
    // Output: 1
    // Explanation:
    // One optimal solution is to not remove any element. Then arr = [2, 10, 14].
    // The split positions are:
    // Split Position i | gcd(arr[0..i])   | gcd(arr[i + 1..m - 1])
    // 0                | 2                | 2
    // 1                | 2                | 14
    // Only the split position at index 0 is valid. Thus, the answer is 1.
    fmt.Println(maxValidSplits([]int{2,10,14})) // 1
    // Example 3:
    // Input: nums = [2,4]
    // Output: 0
    // Explanation:
    // The only remaining array that has a split position is arr = [2, 4].
    // The split positions are:
    // Split Position i | gcd(arr[0..i])   | gcd(arr[i + 1..m - 1])
    // 0                | 2                | 4       
    // There are no valid split positions. Thus, the answer is 0.
    fmt.Println(maxValidSplits([]int{2,4})) // 0

    fmt.Println(maxValidSplits([]int{1,2,3,4,5,6,7,8,9})) // 7
    fmt.Println(maxValidSplits([]int{9,8,7,6,5,4,3,2,1})) // 7

    fmt.Println(maxValidSplits1([]int{10,30,15,10})) // 2
    fmt.Println(maxValidSplits1([]int{2,10,14})) // 1
    fmt.Println(maxValidSplits1([]int{2,4})) // 0
    fmt.Println(maxValidSplits1([]int{1,2,3,4,5,6,7,8,9})) // 7
    fmt.Println(maxValidSplits1([]int{9,8,7,6,5,4,3,2,1})) // 7
}