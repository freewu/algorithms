-- 1709. Biggest Window Between Visits
-- Table: UserVisits
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | user_id     | int  |
-- | visit_date  | date |
-- +-------------+------+
-- This table does not have a primary key, it might contain duplicate rows.
-- This table contains logs of the dates that users visited a certain retailer.

-- Assume today's date is '2021-1-1'.
-- Write a solution that will, for each user_id, find out the largest window of days between each visit and the one right after it (or today if you are considering the last visit).
-- Return the result table ordered by user_id.
-- The query result format is in the following example.
 
-- Example 1:
-- Input: 
-- UserVisits table:
-- +---------+------------+
-- | user_id | visit_date |
-- +---------+------------+
-- | 1       | 2020-11-28 |
-- | 1       | 2020-10-20 |
-- | 1       | 2020-12-3  |
-- | 2       | 2020-10-5  |
-- | 2       | 2020-12-9  |
-- | 3       | 2020-11-11 |
-- +---------+------------+
-- Output: 
-- +---------+---------------+
-- | user_id | biggest_window|
-- +---------+---------------+
-- | 1       | 39            |
-- | 2       | 65            |
-- | 3       | 51            |
-- +---------+---------------+
-- Explanation: 
-- For the first user, the windows in question are between dates:
--     - 2020-10-20 and 2020-11-28 with a total of 39 days. 
--     - 2020-11-28 and 2020-12-3 with a total of 5 days. 
--     - 2020-12-3 and 2021-1-1 with a total of 29 days.
-- Making the biggest window the one with 39 days.
-- For the second user, the windows in question are between dates:
--     - 2020-10-5 and 2020-12-9 with a total of 65 days.
--     - 2020-12-9 and 2021-1-1 with a total of 23 days.
-- Making the biggest window the one with 65 days.
-- For the third user, the only window in question is between dates 2020-11-11 and 2021-1-1 with a total of 51 days.

-- Create table If Not Exists UserVisits(user_id int, visit_date date)
-- Truncate table UserVisits
-- insert into UserVisits (user_id, visit_date) values ('1', '2020-11-28')
-- insert into UserVisits (user_id, visit_date) values ('1', '2020-10-20')
-- insert into UserVisits (user_id, visit_date) values ('1', '2020-12-3')
-- insert into UserVisits (user_id, visit_date) values ('2', '2020-10-5')
-- insert into UserVisits (user_id, visit_date) values ('2', '2020-12-9')
-- insert into UserVisits (user_id, visit_date) values ('3', '2020-11-11')

-- lead over
SELECT 
    user_id, 
    MAX(windows) AS biggest_window
FROM 
(
    SELECT 
        user_id, 
        DATEDIFF(LEAD(visit_date, 1, "2021-01-01") OVER (PARTITION BY user_id ORDER BY visit_date ASC), visit_date) AS windows
    FROM 
        UserVisits
) AS uv
GROUP BY 
    user_id;

-- SELECT 
--     user_id, 
--     DATEDIFF(LEAD(visit_date, 1, "2021-01-01") OVER (PARTITION BY user_id ORDER BY visit_date ASC), visit_date) AS windows
-- FROM 
--     UserVisits
-- | user_id | windows |
-- | ------- | ------- |
-- | 1       | 39      |
-- | 1       | 5       |
-- | 1       | 29      |
-- | 2       | 65      |
-- | 2       | 23      |
-- | 3       | 51      |

-- SELECT 
--     user_id,
--     visit_date, 
--     -- visit_date asc 排序   
--     -- LEAD(visit_date, 1, "2021-01-01") 取 visit_date 的下一个日期,如果没有则使用 "2021-01-01"
--     LEAD(visit_date, 1, "2021-01-01") OVER (PARTITION BY user_id ORDER BY visit_date ASC) AS a 
-- FROM
--     UserVisits
-- | user_id | visit_date | a          |
-- | ------- | ---------- | ---------- |
-- | 1       | 2020-10-20 | 2020-11-28 |
-- | 1       | 2020-11-28 | 2020-12-03 |
-- | 1       | 2020-12-03 | 2021-01-01 |
-- | 2       | 2020-10-05 | 2020-12-09 |
-- | 2       | 2020-12-09 | 2021-01-01 |
-- | 3       | 2020-11-11 | 2021-01-01 |

-- LEAD函数的语法如下：
--     LEAD(expr [, offset [, default_value]]) OVER ([PARTITION BY partition_expression] ORDER BY order_expression [ASC|DESC])
-- SQL
-- LEAD函数接受三个参数：
--     expr：必选参数，表示要查询的列或表达式。
--     offset：可选参数，默认为1，表示要查询的下一行的偏移量。该值可以为正整数或负整数。
--     default_value：可选参数，表示在没有下一行时返回的默认值。

-- LEAD函数还接受两个子句：
--     PARTITION BY：可选子句，表示按指定的表达式进行分组。如果不指定该子句，则对整个结果集进行计算。
--     ORDER BY：必选子句，表示按指定的表达式进行排序。排序方式可以是升序（ASC）或降序（DESC）

