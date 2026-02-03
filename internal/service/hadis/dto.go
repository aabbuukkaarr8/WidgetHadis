package hadis

type Model struct {
	Id    int    `json:"номер хадиса"`
	Title string `json:"title"`
}
type CreateModel struct {
	Title string `json:"title"`
}