package hadis

func (r *Repository) Add(m *Model) (*Model, error) {
	returnedM := &Model{}
	query := `INSERT INTO hadis(title)
	VALUES ($1)
	RETURNING id, title`
	err := r.store.GetConn().QueryRow(query, m.Title).Scan(&returnedM.Id, &returnedM.Title)
	if err != nil {
		return nil, err
	}
	return returnedM, nil
}
