-- 577. Employee Bonus
-- Table: Employee
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | empId       | int     |
-- | name        | varchar |
-- | supervisor  | int     |
-- | salary      | int     |
-- +-------------+---------+
-- empId is the primary key column for this table.
-- Each row of this table indicates the name and the ID of an employee in addition to their salary and the id of their manager.
--  
-- Table: Bonus
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | empId       | int  |
-- | bonus       | int  |
-- +-------------+------+
-- empId is the primary key column for this table.
-- empId is a foreign key to empId from the Employee table.
-- Each row of this table contains the id of an employee and their respective bonus.
--  
-- Write an SQL query to report the name and bonus amount of each employee with a bonus less than 1000.
-- Return the result table in any order.
-- The query result format is in the following example.
--
-- Example 1:
-- Input:
-- Employee table:
-- +-------+--------+------------+--------+
-- | empId | name   | supervisor | salary |
-- +-------+--------+------------+--------+
-- | 3     | Brad   | null       | 4000   |
-- | 1     | John   | 3          | 1000   |
-- | 2     | Dan    | 3          | 2000   |
-- | 4     | Thomas | 3          | 4000   |
-- +-------+--------+------------+--------+
-- Bonus table:
-- +-------+-------+
-- | empId | bonus |
-- +-------+-------+
-- | 2     | 500   |
-- | 4     | 2000  |
-- +-------+-------+
-- Output:
-- +------+-------+
-- | name | bonus |
-- +------+-------+
-- | Brad | null  |
-- | John | null  |
-- | Dan  | 500   |
-- +------+-------+

-- Create table If Not Exists Employee (empId int, name varchar(255), supervisor int, salary int)
-- Create table If Not Exists Bonus (empId int, bonus int)
-- Truncate table Employee
-- insert into Employee (empId, name, supervisor, salary) values ('3', 'Brad', 'None', '4000')
-- insert into Employee (empId, name, supervisor, salary) values ('1', 'John', '3', '1000')
-- insert into Employee (empId, name, supervisor, salary) values ('2', 'Dan', '3', '2000')
-- insert into Employee (empId, name, supervisor, salary) values ('4', 'Thomas', '3', '4000')
-- Truncate table Bonus
-- insert into Bonus (empId, bonus) values ('2', '500')
-- insert into Bonus (empId, bonus) values ('4', '2000')

-- Write your MySQL query statement below
SELECT
    e.name,
    b.bonus
FROM
    Employee AS e
LEFT JOIN
    Bonus AS b
ON
    e.empId = b.empId
WHERE
    b.bonus < 1000 OR
    b.bonus IS NULL