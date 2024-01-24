package main

// 18. 4Sum
// Given an array nums of n integers, 
// return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

// 		0 <= a, b, c, d < n
// 		a, b, c, and d are distinct.
// 		nums[a] + nums[b] + nums[c] + nums[d] == target

// You may return the answer in any order.

// Example 1:
// Input: nums = [1,0,-1,0,-2,2], target = 0
// Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

// Example 2:
// Input: nums = [2,2,2,2,2], target = 8
// Output: [[2,2,2,2]]

// Constraints:
// 		1 <= nums.length <= 200
// 		-10^9 <= nums[i] <= 10^9
// 		-10^9 <= target <= 10^9

import "sort"
import "fmt"

// 双指针
func fourSum(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	// 用 4 格窗口推进
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		// 
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				fmt.Printf("i = %d j = %d left = %d right = %d \n",i,j,left,right)
				fmt.Printf("nums[i] = %d nums[j] = %d nums[left] = %d nums[right] = %d \n",nums[i],nums[j],nums[left],nums[right])
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}

func main() {
	fmt.Println(fourSum([]int{1,0,-1,0,-2,2},0))  // [[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]
	fmt.Println(fourSum([]int{2,2,2,2,2}, 8)) // [[2 2 2 2]]
}