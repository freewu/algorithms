// 2758. Next Day
// Write code that enhances all date objects such that you can call the date.nextDay() method on any date object 
// and it will return the next day in the format YYYY-MM-DD as a string.

// Example 1:
// Input: date = "2014-06-20"
// Output: "2014-06-21"
// Explanation: 
// const date = new Date("2014-06-20");
// date.nextDay(); // "2014-06-21"

// Example 2:
// Input: date = "2017-10-31"
// Output: "2017-11-01"
// Explanation: The day after 2017-10-31 is 2017-11-01.

// Constraints:
//         new Date(date) is a valid date object

/** 
 * @return {string}
 */
Date.prototype.nextDay = function() {
    let t = new Date(this.getTime() + 24 * 60 * 60 * 1000 * 1); // 加一天时间
    return t.toISOString().slice(0, 10); // 按 YYYY-MM-DD 格式输出
}

/**
 * const date = new Date("2014-06-20");
 * date.nextDay(); // "2014-06-21"
 */

console.log((new Date("2014-06-20")).nextDay()) // "2014-06-21"
console.log((new Date("2017-10-31")).nextDay()) // "2017-11-01"