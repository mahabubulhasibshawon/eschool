postgres=# create database eschool;
CREATE DATABASE
postgres=# psql -U postgres -d eschool 
postgres-# ;
ERROR:  syntax error at or near "psql"
LINE 1: psql -U postgres -d eschool
        ^
postgres=# psql -U postgres
postgres-# \l
postgres-# psql -U postgres -d eschool
postgres-# ;
ERROR:  syntax error at or near "psql"
LINE 1: psql -U postgres
        ^
postgres=# psql -U postgres -d eschool;
ERROR:  syntax error at or near "psql"
LINE 1: psql -U postgres -d eschool;
        ^
postgres=# /q
postgres-# exit
Use \q to quit.
postgres-# \q
postgres@TN-AIO-0080:~$ psql -U postgres -d eschool;
psql (16.10 (Ubuntu 16.10-0ubuntu0.24.04.1))
Type "help" for help.

eschool=# CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    instructor TEXT NOT NULL,
    description TEXT,
    category TEXT,
    price NUMERIC(10,2),          -- stores values like 79.99
    duration VARCHAR(50),         -- "10h" fits fine
    level VARCHAR(50),            -- e.g. Beginner, Intermediate
    lessons INT,                  -- number of lessons
    thumbnail TEXT,               -- URL to image
    tags TEXT[]                   -- PostgreSQL array for tags like {"golang","backend","basics"}
);
CREATE TABLE
eschool=# select * from courses
eschool-# ;
 id | title | instructor | description | category | price | duration | level | lessons | thumbnail | tags 
----+-------+------------+-------------+----------+-------+----------+-------+---------+-----------+------
(0 rows)

eschool=# \d
               List of relations
 Schema |      Name      |   Type   |  Owner   
--------+----------------+----------+----------
 public | courses        | table    | postgres
 public | courses_id_seq | sequence | postgres
(2 rows)

eschool=# ALTER TABLE courses
DROP COLUMN price,
DROP COLUMN duration,
DROP COLUMN level,
DROP COLUMN lessons,
DROP COLUMN thumbnail,
DROP COLUMN tags;
ALTER TABLE
eschool=# \dt
          List of relations
 Schema |  Name   | Type  |  Owner   
--------+---------+-------+----------
 public | courses | table | postgres
(1 row)

eschool=# select * from courses;
 id | title | instructor | description | category 
----+-------+------------+-------------+----------
(0 rows)

eschool=# 