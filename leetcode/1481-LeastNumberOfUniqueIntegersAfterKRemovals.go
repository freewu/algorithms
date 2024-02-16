package main

// 1481. Least Number of Unique Integers after K Removals
// Given an array of integers arr and an integer k. 
// Find the least number of unique integers after removing exactly k elements.

// Example 1:
// Input: arr = [5,5,4], k = 1
// Output: 1
// Explanation: Remove the single 4, only 5 is left.

// Example 2:
// Input: arr = [4,3,1,1,3,3,2], k = 3
// Output: 2
// Explanation: Remove 4, 2 and either one of the two 1s or three 3s. 1 and 3 will be left.

// Constraints:
//         1 <= arr.length <= 10^5
//         1 <= arr[i] <= 10^9
//         0 <= k <= arr.length

import "fmt"
import "sort"

// type Pair struct {
//     Key int
//     Value int
// }

// type PairSlice []*Pair

// func (ps PairSlice) Len() int           { return len(ps) }
// func (ps PairSlice) Less(i, j int) bool { return ps[i].Value < ps[j].Value }
// func (ps PairSlice) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

// func (ps *PairSlice) Push(x interface{}) {
//     *ps = append(*ps, x.(*Pair))
// }

// func (ps *PairSlice) Pop() interface{} {
//     n := ps.Len()
//     old := *ps
//     x := old[n - 1]
//     *ps = old[:n - 1]
//     return x
// }


// func findLeastNumOfUniqueInts(arr []int, k int) int {
//     // 生成一个 map  数字 => 出现次数
//     m := map[int]int{}
//     for _, n := range arr {
//         m[n] += 1
//     }

//     var ps PairSlice
//     for k,v := range m {
//         ps = append(ps, &Pair{ Key: k, Value: v })
//     }

//     sort.Sort(sort.Interface(ps))

//     for len(ps) > 0 && k > 0 {
//         p := ps.Pop().(*Pair)
//         fmt.Println(ps)
//         if k >= p.Value {
//             k -= p.Value
//         } else {
//             ps.Push(p)
//             break
//         }
//         fmt.Println(ps)
//     }
//     return len(ps)
// }

func findLeastNumOfUniqueInts(arr []int, k int) int {
    frequencyMap := make(map[int]int)
    
    for i := 0; i < len(arr); i++ {
        frequencyMap[arr[i]]++
    }
    
    var heap [][2]int
    
    for num, frequency := range frequencyMap {
        heap = append(heap, [2]int{num, frequency})
    }
    
    for i := len(heap)/2 - 1; i > -1; i-- {
        heapDown(heap, i, len(heap) - 1)
    }
    
    for i := 0; i < k; i++ {
        heap[0][1]--
        
        if heap[0][1] == 0 {
            heap[0] = heap[len(heap) - 1]
            heap = heap[:len(heap) - 1]
        }
        
        heapDown(heap, 0, len(heap) - 1)
    }
    
    return len(heap)
}

func heapDown(heap [][2]int, pos int, limit int) {
    l, r := 2*pos+1, 2*pos+2
    smaller := pos
    
    if l <= limit && heap[l][1] < heap[smaller][1] {
        smaller = l
    }
    
    if r <= limit && heap[r][1] < heap[smaller][1] {
        smaller = r
    }
    
    if smaller != pos {
        heap[smaller], heap[pos] = heap[pos], heap[smaller]
        heapDown(heap, smaller, limit)
    }
}

// best solution
func findLeastNumOfUniqueInts1(arr []int, k int) int {
    // 生成一个 map  数字 => 出现次数
	count := make(map[int]int)
	for _, v := range arr {
		count[v]++
	}
    // 取出 map 的 value 值 为数组
	minNum := []int{}
	for _, v := range count {
		minNum = append(minNum, v)
	}
    // 排序数组
	sort.Ints(minNum)
    i := 0
	for i = 0; i < len(minNum); i++ {
        // 从数组中恰好移除 k 个元素
		if minNum[i] <= k {
            k -= minNum[i]
            continue
        }
        break
	}
    // 移除后数组中不同整数的最少数目
	return len(minNum[i:])
}

func main() {
    fmt.Println(findLeastNumOfUniqueInts([]int{5,5,4},1)) // 1
    fmt.Println(findLeastNumOfUniqueInts([]int{4,3,1,1,3,3,2},3)) // 2

    fmt.Println(findLeastNumOfUniqueInts1([]int{5,5,4},1)) // 1
    fmt.Println(findLeastNumOfUniqueInts1([]int{4,3,1,1,3,3,2},3)) // 2
} 