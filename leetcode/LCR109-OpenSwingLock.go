package main

// LCR 109. 打开转盘锁
// 一个密码锁由 4 个环形拨轮组成，每个拨轮都有 10 个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
// 每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。
// 每次旋转都只能旋转一个拨轮的一位数字。

// 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
// 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
// 字符串 target 代表可以解锁的数字，请给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。

// 示例 1:
// 输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
// 输出：6
// 解释：
// 可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
// 注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，因为当拨动到 "0102" 时这个锁就会被锁定。

// 示例 2:
// 输入: deadends = ["8888"], target = "0009"
// 输出：1
// 解释：
// 把最后一位反向旋转一次即可 "0000" -> "0009"。

// 示例 3:
// 输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
// 输出：-1
// 解释：
// 无法旋转到目标数字且不被锁定。

// 示例 4:
// 输入: deadends = ["0000"], target = "8888"
// 输出：-1

// 提示：
//     1 <= deadends.length <= 500
//     deadends[i].length == 4
//     target.length == 4
//     target 不在 deadends 之中
//     target 和 deadends[i] 仅由若干位数字组成

import "fmt"
import "strconv"

func openLock(deadends []string, target string) int {
    visited := make(map[int]struct{}) // 将死亡数字加入已访问集合，避免被处理
    for _, d := range deadends {
        if d == "0000" { // 如果死亡数字包含零值，直接短路
            return -1
        }
        di, _ := strconv.Atoi(d)
        visited[di] = struct{}{}
    }

    // 初始化双向搜索列表
    t, _ := strconv.Atoi(target)
    nums1, nums2 := []int{0},[]int{t}
    tens, step := [4]int{1, 10, 100, 1000}, 0 // 用于后续的数学运算 step 记录步数
    for len(nums1) > 0 && len(nums2) > 0 {
        newNum := make([]int, 0, len(nums1)*2) // 不知道每次扩散几倍，先保守取2倍，不够再自动扩容
        for _, num := range nums1 {
            _, ok := visited[num] // 已经访问过，不处理
            if ok {
                continue
            }
            for _, num2 := range nums2 { // 已经在另一侧的搜索列表里，则表示已经找到最小步数
                if num == num2 {
                    return step
                }
            }
            // 加入已访问列表
            visited[num] = struct{}{}
            // 计算数字各个拨轮上下拨动的结果，加入新的搜索列表
            for i := 0; i < 4; i++ {
                // 取出第 i 位数字
                n := (num / tens[i]) % 10

                // 分别计算它加减一的值
                var p1, m1 int
                if n == 0 {
                    m1 = num + (tens[i] * 9)
                } else {
                    m1 = num - tens[i]
                }
                if n == 9 {
                    p1 = num - (tens[i] * 9)
                } else {
                    p1 = num + tens[i]
                }

                // 如果不在已访问列表里，则加入新的搜索列表
                _, ok := visited[p1]
                if !ok {
                    newNum = append(newNum, p1)
                }
                _, ok = visited[m1]
                if !ok {
                    newNum = append(newNum, m1)
                }
            }
        }
        // 如果新的搜索列表比较大或者已经空了，就从另一个方向开始搜，否则用新列表搜
        if len(newNum) == 0 || len(newNum) > len(nums2) {
            nums1 = nums2
            nums2 = newNum
        } else {
            nums1 = newNum
        }
        // 步数递增
        step++
    }
    return -1
}

func main() {
    // A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
    // Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
    // because the wheels of the lock become stuck after the display becomes the dead end "0102".
    fmt.Println(openLock([]string{"0201","0101","0102","1212","2002"},"0202")) // 6
    // Explanation: We can turn the last wheel in reverse to move from "0000" -> "0009".
    fmt.Println(openLock([]string{"8888"},"0009")) // 1
    //Explanation: We cannot reach the target without getting stuck.
    fmt.Println(openLock([]string{"8887","8889","8878","8898","8788","8988","7888","9888"}, "8888")) // -1
}