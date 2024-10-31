package main

// 2353. Design a Food Rating System
// Design a food rating system that can do the following:
//     Modify the rating of a food item listed in the system.
//     Return the highest-rated food item for a type of cuisine in the system.

// Implement the FoodRatings class:
//     FoodRatings(String[] foods, String[] cuisines, int[] ratings) 
//         Initializes the system. 
//         The food items are described by foods, cuisines and ratings, all of which have a length of n.
//             foods[i] is the name of the ith food,
//             cuisines[i] is the type of cuisine of the ith food, and
//             ratings[i] is the initial rating of the ith food.
//     void changeRating(String food, int newRating) 
//         Changes the rating of the food item with the name food.
//     String highestRated(String cuisine) 
//         Returns the name of the food item that has the highest rating for the given type of cuisine. 
//         If there is a tie, return the item with the lexicographically smaller name.

// Note that a string x is lexicographically smaller than string y if x comes before y in dictionary order, 
// that is, either x is a prefix of y, or if i is the first position such that x[i] != y[i], then x[i] comes before y[i] in alphabetic order.

// Example 1:
// Input
// ["FoodRatings", "highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"]
// [[["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], ["korean", "japanese", "japanese", "greek", "japanese", "korean"], [9, 12, 8, 15, 14, 7]], ["korean"], ["japanese"], ["sushi", 16], ["japanese"], ["ramen", 16], ["japanese"]]
// Output
// [null, "kimchi", "ramen", null, "sushi", null, "ramen"]
// Explanation
// FoodRatings foodRatings = new FoodRatings(["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], ["korean", "japanese", "japanese", "greek", "japanese", "korean"], [9, 12, 8, 15, 14, 7]);
// foodRatings.highestRated("korean"); // return "kimchi"
//                                     // "kimchi" is the highest rated korean food with a rating of 9.
// foodRatings.highestRated("japanese"); // return "ramen"
//                                       // "ramen" is the highest rated japanese food with a rating of 14.
// foodRatings.changeRating("sushi", 16); // "sushi" now has a rating of 16.
// foodRatings.highestRated("japanese"); // return "sushi"
//                                       // "sushi" is the highest rated japanese food with a rating of 16.
// foodRatings.changeRating("ramen", 16); // "ramen" now has a rating of 16.
// foodRatings.highestRated("japanese"); // return "ramen"
//                                       // Both "sushi" and "ramen" have a rating of 16.
//                                       // However, "ramen" is lexicographically smaller than "sushi".

// Constraints:
//     1 <= n <= 2 * 10^4
//     n == foods.length == cuisines.length == ratings.length
//     1 <= foods[i].length, cuisines[i].length <= 10
//     foods[i], cuisines[i] consist of lowercase English letters.
//     1 <= ratings[i] <= 10^8
//     All the strings in foods are distinct.
//     food will be the name of a food item in the system across all calls to changeRating.
//     cuisine will be a type of cuisine of at least one food item in the system across all calls to highestRated.
//     At most 2 * 10^4 calls in total will be made to changeRating and highestRated.

import "fmt"
import "container/heap"

type FoodInfo struct {
    Name string // 名称
    Rating int  // 评分
}

type MaxHeap []FoodInfo

func (h MaxHeap) Peek() FoodInfo { return h[0] }
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Swap(i int, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i int, j int) bool {
    if h[i].Rating == h[j].Rating { return h[i].Name < h[j].Name }
    return h[i].Rating > h[j].Rating
}
func (h *MaxHeap) Push(v interface{}) {  *h = append(*h, v.(FoodInfo)) }
func (h *MaxHeap) Pop() interface{} {
    n := len(*h)
    res := (*h)[n - 1]
    *h = (*h)[:n - 1]
    return res
}

type FoodRatings struct {
    Country map[string]string // 记录食物和国家的关联
    Food map[string]int       // 记录食物和评分的关联
    Rating map[string]*MaxHeap // 食物评分排行
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
    res := FoodRatings{ make(map[string]string), make(map[string]int), make(map[string]*MaxHeap), }
    
    for i, v := range foods {
        res.Country[v] = cuisines[i] // Food => Country
        res.Food[v] = ratings[i]     // Food => Rating
        if _, ok := res.Rating[cuisines[i]]; !ok {
            res.Rating[cuisines[i]] = &MaxHeap{}
        }
        heap.Push(res.Rating[cuisines[i]], FoodInfo{ v, ratings[i] })
    }
    return res
}

func (this *FoodRatings) ChangeRating(food string, newRating int)  {
    this.Food[food] = newRating
    heap.Push(this.Rating[this.Country[food]], FoodInfo{ food, newRating })
}

func (this *FoodRatings) HighestRated(cuisine string) string {
    for {
        f := this.Rating[cuisine].Peek()
        if this.Food[f.Name] != f.Rating {
            heap.Pop(this.Rating[cuisine])
            continue
        } else {
            return f.Name
        }
    }
    return ""
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */

func main() {
    // FoodRatings foodRatings = new FoodRatings(["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], ["korean", "japanese", "japanese", "greek", "japanese", "korean"], [9, 12, 8, 15, 14, 7]);
    obj := Constructor(
        []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
        []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
        []int{9, 12, 8, 15, 14, 7},
    )
    fmt.Println(obj)
    // foodRatings.highestRated("korean"); // return "kimchi"
    //                                     // "kimchi" is the highest rated korean food with a rating of 9.
    fmt.Println(obj.HighestRated("korean")) // "kimchi"
    // foodRatings.highestRated("japanese"); // return "ramen"
    //                                       // "ramen" is the highest rated japanese food with a rating of 14.
    fmt.Println(obj.HighestRated("japanese")) // "ramen"
    // foodRatings.changeRating("sushi", 16); // "sushi" now has a rating of 16.
    obj.ChangeRating("sushi", 16)
    fmt.Println(obj)
    // foodRatings.highestRated("japanese"); // return "sushi"
    //                                       // "sushi" is the highest rated japanese food with a rating of 16.
    fmt.Println(obj.HighestRated("japanese")) // "ramen"
    // foodRatings.changeRating("ramen", 16); // "ramen" now has a rating of 16.
    obj.ChangeRating("ramen", 16)
    fmt.Println(obj)
    // foodRatings.highestRated("japanese"); // return "ramen"
    //                                       // Both "sushi" and "ramen" have a rating of 16.
    //                                       // However, "ramen" is lexicographically smaller than "sushi".
    fmt.Println(obj.HighestRated("japanese")) // "ramen"
}