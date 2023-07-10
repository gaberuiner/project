package storage

import (
	"database/sql"
	"errors"
	"log"
)

type Storage struct {
	DB *sql.DB
	StorageI
}

func New() *Storage {
	db, err := sql.Open("sqlite3", "./database/data.db")
	if err != nil {
		log.Fatal(err)
	}
	return &Storage{
		DB: db,
	}
}

type StorageI interface {
	StorageConnect() error
	StorageAdd() error
	StorageDelete() error
	GetPrice() error
	GetTodayStat() error
}

func (s *Storage) StorageConnect() error {

	statement := `CREATE TABLE IF NOT EXISTS Menu(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		price INTEGER NOT NULL);
	)`
	stat, err := s.DB.Prepare(statement)
	if err != nil {
		return errors.New("can't prepare statement")
	}
	_, err = stat.Exec()
	if err != nil {
		return errors.New("can't exec statement")
	}
	return nil
}

func (s *Storage) StorageAdd() error {
	return nil
}

func (s *Storage) StorageDelete() error {
	return nil
}

func (s *Storage) GetPrice() error {
	return nil
}

func (s *Storage) GetTodayStat() error {
	return nil
}
