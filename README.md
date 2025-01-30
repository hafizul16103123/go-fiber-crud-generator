# go-fiber-crud-generator
How to Use the CLI Tool
Build the CLI tool:

go build -o fiber-crud-generator
Generate a CRUD resource:

./fiber-crud-generator generate --name User
The generated project will have the following structure:

Copy
generated-app/
├── models/
│   └── user.go
├── repository/
│   └── repository.go
├── services/
│   └── user_service.go
├── controllers/
│   └── user_controller.go
├── routes/
│   └── user_routes.go
Navigate to the generated project and run it:

cd generated-app
go run main.go