-- 3580. Find Consistently Improving Employees
-- Table: employees
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | employee_id | int     |
-- | name        | varchar |
-- +-------------+---------+
-- employee_id is the unique identifier for this table.
-- Each row contains information about an employee.
-- Table: performance_reviews
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | review_id   | int  |
-- | employee_id | int  |
-- | review_date | date |
-- | rating      | int  |
-- +-------------+------+
-- review_id is the unique identifier for this table.
-- Each row represents a performance review for an employee. The rating is on a scale of 1-5 where 5 is excellent and 1 is poor.
-- Write a solution to find employees who have consistently improved their performance over their last three reviews.

-- An employee must have at least 3 review to be considered
-- The employee's last 3 reviews must show strictly increasing ratings (each review better than the previous)
-- Use the most recent 3 reviews based on review_date for each employee
-- Calculate the improvement score as the difference between the latest rating and the earliest rating among the last 3 reviews
-- Return the result table ordered by improvement score in descending order, then by name in ascending order.

-- The result format is in the following example.

-- Example:
-- Input:
-- employees table:
-- +-------------+----------------+
-- | employee_id | name           |
-- +-------------+----------------+
-- | 1           | Alice Johnson  |
-- | 2           | Bob Smith      |
-- | 3           | Carol Davis    |
-- | 4           | David Wilson   |
-- | 5           | Emma Brown     |
-- +-------------+----------------+
-- performance_reviews table:
-- +-----------+-------------+-------------+--------+
-- | review_id | employee_id | review_date | rating |
-- +-----------+-------------+-------------+--------+
-- | 1         | 1           | 2023-01-15  | 2      |
-- | 2         | 1           | 2023-04-15  | 3      |
-- | 3         | 1           | 2023-07-15  | 4      |
-- | 4         | 1           | 2023-10-15  | 5      |
-- | 5         | 2           | 2023-02-01  | 3      |
-- | 6         | 2           | 2023-05-01  | 2      |
-- | 7         | 2           | 2023-08-01  | 4      |
-- | 8         | 2           | 2023-11-01  | 5      |
-- | 9         | 3           | 2023-03-10  | 1      |
-- | 10        | 3           | 2023-06-10  | 2      |
-- | 11        | 3           | 2023-09-10  | 3      |
-- | 12        | 3           | 2023-12-10  | 4      |
-- | 13        | 4           | 2023-01-20  | 4      |
-- | 14        | 4           | 2023-04-20  | 4      |
-- | 15        | 4           | 2023-07-20  | 4      |
-- | 16        | 5           | 2023-02-15  | 3      |
-- | 17        | 5           | 2023-05-15  | 2      |
-- +-----------+-------------+-------------+--------+
-- Output:
-- +-------------+----------------+-------------------+
-- | employee_id | name           | improvement_score |
-- +-------------+----------------+-------------------+
-- | 2           | Bob Smith      | 3                 |
-- | 1           | Alice Johnson  | 2                 |
-- | 3           | Carol Davis    | 2                 |
-- +-------------+----------------+-------------------+
-- Explanation:
-- Alice Johnson (employee_id = 1):
-- Has 4 reviews with ratings: 2, 3, 4, 5
-- Last 3 reviews (by date): 2023-04-15 (3), 2023-07-15 (4), 2023-10-15 (5)
-- Ratings are strictly increasing: 3 → 4 → 5
-- Improvement score: 5 - 3 = 2
-- Carol Davis (employee_id = 3):
-- Has 4 reviews with ratings: 1, 2, 3, 4
-- Last 3 reviews (by date): 2023-06-10 (2), 2023-09-10 (3), 2023-12-10 (4)
-- Ratings are strictly increasing: 2 → 3 → 4
-- Improvement score: 4 - 2 = 2
-- Bob Smith (employee_id = 2):
-- Has 4 reviews with ratings: 3, 2, 4, 5
-- Last 3 reviews (by date): 2023-05-01 (2), 2023-08-01 (4), 2023-11-01 (5)
-- Ratings are strictly increasing: 2 → 4 → 5
-- Improvement score: 5 - 2 = 3
-- Employees not included:
-- David Wilson (employee_id = 4): Last 3 reviews are all 4 (no improvement)
-- Emma Brown (employee_id = 5): Only has 2 reviews (needs at least 3)
-- The output table is ordered by improvement_score in descending order, then by name in ascending order.

-- CREATE TABLE employees (
--     employee_id INT,
--     name VARCHAR(255)
-- )
-- CREATE TABLE performance_reviews (
--     review_id INT,
--     employee_id INT,
--     review_date DATE,
--     rating INT
-- )
-- Truncate table employees
-- insert into employees (employee_id, name) values ('1', 'Alice Johnson')
-- insert into employees (employee_id, name) values ('2', 'Bob Smith')
-- insert into employees (employee_id, name) values ('3', 'Carol Davis')
-- insert into employees (employee_id, name) values ('4', 'David Wilson')
-- insert into employees (employee_id, name) values ('5', 'Emma Brown')
-- Truncate table performance_reviews
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('1', '1', '2023-01-15', '2')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('2', '1', '2023-04-15', '3')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('3', '1', '2023-07-15', '4')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('4', '1', '2023-10-15', '5')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('5', '2', '2023-02-01', '3')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('6', '2', '2023-05-01', '2')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('7', '2', '2023-08-01', '4')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('8', '2', '2023-11-01', '5')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('9', '3', '2023-03-10', '1')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('10', '3', '2023-06-10', '2')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('11', '3', '2023-09-10', '3')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('12', '3', '2023-12-10', '4')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('13', '4', '2023-01-20', '4')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('14', '4', '2023-04-20', '4')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('15', '4', '2023-07-20', '4')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('16', '5', '2023-02-15', '3')
-- insert into performance_reviews (review_id, employee_id, review_date, rating) values ('17', '5', '2023-05-15', '2')

-- Write your MySQL query statement below
WITH review_rankings AS (
    SELECT
        employee_id, 
        rating, 
        ROW_NUMBER() OVER(PARTITION BY employee_id ORDER BY review_date DESC) AS rn
    FROM 
        performance_reviews
), 
last_three_reviews AS (
    SELECT 
        *
    FROM 
        review_rankings
    WHERE 
        rn <= 3 -- 最近的 3 次评估
), 
improvements AS (
    SELECT 
        r1.employee_id, 
        (r2.rating - r3.rating) AS improvement_1, 
        (r1.rating - r3.rating) AS improvement_score -- 进步分数 为最后 3 次评估中最后一次评分与最早一次评分之间的差值
    FROM 
        last_three_reviews AS r1
    INNER JOIN 
        last_three_reviews AS r2 
    ON 
        r2.rn = r1.rn + 1 AND r1.employee_id = r2.employee_id
    INNER JOIN 
        last_three_reviews AS r3 
    ON 
        r3.rn = r1.rn + 2 AND r1.employee_id = r3.employee_id
)
SELECT 
    e.employee_id, 
    name, 
    improvement_score
FROM 
    improvements AS i
INNER JOIN 
    employees AS e 
ON 
    i.employee_id = e.employee_id
WHERE 
    improvement_score > improvement_1 AND -- 员工过去的 3 次评估，评分必须 严格递增（每次评价都比上一次好）
    improvement_1 > 0
ORDER BY 
    improvement_score DESC, name ASC -- 结果表以 进步分数 降序 排序，然后以 名字 升序 排序