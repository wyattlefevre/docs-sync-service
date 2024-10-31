package repository

import (
	"database/sql"
	"docs-sync-service/models"
)

type FileRepository struct {
	DB *sql.DB
}

func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{DB: db}
}

func (r *FileRepository) Create(file *models.File) (*int, error) {
	// Implementation for inserting a new file into the database
	var id int
	err := r.DB.QueryRow(`
        INSERT INTO files (code_repository, code_file_path, docs_repository, docs_file_path)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (code_repository, code_file_path, docs_repository, docs_file_path) DO NOTHING
        RETURNING id
    `, file.CodeRepository, file.CodeFilePath, file.DocsRepository, file.DocsFilePath).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *FileRepository) GetAll() ([]models.File, error) {
	rows, err := r.DB.Query(`
        SELECT id, code_repository, code_file_path, docs_repository, docs_file_path
        FROM files
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []models.File
	for rows.Next() {
		var file models.File
		err := rows.Scan(&file.ID, &file.CodeRepository, &file.CodeFilePath, &file.DocsRepository, &file.DocsFilePath)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *FileRepository) GetByID(id int) (*models.File, error) {
	var file models.File
	err := r.DB.QueryRow(`
        SELECT id, code_repository, code_file_path, docs_repository, docs_file_path
        FROM files
        WHERE id = $1
    `, id).Scan(&file.ID, &file.CodeRepository, &file.CodeFilePath, &file.DocsRepository, &file.DocsFilePath)

	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *FileRepository) Update(file *models.File) error {
	_, err := r.DB.Exec(`
        UPDATE files
        SET code_repository = $1, code_file_path = $2, docs_repository = $3, docs_file_path = $4
        WHERE id = $5
    `, file.CodeRepository, file.CodeFilePath, file.DocsRepository, file.DocsFilePath, file.ID)

	return err
}

func (r *FileRepository) Delete(id int) error {
	_, err := r.DB.Exec(`
        DELETE FROM files
        WHERE id = $1
    `, id)

	return err
}

func (r *FileRepository) GetByCodeRepo(repo string)([]models.File, error) {
	rows, err := r.DB.Query(`
        SELECT id, code_repository, code_file_path, docs_repository, docs_file_path
        FROM files
		WHERE code_repository = $1
	`, repo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []models.File
	for rows.Next() {
		var file models.File
		err := rows.Scan(&file.ID, &file.CodeRepository, &file.CodeFilePath, &file.DocsRepository, &file.DocsFilePath)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *FileRepository) GetByCodeFilePath(repo string, filePath string)([]models.File, error) {
	rows, err := r.DB.Query(`
        SELECT id, code_repository, code_file_path, docs_repository, docs_file_path
        FROM files
		WHERE code_repository = $1 AND code_file_Path = $2
	`, repo, filePath)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []models.File
	for rows.Next() {
		var file models.File
		err := rows.Scan(&file.ID, &file.CodeRepository, &file.CodeFilePath, &file.DocsRepository, &file.DocsFilePath)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}


func (r *FileRepository) GetByDocsRepo(repo string)([]models.File, error) {
	rows, err := r.DB.Query(`
        SELECT id, code_repository, code_file_path, docs_repository, docs_file_path
        FROM files
		WHERE docs_repository = $1
	`, repo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []models.File
	for rows.Next() {
		var file models.File
		err := rows.Scan(&file.ID, &file.CodeRepository, &file.CodeFilePath, &file.DocsRepository, &file.DocsFilePath)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

