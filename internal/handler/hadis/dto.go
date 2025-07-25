package hadis

type Request struct {
	Title string `json:"title"`
}


type Response struct {
	Id int `json:"Номер Хадиса"`
	Title string `json:"title"`
}