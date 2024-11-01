package main

// 2251. Number of Flowers in Full Bloom
// You are given a 0-indexed 2D integer array flowers, 
// where flowers[i] = [starti, endi] means the ith flower will be in full bloom from starti to endi (inclusive). 
// You are also given a 0-indexed integer array people of size n, where people[i] is the time that the ith person will arrive to see the flowers.

// Return an integer array answer of size n, 
// where answer[i] is the number of flowers that are in full bloom when the ith person arrives.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/02/ex1new.jpg" />
// Input: flowers = [[1,6],[3,7],[9,12],[4,13]], people = [2,3,7,11]
// Output: [1,2,2,2]
// Explanation: The figure above shows the times when the flowers are in full bloom and when the people arrive.
// For each person, we return the number of flowers in full bloom during their arrival.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/02/ex2new.jpg" />
// Input: flowers = [[1,10],[3,3]], people = [3,3,2]
// Output: [2,2,1]
// Explanation: The figure above shows the times when the flowers are in full bloom and when the people arrive.
// For each person, we return the number of flowers in full bloom during their arrival.

// Constraints:
//     1 <= flowers.length <= 5 * 10^4
//     flowers[i].length == 2
//     1 <= starti <= endi <= 10^9
//     1 <= people.length <= 5 * 10^4
//     1 <= people[i] <= 10^9

import "fmt"
import "sort"
import "container/heap"

// 二分法
func fullBloomFlowers(flowers [][]int, people []int) []int {
    res, startBloom, endBloom := []int{}, []int{}, []int{}
    for i := 0; i < len(flowers); i++ {
        startBloom = append(startBloom, flowers[i][0])
    }
    for i := 0; i < len(flowers); i++ {
        endBloom = append(endBloom, flowers[i][1] + 1) // +1 because flowers stop blooming at end+1, not at end
    }
    sort.Slice(startBloom, func(i, j int) bool {
        return startBloom[i] < startBloom[j]
    })
    sort.Slice(endBloom, func(i, j int) bool {
        return endBloom[i] < endBloom[j]
    })
    UpperBound := func(arr []int, target int) int {
        l, r := 0, len(arr) - 1
        for l <= r {
            mid := (r + l) / 2
            if arr[mid] > target {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
        return l
    }
    for _, person := range people {
        countBlooming, countStoppedBlooming := UpperBound(startBloom, person), UpperBound(endBloom, person)
        res = append(res, countBlooming - countStoppedBlooming)
    }
    return res
}

type MinHeap []int
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Peek() any { return (*h)[0] }
func (h *MinHeap) Push(i any) { *h = append(*h, i.(int)) }
func (h *MinHeap) Pop() (res any) {
    n := len(*h)-1
    res, *h = (*h)[n], (*h)[:n]
    return res
}

// 最小堆
func fullBloomFlowers1(flowers [][]int, people []int) []int {
    sort.Slice(flowers, func(i, j int) bool {
        return flowers[i][0] < flowers[j][0]
    })
    sorted := make([]int, len(people))
    copy(sorted, people)
    sort.Ints(sorted)

    pq := MinHeap{}
    heap.Init(&pq)

    index, mp := 0, make(map[int]int)
    for _, p := range sorted {
        for index < len(flowers) && flowers[index][0] <= p {
            heap.Push(&pq, flowers[index][1])
            index++
        }
        for pq.Len() > 0 && pq.Peek().(int) < p {
            heap.Pop(&pq)
        }
        mp[p] = pq.Len()
    }
    res := []int{}
    for _, p := range people {
        res = append(res, mp[p])
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/02/ex1new.jpg" />
    // Input: flowers = [[1,6],[3,7],[9,12],[4,13]], people = [2,3,7,11]
    // Output: [1,2,2,2]
    // Explanation: The figure above shows the times when the flowers are in full bloom and when the people arrive.
    // For each person, we return the number of flowers in full bloom during their arrival.
    fmt.Println(fullBloomFlowers([][]int{{1,6},{3,7},{9,12},{4,13}}, []int{2,3,7,11})) // [1,2,2,2]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/02/ex2new.jpg" />
    // Input: flowers = [[1,10],[3,3]], people = [3,3,2]
    // Output: [2,2,1]
    // Explanation: The figure above shows the times when the flowers are in full bloom and when the people arrive.
    // For each person, we return the number of flowers in full bloom during their arrival.
    fmt.Println(fullBloomFlowers([][]int{{1,10},{3,3}}, []int{3,3,2})) // [2,2,1]

    fmt.Println(fullBloomFlowers1([][]int{{1,6},{3,7},{9,12},{4,13}}, []int{2,3,7,11})) // [1,2,2,2]
    fmt.Println(fullBloomFlowers1([][]int{{1,10},{3,3}}, []int{3,3,2})) // [2,2,1]
}