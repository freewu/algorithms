package main

// LCR 119. 最长连续序列
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

// 示例 1：
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

// 示例 2：
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9

// 提示：
//     0 <= nums.length <= 10^4
//     -10^9 <= nums[i] <= 10^9

// 进阶：可以设计并实现时间复杂度为 O(n) 的解决方案吗？

import "fmt"

// # 解题思路
//     要求找出最长连续序列，输出这个最长的长度。要求时间复杂度为 O(n)
//     把每个数都存在 map 中
//     先删去 map 中没有前一个数 nums[i]-1 也没有后一个数 nums[i]+1 的数 nums[i]，这种数前后都不连续
//     然后在 map 中找到前一个数 nums[i]-1 不存在，但是后一个数 nums[i]+1 存在的数，这种数是连续序列的起点，那么不断的往后搜，直到序列“断”了。最后输出最长序列的长度


// 解法一 map，时间复杂度 O(n)
func longestConsecutive(nums []int) int {
    res, mp := 0, map[int]int{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        if mp[v] == 0 { // 如果不存在, 数组中可以存在多个值相同的
            left, right, sum := 0, 0, 0
            if mp[v-1] > 0 { // 判断前一位是否存在
                left = mp[v-1]
            } else {
                left = 0
            }
            if mp[v+1] > 0 { // 判断后一位是否存在
                right = mp[v+1]
            } else {
                right = 0
            }
            // sum: length of the sequence n is in
            sum = left + right + 1 // 计算出连续的值
            mp[v] = sum //
            // keep track of the max length
            res = max(res, sum)
            // extend the length to the boundary(s) of the sequence
            // will do nothing if n has no neighbors
            mp[v-left] = sum // 本算法的重点
            mp[v+right] = sum // 本算法的重点
        } else {
            continue // 再次出现的不再理会了
        }
    }
    return res
}

// 解法二 暴力解法，时间复杂度 O(n^2)
func longestConsecutive1(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    mp, length, tmp, lcs := map[int]bool{}, 0, 0, 0
    for i := 0; i < len(nums); i++ {
        mp[nums[i]] = true // 以值为 key 写入到  map 中
    }
    for key := range mp {
        if !mp[key-1] && !mp[key+1] { // 删除 前后都不连续的数据
            delete(mp, key)
        }
    }
    if len(mp) == 0 { // 如果都被删除完 说明最长为 1   [1,3,5,7] // 这样的数组
        return 1
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for key := range mp {
        if !mp[key-1] && mp[key+1] { // 找到 起点 前一个不存在 & 后一个存在
            length, tmp = 1, key + 1
            for mp[tmp] { // 连续都存在 就一直累加到 不存在下一个为止
                length++
                tmp++
            }
            lcs = max(lcs, length)
        }
    }
    return max(lcs, length)
}

// best solution
func longestConsecutive2(nums []int) int {
    mp := make(map[int]struct{}, len(nums))
    for _, v := range nums {
        mp[v] = struct{}{}
    }
    res := 0
    for v := range mp {
        if _, ok := mp[v-1]; !ok { // 判断前一位是否存在,不存在说明，当前是开始位
            cur := v
            streak := 0
            for ok2 := true; ok2; _, ok2 = mp[cur] { // 如果后一位存在,则累加
                cur += 1
                streak += 1
            }
            if streak > res { // 判读是否是最长值
                res = streak
            }
        }
    }
    return res
}

func main() {
    fmt.Printf("longestConsecutive([]int{ 100,4,200,1,3,2 }) = %v\n",longestConsecutive([]int{ 100,4,200,1,3,2 })) // 4 [1,2,3,4]
    fmt.Printf("longestConsecutive([]int{ 0,3,7,2,5,8,4,6,0,1 }) = %v\n",longestConsecutive([]int{ 0,3,7,2,5,8,4,6,0,1 })) // 9

    fmt.Printf("longestConsecutive1([]int{ 100,4,200,1,3,2 }) = %v\n",longestConsecutive1([]int{ 100,4,200,1,3,2 })) // 4 [1,2,3,4]
    fmt.Printf("longestConsecutive1([]int{ 0,3,7,2,5,8,4,6,0,1 }) = %v\n",longestConsecutive1([]int{ 0,3,7,2,5,8,4,6,0,1 })) // 9

    fmt.Printf("longestConsecutive2([]int{ 100,4,200,1,3,2 }) = %v\n",longestConsecutive2([]int{ 100,4,200,1,3,2 })) // 4 [1,2,3,4]
    fmt.Printf("longestConsecutive2([]int{ 0,3,7,2,5,8,4,6,0,1 }) = %v\n",longestConsecutive2([]int{ 0,3,7,2,5,8,4,6,0,1 })) // 9
}