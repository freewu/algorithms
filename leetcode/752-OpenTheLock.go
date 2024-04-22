package main

// 752. Open the Lock
// You have a lock in front of you with 4 circular wheels. 
// Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. 
// The wheels can rotate freely and wrap around: 
//     for example we can turn '9' to be '0', or '0' to be '9'. 
// Each move consists of turning one wheel one slot.

// The lock initially starts at '0000', a string representing the state of the 4 wheels.

// You are given a list of deadends dead ends, meaning if the lock displays any of these codes, 
// the wheels of the lock will stop turning and you will be unable to open it.

// Given a target representing the value of the wheels that will unlock the lock, 
// return the minimum total number of turns required to open the lock, or -1 if it is impossible.

// Example 1:
// Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
// Output: 6
// Explanation: 
// A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
// Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
// because the wheels of the lock become stuck after the display becomes the dead end "0102".

// Example 2:
// Input: deadends = ["8888"], target = "0009"
// Output: 1
// Explanation: We can turn the last wheel in reverse to move from "0000" -> "0009".

// Example 3:
// Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
// Output: -1
// Explanation: We cannot reach the target without getting stuck.
 
// Constraints:
//     1 <= deadends.length <= 500
//     deadends[i].length == 4
//     target.length == 4
//     target will not be in the list deadends.
//     target and deadends[i] consist of digits only.

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
