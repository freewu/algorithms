package main

// LCR 120. 寻找文件副本
// 设备中存有 n 个文件，文件 id 记于数组 documents。
// 若文件 id 相同，则定义为该文件存在副本。请返回任一存在副本的文件 id。

// 示例 1：
// 输入：documents = [2, 5, 3, 0, 5, 0]
// 输出：0 或 5

// 提示：
//     0 ≤ documents[i] ≤ n-1
//     2 <= n <= 100000

import "fmt"

func findRepeatDocument(documents []int) int {
    mp := make(map[int]int)
    for _,v := range documents {
        if mp[v] == 1 {
            return v
        }
        mp[v]++
    }
    return -1
}

func findRepeatDocument1(documents []int) int {
    i := 0
    for i < len(documents) {
        if documents[i] == i {
            i++
            continue
        } else if documents[i] == documents[documents[i]] {
            return documents[i]
        } else {
            documents[documents[i]], documents[i] = documents[i], documents[documents[i]]
        }
    }
    return -1
}

func main() {
    // 示例 1：
    // 输入：documents = [2, 5, 3, 0, 5, 0]
    // 输出：0 或 5
    fmt.Println(findRepeatDocument([]int{2, 5, 3, 0, 5, 0})) // 输出：0 或 5

    fmt.Println(findRepeatDocument1([]int{2, 5, 3, 0, 5, 0})) // 输出：0 或 5
}