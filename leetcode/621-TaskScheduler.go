package main

// 621. Task Scheduler
// You are given an array of CPU tasks, each represented by letters A to Z, and a cooling time, n. 
// Each cycle or interval allows the completion of one task. 
// Tasks can be completed in any order, but there's a constraint: identical tasks must be separated by at least n intervals due to cooling time.

// ​Return the minimum number of intervals required to complete all tasks.

// Example 1:
// Input: tasks = ["A","A","A","B","B","B"], n = 2
// Output: 8
// Explanation: A possible sequence is: A -> B -> idle -> A -> B -> idle -> A -> B.
// After completing task A, you must wait two cycles before doing A again. 
// The same applies to task B. In the 3rd interval, neither A nor B can be done, so you idle. 
// By the 4th cycle, you can do A again as 2 intervals have passed.

// Example 2:
// Input: tasks = ["A","C","A","B","D","B"], n = 1
// Output: 6
// Explanation: A possible sequence is: A -> B -> C -> D -> A -> B.
// With a cooling interval of 1, you can repeat a task after just one other task.

// Example 3:
// Input: tasks = ["A","A","A", "B","B","B"], n = 3
// Output: 10
// Explanation: A possible sequence is: A -> B -> idle -> idle -> A -> B -> idle -> idle -> A -> B.

// There are only two types of tasks, A and B, which need to be separated by 3 intervals. This leads to idling twice between repetitions of these tasks.

// Constraints:
//     1 <= tasks.length <= 10^4
//     tasks[i] is an uppercase English letter.
//     0 <= n <= 100

import "fmt"
import "container/heap"

// map
func leastInterval(tasks []byte, n int) int {
    max := func(a int, b int) int { if a > b { return a; }; return b; }
    // 统计 tasks 中每个task的出现次数
    hashMap := make(map[byte]int)
    for _, task := range tasks {
        hashMap[task]++
    }

    // 获取出现次数最多的task，作为桶的数量
    bucketCount := 0
    for _, count := range hashMap {
        bucketCount = max(bucketCount, count)
    }

    // 最后一个桶的任务数，取决于是否有多个出现次数最多的task，从而能填充到最后一个桶中；最后一个桶的任务数 = 出现次数最多的task的种类数量
    lastBucketTaskCount := 0
    for _, count := range hashMap {
        if count == bucketCount {
            lastBucketTaskCount++
        }
    }
    // 如果桶没有被填满：
    // 由于最后一个桶不需要考虑冷却时间n，因此我们将其他桶与最后一个桶分开考虑
    // 在桶没有被填满的情况下，总时间只与桶的个数和桶大小（冷却时间）有关；总时间 = (桶个数 - 1) * 桶大小 + 最后一个桶的任务数
    notFullTime := (bucketCount-1)*(n+1) + lastBucketTaskCount

    // 如果桶被填满或者超出了每个桶的容量：
    // 每个任务之间都不存在冷却时间，冷却时间完全被填满了
    // 此时执行任务任务所需的时间，就是任务的数量
    fullTime := len(tasks)

    // 为了保证在桶是否被填满的情况下都能满足要求，因此上述两种情况的较大值就是我们要求的最终结果
    return max(notFullTime, fullTime)
}

// 双优先队列
type priorityQueue[T any] struct {
	data []T
}

func (pq *priorityQueue[T]) Len() int {
	return len(pq.data)
}

func (pq *priorityQueue[T]) Less(i, j int) bool {
	return cmp(pq.data[i], pq.data[j])
}

func cmp(x interface{}, y interface{}) bool {
	switch x.(type) {
	case pair:
		a, b := x.(pair), y.(pair)
		return a.nextTime < b.nextTime
	case pair2:
		a, b := x.(pair2), y.(pair2)
		return a.cnt > b.cnt
	}
	return true
}

func (pq *priorityQueue[T]) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *priorityQueue[T]) Push(x any) {
	pq.data = append(pq.data, x.(T))
}

func (pq *priorityQueue[T]) Pop() any {
	r := pq.data[pq.Len()-1]
	pq.data = pq.data[:pq.Len()-1]
	return r
}

func (pq *priorityQueue[T]) Top() any {
	r := pq.data[0]
	return r
}

// 按时间排序所有任务, 不考虑任务的优先级
type pair struct {
	cnt      int
	nextTime int
	name string
}
type pair2 struct {
	cnt      int
	nextTime int
	name string
}

