package hadis

import "widjetHadis/internal/service/hadis"

type Service interface {
	Add(input *hadis.CreateModel) (*hadis.Model, error)
}
