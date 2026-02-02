package main

// 3829. Design Ride Sharing System
// A ride sharing system manages ride requests from riders and availability from drivers. 
// Riders request rides, and drivers become available over time. 
// The system should match riders and drivers in the order they arrive.

// Implement the RideSharingSystem class:
//     RideSharingSystem() 
//         Initializes the system.
//     void addRider(int riderId) 
//         Adds a new rider with the given riderId.
//     void addDriver(int driverId) 
//         Adds a new driver with the given driverId.
//     int[] matchDriverWithRider() 
//         Matches the earliest available driver with the earliest waiting rider and removes both of them from the system. 
//         Returns an integer array of size 2 where result = [driverId, riderId] if a match is made. 
//         If no match is available, returns [-1, -1].
//     void cancelRider(int riderId) 
//         Cancels the ride request of the rider with the given riderId if the rider exists and has not yet been matched.
    
// Example 1:
// Input:
// ["RideSharingSystem", "addRider", "addDriver", "addRider", "matchDriverWithRider", "addDriver", "cancelRider", "matchDriverWithRider", "matchDriverWithRider"]
// [[], [3], [2], [1], [], [5], [3], [], []]
// Output:
// [null, null, null, null, [2, 3], null, null, [5, 1], [-1, -1]]
// Explanation
// RideSharingSystem rideSharingSystem = new RideSharingSystem(); // Initializes the system
// rideSharingSystem.addRider(3); // rider 3 joins the queue
// rideSharingSystem.addDriver(2); // driver 2 joins the queue
// rideSharingSystem.addRider(1); // rider 1 joins the queue
// rideSharingSystem.matchDriverWithRider(); // returns [2, 3]
// rideSharingSystem.addDriver(5); // driver 5 becomes available
// rideSharingSystem.cancelRider(3); // rider 3 is already matched, cancel has no effect
// rideSharingSystem.matchDriverWithRider(); // returns [5, 1]
// rideSharingSystem.matchDriverWithRider(); // returns [-1, -1]

// Example 2:
// Input:
// ["RideSharingSystem", "addRider", "addDriver", "addDriver", "matchDriverWithRider", "addRider", "cancelRider", "matchDriverWithRider"]
// [[], [8], [8], [6], [], [2], [2], []]
// Output:
// [null, null, null, null, [8, 8], null, null, [-1, -1]]
// Explanation
// RideSharingSystem rideSharingSystem = new RideSharingSystem(); // Initializes the system
// rideSharingSystem.addRider(8); // rider 8 joins the queue
// rideSharingSystem.addDriver(8); // driver 8 joins the queue
// rideSharingSystem.addDriver(6); // driver 6 joins the queue
// rideSharingSystem.matchDriverWithRider(); // returns [8, 8]
// rideSharingSystem.addRider(2); // rider 2 joins the queue
// rideSharingSystem.cancelRider(2); // rider 2 cancels
// rideSharingSystem.matchDriverWithRider(); // returns [-1, -1]

// Constraints:
//     1 <= riderId, driverId <= 1000
//     Each riderId is unique among riders and is added at most once.
//     Each driverId is unique among drivers and is added at most once.
//     At most 1000 calls will be made in total to addRider​​​​​​​, addDriver, matchDriverWithRider, and cancelRider.

import "fmt"

type RideSharingSystem struct {
    riders        []int
    drivers       []int
    waitingRiders map[int]bool
}

func Constructor() RideSharingSystem {
    return RideSharingSystem{ waitingRiders: map[int]bool{}, riders: []int{}, drivers: []int{}, }
}

func (r *RideSharingSystem) AddRider(riderId int) {
    r.riders = append(r.riders, riderId)
    r.waitingRiders[riderId] = true
}

func (r *RideSharingSystem) AddDriver(driverId int) {
    r.drivers = append(r.drivers, driverId)
}

