package main

import "fmt"

/**
128. Longest Consecutive Sequence
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
You must write an algorithm that runs in O(n) time.

Constraints:

	0 <= nums.length <= 10^5
	-10^9 <= nums[i] <= 10^9

Example 1:

	Input: nums = [100,4,200,1,3,2]
	Output: 4
	Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:

	Input: nums = [0,3,7,2,5,8,4,6,0,1]
	Output: 9

# 解题思路
	要求找出最长连续序列，输出这个最长的长度。要求时间复杂度为 O(n)
	把每个数都存在 map 中
	先删去 map 中没有前一个数 nums[i]-1 也没有后一个数 nums[i]+1 的数 nums[i]，这种数前后都不连续
	然后在 map 中找到前一个数 nums[i]-1 不存在，但是后一个数 nums[i]+1 存在的数，这种数是连续序列的起点，那么不断的往后搜，直到序列“断”了。最后输出最长序列的长度
 */

// 解法一 map，时间复杂度 O(n)
func longestConsecutive(nums []int) int {
	res, numMap := 0, map[int]int{}
	for _, num := range nums {
		if numMap[num] == 0 { // 如果不存在, 数组中可以存在多个值相同的
			left, right, sum := 0, 0, 0
			if numMap[num-1] > 0 { // 判断前一位是否存在
				left = numMap[num-1]
			} else {
				left = 0
			}
			if numMap[num+1] > 0 { // 判断后一位是否存在
				right = numMap[num+1]
			} else {
				right = 0
			}
			// sum: length of the sequence n is in
			sum = left + right + 1 // 计算出连续的值
			numMap[num] = sum //
			// keep track of the max length
			res = max(res, sum)
			// extend the length to the boundary(s) of the sequence
			// will do nothing if n has no neighbors
			numMap[num-left] = sum // 本算法的重点
			numMap[num+right] = sum // 本算法的重点
		} else {
			continue // 再次出现的不再理会了
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法二 暴力解法，时间复杂度 O(n^2)
func longestConsecutive1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numMap, length, tmp, lcs := map[int]bool{}, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = true // 以值为 key 写入到  map 中
	}
	for key := range numMap {
		if !numMap[key-1] && !numMap[key+1] { // 删除 前后都不连续的数据
			delete(numMap, key)
		}
	}
	if len(numMap) == 0 { // 如果都被删除完 说明最长为 1   [1,3,5,7] // 这样的数组
		return 1
	}
	for key := range numMap {
		if !numMap[key-1] && numMap[key+1] { // 找到 起点 前一个不存在 & 后一个存在
			length, tmp = 1, key + 1
			for numMap[tmp] { // 连续都存在 就一直累加到 不存在下一个为止
				length++
				tmp++
			}
			lcs = max(lcs, length)
		}
	}
	return max(lcs, length)
}

// best solution
func longestConsecutiveBest(nums []int) int {
	hm := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		hm[v] = struct{}{}
	}
	longestStreak := 0
	for num := range hm {
		if _, ok := hm[num-1]; !ok { // 判断前一位是否存在,不存在说明，当前是开始位
			curNum := num
			curStreak := 0
			for ok2 := true; ok2; _, ok2 = hm[curNum] { // 如果后一位存在,则累加
				curNum += 1
				curStreak += 1
			}
			if curStreak > longestStreak { // 判读是否是最长值
				longestStreak = curStreak
			}
		}
	}
	return longestStreak
}

func main() {
	fmt.Printf("longestConsecutive([]int{ 100,4,200,1,3,2 }) = %v\n",longestConsecutive([]int{ 100,4,200,1,3,2 })) // 4 [1,2,3,4]
	fmt.Printf("longestConsecutive([]int{ 0,3,7,2,5,8,4,6,0,1 }) = %v\n",longestConsecutive([]int{ 0,3,7,2,5,8,4,6,0,1 })) // 9

	fmt.Printf("longestConsecutive1([]int{ 100,4,200,1,3,2 }) = %v\n",longestConsecutive1([]int{ 100,4,200,1,3,2 })) // 4 [1,2,3,4]
	fmt.Printf("longestConsecutive1([]int{ 0,3,7,2,5,8,4,6,0,1 }) = %v\n",longestConsecutive1([]int{ 0,3,7,2,5,8,4,6,0,1 })) // 9

	fmt.Printf("longestConsecutiveBest([]int{ 100,4,200,1,3,2 }) = %v\n",longestConsecutiveBest([]int{ 100,4,200,1,3,2 })) // 4 [1,2,3,4]
	fmt.Printf("longestConsecutiveBest([]int{ 0,3,7,2,5,8,4,6,0,1 }) = %v\n",longestConsecutiveBest([]int{ 0,3,7,2,5,8,4,6,0,1 })) // 9
}
