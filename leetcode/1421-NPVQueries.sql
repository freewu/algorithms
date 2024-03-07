-- 1421. NPV Queries
-- Table: NPV
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | year          | int     |
-- | npv           | int     |
-- +---------------+---------+
-- (id, year) is the primary key (combination of columns with unique values) of this table.
-- The table has information about the id and the year of each inventory and the corresponding net present value.
 
-- Table: Queries
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | year          | int     |
-- +---------------+---------+
-- (id, year) is the primary key (combination of columns with unique values) of this table.
-- The table has information about the id and the year of each inventory query.
-- Write a solution to find the npv of each query of the Queries table.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- NPV table:
-- +------+--------+--------+
-- | id   | year   | npv    |
-- +------+--------+--------+
-- | 1    | 2018   | 100    |
-- | 7    | 2020   | 30     |
-- | 13   | 2019   | 40     |
-- | 1    | 2019   | 113    |
-- | 2    | 2008   | 121    |
-- | 3    | 2009   | 12     |
-- | 11   | 2020   | 99     |
-- | 7    | 2019   | 0      |
-- +------+--------+--------+
-- Queries table:
-- +------+--------+
-- | id   | year   |
-- +------+--------+
-- | 1    | 2019   |
-- | 2    | 2008   |
-- | 3    | 2009   |
-- | 7    | 2018   |
-- | 7    | 2019   |
-- | 7    | 2020   |
-- | 13   | 2019   |
-- +------+--------+
-- Output: 
-- +------+--------+--------+
-- | id   | year   | npv    |
-- +------+--------+--------+
-- | 1    | 2019   | 113    |
-- | 2    | 2008   | 121    |
-- | 3    | 2009   | 12     |
-- | 7    | 2018   | 0      |
-- | 7    | 2019   | 0      |
-- | 7    | 2020   | 30     |
-- | 13   | 2019   | 40     |
-- +------+--------+--------+
-- Explanation: 
-- The npv value of (7, 2018) is not present in the NPV table, we consider it 0.
-- The npv values of all other queries can be found in the NPV table.

-- Create Table If Not Exists NPV (id int, year int, npv int)
-- Create Table If Not Exists Queries (id int, year int)
-- Truncate table NPV
-- insert into NPV (id, year, npv) values ('1', '2018', '100')
-- insert into NPV (id, year, npv) values ('7', '2020', '30')
-- insert into NPV (id, year, npv) values ('13', '2019', '40')
-- insert into NPV (id, year, npv) values ('1', '2019', '113')
-- insert into NPV (id, year, npv) values ('2', '2008', '121')
-- insert into NPV (id, year, npv) values ('3', '2009', '21')
-- insert into NPV (id, year, npv) values ('11', '2020', '99')
-- insert into NPV (id, year, npv) values ('7', '2019', '0')
-- Truncate table Queries
-- insert into Queries (id, year) values ('1', '2019')
-- insert into Queries (id, year) values ('2', '2008')
-- insert into Queries (id, year) values ('3', '2009')
-- insert into Queries (id, year) values ('7', '2018')
-- insert into Queries (id, year) values ('7', '2019')
-- insert into Queries (id, year) values ('7', '2020')
-- insert into Queries (id, year) values ('13', '2019')

SELECT
    q.id,
    q.year,
    IFNULL(n.npv,0) AS npv -- 净现值不在 NPV 表中, 我们把它看作是 0
FROM
    Queries AS q
LEFT JOIN
    NPV AS n
ON
    n.id = q.id AND 
    n.year = q.year