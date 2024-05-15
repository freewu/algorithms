package main

// LCR 057. 存在重复元素 III
// 给你一个整数数组 nums 和两个整数 k 和 t 。
// 请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
// 如果存在则返回 true，不存在返回 false。

// 示例 1：
// 输入：nums = [1,2,3,1], k = 3, t = 0
// 输出：true

// 示例 2：
// 输入：nums = [1,0,1,1], k = 1, t = 2
// 输出：true

// 示例 3：
// 输入：nums = [1,5,9,1,5,9], k = 2, t = 3
// 输出：false
 
// 提示：
//     0 <= nums.length <= 2 * 10^4
//     -2^31 <= nums[i] <= 2^31 - 1
//     0 <= k <= 10^4
//     0 <= t <= 2^31 - 1

import "fmt"

// 桶排序
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
    if k <= 0 || t < 0 || len(nums) < 2 {
        return false
    }
    buckets := map[int]int{}
    for i := 0; i < len(nums); i++ {
        key := nums[i] / (t + 1) // Get the ID of the bucket from element value nums[i] and bucket width t + 1
        if nums[i] < 0 { // -7/9 = 0, but need -7/9 = -1
            key--
        }
        if _, ok := buckets[key]; ok {
            return true
        }
        if v, ok := buckets[key-1]; ok && nums[i]-v <= t { // check the lower bucket, and have to check the value too
            return true
        }
        if v, ok := buckets[key+1]; ok && v-nums[i] <= t { // check the upper bucket, and have to check the value too
            return true
        }
        if len(buckets) >= k { // maintain k size of window
            delete(buckets, nums[i-k]/(t+1))
        }
        buckets[key] = nums[i]
    }
    return false
}

//  滑动窗口 + 剪枝
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
    n := len(nums)
    if n <= 1 {
        return false
    }
    if k <= 0 {
        return false
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < n; i++ {
        count := 0
        for j := i + 1; j < n && count < k; j++ {
            if abs(nums[i]-nums[j]) <= t {
                return true
            }
            count++
        }
    }
    return false
}

func main() {
    fmt.Printf("containsNearbyAlmostDuplicate([]int{1,2,3,1},3,0) = %v\n",containsNearbyAlmostDuplicate([]int{1,2,3,1},3,0)) // true
    fmt.Printf("containsNearbyAlmostDuplicate([]int{1,0,1,1},1,2) = %v\n",containsNearbyAlmostDuplicate([]int{1,0,1,1},1,2)) // true
    fmt.Printf("containsNearbyAlmostDuplicate([]int{1,5,9,1,5,9},2,3) = %v\n",containsNearbyAlmostDuplicate([]int{1,5,9,1,5,9},2,3)) // false

    fmt.Printf("containsNearbyAlmostDuplicate1([]int{1,2,3,1},3,0) = %v\n",containsNearbyAlmostDuplicate1([]int{1,2,3,1},3,0)) // true
    fmt.Printf("containsNearbyAlmostDuplicate1([]int{1,0,1,1},1,2) = %v\n",containsNearbyAlmostDuplicate1([]int{1,0,1,1},1,2)) // true
    fmt.Printf("containsNearbyAlmostDuplicate1([]int{1,5,9,1,5,9},2,3) = %v\n",containsNearbyAlmostDuplicate1([]int{1,5,9,1,5,9},2,3)) // false
}