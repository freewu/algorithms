-- 3421. Find Students Who Improved
-- Table: Scores
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | student_id  | int     |
-- | subject     | varchar |
-- | score       | int     |
-- | exam_date   | varchar |
-- +-------------+---------+
-- (student_id, subject, exam_date) is the primary key for this table.
-- Each row contains information about a student's score in a specific subject on a particular exam date. score is between 0 and 100 (inclusive).
-- Write a solution to find the students who have shown improvement. A student is considered to have shown improvement if they meet both of these conditions:

-- Have taken exams in the same subject on at least two different dates
-- Their latest score in that subject is higher than their first score
-- Return the result table ordered by student_id, subject in ascending order.

-- The result format is in the following example.

-- Example:
-- Input:
-- Scores table:
-- +------------+----------+-------+------------+
-- | student_id | subject  | score | exam_date  |
-- +------------+----------+-------+------------+
-- | 101        | Math     | 70    | 15-01-2023 |
-- | 101        | Math     | 85    | 15-02-2023 |
-- | 101        | Physics  | 65    | 15-01-2023 |
-- | 101        | Physics  | 60    | 15-02-2023 |
-- | 102        | Math     | 80    | 15-01-2023 |
-- | 102        | Math     | 85    | 15-02-2023 |
-- | 103        | Math     | 90    | 15-01-2023 |
-- | 104        | Physics  | 75    | 15-01-2023 |
-- | 104        | Physics  | 85    | 15-02-2023 |
-- +------------+----------+-------+------------+
-- Output:
-- +------------+----------+-------------+--------------+
-- | student_id | subject  | first_score | latest_score |
-- +------------+----------+-------------+--------------+
-- | 101        | Math     | 70          | 85           |
-- | 102        | Math     | 80          | 85           |
-- | 104        | Physics  | 75          | 85           |
-- +------------+----------+-------------+--------------+
-- Explanation:
-- Student 101 in Math: Improved from 70 to 85
-- Student 101 in Physics: No improvement (dropped from 65 to 60)
-- Student 102 in Math: Improved from 80 to 85
-- Student 103 in Math: Only one exam, not eligible
-- Student 104 in Physics: Improved from 75 to 85
-- Result table is ordered by student_id, subject.

-- CREATE TABLE Scores (
--     student_id INT,
--     subject VARCHAR(50),
--     score INT,
--     exam_date VARCHAR(10)
-- )
-- Truncate table Scores
-- insert into Scores (student_id, subject, score, exam_date) values ('101', 'Math', '70', '15-01-2023')
-- insert into Scores (student_id, subject, score, exam_date) values ('101', 'Math', '85', '15-02-2023')
-- insert into Scores (student_id, subject, score, exam_date) values ('101', 'Physics', '65', '15-01-2023')
-- insert into Scores (student_id, subject, score, exam_date) values ('101', 'Physics', '60', '15-02-2023')
-- insert into Scores (student_id, subject, score, exam_date) values ('102', 'Math', '80', '15-01-2023')
-- insert into Scores (student_id, subject, score, exam_date) values ('102', 'Math', '85', '15-02-2023')
-- insert into Scores (student_id, subject, score, exam_date) values ('103', 'Math', '90', '15-01-2023')
-- insert into Scores (student_id, subject, score, exam_date) values ('104', 'Physics', '75', '15-01-2023')
-- insert into Scores (student_id, subject, score, exam_date) values ('104', 'Physics', '85', '15-02-2023')

-- Write your MySQL query statement below

WITH first_exam AS (
    SELECT
        s.student_id AS student_id,
        s.subject AS subject,
        s.score AS score,
        s.exam_date AS exam_date
    FROM 
        Scores AS s
    RIGHT JOIN
        (
            SELECT 
                student_id,
                subject,
                MIN(exam_date) AS exam_date
            FROM
                Scores
            GROUP BY
                student_id, subject
        ) AS m 
    ON
        s.student_id = m.student_id AND
        s.subject = m.subject AND
        s.exam_date = m.exam_date 
), 
last_exam AS (
    SELECT
        s.student_id AS student_id,
        s.subject AS subject,
        s.score AS score,
        s.exam_date AS exam_date
    FROM 
        Scores AS s
    RIGHT JOIN
        (
            SELECT 
                student_id,
                subject,
                MAX(exam_date) AS exam_date
            FROM
                Scores
            GROUP BY
                student_id, subject
        ) AS m 
    ON
        s.student_id = m.student_id AND
        s.subject = m.subject AND
        s.exam_date = m.exam_date 
)

-- SELECT * FROM first_exam
-- | student_id | subject | score | exam_date  |
-- | ---------- | ------- | ----- | ---------- |
-- | 101        | Math    | 70    | 15-01-2023 |
-- | 101        | Physics | 65    | 15-01-2023 |
-- | 102        | Math    | 80    | 15-01-2023 |
-- | 103        | Math    | 90    | 15-01-2023 |
-- | 104        | Physics | 75    | 15-01-2023 |

-- SELECT * FROM last_exam
-- | student_id | subject | score | exam_date  |
-- | ---------- | ------- | ----- | ---------- |
-- | 101        | Math    | 85    | 15-02-2023 |
-- | 101        | Physics | 60    | 15-02-2023 |
-- | 102        | Math    | 85    | 15-02-2023 |
-- | 103        | Math    | 90    | 15-01-2023 |
-- | 104        | Physics | 85    | 15-02-2023 |

SELECT 
    f.student_id AS student_id,
    f.subject AS subject,
    f.score AS first_score,
    l.score AS latest_score
FROM 
    first_exam AS f
LEFT JOIN
    last_exam AS l
ON 
    f.student_id = l.student_id AND 
    f.subject = l.subject 
WHERE
    f.exam_date < l.exam_date AND
    f.score < l.score 
ORDER BY
    f.student_id, f.subject -- Return the result table ordered by student_id, subject in ascending order.