package models

import (
	"database/sql"
	"time"
)

type SnippetModel struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetRepository struct {
	DB *sql.DB
}

func (s *SnippetRepository) Insert(title string, content string, expires int) (int, error) {
	stmt := `
		INSERT INTO snippets (title, content, created, expires)
		VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))
	`
	result, err := s.DB.Exec(stmt, title, content, expires)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *SnippetRepository) Get(id int) (*SnippetModel, error) {
	return nil, nil
}

func (s *SnippetRepository) Latest() ([]*SnippetModel, error) {
	return nil, nil
}
