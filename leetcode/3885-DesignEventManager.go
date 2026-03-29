package main

// 3885. Design Event Manager
// You are given an initial list of events, where each event has a unique eventId and a priority.

// Implement the EventManager class:
//     1. EventManager(int[][] events) 
//         Initializes the manager with the given events, where events[i] = [eventIdi, priority​​​​​​​i].
//     2. void updatePriority(int eventId, int newPriority) 
//         Updates the priority of the active event with id eventId to newPriority.
//     3. int pollHighest() 
//         Removes and returns the eventId of the active event with the highest priority. 
//         If multiple active events have the same priority, return the smallest eventId among them. 
//         If there are no active events, return -1.

// An event is called active if it has not been removed by pollHighest().

// Example 1:
// Input:
// ["EventManager", "pollHighest", "updatePriority", "pollHighest", "pollHighest"]
// [[[[5, 7], [2, 7], [9, 4]]], [], [9, 7], [], []]
// Output:
// [null, 2, null, 5, 9]
// Explanation
// EventManager eventManager = new EventManager([[5,7], [2,7], [9,4]]); // Initializes the manager with three events
// eventManager.pollHighest(); // both events 5 and 2 have priority 7, so return the smaller id 2
// eventManager.updatePriority(9, 7); // event 9 now has priority 7
// eventManager.pollHighest(); // remaining highest priority events are 5 and 9, return 5
// eventManager.pollHighest(); // return 9

// Example 2:
// Input:
// ["EventManager", "pollHighest", "pollHighest", "pollHighest"]
// [[[[4, 1], [7, 2]]], [], [], []]
// Output:
// [null, 7, 4, -1]
// Explanation
// EventManager eventManager = new EventManager([[4,1], [7,2]]); // Initializes the manager with two events
// eventManager.pollHighest(); // return 7
// eventManager.pollHighest(); // return 4
// eventManager.pollHighest(); // no events remain, return -1

// Constraints:
//     1 <= events.length <= 10^5
//     events[i] = [eventId, priority]
//     1 <= eventId <= 10^9
//     1 <= priority <= 10^9
//     All the values of eventId in events are unique.
//     1 <= newPriority <= 109
//     For every call to updatePriority, eventId refers to an active event.
//     At most 10^5 calls in total will be made to updatePriority and pollHighest.

import "fmt"
import "container/heap"

type Event struct{ priority, id int }
type MinHeap []Event
func (h MinHeap) Len() int      { return len(h) }
func (h MinHeap) Less(i, j int) bool {
    return h[i].priority > h[j].priority || h[i].priority == h[j].priority && h[i].id < h[j].id
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v any)   { *h = append(*h, v.(Event)) }
func (h *MinHeap) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type EventManager struct {
	mp map[int]int
	h  *MinHeap
}

func Constructor(events [][]int) EventManager {
    n := len(events)
    mp := make(map[int]int, n) // 预分配空间
    h := make(MinHeap, n)
    for i, e := range events {
        id, p := e[0], e[1]
        mp[id] = p
        h[i] = Event{p, id}
    }
    heap.Init(&h)
    return EventManager{mp, &h}
}

func (m EventManager) UpdatePriority(eventId, newPriority int) {
    m.mp[eventId] = newPriority
    heap.Push(m.h, Event{newPriority, eventId})
}

func (m EventManager) PollHighest() int {
    for m.h.Len() > 0 {
        e := heap.Pop(m.h).(Event)
        if m.mp[e.id] == e.priority {
            delete(m.mp, e.id)
            return e.id
        }
        // 不对，继续找下一个
    }
    return -1
}

func main() {
    // Example 1:
    // Input:
    // ["EventManager", "pollHighest", "updatePriority", "pollHighest", "pollHighest"]
    // [[[[5, 7], [2, 7], [9, 4]]], [], [9, 7], [], []]
    // Output:
    // [null, 2, null, 5, 9]
    // Explanation
    // EventManager eventManager = new EventManager([[5,7], [2,7], [9,4]]); // Initializes the manager with three events
    // eventManager.pollHighest(); // both events 5 and 2 have priority 7, so return the smaller id 2
    // eventManager.updatePriority(9, 7); // event 9 now has priority 7
    // eventManager.pollHighest(); // remaining highest priority events are 5 and 9, return 5
    // eventManager.pollHighest(); // return 9
    obj1 := Constructor([][]int{{5,7}, {2,7}, {9,4}})
    fmt.Println(obj1.PollHighest()) // 2
    obj1.UpdatePriority(9, 7)
    fmt.Println(obj1.PollHighest()) // 5
    fmt.Println(obj1.PollHighest()) // 9

    // Example 2:
    // Input:
    // ["EventManager", "pollHighest", "pollHighest", "pollHighest"]
    // [[[[4, 1], [7, 2]]], [], [], []]
    // Output:
    // [null, 7, 4, -1]
    // Explanation
    // EventManager eventManager = new EventManager([[4,1], [7,2]]); // Initializes the manager with two events
    // eventManager.pollHighest(); // return 7
    // eventManager.pollHighest(); // return 4
    // eventManager.pollHighest(); // no events remain, return -1
    obj2 := Constructor([][]int{{4,1}, {7,2}})
    fmt.Println(obj2.PollHighest()) // 7
    fmt.Println(obj2.PollHighest()) // 4
    fmt.Println(obj2.PollHighest()) // -1
}