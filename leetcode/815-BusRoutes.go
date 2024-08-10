package main

// 815. Bus Routes
// You are given an array routes representing bus routes where routes[i] is a bus route that the ith bus repeats forever.
//     For example, if routes[0] = [1, 5, 7], this means that the 0th bus travels in the sequence 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... forever.

// You will start at the bus stop source (You are not on any bus initially), 
// and you want to go to the bus stop target. You can travel between bus stops by buses only.

// Return the least number of buses you must take to travel from source to target. 
// Return -1 if it is not possible.

// Example 1:
// Input: routes = [[1,2,7],[3,6,7]], source = 1, target = 6
// Output: 2
// Explanation: The best strategy is take the first bus to the bus stop 7, then take the second bus to the bus stop 6.

// Example 2:
// Input: routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
// Output: -1

// Constraints:
//     1 <= routes.length <= 500.
//     1 <= routes[i].length <= 10^5
//     All the values of routes[i] are unique.
//     sum(routes[i].length) <= 10^5
//     0 <= routes[i][j] < 10^6
//     0 <= source, target < 10^6

import "fmt"

// bfs
func numBusesToDestination(routes [][]int, source int, target int) int {
    n := len(routes)
    // Create hashmap using routes, such as bus stop number is used as key
    // and it tells that how many buses travel from that bus stop
    mp := make(map[int][]int)
    for i := 0; i < n; i++ {
        for j := 0; j < len(routes[i]); j++ {
            busstop := routes[i][j]
            buses := mp[busstop]
            buses = append(buses, i)
            mp[busstop] = buses
        }
    }
    // Initialize the two more sets, that one contains visited busstop and another
    // contains visited buses
    mbuses, mstops := make(map[int]bool), make(map[int]bool)
    queue, count := []int{}, 0
    // Push the source to get started
    // NOTE: Queue contains the stops, not buses
    queue = append(queue, source)
    mstops[source] = true
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            busstop := queue[0]
            queue = queue[1:]
            if busstop == target { // If this is the target, return the level
                return count
            }
            buses := mp[busstop] // Bring out all the buses originating from this stop
            // Now loop over all the buses and check all the stations for those buses
            // If any stop is not visited, then put them into the queue
            for _, bus := range buses {
                if _, ok := mbuses[bus]; ok { continue } // If the bus is visited, continue
                for _, stop := range routes[bus] { // Check all the stops for this bus
                    if _, ok := mstops[stop]; ok {
                        continue
                    }
                    // Put this stop into queue to be proccessed and mark it visited
                    queue = append(queue, stop)
                    mstops[stop] = true
                }
                mbuses[bus] = true // Mark the bus as visited
            }
        }
        count += 1 // Increment the count as we are entering into new level
    }
    return -1
}

func main() {
    // Example 1:
    // Input: routes = [[1,2,7],[3,6,7]], source = 1, target = 6
    // Output: 2
    // Explanation: The best strategy is take the first bus to the bus stop 7, then take the second bus to the bus stop 6.
    fmt.Println(numBusesToDestination([][]int{{1,2,7},{3,6,7}}, 1, 6)) // 2
    // Example 2:
    // Input: routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
    // Output: -1
    fmt.Println(numBusesToDestination([][]int{{7,12},{4,5,15},{6},{15,19},{9,12,13}}, 15, 12)) // -1
}