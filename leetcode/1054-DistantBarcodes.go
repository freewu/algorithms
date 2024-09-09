package main

// 1054. Distant Barcodes
// In a warehouse, there is a row of barcodes, where the ith barcode is barcodes[i].

// Rearrange the barcodes so that no two adjacent barcodes are equal. 
// You may return any answer, and it is guaranteed an answer exists.

// Example 1:
// Input: barcodes = [1,1,1,2,2,2]
// Output: [2,1,2,1,2,1]

// Example 2:
// Input: barcodes = [1,1,1,1,2,2,3,3]
// Output: [1,3,1,3,1,2,1,2]

// Constraints:
//     1 <= barcodes.length <= 10000
//     1 <= barcodes[i] <= 10000

import "fmt"
import "sort"

func rearrangeBarcodes(barcodes []int) []int {
    mp := map[int]int{}
    for _, v := range barcodes { // 统计出现频次
        mp[v]++
    }
    arr := make([][]int, 0, len(mp))
    for k, v := range mp {
        arr = append(arr, []int{ k, v })
    }
    sort.Slice(arr, func(i, j int) bool { // 出现次数从大到小排
        return arr[i][1] > arr[j][1]
    })
    index := 0
    for i := 0; i < len(barcodes); i += 2 {
        if arr[index][1] == 0 {
            index++
        }
        arr[index][1]--
        barcodes[i] = arr[index][0]
    }
    for i := 1; i < len(barcodes); i += 2 {
        if arr[index][1] == 0 {
            index++
        }
        arr[index][1]--
        barcodes[i] = arr[index][0]
    }
    return barcodes
}

func rearrangeBarcodes1(barcodes []int) []int {
    n := len(barcodes)
    mp := map[int]int{}
    for _, v := range barcodes { // 统计出现频次
        mp[v]++
    }
    val, freq := 0, 0 
    for k, v := range mp {
        if freq < v {
            val, freq = k, v
        }
    }
    res, i := make([]int, n), 0
    for freq > 0 {
        res[i] = val
        i += 2
        freq--
    }
    delete(mp, val)
    for k, v := range mp {
        for v > 0 {
            if i >= n {
                i = 1
            }
            res[i] = k
            i += 2
            v--
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: barcodes = [1,1,1,2,2,2]
    // Output: [2,1,2,1,2,1]
    fmt.Println(rearrangeBarcodes([]int{1,1,1,2,2,2})) // [2,1,2,1,2,1]
    // Example 2:
    // Input: barcodes = [1,1,1,1,2,2,3,3]
    // Output: [1,3,1,3,1,2,1,2]
    fmt.Println(rearrangeBarcodes([]int{1,1,1,1,2,2,3,3})) // [1,3,1,3,1,2,1,2]

    fmt.Println(rearrangeBarcodes1([]int{1,1,1,2,2,2})) // [2,1,2,1,2,1]
    fmt.Println(rearrangeBarcodes1([]int{1,1,1,1,2,2,3,3})) // [1,3,1,3,1,2,1,2]
}