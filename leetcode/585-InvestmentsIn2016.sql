-- 585. Investments in 2016
-- Table: Insurance
-- +-------------+-------+
-- | Column Name | Type  |
-- +-------------+-------+
-- | pid         | int   |
-- | tiv_2015    | float |
-- | tiv_2016    | float |
-- | lat         | float |
-- | lon         | float |
-- +-------------+-------+
-- pid is the primary key column for this table.
-- Each row of this table contains information about one policy where:
-- pid is the policyholder's policy ID.
-- tiv_2015 is the total investment value in 2015 and tiv_2016 is the total investment value in 2016.
-- lat is the latitude of the policy holder's city.
-- lon is the longitude of the policy holder's city.
--  
-- Write an SQL query to report the sum of all total investment values in 2016 tiv_2016, for all policyholders who:
--      have the same tiv_2015 value as one or more other policyholders, 
--      and are not located in the same city like any other policyholder (i.e., the (lat, lon) attribute pairs must be unique).

-- Round tiv_2016 to two decimal places.
-- The query result format is in the following example.
--
-- Example 1:
-- Input:
-- Insurance table:
-- +-----+----------+----------+-----+-----+
-- | pid | tiv_2015 | tiv_2016 | lat | lon |
-- +-----+----------+----------+-----+-----+
-- | 1   | 10       | 5        | 10  | 10  |
-- | 2   | 20       | 20       | 20  | 20  |
-- | 3   | 10       | 30       | 20  | 20  |
-- | 4   | 10       | 40       | 40  | 40  |
-- +-----+----------+----------+-----+-----+
-- Output:
-- +----------+
-- | tiv_2016 |
-- +----------+
-- | 45.00    |
-- +----------+
-- Explanation:
-- The first record in the table, like the last record, meets both of the two criteria.
-- The tiv_2015 value 10 is the same as the third and fourth records, and its location is unique.
-- The second record does not meet any of the two criteria.
-- Its tiv_2015 is not like any other policyholders and its location is the same as the third record, which makes the third record fail, too.
-- So, the result is the sum of tiv_2016 of the first and last record, which is 45.

-- Create Table If Not Exists Insurance (pid int, tiv_2015 float, tiv_2016 float, lat float, lon float)
-- Truncate table Insurance
-- insert into Insurance (pid, tiv_2015, tiv_2016, lat, lon) values ('1', '10', '5', '10', '10')
-- insert into Insurance (pid, tiv_2015, tiv_2016, lat, lon) values ('2', '20', '20', '20', '20')
-- insert into Insurance (pid, tiv_2015, tiv_2016, lat, lon) values ('3', '10', '30', '20', '20')
-- insert into Insurance (pid, tiv_2015, tiv_2016, lat, lon) values ('4', '10', '40', '40', '40')

-- Write your MySQL query statement below
SELECT
    ROUND(SUM(t.TIV_2016),2) AS TIV_2016
FROM
    (
        SELECT
            *,
            count(PID) over(partition by TIV_2015) cn1, -- 2015年相同投资
            count(PID) over(partition by LAT, LON) cn2 -- 相同城市
        FROM
            `insurance`
    ) AS t
WHERE
    t.cn1 > 1 AND -- 2015年的投资要存在和别人一样
    t.cn2 = 1 -- 一个城市只能有一个