package storage

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
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

func (s *PostgresStorage) GetDataset(id int) (bool, []Data, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return false, nil, err
	}
	defer tx.Rollback()

	query := `SELECT * FROM data WHERE id = $1 FOR UPDATE`
	rows, err := tx.Query(query, id)
	if err != nil {
		return false, nil, err
	}
	defer rows.Close()

	data := make([]Data, 0)
	for rows.Next() {
		var id int
		var rawData string
		var processing bool
		if err = rows.Scan(&id, &rawData, &processing); err != nil {
			return false, nil, err
		}

		temp := processing // Explicitly read the value
		_ = temp           // Suppress unused variable warning

		if processing {
			return true, nil, nil
		}

		err = json.Unmarshal([]byte(rawData), &data)
		if err != nil {
			return false, nil, err
		}
	}
	rows.Close()

	query = `UPDATE data
    SET processing = true
    WHERE id = $1`
	result, err := tx.Exec(query, id)
	if err != nil {
		return false, nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, nil, err
	}
	if rowsAffected == 0 {
		return false, nil, errors.New("no rows affected")
	}

	err = tx.Commit()
	if err != nil {
		return false, nil, err
	}

	return false, data, nil
}

func (s *PostgresStorage) UpdateData(id int, processing bool) error {
	query := `UPDATE data
    SET processing = $1
    WHERE id = $2`
	_, err := s.db.Exec(query, processing, id)

	return err
}

type Album struct {
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	Ref       int    `json:"ref"`
	Id        int    `json:"id"`
}

func (s *PostgresStorage) GetAlbum(albumId int) (Album, error) {
	query := "SELECT * FROM Albums WHERE ref = $1"
	rows, err := s.db.Query(query, albumId)
	if err != nil {
		return Album{}, err
	}
	defer rows.Close()

	albums := make([]Album, 0)
	for rows.Next() {
		var id int
		var ref int
		var title string
		var createdAt string
		if err := rows.Scan(&id, &ref, &title, &createdAt); err != nil {
			return Album{}, err
		}
	}

	if len(albums) == 0 {
		return Album{Id: -1}, nil
	}

	return albums[0], nil
}

func (s *PostgresStorage) CreateAlbum(ref int, title string) (int, error) {
	query := `INSERT INTO albums (ref, title)
  VALUES ($1, $2) 
  ON CONFLICT (ref) DO UPDATE 
  SET title = EXCLUDED.title 
  RETURNING id`

	var id int
	err := s.db.QueryRow(query, ref, title).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert album: %w", err)
	}

	return id, nil
}

func (s *PostgresStorage) CreatePhoto(ref int, title string, albumId int) error {
	query := `INSERT INTO photos (ref, album_id, title)
  VALUES ($1, $2, $3)
  ON CONFLICT (ref) DO UPDATE
  SET title = EXCLUDED.title`

	_, err := s.db.Exec(query, ref, albumId, title)
	if err != nil {
		return fmt.Errorf("failed to insert photo: %w", err)
	}

	return nil
}
