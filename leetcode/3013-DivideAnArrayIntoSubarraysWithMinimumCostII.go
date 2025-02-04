package main

// 3013. Divide an Array Into Subarrays With Minimum Cost II
// You are given a 0-indexed array of integers nums of length n, and two positive integers k and dist.

// The cost of an array is the value of its first element. 
// For example, the cost of [1,2,3] is 1 while the cost of [3,4,1] is 3.

// You need to divide nums into k disjoint contiguous subarrays, 
// such that the difference between the starting index of the second subarray and the starting index of the kth subarray should be less than or equal to dist. 
// In other words, if you divide nums into the subarrays nums[0..(i1 - 1)], nums[i1..(i2 - 1)], ..., nums[ik-1..(n - 1)], then ik-1 - i1 <= dist.

// Return the minimum possible sum of the cost of these subarrays.

// Example 1:
// Input: nums = [1,3,2,6,4,2], k = 3, dist = 3
// Output: 5
// Explanation: The best possible way to divide nums into 3 subarrays is: [1,3], [2,6,4], and [2]. This choice is valid because ik-1 - i1 is 5 - 2 = 3 which is equal to dist. The total cost is nums[0] + nums[2] + nums[5] which is 1 + 2 + 2 = 5.
// It can be shown that there is no possible way to divide nums into 3 subarrays at a cost lower than 5.

// Example 2:
// Input: nums = [10,1,2,2,2,1], k = 4, dist = 3
// Output: 15
// Explanation: The best possible way to divide nums into 4 subarrays is: [10], [1], [2], and [2,2,1]. This choice is valid because ik-1 - i1 is 3 - 1 = 2 which is less than dist. The total cost is nums[0] + nums[1] + nums[2] + nums[3] which is 10 + 1 + 2 + 2 = 15.
// The division [10], [1], [2,2,2], and [1] is not valid, because the difference between ik-1 and i1 is 5 - 1 = 4, which is greater than dist.
// It can be shown that there is no possible way to divide nums into 4 subarrays at a cost lower than 15.

// Example 3:
// Input: nums = [10,8,18,9], k = 3, dist = 1
// Output: 36
// Explanation: The best possible way to divide nums into 4 subarrays is: [10], [8], and [18,9]. This choice is valid because ik-1 - i1 is 2 - 1 = 1 which is equal to dist.The total cost is nums[0] + nums[1] + nums[2] which is 10 + 8 + 18 = 36.
// The division [10], [8,18], and [9] is not valid, because the difference between ik-1 and i1 is 3 - 1 = 2, which is greater than dist.
// It can be shown that there is no possible way to divide nums into 3 subarrays at a cost lower than 36.

// Constraints:
//     3 <= n <= 10^5
//     1 <= nums[i] <= 10^9
//     3 <= k <= n
//     k - 2 <= dist <= n - 2

import "fmt"
import "github.com/emirpasic/gods/v2/maps/treemap"
import "github.com/emirpasic/gods/trees/redblacktree"

func minimumCost(nums []int, k int, dist int) int64 {
    k -= 1
    primary := treemap.NewWith[int, int](func(x, y int) int {
        if nums[x] == nums[y] {
            return x - y
        }
        return nums[x] - nums[y]
    })
    reserve := treemap.NewWith[int, int](func(x, y int) int {
        if nums[x] == nums[y] {
            return x - y
        }
        return nums[x] - nums[y]
    })
    min := func (x, y int64) int64 { if x < y { return x; }; return y; }
    res, sum, n := int64(1 << 61), int64(0), len(nums)
    for right := 1; right < n; right++ {
        if right > dist {
            left := right - dist - 1
            if _, ok := primary.Get(left); ok {
                primary.Remove(left)
                sum -= int64(nums[left])
            } else {
                reserve.Remove(left)
            }
        }
        reserve.Put(right, right)
        mn, _, _ := reserve.Min()
        reserve.Remove(mn)
        primary.Put(mn, mn)
        sum += int64(nums[mn])
        if primary.Size() > k {
            mx, _, _ := primary.Max()
            reserve.Put(mx, mx)
            primary.Remove(mx)
            sum -= int64(nums[mx])
        }
        if primary.Size() == k {
            res = min(res, sum)
        }
    }
    return res + int64(nums[0])
}

