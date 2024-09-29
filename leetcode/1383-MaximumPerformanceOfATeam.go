package main

// 1383. Maximum Performance of a Team
// You are given two integers n and k and two integer arrays speed and efficiency both of length n. 
// There are n engineers numbered from 1 to n. speed[i] and efficiency[i] represent the speed and efficiency of the ith engineer respectively.

// Choose at most k different engineers out of the n engineers to form a team with the maximum performance.

// The performance of a team is the sum of its engineers' speeds multiplied by the minimum efficiency among its engineers.

// Return the maximum performance of this team. 
// Since the answer can be a huge number, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 2
// Output: 60
// Explanation: 
// We have the maximum performance of the team by selecting engineer 2 (with speed=10 and efficiency=4) and engineer 5 (with speed=5 and efficiency=7). That is, performance = (10 + 5) * min(4, 7) = 60.

// Example 2:
// Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 3
// Output: 68
// Explanation:
// This is the same example as the first but k = 3. We can select engineer 1, engineer 2 and engineer 5 to get the maximum performance of the team. That is, performance = (2 + 10 + 5) * min(5, 4, 7) = 68.

// Example 3:
// Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 4
// Output: 72

// Constraints:
//     1 <= k <= n <= 10^5
//     speed.length == n
//     efficiency.length == n
//     1 <= speed[i] <= 10^5
//     1 <= efficiency[i] <= 10^8

import "fmt"
import "sort"
import "container/heap"
import "slices"

type MinHeap []int

func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Pop() (v interface{}) {
    *h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
    return v
}
func (h *MinHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *MinHeap) Peak() int          { return (*h)[0] }
func (h *MinHeap) Sum() int           { 
    res := 0 
    for _, v := range (*h) {
        res +=v
    }
    return res
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
    type Person struct{
        speed int
        efficiency int
    }
    persons, minHeap := []Person{}, new(MinHeap)
    for i := 0 ; i < len(speed) ; i++{
        persons = append(persons , Person{ speed[i], efficiency[i] })
    }
    sort.Slice(persons, func(i, j int) bool {
        return persons[i].efficiency >= persons[j].efficiency
    })
    res, totalSpeed := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range persons {
        if minHeap.Len() >= k {
            totalSpeed -= heap.Pop(minHeap).(int)
        }
        totalSpeed += v.speed
        heap.Push(minHeap, v.speed)
        currentPerfomance := minHeap.Sum() * v.efficiency
        res = max(res, currentPerfomance)
    }
    return res % 1_000_000_007
}

type SpeedHeap struct{ sort.IntSlice }
func (h *SpeedHeap) Push(v any)   {}
func (h *SpeedHeap) Pop() (v any) { return }

func maxPerformance1(n int, speed []int, efficiency []int, k int) int {
    // 门槛堆 topK + 排序 同LC2542
    // sum(efficiency) <= 10^13 *speed  <= 10^18,不会溢出
    ids := make([]int, n)
    for i := range ids {
        ids[i] = i
    }
    slices.SortFunc(ids, func(a, b int) int { return efficiency[b] - efficiency[a] }) // 按效率逆序排序
    hp := SpeedHeap{make(sort.IntSlice, 0, n)}  // topK小根堆,维护最大k个speed
    res, sum := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, idx := range ids[0:k] {
        sum += speed[idx]
        hp.IntSlice = append(hp.IntSlice, speed[idx])
        res = max(res, sum*efficiency[idx]) // 重大bug!! 不同与LC2542,是最多k个,因为乘法两边随着遍历一个增加,一个减小,ans可能在不满k个时取得
    }
    heap.Init(&hp)
    for _, idx := range ids[k:] {
        // 不会错过最佳答案的做法,具体参见LC2542,
        // 最佳答案的最小值的索引为i,其余索引为a,b,c,d.. 则speed数组在这些索引中挑选topK,那么i必在其中
        // 反证:如果i不在其中, 那么挑选的这topK的sum比产生最近答案的sum大, 而在efficiency中,i本身就是最小的,会产生一个"更好"的答案,违背假设
        if sp := speed[idx]; sp > hp.IntSlice[0] {
            sum += sp - hp.IntSlice[0]
            hp.IntSlice[0] = sp
            heap.Fix(&hp, 0)
            res = max(res, sum * efficiency[idx])
        }
    }
    return res % 1_000_000_007
}

func main() {
    // Example 1:
    // Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 2
    // Output: 60
    // Explanation: 
    // We have the maximum performance of the team by selecting engineer 2 (with speed=10 and efficiency=4) and engineer 5 (with speed=5 and efficiency=7). That is, performance = (10 + 5) * min(4, 7) = 60.
    fmt.Println(maxPerformance(6, []int{2,10,3,1,5,8}, []int{5,4,3,9,7,2}, 2)) // 60
    // Example 2:
    // Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 3
    // Output: 68
    // Explanation:
    // This is the same example as the first but k = 3. We can select engineer 1, engineer 2 and engineer 5 to get the maximum performance of the team. That is, performance = (2 + 10 + 5) * min(5, 4, 7) = 68.
    fmt.Println(maxPerformance(6, []int{2,10,3,1,5,8}, []int{5,4,3,9,7,2}, 3)) // 68
    // Example 3:
    // Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 4
    // Output: 72
    fmt.Println(maxPerformance(6, []int{2,10,3,1,5,8}, []int{5,4,3,9,7,2}, 4)) // 72

    fmt.Println(maxPerformance1(6, []int{2,10,3,1,5,8}, []int{5,4,3,9,7,2}, 2)) // 60
    fmt.Println(maxPerformance1(6, []int{2,10,3,1,5,8}, []int{5,4,3,9,7,2}, 3)) // 68
    fmt.Println(maxPerformance1(6, []int{2,10,3,1,5,8}, []int{5,4,3,9,7,2}, 4)) // 72
}