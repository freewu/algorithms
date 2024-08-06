package main

// LCR 170. 交易逆序对的总数
// 在股票交易中，如果前一天的股价高于后一天的股价，则可以认为存在一个「交易逆序对」。
// 请设计一个程序，输入一段时间内的股票交易记录 record，返回其中存在的「交易逆序对」总数。

// 示例 1:
// 输入：record = [9, 7, 5, 4, 6]
// 输出：8
// 解释：交易中的逆序对为 (9, 7), (9, 5), (9, 4), (9, 6), (7, 5), (7, 4), (7, 6), (5, 4)。

// 限制：
//     0 <= record.length <= 50000

import "fmt"

// 升序归并排序，过程中统计逆序对个数 -- 其实就是求每个数的后面有多少个比其小的数字
func reversePairs(record []int) int {
    // mergeSort 升序归并排序，过程中统计逆序对个数，返回排好序的数组和该数组中逆序对个数
    var mergeSort func(nums []int) ([]int, int)
    mergeSort = func(nums []int) ([]int, int) {
        // 若数组中只剩一个元素则直接返回
        if len(nums) <= 1 {
            return nums, 0
        }
        // 每对将数组平均分成两个子数组，递归进行归并排序
        leftNums, leftReverseNum := mergeSort(nums[:len(nums)/2])
        rightNums, rightReverseNum := mergeSort(nums[len(nums)/2:])
        // 两个子数组排序完成后，合并两个子数组，同时统计逆序对个数，左侧子数组的所有元素在原数组中一定都位于右侧子数组前面，
        // 同时按下标从小到大遍历两个子数组，数组内部已经是排好序的，按照大小依次将左右数组元素添加到结果集，同时统计逆序对个数，
        // 若右侧元素先入结果集，则左侧剩余的所有元素都与它组成逆序对，相反如果左边元素先插入，则说明不存在逆序对
        resultNums := make([]int, 0, len(nums))
        i, j, reverseNum := 0, 0, leftReverseNum+rightReverseNum
        for i < len(leftNums) || j < len(rightNums) {
            if i == len(leftNums) {
                // 若左侧数组已经遍历完成，则直接将右侧数组全部追加
                resultNums = append(resultNums, rightNums[j:]...)
                break
            } else if j == len(rightNums) {
                // 若右侧数组已经遍历完成，则直接将左侧数组全部追加
                resultNums = append(resultNums, leftNums[i:]...)
                break
            } else if leftNums[i] <= rightNums[j] {
                // 若左侧数组先入结果集，说明不存在逆序对，因为说明该数字比右侧所有剩余元素都要小，无法构成逆序对
                resultNums = append(resultNums, leftNums[i])
                i++
            } else {
                resultNums = append(resultNums, rightNums[j])
                j++
                // 若右侧元素先入结果集，说明左侧剩余元素都比其大且都在它左面，都能与其构成逆序对
                reverseNum += len(leftNums) - i
            }
        }
        // 返回当前数组的排序结果以及逆序对个数
        return resultNums, reverseNum
    }
    _, reverseNum := mergeSort(record)
    return reverseNum
}

func main() {
    // 示例 1:
    // 输入：record = [9, 7, 5, 4, 6]
    // 输出：8
    // 解释：交易中的逆序对为 (9, 7), (9, 5), (9, 4), (9, 6), (7, 5), (7, 4), (7, 6), (5, 4)。
    fmt.Println(reversePairs([]int{9, 7, 5, 4, 6})) // 8
}