package decryptor

type OnDecrypted func(Clip, *string)

type Service struct {
	Decoder Decoder
	Storage Storage
	Courses CourseRepository
	Clips   ClipRepository
}

func (s *Service) DecryptAll(evt OnDecrypted) error {
	courses, err := s.Courses.FindAll()
	if err != nil {
		return err
	}
	for _, course := range courses {
		for _, module := range course.Modules {
			for _, clip := range module.Clips {
				if !s.Clips.ExistsByID(clip.ID) {
					if evt != nil {
						evt(clip, nil)
					}
					continue
				}
				r, err := s.Clips.GetContentByID(clip.ID)
				if err != nil {
					return err
				}
				dec := s.Decoder.Decode(r)
				file, err := s.Storage.Save(clip, dec)
				if err != nil {
					return err
				}
				r.Close()
				if evt != nil {
					evt(clip, &file)
				}
			}
		}
	}
	return nil
}
