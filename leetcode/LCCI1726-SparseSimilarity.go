package main

// 面试题 17.26. Sparse Similarity LCCI
// The similarity of two documents (each with distinct words) is defined to be the size of the intersection divided by the size of the union. 
// For example, if the documents consist of integers, the similarity of {1, 5, 3} and {1, 7, 2, 3} is 0.4, because the intersection has size 2 and the union has size 5. 
// We have a long list of documents (with distinct values and each with an associated ID) where the similarity is believed to be "sparse". 
// That is, any two arbitrarily selected documents are very likely to have similarity 0. 
// Design an algorithm that returns a list of pairs of document IDs and the associated similarity.

// Input is a 2D array docs, where docs[i] is the document with id i. 
// Return an array of strings, where each string represents a pair of documents with similarity greater than 0. 
// The string should be formatted as  {id1},{id2}: {similarity}, where id1 is the smaller id in the two documents, and similarity is the similarity rounded to four decimal places. 
// You can return the array in any order.

// Example:
// Input: 
// [
//   [14, 15, 100, 9, 3],
//   [32, 1, 9, 3, 5],
//   [15, 29, 2, 6, 8, 7],
//   [7, 10]
// ]
// Output:
// [
//   "0,1: 0.2500",
//   "0,2: 0.1000",
//   "2,3: 0.1429"
// ]
// Note:
//     docs.length <= 500
//     docs[i].length <= 500
//     The number of document pairs with similarity greater than 0 will not exceed 1000.

import "fmt"

func computeSimilarities(docs [][]int) []string {
    n := len(docs)
    mp := make(map[int][]int)
    arr := [][]int{}
    for i:=0; i<n; i++ {
        x := make([]int, n)
        arr = append(arr, x)
    }
    for i := range docs {
        for j := range docs[i] {
            v := docs[i][j]
            if _, ok := mp[v]; !ok {
                mp[v] = []int{ i }
            } else {
                for _, x := range mp[v] {
                    arr[x][i]++
                }
                mp[v] = append(mp[v], i)
            }
        } 
    }
    res := []string{}
    for i := range arr {
        for j := i + 1; j < n; j++ {
            if arr[i][j] > 0 {
                rate := float64(arr[i][j]) / float64(len(docs[i]) + len(docs[j]) - arr[i][j])
                // golang的fmt.Sprintf中的小数截取操作是“四舍六入五成双”，1/32 = 0.03125 ，取四位小数是 0.0312, 因为5前面是偶数，后面没有更小位，所以被舍掉，但是题目是根据“四舍五入”算的
                // 所以在rate后加入极小浮点数1e-9将这种结果进位，不然过不了13号测试结果
                res = append(res, fmt.Sprintf("%d,%d: %.4f", i, j, rate + 1e-9))
            }
        }
    }
    return res
}

func main() {
    // Example:
    // Input: 
    // [
    //   [14, 15, 100, 9, 3],
    //   [32, 1, 9, 3, 5],
    //   [15, 29, 2, 6, 8, 7],
    //   [7, 10]
    // ]
    // Output:
    // [
    //   "0,1: 0.2500",
    //   "0,2: 0.1000",
    //   "2,3: 0.1429"
    // ]
    // Note:
    //     docs.length <= 500
    //     docs[i].length <= 500
    //     The number of document pairs with similarity greater than 0 will not exceed 1000.
    grid1 := [][]int{
        {14, 15, 100, 9, 3},
        {32, 1, 9, 3, 5},
        {15, 29, 2, 6, 8, 7},
        {7, 10},
    }
    fmt.Println(computeSimilarities(grid1)) // ["0,1: 0.2500", "0,2: 0.1000", "2,3: 0.1429"]
}