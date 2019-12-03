package main

import (
	"flag"

	"github.com/ajdnik/decrypo/decryptor"
	"github.com/ajdnik/decrypo/file"
	"github.com/ajdnik/decrypo/pluralsight"
)

func main() {
	defClip, err := pluralsight.GetClipPath()
	if err != nil {
		panic(err)
	}
	defDb, err := pluralsight.GetDbPath()
	if err != nil {
		panic(err)
	}
	clips := flag.String("clips", defClip, "location of clip .psv files")
	db := flag.String("db", defDb, "location of sqlite file")
	output := flag.String("output", "./Pluralsight Courses/", "location of decrypted courses")
	flag.Parse()

	svc := decryptor.Service{
		Decoder: &pluralsight.Decoder{},
		Storage: &file.Storage{
			Path: *output,
		},
		Clips: &pluralsight.ClipRepository{
			Path: *clips,
		},
		Courses: &pluralsight.CourseRepository{
			Path: *db,
		},
	}
	err = svc.DecryptAll()
	if err != nil {
		panic(err)
	}
}
