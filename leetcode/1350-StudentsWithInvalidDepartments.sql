-- 1350. Students With Invalid Departments
-- Table: Departments
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | name          | varchar |
-- +---------------+---------+
-- In SQL, id is the primary key of this table.
-- The table has information about the id of each department of a university.

-- Table: Students
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | name          | varchar |
-- | department_id | int     |
-- +---------------+---------+
-- In SQL, id is the primary key of this table.
-- The table has information about the id of each student at a university and the id of the department he/she studies at.
 
-- Find the id and the name of all students who are enrolled in departments that no longer exist.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Departments table:
-- +------+--------------------------+
-- | id   | name                     |
-- +------+--------------------------+
-- | 1    | Electrical Engineering   |
-- | 7    | Computer Engineering     |
-- | 13   | Bussiness Administration |
-- +------+--------------------------+
-- Students table:
-- +------+----------+---------------+
-- | id   | name     | department_id |
-- +------+----------+---------------+
-- | 23   | Alice    | 1             |
-- | 1    | Bob      | 7             |
-- | 5    | Jennifer | 13            |
-- | 2    | John     | 14            |
-- | 4    | Jasmine  | 77            |
-- | 3    | Steve    | 74            |
-- | 6    | Luis     | 1             |
-- | 8    | Jonathan | 7             |
-- | 7    | Daiana   | 33            |
-- | 11   | Madelynn | 1             |
-- +------+----------+---------------+
-- Output: 
-- +------+----------+
-- | id   | name     |
-- +------+----------+
-- | 2    | John     |
-- | 7    | Daiana   |
-- | 4    | Jasmine  |
-- | 3    | Steve    |
-- +------+----------+
-- Explanation: 
-- John, Daiana, Steve, and Jasmine are enrolled in departments 14, 33, 74, and 77 respectively. department 14, 33, 74, and 77 do not exist in the Departments table.

-- Create table If Not Exists Departments (id int, name varchar(30))
-- Create table If Not Exists Students (id int, name varchar(30), department_id int)
-- Truncate table Departments
-- insert into Departments (id, name) values ('1', 'Electrical Engineering')
-- insert into Departments (id, name) values ('7', 'Computer Engineering')
-- insert into Departments (id, name) values ('13', 'Bussiness Administration')
-- Truncate table Students
-- insert into Students (id, name, department_id) values ('23', 'Alice', '1')
-- insert into Students (id, name, department_id) values ('1', 'Bob', '7')
-- insert into Students (id, name, department_id) values ('5', 'Jennifer', '13')
-- insert into Students (id, name, department_id) values ('2', 'John', '14')
-- insert into Students (id, name, department_id) values ('4', 'Jasmine', '77')
-- insert into Students (id, name, department_id) values ('3', 'Steve', '74')
-- insert into Students (id, name, department_id) values ('6', 'Luis', '1')
-- insert into Students (id, name, department_id) values ('8', 'Jonathan', '7')
-- insert into Students (id, name, department_id) values ('7', 'Daiana', '33')
-- insert into Students (id, name, department_id) values ('11', 'Madelynn', '1')

-- Write your MySQL query statement below
SELECT 
    s.id AS id,
    s.name AS name
FROM
    Students AS s
LEFT JOIN
    Departments AS d 
ON 
    s.department_id = d.id
WHERE
    d.name IS NULL

SELECT 
    id,name
FROM 
    Students
WHERE 
    department_id NOT IN (SELECT id FROM Departments)
