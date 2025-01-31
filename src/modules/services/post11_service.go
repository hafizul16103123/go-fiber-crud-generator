package services

import (
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/hafizul16103123/go-fiber-crud-generator/modules/repository"
    "github.com/hafizul16103123/go-fiber-crud-generator/modules/models"
)

type Post11Service struct {
    Repo *repository.Repository
}

func NewPost11Service(repo *repository.Repository) *Post11Service {
    return &Post11Service{Repo: repo}
}

func (s *Post11Service) CreatePost11(post11 *models.Post11) (interface{}, error) {
    return s.Repo.Create("post11s", post11)
}

func (s *Post11Service) GetAllPost11s() ([]bson.M, error) {
    return s.Repo.FindAll("post11s", bson.M{})
}

func (s *Post11Service) GetPost11ByID(id string) (interface{}, error) {
    objectID, _ := primitive.ObjectIDFromHex(id)
    return s.Repo.FindOne("post11s", bson.M{"_id": objectID})
}

func (s *Post11Service) UpdatePost11(id string, update bson.M) error {
    objectID, _ := primitive.ObjectIDFromHex(id)
    return s.Repo.Update("post11s", bson.M{"_id": objectID}, update)
}

func (s *Post11Service) DeletePost11(id string) error {
    objectID, _ := primitive.ObjectIDFromHex(id)
    return s.Repo.Delete("post11s", bson.M{"_id": objectID})
}