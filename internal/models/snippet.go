package models

import (
	"database/sql"
	"errors"
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
	stmt := `
		SELECT * 
		FROM snippets
		WHERE expires > UTC_TIMESTAMP()
		AND id = ?
	`

	row := s.DB.QueryRow(stmt, id)

	snp := &SnippetModel{}

	err := row.Scan(&snp.ID, &snp.Title, &snp.Content, &snp.Created, &snp.Expires)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNoRecord
	}

	if err != nil {
		return nil, err
	}

	return snp, nil
}

func (s *SnippetRepository) Latest() ([]*SnippetModel, error) {
	stmt := `
		SELECT *
		FROM snippets
		WHERE expires > UTC_TIMESTAMP()
		ORDER BY id DESC
		LIMIT 10
	`

	rows, err := s.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snps := []*SnippetModel{}

	for rows.Next() {
		snp := &SnippetModel{}

		err := rows.Scan(&snp.ID, &snp.Title, &snp.Content, &snp.Created, &snp.Expires)

		if err != nil {
			return nil, err
		}

		snps = append(snps, snp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snps, nil
}
