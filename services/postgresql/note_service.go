package services

import (
	"context"

	"github.com/gnotnek/golang-noteapp-backend/models"
)

func GetAllNotes() ([]models.Notes, error) {
	rows, err := DB.Query(context.Background(), "SELECT * FROM notes")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notes := []models.Notes{}
	for rows.Next() {
		var note models.Notes
		err := rows.Scan(&note.ID, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func AddNote(note models.Notes) error {
	_, err := DB.Exec(context.Background(), "INSERT INTO notes (title, body) VALUES ($1, $2)", note.Title, note.Body)
	return err
}
