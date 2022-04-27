-- 1919 · Members Whose Scores Exceed the Group leader
-- The team table contains all the group members' name (name), their group leader (group_leader) also belongs to the group members, each group member has an id, and there is also a list of the id of the group leader corresponding to the group members.

-- Table definition: group_members

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	group member's name
-- score	int	group member's score
-- group_leader_id	int	group leader id
-- Given the group_members table, write a SQL query that can get the names of group members whose score exceeds their group leader (group_leader).

-- Example
-- Eample 1:

-- Table content: group_members

-- id	name	score	group_leader_id
-- 1	Bryant	81	4
-- 2	Iverson	60	1
-- 3	Carter	51	null
-- 4	McGrady	62	null
-- In the above table, Bryant is the only group member whose score exceeds his group leader.

-- name
-- Bryant
-- Eample 2:

-- Table content: group_members

-- id	name	score	group_leader_id
-- 1	Bryant	81	null
-- 2	Iverson	60	3
-- 3	Carter	51	1
-- In the above table, Iverson is the only group member whose score exceeds his group leader.

-- name
-- Iverson
SELECT 
	a.name AS name
FROM 
	group_members AS a
WHERE
	a.group_leader_id IS NOT NULL AND
	a.score > (
		SELECT
			b.score
		FROM
			group_members AS b
		WHERE
			b.id = a.group_leader_id
	)