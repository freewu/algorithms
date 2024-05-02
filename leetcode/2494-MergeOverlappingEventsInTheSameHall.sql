-- 2494. Merge Overlapping Events in the Same Hall
-- Table: HallEvents
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | hall_id     | int  |
-- | start_day   | date |
-- | end_day     | date |
-- +-------------+------+
-- This table may contain duplicates rows.
-- Each row of this table indicates the start day and end day of an event and the hall in which the event is held.
 
-- Write a solution to merge all the overlapping events that are held in the same hall. Two events overlap if they have at least one day in common.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- HallEvents table:
-- +---------+------------+------------+
-- | hall_id | start_day  | end_day    |
-- +---------+------------+------------+
-- | 1       | 2023-01-13 | 2023-01-14 |
-- | 1       | 2023-01-14 | 2023-01-17 |
-- | 1       | 2023-01-18 | 2023-01-25 |
-- | 2       | 2022-12-09 | 2022-12-23 |
-- | 2       | 2022-12-13 | 2022-12-17 |
-- | 3       | 2022-12-01 | 2023-01-30 |
-- +---------+------------+------------+
-- Output: 
-- +---------+------------+------------+
-- | hall_id | start_day  | end_day    |
-- +---------+------------+------------+
-- | 1       | 2023-01-13 | 2023-01-17 |
-- | 1       | 2023-01-18 | 2023-01-25 |
-- | 2       | 2022-12-09 | 2022-12-23 |
-- | 3       | 2022-12-01 | 2023-01-30 |
-- +---------+------------+------------+
-- Explanation: There are three halls.
-- Hall 1:
-- - The two events ["2023-01-13", "2023-01-14"] and ["2023-01-14", "2023-01-17"] overlap. We merge them in one event ["2023-01-13", "2023-01-17"].
-- - The event ["2023-01-18", "2023-01-25"] does not overlap with any other event, so we leave it as it is.
-- Hall 2:
-- - The two events ["2022-12-09", "2022-12-23"] and ["2022-12-13", "2022-12-17"] overlap. We merge them in one event ["2022-12-09", "2022-12-23"].
-- Hall 3:
-- - The hall has only one event, so we return it. Note that we only consider the events of each hall separately.

-- Create table If Not Exists HallEvents (hall_id int, start_day date, end_day date)
-- Truncate table HallEvents
-- insert into HallEvents (hall_id, start_day, end_day) values ('1', '2023-01-13', '2023-01-14')
-- insert into HallEvents (hall_id, start_day, end_day) values ('1', '2023-01-14', '2023-01-17')
-- insert into HallEvents (hall_id, start_day, end_day) values ('1', '2023-01-18', '2023-01-25')
-- insert into HallEvents (hall_id, start_day, end_day) values ('2', '2022-12-09', '2022-12-23')
-- insert into HallEvents (hall_id, start_day, end_day) values ('2', '2022-12-13', '2022-12-17')
-- insert into HallEvents (hall_id, start_day, end_day) values ('3', '2022-12-01', '2023-01-30')

-- Write your MySQL query statement below
SELECT 
    hall_id, 
    MIN(start_day) AS start_day, 
    MAX(end_day) AS end_day
FROM 
(
    SELECT 
        *, 
        @x := IFNULL(
            IF(
                start_day > MAX(end_day) OVER (partition by hall_id order by hall_id,start_day,end_day rows between unbounded preceding and 1 preceding),
                @x := @x + 1, @x
            ),
            @x := 1
        ) AS a
    FROM 
        HallEvents
) AS t
GROUP BY
    hall_id,a

-- best solution
SELECT 
    hall_id, 
    MIN(start_day) AS start_day, 
    MAX(end_day) AS end_day
FROM 
(
    SELECT 
        e.*,
        @period := IF
        (
            @id = hall_id, 
            IF( @end >= start_day, @period, @period + 1 ),
            @period + 1
        ) AS period,
        @end := IF(@id = hall_id AND @end > end_day, @end, end_day),
        @id:= hall_id
    FROM 
        hallevents AS e, 
        (SELECT @id:=0, @period:=0, @end:= null) AS v
    ORDER BY 
        hall_id, start_day
) AS t
GROUP BY
    hall_id, period