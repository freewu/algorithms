package main

// 1756. Design Most Recently Used Queue
// Design a queue-like data structure that moves the most recently used element to the end of the queue.

// Implement the MRUQueue class:
//     MRUQueue(int n) 
//         constructs the MRUQueue with n elements: [1,2,3,...,n].
//     int fetch(int k) 
//         moves the kth element (1-indexed) to the end of the queue and returns it.

// Example 1:
// Input:
// ["MRUQueue", "fetch", "fetch", "fetch", "fetch"]
// [[8], [3], [5], [2], [8]]
// Output:
// [null, 3, 6, 2, 2]
// Explanation:
// MRUQueue mRUQueue = new MRUQueue(8); // Initializes the queue to [1,2,3,4,5,6,7,8].
// mRUQueue.fetch(3); // Moves the 3rd element (3) to the end of the queue to become [1,2,4,5,6,7,8,3] and returns it.
// mRUQueue.fetch(5); // Moves the 5th element (6) to the end of the queue to become [1,2,4,5,7,8,3,6] and returns it.
// mRUQueue.fetch(2); // Moves the 2nd element (2) to the end of the queue to become [1,4,5,7,8,3,6,2] and returns it.
// mRUQueue.fetch(8); // The 8th element (2) is already at the end of the queue so just return it.

// Constraints:
//     1 <= n <= 2000
//     1 <= k <= n
//     At most 2000 calls will be made to fetch.
    
// Follow up: Finding an O(n) algorithm per fetch is a bit easy. 
// Can you find an algorithm with a better complexity for each fetch call?

import "fmt"

type MRUQueue struct {
    data []int
}

func Constructor(n int) MRUQueue {
    arr := []int{}
    for i := 1; i <= n; i++ {
        arr = append(arr, i)
    }
    return MRUQueue{ arr }
}

func (this *MRUQueue) Fetch(k int) int {
    k = k - 1
    cur, t2 := this.data[k],  this.data[k+1:]
    t := this.data[:k]
    t = append(t, t2...)
    t = append(t, cur)
    this.data = t
    return cur
}

/**
 * Your MRUQueue object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Fetch(k);
 */

func main() {
    // MRUQueue mRUQueue = new MRUQueue(8); // Initializes the queue to [1,2,3,4,5,6,7,8].
    obj := Constructor(8)
    fmt.Println(obj) // [1,2,3,4,5,6,7,8]
    // mRUQueue.fetch(3); // Moves the 3rd element (3) to the end of the queue to become [1,2,4,5,6,7,8,3] and returns it.
    fmt.Println(obj.Fetch(3)) // 3
    fmt.Println(obj) // [1,2,4,5,6,7,8,3]
    // mRUQueue.fetch(5); // Moves the 5th element (6) to the end of the queue to become [1,2,4,5,7,8,3,6] and returns it.
    fmt.Println(obj.Fetch(5)) // 6
    fmt.Println(obj) // [1,2,4,5,7,8,3,6]
    // mRUQueue.fetch(2); // Moves the 2nd element (2) to the end of the queue to become [1,4,5,7,8,3,6,2] and returns it.
    fmt.Println(obj.Fetch(2)) // 2
    fmt.Println(obj) // [1,4,5,7,8,3,6,2]
    // mRUQueue.fetch(8); // The 8th element (2) is already at the end of the queue so just return it.
    fmt.Println(obj.Fetch(8)) // 2
    fmt.Println(obj) // [1,4,5,7,8,3,6,2]
}