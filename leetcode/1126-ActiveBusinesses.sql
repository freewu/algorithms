-- 1126. Active Businesses
-- Table: Events
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | business_id   | int     |
-- | event_type    | varchar |
-- | occurrences   | int     | 
-- +---------------+---------+
-- (business_id, event_type) is the primary key (combination of columns with unique values) of this table.
-- Each row in the table logs the info that an event of some type occurred at some business for a number of times.
-- The average activity for a particular event_type is the average occurrences across all companies that have this event.
-- An active business is a business that has more than one event_type such that their occurrences is strictly greater than the average activity for that event.
-- Write a solution to find all active businesses.

-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Events table:
-- +-------------+------------+-------------+
-- | business_id | event_type | occurrences |
-- +-------------+------------+-------------+
-- | 1           | reviews    | 7           |
-- | 3           | reviews    | 3           |
-- | 1           | ads        | 11          |
-- | 2           | ads        | 7           |
-- | 3           | ads        | 6           |
-- | 1           | page views | 3           |
-- | 2           | page views | 12          |
-- +-------------+------------+-------------+
-- Output: 
-- +-------------+
-- | business_id |
-- +-------------+
-- | 1           |
-- +-------------+
-- Explanation:  
-- The average activity for each event can be calculated as follows:
-- - 'reviews': (7+3)/2 = 5
-- - 'ads': (11+7+6)/3 = 8
-- - 'page views': (3+12)/2 = 7.5
-- The business with id=1 has 7 'reviews' events (more than 5) and 11 'ads' events (more than 8), so it is an active business.

-- Create table If Not Exists Events (business_id int, event_type varchar(10), occurrences int)
-- Truncate table Events
-- insert into Events (business_id, event_type, occurrences) values ('1', 'reviews', '7')
-- insert into Events (business_id, event_type, occurrences) values ('3', 'reviews', '3')
-- insert into Events (business_id, event_type, occurrences) values ('1', 'ads', '11')
-- insert into Events (business_id, event_type, occurrences) values ('2', 'ads', '7')
-- insert into Events (business_id, event_type, occurrences) values ('3', 'ads', '6')
-- insert into Events (business_id, event_type, occurrences) values ('1', 'page views', '3')
-- insert into Events (business_id, event_type, occurrences) values ('2', 'page views', '12')

-- Write your MySQL query statement below
SELECT 
    e.business_id
FROM 
    Events  AS e
JOIN
( -- occurrences 均值
    SELECT 
        event_type, 
        AVG(occurrences) AS eventAvg
    FROM 
        Events
    GROUP BY 
        event_type
) AS a 
ON e.event_type = a.event_type
WHERE
    e.occurrences > a.eventAvg -- occurrences 严格大于 该事件的平均活动次数
GROUP BY 
    e.business_id
HAVING 
    COUNT(*) >= 2 -- 活跃业务 是指具有 多个 event_type 的业务，它们的 occurrences 严格大于 该事件的平均活动次数


-- best solution
SELECT 
    business_id
FROM 
(
    SELECT 
        *, 
        AVG(occurrences) OVER(PARTITION BY event_type) AS avg_occ
    FROM 
        events
) t
GROUP BY 
    business_id
HAVING 
    SUM(t.occurrences > avg_occ) > 1