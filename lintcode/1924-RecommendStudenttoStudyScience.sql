-- 1924 · Recommend Student to Study Science
-- # Description
-- If a student scores over 90 in mathematics or over 95 in science , then that student is recommended to study Science.
-- Please write a SQL query to output the names, mathematics scores and science comprehensive scores of all suitable science students in the table.

-- cst (comprehensive science test)
-- Table definition: student_scores

-- columns_name	type	explaination
-- name	varchar	student's name
-- chinese	int	student's Chinese score
-- math	int	student's math score
-- english	int	student's English score
-- cst	int	student's cst score
-- Example
-- Example 1:

-- Table content: students

-- name	chinese	math	english	cst
-- KangKang	95	91	89	97
-- Jane	90	93	98	98
-- Micheal	85	76	93	92
-- Maria	88	89	95	94
-- LiHua	30	13	19	23
-- According to the above table, we should output:

-- name	math	cst
-- KangKang	91	97
-- Jane	93	98
-- Example 2:

-- Table content: students

-- name	chinese	math	english	cst
-- KangKang	95	90	89	97
-- Jane	90	90	98	95
-- According to the above table, we should output:

-- name	math	cst
-- KangKang	90	97

SELECT 
	name,
	math,
	cst
FROM 
	student_scores
WHERE
	math > 90 OR
	cst > 95