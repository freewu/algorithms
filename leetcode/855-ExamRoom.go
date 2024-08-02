package main

// 855. Exam Room
// There is an exam room with n seats in a single row labeled from 0 to n - 1.

// When a student enters the room, they must sit in the seat that maximizes the distance to the closest person. 
// If there are multiple such seats, they sit in the seat with the lowest number. 
// If no one is in the room, then the student sits at seat number 0.

// Design a class that simulates the mentioned exam room.

// Implement the ExamRoom class:
//     ExamRoom(int n) 
//         Initializes the object of the exam room with the number of the seats n.
//     int seat() 
//         Returns the label of the seat at which the next student will set.
//     void leave(int p) 
//         Indicates that the student sitting at seat p will leave the room. 
//         It is guaranteed that there will be a student sitting at seat p.

// Example 1:
// Input
// ["ExamRoom", "seat", "seat", "seat", "seat", "leave", "seat"]
// [[10], [], [], [], [], [4], []]
// Output
// [null, 0, 9, 4, 2, null, 5]
// Explanation
// ExamRoom examRoom = new ExamRoom(10);
// examRoom.seat(); // return 0, no one is in the room, then the student sits at seat number 0.
// examRoom.seat(); // return 9, the student sits at the last seat number 9.
// examRoom.seat(); // return 4, the student sits at the last seat number 4.
// examRoom.seat(); // return 2, the student sits at the last seat number 2.
// examRoom.leave(4);
// examRoom.seat(); // return 5, the student sits at the last seat number 5.

// Constraints:
//     1 <= n <= 10^9
//     It is guaranteed that there is a student sitting at seat p.
//     At most 10^4 calls will be made to seat and leave.

import "fmt"
import "container/heap"

type ExamRoom struct {
    n int
    seats []int
}

func Constructor(n int) ExamRoom {
    return ExamRoom{n, []int{}}
}

func (this *ExamRoom) Seat() int {
    seats := this.seats
    if len(seats) == 0 {
        this.seats = append(this.seats, 0)
        return 0
    }
    maxDist, seat, behind := seats[0], 0, 0
    for i := 0; i < len(seats) - 1; i++ {
        f, b := seats[i], seats[i + 1]
        curDist := (b - f) / 2
        if curDist > maxDist {
            maxDist = curDist
            seat = f + curDist
            behind = i + 1
        }
    }
    last := seats[len(seats) - 1]
    if this.n - 1 - last > maxDist {
        this.seats = append(this.seats, this.n - 1)
        return this.n - 1
    }
    newSeats := make([]int, len(seats) + 1)
    copy(newSeats, seats[:behind])
    copy(newSeats[behind + 1:], seats[behind:])
    newSeats[behind] = seat
    this.seats = newSeats
    return seat
}

func (this *ExamRoom) Leave(p int)  {
    seats, index := this.seats, 0
    for i := 0; i < len(seats); i++ {
        if seats[i] == p {
            index = i
            break
        }
    }
    newSeats := make([]int, len(seats) - 1)
    copy(newSeats, seats[:index])
    copy(newSeats[index:], seats[index + 1:])
    this.seats = newSeats
}

