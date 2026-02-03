package hadis

import "widjetHadis/internal/repository/hadis"

func (s *Service) Add(input CreateModel) (*Model, error) {
	toDB := hadis.Model{
		Title: input.Title,
	}

	created, err := s.repo.Add(&toDB)
	if err != nil {
		return nil, err
	}
	fromDB := Model{
		Id:    created.Id,
		Title: created.Title,
	}
	return &fromDB, nil

}
