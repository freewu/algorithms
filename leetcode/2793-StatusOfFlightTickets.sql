-- 2793. Status of Flight Tickets
-- Table: Flights
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | flight_id   | int  |
-- | capacity    | int  |
-- +-------------+------+
-- flight_id column contains distinct values.
-- Each row of this table contains flight id and capacity.

-- Table: Passengers
-- +--------------+----------+
-- | Column Name  | Type     |
-- +--------------+----------+
-- | passenger_id | int      |
-- | flight_id    | int      |
-- | booking_time | datetime |
-- +--------------+----------+
-- passenger_id column contains distinct values.
-- booking_time column contains distinct values.
-- Each row of this table contains passenger id, booking time, and their flight id.
-- Passengers book tickets for flights in advance. If a passenger books a ticket for a flight and there are still empty seats available on the flight, the passenger's ticket will be confirmed. However, the passenger will be on a waitlist if the flight is already at full capacity.

-- Write a solution to determine the current status of flight tickets for each passenger.
-- Return the result table ordered by passenger_id in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Flights table:
-- +-----------+----------+
-- | flight_id | capacity |
-- +-----------+----------+
-- | 1         | 2        |
-- | 2         | 2        |
-- | 3         | 1        |
-- +-----------+----------+
-- Passengers table:
-- +--------------+-----------+---------------------+
-- | passenger_id | flight_id | booking_time        |
-- +--------------+-----------+---------------------+
-- | 101          | 1         | 2023-07-10 16:30:00 |
-- | 102          | 1         | 2023-07-10 17:45:00 |
-- | 103          | 1         | 2023-07-10 12:00:00 |
-- | 104          | 2         | 2023-07-05 13:23:00 |
-- | 105          | 2         | 2023-07-05 09:00:00 |
-- | 106          | 3         | 2023-07-08 11:10:00 |
-- | 107          | 3         | 2023-07-08 09:10:00 |
-- +--------------+-----------+---------------------+
-- Output: 
-- +--------------+-----------+
-- | passenger_id | Status    |
-- +--------------+-----------+
-- | 101          | Confirmed | 
-- | 102          | Waitlist  | 
-- | 103          | Confirmed | 
-- | 104          | Confirmed | 
-- | 105          | Confirmed | 
-- | 106          | Waitlist  | 
-- | 107          | Confirmed | 
-- +--------------+-----------+
-- Explanation: 
-- - Flight 1 has a capacity of 2 passengers. Passenger 101 and Passenger 103 were the first to book tickets, securing the available seats. Therefore, their bookings are confirmed. However, Passenger 102 was the third person to book a ticket for this flight, which means there are no more available seats. Passenger 102 is now placed on the waitlist, 
-- - Flight 2 has a capacity of 2 passengers, Flight 2 has exactly two passengers who booked tickets,  Passenger 104 and Passenger 105. Since the number of passengers who booked tickets matches the available seats, both bookings are confirmed.
-- - Flight 3 has a capacity of 1 passenger. Passenger 107 booked earlier and secured the only available seat, confirming their booking. Passenger 106, who booked after Passenger 107, is on the waitlist.

-- Create table if not exists Flights(flight_id int,capacity int)
-- Create table if not exists Passengers (passenger_id int,flight_id int,booking_time datetime)
-- Truncate table Flights
-- insert into Flights (flight_id, capacity) values ('1', '2')
-- insert into Flights (flight_id, capacity) values ('2', '2')
-- insert into Flights (flight_id, capacity) values ('3', '1')
-- Truncate table Passengers
-- insert into Passengers (passenger_id, flight_id, booking_time) values ('101', '1', '2023-07-10 16:30:00')
-- insert into Passengers (passenger_id, flight_id, booking_time) values ('102', '1', '2023-07-10 17:45:00')
-- insert into Passengers (passenger_id, flight_id, booking_time) values ('103', '1', '2023-07-10 12:00:00')
-- insert into Passengers (passenger_id, flight_id, booking_time) values ('104', '2', '2023-07-05 13:23:00')
-- insert into Passengers (passenger_id, flight_id, booking_time) values ('105', '2', '2023-07-05 09:00:00')
-- insert into Passengers (passenger_id, flight_id, booking_time) values ('106', '3', '2023-07-08 11:10:00')
-- insert into Passengers (passenger_id, flight_id, booking_time) values ('107', '3', '2023-07-08 09:10:00')

-- Write your MySQL query statement below
WITH t AS ( -- 按航班号分组按预定时间(小->大)给旅客编号
    SELECT 
        *,
        ROW_NUMBER() OVER(PARTITION BY flight_id ORDER BY booking_time) AS rn
    FROM
        Passengers
)
-- SELECT * FROM t 
-- | passenger_id | flight_id | booking_time        | rn |
-- | ------------ | --------- | ------------------- | -- |
-- | 103          | 1         | 2023-07-10 12:00:00 | 1  |
-- | 101          | 1         | 2023-07-10 16:30:00 | 2  |
-- | 102          | 1         | 2023-07-10 17:45:00 | 3  |
-- | 105          | 2         | 2023-07-05 09:00:00 | 1  |
-- | 104          | 2         | 2023-07-05 13:23:00 | 2  |
-- | 107          | 3         | 2023-07-08 09:10:00 | 1  |
-- | 106          | 3         | 2023-07-08 11:10:00 | 2  |

SELECT 
    t.passenger_id,
    CASE 
        WHEN t.rn <= f.capacity THEN "Confirmed"
        ELSE "Waitlist"
    END AS Status
FROM 
    t
LEFT JOIN 
    Flights AS f 
ON 
    t.flight_id  = f.flight_id 
ORDER BY
    passenger_id 