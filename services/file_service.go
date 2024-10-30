package services

import (
    "docs-sync-service/models"
    "docs-sync-service/repository"
)

type FileService struct {
    Repo *repository.FileRepository
}

func NewFileService(repo *repository.FileRepository) *FileService {
    return &FileService{Repo: repo}
}

func (s *FileService) Create(file *models.File) error {
    return s.Repo.Create(file)
}

func (s *FileService) GetAll() ([]models.File, error) {
    return s.Repo.GetAll()
}

func (s *FileService) GetByID(id int) (*models.File, error) {
    return s.Repo.GetByID(id)
}

func (s *FileService) Update(file *models.File) error {
    return s.Repo.Update(file)
}

func (s *FileService) Delete(id int) error {
    return s.Repo.Delete(id)
}