type interval struct {
    l,r int
}
type myHeap []interval
func (h myHeap) Len() int { return len(h) }
func (h myHeap) Swap(i,j int) { h[i],h[j]=h[j],h[i] }
func (h myHeap) Less(i,j int) bool {
    a:=(h[i].r-h[i].l)/2
    b:=(h[j].r-h[j].l)/2
    if a == b {
        return h[i].l < h[j].l
    }
    return a > b
}
func (h *myHeap) Push(x any) { *h = append(*h, x.(interval))}
func (h *myHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type ExamRoom1 struct {
    h *myHeap
    m map[int]interval
    del map[interval]bool
    cap,size int
}

func Constructor1(n int) ExamRoom1 {
    h := &myHeap{}
    heap.Init(h)
    return ExamRoom1{h , map[int]interval{}, map[interval]bool{}, n, 0}
}

func (this *ExamRoom1) Seat() int {
    if this.size == this.cap {
        return 0
    } else if this.size == 0 {
        this.h = &myHeap{}
        this.m = map[int]interval{}
        this.del = map[interval]bool{}
        this.size++
        this.m[0] = interval{-1, 1 << 32 - 1 }
        heap.Push(this.h, interval{0, (this.cap-1)*2})
        return 0
    } 
    for this.h.Len() > 0 {
        i:=heap.Pop(this.h).(interval)
        if this.del[i] {
            this.del[i]=false
            continue
        }
        mid := i.l+(i.r-i.l)/2
        if i.r < this.cap {
            heap.Push(this.h, interval{mid, i.r})
        }
        if i.l >= 0 {
            heap.Push(this.h, interval{i.l, mid})
        }
        this.m[mid] = i
        this.m[i.l] = interval{this.m[i.l].l, mid}
        this.m[i.r] = interval{mid, this.m[i.r].r}
        this.size++
        return mid
    }
    return 0
}

func (this *ExamRoom1) Leave(p int)  {
    i,ok:=this.m[p]
    if ok {
        a:=interval{i.l, p}
        b:=interval{p, i.r}
        this.del[a]=true
        this.del[b]=true
        nl,nr:=i.l, i.r
        if i.r>=this.cap {
            nr=i.l+(this.cap-1-i.l)*2
        }
        if i.l<0 {
            nl=0-i.r
        }
        this.m[i.l]=interval{this.m[i.l].l, nr}
        this.m[i.r]=interval{nl, this.m[i.r].r}
        heap.Push(this.h, interval{nl, nr})
    }
    delete(this.m, p)
    this.size--
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */

func main() {
    // ExamRoom examRoom = new ExamRoom(10);
    obj := Constructor(10)
    fmt.Println(obj)
    // examRoom.seat(); // return 0, no one is in the room, then the student sits at seat number 0.
    fmt.Println(obj.Seat()) // 0
    fmt.Println(obj)
    // examRoom.seat(); // return 9, the student sits at the last seat number 9.
    fmt.Println(obj.Seat()) // 9
    fmt.Println(obj)
    // examRoom.seat(); // return 4, the student sits at the last seat number 4.
    fmt.Println(obj.Seat()) // 4
    fmt.Println(obj)
    // examRoom.seat(); // return 2, the student sits at the last seat number 2.
    fmt.Println(obj.Seat()) // 2
    fmt.Println(obj)
    // examRoom.leave(4);
    obj.Leave(4)
    fmt.Println(obj)
    // examRoom.seat(); // return 5, the student sits at the last seat number 5.
    fmt.Println(obj.Seat()) // 5
    fmt.Println(obj)


    // ExamRoom examRoom = new ExamRoom(10);
    obj1 := Constructor1(10)
    fmt.Println(obj1)
    // examRoom.seat(); // return 0, no one is in the room, then the student sits at seat number 0.
    fmt.Println(obj1.Seat()) // 0
    fmt.Println(obj1)
    // examRoom.seat(); // return 9, the student sits at the last seat number 9.
    fmt.Println(obj1.Seat()) // 9
    fmt.Println(obj1)
    // examRoom.seat(); // return 4, the student sits at the last seat number 4.
    fmt.Println(obj1.Seat()) // 4
    fmt.Println(obj1)
    // examRoom.seat(); // return 2, the student sits at the last seat number 2.
    fmt.Println(obj1.Seat()) // 2
    fmt.Println(obj1)
    // examRoom.leave(4);
    obj.Leave(4)
    fmt.Println(obj1)
    // examRoom.seat(); // return 5, the student sits at the last seat number 5.
    fmt.Println(obj1.Seat()) // 5
    fmt.Println(obj1)
}