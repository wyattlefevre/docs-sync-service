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

func (s *FileService) Create(file *models.File) (*int, error) {
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

func (s *FileService) GetByCodeRepo(repo string) ([]models.File, error){
    return s.Repo.GetByCodeRepo(repo)
}

func (s *FileService) GetByCodeFilePath(repo string, filePath string) ([]models.File, error){
    return s.Repo.GetByCodeFilePath(repo, filePath)
}

func (s *FileService) GetByDocsRepo(repo string) ([]models.File, error){
    return s.Repo.GetByDocsRepo(repo)
}
