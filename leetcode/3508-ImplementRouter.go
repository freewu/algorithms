package main

// 3508. Implement Router
// Design a data structure that can efficiently manage data packets in a network router. 
// Each data packet consists of the following attributes:
//     source: A unique identifier for the machine that generated the packet.
//     destination: A unique identifier for the target machine.
//     timestamp: The time at which the packet arrived at the router.
// Implement the Router class:
//     Router(int memoryLimit): Initializes the Router object with a fixed memory limit.
//         1. memoryLimit is the maximum number of packets the router can store at any given time.
//         2. If adding a new packet would exceed this limit, the oldest packet must be removed to free up space.
//     bool addPacket(int source, int destination, int timestamp): Adds a packet with the given attributes to the router.
//         1. A packet is considered a duplicate if another packet with the same source, destination, and timestamp already exists in the router.
//         2. Return true if the packet is successfully added (i.e., it is not a duplicate); otherwise return false.
//     int[] forwardPacket(): Forwards the next packet in FIFO (First In First Out) order.
//         1. Remove the packet from storage.
//         2. Return the packet as an array [source, destination, timestamp].
//         3. If there are no packets to forward, return an empty array.
//     int getCount(int destination, int startTime, int endTime):
//         1. Returns the number of packets currently stored in the router (i.e., not yet forwarded) that have the specified destination and have timestamps in the inclusive range [startTime, endTime].

// Note that queries for addPacket will be made in increasing order of timestamp.

// Example 1:
// Input:
// ["Router", "addPacket", "addPacket", "addPacket", "addPacket", "addPacket", "forwardPacket", "addPacket", "getCount"]
// [[3], [1, 4, 90], [2, 5, 90], [1, 4, 90], [3, 5, 95], [4, 5, 105], [], [5, 2, 110], [5, 100, 110]]
// Output:
// [null, true, true, false, true, true, [2, 5, 90], true, 1]
// Explanation
// Router router = new Router(3); // Initialize Router with memoryLimit of 3.
// router.addPacket(1, 4, 90); // Packet is added. Return True.
// router.addPacket(2, 5, 90); // Packet is added. Return True.
// router.addPacket(1, 4, 90); // This is a duplicate packet. Return False.
// router.addPacket(3, 5, 95); // Packet is added. Return True
// router.addPacket(4, 5, 105); // Packet is added, [1, 4, 90] is removed as number of packets exceeds memoryLimit. Return True.
// router.forwardPacket(); // Return [2, 5, 90] and remove it from router.
// router.addPacket(5, 2, 110); // Packet is added. Return True.
// router.getCount(5, 100, 110); // The only packet with destination 5 and timestamp in the inclusive range [100, 110] is [4, 5, 105]. Return 1.

// Example 2:
// Input:
// ["Router", "addPacket", "forwardPacket", "forwardPacket"]
// [[2], [7, 4, 90], [], []]
// Output:
// [null, true, [7, 4, 90], []]
// Explanation
// Router router = new Router(2); // Initialize Router with memoryLimit of 2.
// router.addPacket(7, 4, 90); // Return True.
// router.forwardPacket(); // Return [7, 4, 90].
// router.forwardPacket(); // There are no packets left, return [].
 
// Constraints:
//     2 <= memoryLimit <= 10^5
//     1 <= source, destination <= 2 * 10^5
//     1 <= timestamp <= 10^9
//     1 <= startTime <= endTime <= 10^9
//     At most 10^5 calls will be made to addPacket, forwardPacket, and getCount methods altogether.
//     queries for addPacket will be made in increasing order of timestamp.

import "fmt"
// import "strconv"

// type Router struct {
//     routers     [][]int              // Store all Packets
//     hash        map[string]struct{}  // Duplicate handle
//     des         map[int][]int        // Store destination timestamps
//     memoryLimit int                  // memoryLimit to add into the routers
// }

// func Constructor(memoryLimit int) Router {
//     return Router{
//         routers:     make([][]int, 0),
//         hash:        make(map[string]struct{}),
//         des:         make(map[int][]int),
//         memoryLimit: memoryLimit,
//     }
// }

// // adding packets into routers, des and hash
// func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
//     key := strconv.Itoa(source) + "#" + strconv.Itoa(destination) + "#" + strconv.Itoa(timestamp)
//     // if packet is not presnt in the routers
//     if _, exists := this.hash[key]; !exists {
//         // if routers have memorylimit add directly
//         if len(this.routers) < this.memoryLimit {
//             this.hash[key] = struct{}{}
//             this.routers = append(this.routers, []int{source, destination, timestamp})
//             this.des[destination] = append(this.des[destination], timestamp)
//         } else {
//             // first removing from routers, des and hash because limit exceeds 
//             // and adding the new packet
//             this.ForwardPacket()
//             this.hash[key] = struct{}{}
//             this.routers = append(this.routers, []int{source, destination, timestamp})
//             this.des[destination] = append(this.des[destination], timestamp)
//         }
//         return true
//     }
//     return false
// }

// // Removing in FIFO order from routers, des and hash
// func (this *Router) ForwardPacket() []int {
//     if len(this.routers) == 0 {
//         return nil
//     }
//     packet := this.routers[0]
//     this.routers = this.routers[1:]
//     // Remove from des
//     dest := packet[1]
//     if timestamps, ok := this.des[dest]; ok && len(timestamps) > 0 {
//         this.des[dest] = timestamps[1:]
//     }
//     // Remove from hash
//     key := strconv.Itoa(packet[0]) + "#" + strconv.Itoa(packet[1]) + "#" + strconv.Itoa(packet[2])
//     delete(this.hash, key)
//     return packet
// }

