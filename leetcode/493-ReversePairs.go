package main

// 493. Reverse Pairs
// Given an integer array nums, return the number of reverse pairs in the array.
// A reverse pair is a pair (i, j) where:
//     0 <= i < j < nums.length and
//     nums[i] > 2 * nums[j].
 
// Example 1:
// Input: nums = [1,3,2,3,1]
// Output: 2
// Explanation: The reverse pairs are:
// (1, 4) --> nums[1] = 3, nums[4] = 1, 3 > 2 * 1
// (3, 4) --> nums[3] = 3, nums[4] = 1, 3 > 2 * 1

// Example 2:
// Input: nums = [2,4,3,5,1]
// Output: 3
// Explanation: The reverse pairs are:
// (1, 4) --> nums[1] = 4, nums[4] = 1, 4 > 2 * 1
// (2, 4) --> nums[2] = 3, nums[4] = 1, 3 > 2 * 1
// (3, 4) --> nums[3] = 5, nums[4] = 1, 5 > 2 * 1
 
// Constraints:
//     1 <= nums.length <= 5 * 10^4
//     -2^31 <= nums[i] <= 2^31 - 1

import "fmt"
import "sort"

// 归并排序
func reversePairs(nums []int) int {
    var merge_sort func (nums []int, l, r int) int
    merge_sort = func (nums []int, l, r int) int {
        if l >= r {
            return 0
        }
        m := (l + r) / 2
        cnt := merge_sort(nums, l, m) + merge_sort(nums, m+1, r)
        i, j := l, m + 1 // count reverse pair from sorted list
        // (l) 0 <= i < j < nums.length (r)
        for i <= m && j <= r {
            // nums[i] > 2 * nums[j].
            if nums[i] > nums[j] * 2 {
                cnt += m - i + 1
                j++
            } else {
                i++
            }
        }
        sort.Sort(sort.IntSlice(nums[l : r+1])) // sort the elements from l to r
        return cnt
    }
    return merge_sort(nums, 0, len(nums)-1)
}

// // 树状数组，时间复杂度 O(n log n)
// func reversePairs1(nums []int) (cnt int) {
//     n := len(nums)
//     if n <= 1 {
//         return
//     }
//     // 离散化所有下面统计时会出现的元素
//     allNums := make([]int, 0, 2*n)
//     for _, v := range nums {
//         allNums = append(allNums, v, 2*v)
//     }
//     sort.Ints(allNums)
//     k := 1
//     kth := map[int]int{allNums[0]: k}
//     for i := 1; i < 2*n; i++ {
//         if allNums[i] != allNums[i-1] {
//             k++
//             kth[allNums[i]] = k
//         }
//     }
//     bit := template.BinaryIndexedTree{}
//     bit.Init(k)
//     for i, v := range nums {
//         cnt += i - bit.Query(kth[2*v])
//         bit.Add(kth[v], 1)
//     }
//     return
// }

// //  线段树，时间复杂度 O(n log n)
// func reversePairs2(nums []int) int {
//     if len(nums) < 2 {
//         return 0
//     }
//     st, numsMap, indexMap, numsArray, res := template.SegmentCountTree{}, make(map[int]int, 0), make(map[int]int, 0), []int{}, 0
//     numsMap[nums[0]] = nums[0]
//     for _, num := range nums {
//         numsMap[num] = num
//         numsMap[2*num+1] = 2*num + 1
//     }
//     // numsArray 是 prefixSum 去重之后的版本，利用 numsMap 去重
//     for _, v := range numsMap {
//         numsArray = append(numsArray, v)
//     }
//     // 排序是为了使得线段树中的区间 left <= right，如果此处不排序，线段树中的区间有很多不合法。
//     sort.Ints(numsArray)
//     // 离散化，构建映射关系
//     for i, n := range numsArray {
//         indexMap[n] = i
//     }
//     numsArray = []int{}
//     // 离散化，此题如果不离散化，MaxInt32 的数据会使得数字越界。
//     for i := 0; i < len(indexMap); i++ {
//         numsArray = append(numsArray, i)
//     }
//     // 初始化线段树，节点内的值都赋值为 0，即计数为 0
//     st.Init(numsArray, func(i, j int) int {
//         return 0
//     })
//     for _, num := range nums {
//         res += st.Query(indexMap[num*2+1], len(indexMap)-1)
//         st.UpdateCount(indexMap[num])
//     }
//     return res
// }

func reversePairs3(nums []int) int {
    help := make([]int, len(nums))
    var merge func(nums, help []int, l, m, r int) int
    merge = func(nums, help []int, l, m, r int) int {
        res := 0
        for i, j := l, m+1; i <= m; i++ {
            for j <= r && nums[i] > nums[j]*2 {
                j++
            }
            res += j - m - 1
        }
        // merge process
        i, p, q := l, l, m+1
        for p <= m && q <= r {
            if nums[p] <= nums[q] {
                help[i] = nums[p]
                p++
            } else {
                help[i] = nums[q]
                q++
            }
            i++
        }
        if p <= m {
            copy(help[i:r+1], nums[p:m+1])
        } else {
            copy(help[i:r+1], nums[q:r+1])
        }
        copy(nums[l:r+1], help[l:r+1])
        return res
    }
    var count func(nums, help []int, l, r int) int
    count = func(nums, help []int, l, r int) int {
        if l == r {
            return 0
        }
        m := (r + l) >> 1
        return count(nums, help, l, m) + count(nums, help, m+1, r) + merge(nums, help, l, m, r)
    }
    return count(nums, help, 0, len(nums)-1)
}

func main() {
    // Explanation: The reverse pairs are:
    // (1, 4) --> nums[1] = 3, nums[4] = 1, 3 > 2 * 1
    // (3, 4) --> nums[3] = 3, nums[4] = 1, 3 > 2 * 1
    fmt.Println(reversePairs([]int{1,3,2,3,1})) // 2
    // Explanation: The reverse pairs are:
    // (1, 4) --> nums[1] = 4, nums[4] = 1, 4 > 2 * 1
    // (2, 4) --> nums[2] = 3, nums[4] = 1, 3 > 2 * 1
    // (3, 4) --> nums[3] = 5, nums[4] = 1, 5 > 2 * 1
    fmt.Println(reversePairs([]int{2,4,3,5,1})) // 3

    // fmt.Println(reversePairs1([]int{1,3,2,3,1})) // 2
    // fmt.Println(reversePairs1([]int{2,4,3,5,1})) // 3

    // fmt.Println(reversePairs2([]int{1,3,2,3,1})) // 2
    // fmt.Println(reversePairs2([]int{2,4,3,5,1})) // 3

    fmt.Println(reversePairs3([]int{1,3,2,3,1})) // 2
    fmt.Println(reversePairs3([]int{2,4,3,5,1})) // 3
}