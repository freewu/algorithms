-- 2238. Number of Times a Driver Was a Passenger
-- Table: Rides
-- +--------------+------+
-- | Column Name  | Type |
-- +--------------+------+
-- | ride_id      | int  |
-- | driver_id    | int  |
-- | passenger_id | int  |
-- +--------------+------+
-- ride_id is the column with unique values for this table.
-- Each row of this table contains the ID of the driver and the ID of the passenger that rode in ride_id.
-- Note that driver_id != passenger_id.
 
-- Write a solution to report the ID of each driver and the number of times they were a passenger.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Rides table:
-- +---------+-----------+--------------+
-- | ride_id | driver_id | passenger_id |
-- +---------+-----------+--------------+
-- | 1       | 7         | 1            |
-- | 2       | 7         | 2            |
-- | 3       | 11        | 1            |
-- | 4       | 11        | 7            |
-- | 5       | 11        | 7            |
-- | 6       | 11        | 3            |
-- +---------+-----------+--------------+
-- Output: 
-- +-----------+-----+
-- | driver_id | cnt |
-- +-----------+-----+
-- | 7         | 2   |
-- | 11        | 0   |
-- +-----------+-----+
-- Explanation: 
-- There are two drivers in all the given rides: 7 and 11.
-- The driver with ID = 7 was a passenger two times.
-- The driver with ID = 11 was never a passenger.

-- Create table If Not Exists Rides (ride_id int, driver_id int, passenger_id int)
-- Truncate table Rides
-- insert into Rides (ride_id, driver_id, passenger_id) values ('1', '7', '1')
-- insert into Rides (ride_id, driver_id, passenger_id) values ('2', '7', '2')
-- insert into Rides (ride_id, driver_id, passenger_id) values ('3', '11', '1')
-- insert into Rides (ride_id, driver_id, passenger_id) values ('4', '11', '7')
-- insert into Rides (ride_id, driver_id, passenger_id) values ('5', '11', '7')
-- insert into Rides (ride_id, driver_id, passenger_id) values ('6', '11', '3')

-- Write your MySQL query statement below
WITH 
d AS ( -- 汇总司机信息
    SELECT DISTINCT driver_id FROM Rides 
),
r AS ( -- 汇总乘客 & 乘坐次数
    SELECT 
        passenger_id,
        COUNT(*) AS cnt
    FROM
        Rides
    GROUP BY
        passenger_id
)
SELECT
    d.driver_id,
    IFNULL(r.cnt, 0) AS cnt 
FROM
    d
LEFT JOIN 
    r 
ON 
    d.driver_id = r.passenger_id