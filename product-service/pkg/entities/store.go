package entity

type Store struct {
	Store_id string `json:"store_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
}
