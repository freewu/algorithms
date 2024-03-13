-- 1445. Apples & Oranges
-- Table: Sales
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | sale_date     | date    |
-- | fruit         | enum    | 
-- | sold_num      | int     | 
-- +---------------+---------+
-- (sale_date, fruit) is the primary key (combination of columns with unique values) of this table.
-- This table contains the sales of "apples" and "oranges" sold each day.

-- Write a solution to report the difference between the number of apples and oranges sold each day.
-- Return the result table ordered by sale_date.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Sales table:
-- +------------+------------+-------------+
-- | sale_date  | fruit      | sold_num    |
-- +------------+------------+-------------+
-- | 2020-05-01 | apples     | 10          |
-- | 2020-05-01 | oranges    | 8           |
-- | 2020-05-02 | apples     | 15          |
-- | 2020-05-02 | oranges    | 15          |
-- | 2020-05-03 | apples     | 20          |
-- | 2020-05-03 | oranges    | 0           |
-- | 2020-05-04 | apples     | 15          |
-- | 2020-05-04 | oranges    | 16          |
-- +------------+------------+-------------+
-- Output: 
-- +------------+--------------+
-- | sale_date  | diff         |
-- +------------+--------------+
-- | 2020-05-01 | 2            |
-- | 2020-05-02 | 0            |
-- | 2020-05-03 | 20           |
-- | 2020-05-04 | -1           |
-- +------------+--------------+
-- Explanation: 
-- Day 2020-05-01, 10 apples and 8 oranges were sold (Difference  10 - 8 = 2).
-- Day 2020-05-02, 15 apples and 15 oranges were sold (Difference 15 - 15 = 0).
-- Day 2020-05-03, 20 apples and 0 oranges were sold (Difference 20 - 0 = 20).
-- Day 2020-05-04, 15 apples and 16 oranges were sold (Difference 15 - 16 = -1).

-- Create table If Not Exists Sales (sale_date date, fruit ENUM('apples', 'oranges'), sold_num int)
-- Truncate table Sales
-- insert into Sales (sale_date, fruit, sold_num) values ('2020-05-01', 'apples', '10')
-- insert into Sales (sale_date, fruit, sold_num) values ('2020-05-01', 'oranges', '8')
-- insert into Sales (sale_date, fruit, sold_num) values ('2020-05-02', 'apples', '15')
-- insert into Sales (sale_date, fruit, sold_num) values ('2020-05-02', 'oranges', '15')
-- insert into Sales (sale_date, fruit, sold_num) values ('2020-05-03', 'apples', '20')
-- insert into Sales (sale_date, fruit, sold_num) values ('2020-05-03', 'oranges', '0')
-- insert into Sales (sale_date, fruit, sold_num) values ('2020-05-04', 'apples', '15')
-- insert into Sales (sale_date, fruit, sold_num) values ('2020-05-04', 'oranges', '16')

# Write your MySQL query statement below
SELECT 
    sale_date,
    (apples - oranges) AS diff
FROM
(
    SELECT
        sale_date,
        SUM(IF(fruit = 'apples',sold_num, 0)) AS apples,
        SUM(IF(fruit = 'oranges',sold_num, 0)) AS oranges
    FROM
        Sales
    GROUP BY      
        sale_date
) AS t
ORDER BY
    sale_date -- 结果表, 按照格式为 ('YYYY-MM-DD') 的 sale_date 排序

-- best solution
SELECT 
    sale_date, 
    SUM(IF( fruit='apples', sold_num, -sold_num)) AS diff -- 不为 apples 就是 apples 要减掉的数
FROM
    Sales
GROUP BY      
    sale_date
ORDER BY
    sale_date