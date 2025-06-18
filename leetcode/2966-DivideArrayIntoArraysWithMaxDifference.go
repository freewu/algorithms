package main

// 2966. Divide Array Into Arrays With Max Difference
// You are given an integer array nums of size n and a positive integer k.
// Divide the array into one or more arrays of size 3 satisfying the following conditions:
//     Each element of nums should be in exactly one array.
//     The difference between any two elements in one array is less than or equal to k.

// Return a 2D array containing all the arrays. If it is impossible to satisfy the conditions, 
// return an empty array. And if there are multiple answers, return any of them.

// Example 1:
// Input: nums = [1,3,4,8,7,9,3,5,1], k = 2
// Output: [[1,1,3],[3,4,5],[7,8,9]]
// Explanation: We can divide the array into the following arrays: [1,1,3], [3,4,5] and [7,8,9].
// The difference between any two elements in each array is less than or equal to 2.
// Note that the order of elements is not important.

// Example 2:
// Input: nums = [1,3,3,2,7,3], k = 3
// Output: []
// Explanation: It is not possible to divide the array satisfying all the conditions.

// Constraints:
//     n == nums.length
//     1 <= n <= 10^5
//     n is a multiple of 3.
//     1 <= nums[i] <= 10^5
//     1 <= k <= 10^5

import "fmt"
import "sort"

func divideArray(nums []int, k int) [][]int {
    n := len(nums)
    if n % 3 != 0 { return [][]int{} }
    // 先将数组排序
    sort.Ints(nums)
    for i := 0; i < n; i= i + 3 {
        // 子数组中 任意 两个元素的差必须小于或等于 k [n1,n2,n3] 取 n1 n3 做判断即可
        if (nums[i + 2] - nums[i]) > k {
            return [][]int{}
        }
    }
    res := make([][]int, 0, n / 3)
    for i := 0; i < n; i = i + 3 {
        //res = append(res,[]int{nums[i],nums[i+1],nums[i+2]})
        res = append(res, nums[i:i+3])
    }
    return res
}

// best solution
func divideArray1(nums []int, k int) [][]int {
    res := [][]int{}
    sort.Ints(nums)
    for i := 2; i < len(nums); i += 3 {
        if nums[i] - nums[i-2] > k {
            return nil
        }
        res = append(res, nums[i - 2 : i + 1])
    }
    return res
}

func divideArray2(nums []int, k int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    res := make([][]int, n / 3)
    for i := 0; i < n; i += 3 {
        if nums[i + 2] - nums[i] > k {
            return [][]int{}
        }
        res[i/3] = []int{ nums[i], nums[i + 1], nums[i + 2] }
    }
    return res
}

func divideArray3(nums []int, k int) [][]int {
    mn, mx := nums[0], nums[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums { // 找到最大最小值
        mn = min(mn, v)
        mx = max(mx, v)
    }
    count := make([]int, mx - mn + 1)
    for _, v := range nums {
        count[v - mn]++
    }
    sorted := make([]int, 0, len(nums))
    for i, c := range count {
        for j := 0; j < c; j++ {
            sorted = append(sorted, i + mn)
        }
    }
    res := make([][]int, 0, len(nums) / 3)
    for i := 0; i < len(sorted); i += 3 {
        if sorted[i+2]-sorted[i] > k {
            return [][]int{}
        }
        res = append(res, sorted[i:i+3])
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,4,8,7,9,3,5,1], k = 2
    // Output: [[1,1,3],[3,4,5],[7,8,9]]
    // Explanation: We can divide the array into the following arrays: [1,1,3], [3,4,5] and [7,8,9].
    // The difference between any two elements in each array is less than or equal to 2.
    // Note that the order of elements is not important.
    fmt.Println(divideArray([]int{1,3,4,8,7,9,3,5,1},2)) // [[1,1,3],[3,4,5],[7,8,9]]
    // Example 2:
    // Input: nums = [1,3,3,2,7,3], k = 3
    // Output: []
    // Explanation: It is not possible to divide the array satisfying all the conditions.
    fmt.Println(divideArray([]int{1,3,3,2,7,3},3)) // []

    fmt.Println(divideArray([]int{1,2,3,4,5,6,7,8,9},3)) // [[1 2 3] [4 5 6] [7 8 9]]
    fmt.Println(divideArray([]int{9,8,7,6,5,4,3,2,1},3)) // [[1 2 3] [4 5 6] [7 8 9]]

    fmt.Println(divideArray1([]int{1,3,4,8,7,9,3,5,1},2)) // [[1,1,3],[3,4,5],[7,8,9]]
    fmt.Println(divideArray1([]int{1,3,3,2,7,3},3)) // []
    fmt.Println(divideArray1([]int{1,2,3,4,5,6,7,8,9},3)) // [[1 2 3] [4 5 6] [7 8 9]]
    fmt.Println(divideArray1([]int{9,8,7,6,5,4,3,2,1},3)) // [[1 2 3] [4 5 6] [7 8 9]]
    
    fmt.Println(divideArray2([]int{1,3,4,8,7,9,3,5,1},2)) // [[1,1,3],[3,4,5],[7,8,9]]
    fmt.Println(divideArray2([]int{1,3,3,2,7,3},3)) // []
    fmt.Println(divideArray2([]int{1,2,3,4,5,6,7,8,9},3)) // [[1 2 3] [4 5 6] [7 8 9]]
    fmt.Println(divideArray2([]int{9,8,7,6,5,4,3,2,1},3)) // [[1 2 3] [4 5 6] [7 8 9]]

    fmt.Println(divideArray3([]int{1,3,4,8,7,9,3,5,1},2)) // [[1,1,3],[3,4,5],[7,8,9]]
    fmt.Println(divideArray3([]int{1,3,3,2,7,3},3)) // []
    fmt.Println(divideArray3([]int{1,2,3,4,5,6,7,8,9},3)) // [[1 2 3] [4 5 6] [7 8 9]]
    fmt.Println(divideArray3([]int{9,8,7,6,5,4,3,2,1},3)) // [[1 2 3] [4 5 6] [7 8 9]]
}