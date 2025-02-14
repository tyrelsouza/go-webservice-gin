package models

type Album struct {
	ID     string  `json:"id" gorm:"primary_key"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type CreateAlbumInput struct {
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	Price  string `json:"price" binding:"required"`
}

type UpdateAlbumInput struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
