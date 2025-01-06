-- 3415. Find Products with Three Consecutive Digits 
-- Table: Products
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | product_id  | int     |
-- | name        | varchar |
-- +-------------+---------+
-- product_id is the unique key for this table.
-- Each row of this table contains the ID and name of a product.
-- Write a solution to find all products whose names contain a sequence of exactly three digits in a row. 

-- Return the result table ordered by product_id in ascending order.

-- The result format is in the following example.

-- Example:
-- Input:
-- products table:
-- +-------------+--------------------+
-- | product_id  | name               |
-- +-------------+--------------------+
-- | 1           | ABC123XYZ          |
-- | 2           | A12B34C            |
-- | 3           | Product56789       |
-- | 4           | NoDigitsHere       |
-- | 5           | 789Product         |
-- | 6           | Item003Description |
-- | 7           | Product12X34       |
-- +-------------+--------------------+
-- Output:
-- +-------------+--------------------+
-- | product_id  | name               |
-- +-------------+--------------------+
-- | 1           | ABC123XYZ          |
-- | 5           | 789Product         |
-- | 6           | Item003Description |
-- +-------------+--------------------+
-- Explanation:
-- Product 1: ABC123XYZ contains the digits 123.
-- Product 5: 789Product contains the digits 789.
-- Product 6: Item003Description contains 003, which is exactly three digits.
-- Note:
-- Results are ordered by product_id in ascending order.
-- Only products with exactly three consecutive digits in their names are included in the result.

-- CREATE TABLE if not exists products  (
--     product_id INT,
--     name VARCHAR(255)
-- )

-- Truncate table Products
-- insert into Products (product_id, name) values ('1', 'ABC123XYZ')
-- insert into Products (product_id, name) values ('2', 'A12B34C')
-- insert into Products (product_id, name) values ('3', 'Product56789')
-- insert into Products (product_id, name) values ('4', 'NoDigitsHere')
-- insert into Products (product_id, name) values ('5', '789Product')
-- insert into Products (product_id, name) values ('6', 'Item003Description')
-- insert into Products (product_id, name) values ('7', 'Product12X34')

-- RLIKE
-- 语法格式为A [NOT] RLIKE B，基于java的正则表达式接口实现，如果A中有与B匹配则返回TRUE，否则返回FALSE
-- ^：用来匹配字符串的开始
-- $：用来匹配字符串的结尾。
-- []：方括号中的任何字符都可以匹配，例如[0-9a-z其他]可以匹配‘0-9’数字任意一个，小写字母‘a-z’任意一个，‘其’，‘他’。
-- -：连接符用来表示字符串的范围，如上面的[0-9]
-- +：表示匹配次数出现一个或多个。
-- *：表示匹配该字符出现0个或多个，比如[0-9]*表示匹配0个或多个数字。
-- ()：在圆括号中的内容将被看做一个整体。（ab）匹配ab。
-- {m}：整数m表示花括号前的字符串出现的次数。比如(ab){2}匹配abab，(ab){1,3}表示匹配1到3个‘ab’。
-- |：表示或‘or’，比如’^张|^李’表示匹配张开头，或者李开头。

-- NOT REGEXP	REGEXP的否定
-- REGEXP	string是否匹配正则表达式
-- REGEXP_INSTR()	匹配正则表达式的子串的起始索引
-- REGEXP_LIKE()	string是否匹配正则表达式
-- REGEXP_REPLACE()	替换匹配正则表达式的子字符串
-- REGEXP_SUBSTR()	返回子串匹配正则表达式
-- RLIKE	string是否匹配正则表达式

SELECT 
    *
FROM
    Products
WHERE
    name REGEXP '[0-9]{3}' And
    name NOT REGEXP '[0-9]{4}'
ORDER BY 
    product_id -- Return the result table ordered by product_id in ascending order.


SELECT 
    *
FROM
    Products
WHERE
    name RLIKE '[0-9]{3}' And
    name NOT RLIKE '[0-9]{4}'
ORDER BY 
    product_id -- Return the result table ordered by product_id in ascending order.