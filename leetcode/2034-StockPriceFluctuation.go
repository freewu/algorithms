package main

// 2034. Stock Price Fluctuation
// You are given a stream of records about a particular stock. 
// Each record contains a timestamp and the corresponding price of the stock at that timestamp.

// Unfortunately due to the volatile nature of the stock market, the records do not come in order. 
// Even worse, some records may be incorrect. 
// Another record with the same timestamp may appear later in the stream correcting the price of the previous wrong record.

// Design an algorithm that:
//     Updates the price of the stock at a particular timestamp, correcting the price from any previous records at the timestamp.
//     Finds the latest price of the stock based on the current records. The latest price is the price at the latest timestamp recorded.
//     Finds the maximum price the stock has been based on the current records.
//     Finds the minimum price the stock has been based on the current records.

// Implement the StockPrice class:
//     StockPrice() Initializes the object with no price records.
//     void update(int timestamp, int price) Updates the price of the stock at the given timestamp.
//     int current() Returns the latest price of the stock.
//     int maximum() Returns the maximum price of the stock.
//     int minimum() Returns the minimum price of the stock.
 

// Example 1:
// Input
// ["StockPrice", "update", "update", "current", "maximum", "update", "maximum", "update", "minimum"]
// [[], [1, 10], [2, 5], [], [], [1, 3], [], [4, 2], []]
// Output
// [null, null, null, 5, 10, null, 5, null, 2]
// Explanation
// StockPrice stockPrice = new StockPrice();
// stockPrice.update(1, 10); // Timestamps are [1] with corresponding prices [10].
// stockPrice.update(2, 5);  // Timestamps are [1,2] with corresponding prices [10,5].
// stockPrice.current();     // return 5, the latest timestamp is 2 with the price being 5.
// stockPrice.maximum();     // return 10, the maximum price is 10 at timestamp 1.
// stockPrice.update(1, 3);  // The previous timestamp 1 had the wrong price, so it is updated to 3.
//                           // Timestamps are [1,2] with corresponding prices [3,5].
// stockPrice.maximum();     // return 5, the maximum price is 5 after the correction.
// stockPrice.update(4, 2);  // Timestamps are [1,2,4] with corresponding prices [3,5,2].
// stockPrice.minimum();     // return 2, the minimum price is 2 at timestamp 4.
 
// Constraints:
//     1 <= timestamp, price <= 10^9
//     At most 10^5 calls will be made in total to update, current, maximum, and minimum.
//     current, maximum, and minimum will be called only after update has been called at least once.

import "fmt"
import "container/heap"

type Stock struct {
    Timestamp, Price int
}

type MinHeap []Stock
func (h MinHeap)  Len() int           { return len(h) }
func (h MinHeap)  Less(i, j int) bool { return h[i].Price < h[j].Price }
func (h MinHeap)  Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Stock)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    x := old[len(old)-1]
    *h = old[0:len(old)-1]
    return x
}

type StockPrice struct {
    curTime int
    prices map[int]int
    maxPriceHeap MinHeap
    minPriceHeap MinHeap
}

func Constructor() StockPrice {
    mx, mn := MinHeap{}, MinHeap{}
    heap.Init(&mx)
    heap.Init(&mn)
    return StockPrice{ 0, map[int]int{}, mx, mn, }
}

func (this *StockPrice) Update(timestamp int, price int)  {
    if this.curTime < timestamp { this.curTime = timestamp }
    heap.Push(&this.maxPriceHeap, Stock{timestamp, -1 * price})
    heap.Push(&this.minPriceHeap, Stock{timestamp, price })
    this.prices[timestamp] = price
}

func (this *StockPrice) Current() int {
    return this.prices[this.curTime]   
}

func (this *StockPrice) Maximum() int {
    for this.prices[this.maxPriceHeap[0].Timestamp] != -1 * this.maxPriceHeap[0].Price {
        heap.Pop(&this.maxPriceHeap)
    }
    return -1 * this.maxPriceHeap[0].Price
}

func (this *StockPrice) Minimum() int {
    for this.prices[this.minPriceHeap[0].Timestamp] != this.minPriceHeap[0].Price {
        heap.Pop(&this.minPriceHeap)
    }
    return this.minPriceHeap[0].Price
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */

func main() {
    // Explanation
    // StockPrice stockPrice = new StockPrice();
    obj := Constructor()
    fmt.Println(obj)
    // stockPrice.update(1, 10); // Timestamps are [1] with corresponding prices [10].
    obj.Update(1,10)
    fmt.Println(obj)
    // stockPrice.update(2, 5);  // Timestamps are [1,2] with corresponding prices [10,5].
    obj.Update(2,5)
    fmt.Println(obj)
    // stockPrice.current();     // return 5, the latest timestamp is 2 with the price being 5.
    fmt.Println(obj.Current()) // 5
    // stockPrice.maximum();     // return 10, the maximum price is 10 at timestamp 1.
    fmt.Println(obj.Maximum()) // 10
    // stockPrice.update(1, 3);  // The previous timestamp 1 had the wrong price, so it is updated to 3.
    //                           // Timestamps are [1,2] with corresponding prices [3,5].
    obj.Update(1,3)
    fmt.Println(obj)
    // stockPrice.maximum();     // return 5, the maximum price is 5 after the correction.
    fmt.Println(obj.Maximum()) // 5
    // stockPrice.update(4, 2);  // Timestamps are [1,2,4] with corresponding prices [3,5,2].
    obj.Update(4,2)
    fmt.Println(obj)
    // stockPrice.minimum();     // return 2, the minimum price is 2 at timestamp 4.
    fmt.Println(obj.Minimum()) // 2
}