package decryptor

type Course struct {
	Title   string
	Modules []Module
}

type CourseRepository interface {
	FindAll() ([]Course, error)
}