func leastInterval1(tasks []byte, n int) int {
	hash := make(map[byte]int)

	for _, v := range tasks {
		hash[v]++
	}

	pq := &priorityQueue[pair]{}
	pq2 := &priorityQueue[pair2]{}
	heap.Init(pq)

	for k, v := range hash {
		heap.Push(pq, pair{v, 1, string(k)})
	}

	cur := 0

	for pq.Len() > 0 || pq2.Len() > 0 {
		// 可执行队列为空, 那么我们要拿pq的头顶, 有必要的话需要更新时间
		if pq2.Len() == 0 {
			cur = max(cur, pq.Top().(pair).nextTime)
		} else {
            cur ++
        }
        for pq.Len() > 0 && pq.Top().(pair).nextTime <= cur {
            t := heap.Pop(pq).(pair)
            heap.Push(pq2, pair2{t.cnt, t.nextTime, t.name})
        }
        
		t := heap.Pop(pq2).(pair2)
		if t.cnt > 1 {
			heap.Push(pq, pair{t.cnt - 1, cur + n + 1, t.name})
		}
	}

	return cur
}

// array
func leastInterval2(tasks []byte, n int) int {
    max := func(a int, b int) int { if a > b { return a; }; return b; }
    // 统计不同任务数量
    arr := make([]int, 26)
    for _, task := range tasks {
        arr[task - 'A']++
    }
    // 得到最多的任务数量 和 相同任务数据的
    m, count := 0, 0
    for _, num := range arr {
        if num > m {
            m = num
            count = 1
        } else if num == m {
            count++
        }
    }
    return max(len(tasks), (n + 1) * ( m-1 ) + count)
}


func main() {
    // Explanation: A possible sequence is: A -> B -> idle -> A -> B -> idle -> A -> B.
    // After completing task A, you must wait two cycles before doing A again. 
    // The same applies to task B. In the 3rd interval, neither A nor B can be done, so you idle. 
    // By the 4th cycle, you can do A again as 2 intervals have passed.
    fmt.Println(leastInterval([]byte{'A','A','A','B','B','B'},2)) // 8
    // Explanation: A possible sequence is: A -> B -> C -> D -> A -> B.
    // With a cooling interval of 1, you can repeat a task after just one other task.
    fmt.Println(leastInterval([]byte{'A','C','A','B','D','B'},1)) // 6
    // Explanation: A possible sequence is: A -> B -> idle -> idle -> A -> B -> idle -> idle -> A -> B.
    fmt.Println(leastInterval([]byte{'A','A','A','B','B','B'}, 3)) // 10
    // A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> (待命) -> (待命) -> A -> (待命) -> (待命) -> A
    fmt.Println(leastInterval([]byte{'A','A','A','A','A','A','B','C','D','E','F','G'}, 2)) // 16


    fmt.Println(leastInterval1([]byte{'A','A','A','B','B','B'},2)) // 8
    // Explanation: A possible sequence is: A -> B -> C -> D -> A -> B.
    // With a cooling interval of 1, you can repeat a task after just one other task.
    fmt.Println(leastInterval1([]byte{'A','C','A','B','D','B'},1)) // 6
    // Explanation: A possible sequence is: A -> B -> idle -> idle -> A -> B -> idle -> idle -> A -> B.
    fmt.Println(leastInterval1([]byte{'A','A','A','B','B','B'}, 3)) // 10
    // A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> (待命) -> (待命) -> A -> (待命) -> (待命) -> A
    fmt.Println(leastInterval1([]byte{'A','A','A','A','A','A','B','C','D','E','F','G'}, 2)) // 16

    
    fmt.Println(leastInterval2([]byte{'A','A','A','B','B','B'},2)) // 8
    // Explanation: A possible sequence is: A -> B -> C -> D -> A -> B.
    // With a cooling interval of 1, you can repeat a task after just one other task.
    fmt.Println(leastInterval2([]byte{'A','C','A','B','D','B'},1)) // 6
    // Explanation: A possible sequence is: A -> B -> idle -> idle -> A -> B -> idle -> idle -> A -> B.
    fmt.Println(leastInterval2([]byte{'A','A','A','B','B','B'}, 3)) // 10
    // A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> (待命) -> (待命) -> A -> (待命) -> (待命) -> A
    fmt.Println(leastInterval2([]byte{'A','A','A','A','A','A','B','C','D','E','F','G'}, 2)) // 16
}