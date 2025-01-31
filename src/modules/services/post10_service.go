package services

import (
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go-crud/modules/repository"
    "go-crud/modules/models"
)

type Post10Service struct {
    Repo *repository.Repository
}

func NewPost10Service(repo *repository.Repository) *Post10Service {
    return &Post10Service{Repo: repo}
}

func (s *Post10Service) CreatePost10(post10 *models.Post10) (interface{}, error) {
    return s.Repo.Create("post10s", post10)
}

func (s *Post10Service) GetAllPost10s() ([]bson.M, error) {
    return s.Repo.FindAll("post10s", bson.M{})
}

func (s *Post10Service) GetPost10ByID(id string) (interface{}, error) {
    objectID, _ := primitive.ObjectIDFromHex(id)
    return s.Repo.FindOne("post10s", bson.M{"_id": objectID})
}

func (s *Post10Service) UpdatePost10(id string, update bson.M) error {
    objectID, _ := primitive.ObjectIDFromHex(id)
    return s.Repo.Update("post10s", bson.M{"_id": objectID}, update)
}

func (s *Post10Service) DeletePost10(id string) error {
    objectID, _ := primitive.ObjectIDFromHex(id)
    return s.Repo.Delete("post10s", bson.M{"_id": objectID})
}