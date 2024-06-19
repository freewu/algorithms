package main

// 697. Degree of an Array
// Given a non-empty array of non-negative integers nums, 
// the degree of this array is defined as the maximum frequency of any one of its elements.

// Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

// Example 1:
// Input: nums = [1,2,2,3,1]
// Output: 2
// Explanation: 
// The input array has a degree of 2 because both elements 1 and 2 appear twice.
// Of the subarrays that have the same degree:
// [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
// The shortest length is 2. So return 2.

// Example 2:
// Input: nums = [1,2,2,3,1,4,2]
// Output: 6
// Explanation: 
// The degree is 3 because the element 2 is repeated 3 times.
// So [2,2,3,1,4,2] is the shortest subarray, therefore returning 6.
 
// Constraints:
//     nums.length will be between 1 and 50,000.
//     nums[i] will be an integer between 0 and 49,999.

import "fmt"

func findShortestSubArray(nums []int) int {
    tmp := make(map[int][]int)
    degree, res := 0, 1 << 32 - 1
    for i, v := range nums {
        tmp[v] = append(tmp[v], i)
        if len(tmp[v]) > degree { // 得到最大的出现次数
            degree = len(tmp[v])
        }
    }
    for _, v := range tmp {
        if len(v) == degree {
            l := v[len(v)-1] - v[0] + 1
            if l < res {
                res = l
            }
        }
    }
    return res
}

func findShortestSubArray1(nums []int) int {
    type entry struct {
        count int
        left  int
        right int
    }
    element := map[int]entry{} // 首先，先找出数组的度，也就是出现次数最多的数，并将其子串的首与尾的位置记录下来
    for i, v := range nums {
        if e, ok := element[v]; ok { // 判断元素是否在map中出现过
            // 若以前存在就在原基础上修改
            e.count++
            e.right = i
            element[v] = e //更新原有的值
        } else {// 第一次出现的元素，加入 map 中
            element[v] = entry{1, i, i} // 默认是没有值的，所以要将与i给元素更新进去
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res, maxCount := 0, 0
    for _, e := range element { // 其次，，选取最高的度的数之间的子串
        if e.count > maxCount { // 如果度数大于当前最高，就将其子串提取出来
            maxCount = e.count
            res = e.right - e.left + 1 // 子串的最后一位减去第一位
        } else if e.count == maxCount { // 如果等于，就选取其中小的一个
            res = min(res, e.right - e.left + 1)
        }
    }
    return res
}


func main() {
    // Example 1:
    // Input: nums = [1,2,2,3,1]
    // Output: 2
    // Explanation: 
    // The input array has a degree of 2 because both elements 1 and 2 appear twice.
    // Of the subarrays that have the same degree:
    // [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
    // The shortest length is 2. So return 2.
    fmt.Println(findShortestSubArray([]int{1,2,2,3,1})) // 2
    // Example 2:
    // Input: nums = [1,2,2,3,1,4,2]
    // Output: 6
    // Explanation: 
    // The degree is 3 because the element 2 is repeated 3 times.
    // So [2,2,3,1,4,2] is the shortest subarray, therefore returning 6.
    fmt.Println(findShortestSubArray([]int{1,2,2,3,1,4,2})) // 6

    fmt.Println(findShortestSubArray1([]int{1,2,2,3,1})) // 2
    fmt.Println(findShortestSubArray1([]int{1,2,2,3,1,4,2})) // 6
}