package main

// 15. 3Sum
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.

// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation: 
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.

// Example 2:
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.

// Example 3:
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.
 
// Constraints:
//     3 <= nums.length <= 3000
//     -10^5 <= nums[i] <= 10^5

import "fmt"
import "sort"

// 最优解，双指针 + 排序
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
    for index = 1; index < length-1; index++ {
        start, end = 0, length-1
        if index > 1 && nums[index] == nums[index-1] {
            start = index - 1
        }
        for start < index && end > index {
            if start > 0 && nums[start] == nums[start-1] {
                start++
                continue
            }
            if end < length-1 && nums[end] == nums[end+1] {
                end--
                continue
            }
            addNum = nums[start] + nums[end] + nums[index]
            if addNum == 0 {
                res = append(res, []int{nums[start], nums[index], nums[end]})
                start++
                end--
            } else if addNum > 0 {
                end--
            } else {
                start++
            }
        }
    }
    return res
}

func threeSum1(nums []int) [][]int {
    res := [][]int{}
    counter := map[int]int{}
    for _, value := range nums {
        counter[value]++
    }

    uniqNums := []int{}
    for key := range counter {
        uniqNums = append(uniqNums, key)
    }
    sort.Ints(uniqNums)

    for i := 0; i < len(uniqNums); i++ {
        if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
            res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
        }
        for j := i + 1; j < len(uniqNums); j++ {
            if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
                res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
            }
            if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
                res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
            }
            c := 0 - uniqNums[i] - uniqNums[j]
            if c > uniqNums[j] && counter[c] > 0 {
                res = append(res, []int{uniqNums[i], uniqNums[j], c})
            }
        }
    }
    return res
}

// best solution
func threeSum2(nums []int) [][]int {
    var res [][]int
    if len(nums) < 3 {
        return res
    }
    sort.Ints(nums)
    for i := 0; i < len(nums)-1; i++ {
        if i != 0 && nums[i-1] == nums[i] {
            continue
        }
        for k, j := i+1, len(nums)-1; k < j; {
            n := nums[i] + nums[k] + nums[j]
            if n == 0 {
                res = append(res, []int{nums[i], nums[k], nums[j]})
                p := k
                for p < j && nums[p] == nums[k] {
                    p++
                }
                k = p
            } else if n > 0 {
                j--
            } else {
                k++
            }
        }
    }
    return res
}

func main() {
    // fmt.Printf("threeSum([]int{-1, 0, 1, 2, -1, -4}) = %v\n",threeSum([]int{-1, 0, 1, 2, -1, -4}))
    // fmt.Printf("threeSum1([]int{-1, 0, 1, 2, -1, -4}) = %v\n",threeSum1([]int{-1, 0, 1, 2, -1, -4}))
    // fmt.Printf("threeSum2([]int{-1, 0, 1, 2, -1, -4}) = %v\n",threeSum2([]int{-1, 0, 1, 2, -1, -4}))
    
    // nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
    // nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
    // nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
    fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
    fmt.Println(threeSum([]int{0,1,1})) // []
    fmt.Println(threeSum([]int{0,0,0})) // [[0,0,0]]
    fmt.Println(threeSum([]int{0,1,1,-2,3,2,-1})) // [[-2 -1 3] [-2 0 2] [-1 0 1] [-2 1 1]]

    fmt.Println(threeSum1([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
    fmt.Println(threeSum1([]int{0,1,1})) // []
    fmt.Println(threeSum1([]int{0,0,0})) // [[0,0,0]]
    fmt.Println(threeSum1([]int{0,1,1,-2,3,2,-1})) // [[-2 -1 3] [-2 0 2] [-1 0 1] [-2 1 1]]

    fmt.Println(threeSum2([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
    fmt.Println(threeSum2([]int{0,1,1})) // []
    fmt.Println(threeSum2([]int{0,0,0})) // [[0,0,0]]
    fmt.Println(threeSum2([]int{0,1,1,-2,3,2,-1})) // [[-2 -1 3] [-2 0 2] [-1 0 1] [-2 1 1]]

}
