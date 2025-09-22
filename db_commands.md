
# 📚 eSchool PostgreSQL Database Setup

This document outlines the steps taken to create and configure a PostgreSQL database for an e-learning platform named **eSchool**. The setup includes database creation, table creation, schema modifications, and basic queries.

---

## 🛠️ 1. Create the Database

```sql
-- Log into PostgreSQL and create a new database named "eschool"
CREATE DATABASE eschool;
````

---

## 🔁 2. Connect to the Database

> 💡 Note: The following was incorrectly attempted inside the PostgreSQL shell.
> Commands like `psql -U postgres -d eschool` must be run **from the terminal**, not within `psql`.

Incorrect attempts inside `psql` shell:

```sql
psql -U postgres -d eschool;
```

✅ Correct way to connect **from the terminal (bash)**:

```bash
psql -U postgres -d eschool
```

---

## 📋 3. Create the `courses` Table

```sql
CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    instructor TEXT NOT NULL,
    description TEXT,
    category TEXT,
    price NUMERIC(10,2),       -- e.g., 79.99
    duration VARCHAR(50),      -- e.g., "10h"
    level VARCHAR(50),         -- e.g., Beginner, Intermediate
    lessons INT,               -- Number of lessons
    thumbnail TEXT,            -- URL to thumbnail image
    tags TEXT[]                -- Array of tags: {"golang", "backend", "basics"}
);
```

---

## 🔍 4. Check if Table Exists and View It

List all tables:

```sql
\dt
```

Check table structure:

```sql
\d
```

Select all rows from the `courses` table (currently empty):

```sql
SELECT * FROM courses;
```

---

## 🔧 5. Alter the Table: Remove Unneeded Columns

In this step, several columns were dropped from the `courses` table for simplification:

```sql
ALTER TABLE courses
DROP COLUMN price,
DROP COLUMN duration,
DROP COLUMN level,
DROP COLUMN lessons,
DROP COLUMN thumbnail,
DROP COLUMN tags;
```

---

## ✅ 6. Final Table Structure

Check updated structure:

```sql
\dt
\d courses
```

Resulting columns:

* `id` (serial, primary key)
* `title` (text, not null)
* `instructor` (text, not null)
* `description` (text)
* `category` (text)

---

## 🔍 7. Query Final Table

```sql
SELECT * FROM courses;
```

Output:

```
 id | title | instructor | description | category 
----+-------+------------+-------------+----------
(0 rows)
```

---

## 📌 Notes

* Always exit the PostgreSQL shell using `\q`, not commands like `exit` or `/q`.
* Use proper CLI for `psql` commands, and SQL syntax **only inside the `psql` shell**.

---

## 🚀 Next Steps

* Insert sample data into the `courses` table.
* Create additional tables like `students`, `enrollments`, etc.
* Define relationships between tables (e.g., `foreign keys`).
* Implement data validation or constraints as needed.

```
