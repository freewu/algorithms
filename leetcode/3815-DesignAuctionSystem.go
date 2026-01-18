package main 

// 3815. Design Auction System
// You are asked to design an auction system that manages bids from multiple users in real time.

// Each bid is associated with a userId, an itemId, and a bidAmount.

// Implement the AuctionSystem class:​​​​​​​
//     AuctionSystem(): 
//         Initializes the AuctionSystem object.
//     void addBid(int userId, int itemId, int bidAmount): 
//         Adds a new bid for itemId by userId with bidAmount. 
//         If the same userId already has a bid on itemId, replace it with the new bidAmount.
//     void updateBid(int userId, int itemId, int newAmount): 
//         Updates the existing bid of userId for itemId to newAmount. 
//         It is guaranteed that this bid exists.
//     void removeBid(int userId, int itemId): 
//         Removes the bid of userId for itemId. 
//         It is guaranteed that this bid exists.
//     int getHighestBidder(int itemId): 
//         Returns the userId of the highest bidder for itemId. 
//         If multiple users have the same highest bidAmount, return the user with the highest userId. 
//         If no bids exist for the item, return -1.

// Example 1:
// Input:
// ["AuctionSystem", "addBid", "addBid", "getHighestBidder", "updateBid", "getHighestBidder", "removeBid", "getHighestBidder", "getHighestBidder"]
// [[], [1, 7, 5], [2, 7, 6], [7], [1, 7, 8], [7], [2, 7], [7], [3]]
// Output:
// [null, null, null, 2, null, 1, null, 1, -1]
// Explanation
// AuctionSystem auctionSystem = new AuctionSystem(); // Initialize the Auction system
// auctionSystem.addBid(1, 7, 5); // User 1 bids 5 on item 7
// auctionSystem.addBid(2, 7, 6); // User 2 bids 6 on item 7
// auctionSystem.getHighestBidder(7); // return 2 as User 2 has the highest bid
// auctionSystem.updateBid(1, 7, 8); // User 1 updates bid to 8 on item 7
// auctionSystem.getHighestBidder(7); // return 1 as User 1 now has the highest bid
// auctionSystem.removeBid(2, 7); // Remove User 2's bid on item 7
// auctionSystem.getHighestBidder(7); // return 1 as User 1 is the current highest bidder
// auctionSystem.getHighestBidder(3); // return -1 as no bids exist for item 3
 
// Constraints:
//     1 <= userId, itemId <= 5 * 10^4
//     1 <= bidAmount, newAmount <= 10^9
//     At most 5 * 10^4 total calls to addBid, updateBid, removeBid, and getHighestBidder.
//     The input is generated such that for updateBid and removeBid, the bid from the given userId for the given itemId will be valid.

import "fmt"
import "container/heap"

type HeapPair struct{ bidAmount, userId int }
type Heap []HeapPair
func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
    a, b := h[i], h[j]
    return a.bidAmount > b.bidAmount || a.bidAmount == b.bidAmount && a.userId > b.userId
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(v any)   { *h = append(*h, v.(HeapPair)) }
func (h *Heap) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a) - 1]; return v }

type Pair struct{ userId, itemId int }
type AuctionSystem struct {
    amount map[Pair]int // (userId, itemId) -> bidAmount
    itemH  map[int]*Heap  // itemId -> [(bidAmount, userId)]
}

func Constructor() AuctionSystem {
    return AuctionSystem{map[Pair]int{}, map[int]*Heap{}}
}

func (a AuctionSystem) AddBid(userId, itemId, bidAmount int) {
    a.amount[Pair{userId, itemId}] = bidAmount
    if a.itemH[itemId] == nil {
        a.itemH[itemId] = &Heap{}
    }
    heap.Push(a.itemH[itemId], HeapPair{bidAmount, userId})
}

func (a AuctionSystem) UpdateBid(userId, itemId, newAmount int) {
    a.AddBid(userId, itemId, newAmount)
    // 堆中重复的元素在 GetHighestBidder 中删除（懒更新）
}

func (a AuctionSystem) RemoveBid(userId, itemId int) {
    delete(a.amount, Pair{userId, itemId})
    // 堆中元素在 GetHighestBidder 中删除（懒删除）
}

func (a AuctionSystem) GetHighestBidder(itemId int) int {
    h := a.itemH[itemId]
    if h == nil {
        return -1
    }
    for h.Len() > 0 {
        if (*h)[0].bidAmount == a.amount[Pair{(*h)[0].userId, itemId}] {
            return (*h)[0].userId
        }
        // 货不对板，堆顶出价与实际出价不符
        heap.Pop(h)
    }
    return -1
}

/**
 * Your AuctionSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddBid(userId,itemId,bidAmount);
 * obj.UpdateBid(userId,itemId,newAmount);
 * obj.RemoveBid(userId,itemId);
 * param_4 := obj.GetHighestBidder(itemId);
 */

func main() {
    // Example 1:
    // Input:
    // ["AuctionSystem", "addBid", "addBid", "getHighestBidder", "updateBid", "getHighestBidder", "removeBid", "getHighestBidder", "getHighestBidder"]
    // [[], [1, 7, 5], [2, 7, 6], [7], [1, 7, 8], [7], [2, 7], [7], [3]]
    // Output:
    // [null, null, null, 2, null, 1, null, 1, -1]
    // Explanation
    // AuctionSystem auctionSystem = new AuctionSystem(); // Initialize the Auction system
    obj := Constructor();
    fmt.Println(obj);
    // auctionSystem.addBid(1, 7, 5); // User 1 bids 5 on item 7
    obj.AddBid(1, 7, 5);
    fmt.Println(obj);
    // auctionSystem.addBid(2, 7, 6); // User 2 bids 6 on item 7
    obj.AddBid(2, 7, 6);
    fmt.Println(obj);
    // auctionSystem.getHighestBidder(7); // return 2 as User 2 has the highest bid
    fmt.Println(obj.GetHighestBidder(7)); // 2
    // auctionSystem.updateBid(1, 7, 8); // User 1 updates bid to 8 on item 7
    obj.UpdateBid(1, 7, 8);
    fmt.Println(obj);
    // auctionSystem.getHighestBidder(7); // return 1 as User 1 now has the highest bid
    fmt.Println(obj.GetHighestBidder(7)); // 1
    // auctionSystem.removeBid(2, 7); // Remove User 2's bid on item 7
    obj.RemoveBid(2, 7);
    fmt.Println(obj);
    // auctionSystem.getHighestBidder(7); // return 1 as User 1 is the current highest bidder
    fmt.Println(obj.GetHighestBidder(7)); // 1
    // auctionSystem.getHighestBidder(3); // return -1 as no bids exist for item 3
    fmt.Println(obj.GetHighestBidder(3)); // -1
}