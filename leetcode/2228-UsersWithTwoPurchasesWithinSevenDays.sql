-- 2228. Users With Two Purchases Within Seven Days
-- Table: Purchases
-- +---------------+------+
-- | Column Name   | Type |
-- +---------------+------+
-- | purchase_id   | int  |
-- | user_id       | int  |
-- | purchase_date | date |
-- +---------------+------+
-- purchase_id contains unique values.
-- This table contains logs of the dates that users purchased from a certain retailer.

-- Write a solution to report the IDs of the users that made any two purchases at most 7 days apart.
-- Return the result table ordered by user_id.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Purchases table:
-- +-------------+---------+---------------+
-- | purchase_id | user_id | purchase_date |
-- +-------------+---------+---------------+
-- | 4           | 2       | 2022-03-13    |
-- | 1           | 5       | 2022-02-11    |
-- | 3           | 7       | 2022-06-19    |
-- | 6           | 2       | 2022-03-20    |
-- | 5           | 7       | 2022-06-19    |
-- | 2           | 2       | 2022-06-08    |
-- +-------------+---------+---------------+
-- Output: 
-- +---------+
-- | user_id |
-- +---------+
-- | 2       |
-- | 7       |
-- +---------+
-- Explanation: 
-- User 2 had two purchases on 2022-03-13 and 2022-03-20. Since the second purchase is within 7 days of the first purchase, we add their ID.
-- User 5 had only 1 purchase.
-- User 7 had two purchases on the same day so we add their ID.

-- Create table If Not Exists Purchases (purchase_id int, user_id int, purchase_date date)
-- Truncate table Purchases
-- insert into Purchases (purchase_id, user_id, purchase_date) values ('4', '2', '2022-03-13')
-- insert into Purchases (purchase_id, user_id, purchase_date) values ('1', '5', '2022-02-11')
-- insert into Purchases (purchase_id, user_id, purchase_date) values ('3', '7', '2022-06-19')
-- insert into Purchases (purchase_id, user_id, purchase_date) values ('6', '2', '2022-03-20')
-- insert into Purchases (purchase_id, user_id, purchase_date) values ('5', '7', '2022-06-19')
-- insert into Purchases (purchase_id, user_id, purchase_date) values ('2', '2', '2022-06-08')

SELECT 
    DISTINCT a.user_id
    -- a.*,
    -- b.*,
    -- ABS(DATEDIFF(b.purchase_date, a.purchase_date)) AS diff 
FROM
    Purchases AS a
LEFT JOIN 
    Purchases AS b
ON
   a.user_id = b.user_id AND 
   a.purchase_id  != b.purchase_id -- 不关联相同的
WHERE 
    b.purchase_id IS NOT NULL AND
    ABS(DATEDIFF(b.purchase_date, a.purchase_date)) <= 7 --  any two purchases at most 7 days apart.
ORDER BY 
    user_id -- Return the result table ordered by user_id.