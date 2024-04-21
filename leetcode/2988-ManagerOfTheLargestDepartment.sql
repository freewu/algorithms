-- 2988. Manager of the Largest Department
-- Table: Employees
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | emp_id      | int     |
-- | emp_name    | varchar |
-- | dep_id      | int     |
-- | position    | varchar |
-- +-------------+---------+
-- emp_id is column of unique values for this table.
-- This table contains emp_id, emp_name, dep_id, and position.
-- Write a solution to find the name of the manager from the largest department. There may be multiple largest departments when the number of employees in those departments is the same.

-- Return the result table sorted by dep_id in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Employees table:
-- +--------+----------+--------+---------------+
-- | emp_id | emp_name | dep_id | position      | 
-- +--------+----------+--------+---------------+
-- | 156    | Michael  | 107    | Manager       |
-- | 112    | Lucas    | 107    | Consultant    |    
-- | 8      | Isabella | 101    | Manager       | 
-- | 160    | Joseph   | 100    | Manager       | 
-- | 80     | Aiden    | 100    | Engineer      | 
-- | 190    | Skylar   | 100    | Freelancer    | 
-- | 196    | Stella   | 101    | Coordinator   |
-- | 167    | Audrey   | 100    | Consultant    |
-- | 97     | Nathan   | 101    | Supervisor    |
-- | 128    | Ian      | 101    | Administrator |
-- | 81     | Ethan    | 107    | Administrator |
-- +--------+----------+--------+---------------+
-- Output
-- +--------------+--------+
-- | manager_name | dep_id | 
-- +--------------+--------+
-- | Joseph       | 100    | 
-- | Isabella     | 101    | 
-- +--------------+--------+
-- Explanation
-- - Departments with IDs 100 and 101 each has a total of 4 employees, while department 107 has 3 employees. Since both departments 100 and 101 have an equal number of employees, their respective managers will be included.
-- Output table is ordered by dep_id in ascending order.

-- Create table if not exists Employees ( emp_id int, emp_name varchar(50), dep_id int, position varchar(30))
-- Truncate table Employees
-- insert into Employees (emp_id, emp_name, dep_id, position) values ('156', 'Michael', '107', 'Manager')
-- insert into Employees (emp_id, emp_name, dep_id, position) values ('112', 'Lucas', '107', 'Consultant')
-- insert into Employees (emp_id, emp_name, dep_id, position) values ('8', 'Isabella', '101', 'Manager')
-- insert into Employees (emp_id, emp_name, dep_id, position) values ('160', 'Joseph', '100', 'Manager')
-- insert into Employees (emp_id, emp_name, dep_id, position) values ('80', 'Aiden', '100', 'Engineer')
-- insert into Employees (emp_id, emp_name, dep_id, position) values ('190', 'Skylar', '100', 'Freelancer')
-- insert into Employees (emp_id, emp_name, dep_id, position) values ('196', 'Stella', '101', 'Coordinator')
-- insert into Employees (emp_id, emp_name, dep_id, position) values ('167', 'Audrey', '100', 'Consultant')
-- insert into Employees (emp_id, emp_name, dep_id, position) values ('97', 'Nathan', '101', 'Supervisor')
-- insert into Employees (emp_id, emp_name, dep_id, position) values ('128', 'Ian', '101', 'Administrator')
-- insert into Employees (emp_id, emp_name, dep_id, position) values ('81', 'Ethan', '107', 'Administrator')

-- Write your MySQL query statement below
SELECT 
    m.manager_name,
    d.dep_id
FROM
(
    SELECT 
        COUNT(*) AS cnt,
        RANK() OVER(ORDER BY COUNT(*) DESC) AS rk,
        dep_id
    FROM 
        Employees 
    GROUP BY 
        dep_id
) AS d 
LEFT JOIN 
(
    SELECT 
        dep_id,
        emp_name AS manager_name 
    FROM 
        Employees 
    WHERE
        position = "Manager"
) AS m
ON
    d.dep_id = m.dep_id
WHERE
    d.rk = 1
ORDER BY
    d.dep_id
