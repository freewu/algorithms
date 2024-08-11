package main

// 1357. Apply Discount Every n Orders
// There is a supermarket that is frequented by many customers. 
// The products sold at the supermarket are represented as two parallel integer arrays products and prices, 
// where the ith product has an ID of products[i] and a price of prices[i].

// When a customer is paying, their bill is represented as two parallel integer arrays product and amount, 
// where the jth product they purchased has an ID of product[j], and amount[j] is how much of the product they bought. 
// Their subtotal is calculated as the sum of each amount[j] * (price of the jth product).

// The supermarket decided to have a sale. 
// Every nth customer paying for their groceries will be given a percentage discount. 
// The discount amount is given by discount, where they will be given discount percent off their subtotal. 
// More formally, if their subtotal is bill, then they would actually pay bill * ((100 - discount) / 100).

// Implement the Cashier class:
//     Cashier(int n, int discount, int[] products, int[] prices) 
//         Initializes the object with n, the discount, and the products and their prices.
//     double getBill(int[] product, int[] amount) 
//         Returns the final total of the bill with the discount applied (if any). 
//         Answers within 10^-5 of the actual value will be accepted.

// Example 1:
// Input
// ["Cashier","getBill","getBill","getBill","getBill","getBill","getBill","getBill"]
// [[3,50,[1,2,3,4,5,6,7],[100,200,300,400,300,200,100]],[[1,2],[1,2]],[[3,7],[10,10]],[[1,2,3,4,5,6,7],[1,1,1,1,1,1,1]],[[4],[10]],[[7,3],[10,10]],[[7,5,3,1,6,4,2],[10,10,10,9,9,9,7]],[[2,3,5],[5,3,2]]]
// Output
// [null,500.0,4000.0,800.0,4000.0,4000.0,7350.0,2500.0]
// Explanation
// Cashier cashier = new Cashier(3,50,[1,2,3,4,5,6,7],[100,200,300,400,300,200,100]);
// cashier.getBill([1,2],[1,2]);                        // return 500.0. 1st customer, no discount.
//                                                      // bill = 1 * 100 + 2 * 200 = 500.
// cashier.getBill([3,7],[10,10]);                      // return 4000.0. 2nd customer, no discount.
//                                                      // bill = 10 * 300 + 10 * 100 = 4000.
// cashier.getBill([1,2,3,4,5,6,7],[1,1,1,1,1,1,1]);    // return 800.0. 3rd customer, 50% discount.
//                                                      // Original bill = 1600
//                                                      // Actual bill = 1600 * ((100 - 50) / 100) = 800.
// cashier.getBill([4],[10]);                           // return 4000.0. 4th customer, no discount.
// cashier.getBill([7,3],[10,10]);                      // return 4000.0. 5th customer, no discount.
// cashier.getBill([7,5,3,1,6,4,2],[10,10,10,9,9,9,7]); // return 7350.0. 6th customer, 50% discount.
//                                                      // Original bill = 14700, but with
//                                                      // Actual bill = 14700 * ((100 - 50) / 100) = 7350.
// cashier.getBill([2,3,5],[5,3,2]);                    // return 2500.0.  7th customer, no discount.

// Constraints:
//     1 <= n <= 10^4
//     0 <= discount <= 100
//     1 <= products.length <= 200
//     prices.length == products.length
//     1 <= products[i] <= 200
//     1 <= prices[i] <= 1000
//     The elements in products are unique.
//     1 <= product.length <= products.length
//     amount.length == product.length
//     product[j] exists in products.
//     1 <= amount[j] <= 1000
//     The elements of product are unique.
//     At most 1000 calls will be made to getBill.
//     Answers within 10^-5 of the actual value will be accepted.

import "fmt"

type Cashier struct {
    LuckyNumber     int // 幸运数
    CustomerCount   int // 客户数
    Discount        int // 折扣
    ProductPrice    map[int] int // [产品编号][价格]
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
    record := map[int] int {}
    for i := 0; i < len(products); i++ { 
        record[products[i]] = prices[i]
    }
    return Cashier {
        LuckyNumber: n,
        CustomerCount: 0,
        Discount: discount,
        ProductPrice: record,
    }
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
    this.CustomerCount++
    sum := 0
    for i:=0; i<len(product); i++ {
        sum += this.ProductPrice[product[i]] * amount[i]
    }
    if this.CustomerCount == this.LuckyNumber { // 幸运用户 计算折扣
        this.CustomerCount = 0
        return float64(sum) - (float64(sum) * float64(this.Discount)) / 100
    }
    return float64(sum)
}

type Cashier1 struct {
    LuckyNumber     int // 幸运数
    CustomerCount   int // 客户数
    Discount        int // 折扣
    ProductPrice    []int // <-.商品价格表: ↑ 
}

func Constructor1(n int, discount int, products []int, prices []int) Cashier1 {
    mn := products[0]
    for _, p := range products[1:] {
        if mn < p { mn = p }
    }
    table := make([]int, 1 + mn) 
    for i, pro := range products {
        table[pro] = prices[i]  // <-.商品ID映射到单品价格: ↑. 
    }
    return Cashier1{n, 0, discount, table }
}

