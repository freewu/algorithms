-- 1412. Find the Quiet Students in All Exams
-- Table: Student
-- +---------------------+---------+
-- | Column Name         | Type    |
-- +---------------------+---------+
-- | student_id          | int     |
-- | student_name        | varchar |
-- +---------------------+---------+
-- student_id is the primary key (column with unique values) for this table.
-- student_name is the name of the student.
 
-- Table: Exam
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | exam_id       | int     |
-- | student_id    | int     |
-- | score         | int     |
-- +---------------+---------+
-- (exam_id, student_id) is the primary key (combination of columns with unique values) for this table.
-- Each row of this table indicates that the student with student_id had a score points in the exam with id exam_id.
 
-- A quiet student is the one who took at least one exam and did not score the highest or the lowest score.
-- Write a solution to report the students (student_id, student_name) being quiet in all exams. 
-- Do not return the student who has never taken any exam.
-- Return the result table ordered by student_id.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Student table:
-- +-------------+---------------+
-- | student_id  | student_name  |
-- +-------------+---------------+
-- | 1           | Daniel        |
-- | 2           | Jade          |
-- | 3           | Stella        |
-- | 4           | Jonathan      |
-- | 5           | Will          |
-- +-------------+---------------+
-- Exam table:
-- +------------+--------------+-----------+
-- | exam_id    | student_id   | score     |
-- +------------+--------------+-----------+
-- | 10         |     1        |    70     |
-- | 10         |     2        |    80     |
-- | 10         |     3        |    90     |
-- | 20         |     1        |    80     |
-- | 30         |     1        |    70     |
-- | 30         |     3        |    80     |
-- | 30         |     4        |    90     |
-- | 40         |     1        |    60     |
-- | 40         |     2        |    70     |
-- | 40         |     4        |    80     |
-- +------------+--------------+-----------+
-- Output: 
-- +-------------+---------------+
-- | student_id  | student_name  |
-- +-------------+---------------+
-- | 2           | Jade          |
-- +-------------+---------------+
-- Explanation: 
-- For exam 1: Student 1 and 3 hold the lowest and high scores respectively.
-- For exam 2: Student 1 hold both highest and lowest score.
-- For exam 3 and 4: Studnet 1 and 4 hold the lowest and high scores respectively.
-- Student 2 and 5 have never got the highest or lowest in any of the exams.
-- Since student 5 is not taking any exam, he is excluded from the result.
-- So, we only return the information of Student 2.

-- Create table If Not Exists Student (student_id int, student_name varchar(30))
-- Create table If Not Exists Exam (exam_id int, student_id int, score int)
-- Truncate table Student
-- insert into Student (student_id, student_name) values ('1', 'Daniel')
-- insert into Student (student_id, student_name) values ('2', 'Jade')
-- insert into Student (student_id, student_name) values ('3', 'Stella')
-- insert into Student (student_id, student_name) values ('4', 'Jonathan')
-- insert into Student (student_id, student_name) values ('5', 'Will')
-- Truncate table Exam
-- insert into Exam (exam_id, student_id, score) values ('10', '1', '70')
-- insert into Exam (exam_id, student_id, score) values ('10', '2', '80')
-- insert into Exam (exam_id, student_id, score) values ('10', '3', '90')
-- insert into Exam (exam_id, student_id, score) values ('20', '1', '80')
-- insert into Exam (exam_id, student_id, score) values ('30', '1', '70')
-- insert into Exam (exam_id, student_id, score) values ('30', '3', '80')
-- insert into Exam (exam_id, student_id, score) values ('30', '4', '90')
-- insert into Exam (exam_id, student_id, score) values ('40', '1', '60')
-- insert into Exam (exam_id, student_id, score) values ('40', '2', '70')
-- insert into Exam (exam_id, student_id, score) values ('40', '4', '80')

-- 取每门考试的最高分和最低分
-- SELECT
--     MAX(score) AS max_score,
--     MIN(score) AS min_score,
--     exam_id    
-- FROM
--     Exam 
-- GROUP BY
--     exam_id

-- max & min
SELECT
    student_id,
    student_name
FROM
    Student 
WHERE
    student_id IN 
    ( -- 都参加了考试的学生
        SELECT DISTINCT student_id FROM Exam 
    ) AND 
    student_id NOT IN 
    (-- 获得最高分或最低分的同学
        SELECT 
           DISTINCT e.student_id
        FROM
            Exam  AS e,
            (
                SELECT
                    MAX(score) AS max_score,
                    MIN(score) AS min_score,
                    exam_id    
                FROM
                    Exam 
                GROUP BY
                    exam_id
            ) AS r
        WHERE
            e.exam_id = r.exam_id  AND 
            (
                e.score = r.max_score OR
                e.score = r.min_score
            )
    )

-- RANK
SELECT 
    *
FROM 
    student
WHERE 
    student_id IN ( SELECT DISTINCT(student_id) FROM exam )AND 
    student_id NOT IN (
        SELECT 
            student_id
        FROM (
            SELECT 
                student_id,
                RANK() OVER(PARTITION BY exam_id ORDER BY score DESC) AS high, -- 每门考试按分数由高到低排名
                RANK() OVER(PARTITION BY exam_id ORDER BY score ASC)  AS low   -- 每门考试按分数由低到高排名
            FROM 
                exam
        ) AS a
        WHERE -- 考试排名不为最高或最低
            high = 1 OR 
            low = 1
    )