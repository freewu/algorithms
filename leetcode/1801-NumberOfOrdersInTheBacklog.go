package main

// 1801. Number of Orders in the Backlog
// You are given a 2D integer array orders, where each orders[i] = [pricei, amounti, orderTypei] denotes 
// that amounti orders have been placed of type orderTypei at the price pricei. 
// The orderTypei is:
//     0 if it is a batch of buy orders, or
//     1 if it is a batch of sell orders.

// Note that orders[i] represents a batch of amounti independent orders with the same price and order type. 
// All orders represented by orders[i] will be placed before all orders represented by orders[i+1] for all valid i.

// There is a backlog that consists of orders that have not been executed. 
// The backlog is initially empty. 
// When an order is placed, the following happens:
//     1. If the order is a buy order, you look at the sell order with the smallest price in the backlog. 
//        If that sell order's price is smaller than or equal to the current buy order's price, 
//        they will match and be executed, and that sell order will be removed from the backlog. 
//        Else, the buy order is added to the backlog.
//     2. Vice versa, if the order is a sell order, you look at the buy order with the largest price in the backlog. 
//        If that buy order's price is larger than or equal to the current sell order's price, they will match and be executed, and that buy order will be removed from the backlog. 
//        Else, the sell order is added to the backlog.

// Return the total amount of orders in the backlog after placing all the orders from the input. 
// Since this number can be large, return it modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/11/ex1.png" />
// Input: orders = [[10,5,0],[15,2,1],[25,1,1],[30,4,0]]
// Output: 6
// Explanation: Here is what happens with the orders:
// - 5 orders of type buy with price 10 are placed. There are no sell orders, so the 5 orders are added to the backlog.
// - 2 orders of type sell with price 15 are placed. There are no buy orders with prices larger than or equal to 15, so the 2 orders are added to the backlog.
// - 1 order of type sell with price 25 is placed. There are no buy orders with prices larger than or equal to 25 in the backlog, so this order is added to the backlog.
// - 4 orders of type buy with price 30 are placed. The first 2 orders are matched with the 2 sell orders of the least price, which is 15 and these 2 sell orders are removed from the backlog. The 3rd order is matched with the sell order of the least price, which is 25 and this sell order is removed from the backlog. Then, there are no more sell orders in the backlog, so the 4th order is added to the backlog.
// Finally, the backlog has 5 buy orders with price 10, and 1 buy order with price 30. So the total number of orders in the backlog is 6.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/03/11/ex2.png" />
// Input: orders = [[7,1000000000,1],[15,3,0],[5,999999995,0],[5,1,1]]
// Output: 999999984
// Explanation: Here is what happens with the orders:
// - 109 orders of type sell with price 7 are placed. There are no buy orders, so the 109 orders are added to the backlog.
// - 3 orders of type buy with price 15 are placed. They are matched with the 3 sell orders with the least price which is 7, and those 3 sell orders are removed from the backlog.
// - 999999995 orders of type buy with price 5 are placed. The least price of a sell order is 7, so the 999999995 orders are added to the backlog.
// - 1 order of type sell with price 5 is placed. It is matched with the buy order of the highest price, which is 5, and that buy order is removed from the backlog.
// Finally, the backlog has (1000000000-3) sell orders with price 7, and (999999995-1) buy orders with price 5. So the total number of orders = 1999999991, which is equal to 999999984 % (109 + 7).

// Constraints:
//     1 <= orders.length <= 10^5
//     orders[i].length == 3
//     1 <= pricei, amounti <= 10^9
//     orderTypei is either 0 or 1.

import "fmt"
import "container/heap"

type Item struct {
    price, amount int
}
type MinHeap []Item
func (hp MinHeap) Len() int { return len(hp) }
func (hp MinHeap) Less(i, j int) bool { return hp[i].price < hp[j].price }
func (hp MinHeap) Swap(i, j int) { hp[i], hp[j] = hp[j], hp[i] }
func (hp *MinHeap) Push(x interface{}) { *hp = append(*hp, x.(Item)) }
func (hp *MinHeap) Pop() interface{} {
    arr := *hp
    *hp = arr[:len(arr) - 1]
    return arr[len(arr) - 1]
}

type MaxHeap []Item
func (hp MaxHeap) Len() int { return len(hp) }
func (hp MaxHeap) Less(i, j int) bool { return hp[i].price > hp[j].price }
func (hp MaxHeap) Swap(i, j int) { hp[i], hp[j] = hp[j], hp[i] }
func (hp *MaxHeap) Push(x interface{}) { *hp = append(*hp, x.(Item)) }
func (hp *MaxHeap) Pop() interface{} {
    arr := *hp
    *hp = arr[:len(arr)-1]
    return arr[len(arr)-1]
}

