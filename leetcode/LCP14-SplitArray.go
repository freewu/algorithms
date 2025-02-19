package main

// LCP 14. 切分数组
// 给定一个整数数组 nums ，小李想将 nums 切割成若干个非空子数组，使得每个子数组最左边的数和最右边的数的最大公约数大于 1 。
// 为了减少他的工作量，请求出最少可以切成多少个子数组。

// 示例 1：
// 输入：nums = [2,3,3,2,3,3]
// 输出：2
// 解释：最优切割为 [2,3,3,2] 和 [3,3] 。第一个子数组头尾数字的最大公约数为 2 ，第二个子数组头尾数字的最大公约数为 3 。

// 示例 2：
// 输入：nums = [2,3,5,7]
// 输出：4
// 解释：只有一种可行的切割：[2], [3], [5], [7]

// 限制：
//     1 <= nums.length <= 10^5
//     2 <= nums[i] <= 10^6

import "fmt"

func splitArray(nums []int) int {
    n := len(nums)
    if n == 1 { return 1 }
    // 1、动态规划-记录位置i的最少切分数
    divides := make([]int, n)
    // 2、难点：记录素数prime的最佳位置， 对于[6, 7, 11, 13, 2]: 6 的素数集: 2和3
    // 则表示约数中含有2和3的数字，可以和位置0的数字6组成子数组，即：
    // indexs[2] = 0, indexs[3] = indexs[7] = 1
    indexs := make(map[int]int)
    // 3、难点：通过构造素数速查表：mp[num] 表示任意数字 num 的最小质数
    mx := 0
    for i := 0; i < n; i++ {
        if nums[i] > mx {
            mx = nums[i]
        }
    }
    mp := make([]int, mx + 10)
    for i := 2; i <= mx; i++ { // 素数筛构造速查表
        if mp[i] == 0 {
            for j := i; j <= mx; j += i {
                if mp[j] == 0 {
                    mp[j] = i
                }
            }
        }
    }
    // 初始化位置0
    divides[0] = 1
    num := nums[0]
    for {
        // a：通过速查表，快速找到数字 num 的最小素数。
        prime := mp[num]
        indexs[prime] = 0
        // b：num / prime后得到新的数字 num2, 继续执行步骤a
        for num % prime == 0 {
            num = num / prime
        }
        if num == 1 { break }
    }
    for i := 1; i < n; i++ {
        divides[i] = i + 1 // 动态更新
        num = nums[i]
        for {
            prime := mp[num]
            if _, ok := indexs[prime]; !ok {
                indexs[prime] = i
            } else if divides[indexs[prime]] > divides[i - 1] + 1 {
                // 难点2：动态更新 indexs，[2, 3, 5, 7, 11, 13, 17, 4, 49, 41, 51]中，
                // 一开始 indexs[7] = 3，之后由于4和49的加入，indexs[7] = 8
                indexs[prime] = i
            }
            curr := indexs[prime]
            if curr == 0 {
                divides[i] = 1
                break
            } else if divides[curr - 1] + 1 < divides[i] {
                divides[i] = divides[curr - 1] + 1
            }
            for num % prime == 0 {
                num = num / prime
            }
            if num <= 1 { break }
        }
    }
    return divides[n - 1]
}

func main() {
    // 示例 1：
    // 输入：nums = [2,3,3,2,3,3]
    // 输出：2
    // 解释：最优切割为 [2,3,3,2] 和 [3,3] 。第一个子数组头尾数字的最大公约数为 2 ，第二个子数组头尾数字的最大公约数为 3 。
    fmt.Println(splitArray([]int{2,3,3,2,3,3})) // 2
    // 示例 2：
    // 输入：nums = [2,3,5,7]
    // 输出：4
    // 解释：只有一种可行的切割：[2], [3], [5], [7]
    fmt.Println(splitArray([]int{2,3,5,7})) // 4
}