package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type AlbumService struct {
	DB *sql.DB
}

func (s *AlbumService) fetchAlbums(p *Pagination) ([]Album, error) {
	var (
		rows *sql.Rows
		err  error
	)
	if p != nil {
		rows, err = s.DB.Query("select * from album limit $1 offset $2", p.Limit, p.Offset)
	} else {
		rows, err = s.DB.Query("select * from album")
	}

	if err != nil {
		errorLog.Error(err)
		return nil, errors.New("could not fetch albums")
	}
	defer rows.Close()
	albums := []Album{}

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			errorLog.Errorf("Could not scan")
		}
		albums = append(albums, alb)
	}
	return albums, nil
}

func (s *AlbumService) fetchAlbumById(id int64) (*Album, error) {

	row := s.DB.QueryRow("select * from album where id = $1", id)
	err := row.Err()

	if err != nil {
		errorLog.Error(err)
		return nil, fmt.Errorf("could not provide album with id %d", id)
	}

	alb := Album{}
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		return nil, nil
	}
	return &alb, nil
}

func (s *AlbumService) createAlbum(body CreateAlbumBody) (*Album, error) {
	row := s.DB.QueryRow(
		"INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING *",
		body.Title,
		body.Artist,
		body.Price,
	)

	err := row.Err()

	if err != nil {
		errorLog.Error(err)
		return nil, errors.New("could not save new album")
	} else {
		alb := Album{}
		if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			// in case RETURNING clause failed or mistyped
			return nil, nil
		} else {
			return &alb, nil
		}
	}
}
