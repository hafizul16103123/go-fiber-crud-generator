package repository

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
    DB *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
    return &Repository{DB: db}
}

func (r *Repository) Create(collection string, document interface{}) (interface{}, error) {
    result, err := r.DB.Collection(collection).InsertOne(context.Background(), document)
    if err != nil {
        return nil, err
    }
    return result.InsertedID, nil
}

func (r *Repository) FindAll(collection string, filter bson.M) ([]bson.M, error) {
    cursor, err := r.DB.Collection(collection).Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var results []bson.M
    if err := cursor.All(context.Background(), &results); err != nil {
        return nil, err
    }
    return results, nil
}

func (r *Repository) FindOne(collection string, filter bson.M) (bson.M, error) {
    var result bson.M
    err := r.DB.Collection(collection).FindOne(context.Background(), filter).Decode(&result)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func (r *Repository) Update(collection string, filter bson.M, update bson.M) error {
    _, err := r.DB.Collection(collection).UpdateOne(context.Background(), filter, update)
    return err
}

func (r *Repository) Delete(collection string, filter bson.M) error {
    _, err := r.DB.Collection(collection).DeleteOne(context.Background(), filter)
    return err
}