// 2705. Compact Object
// Given an object or array obj, return a compact object. 
// A compact object is the same as the original object, except with keys containing falsy values removed. 
// This operation applies to the object and any nested objects. 
// Arrays are considered objects where the indices are keys. 
// A value is considered falsy when Boolean(value) returns false.
// You may assume the obj is the output of JSON.parse. In other words, it is valid JSON.

// Example 1:
// Input: obj = [null, 0, false, 1]
// Output: [1]
// Explanation: All falsy values have been removed from the array.

// Example 2:
// Input: obj = {"a": null, "b": [false, 1]}
// Output: {"b": [1]}
// Explanation: obj["a"] and obj["b"][0] had falsy values and were removed.

// Example 3:
// Input: obj = [null, 0, 5, [0], [false, 16]]
// Output: [5, [], [16]]
// Explanation: obj[0], obj[1], obj[3][0], and obj[4][0] were falsy and removed.

// Constraints:
//         obj is a valid JSON object
//         2 <= JSON.stringify(obj).length <= 10^6

/**
 * @param {Object|Array} obj
 * @return {Object|Array}
 */
var compactObject = function(obj) {
    // 不是对象就可能是null或者字符，数字（因为题目说是JSON转化，排除函数和奇怪的东西）
    if (obj == null || typeof obj !== 'object') {
        return obj;
    }
    // 数组的话可以直接枚举
    if (Array.isArray(obj)) {
        const res = [];
        for (let it of obj) {
            const val = compactObject(it);
            // 不为 0 / false / null / undefined 重新入到新的数组中 
            if (val) res.push(val);
        }
        return res;
    }
    // 对象需要把key拿出来
    const res = {};
    const keys = Object.keys(obj);
    for (let key of keys) {
        const val = compactObject(obj[key]);
        // 值 不为 0 / false / null / undefined 重新入到新的对象中 
        if (val) res[key] = val;
    }
    return res;
};

console.log(compactObject([null, 0, false, 1])) // [1]
console.log(compactObject({"a": null, "b": [false, 1]})) // {"b": [1]}
console.log(compactObject([null, 0, 5, [0], [false, 16]])) // [5, [], [16]]
