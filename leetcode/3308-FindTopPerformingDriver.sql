-- 3308. Find Top Performing Driver
-- Table: Drivers
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | driver_id    | int     |
-- | name         | varchar |
-- | age          | int     |
-- | experience   | int     |
-- | accidents    | int     |
-- +--------------+---------+
-- (driver_id) is the unique key for this table.
-- Each row includes a driver's ID, their name, age, years of driving experience, and the number of accidents they’ve had.
-- Table: Vehicles
-- +--------------+---------+
-- | vehicle_id   | int     |
-- | driver_id    | int     |
-- | model        | varchar |
-- | fuel_type    | varchar |
-- | mileage      | int     |
-- +--------------+---------+
-- (vehicle_id, driver_id, fuel_type) is the unique key for this table.
-- Each row includes the vehicle's ID, the driver who operates it, the model, fuel type, and mileage.
-- Table: Trips
-- +--------------+---------+
-- | trip_id      | int     |
-- | vehicle_id   | int     |
-- | distance     | int     |
-- | duration     | int     |
-- | rating       | int     |
-- +--------------+---------+
-- (trip_id) is the unique key for this table.
-- Each row includes a trip's ID, the vehicle used, the distance covered (in miles), the trip duration (in minutes), and the passenger's rating (1-5).
-- Uber is analyzing drivers based on their trips. Write a solution to find the top-performing driver for each fuel type based on the following criteria:

-- A driver's performance is calculated as the average rating across all their trips. Average rating should be rounded to 2 decimal places.
-- If two drivers have the same average rating, the driver with the longer total distance traveled should be ranked higher.
-- If there is still a tie, choose the driver with the fewest accidents.
-- Return the result table ordered by fuel_type in ascending order.

-- The result format is in the following example.

-- Example:
-- Input:
-- Drivers table:
-- +-----------+----------+-----+------------+-----------+
-- | driver_id | name     | age | experience | accidents |
-- +-----------+----------+-----+------------+-----------+
-- | 1         | Alice    | 34  | 10         | 1         |
-- | 2         | Bob      | 45  | 20         | 3         |
-- | 3         | Charlie  | 28  | 5          | 0         |
-- +-----------+----------+-----+------------+-----------+
-- Vehicles table:
-- +------------+-----------+---------+-----------+---------+
-- | vehicle_id | driver_id | model   | fuel_type | mileage |
-- +------------+-----------+---------+-----------+---------+
-- | 100        | 1         | Sedan   | Gasoline  | 20000   |
-- | 101        | 2         | SUV     | Electric  | 30000   |
-- | 102        | 3         | Coupe   | Gasoline  | 15000   |
-- +------------+-----------+---------+-----------+---------+
-- Trips table:
-- +---------+------------+----------+----------+--------+
-- | trip_id | vehicle_id | distance | duration | rating |
-- +---------+------------+----------+----------+--------+
-- | 201     | 100        | 50       | 30       | 5      |
-- | 202     | 100        | 30       | 20       | 4      |
-- | 203     | 101        | 100      | 60       | 4      |
-- | 204     | 101        | 80       | 50       | 5      |
-- | 205     | 102        | 40       | 30       | 5      |
-- | 206     | 102        | 60       | 40       | 5      |
-- +---------+------------+----------+----------+--------+
-- Output:
-- +-----------+-----------+--------+----------+
-- | fuel_type | driver_id | rating | distance |
-- +-----------+-----------+--------+----------+
-- | Electric  | 2         | 4.50   | 180      |
-- | Gasoline  | 3         | 5.00   | 100      |
-- +-----------+-----------+--------+----------+
-- Explanation:
-- For fuel type Gasoline, both Alice (Driver 1) and Charlie (Driver 3) have trips. Charlie has an average rating of 5.0, while Alice has 4.5. Therefore, Charlie is selected.
-- For fuel type Electric, Bob (Driver 2) is the only driver with an average rating of 4.5, so he is selected.
-- The output table is ordered by fuel_type in ascending order.