func getNumberOfBacklogOrders(orders [][]int) int {
    res, buyOrders, sellOrders := 0, MaxHeap{}, MinHeap{}
    for _, order := range orders {
        orderType, amount, price := order[2], order[1], order[0]
        if orderType == 0 {
            for amount > 0 {
                // 满足出售订单的价格小于购买订单, 且出售订单不为空
                if len(sellOrders) > 0 && sellOrders[0].price <= price {
                    preAmount := sellOrders[0].amount // 存储最低出售订单的价格(小根堆堆顶元素的价格)
                    sellOrders[0].amount -= amount // 最低出售订单的数量减去满足条件的购买订单的数量
                    if sellOrders[0].amount <= 0 {// 最低出售订单的数量不够
                        amount -= preAmount // 购买订单的数量 等于 和数量不够的出售订单 之差
                        heap.Pop(&sellOrders) // 将数量不够最低价格出售订单 弹出堆
                    } else if sellOrders[0].amount > 0 {// 最低出售订单的数量绰绰有余
                        amount = 0  // 直接将购买订单的数量设置为0跳出循环  也可以直接break
                    }
                } else { // 没有找到满足条件的出售订单,将购买订单被压入进大根堆, 并跳出for循环
                    heap.Push(&buyOrders, Item{ price, amount })
                    break
                }
                //购买订单的数量如果还>0 则继续和小根堆堆顶元素比较, 直到amount=0或者购买订单被压入进大根堆
            }
        } else if orderType == 1 {  // 和orderType ==0 逻辑类似
            for amount > 0 {
                if len(buyOrders) > 0 && buyOrders[0].price >= price {
                    preAmount := buyOrders[0].amount
                    buyOrders[0].amount -= amount
                    if buyOrders[0].amount <= 0 {
                        amount -= preAmount
                        heap.Pop(&buyOrders)
                    } else if buyOrders[0].amount > 0 {
                        amount = 0
                    }
                } else {
                    heap.Push(&sellOrders, Item{price, amount})
                    break
                }
            }
        }
    }
    for _, sellOrder := range sellOrders { res += sellOrder.amount }
    for _, buyOrder := range buyOrders {  res += buyOrder.amount }
    return res % 1_000_000_007
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/11/ex1.png" />
    // Input: orders = [[10,5,0],[15,2,1],[25,1,1],[30,4,0]]
    // Output: 6
    // Explanation: Here is what happens with the orders:
    // - 5 orders of type buy with price 10 are placed. There are no sell orders, so the 5 orders are added to the backlog.
    // - 2 orders of type sell with price 15 are placed. There are no buy orders with prices larger than or equal to 15, so the 2 orders are added to the backlog.
    // - 1 order of type sell with price 25 is placed. There are no buy orders with prices larger than or equal to 25 in the backlog, so this order is added to the backlog.
    // - 4 orders of type buy with price 30 are placed. The first 2 orders are matched with the 2 sell orders of the least price, which is 15 and these 2 sell orders are removed from the backlog. The 3rd order is matched with the sell order of the least price, which is 25 and this sell order is removed from the backlog. Then, there are no more sell orders in the backlog, so the 4th order is added to the backlog.
    // Finally, the backlog has 5 buy orders with price 10, and 1 buy order with price 30. So the total number of orders in the backlog is 6.
    fmt.Println(getNumberOfBacklogOrders([][]int{{10,5,0},{15,2,1},{25,1,1},{30,4,0}})) // 6
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/03/11/ex2.png" />
    // Input: orders = [[7,1000000000,1],[15,3,0],[5,999999995,0],[5,1,1]]
    // Output: 999999984
    // Explanation: Here is what happens with the orders:
    // - 109 orders of type sell with price 7 are placed. There are no buy orders, so the 109 orders are added to the backlog.
    // - 3 orders of type buy with price 15 are placed. They are matched with the 3 sell orders with the least price which is 7, and those 3 sell orders are removed from the backlog.
    // - 999999995 orders of type buy with price 5 are placed. The least price of a sell order is 7, so the 999999995 orders are added to the backlog.
    // - 1 order of type sell with price 5 is placed. It is matched with the buy order of the highest price, which is 5, and that buy order is removed from the backlog.
    // Finally, the backlog has (1000000000-3) sell orders with price 7, and (999999995-1) buy orders with price 5. So the total number of orders = 1999999991, which is equal to 999999984 % (109 + 7).
    fmt.Println(getNumberOfBacklogOrders([][]int{{7,1000000000,1},{15,3,0},{5,999999995,0},{5,1,1}})) // 999999984
}