package hadis

import "widjetHadis/internal/service/hadis"

type Request struct {
	Title string `json:"title"`
}

type Response struct {
	Id    int    `json:"Номер Хадиса"`
	Title string `json:"title"`
}

func (m *Request) ToSrv() hadis.CreateModel {
	return hadis.CreateModel{
		Title: m.Title,
	}
}
func (m *Response) FromSrv(srv *hadis.Model) {
	m.Id = srv.Id
	m.Title = srv.Title
}
