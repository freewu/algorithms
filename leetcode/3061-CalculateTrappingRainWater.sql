-- 3061. Calculate Trapping Rain Water
-- Table: Heights
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | id          | int  |
-- | height      | int  |
-- +-------------+------+
-- id is the primary key (column with unique values) for this table, and it is guaranteed to be in sequential order.
-- Each row of this table contains an id and height.
-- Write a solution to calculate the amount of rainwater can be trapped between the bars in the landscape, considering that each bar has a width of 1 unit.

-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Heights table:
-- +-----+--------+
-- | id  | height |
-- +-----+--------+
-- | 1   | 0      |
-- | 2   | 1      |
-- | 3   | 0      |
-- | 4   | 2      |
-- | 5   | 1      |
-- | 6   | 0      |
-- | 7   | 1      |
-- | 8   | 3      |
-- | 9   | 2      |
-- | 10  | 1      |
-- | 11  | 2      |
-- | 12  | 1      |
-- +-----+--------+
-- Output: 
-- +---------------------+
-- | total_trapped_water | 
-- +---------------------+
-- | 6                   | 
-- +---------------------+
-- Explanation: 
-- <img src="https://assets.leetcode.com/uploads/2024/02/26/trapping_rain_water.png" />
-- The elevation map depicted above (in the black section) is graphically represented with the x-axis denoting the id and the y-axis representing the heights [0,1,0,2,1,0,1,3,2,1,2,1]. In this scenario, 6 units of rainwater are trapped within the blue section.

-- Create Table if not Exists Heights(id int, height int)
-- Truncate table Heights
-- insert into Heights (id, height) values ('1', '0')
-- insert into Heights (id, height) values ('2', '1')
-- insert into Heights (id, height) values ('3', '0')
-- insert into Heights (id, height) values ('4', '2')
-- insert into Heights (id, height) values ('5', '1')
-- insert into Heights (id, height) values ('6', '0')
-- insert into Heights (id, height) values ('7', '1')
-- insert into Heights (id, height) values ('8', '3')
-- insert into Heights (id, height) values ('9', '2')
-- insert into Heights (id, height) values ('10', '1')
-- insert into Heights (id, height) values ('11', '2')
-- insert into Heights (id, height) values ('12', '1')

-- Write your PostgreSQL query statement below
SELECT 
    COALESCE(sum(val), 0) AS total_trapped_water
FROM 
(
    SELECT 
        CASE
            WHEN lh > rh THEN rh - height
            ELSE lh - height 
        END AS val
    FROM
    (
        SELECT 
            b.id,
            b.height, 
            MAX(a.height) AS lh, 
            MAX(c.height) AS rh
        FROM 
            Heights AS a 
        JOIN 
            Heights AS  b 
        ON 
            a.id < b.id AND a.height > b.height
        JOIN 
            Heights AS c 
        ON  
            c.id > b.id AND c.height > b.height
        GROUP BY 
            b.id, b.height
    ) AS d 
) AS  e


select
    sum(case when k > height then k - height else 0 end) as total_trapped_water
from
(   select
        id,
        height,
        least( max(height)over(order by id), max(height)over(order by id desc) ) as k
    from Heights
) as t