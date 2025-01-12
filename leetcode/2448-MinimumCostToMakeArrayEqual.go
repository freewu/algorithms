package main

// 2448. Minimum Cost to Make Array Equal
// You are given two 0-indexed arrays nums and cost consisting each of n positive integers.

// You can do the following operation any number of times:
//     Increase or decrease any element of the array nums by 1.

// The cost of doing one operation on the ith element is cost[i].

// Return the minimum total cost such that all the elements of the array nums become equal.

// Example 1:
// Input: nums = [1,3,5,2], cost = [2,3,1,14]
// Output: 8
// Explanation: We can make all the elements equal to 2 in the following way:
// - Increase the 0th element one time. The cost is 2.
// - Decrease the 1st element one time. The cost is 3.
// - Decrease the 2nd element three times. The cost is 1 + 1 + 1 = 3.
// The total cost is 2 + 3 + 3 = 8.
// It can be shown that we cannot make the array equal with a smaller cost.

// Example 2:
// Input: nums = [2,2,2,2,2], cost = [4,2,8,1,3]
// Output: 0
// Explanation: All the elements are already equal, so no operations are needed.

// Constraints:
//     n == nums.length == cost.length
//     1 <= n <= 10^5
//     1 <= nums[i], cost[i] <= 10^6
//     Test cases are generated in a way that the output doesn't exceed 2^53-1

import "fmt"
import "sort"

func minCost(nums []int, cost []int) int64 {
    n := len(nums)
    arr := make([][2]int, n)
    for i := 0; i < n; i++ {
        arr[i] = [2]int{ nums[i], cost[i] }
    }
    sort.Slice(arr, func(i, j int) bool { 
        return arr[i][0] < arr[j][0] 
    })
    prefix := make([]int, n)
    prefix[0] = arr[0][1]
    for i := 1; i < n; i++ {
        prefix[i] = prefix[i-1] + arr[i][1]
    }
    res := 0
    for i := 1; i < n; i++ {
        res += ((arr[i][0] - arr[0][0]) * arr[i][1])
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < n; i++ {
        jump := (arr[i][0] - arr[i-1][0])
        res = min(res, res - jump * prefix[n - 1] + 2 * jump * prefix[i-1])
    }
    return int64(res)
}

func minCost1(nums []int, costs []int) int64 {
    type Node struct { num, cost int64 }
    nodes := make([]Node, len(nums))
    for i := range nums {
        nodes[i] = Node{num: int64(nums[i]), cost: int64(costs[i])}
    }
    sort.Slice(nodes, func(i, j int) bool {
        return nodes[i].num < nodes[j].num
    })
    sumCost, sumDiff, minCost := int64(0), int64(0), int64(0)
     for i := range nodes {
        diff := nodes[i].num - nodes[0].num
        sumDiff += diff
        sumCost += nodes[i].cost
        minCost += nodes[i].cost * diff
    }
    prevSumCost, currCost := nodes[0].cost, minCost
    for i := 1; i < len(nodes); i++ {
        diff := nodes[i].num - nodes[i-1].num
        currCost = currCost - (sumCost-prevSumCost) * diff + prevSumCost*diff
        prevSumCost += nodes[i].cost
        if currCost < minCost {
            minCost = currCost
        }
    }
    return minCost
}

func minCost2(nums []int, cost []int) int64 {
    arr := [][2]int{}
    for i, v := range nums {
        arr = append(arr, [2]int{ v, cost[i] })
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    l, r := 0, 0
    for _, v := range nums {
        l, r = min(l, v), max(r, v)
    }
    getCost := func(arr [][2]int, target int) int64 {
        res := 0
        for _, v := range arr {
            num, cost := v[0], v[1]
            res += (abs((num - target)) * cost)
        }
        return int64(res)
    }
    for l < r {
        mid := l + (r - l) / 2
        cost1, cost2 := getCost(arr, mid), getCost(arr, mid + 1)
        if cost1 < cost2 { // minimum exists in left side
            r = mid
        } else { // minimum exists in right side
            l = mid + 1
        }
    }
    return getCost(arr, l)
}

func main() {
    // Example 1:
    // Input: nums = [1,3,5,2], cost = [2,3,1,14]
    // Output: 8
    // Explanation: We can make all the elements equal to 2 in the following way:
    // - Increase the 0th element one time. The cost is 2.
    // - Decrease the 1st element one time. The cost is 3.
    // - Decrease the 2nd element three times. The cost is 1 + 1 + 1 = 3.
    // The total cost is 2 + 3 + 3 = 8.
    // It can be shown that we cannot make the array equal with a smaller cost.
    fmt.Println(minCost([]int{1,3,5,2}, []int{2,3,1,14})) // 8
    // Example 2:
    // Input: nums = [2,2,2,2,2], cost = [4,2,8,1,3]
    // Output: 0
    // Explanation: All the elements are already equal, so no operations are needed.
    fmt.Println(minCost([]int{2,2,2,2,2}, []int{4,2,8,1,3})) // 0

    fmt.Println(minCost([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 82
    fmt.Println(minCost([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 82
    fmt.Println(minCost([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 82
    fmt.Println(minCost([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 82

    fmt.Println(minCost1([]int{1,3,5,2}, []int{2,3,1,14})) // 8
    fmt.Println(minCost1([]int{2,2,2,2,2}, []int{4,2,8,1,3})) // 0
    fmt.Println(minCost1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 82
    fmt.Println(minCost1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 82
    fmt.Println(minCost1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 82
    fmt.Println(minCost1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 82

    fmt.Println(minCost2([]int{1,3,5,2}, []int{2,3,1,14})) // 8
    fmt.Println(minCost2([]int{2,2,2,2,2}, []int{4,2,8,1,3})) // 0
    fmt.Println(minCost2([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 82
    fmt.Println(minCost2([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 82
    fmt.Println(minCost2([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 82
    fmt.Println(minCost2([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 82
}