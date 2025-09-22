package database



type Course struct {
	ID          int      `db:"id" json:"id"`
	Title       string   `db:"title" json:"title"`
	Instructor  string   `db:"instructor" json:"instructor"`
	Description string   `db:"description" json:"description"`
	Category    string   `db:"category" json:"category"`
}

// StoreCourse inserts a new course into the DB
// func StoreCourse(db *sqlx.DB, c Course) (Course, error) {
// 	query := `
//         INSERT INTO courses
//         (title, instructor, description, category, price, duration, level, lessons, thumbnail, tags)
//         VALUES
//         (:title, :instructor, :description, :category, :price, :duration, :level, :lessons, :thumbnail, :tags)
//         RETURNING id
//     `

// 	// NamedQuery binds struct fields to query params
// 	var id int
// 	err := db.QueryRowx(query, c).Scan(&id)
// 	if err != nil {
// 		return Course{}, err
// 	}
// 	c.ID = id
// 	return c, nil
// }
