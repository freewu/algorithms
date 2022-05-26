-- 626. Exchange Seats
-- Table: Seat
--
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | name        | varchar |
-- +-------------+---------+
-- id is the primary key column for this table.
-- Each row of this table indicates the name and the ID of a student.
-- id is a continuous increment.
--  
-- Write an SQL query to swap the seat id of every two consecutive students.
-- If the number of students is odd, the id of the last student is not swapped.
-- Return the result table ordered by id in ascending order.
--
-- The query result format is in the following example.
--
-- Example 1:
--
-- Input:
-- Seat table:
-- +----+---------+
-- | id | student |
-- +----+---------+
-- | 1  | Abbot   |
-- | 2  | Doris   |
-- | 3  | Emerson |
-- | 4  | Green   |
-- | 5  | Jeames  |
-- +----+---------+
-- Output:
-- +----+---------+
-- | id | student |
-- +----+---------+
-- | 1  | Doris   |
-- | 2  | Abbot   |
-- | 3  | Green   |
-- | 4  | Emerson |
-- | 5  | Jeames  |
-- +----+---------+
-- Explanation:
-- Note that if the number of students is odd, there is no need to change the last one's seat.
--
-- Write your MySQL query statement below
SELECT
    (
        CASE
            WHEN MOD(id,2) = 1 AND id = (SELECT COUNT(*) FROM seat) THEN id -- 如果是奇数最后一个同学保持不变
            WHEN MOD(id,2) = 1 THEN id + 1 -- 奇数向后
            ElSE id - 1 -- 偶数向前
        END
    ) AS id,
    student
FROM
    seat
ORDER BY
    id;