func (r *RideSharingSystem) MatchDriverWithRider() []int {
    // 弹出队列中的已取消乘客
    for len(r.riders) > 0 && !r.waitingRiders[r.riders[0]] {
        r.riders = r.riders[1:]
    }
    // 没有乘客或者司机
    if len(r.riders) == 0 || len(r.drivers) == 0 {
        return []int{-1, -1}
    }
    // 配对（这里没有删除 waitingRiders 中的乘客编号，面试的话建议写上删除的逻辑）
    res := []int{r.drivers[0], r.riders[0]}
    r.riders = r.riders[1:]
    r.drivers = r.drivers[1:]
    return res
}

func (r *RideSharingSystem) CancelRider(riderId int) {
    delete(r.waitingRiders, riderId)
}

/**
 * Your RideSharingSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRider(riderId);
 * obj.AddDriver(driverId);
 * param_3 := obj.MatchDriverWithRider();
 * obj.CancelRider(riderId);
 */

func main() {
    // Example 1:
    // Input:
    // ["RideSharingSystem", "addRider", "addDriver", "addRider", "matchDriverWithRider", "addDriver", "cancelRider", "matchDriverWithRider", "matchDriverWithRider"]
    // [[], [3], [2], [1], [], [5], [3], [], []]
    // Output:
    // [null, null, null, null, [2, 3], null, null, [5, 1], [-1, -1]]
    // Explanation
    // RideSharingSystem rideSharingSystem = new RideSharingSystem(); // Initializes the system
    obj1 := Constructor()
    // rideSharingSystem.addRider(3); // rider 3 joins the queue
    obj1.AddRider(3)
    fmt.Println(obj1)
    // rideSharingSystem.addDriver(2); // driver 2 joins the queue
    obj1.AddDriver(2)
    fmt.Println(obj1)
    // rideSharingSystem.addRider(1); // rider 1 joins the queue
    obj1.AddRider(1)
    fmt.Println(obj1)
    // rideSharingSystem.matchDriverWithRider(); // returns [2, 3]
    fmt.Println(obj1.MatchDriverWithRider()) // [2, 3]
    // rideSharingSystem.addDriver(5); // driver 5 becomes available
    obj1.AddDriver(5)
    fmt.Println(obj1)
    // rideSharingSystem.cancelRider(3); // rider 3 is already matched, cancel has no effect
    obj1.CancelRider(3)
    fmt.Println(obj1)
    // rideSharingSystem.matchDriverWithRider(); // returns [5, 1]
    fmt.Println(obj1.MatchDriverWithRider()) // [5, 1]
    // rideSharingSystem.matchDriverWithRider(); // returns [-1, -1]
    fmt.Println(obj1.MatchDriverWithRider()) // [-1, -1]

    // Example 2:
    // Input:
    // ["RideSharingSystem", "addRider", "addDriver", "addDriver", "matchDriverWithRider", "addRider", "cancelRider", "matchDriverWithRider"]
    // [[], [8], [8], [6], [], [2], [2], []]
    // Output:
    // [null, null, null, null, [8, 8], null, null, [-1, -1]]
    // Explanation
    // RideSharingSystem rideSharingSystem = new RideSharingSystem(); // Initializes the system
    obj2 := Constructor()
    // rideSharingSystem.addRider(8); // rider 8 joins the queue
    obj2.AddRider(8)
    fmt.Println(obj2)
    // rideSharingSystem.addDriver(8); // driver 8 joins the queue
    obj2.AddDriver(8)
    fmt.Println(obj2)
    // rideSharingSystem.addDriver(6); // driver 6 joins the queue
    obj2.AddDriver(6)
    fmt.Println(obj2)
    // rideSharingSystem.matchDriverWithRider(); // returns [8, 8]
    fmt.Println(obj2.MatchDriverWithRider()) // [8, 8]
    // rideSharingSystem.addRider(2); // rider 2 joins the queue
    obj2.AddRider(2)
    fmt.Println(obj2)
    // rideSharingSystem.cancelRider(2); // rider 2 cancels
    obj2.CancelRider(2)
    fmt.Println(obj2)
    // rideSharingSystem.matchDriverWithRider(); // returns [-1, -1]
    fmt.Println(obj2.MatchDriverWithRider()) // [-1, -1]
    }