package decryptor

// Module represents a video course module
type Module struct {
	Order  int
	Title  string
	Clips  []Clip
	Course *Course
}
