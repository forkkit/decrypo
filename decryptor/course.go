package decryptor

// Course represents a video course
type Course struct {
	Title   string
	Modules []Module
}

// CourseRepository defined an interface for fetching courses
type CourseRepository interface {
	FindAll() ([]Course, error)
}
