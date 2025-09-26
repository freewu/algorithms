package main

// 611. Valid Triangle Number
// Given an integer array nums, return the number of triplets chosen from the array that can make triangles if we take them as side lengths of a triangle.

// Example 1:
// Input: nums = [2,2,3,4]
// Output: 3
// Explanation: Valid combinations are: 
// 2,3,4 (using the first 2)
// 2,3,4 (using the second 2)
// 2,2,3

// Example 2:
// Input: nums = [4,2,3,4]
// Output: 4
 
// Constraints:
//     1 <= nums.length <= 1000
//     0 <= nums[i] <= 1000

import "fmt"
import "sort"

func triangleNumber(nums []int) int {
    sort.Ints(nums)
    res, n := 0, len(nums)
    for i := n - 1; i >= 2; i-- {
        l, r := 0, i - 1
        for l < r {
            if nums[l] + nums[r] > nums[i] {
                res += (r - l)
                r--
            } else {
                l++
            }
        }
    }
    return res
}

// 构成有效三角形的条件是任意两边的长度大于第三边
// 对数组进行升序排列后即满足a <= b <= c,即已经满足a + c > b，b + c > a 是一定成立的。
// 所以原问题便转化为只需关心 a + b > c 这一个不等式是否成立即可
func triangleNumber1(nums []int) int {
    sort.Ints(nums)
    res, l := 0, len(nums)
    for i := 0; i < l - 2; i++ {  //此时取i=0,j=1,k=2作为三边，往下遍历，只需要符合i+j>k即可
        k := i + 2
        for j := i + 1; j < l-1 && nums[i] != 0; j++{
            for k < l && nums[k] < nums[i] + nums[j] {
                k++  // 此处的k不需要从最前面重新遍历，当j往前走，k之前的那些数肯定全部都满足，可以减少时间复杂度
            }
            // res += max(k-j, 0)
            // 减去1的原因：距离，当i=1,j=2,k=3假设符合条件，当符合条件的时候，进入第三层循环，此时k++,即k=4,然后退出循环
            // 此时如果只计算k-j则为2，实际只有一个值，当不满足条件的话，k-j也会有一个结果，减去1则不符合也没有结果
            res += k - j - 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,2,3,4]
    // Output: 3
    // Explanation: Valid combinations are: 
    // 2,3,4 (using the first 2)
    // 2,3,4 (using the second 2)
    // 2,2,3
    fmt.Println(triangleNumber([]int{2,2,3,4})) // 3
    // Example 2:
    // Input: nums = [4,2,3,4]
    // Output: 4
    fmt.Println(triangleNumber([]int{4,2,3,4})) // 4

    fmt.Println(triangleNumber([]int{1,2,3,4,5,6,7,8,9})) // 34
    fmt.Println(triangleNumber([]int{9,8,7,6,5,4,3,2,1})) // 34

    fmt.Println(triangleNumber1([]int{2,2,3,4})) // 3
    fmt.Println(triangleNumber1([]int{4,2,3,4})) // 4
    fmt.Println(triangleNumber1([]int{1,2,3,4,5,6,7,8,9})) // 34
    fmt.Println(triangleNumber1([]int{9,8,7,6,5,4,3,2,1})) // 34
}