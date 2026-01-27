package main

// 3822. Design Order Management System
// You are asked to design a simple order management system for a trading platform.

// Each order is associated with an orderId, an orderType ("buy" or "sell"), and a price.

// An order is considered active unless it is canceled.

// Implement the OrderManagementSystem class:
//     OrderManagementSystem(): 
//         Initializes the order management system.
//     void addOrder(int orderId, string orderType, int price): 
//         Adds a new active order with the given attributes. 
//         It is guaranteed that orderId is unique.
//     void modifyOrder(int orderId, int newPrice): 
//         Modifies the price of an existing order. 
//         It is guaranteed that the order exists and is active.
//     void cancelOrder(int orderId): 
//         Cancels an existing order. It is guaranteed that the order exists and is active.
//     vector<int> getOrdersAtPrice(string orderType, int price): 
//         Returns the orderIds of all active orders that match the given orderType and price. 
//         If no such orders exist, return an empty list.

// Note: The order of returned orderIds does not matter.

// Example 1:
// Input:
// ["OrderManagementSystem", "addOrder", "addOrder", "addOrder", "getOrdersAtPrice", "modifyOrder", "modifyOrder", "getOrdersAtPrice", "cancelOrder", "cancelOrder", "getOrdersAtPrice"]
// [[], [1, "buy", 1], [2, "buy", 1], [3, "sell", 2], ["buy", 1], [1, 3], [2, 1], ["buy", 1], [3], [2], ["buy", 1]]
// Output:
// [null, null, null, null, [2, 1], null, null, [2], null, null, []]
// Explanation
// OrderManagementSystem orderManagementSystem = new OrderManagementSystem();
// orderManagementSystem.addOrder(1, "buy", 1); // A buy order with ID 1 is added at price 1.
// orderManagementSystem.addOrder(2, "buy", 1); // A buy order with ID 2 is added at price 1.
// orderManagementSystem.addOrder(3, "sell", 2); // A sell order with ID 3 is added at price 2.
// orderManagementSystem.getOrdersAtPrice("buy", 1); // Both buy orders (IDs 1 and 2) are active at price 1, so the result is [2, 1].
// orderManagementSystem.modifyOrder(1, 3); // Order 1 is updated: its price becomes 3.
// orderManagementSystem.modifyOrder(2, 1); // Order 2 is updated, but its price remains 1.
// orderManagementSystem.getOrdersAtPrice("buy", 1); // Only order 2 is still an active buy order at price 1, so the result is [2].
// orderManagementSystem.cancelOrder(3); // The sell order with ID 3 is canceled and removed from active orders.
// orderManagementSystem.cancelOrder(2); // The buy order with ID 2 is canceled and removed from active orders.
// orderManagementSystem.getOrdersAtPrice("buy", 1); // There are no active buy orders left at price 1, so the result is [].

// Constraints:
//     1 <= orderId <= 2000
//     orderId is unique across all orders.
//     orderType is either "buy" or "sell".
//     1 <= price <= 10^9
//     The total number of calls to addOrder, modifyOrder, cancelOrder, and getOrdersAtPrice does not exceed 2000.
//     For modifyOrder and cancelOrder, the specified orderId is guaranteed to exist and be active.

import "fmt"

// 定义订单结构体，存储单个订单的信息
type Order struct {
    orderType string // 订单类型：buy/sell
    price     int    // 订单价格
    isActive  bool   // 是否活跃（未取消）
}

type OrderManagementSystem struct {
    // 核心存储：key=订单ID，value=订单详情
    orderMap map[int]*Order
}

// 构造函数，初始化订单管理系统
func Constructor() OrderManagementSystem {
    return OrderManagementSystem{ make(map[int]*Order) }
}

// 添加新订单（保证orderId唯一）
func (this *OrderManagementSystem) AddOrder(orderId int, orderType string, price int) {
    this.orderMap[orderId] = &Order{
        orderType: orderType,
        price:     price,
        isActive:  true, // 新订单默认活跃
    }
}

// 修改现有活跃订单的价格（保证订单存在且活跃）
func (this *OrderManagementSystem) ModifyOrder(orderId int, newPrice int) {
    if order, exists := this.orderMap[orderId]; exists && order.isActive {
        order.price = newPrice
    }
}

// 取消现有活跃订单（保证订单存在且活跃）
func (this *OrderManagementSystem) CancelOrder(orderId int) {
    if order, exists := this.orderMap[orderId]; exists && order.isActive {
        order.isActive = false
    }
}

// 查询指定类型和价格的所有活跃订单ID
func (this *OrderManagementSystem) GetOrdersAtPrice(orderType string, price int) []int {
    var res []int
    // 遍历所有订单，筛选符合条件的活跃订单
    for orderId, order := range this.orderMap {
        if order.isActive && order.orderType == orderType && order.price == price {
            res = append(res, orderId)
        }
    }
    return res
}

/**
 * Your OrderManagementSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddOrder(orderId,orderType,price);
 * obj.ModifyOrder(orderId,newPrice);
 * obj.CancelOrder(orderId);
 * param_4 := obj.GetOrdersAtPrice(orderType,price);
 */

func main() {
    // OrderManagementSystem orderManagementSystem = new OrderManagementSystem();
    obj := Constructor();
    // orderManagementSystem.addOrder(1, "buy", 1); // A buy order with ID 1 is added at price 1.
    obj.AddOrder(1, "buy", 1);
    fmt.Println(obj);
    // orderManagementSystem.addOrder(2, "buy", 1); // A buy order with ID 2 is added at price 1.
    obj.AddOrder(2, "buy", 1);
    fmt.Println(obj);
    // orderManagementSystem.addOrder(3, "sell", 2); // A sell order with ID 3 is added at price 2.
    obj.AddOrder(3, "sell", 2);
    fmt.Println(obj);
    // orderManagementSystem.getOrdersAtPrice("buy", 1); // Both buy orders (IDs 1 and 2) are active at price 1, so the result is [2, 1].
    fmt.Println(obj.GetOrdersAtPrice("buy", 1)); // [2, 1]
    // orderManagementSystem.modifyOrder(1, 3); // Order 1 is updated: its price becomes 3.
    obj.ModifyOrder(1, 3);
    fmt.Println(obj);
    // orderManagementSystem.modifyOrder(2, 1); // Order 2 is updated, but its price remains 1.
    obj.ModifyOrder(2, 1);
    fmt.Println(obj);
    // orderManagementSystem.getOrdersAtPrice("buy", 1); // Only order 2 is still an active buy order at price 1, so the result is [2].
    fmt.Println(obj.GetOrdersAtPrice("buy", 1)); // [2]
    // orderManagementSystem.cancelOrder(3); // The sell order with ID 3 is canceled and removed from active orders.
    obj.CancelOrder(3);
    fmt.Println(obj);
    // orderManagementSystem.cancelOrder(2); // The buy order with ID 2 is canceled and removed from active orders.
    obj.CancelOrder(2);
    fmt.Println(obj);
    // orderManagementSystem.getOrdersAtPrice("buy", 1); // There are no active buy orders left at price 1, so the result is [].
    fmt.Println(obj.GetOrdersAtPrice("buy", 1)); // []
}