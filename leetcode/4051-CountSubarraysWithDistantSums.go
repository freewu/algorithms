package main

// 4051. Count Subarrays with Distant Sums
// You are given an integer array nums and two integers goal and k.

// A subarray nums[i..j] is considered distant if the absolute difference between its sum and goal is at least k.

// Return the number of distant subarrays.

// Example 1:
// Input: nums = [1,2,1], goal = 4, k = 1
// Output: 5
// Explanation:
// The distant subarrays for k = 1 are:
// i | j | nums[i..j]  | Sum | abs(sum - goal)
// 0 | 0 | [1]         | 1   | 3
// 1 | 1 | [2]         | 2   | 2
// 2 | 2 | [1]         | 1   | 3
// 0 | 1 | [1, 2]      | 3   | 1
// 1 | 2 | [2, 1]      | 3   | 1
// Thus, the answer is 5.

// Example 2:
// Input: nums = [2,-1,3], goal = 2, k = 2
// Output: 2
// Explanation:
// The distant subarrays for k = 2 are:
// i | j | nums[i..j]  | Sum | abs(sum - goal)
// 0 | 1 | [-1]        | -1  | 3
// 0 | 2 | [2, -1, 3]  | 4   | 2
// Thus, the answer is 2.

// Example 3:
// Input: nums = [-3,1,2], goal = 0, k = 3
// Output: 2
// Explanation:
// The distant subarrays for k = 3 are:
// i | j | nums[i..j]  | Sum | abs(sum - goal)
// 0 | 0 | [-3]        | -3  | 3
// 1 | 2 | [1, 2]      | 3   | 3
// Thus, the answer is 2.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9
//     -10^9 <= goal <= 10^9
//     0 <= k <= 10^9

import "fmt"
import "sort"
import "slices"

func distantSubarrays(nums []int, goal int, k int) int64 {
    n, m := len(nums), 0
    if k == 0 {
        return int64(n) * int64(n+1) / 2
    }
    pref := make([]int, n+1)
    for i := 0; i < n; i++ {
        pref[i+1] = pref[i] + nums[i]
    }
    temp := append([]int(nil), pref...)
    sort.Ints(temp)
    for i := 0; i < len(temp); i++ {
        if i == 0 || temp[i] != temp[m-1] {
            temp[m] = temp[i]
            m++
        }
    }
    temp = temp[:m]
    tree := make([]int, 4 * m)
    u, l := goal + k, goal - k

    var update func(node, start, end, idx int)
    update = func(node, start, end, idx int) {
        if start == end {
            tree[node]++
            return
        }
        mid := start + (end-start)/2
        if idx <= mid {
            update(2*node+1, start, mid, idx)
        } else {
            update(2*node+2, mid+1, end, idx)
        }
        tree[node] = tree[2*node+1] + tree[2*node+2]
    }
    var query func(node, start, end, l, r int) int
    query = func (node, start, end, l, r int) int {
        if r < start || end < l {
            return 0
        }
        if l <= start && end <= r {
            return tree[node]
        }
        mid := start + (end-start)/2
        p1 := query(2*node+1, start, mid, l, r)
        p2 := query(2*node+2, mid+1, end, l, r)
        return p1 + p2
    }
    lowerBound := func(arr []int, key int) int {
        low, high := 0, len(arr)
        for low < high {
            mid := (low + high) / 2
            if arr[mid] < key {
                low = mid + 1
            } else {
                high = mid
            }
        }
        return low
    }
    upperBound := func(arr []int, key int) int {
        low, high := 0, len(arr)
        for low < high {
            mid := (low + high) / 2
            if arr[mid] <= key {
                low = mid + 1
            } else {
                high = mid
            }
        }
        return low
    }
    rank0 := lowerBound(temp, pref[0])
    update(0, 0, m-1, rank0)
    res := int64(0)
    for i := 1; i <= n; i++ {
        val1 := pref[i] - u
        idx1 := upperBound(temp, val1) - 1
        if idx1 >= 0 {
            res += int64(query(0, 0, m-1, 0, idx1))
        }
        val2 := pref[i] - l
        idx2 := lowerBound(temp, val2)
        if idx2 < m {
            res += int64(query(0, 0, m-1, idx2, m-1))
        }
        r := lowerBound(temp, pref[i])
        update(0, 0, m-1, r)
    }
    return res
}

