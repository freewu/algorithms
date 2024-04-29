-- 2153. The Number of Passengers in Each Bus II
-- Table: Buses
-- +--------------+------+
-- | Column Name  | Type |
-- +--------------+------+
-- | bus_id       | int  |
-- | arrival_time | int  |
-- | capacity     | int  |
-- +--------------+------+
-- bus_id contains unique values.
-- Each row of this table contains information about the arrival time of a bus at the LeetCode station and its capacity (the number of empty seats it has).
-- No two buses will arrive at the same time and all bus capacities will be positive integers.
 
-- Table: Passengers
-- +--------------+------+
-- | Column Name  | Type |
-- +--------------+------+
-- | passenger_id | int  |
-- | arrival_time | int  |
-- +--------------+------+
-- passenger_id contains unique values.
-- Each row of this table contains information about the arrival time of a passenger at the LeetCode station.
 
-- Buses and passengers arrive at the LeetCode station. 
-- If a bus arrives at the station at a time tbus and a passenger arrived at a time tpassenger where tpassenger <= tbus and the passenger did not catch any bus, the passenger will use that bus. In addition, each bus has a capacity. If at the moment the bus arrives at the station there are more passengers waiting than its capacity capacity, only capacity passengers will use the bus.
-- Write a solution to report the number of users that used each bus.
-- Return the result table ordered by bus_id in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Buses table:
-- +--------+--------------+----------+
-- | bus_id | arrival_time | capacity |
-- +--------+--------------+----------+
-- | 1      | 2            | 1        |
-- | 2      | 4            | 10       |
-- | 3      | 7            | 2        |
-- +--------+--------------+----------+
-- Passengers table:
-- +--------------+--------------+
-- | passenger_id | arrival_time |
-- +--------------+--------------+
-- | 11           | 1            |
-- | 12           | 1            |
-- | 13           | 5            |
-- | 14           | 6            |
-- | 15           | 7            |
-- +--------------+--------------+
-- Output: 
-- +--------+----------------+
-- | bus_id | passengers_cnt |
-- +--------+----------------+
-- | 1      | 1              |
-- | 2      | 1              |
-- | 3      | 2              |
-- +--------+----------------+
-- Explanation: 
-- - Passenger 11 arrives at time 1.
-- - Passenger 12 arrives at time 1.
-- - Bus 1 arrives at time 2 and collects passenger 11 as it has one empty seat.
-- - Bus 2 arrives at time 4 and collects passenger 12 as it has ten empty seats.
-- - Passenger 12 arrives at time 5.
-- - Passenger 13 arrives at time 6.
-- - Passenger 14 arrives at time 7.
-- - Bus 3 arrives at time 7 and collects passengers 12 and 13 as it has two empty seats.

-- Create table If Not Exists Buses (bus_id int, arrival_time int, capacity int)
-- Create table If Not Exists Passengers (passenger_id int, arrival_time int)
-- Truncate table Buses
-- insert into Buses (bus_id, arrival_time, capacity) values ('1', '2', '1')
-- insert into Buses (bus_id, arrival_time, capacity) values ('2', '4', '10')
-- insert into Buses (bus_id, arrival_time, capacity) values ('3', '7', '2')
-- Truncate table Passengers
-- insert into Passengers (passenger_id, arrival_time) values ('11', '1')
-- insert into Passengers (passenger_id, arrival_time) values ('12', '1')
-- insert into Passengers (passenger_id, arrival_time) values ('13', '5')
-- insert into Passengers (passenger_id, arrival_time) values ('14', '6')
-- insert into Passengers (passenger_id, arrival_time) values ('15', '7')


# Write your MySQL query statement below
WITH RECURSIVE bus_lag_info AS (
    SELECT 
        *,
        LAG(arrival_time, 1, 0) OVER(ORDER BY arrival_time) AS prev_arrival_time -- 先通过LAG()开窗获得上一班车的到达时间
    FROM 
        Buses 
),
passenger_arrival_between_buses_info AS ( -- 统计出两辆bus间的到达的乘客
    SELECT
        b.bus_id,
        b.arrival_time,
        b.capacity,
        COUNT(p.passenger_id) AS passenger_arrival_between_buses_cnt,
        ROW_NUMBER() OVER(ORDER BY b.arrival_time) AS rn
    FROM 
        bus_lag_info AS b
    LEFT JOIN 
        Passengers AS p
    ON 
        p.arrival_time > b.prev_arrival_time AND p.arrival_time <= b.arrival_time
    GROUP BY 
        1, 2, 3
),
boarded_info AS (
    ( -- 第一辆bus到达
        SELECT
            bus_id,
            LEAST(capacity, passenger_arrival_between_buses_cnt) AS boarded_cnt,
            GREATEST(0, passenger_arrival_between_buses_cnt - capacity) AS left_cnt,
            rn
        FROM 
            passenger_arrival_between_buses_info
        WHERE 
            rn = 1
    )
    UNION ALL
    (
        SELECT
            p.bus_id,
            LEAST(capacity, left_cnt + passenger_arrival_between_buses_cnt) AS boarded_cnt,
            GREATEST(0, left_cnt + passenger_arrival_between_buses_cnt - capacity) AS left_cnt,
            p.rn
        FROM 
            passenger_arrival_between_buses_info p, boarded_info b
        WHERE 
            p.rn = b.rn + 1
    )
)

SELECT 
    bus_id,
    boarded_cnt AS passengers_cnt
FROM 
    boarded_info
ORDER BY 
    bus_id

-- use variable
WITH t AS (
    SELECT
        bus_id, 
        capacity, 
        b.arrival_time AS time,
        count(passenger_id) AS total_passenger
    FROM 
        Buses AS b
    LEFT JOIN 
        Passengers AS p 
    ON 
        b.arrival_time >= p.arrival_time
    GROUP BY
        bus_id
)

SELECT 
    bus_id, 
    passengers_cnt
FROM 
    (
        SELECT 
            bus_id, 
            LEAST(total_passenger - @passenger_before, capacity)  AS passengers_cnt, 
            capacity, 
            time,
            @passenger_before:= LEAST(total_passenger, @passenger_before+capacity) AS p
        FROM 
            (select @passenger_before := 0) as init, 
            t
    ) AS a
ORDER BY 
    bus_id