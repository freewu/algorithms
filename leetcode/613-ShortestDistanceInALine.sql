-- 613. Shortest Distance in a Line
-- Table: Point
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | x           | int  |
-- +-------------+------+
-- x is the primary key column for this table.
-- Each row of this table indicates the position of a point on the X-axis.
--  
-- Write an SQL query to report the shortest distance between any two points from the Point table.
-- The query result format is in the following example.

-- Example 1:
-- Input:
-- Point table:
-- +----+
-- | x  |
-- +----+
-- | -1 |
-- | 0  |
-- | 2  |
-- +----+
-- Output:
-- +----------+
-- | shortest |
-- +----------+
-- | 1        |
-- +----------+
-- Explanation: The shortest distance is between points -1 and 0 which is |(-1) - 0| = 1.
-- Follow up: How could you optimize your query if the Point table is ordered in ascending order?

-- Create Table If Not Exists Point (x int not null)
-- Truncate table Point
-- insert into Point (x) values ('-1')
-- insert into Point (x) values ('0')
-- insert into Point (x) values ('2')

-- Write your MySQL query statement below
SELECT
    MIN(ABS(a.x - b.x)) AS shortest -- 取最小的即可
FROM
    `point` AS a,
    `point` AS b
WHERE
    a.x != b.x -- 排除自己和自己(自己就是 0 了)

-- best solution
SELECT 
    MIN(diff) AS shortest
FROM 
(
	SELECT 
        x,
	    ABS(LEAD(x, 1) OVER(ORDER BY x) - x)  AS diff
	FROM 
        point
) AS t

--Lag 和 Lead 分析函数可以在同一次查询中取出同一字段的前 N 行的数据(Lag)和后 N 行的数据(Lead)作为独立的列。
-- 在实际应用当中，若要用到取今天和昨天的某字段差值时，Lag 和 Lead 函数的应用就显得尤为重要。
-- 当然，这种操作可以用表的自连接实现，但是 LAG 和 LEAD 与 left join、right join 等自连接相比，效率更高，SQL 更简洁。

-- # LAG 函数
-- LAG函数用于获取结果集中当前行之前的某一行的值。语法如下：

--         LAG (expression, offset, default) OVER (PARTITION BY partition_expression ORDER BY sort_expression)
--         expression: 要检索的列或表达式。
--         offset: 要返回的行数，如果不指定，默认为 1，即上一行。
--         efault: 当指定的行数超出结果集范围时，返回的默认值。

-- # LEAD 函数
-- LEAD函数用于获取结果集中当前行之后的某一行的值。语法如下：

--         LEAD (expression, offset, default) OVER (PARTITION BY partition_expression ORDER BY sort_expression)
--         expression: 要检索的列或表达式。
--         offset: 要返回的行数，如果不指定，默认为 1，即下一行。
--         default: 当指定的行数超出结果集范围时，返回的默认值。

-- # 使用案例
-- ## LAG 示例
-- 考虑一个名为sales的表，包含销售数据：
-- CREATE TABLE sales (
--     sale_date DATE,
--     revenue INT
-- );
 
-- INSERT INTO sales VALUES
-- ('2023-01-01', 100),
-- ('2023-01-02', 150),
-- ('2023-01-03', 200),
-- ('2023-01-04', 120);
-- 使用LAG函数，你可以获取前一天的销售额：
--         SELECT sale_date                                             -- 日期
--             , revenue                                               -- 当前销售额
--             , LAG(revenue) OVER (ORDER BY sale_date) AS lag_revenue -- 前一行的销售额
--         FROM sales;
-- 结果：
-- | sale_date  | revenue | lag_revenue |
-- |------------|---------|-------------|
-- | 2023-01-01 | 100     | NULL        |
-- | 2023-01-02 | 150     | 100         |
-- | 2023-01-03 | 200     | 150         |
-- | 2023-01-04 | 120     | 200         |

-- ## LEAD 示例
-- 使用LEAD函数，你可以获取后一天的销售额：
--         SELECT sale_date                                               -- 日期
--             , revenue                                                 -- 日期
--             , LEAD(revenue) OVER (ORDER BY sale_date) AS lead_revenue -- 后一行的销售额
--         FROM sales;
-- 结果：
-- | sale_date  | revenue | lead_revenue |
-- |------------|---------|--------------|
-- | 2023-01-01 | 100     | 150          |
-- | 2023-01-02 | 150     | 200          |
-- | 2023-01-03 | 200     | 120          |
-- | 2023-01-04 | 120     | NULL         |

-- 在这两个示例中，LAG和LEAD函数通过ORDER BY子句按销售日期对结果集进行排序。
-- 这允许你在时间序列数据中访问前一行或后一行的值，以进行比较或计算差异等操作