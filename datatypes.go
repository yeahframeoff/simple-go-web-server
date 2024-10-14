package main

type CreateAlbumBody struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Album struct {
	ID int64 `json:"id"`
	CreateAlbumBody
}
