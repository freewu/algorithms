// 2777. Date Range Generator
// Given a start date start, an end date end, and a positive integer step, 
// return a generator object that yields dates in the range from start to end inclusive. 
// All dates are in the string format YYYY-MM-DD. 
// The value of step indicates the number of days between consecutive yielded values.

// Example 1:
// Input: start = "2023-04-01", end = "2023-04-04", step = 1
// Output: ["2023-04-01","2023-04-02","2023-04-03","2023-04-04"]
// Explanation: 
// const g = dateRangeGenerator(start, end, step);
// g.next().value // '2023-04-01'
// g.next().value // '2023-04-02'
// g.next().value // '2023-04-03'
// g.next().value // '2023-04-04'

// Example 2:
// Input: start = "2023-04-10", end = "2023-04-20", step = 3
// Output: ["2023-04-10","2023-04-13","2023-04-16","2023-04-19"]
// Explanation: 
// const g = dateRangeGenerator(start, end, step);
// g.next().value // '2023-04-10'
// g.next().value // '2023-04-13'
// g.next().value // '2023-04-16'
// g.next().value // '2023-04-19'
// Example 3:

// Input: start = "2023-04-10", end = "2023-04-10", step = 1
// Output: ["2023-04-10"]
// Explanation: 
// const g = dateRangeGenerator(start, end, step);
// g.next().value // '2023-04-10'
 
// Constraints:
//         new Date(start) <= new Date(end)
//         0 <= The difference in days between the start date and the end date <= 1000
//         1 <= step <= 100

var dateRangeGenerator = function* (start, end, step) {
    const s = new Date(start), e = new Date(end)
    const a = 24 * 60 * 60 * 1000 * step
    while (s <= e) {
        // All dates are in the string format YYYY-MM-DD. 
        // 返回 YYYY-MM-DD 格式
        yield s.toISOString().split("T")[0]
        // 加上一个 step 的时间
        s.setTime(s.getTime() + a)
    }
};

var dateRangeGenerator1 = function* (start, end, step) {
	let s = new Date(start)
    const e = new Date(end)

    while(s.getTime() <= e.getTime()) {
        yield s.toISOString().slice(0, 10)
        s = new Date(s.getTime() + 24 * 60 * 60 * 1000 * step)
    }
};

/**
 * const g = dateRangeGenerator('2023-04-01', '2023-04-04', 1);
 * g.next().value; // '2023-04-01'
 * g.next().value; // '2023-04-02'
 * g.next().value; // '2023-04-03'
 * g.next().value; // '2023-04-04'
 * g.next().done; // true
 */

// Example 1:
let g = dateRangeGenerator('2023-04-01', '2023-04-04', 1);
console.log(g.next().value); // '2023-04-01'
console.log(g.next().value); // '2023-04-02'
console.log(g.next().value); // '2023-04-03'
console.log(g.next().value); // '2023-04-04'
console.log(g.next().value); // undefined
console.log(g.next().done); // true

// Example 2:
g = dateRangeGenerator("2023-04-10",  "2023-04-20", 3);
console.log(g.next().value); // '2023-04-10'
console.log(g.next().value); // '2023-04-13'
console.log(g.next().value); // '2023-04-16'
console.log(g.next().value); // '2023-04-19'
console.log(g.next().done); // true

// Example 3:
g = dateRangeGenerator("2023-04-10",  "2023-04-10", 1);
console.log(g.next().value); // '2023-04-10'
console.log(g.next().done); // true

// Example 1:
g = dateRangeGenerator1('2023-04-01', '2023-04-04', 1);
console.log(g.next().value); // '2023-04-01'
console.log(g.next().value); // '2023-04-02'
console.log(g.next().value); // '2023-04-03'
console.log(g.next().value); // '2023-04-04'
console.log(g.next().value); // undefined
console.log(g.next().done); // true

// Example 2:
g = dateRangeGenerator1("2023-04-10",  "2023-04-20", 3);
console.log(g.next().value); // '2023-04-10'
console.log(g.next().value); // '2023-04-13'
console.log(g.next().value); // '2023-04-16'
console.log(g.next().value); // '2023-04-19'
console.log(g.next().done); // true

// Example 3:
g = dateRangeGenerator1("2023-04-10",  "2023-04-10", 1);
console.log(g.next().value); // '2023-04-10'
console.log(g.next().done); // true