-- 1613. Find the Missing IDs
-- Table: Customers
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | customer_id   | int     |
-- | customer_name | varchar |
-- +---------------+---------+
-- customer_id is the column with unique values for this table.
-- Each row of this table contains the name and the id customer.
 
-- Write a solution to find the missing customer IDs. 
-- The missing IDs are ones that are not in the Customers table but are in the range between 1 and the maximum customer_id present in the table.

-- Notice that the maximum customer_id will not exceed 100.

-- Return the result table ordered by ids in ascending order.
-- The result format is in the following example.
-- Example 1:

-- Input: 
-- Customers table:
-- +-------------+---------------+
-- | customer_id | customer_name |
-- +-------------+---------------+
-- | 1           | Alice         |
-- | 4           | Bob           |
-- | 5           | Charlie       |
-- +-------------+---------------+
-- Output: 
-- +-----+
-- | ids |
-- +-----+
-- | 2   |
-- | 3   |
-- +-----+
-- Explanation: 
-- The maximum customer_id present in the table is 5, so in the range [1,5], IDs 2 and 3 are missing from the table.

-- Create table If Not Exists Customers (customer_id int, customer_name varchar(20))
-- Truncate table Customers
-- insert into Customers (customer_id, customer_name) values ('1', 'Alice')
-- insert into Customers (customer_id, customer_name) values ('4', 'Bob')
-- insert into Customers (customer_id, customer_name) values ('5', 'Charlie')

-- 生成 1 - MAX(customer_id) 的I
-- SELECT 
--     (0 + (n+1)) AS ids
-- FROM 
-- (
--     SELECT 
--         @rownum:=@rownum+1 as n
--     FROM (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5) t1,
--          (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5) t2,
--          (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5) t3,
--          (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5) t5,
--          (SELECT @rownum:=-1) t4
-- ) AS numbers
-- WHERE 
--     (0 + (n+1)) <= (SELECT MAX(customer_id) FROM Customers);

SELECT
    i.ids AS ids
FROM
    (
        SELECT 
            (0 + (n+1)) AS ids
        FROM 
        (
            SELECT 
                @rownum:=@rownum+1 as n
            FROM (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5) t1,
                (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5) t2,
                (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5) t3,
                (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5) t5,
                (SELECT @rownum:=-1) t4
        ) AS numbers
        WHERE 
            (0 + (n+1)) <= (SELECT MAX(customer_id) FROM Customers)
    ) AS i
LEFT JOIN
    Customers AS c
ON
    i.ids = c.customer_id 
WHERE
    c.customer_name IS NULL
ORDER BY
    ids

-- 除了使用递归得到1-100的自然数,也可以使用下面的方式
with t1 as (
  select 1 as number
  union all
  select 2
  union all
  select 3
  union all
  select 4
  union all
  select 5
  union all
  select 6
  union all
  select 7
  union all
  select 8
  union all
  select 9
),
t2 as ( -- 创造1-100的自然数
  select CAST(CONCAT(cast(a.number AS CHAR),cast(b.number AS CHAR)) AS SIGNED) as num
  from t1 as a, t1 as b
  union all
  select 10 as num
  union all
  select 20 as num
  union all
  select 30 as num
  union all
  select 40 as num
  union all
  select 50 as num
  union all
  select 60 as num
  union all
  select 70 as num
  union all
  select 80 as num
  union all
  select 90 as num
  union all
  select 100 as num
  union all
  select * from t1
)
select 
    num as ids
from
     t2
where 
    num not in (select customer_id from Customers) and 
    num < (select max(customer_id) from Customers)
order by 
    num

-- 通过递归生成ID
WITH RECURSIVE table1(ids) AS 
(
    SELECT 1
    UNION ALL
    SELECT 
        ids + 1 
    FROM 
        table1 
    WHERE 
        ids < (SELECT MAX(customer_id) FROM customers)
)

SELECT 
    ids
FROM 
    table1
WHERE 
    ids NOT IN ( SELECT customer_id FROM customers )
ORDER BY 
    ids