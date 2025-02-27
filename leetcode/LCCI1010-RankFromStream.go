package main

// 面试题 10.10. Rank from Stream LCCI
// Imagine you are reading in a stream of integers. 
// Periodically, you wish to be able to look up the rank of a number x (the number of values less than or equal to x).
// lmplement the data structures and algorithms to support these operations. 
// That is, implement the method track (int x), which is called when each number is generated, 
// and the method getRankOfNumber(int x), which returns the number of values less than or equal to x.

// Note: This problem is slightly different from the original one in the book.

// Example:
// Input:
// ["StreamRank", "getRankOfNumber", "track", "getRankOfNumber"]
// [[], [1], [0], [0]]
// Output:
// [null,0,null,1]

// Note:
//     x <= 50000
//     The number of calls of both track and getRankOfNumber methods are less than or equal to 2000.

import "fmt"
import "sort"

type StreamRank struct {
    data []int
}

func Constructor() StreamRank {
    return StreamRank{}
}

func (s *StreamRank) Track(x int) {
    index := sort.Search(len(s.data), func(i int) bool {
        return s.data[i] > x
    })
    data := make([]int, len(s.data) + 1)
    copy(data, s.data[:index])
    data[index] = x
    copy(data[index + 1:], s.data[index:])
    s.data = data
}

func (s *StreamRank) GetRankOfNumber(x int) int {
    return sort.Search(len(s.data), func(i int) bool {
        return s.data[i] > x
    })
}

// map
// type StreamRank struct {
//     data map[int]int
// }

// func Constructor() StreamRank {
//     return StreamRank{ data: make(map[int]int) }
// }

// func (this *StreamRank) Track(x int)  {
//     this.data[x]++
// }

// func (this *StreamRank) GetRankOfNumber(x int) int {
//         res := 0
//         for k, v := range this.data {
//             if x >= k {
//                 res += v
//             }
//         }
//         return res
// }

/**
 * Your StreamRank object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Track(x);
 * param_2 := obj.GetRankOfNumber(x);
 */

func main() {
    obj := Constructor()
    fmt.Println(obj)
    fmt.Println(obj.GetRankOfNumber(1)) // 0
    obj.Track(0)
    fmt.Println(obj)
    fmt.Println(obj.GetRankOfNumber(0)) // 1
}