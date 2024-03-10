-- 2072. The Winner University
-- Table: NewYork
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | student_id  | int  |
-- | score       | int  |
-- +-------------+------+
-- In SQL, student_id is the primary key for this table.
-- Each row contains information about the score of one student from New York University in an exam.
 
-- Table: California
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | student_id  | int  |
-- | score       | int  |
-- +-------------+------+
-- In SQL, student_id is the primary key for this table.
-- Each row contains information about the score of one student from California University in an exam.
-- There is a competition between New York University and California University. 
-- The competition is held between the same number of students from both universities. 
-- The university that has more excellent students wins the competition. 
-- If the two universities have the same number of excellent students, the competition ends in a draw.
-- An excellent student is a student that scored 90% or more in the exam.

-- Return:
--     "New York University" if New York University wins the competition.
--     "California University" if California University wins the competition.
--     "No Winner" if the competition ends in a draw.
--     The result format is in the following example.

-- Example 1:
-- Input: 
-- NewYork table:
-- +------------+-------+
-- | student_id | score |
-- +------------+-------+
-- | 1          | 90    |
-- | 2          | 87    |
-- +------------+-------+
-- California table:
-- +------------+-------+
-- | student_id | score |
-- +------------+-------+
-- | 2          | 89    |
-- | 3          | 88    |
-- +------------+-------+
-- Output: 
-- +---------------------+
-- | winner              |
-- +---------------------+
-- | New York University |
-- +---------------------+
-- Explanation:
-- New York University has 1 excellent student, and California University has 0 excellent students.

-- Example 2:
-- Input: 
-- NewYork table:
-- +------------+-------+
-- | student_id | score |
-- +------------+-------+
-- | 1          | 89    |
-- | 2          | 88    |
-- +------------+-------+
-- California table:
-- +------------+-------+
-- | student_id | score |
-- +------------+-------+
-- | 2          | 90    |
-- | 3          | 87    |
-- +------------+-------+
-- Output: 
-- +-----------------------+
-- | winner                |
-- +-----------------------+
-- | California University |
-- +-----------------------+
-- Explanation:
-- New York University has 0 excellent students, and California University has 1 excellent student.

-- Example 3:
-- Input: 
-- NewYork table:
-- +------------+-------+
-- | student_id | score |
-- +------------+-------+
-- | 1          | 89    |
-- | 2          | 90    |
-- +------------+-------+
-- California table:
-- +------------+-------+
-- | student_id | score |
-- +------------+-------+
-- | 2          | 87    |
-- | 3          | 99    |
-- +------------+-------+
-- Output: 
-- +-----------+
-- | winner    |
-- +-----------+
-- | No Winner |
-- +-----------+
-- Explanation:
-- Both New York University and California University have 1 excellent student.

-- Create table If Not Exists NewYork (student_id int, score int)
-- Create table If Not Exists California (student_id int, score int)
-- Truncate table NewYork
-- insert into NewYork (student_id, score) values ('1', '90')
-- insert into NewYork (student_id, score) values ('2', '87')
-- Truncate table California
-- insert into California (student_id, score) values ('2', '89')
-- insert into California (student_id, score) values ('3', '88')

-- Write your MySQL query statement below
-- IF
SELECT 
    IF(a.num = b. num, "No Winner",
        IF(a.num > b. num, "New York University","California University")
    ) AS winner
FROM 
(
    SELECT
        COUNT(DISTINCT student_id) AS num
    FROM
        NewYork
    WHERE
        score >= 90
) AS a,
(
    SELECT
        COUNT(DISTINCT student_id) AS num
    FROM
        California
    WHERE
        score >= 90
) AS b

-- CASE WHEN
SELECT
    CASE
        WHEN (SELECT count(*) FROM NewYork WHERE score >= 90) > (SELECT count(*) FROM California WHERE score >= 90) THEN 'New York University'
        WHEN (SELECT count(*) FROM California WHERE >= 90)> (SELECT count(*) FROM NewYork WHERE score >= 90)  THEN 'California University'
        ELSE 'No Winner'
    END AS winner

SELECT 
    CASE
        WHEN n > c THEN 'New York University'
        WHEN n < c THEN 'California University'
        ELSE 'No Winner'
    END AS winner 
FROM
    (SELECT COUNT(*) AS n FROM NewYork WHERE score >= 90) AS a,
    (SELECT COUNT(*) AS c FROM California WHERE score >= 90) AS b

-- with + case then
WITH 
    a AS (SELECT SUM(IF(score >= 90, 1, 0)) cnt_ny FROM NewYork), 
    b AS (SELECT SUM(IF(score >= 90, 1, 0)) cnt_ca FROM California)

SELECT 
    CASE 
        WHEN a.cnt_ny > b.cnt_ca THEN 'New York University'
        WHEN a.cnt_ny < b.cnt_ca THEN 'California University'
        WHEN a.cnt_ny = b.cnt_ca THEN 'No Winner'
    END AS winner
FROM 
    a, b 