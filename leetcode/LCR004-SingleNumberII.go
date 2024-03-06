package main

// LCR 004. 只出现一次的数字 II
// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

// 示例 1：
// 输入：nums = [2,2,3,2]
// 输出：3

// 示例 2：
// 输入：nums = [0,1,0,1,0,1,100]
// 输出：100

// 提示：
//         1 <= nums.length <= 3 * 10^4
//         -2^31 <= nums[i] <= 2^31 - 1
//         nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次

// 进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

import "fmt"

func singleNumber(nums []int) int {
    m := make(map[int]int, len(nums) / 3 + 1)
    for _,v := range nums {
        m[v]++
    }
    for k,v := range m {
        // 除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
        if v != 3 {
            return k
        }
    }
    return -1
}

// Space O(1)
// &^ 表示 AND NOT 的意思。这里的 ^ 作为一元操作符，表示按位取反 (^0001 0100 = 1110 1011)，X &^ Y 的意思是将 X 中与 Y 相异的位保留，相同的位清零
func singleNumber1(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}

// best solution
func singleNumber2(nums []int) int {
	var ans int32
	for i := 0; i < 32; i++ {
		cnt := 0
		for _, num := range nums {
			if (num >> i) & 1 == 1 {
				cnt++
			}
		}
		if cnt % 3 == 1 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

func main() {
    fmt.Println(singleNumber([]int{2,2,3,2})) // 3
    fmt.Println(singleNumber([]int{0,1,0,1,0,1,100})) // 100

    fmt.Println(singleNumber1([]int{2,2,3,2})) // 3
    fmt.Println(singleNumber1([]int{0,1,0,1,0,1,100})) // 100

    fmt.Println(singleNumber2([]int{2,2,3,2})) // 3
    fmt.Println(singleNumber2([]int{0,1,0,1,0,1,100})) // 100
}