func minimumCost1(nums []int, k, dist int) int64 {
    k--
    L := redblacktree.NewWithIntComparator()
    R := redblacktree.NewWithIntComparator()
    add := func(t *redblacktree.Tree, x int) {
        if v, ok := t.Get(x); ok {
            t.Put(x, v.(int) + 1)
        } else {
            t.Put(x, 1)
        }
    }
    del := func(t *redblacktree.Tree, x int) {
        c, _ := t.Get(x)
        if c.(int) > 1 {
            t.Put(x, c.(int)-1)
        } else {
            t.Remove(x)
        }
    }
    sumL := nums[0]
    for _, x := range nums[1 : dist+2] {
        sumL += x
        add(L, x)
    }
    sizeL := dist + 1
    l2r := func() {
        x := L.Right().Key.(int)
        sumL -= x
        sizeL--
        del(L, x)
        add(R, x)
    }
    r2l := func() {
        x := R.Left().Key.(int)
        sumL += x
        sizeL++
        del(R, x)
        add(L, x)
    }
    for sizeL > k {
        l2r()
    }
    res := sumL
    for i := dist + 2; i < len(nums); i++ {
        // 移除 out
        out := nums[i-dist-1]
        if _, ok := L.Get(out); ok {
            sumL -= out
            sizeL--
            del(L, out)
        } else {
            del(R, out)
        }
        // 添加 in
        in := nums[i]
        if in < L.Right().Key.(int) {
            sumL += in
            sizeL++
            add(L, in)
        } else {
            add(R, in)
        }
        // 维护大小
        if sizeL == k-1 {
            r2l()
        } else if sizeL == k+1 {
            l2r()
        }
        res = min(res, sumL)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2,6,4,2], k = 3, dist = 3
    // Output: 5
    // Explanation: The best possible way to divide nums into 3 subarrays is: [1,3], [2,6,4], and [2]. This choice is valid because ik-1 - i1 is 5 - 2 = 3 which is equal to dist. The total cost is nums[0] + nums[2] + nums[5] which is 1 + 2 + 2 = 5.
    // It can be shown that there is no possible way to divide nums into 3 subarrays at a cost lower than 5.
    fmt.Println(minimumCost([]int{1,3,2,6,4,2}, 3, 3)) // 5
    // Example 2:
    // Input: nums = [10,1,2,2,2,1], k = 4, dist = 3
    // Output: 15
    // Explanation: The best possible way to divide nums into 4 subarrays is: [10], [1], [2], and [2,2,1]. This choice is valid because ik-1 - i1 is 3 - 1 = 2 which is less than dist. The total cost is nums[0] + nums[1] + nums[2] + nums[3] which is 10 + 1 + 2 + 2 = 15.
    // The division [10], [1], [2,2,2], and [1] is not valid, because the difference between ik-1 and i1 is 5 - 1 = 4, which is greater than dist.
    // It can be shown that there is no possible way to divide nums into 4 subarrays at a cost lower than 15.
    fmt.Println(minimumCost([]int{10,1,2,2,2,1}, 4, 3)) // 15
    // Example 3:
    // Input: nums = [10,8,18,9], k = 3, dist = 1
    // Output: 36
    // Explanation: The best possible way to divide nums into 4 subarrays is: [10], [8], and [18,9]. This choice is valid because ik-1 - i1 is 2 - 1 = 1 which is equal to dist.The total cost is nums[0] + nums[1] + nums[2] which is 10 + 8 + 18 = 36.
    // The division [10], [8,18], and [9] is not valid, because the difference between ik-1 and i1 is 3 - 1 = 2, which is greater than dist.
    // It can be shown that there is no possible way to divide nums into 3 subarrays at a cost lower than 36.
    fmt.Println(minimumCost([]int{10,8,18,9}, 3, 1)) // 36

    fmt.Println(minimumCost([]int{1,2,3,4,5,6,7,8,9}, 3, 1)) // 
    fmt.Println(minimumCost([]int{9,8,7,6,5,4,3,2,1}, 3, 1)) // 
}