// // Search from the des array the timestamp
// func (this *Router) GetCount(destination int, startTime int, endTime int) int {
//     count := 0
//     if timestamps, ok := this.des[destination]; ok {
//         for _, t := range timestamps {
//             if t >= startTime && t <= endTime {
//                 count++
//             }
//         }
//     }
//     return count
// }

import "sort"

type Router struct {
    queue      [][2]int
    head, tail int
    vmPackets  map[int][]int
    packets    map[int]struct{}
}

func Constructor(memoryLimit int) Router {
    return Router{
        queue:     make([][2]int, memoryLimit),
        head:      0,
        tail:      0,
        vmPackets: map[int][]int{},
        packets:   map[int]struct{}{},
    }
}

func hash(source, dest, timestamp int) int {
    return source + dest*1e6 + timestamp*1e12
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
    packetID := hash(source, destination, timestamp)
    if _, ok := this.packets[packetID]; ok {
        return false
    }
    n := len(this.queue)
    if this.tail-this.head == n {
        // remove oldest
        item := this.queue[this.head%len(this.queue)]
        timestamp := this.vmPackets[item[0]][0]
        this.head++
        this.vmPackets[item[0]] = this.vmPackets[item[0]][1:]
        packetID := hash(item[1], item[0], timestamp)
        delete(this.packets, packetID)
    }
    this.queue[this.tail%n] = [2]int{destination, source}
    this.tail++
    this.packets[packetID] = struct{}{}
    this.vmPackets[destination] = append(this.vmPackets[destination], timestamp)
    return true
}

func (this *Router) ForwardPacket() []int {
    if this.head == this.tail {
        return nil
    }
    item := this.queue[this.head%len(this.queue)]
    this.head++
    timestamp := this.vmPackets[item[0]][0]
    this.vmPackets[item[0]] = this.vmPackets[item[0]][1:]
    packetID := hash(item[1], item[0], timestamp)
    delete(this.packets, packetID)
    return []int{item[1], item[0], timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
    packets := this.vmPackets[destination]
    left := sort.SearchInts(packets, startTime)
    right := sort.Search(len(packets), func(i int) bool {
        return packets[i] > endTime
    })
    return right - left
}

func main() {
    // Example 1:
    // Input:
    // ["Router", "addPacket", "addPacket", "addPacket", "addPacket", "addPacket", "forwardPacket", "addPacket", "getCount"]
    // [[3], [1, 4, 90], [2, 5, 90], [1, 4, 90], [3, 5, 95], [4, 5, 105], [], [5, 2, 110], [5, 100, 110]]
    // Output:
    // [null, true, true, false, true, true, [2, 5, 90], true, 1]
    // Explanation
    // Router router = new Router(3); // Initialize Router with memoryLimit of 3.
    obj1 := Constructor(3)
    fmt.Println(obj1)
    // router.addPacket(1, 4, 90); // Packet is added. Return True.
    fmt.Println(obj1.AddPacket(1, 4, 90)) // true
    fmt.Println(obj1)
    // router.addPacket(2, 5, 90); // Packet is added. Return True.
    fmt.Println(obj1.AddPacket(2, 5, 90)) // true
    fmt.Println(obj1)
    // router.addPacket(1, 4, 90); // This is a duplicate packet. Return False.
    fmt.Println(obj1.AddPacket(1, 4, 90)) // false duplicate packet
    fmt.Println(obj1)
    // router.addPacket(3, 5, 95); // Packet is added. Return True
    fmt.Println(obj1.AddPacket(3, 5, 95)) // true
    fmt.Println(obj1)
    // router.addPacket(4, 5, 105); // Packet is added, [1, 4, 90] is removed as number of packets exceeds memoryLimit. Return True.
    fmt.Println(obj1.AddPacket(4, 5, 105)) // true
    fmt.Println(obj1)
    // router.forwardPacket(); // Return [2, 5, 90] and remove it from router.
    fmt.Println(obj1.ForwardPacket()) // [2, 5, 90]
    // router.addPacket(5, 2, 110); // Packet is added. Return True.
    fmt.Println(obj1.AddPacket(5, 2, 110)) // true
    fmt.Println(obj1)
    // router.getCount(5, 100, 110); // The only packet with destination 5 and timestamp in the inclusive range [100, 110] is [4, 5, 105]. Return 1.
    fmt.Println(obj1.GetCount(5, 100, 110)) // 1

    // Example 2:
    // Input:
    // ["Router", "addPacket", "forwardPacket", "forwardPacket"]
    // [[2], [7, 4, 90], [], []]
    // Output:
    // [null, true, [7, 4, 90], []]
    // Explanation
    // Router router = new Router(2); // Initialize Router with memoryLimit of 2.
    obj2 := Constructor(2)
    fmt.Println(obj2)
    // router.addPacket(7, 4, 90); // Return True.
    fmt.Println(obj2.AddPacket(7, 4, 90)) // true
    fmt.Println(obj2)
    // router.forwardPacket(); // Return [7, 4, 90].
    fmt.Println(obj2.ForwardPacket()) // [7, 4, 90]
    // router.forwardPacket(); // There are no packets left, return [].
    fmt.Println(obj2.ForwardPacket()) // []
}