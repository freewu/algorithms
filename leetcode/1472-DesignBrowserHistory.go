package main

// 1472. Design Browser History
// You have a browser of one tab where you start on the homepage and you can visit another url, get back in the history number of steps or move forward in the history number of steps.

// Implement the BrowserHistory class:
//     BrowserHistory(string homepage) Initializes the object with the homepage of the browser.
//     void visit(string url) Visits url from the current page. It clears up all the forward history.
//     string back(int steps) Move steps back in history. If you can only return x steps in the history and steps > x, you will return only x steps. Return the current url after moving back in history at most steps.
//     string forward(int steps) Move steps forward in history. If you can only forward x steps in the history and steps > x, you will forward only x steps. Return the current url after forwarding in history at most steps.
    

// Example:
// Input:
// ["BrowserHistory","visit","visit","visit","back","back","forward","visit","forward","back","back"]
// [["leetcode.com"],["google.com"],["facebook.com"],["youtube.com"],[1],[1],[1],["linkedin.com"],[2],[2],[7]]
// Output:
// [null,null,null,null,"facebook.com","google.com","facebook.com",null,"linkedin.com","google.com","leetcode.com"]
// Explanation:
// BrowserHistory browserHistory = new BrowserHistory("leetcode.com");
// browserHistory.visit("google.com");       // You are in "leetcode.com". Visit "google.com"
// browserHistory.visit("facebook.com");     // You are in "google.com". Visit "facebook.com"
// browserHistory.visit("youtube.com");      // You are in "facebook.com". Visit "youtube.com"
// browserHistory.back(1);                   // You are in "youtube.com", move back to "facebook.com" return "facebook.com"
// browserHistory.back(1);                   // You are in "facebook.com", move back to "google.com" return "google.com"
// browserHistory.forward(1);                // You are in "google.com", move forward to "facebook.com" return "facebook.com"
// browserHistory.visit("linkedin.com");     // You are in "facebook.com". Visit "linkedin.com"
// browserHistory.forward(2);                // You are in "linkedin.com", you cannot move forward any steps.
// browserHistory.back(2);                   // You are in "linkedin.com", move back two steps to "facebook.com" then to "google.com". return "google.com"
// browserHistory.back(7);                   // You are in "google.com", you can move back only one step to "leetcode.com". return "leetcode.com"
 
// Constraints:
//     1 <= homepage.length <= 20
//     1 <= url.length <= 20
//     1 <= steps <= 100
//     homepage and url consist of  '.' or lower case English letters.
//     At most 5000 calls will be made to visit, back, and forward.

import "fmt"

type BrowserHistory struct {
    index int // 当前的位置
    sites []string // 访问过的网站
}

func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{index: 0, sites: []string{ homepage }}
}

func (this *BrowserHistory) Visit(url string) {
    this.sites = append(this.sites[:this.index + 1], url)
    this.index = len(this.sites) - 1
}

func (this *BrowserHistory) Back(steps int) string {
    // 回退步子太大直接回到最开始 0
    if this.index - steps <= 0 {
        this.index = 0
        return this.sites[0]
    }
    this.index -= steps
    return this.sites[this.index]
}

func (this *BrowserHistory) Forward(steps int) string {
    // 向前的步子太大直接到最后
    if len(this.sites) <= this.index + steps {
        this.index = len(this.sites) - 1
        return this.sites[this.index]
    }
    this.index += steps
    return this.sites[this.index]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */

func main() {
    // BrowserHistory browserHistory = new BrowserHistory("leetcode.com");
    obj := Constructor("leetcode.com")
    fmt.Println(obj) // {0 [leetcode.com]}
    // browserHistory.visit("google.com");       // You are in "leetcode.com". Visit "google.com"
    obj.Visit("google.com")
    fmt.Println(obj) // {1 [leetcode.com google.com]}
    // browserHistory.visit("facebook.com");     // You are in "google.com". Visit "facebook.com"
    obj.Visit("facebook.com")
    fmt.Println(obj) // {2 [leetcode.com google.com facebook.com]}
    // browserHistory.visit("youtube.com");      // You are in "facebook.com". Visit "youtube.com"
    obj.Visit("youtube.com")
    fmt.Println(obj) // {3 [leetcode.com google.com facebook.com youtube.com]}
    // browserHistory.back(1);                   // You are in "youtube.com", move back to "facebook.com" return "facebook.com"
    fmt.Println(obj.Back(1)) // facebook.com
    // browserHistory.back(1);                   // You are in "facebook.com", move back to "google.com" return "google.com"
    fmt.Println(obj.Back(1)) // google.com
    // browserHistory.forward(1);                // You are in "google.com", move forward to "facebook.com" return "facebook.com"
    fmt.Println(obj.Forward(1)) // facebook.com
    // browserHistory.visit("linkedin.com");     // You are in "facebook.com". Visit "linkedin.com"
    obj.Visit("linkedin.com")
    fmt.Println(obj) // {3 [leetcode.com google.com facebook.com linkedin.com]}
    // browserHistory.forward(2);                // You are in "linkedin.com", you cannot move forward any steps.
    fmt.Println(obj.Forward(2)) // linkedin.com
    // browserHistory.back(2);                   // You are in "linkedin.com", move back two steps to "facebook.com" then to "google.com". return "google.com"
    fmt.Println(obj.Back(2)) // google.com
    // browserHistory.back(7);                   // You are in "google.com", you can move back only one step to "leetcode.com". return "leetcode.com"
    fmt.Println(obj.Back(7)) // leetcode.com
}