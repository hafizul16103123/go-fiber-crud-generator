package services

import (
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go-crud/modules/repository"
    "go-crud/modules/models"
)

type UserService struct {
    Repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
    return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *models.User) (interface{}, error) {
    return s.Repo.Create("users", user)
}

func (s *UserService) GetAllUsers() ([]bson.M, error) {
    return s.Repo.FindAll("users", bson.M{})
}

func (s *UserService) GetUserByID(id string) (interface{}, error) {
    objectID, _ := primitive.ObjectIDFromHex(id)
    return s.Repo.FindOne("users", bson.M{"_id": objectID})
}

func (s *UserService) UpdateUser(id string, update bson.M) error {
    objectID, _ := primitive.ObjectIDFromHex(id)
    return s.Repo.Update("users", bson.M{"_id": objectID}, update)
}

func (s *UserService) DeleteUser(id string) error {
    objectID, _ := primitive.ObjectIDFromHex(id)
    return s.Repo.Delete("users", bson.M{"_id": objectID})
}