-- 2669. Count Artist Occurrences On Spotify Ranking List
-- Table: Spotify
-- +-------------+---------+ 
-- | Column Name | Type    | 
-- +-------------+---------+ 
-- | id          | int     | 
-- | track_name  | varchar |
-- | artist      | varchar |
-- +-------------+---------+
-- id is the primary key (column with unique values) for this table.
-- Each row contains an id, track_name, and artist.
-- Write a solution to find how many times each artist appeared on the Spotify ranking list.

-- Return the result table having the artist's name along with the corresponding number of occurrences ordered by occurrence count in descending order. 
-- If the occurrences are equal, then it’s ordered by the artist’s name in ascending order.
-- The result format is in the following example​​​​​.

-- Example 1:
-- Input:
-- Spotify table: 
-- +---------+--------------------+------------+ 
-- | id      | track_name         | artist     |  
-- +---------+--------------------+------------+
-- | 303651  | Heart Won't Forget | Sia        |
-- | 1046089 | Shape of you       | Ed Sheeran |
-- | 33445   | I'm the one        | DJ Khalid  |
-- | 811266  | Young Dumb & Broke | DJ Khalid  | 
-- | 505727  | Happier            | Ed Sheeran |
-- +---------+--------------------+------------+ 
-- Output:
-- +------------+-------------+
-- | artist     | occurrences | 
-- +------------+-------------+
-- | DJ Khalid  | 2           |
-- | Ed Sheeran | 2           |
-- | Sia        | 1           | 
-- +------------+-------------+ 
-- Explanation: The count of occurrences is listed in descending order under the column name "occurrences". 
-- If the number of occurrences is the same, the artist's names are sorted in ascending order.
-- Create table If Not Exists Spotify (id int,track_name varchar(100),artist varchar(100))
-- Truncate table Spotify
-- insert into Spotify (id, track_name, artist) values ('303651', 'Heart Won't Forget', 'Ed Sheeran')
-- insert into Spotify (id, track_name, artist) values ('1046089', 'Shape of you', 'Sia')
-- insert into Spotify (id, track_name, artist) values ('33445', 'I'm the one', 'DJ Khalid')
-- insert into Spotify (id, track_name, artist) values ('811266', 'Young Dumb & Broke', 'DJ Khalid')
-- insert into Spotify (id, track_name, artist) values ('505727', 'Happier', 'Ed Sheeran')

-- Write your MySQL query statement below
SELECT
    artist,
    COUNT(*) AS occurrences
FROM
    Spotify 
GROUP BY
    artist
ORDER BY
    occurrences DESC, -- occurrences ordered by occurrence count in descending order
    artist ASC -- If the occurrences are equal, then it’s ordered by the artist’s name in ascending order.