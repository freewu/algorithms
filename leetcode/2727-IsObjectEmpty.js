// 2727. Is Object Empty
// Given an object or an array, return if it is empty.
//     An empty object contains no key-value pairs.
//     An empty array contains no elements.

// You may assume the object or array is the output of JSON.parse.

// Example 1:
// Input: obj = {"x": 5, "y": 42}
// Output: false
// Explanation: The object has 2 key-value pairs so it is not empty.

// Example 2:
// Input: obj = {}
// Output: true
// Explanation: The object doesn't have any key-value pairs so it is empty.

// Example 3:
// Input: obj = [null, false, 0]
// Output: false
// Explanation: The array has 3 elements so it is not empty.

// Constraints:
//     obj is a valid JSON object or array
//     2 <= JSON.stringify(obj).length <= 10^5

/**
 * @param {Object|Array} obj
 * @return {boolean}
 */
var isEmpty = function(obj) {
    // // 如果是数组判断，数组长度
    // if (Array.isArray(obj)) return obj.length === 0;
    // // 如果是对象，判读对象的 keys 数组的长度
    // return Object.keys(obj).length === 0;

    // 判断是否对象是否能循环
    for (let key in obj) return false;
    return true;

    // 判断是否是 [] 数组 {} 对象
    // return JSON.stringify(obj).length<=2?true:false;
};