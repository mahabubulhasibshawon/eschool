package database

// var courseList []Course

// type Course struct {
// 	ID          int      `json:"id"`
// 	Title       string   `json:"title"`
// 	Instructor  string   `json:"instructor"`  // Who teaches the course
// 	Description string   `json:"description"`
// 	Category    string   `json:"category"`    // e.g. "Programming", "Design"
// 	Price       float64  `json:"price"`
// 	Duration    string   `json:"duration"`    // e.g. "6h 30m" or "8 weeks"
// 	Level       string   `json:"level"`       // e.g. "Beginner", "Intermediate", "Advanced"
// 	Lessons     int      `json:"lessons"`     // Number of lessons or modules
// 	Thumbnail   string   `json:"thumbnail"`   // Cover image URL
// 	Tags        []string `json:"tags"`        // Extra keywords
// }

// // Create / Store a new course
// func StoreCourse(c Course) Course {
// 	c.ID = len(courseList) + 1
// 	courseList = append(courseList, c)
// 	return c
// }

// // List all courses
// func ListCourses() []Course {
// 	return courseList
// }

// // Get a course by ID
// func GetCourse(courseID int) *Course {
// 	for i := range courseList {
// 		if courseList[i].ID == courseID {
// 			return &courseList[i]
// 		}
// 	}
// 	return nil
// }

// // Update an existing course
// func UpdateCourse(course Course) {
// 	for idx := range courseList {
// 		if courseList[idx].ID == course.ID {
// 			courseList[idx] = course
// 		}
// 	}
// }

// // Delete a course by ID
// func DeleteCourse(courseID int) {
// 	var tempList []Course
// 	for _, c := range courseList {
// 		if c.ID != courseID {
// 			tempList = append(tempList, c)
// 		}
// 	}
// 	courseList = tempList
// }

// // Seed initial data
// func init() {
// 	c1 := Course{
// 		ID:          1,
// 		Title:       "Go Language Fundamentals",
// 		Instructor:  "Alice Rahman",
// 		Description: "Learn Go from scratch: syntax, control flow, and tooling.",
// 		Category:    "Programming",
// 		Price:       79.99,
// 		Duration:    "10h",
// 		Level:       "Beginner",
// 		Lessons:     45,
// 		Thumbnail:   "https://example.com/images/go-fundamentals.jpg",
// 		Tags:        []string{"golang", "backend", "basics"},
// 	}

// 	c2 := Course{
// 		ID:          2,
// 		Title:       "REST API Development with Go",
// 		Instructor:  "Mahabub Hasan",
// 		Description: "Build production-ready REST APIs using the Go standard library.",
// 		Category:    "Backend Development",
// 		Price:       129.00,
// 		Duration:    "14h",
// 		Level:       "Intermediate",
// 		Lessons:     60,
// 		Thumbnail:   "https://example.com/images/go-rest-api.jpg",
// 		Tags:        []string{"api", "web", "golang"},
// 	}

// 	c3 := Course{
// 		ID:          3,
// 		Title:       "Cloud Deployment with Docker & Kubernetes",
// 		Instructor:  "Farhana Karim",
// 		Description: "Containerize Go apps and deploy to Kubernetes clusters.",
// 		Category:    "DevOps",
// 		Price:       149.50,
// 		Duration:    "12h",
// 		Level:       "Advanced",
// 		Lessons:     55,
// 		Thumbnail:   "https://example.com/images/go-docker-k8s.jpg",
// 		Tags:        []string{"docker", "k8s", "deployment"},
// 	}

// 	courseList = append(courseList, c1, c2, c3)
// }
