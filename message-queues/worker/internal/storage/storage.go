package storage

import (
	"database/sql"
	"encoding/json"
	"log/slog"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db  *sql.DB
	log *slog.Logger
}

func NewPostgresStorage(connectionString string, logger *slog.Logger) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db:  db,
		log: logger,
	}, nil
}

func (s *PostgresStorage) Close() error {
	return s.db.Close()
}

type Data struct {
	Title   string `json:"title"`
	Id      int    `json:"id"`
	AlbumId int    `json:"albumId"`
}

func (s *PostgresStorage) GetDataset(id int) ([]Data, error) {
	query := "SELECT * FROM data WHERE id = $1"
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]Data, 0)
	for rows.Next() {
		var id int
		var rawData string
		if err := rows.Scan(&id, &rawData); err != nil {
			return nil, err
		}

		err := json.Unmarshal([]byte(rawData), &data)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

type Album struct {
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	Id        int    `json:"id"`
}

func (s *PostgresStorage) GetAlbum(albumId int) (Album, error) {
	query := "SELECT * FROM Albums WHERE id = $1"
	rows, err := s.db.Query(query, albumId)
	if err != nil {
		return Album{}, err
	}
	defer rows.Close()

	albums := make([]Album, 0)
	for rows.Next() {
		var id int
		var title string
		var createdAt string
		if err := rows.Scan(&id, &title, &createdAt); err != nil {
			return Album{}, err
		}
	}

	if len(albums) == 0 {
		return Album{Id: -1}, nil
	}

	return albums[0], nil
}
