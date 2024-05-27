-- 3166. Calculate Parking Fees and Duration
-- Table: ParkingTransactions
-- +--------------+-----------+
-- | Column Name  | Type      |
-- +--------------+-----------+
-- | lot_id       | int       |
-- | car_id       | int       |
-- | entry_time   | datetime  |
-- | exit_time    | datetime  |
-- | fee_paid     | decimal   |
-- +--------------+-----------+
-- (lot_id, car_id, entry_time) is the primary key (combination of columns with unique values) for this table.
-- Each row of this table contains the ID of the parking lot, the ID of the car, the entry and exit times, and the fee paid for the parking duration.
-- Write a solution to find the total parking fee paid by each car across all parking lots, and the average hourly fee (rounded to 2 decimal places) paid by each car. Also, find the parking lot where each car spent the most total time.

-- Return the result table ordered by car_id in ascending order.
-- Note: Test cases are generated in such a way that an individual car cannot be in multiple parking lots at the same time.
-- The result format is in the following example.

-- Example:
-- Input:
-- ParkingTransactions table:
-- +--------+--------+---------------------+---------------------+----------+
-- | lot_id | car_id | entry_time          | exit_time           | fee_paid |
-- +--------+--------+---------------------+---------------------+----------+
-- | 1      | 1001   | 2023-06-01 08:00:00 | 2023-06-01 10:30:00 | 5.00     |
-- | 1      | 1001   | 2023-06-02 11:00:00 | 2023-06-02 12:45:00 | 3.00     |
-- | 2      | 1001   | 2023-06-01 10:45:00 | 2023-06-01 12:00:00 | 6.00     |
-- | 2      | 1002   | 2023-06-01 09:00:00 | 2023-06-01 11:30:00 | 4.00     |
-- | 3      | 1001   | 2023-06-03 07:00:00 | 2023-06-03 09:00:00 | 4.00     |
-- | 3      | 1002   | 2023-06-02 12:00:00 | 2023-06-02 14:00:00 | 2.00     |
-- +--------+--------+---------------------+---------------------+----------+
-- Output:
-- +--------+----------------+----------------+---------------+
-- | car_id | total_fee_paid | avg_hourly_fee | most_time_lot |
-- +--------+----------------+----------------+---------------+
-- | 1001   | 18.00          | 2.40           | 1             |
-- | 1002   | 6.00           | 1.33           | 2             |
-- +--------+----------------+----------------+---------------+
-- Explanation:
-- For car ID 1001:
-- From 2023-06-01 08:00:00 to 2023-06-01 10:30:00 in lot 1: 2.5 hours, fee 5.00
-- From 2023-06-02 11:00:00 to 2023-06-02 12:45:00 in lot 1: 1.75 hours, fee 3.00
-- From 2023-06-01 10:45:00 to 2023-06-01 12:00:00 in lot 2: 1.25 hours, fee 6.00
-- From 2023-06-03 07:00:00 to 2023-06-03 09:00:00 in lot 3: 2 hours, fee 4.00
-- Total fee paid: 18.00, total hours: 7.5, average hourly fee: 2.40, most time spent in lot 1: 2.5 hours.
-- For car ID 1002:
-- From 2023-06-01 09:00:00 to 2023-06-01 11:30:00 in lot 2: 2.5 hours, fee 4.00
-- From 2023-06-02 12:00:00 to 2023-06-02 14:00:00 in lot 3: 2 hours, fee 2.00
-- Total fee paid: 6.00, total hours: 4.5, average hourly fee: 1.33, most time spent in lot 2: 2.5 hours.
-- Note: Output table is ordered by car_id in ascending order.

-- CREATE TABLE If not exists ParkingTransactions (
--     lot_id INT,
--     car_id INT,
--     entry_time DATETIME,
--     exit_time DATETIME,
--     fee_paid DECIMAL(10, 2)
-- )

-- Truncate table ParkingTransactions
-- insert into ParkingTransactions (lot_id, car_id, entry_time, exit_time, fee_paid) values ('1', '1001', '2023-06-01 08:00:00', '2023-06-01 10:30:00', '5.0')
-- insert into ParkingTransactions (lot_id, car_id, entry_time, exit_time, fee_paid) values ('1', '1001', '2023-06-02 11:00:00', '2023-06-02 12:45:00', '3.0')
-- insert into ParkingTransactions (lot_id, car_id, entry_time, exit_time, fee_paid) values ('2', '1001', '2023-06-01 10:45:00', '2023-06-01 12:00:00', '6.0')
-- insert into ParkingTransactions (lot_id, car_id, entry_time, exit_time, fee_paid) values ('2', '1002', '2023-06-01 09:00:00', '2023-06-01 11:30:00', '4.0')
-- insert into ParkingTransactions (lot_id, car_id, entry_time, exit_time, fee_paid) values ('3', '1001', '2023-06-03 07:00:00', '2023-06-03 09:00:00', '4.0')
-- insert into ParkingTransactions (lot_id, car_id, entry_time, exit_time, fee_paid) values ('3', '1002', '2023-06-02 12:00:00', '2023-06-02 14:00:00', '2.0')

WITH r AS ( -- 统计每次停车时长
    SELECT 
        lot_id,
        car_id,
        TIMESTAMPDIFF(SECOND, entry_time, exit_time) / 3600 AS amount_time,
        fee_paid
    FROM
       ParkingTransactions  
), 
p AS ( -- 每车辆在不同停车场的停车时长 & 排名
    SELECT 
        lot_id,
        car_id,
        SUM(amount_time),
        RANK() OVER(PARTITION BY car_id ORDER BY SUM(amount_time) DESC) AS rk 
    FROM 
        r
    GROUP BY 
        car_id, lot_id
)
-- SELECT * FROM r
-- SELECT * FROM p

SELECT 
    r.car_id,
    SUM(fee_paid) AS total_fee_paid,
    ROUND(SUM(fee_paid) / SUM(amount_time), 2) AS  avg_hourly_fee,
    p.lot_id AS most_time_lot 
FROM
    r
LEFT JOIN
    p
ON 
    r.car_id = p.car_id AND p.rk = 1
GROUP BY
    r.car_id
ORDER BY 
    r.car_id