func distantSubarrays1(nums []int, goal int, k int) int64 {
    n := int64(len(nums))
    if k == 0 {
        return n * (n + 1) / 2
    }
    pref := make([]int64, len(nums)+1)
    for i, x := range nums {
        pref[i+1] = pref[i] + int64(x)
    }
    vals := slices.Clone(pref)
    slices.Sort(vals)
    vals = slices.Compact(vals)
    bit := make([]int64, len(vals)+1)
    add := func(i int) {
        for i++; i < len(bit); i += i & -i {
            bit[i]++
        }
    }
    sum := func(i int) (s int64) {
        for ; i > 0; i -= i & -i {
            s += bit[i]
        }
        return
    }
    near, g, kk := int64(0), int64(goal), int64(k)
    for _, p := range pref {
        low, high := p-g-kk, p-g+kk
        l := sort.Search(len(vals), func(i int) bool { return vals[i] > low })
        r := sort.Search(len(vals), func(i int) bool { return vals[i] >= high })
        near += sum(r) - sum(l) 
        i, _ := slices.BinarySearch(vals, p)
        add(i)
    }
    return n * (n + 1) / 2 - near
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1], goal = 4, k = 1
    // Output: 5
    // Explanation:
    // The distant subarrays for k = 1 are:
    // i | j | nums[i..j]  | Sum | abs(sum - goal)
    // 0 | 0 | [1]         | 1   | 3
    // 1 | 1 | [2]         | 2   | 2
    // 2 | 2 | [1]         | 1   | 3
    // 0 | 1 | [1, 2]      | 3   | 1
    // 1 | 2 | [2, 1]      | 3   | 1
    // Thus, the answer is 5.
    fmt.Println(distantSubarrays([]int{1,2,1}, 4, 1)) // 5
    // Example 2:
    // Input: nums = [2,-1,3], goal = 2, k = 2
    // Output: 2
    // Explanation:
    // The distant subarrays for k = 2 are:
    // i | j | nums[i..j]  | Sum | abs(sum - goal)
    // 0 | 1 | [-1]        | -1  | 3
    // 0 | 2 | [2, -1, 3]  | 4   | 2
    // Thus, the answer is 2.
    fmt.Println(distantSubarrays([]int{2, -1, 3}, 2, 2)) // 2
    // Example 3:
    // Input: nums = [-3,1,2], goal = 0, k = 3
    // Output: 2
    // Explanation:
    // The distant subarrays for k = 3 are:
    // i | j | nums[i..j]  | Sum | abs(sum - goal)
    // 0 | 0 | [-3]        | -3  | 3
    // 1 | 2 | [1, 2]      | 3   | 3
    // Thus, the answer is 2.
    fmt.Println(distantSubarrays([]int{-3,1,2}, 0, 3)) // 2

    fmt.Println(distantSubarrays([]int{1,2,3,4,5,6,7,8,9}, 2, 2)) // 41
    fmt.Println(distantSubarrays([]int{9,8,7,6,5,4,3,2,1}, 2, 2)) // 41

    fmt.Println(distantSubarrays1([]int{1,2,1}, 4, 1)) // 5
    fmt.Println(distantSubarrays1([]int{2, -1, 3}, 2, 2)) // 2
    fmt.Println(distantSubarrays1([]int{-3,1,2}, 0, 3)) // 2
    fmt.Println(distantSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 2, 2)) // 41
    fmt.Println(distantSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 2, 2)) // 41
}