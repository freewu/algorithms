-- 2092 · Search for a lecturer number
-- # Description
-- Please write an SQL statement to query the different teacher id teacher_id from the courses table.
-- Table definition : courses (course table)

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id

-- Note that the query returns a non-repeating teacher number.
-- If the query returns no results, it should return null.
-- Example
-- Example 1:

-- Table Contents : courses

-- | id   | name                    | student_count | created_at | teacher_id |
-- | ---- | ----------------------- | ------------- | ---------- | ---------- |
-- | 1    | Advanced Algorithms     | 880           | 2020/6/1   | 4          |
-- | 2    | System Design           | 1350          | 2020/7/18  | 3          |
-- | 3    | Django                  | 780           | 2020/2/29  | 3          |
-- | 4    | Web                     | 340           | 2020/4/22  | 4          |
-- | 5    | Big Data                | 700           | 2020/9/11  | 1          |
-- | 6    | Artificial Intelligence | 1660          | 2018/5/13  | 3          |

-- After running your SQL statement, the table should return.

-- teacher_id
-- 4
-- 3
-- 1
-- Example 2:

-- Table Contents : courses

-- | id   | name                    | student_count | created_at | teacher_id |
-- | ---- | ----------------------- | ------------- | ---------- | ---------- |
-- | 1    | Advanced Algorithms     | 880           | 2020/6/1   | NULL       |
-- | 2    | System Design           | 1350          | 2020/7/18  | NULL       |
-- | 3    | Django                  | 880           | 2020/2/29  | NULL       |
-- | 4    | Web                     | 340           | 2020/4/22  | NULL       |
-- | 5    | Big Data                | 780           | 2020/9/11  | NULL       |

-- After running your SQL statement, the table should return.

-- teacher_id
-- The teacher number teacher_id in example 2 is empty, so the output contains only the table header and no data.
SELECT
	DISTINCT(teacher_id) AS teacher_id 
FROM
	courses