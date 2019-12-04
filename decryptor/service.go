package decryptor

type Service struct {
	Decoder Decoder
	Storage Storage
	Courses CourseRepository
	Clips   ClipRepository
}

func (s *Service) DecryptAll() error {
	courses, err := s.Courses.FindAll()
	if err != nil {
		return err
	}
	for _, course := range courses {
		for _, module := range course.Modules {
			for _, clip := range module.Clips {
				if !s.Clips.ExistsByID(clip.ID) {
					err := s.Storage.SavePlaceholder(clip)
					if err != nil {
						return err
					}
					continue
				}
				r, err := s.Clips.GetContentByID(clip.ID)
				if err != nil {
					return err
				}
				defer r.Close()
				dec := s.Decoder.Decode(r)
				err = s.Storage.Save(clip, dec)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
