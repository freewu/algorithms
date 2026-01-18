package main

// 3814. Maximum Capacity Within Budget
// You are given two integer arrays costs and capacity, both of length n, 
// where costs[i] represents the purchase cost of the ith machine and capacity[i] represents its performance capacity.

// You are also given an integer budget.

// You may select at most two distinct machines such that the total cost of the selected machines is strictly less than budget.

// Return the maximum achievable total capacity of the selected machines.

// Example 1:
// Input: costs = [4,8,5,3], capacity = [1,5,2,7], budget = 8
// Output: 8
// Explanation:
// Choose two machines with costs[0] = 4 and costs[3] = 3.
// The total cost is 4 + 3 = 7, which is strictly less than budget = 8.
// The maximum total capacity is capacity[0] + capacity[3] = 1 + 7 = 8.

// Example 2:
// Input: costs = [3,5,7,4], capacity = [2,4,3,6], budget = 7
// Output: 6
// Explanation:
// Choose one machine with costs[3] = 4.
// The total cost is 4, which is strictly less than budget = 7.
// The maximum total capacity is capacity[3] = 6.

// Example 3:
// Input: costs = [2,2,2], capacity = [3,5,4], budget = 5
// Output: 9
// Explanation:
// Choose two machines with costs[1] = 2 and costs[2] = 2.
// The total cost is 2 + 2 = 4, which is strictly less than budget = 5.
// The maximum total capacity is capacity[1] + capacity[2] = 5 + 4 = 9.

// Constraints:
//     1 <= n == costs.length == capacity.length <= 10^5
//     1 <= costs[i], capacity[i] <= 10^5
//     1 <= budget <= 2 * 10^5

import "fmt"
import "sort"

func maxCapacity(costs []int, capacity []int, budget int) int {
    res, n := 0, len(costs)
    arr := make([][2]int, n)
    for i:= 0; i < n; i++ { 
        arr[i] = [2]int{ costs[i], capacity[i] } 
    }
    sort.Slice(arr, func(i,j int) bool { 
        return arr[i][0] < arr[j][0] 
    })
    pref := make([]int,n)
    pref[0] = arr[0][1]
    for i := 1; i < n; i++ {
        if pref[i-1] > arr[i][1] { 
            pref[i] = pref[i-1] 
        } else { 
            pref[i] = arr[i][1] 
        }
    }
    for i := 0; i < n; i++ {
        c, ca := arr[i][0], arr[i][1]
        if c < budget && ca > res { 
            res = ca 
        }
        remain := budget - c - 1
        best, low, high := -1, 0, i - 1
        for low <= high {
            mid := (low + high) / 2
            if arr[mid][0] <= remain { 
                best = mid 
                low = mid + 1 
            } else { 
                high = mid - 1 
            }
        }
        if best >= 0 && ca + pref[best] > res { 
            res = ca + pref[best] 
        }
    }
    return res
}

func maxCapacity1(costs []int, capacity []int, budget int) int {
    mx, n := 0, len(costs)
    for i := 0; i < n; i++ { // Find max cost to size Fenwick tree
        if costs[i] > mx {
            mx = costs[i]
        }
    }
    tree := make([]int, mx + 1) // Fenwick tree for prefix maximum (1-indexed)
    update := func(i, val int) {
        for i <= mx {
            if val > tree[i] {
                tree[i] = val
            }
            i += i & -i
        }
    }
    query := func(i int) int {
        res := 0
        for i > 0 {
            if tree[i] > res {
                res = tree[i]
            }
            i -= i & -i
        }
        return res
    }
    res, budgetMinus1 := 0, budget - 1
    for i := 0; i < n; i++ {
        c, cap := costs[i], capacity[i]
        // Choose only this machine
        if c < budget && cap > res {
            res = cap
        }
        // Choose this machine + one previous machine (distinct) with total cost < budget
        limit := budgetMinus1 - c
        if limit > 0 {
            if limit > mx {
                limit = mx
            }
            bestPrev := query(limit)
            if cap + bestPrev > res {
                res = cap + bestPrev
            }
        }
        // Add current machine for future pairing
        update(c, cap)
    }
    return res
}

func main() {
    // Example 1:
    // Input: costs = [4,8,5,3], capacity = [1,5,2,7], budget = 8
    // Output: 8
    // Explanation:
    // Choose two machines with costs[0] = 4 and costs[3] = 3.
    // The total cost is 4 + 3 = 7, which is strictly less than budget = 8.
    // The maximum total capacity is capacity[0] + capacity[3] = 1 + 7 = 8.
    fmt.Println(maxCapacity([]int{4,8,5,3}, []int{1,5,2,7}, 8)) // 8
    // Example 2:
    // Input: costs = [3,5,7,4], capacity = [2,4,3,6], budget = 7
    // Output: 6
    // Explanation:
    // Choose one machine with costs[3] = 4.
    // The total cost is 4, which is strictly less than budget = 7.
    // The maximum total capacity is capacity[3] = 6.
    fmt.Println(maxCapacity([]int{3,5,7,4}, []int{2,4,3,6}, 7)) // 6
    // Example 3:
    // Input: costs = [2,2,2], capacity = [3,5,4], budget = 5
    // Output: 9
    // Explanation:
    // Choose two machines with costs[1] = 2 and costs[2] = 2.
    // The total cost is 2 + 2 = 4, which is strictly less than budget = 5.
    // The maximum total capacity is capacity[1] + capacity[2] = 5 + 4 = 9.
    fmt.Println(maxCapacity([]int{2,2,2}, []int{3,5,4}, 5)) // 9   

    fmt.Println(maxCapacity([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 5)) // 4
    fmt.Println(maxCapacity([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 5)) // 17
    fmt.Println(maxCapacity([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 5)) // 17
    fmt.Println(maxCapacity([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 5)) // 4
    fmt.Println(maxCapacity([]int{4,6}, []int{2,4}, 5)) // 2

    fmt.Println(maxCapacity1([]int{4,8,5,3}, []int{1,5,2,7}, 8)) // 8
    fmt.Println(maxCapacity1([]int{3,5,7,4}, []int{2,4,3,6}, 7)) // 6
    fmt.Println(maxCapacity1([]int{2,2,2}, []int{3,5,4}, 5)) // 9   
    fmt.Println(maxCapacity1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 5)) // 4
    fmt.Println(maxCapacity1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 5)) // 17
    fmt.Println(maxCapacity1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 5)) // 17
    fmt.Println(maxCapacity1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 5)) // 4
    fmt.Println(maxCapacity1([]int{4,6}, []int{2,4}, 5)) // 2
}