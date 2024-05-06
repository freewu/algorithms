-- 100325. Consecutive Available Seats II
-- Table: Cinema
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | seat_id     | int  |
-- | free        | bool |
-- +-------------+------+
-- seat_id is an auto-increment column for this table.
-- Each row of this table indicates whether the ith seat is free or not. 1 means free while 0 means occupied.
-- Write a solution to find the length of longest consecutive sequence of available seats in the cinema.

-- Note:
--     There will always be at most one longest consecutive sequence.
--     If there are multiple consecutive sequences with the same length, include all of them in the output.

-- Return the result table ordered by first_seat_id in ascending order.
-- The result format is in the following example.

-- Example:
-- Input:
-- Cinema table:
-- +---------+------+
-- | seat_id | free |
-- +---------+------+
-- | 1       | 1    |
-- | 2       | 0    |
-- | 3       | 1    |
-- | 4       | 1    |
-- | 5       | 1    |
-- +---------+------+
-- Output:
-- +-----------------+----------------+-----------------------+
-- | first_seat_id   | last_seat_id   | consecutive_seats_len |
-- +-----------------+----------------+-----------------------+
-- | 3               | 5              | 3                     |
-- +-----------------+----------------+-----------------------+
-- Explanation:
-- Longest consecutive sequence of available seats starts from seat 3 and ends at seat 5 with a length of 3.
-- Output table is ordered by first_seat_id in ascending order.

-- CREATE TABLE if Not exists Cinema (
--     seat_id INT PRIMARY KEY AUTO_INCREMENT,
--     free BOOLEAN
-- )
-- Truncate table Cinema
-- insert into Cinema (seat_id, free) values ('1', '1')
-- insert into Cinema (seat_id, free) values ('2', '0')
-- insert into Cinema (seat_id, free) values ('3', '1')
-- insert into Cinema (seat_id, free) values ('4', '1')
-- insert into Cinema (seat_id, free) values ('5', '1')

-- # Write your MySQL query statement below
-- WITH t AS (
--     SELECT 
--         *,
--         seat_id - (RANK() OVER(ORDER BY seat_id )) AS diff 
--     FROM 
--         (SELECT * FROM Cinema WHERE free = 1) AS c
-- )
-- -- SELECT * FROM t 
-- -- | seat_id | free | diff |
-- -- | ------- | ---- | ---- |
-- -- | 1       | 1    | 0    |
-- -- | 3       | 1    | 1    |
-- -- | 4       | 1    | 1    |
-- -- | 5       | 1    | 1    |
-- SELECT 
--     MIN(seat_id) AS first_seat_id,
--     MAX(seat_id) AS last_seat_id,
--     COUNT(*) AS consecutive_seats_len 
-- FROM
--     t 
-- WHERE
--     diff =  (
--                 SELECT 
--                     diff 
--                 FROM 
--                     t 
--                 GROUP BY 
--                     diff 
--                 ORDER BY 
--                     COUNT(*) DESC
--                 LIMIT 1 
--             )

WITH t AS (
    SELECT 
        *,
        seat_id - (RANK() OVER(ORDER BY seat_id )) AS diff 
    FROM 
        (SELECT * FROM Cinema WHERE free = 1) AS c
),
s AS (
    SELECT 
        diff,
        COUNT(*) AS cnt
    FROM 
        t 
    GROUP BY
        diff
),
c AS (
    SELECT 
        *,
        RANK() OVER(PARTITION BY diff ORDER BY seat_id) AS rk
    FROM 
        t
    WHERE
        diff IN ( 
            SELECT diff FROM s WHERE cnt = (SELECT MAX(cnt) FROM s)
        )
)
-- SELECT * FROM s

-- SELECT 
--     MIN(seat_id) AS first_seat_id,
--     MAX(seat_id) AS last_seat_id,
--     COUNT(*) AS consecutive_seats_len 
-- FROM
--     t 
-- WHERE
--     diff IN ( 
--         SELECT diff FROM s WHERE cnt = (SELECT MAX(cnt) FROM s)
--     )

-- SELECT 
--     *,
--     RANK() OVER(PARTITION BY diff ORDER BY seat_id) AS rk
-- FROM 
--     t
-- WHERE
--     diff IN ( 
--         SELECT diff FROM s WHERE cnt = (SELECT MAX(cnt) FROM s)
--     )
-- | seat_id | free | diff | rk |
-- | ------- | ---- | ---- | -- |
-- | 21      | 1    | 10   | 1  |
-- | 22      | 1    | 10   | 2  |
-- | 23      | 1    | 10   | 3  |
-- | 24      | 1    | 10   | 4  |
-- | 31      | 1    | 14   | 1  |
-- | 32      | 1    | 14   | 2  |
-- | 33      | 1    | 14   | 3  |
-- | 34      | 1    | 14   | 4  |

SELECT 
    MIN(seat_id) AS first_seat_id,
    MAX(seat_id) AS last_seat_id,
    COUNT(*) AS consecutive_seats_len 
FROM
    c 
GROUP BY 
    diff 