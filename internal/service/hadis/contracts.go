package hadis

import "widjetHadis/internal/repository/hadis"

type Repo interface {
	Add(m *hadis.Model) (*hadis.Model, error)
}
