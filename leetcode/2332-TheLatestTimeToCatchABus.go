package main

// 2332. The Latest Time to Catch a Bus
// You are given a 0-indexed integer array buses of length n, where buses[i] represents the departure time of the ith bus. 
// You are also given a 0-indexed integer array passengers of length m, where passengers[j] represents the arrival time of the jth passenger. 
// All bus departure times are unique. All passenger arrival times are unique.

// You are given an integer capacity, which represents the maximum number of passengers that can get on each bus.

// When a passenger arrives, they will wait in line for the next available bus. 
// You can get on a bus that departs at x minutes if you arrive at y minutes where y <= x, and the bus is not full. 
// Passengers with the earliest arrival times get on the bus first.

// More formally when a bus arrives, either:
//     If capacity or fewer passengers are waiting for a bus, they will all get on the bus, or
//     The capacity passengers with the earliest arrival times will get on the bus.

// Return the latest time you may arrive at the bus station to catch a bus. 
// You cannot arrive at the same time as another passenger.

// Note: The arrays buses and passengers are not necessarily sorted.

// Example 1:
// Input: buses = [10,20], passengers = [2,17,18,19], capacity = 2
// Output: 16
// Explanation: Suppose you arrive at time 16.
// At time 10, the first bus departs with the 0th passenger. 
// At time 20, the second bus departs with you and the 1st passenger.
// Note that you may not arrive at the same time as another passenger, which is why you must arrive before the 1st passenger to catch the bus.

// Example 2:
// Input: buses = [20,30,10], passengers = [19,13,26,4,25,11,21], capacity = 2
// Output: 20
// Explanation: Suppose you arrive at time 20.
// At time 10, the first bus departs with the 3rd passenger. 
// At time 20, the second bus departs with the 5th and 1st passengers.
// At time 30, the third bus departs with the 0th passenger and you.
// Notice if you had arrived any later, then the 6th passenger would have taken your seat on the third bus.

// Constraints:
//     n == buses.length
//     m == passengers.length
//     1 <= n, m, capacity <= 10^5
//     2 <= buses[i], passengers[i] <= 10^9
//     Each element in buses is unique.
//     Each element in passengers is unique.

import "fmt"
import "sort"

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
    sort.Ints(buses)
    sort.Ints(passengers)
    n, m := len(buses), len(passengers)

    binarySearch := func(arr []int, left, target int) int{
        mid, right := 0, len(arr)
        for left < right {
            mid = (right + left) / 2
            if mid >= len(arr) {
                return mid
            }
            if arr[mid] <= target {
                left = mid + 1
            } else {
                right = mid
            }
        }
        return left
    }
    
    occur := make(map[int]bool, m)
    for _, p := range passengers  {
        occur[p] = true
    }
    left, lastBusNotFull := -1, false
    // find the index of last person who
    // is able to catch a bus
    // "I" must come early than the last person
    for i, bus := range buses {
        // binary search from index left + 1
        // since every person arrives earlier
        // than left is already leaving
        index := binarySearch(passengers, left + 1, bus)
        newPassengers := index - left - 1
        if newPassengers >= capacity {
            newPassengers = capacity
        } else if i == n - 1 {
            lastBusNotFull = true
        }
        left += newPassengers
    }
    // every bus leaves before the earliest person
    // just return the departure time of the last bus
    if left == -1 {
        return buses[n - 1]
    }
    // possible result, try from arrive time 
    // of last person who can catch a bus
    possible := passengers[left] - 1
    if lastBusNotFull {
        // if the last bus leaves with non-zero capacity
        // try from its' departure time
        possible = buses[n - 1]
    }
    for {
        _, ok := occur[possible]
        if ok {
            possible--
        } else {
            break
        }
    }
    return possible
}

func latestTimeCatchTheBus1(buses []int, passengers []int, capacity int) int {
    sort.Ints(buses)
    sort.Ints(passengers)
    position, space := 0, 0
    for _, arrive := range buses {
        space = capacity
        for space > 0 && position < len(passengers) && passengers[position] <= arrive {
            space--
            position++
        }
    }
    position--
    lastCatchTime := buses[len(buses)-1]
    if space <= 0 {
        lastCatchTime = passengers[position]
    }
    for position >= 0 && passengers[position] == lastCatchTime {
        position--
        lastCatchTime--
    }
    return lastCatchTime
}

func latestTimeCatchTheBus2(buses []int, passengers []int, capacity int) int {
    sort.Ints(buses)
    sort.Ints(passengers)
    paxIdx, busIdx, latestTime, count, lastPass := 0, 0, 0, 0, 0
    for paxIdx < len(passengers) && busIdx < len(buses) {
        count = 0
        for count < capacity && paxIdx < len(passengers) && passengers[paxIdx] <= buses[busIdx] {
            if passengers[paxIdx] > 1 && (passengers[paxIdx] - passengers[lastPass] > 1 || paxIdx==0) {
                latestTime = passengers[paxIdx]-1
            }
            lastPass = paxIdx
            paxIdx++
            count++
        }
        if count < capacity && passengers[lastPass] != buses[busIdx]  {
            latestTime = buses[busIdx]
        }
        busIdx++
    }
    if busIdx < len(buses) && paxIdx >= len(passengers) {
        latestTime = buses[len(buses)-1]
    }
    return latestTime
}

func main() {
    // Example 1:
    // Input: buses = [10,20], passengers = [2,17,18,19], capacity = 2
    // Output: 16
    // Explanation: Suppose you arrive at time 16.
    // At time 10, the first bus departs with the 0th passenger. 
    // At time 20, the second bus departs with you and the 1st passenger.
    // Note that you may not arrive at the same time as another passenger, which is why you must arrive before the 1st passenger to catch the bus.
    fmt.Println(latestTimeCatchTheBus([]int{10,20}, []int{2,17,18,19}, 2)) // 16
    // Example 2:
    // Input: buses = [20,30,10], passengers = [19,13,26,4,25,11,21], capacity = 2
    // Output: 20
    // Explanation: Suppose you arrive at time 20.
    // At time 10, the first bus departs with the 3rd passenger. 
    // At time 20, the second bus departs with the 5th and 1st passengers.
    // At time 30, the third bus departs with the 0th passenger and you.
    // Notice if you had arrived any later, then the 6th passenger would have taken your seat on the third bus.
    fmt.Println(latestTimeCatchTheBus([]int{20,30,10}, []int{19,13,26,4,25,11,21}, 2)) // 20

    fmt.Println(latestTimeCatchTheBus1([]int{10,20}, []int{2,17,18,19}, 2)) // 16
    fmt.Println(latestTimeCatchTheBus1([]int{20,30,10}, []int{19,13,26,4,25,11,21}, 2)) // 20

    fmt.Println(latestTimeCatchTheBus2([]int{10,20}, []int{2,17,18,19}, 2)) // 16
    fmt.Println(latestTimeCatchTheBus2([]int{20,30,10}, []int{19,13,26,4,25,11,21}, 2)) // 20
}