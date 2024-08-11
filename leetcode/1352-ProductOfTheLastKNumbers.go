package main

// 1352. Product of the Last K Numbers
// Design an algorithm that accepts a stream of integers and retrieves the product of the last k integers of the stream.

// Implement the ProductOfNumbers class:
//     ProductOfNumbers() 
//         Initializes the object with an empty stream.
//     void add(int num) 
//         Appends the integer num to the stream.
//     int getProduct(int k) 
//         Returns the product of the last k numbers in the current list. 
//         You can assume that always the current list has at least k numbers.

// The test cases are generated so that, at any time, 
// the product of any contiguous sequence of numbers will fit into a single 32-bit integer without overflowing.

// Example:
// Input
// ["ProductOfNumbers","add","add","add","add","add","getProduct","getProduct","getProduct","add","getProduct"]
// [[],[3],[0],[2],[5],[4],[2],[3],[4],[8],[2]]
// Output
// [null,null,null,null,null,null,20,40,0,null,32]
// Explanation
// ProductOfNumbers productOfNumbers = new ProductOfNumbers();
// productOfNumbers.add(3);        // [3]
// productOfNumbers.add(0);        // [3,0]
// productOfNumbers.add(2);        // [3,0,2]
// productOfNumbers.add(5);        // [3,0,2,5]
// productOfNumbers.add(4);        // [3,0,2,5,4]
// productOfNumbers.getProduct(2); // return 20. The product of the last 2 numbers is 5 * 4 = 20
// productOfNumbers.getProduct(3); // return 40. The product of the last 3 numbers is 2 * 5 * 4 = 40
// productOfNumbers.getProduct(4); // return 0. The product of the last 4 numbers is 0 * 2 * 5 * 4 = 0
// productOfNumbers.add(8);        // [3,0,2,5,4,8]
// productOfNumbers.getProduct(2); // return 32. The product of the last 2 numbers is 4 * 8 = 32 

// Constraints:
//     0 <= num <= 100
//     1 <= k <= 4 * 10^4
//     At most 4 * 10^4 calls will be made to add and getProduct.
//     The product of the stream at any point in time will fit in a 32-bit integer.

import "fmt"

type ProductOfNumbers struct {
    data []int
    product int
}

func Constructor() ProductOfNumbers {
    return ProductOfNumbers{ data: []int{}, product: 0 }
}

func (this *ProductOfNumbers) Add(num int)  {
    this.data = append(this.data, num)
    if len(this.data) == 1 {
        this.product = num
    } else {
        this.product = this.product * num
    }
}

func (this *ProductOfNumbers) GetProduct(k int) int {
    n := len(this.data)
    if k >= n {
        return this.product
    }
    res := this.data[n - 1]
    for i := n - 2; i >= n - k; i--{
        res = res * this.data[i]
    }
    return res
}


type ProductOfNumbers1 struct {
    Product []int
    Data    []int
}

func Constructor1() ProductOfNumbers1 {
    return ProductOfNumbers1{ Product: []int{1}, Data: make([]int, 0) }
}

func (this *ProductOfNumbers1) Add(num int) {
    if num == 0 {
        this.Product = []int{1}
        return
    }
    n := len(this.Product)
    this.Product = append(this.Product, this.Product[n-1] * num)
}

func (this *ProductOfNumbers1) GetProduct(k int) int {
    if k >= len(this.Product) { // 说明中途遇到0了
        return 0
    }
    last := len(this.Product) - 1
    start := last - k
    return this.Product[last] / this.Product[start]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */

func main() {
    // ProductOfNumbers productOfNumbers = new ProductOfNumbers();
    obj := Constructor()
    fmt.Println(obj)
    // productOfNumbers.add(3);        // [3]
    obj.Add(3)
    fmt.Println(obj) // [3]
    // productOfNumbers.add(0);        // [3,0]
    obj.Add(0)
    fmt.Println(obj) // [3,0]
    // productOfNumbers.add(2);        // [3,0,2]
    obj.Add(2)
    fmt.Println(obj) // [3,0,2]
    // productOfNumbers.add(5);        // [3,0,2,5]
    obj.Add(5)
    fmt.Println(obj) // [3,0,2,5]
    // productOfNumbers.add(4);        // [3,0,2,5,4]
    obj.Add(4)
    fmt.Println(obj) // [3,0,2,5,4]
    // productOfNumbers.getProduct(2); // return 20. The product of the last 2 numbers is 5 * 4 = 20
    fmt.Println(obj.GetProduct(2)) // 20
    // productOfNumbers.getProduct(3); // return 40. The product of the last 3 numbers is 2 * 5 * 4 = 40
    fmt.Println(obj.GetProduct(3)) // 40
    // productOfNumbers.getProduct(4); // return 0. The product of the last 4 numbers is 0 * 2 * 5 * 4 = 0
    fmt.Println(obj.GetProduct(4)) // 0
    // productOfNumbers.add(8);        // [3,0,2,5,4,8]
    obj.Add(8)
    fmt.Println(obj) // [3,0,2,5,4,8]
    // productOfNumbers.getProduct(2); // return 32. The product of the last 2 numbers is 4 * 8 = 32 
    fmt.Println(obj.GetProduct(2)) // 32

    // ProductOfNumbers productOfNumbers = new ProductOfNumbers();
    obj1 := Constructor1()
    fmt.Println(obj1)
    // productOfNumbers.add(3);        // [3]
    obj1.Add(3)
    fmt.Println(obj1) // [3]
    // productOfNumbers.add(0);        // [3,0]
    obj1.Add(0)
    fmt.Println(obj1) // [3,0]
    // productOfNumbers.add(2);        // [3,0,2]
    obj1.Add(2)
    fmt.Println(obj1) // [3,0,2]
    // productOfNumbers.add(5);        // [3,0,2,5]
    obj1.Add(5)
    fmt.Println(obj1) // [3,0,2,5]
    // productOfNumbers.add(4);        // [3,0,2,5,4]
    obj1.Add(4)
    fmt.Println(obj1) // [3,0,2,5,4]
    // productOfNumbers.getProduct(2); // return 20. The product of the last 2 numbers is 5 * 4 = 20
    fmt.Println(obj1.GetProduct(2)) // 20
    // productOfNumbers.getProduct(3); // return 40. The product of the last 3 numbers is 2 * 5 * 4 = 40
    fmt.Println(obj1.GetProduct(3)) // 40
    // productOfNumbers.getProduct(4); // return 0. The product of the last 4 numbers is 0 * 2 * 5 * 4 = 0
    fmt.Println(obj1.GetProduct(4)) // 0
    // productOfNumbers.add(8);        // [3,0,2,5,4,8]
    obj1.Add(8)
    fmt.Println(obj1) // [3,0,2,5,4,8]
    // productOfNumbers.getProduct(2); // return 32. The product of the last 2 numbers is 4 * 8 = 32 
    fmt.Println(obj1.GetProduct(2)) // 32
}