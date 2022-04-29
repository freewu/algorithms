-- 1993 · Delete the course called Dynamic Programming
-- # Write an SQL statement to delete the information of the course named Dynamic Programming from the course table courses.

-- Table definition: courses

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	number of students
-- created_at	date	start time
-- teacher_id	int	teacher id



-- Example
-- Table Contents: courses

-- | id | name | student_count | created_at | teacher_id |
-- | ---- | ---- | ---- | ---- | ---- | ---- | ----
-- | 1 | Senior Algorithm | 880 | 2020-06-01 | 4 |
-- | 2 | System Design | 1350 | 2020-07-18 | 3 |
-- | 3 | Django | 780 | 2020-02-29 | 3 |
-- | 4 | Web | 340 | 2020-04-22 | 4 |
-- | 5 | Big Data | 700 | 2020-09-11 | 1 |
-- | 6 | Artificial Intelligence | 1660 | 2018-05-13 | 3 |
-- | 7 | Java P6+ | 780 | 2019-01-19 | 3 |
-- | 8 | Data Analysis | 500 | 2019-07-12 | 1 |
-- | 10 | Object Oriented Design | 300 | 2020-08-08 | 4 |
-- | 12 | Dynamic Programming | 2000 | 2018-08-18 | 1 |

-- After the execution of the DELETE code, we will execute SELECT * FROM courses and the table should return.

-- | id | name | student_count | created_at | teacher_id |
-- | ---- | ---- | ---- | ---- | ---- | ---- | ----
-- | 1 | Senior Algorithm | 880 | 2020-06-01 | 4 |
-- | 2 | System Design | 1350 | 2020-07-18 | 3 |
-- | 3 | Django | 780 | 2020-02-29 | 3 |
-- | 4 | Web | 340 | 2020-04-22 | 4 |
-- | 5 | Big Data | 700 | 2020-09-11 | 1 |
-- | 6 | Artificial Intelligence | 1660 | 2018-05-13 | 3 |
-- | 7 | Java P6+ | 780 | 2019-01-19 | 3 |
-- | 8 | Data Analysis | 500 | 2019-07-12 | 1 |
-- | 10 | Object Oriented Design | 300 | 2020-08-08 | 4 |

DELETE FROM 
	courses
WHERE 
	name = 'Dynamic Programming'