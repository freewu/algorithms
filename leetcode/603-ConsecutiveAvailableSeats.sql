-- 603. Consecutive Available Seats
-- Table: Cinema
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | seat_id     | int  |
-- | free        | bool |
-- +-------------+------+
-- seat_id is an auto-increment primary key column for this table.
-- Each row of this table indicates whether the ith seat is free or not. 1 means free while 0 means occupied.
--  
-- Write an SQL query to report all the consecutive available seats in the cinema.
-- Return the result table ordered by seat_id in ascending order.
-- The test cases are generated so that more than two seats are consecutively available.
-- The query result format is in the following example.
--
-- Example 1:
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
-- +---------+
-- | seat_id |
-- +---------+
-- | 3       |
-- | 4       |
-- | 5       |
-- +---------+

-- Create table If Not Exists Cinema (seat_id int primary key auto_increment, free bool)
-- Truncate table Cinema
-- insert into Cinema (seat_id, free) values ('1', '1')
-- insert into Cinema (seat_id, free) values ('2', '0')
-- insert into Cinema (seat_id, free) values ('3', '1')
-- insert into Cinema (seat_id, free) values ('4', '1')
-- insert into Cinema (seat_id, free) values ('5', '1')


-- Write your MySQL query statement below
SELECT
    DISTINCT a.seat_id AS seat_id
FROM
    Cinema AS a,
    Cinema AS b
WHERE
    (a.seat_id = b.seat_id + 1 OR a.seat_id = b.seat_id - 1) AND
    a.free = 1 AND b.free = 1
ORDER BY
    a.seat_id

-- best solution
SELECT 
    DISTINCT c1.seat_id AS seat_id
FROM 
    cinema AS c1,cinema AS c2
WHERE 
    c1.free =1 AND 
    c2.free=1 AND 
    ABS(c1.seat_id - c2.seat_id) = 1
ORDER BY 
    c1.seat_id;