func (this *Cashier1) GetBill(product []int, amount []int) float64 {
    this.CustomerCount++
    bill := float64(0) 
    for i, v := range product {
        bill += float64(amount[i]) * float64(this.ProductPrice[v]) 
    }
    if this.CustomerCount % this.LuckyNumber == 0 {
        bill -= bill * (float64(this.Discount) / 100.0) 
    }
    return bill 
}

/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
 */

func main() {
    // Cashier cashier = new Cashier(3,50,[1,2,3,4,5,6,7],[100,200,300,400,300,200,100]);
    obj := Constructor(3, 50, []int{1,2,3,4,5,6,7}, []int{100,200,300,400,300,200,100})
    // cashier.getBill([1,2],[1,2]);                        // return 500.0. 1st customer, no discount.
    //                                                      // bill = 1 * 100 + 2 * 200 = 500.
    fmt.Println(obj.GetBill([]int{1, 2}, []int{1, 2})) // 500.0
    // cashier.getBill([3,7],[10,10]);                      // return 4000.0. 2nd customer, no discount.
    //                                                      // bill = 10 * 300 + 10 * 100 = 4000.
    fmt.Println(obj.GetBill([]int{3, 7}, []int{10, 10})) // 4000.0
    // cashier.getBill([1,2,3,4,5,6,7],[1,1,1,1,1,1,1]);    // return 800.0. 3rd customer, 50% discount.
    //                                                      // Original bill = 1600
    //                                                      // Actual bill = 1600 * ((100 - 50) / 100) = 800.
    fmt.Println(obj.GetBill([]int{1,2,3,4,5,6,7}, []int{1,1,1,1,1,1,1})) // 800.0
    // cashier.getBill([4],[10]);                           // return 4000.0. 4th customer, no discount.
    fmt.Println(obj.GetBill([]int{4}, []int{10})) // 4000.0
    // cashier.getBill([7,3],[10,10]);                      // return 4000.0. 5th customer, no discount.
    fmt.Println(obj.GetBill([]int{3, 7}, []int{10, 10})) // 4000.0
    // cashier.getBill([7,5,3,1,6,4,2],[10,10,10,9,9,9,7]); // return 7350.0. 6th customer, 50% discount.
    //                                                      // Original bill = 14700, but with
    //                                                      // Actual bill = 14700 * ((100 - 50) / 100) = 7350.
    fmt.Println(obj.GetBill([]int{7,5,3,1,6,4,2}, []int{10,10,10,9,9,9,7})) // 7350.0
    // cashier.getBill([2,3,5],[5,3,2]);                    // return 2500.0.  7th customer, no discount.
    fmt.Println(obj.GetBill([]int{2,3,5}, []int{5,3,2})) // 2500.0

    // Cashier cashier = new Cashier(3,50,[1,2,3,4,5,6,7],[100,200,300,400,300,200,100]);
    obj1 := Constructor1(3, 50, []int{1,2,3,4,5,6,7}, []int{100,200,300,400,300,200,100})
    // cashier.getBill([1,2],[1,2]);                        // return 500.0. 1st customer, no discount.
    //                                                      // bill = 1 * 100 + 2 * 200 = 500.
    fmt.Println(obj1.GetBill([]int{1, 2}, []int{1, 2})) // 500.0
    // cashier.getBill([3,7],[10,10]);                      // return 4000.0. 2nd customer, no discount.
    //                                                      // bill = 10 * 300 + 10 * 100 = 4000.
    fmt.Println(obj1.GetBill([]int{3, 7}, []int{10, 10})) // 4000.0
    // cashier.getBill([1,2,3,4,5,6,7],[1,1,1,1,1,1,1]);    // return 800.0. 3rd customer, 50% discount.
    //                                                      // Original bill = 1600
    //                                                      // Actual bill = 1600 * ((100 - 50) / 100) = 800.
    fmt.Println(obj1.GetBill([]int{1,2,3,4,5,6,7}, []int{1,1,1,1,1,1,1})) // 800.0
    // cashier.getBill([4],[10]);                           // return 4000.0. 4th customer, no discount.
    fmt.Println(obj1.GetBill([]int{4}, []int{10})) // 4000.0
    // cashier.getBill([7,3],[10,10]);                      // return 4000.0. 5th customer, no discount.
    fmt.Println(obj1.GetBill([]int{3, 7}, []int{10, 10})) // 4000.0
    // cashier.getBill([7,5,3,1,6,4,2],[10,10,10,9,9,9,7]); // return 7350.0. 6th customer, 50% discount.
    //                                                      // Original bill = 14700, but with
    //                                                      // Actual bill = 14700 * ((100 - 50) / 100) = 7350.
    fmt.Println(obj1.GetBill([]int{7,5,3,1,6,4,2}, []int{10,10,10,9,9,9,7})) // 7350.0
    // cashier.getBill([2,3,5],[5,3,2]);                    // return 2500.0.  7th customer, no discount.
    fmt.Println(obj1.GetBill([]int{2,3,5}, []int{5,3,2})) // 2500.0
}