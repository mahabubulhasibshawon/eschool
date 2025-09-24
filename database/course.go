package database

import "time"

// Course struct (updated)
type Course struct {
	ID           int       `db:"id" json:"id"`
	Title        string    `db:"title" json:"title"`
	InstructorID int       `db:"instructor_id" json:"instructor_id"`
	Description  string    `db:"description" json:"description"`
	Category     string    `db:"category" json:"category"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
