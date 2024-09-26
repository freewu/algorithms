package main

// 1845. Seat Reservation Manager
// Design a system that manages the reservation state of n seats that are numbered from 1 to n.

// Implement the SeatManager class:
//     SeatManager(int n) 
//         Initializes a SeatManager object that will manage n seats numbered from 1 to n. 
//         All seats are initially available.
//     int reserve() 
//         Fetches the smallest-numbered unreserved seat, reserves it, and returns its number.
//     void unreserve(int seatNumber) 
//         Unreserves the seat with the given seatNumber.

// Example 1:
// Input
// ["SeatManager", "reserve", "reserve", "unreserve", "reserve", "reserve", "reserve", "reserve", "unreserve"]
// [[5], [], [], [2], [], [], [], [], [5]]
// Output
// [null, 1, 2, null, 2, 3, 4, 5, null]
// Explanation
// SeatManager seatManager = new SeatManager(5); // Initializes a SeatManager with 5 seats.
// seatManager.reserve();    // All seats are available, so return the lowest numbered seat, which is 1.
// seatManager.reserve();    // The available seats are [2,3,4,5], so return the lowest of them, which is 2.
// seatManager.unreserve(2); // Unreserve seat 2, so now the available seats are [2,3,4,5].
// seatManager.reserve();    // The available seats are [2,3,4,5], so return the lowest of them, which is 2.
// seatManager.reserve();    // The available seats are [3,4,5], so return the lowest of them, which is 3.
// seatManager.reserve();    // The available seats are [4,5], so return the lowest of them, which is 4.
// seatManager.reserve();    // The only available seat is seat 5, so return 5.
// seatManager.unreserve(5); // Unreserve seat 5, so now the available seats are [5].

// Constraints:
//     1 <= n <= 10^5
//     1 <= seatNumber <= n
//     For each call to reserve, it is guaranteed that there will be at least one unreserved seat.
//     For each call to unreserve, it is guaranteed that seatNumber will be reserved.
//     At most 10^5 calls in total will be made to reserve and unreserve.

import "fmt"
import "sort"
import "container/heap"

// 超出时间限制 68 / 69 
type SeatManager1 struct {
    seats []bool
}

func Constructor1(n int) SeatManager1 {
    return SeatManager1{ make([]bool, n) }
}

func (this *SeatManager1) Reserve() int {
    for i := 0; i < len(this.seats); i++ {
        if this.seats[i] == false { // 找到第1个空的座位
            this.seats[i] = true
            return i + 1
        }
    }
    return -1
}

func (this *SeatManager1) Unreserve(seatNumber int)  {
    this.seats[seatNumber - 1] = false
}



type SeatManager struct {
    sort.IntSlice // 继承 Len, Less, Swap
    seats int
}

func Constructor(int) SeatManager {
    return SeatManager{}
}

func (this *SeatManager) Reserve() int {
    if len(this.IntSlice) > 0 { // 有空出来的椅子
        return heap.Pop(this).(int) // 坐编号最小的
    }
    this.seats += 1 // 添加一把新的椅子
    return this.seats
}

func (this *SeatManager) Unreserve(seatNumber int) {
    heap.Push(this, seatNumber) // 有人离开了椅子
}

func (this *SeatManager) Push(v any) { 
    this.IntSlice = append(this.IntSlice, v.(int)) 
}

func (this *SeatManager) Pop() any   { 
    arr := this.IntSlice
    res := arr[len(arr)-1]
    this.IntSlice = arr[:len(arr)-1]
    return res 
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */

func main() {
    // SeatManager seatManager = new SeatManager(5); // Initializes a SeatManager with 5 seats.
    obj := Constructor(5)
    fmt.Println(obj)
    // seatManager.reserve();    // All seats are available, so return the lowest numbered seat, which is 1.
    fmt.Println(obj.Reserve()) // 1 
    fmt.Println(obj)
    // seatManager.reserve();    // The available seats are [2,3,4,5], so return the lowest of them, which is 2.
    fmt.Println(obj.Reserve()) // 2
    fmt.Println(obj)
    // seatManager.unreserve(2); // Unreserve seat 2, so now the available seats are [2,3,4,5].
    obj.Unreserve(2)
    fmt.Println(obj)
    // seatManager.reserve();    // The available seats are [2,3,4,5], so return the lowest of them, which is 2.
    fmt.Println(obj.Reserve()) // 2
    fmt.Println(obj)
    // seatManager.reserve();    // The available seats are [3,4,5], so return the lowest of them, which is 3.
    fmt.Println(obj.Reserve()) // 3
    fmt.Println(obj)
    // seatManager.reserve();    // The available seats are [4,5], so return the lowest of them, which is 4.
    fmt.Println(obj.Reserve()) // 4
    fmt.Println(obj)
    // seatManager.reserve();    // The only available seat is seat 5, so return 5.
    fmt.Println(obj.Reserve()) // 5
    fmt.Println(obj)
    // seatManager.unreserve(5); // Unreserve seat 5, so now the available seats are [5].
    obj.Unreserve(5)
    fmt.Println(obj)

    // SeatManager seatManager = new SeatManager(5); // Initializes a SeatManager with 5 seats.
    obj1 := Constructor1(5)
    fmt.Println(obj1)
    // seatManager.reserve();    // All seats are available, so return the lowest numbered seat, which is 1.
    fmt.Println(obj1.Reserve()) // 1 
    fmt.Println(obj1)
    // seatManager.reserve();    // The available seats are [2,3,4,5], so return the lowest of them, which is 2.
    fmt.Println(obj1.Reserve()) // 2
    fmt.Println(obj1)
    // seatManager.unreserve(2); // Unreserve seat 2, so now the available seats are [2,3,4,5].
    obj1.Unreserve(2)
    fmt.Println(obj1)
    // seatManager.reserve();    // The available seats are [2,3,4,5], so return the lowest of them, which is 2.
    fmt.Println(obj1.Reserve()) // 2
    fmt.Println(obj1)
    // seatManager.reserve();    // The available seats are [3,4,5], so return the lowest of them, which is 3.
    fmt.Println(obj1.Reserve()) // 3
    fmt.Println(obj1)
    // seatManager.reserve();    // The available seats are [4,5], so return the lowest of them, which is 4.
    fmt.Println(obj1.Reserve()) // 4
    fmt.Println(obj1)
    // seatManager.reserve();    // The only available seat is seat 5, so return 5.
    fmt.Println(obj1.Reserve()) // 5
    fmt.Println(obj1)
    // seatManager.unreserve(5); // Unreserve seat 5, so now the available seats are [5].
    obj1.Unreserve(5)
    fmt.Println(obj1)
}