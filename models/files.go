package models

import (
	_ "github.com/lib/pq"
)

type File struct {
	ID       string
	UserID   string
	Filename string
	URL      string
}

// SaveFileMetadata inserts file metadata into the database
func SaveFileMetadata(userID, filename, url string) error {
	_, err := db.Exec("INSERT INTO files (user_id, filename, url) VALUES ($1, $2, $3)", userID, filename, url)
	return err
}

// GetFilesByUserID retrieves files by user ID
func GetFilesByUserID(userID string) ([]File, error) {
	rows, err := db.Query("SELECT id, filename, url FROM files WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []File
	for rows.Next() {
		var file File
		err := rows.Scan(&file.ID, &file.Filename, &file.URL)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}