-- CREATE TABLE If not exists Drivers (
--     driver_id INT ,
--     name VARCHAR(100),
--     age INT,
--     experience INT,
--     accidents INT
-- )
-- CREATE TABLE If not exists Vehicles (
--     vehicle_id INT ,
--     driver_id INT,
--     model VARCHAR(100),
--     fuel_type VARCHAR(50),
--     mileage INT)
-- CREATE TABLE  If not exists Trips (
--     trip_id INT ,
--     vehicle_id INT,
--     distance INT,
--     duration INT,
--     rating INT
-- )
-- Truncate table Drivers
-- insert into Drivers (driver_id, name, age, experience, accidents) values ('1', 'Alice', '34', '10', '1')
-- insert into Drivers (driver_id, name, age, experience, accidents) values ('2', 'Bob', '45', '20', '3')
-- insert into Drivers (driver_id, name, age, experience, accidents) values ('3', 'Charlie', '28', '5', '0')
-- Truncate table Vehicles
-- insert into Vehicles (vehicle_id, driver_id, model, fuel_type, mileage) values ('100', '1', 'Sedan', 'Gasoline', '20000')
-- insert into Vehicles (vehicle_id, driver_id, model, fuel_type, mileage) values ('101', '2', 'SUV', 'Electric', '30000')
-- insert into Vehicles (vehicle_id, driver_id, model, fuel_type, mileage) values ('102', '3', 'Coupe', 'Gasoline', '15000')
-- Truncate table Trips
-- insert into Trips (trip_id, vehicle_id, distance, duration, rating) values ('201', '100', '50', '30', '5')
-- insert into Trips (trip_id, vehicle_id, distance, duration, rating) values ('202', '100', '30', '20', '4')
-- insert into Trips (trip_id, vehicle_id, distance, duration, rating) values ('203', '101', '100', '60', '4')
-- insert into Trips (trip_id, vehicle_id, distance, duration, rating) values ('204', '101', '80', '50', '5')
-- insert into Trips (trip_id, vehicle_id, distance, duration, rating) values ('205', '102', '40', '30', '5')
-- insert into Trips (trip_id, vehicle_id, distance, duration, rating) values ('206', '102', '60', '40', '5')

-- Write your MySQL query statement below
WITH t AS (
    SELECT
        d.driver_id, 
        d.accidents,
        v.vehicle_id,
        v.fuel_type,
        SUM(t.distance) AS distance,
        ROUND(SUM(t.rating) / COUNT(*), 2) AS rating 
    FROM 
        Trips AS t
    LEFT JOIN  
        Vehicles AS v 
    ON
        t.vehicle_id = v.vehicle_id 
    LEFT JOIN  
        Drivers AS d 
    ON
        d.driver_id = v.driver_id
    GROUP BY
        d.driver_id, v.fuel_type 
),
-- SELECT * FROM t;
-- | driver_id | accidents | vehicle_id | fuel_type | distance | rating |
-- | --------- | --------- | ---------- | --------- | -------- | ------ |
-- | 1         | 1         | 100        | Gasoline  | 80       | 4.5    |
-- | 2         | 3         | 101        | Electric  | 180      | 4.5    |
-- | 3         | 0         | 102        | Gasoline  | 100      | 5      |
t1 AS (
    SELECT
        fuel_type,
        driver_id,
        rating,
        distance,
        RANK() OVER (
            PARTITION BY fuel_type 
            ORDER BY  rating DESC, distance DESC,  accidents ASC
        ) AS rk
    FROM 
        t
)
-- SELECT * FROM t1;
-- | fuel_type | driver_id | rating | distance | rk |
-- | --------- | --------- | ------ | -------- | -- |
-- | Electric  | 2         | 4.5    | 180      | 1  |
-- | Gasoline  | 3         | 5      | 100      | 1  |
-- | Gasoline  | 1         | 4.5    | 80       | 2  |
SELECT  
    fuel_type,
    driver_id,
    rating,
    distance
FROM 
    t1
WHERE
    rk = 1
ORDER BY 
    fuel_type
