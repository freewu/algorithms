-- 571. Find Median Given Frequency of Numbers
-- Table: Numbers
--
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | num         | int  |
-- | frequency   | int  |
-- +-------------+------+
-- num is the primary key for this table.
-- Each row of this table shows the frequency of a number in the database.
--  
--
-- The median is the value separating the higher half from the lower half of a data sample.
--
-- Write an SQL query to report the median of all the numbers in the database after decompressing the Numbers table. Round the median to one decimal point.
--
-- The query result format is in the following example.
--
--  
--
-- Example 1:
--
-- Input:
-- Numbers table:
-- +-----+-----------+
-- | num | frequency |
-- +-----+-----------+
-- | 0   | 7         |
-- | 1   | 1         |
-- | 2   | 3         |
-- | 3   | 1         |
-- +-----+-----------+
-- Output:
-- +--------+
-- | median |
-- +--------+
-- | 0.0    |
-- +--------+
-- Explanation:
-- If we decompress the Numbers table, we will get [0, 0, 0, 0, 0, 0, 0, 1, 2, 2, 2, 3], so the median is (0 + 0) / 2 = 0.
--

--  Write your MySQL query statement below
SELECT
    ROUND(AVG(num),1) AS median
FROM
    (
        SELECT
            num,
            SUM(frequency) OVER(ORDER BY num asc) AS asc_amount,  -- 从小到大
            SUM(frequency) OVER(ORDER BY num desc) AS desc_amount, -- 从大到小
            SUM(frequency) OVER() AS total_num
        FROM
            Numbers
    ) AS a
WHERE
    asc_amount >= total_num / 2 AND
    desc_amount >= total_num / 2