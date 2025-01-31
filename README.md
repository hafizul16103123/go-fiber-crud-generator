
# Hi, I'm Md Hafizul Islam! 👋


# A CLI tool to generate CRUD resources for Go Fiber with MongoDB using Repository pattern.



## Features

**Generate CRUD Controllers:** Quickly create controllers for CRUD operations.

**Customizable Templates:** Use predefined templates or customize them to fit your needs.

**Easy Integration:** Integrate the generated controllers into your Fiber app with minimal effort.

## Available command:
```
Usage:
  fiber-crud-generator [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generate a CRUD resource
  help        Help about any command
  init        Initialize a simple Fiber app with MongoDB connection

Flags:
  -h, --help   help for fiber-crud-generator

Use "fiber-crud-generator [command] --help" for more information about a command.
```


## Run Locally

Clone the project

```bash
 git clone https://github.com/hafizul16103123/go-fiber-crud-generator
```

Go to the project directory

```bash
  cd fiber-crud-generator
```

Build the CLI tool:

```bash
  go build -o fiber-crud-generator
```

Move the binary to a directory in your PATH (optional):

```bash
  sudo mv fiber-crud-generator /usr/local/bin/
```
Create a mod
```
go mod init <module_name>
```

## Uses
To generate a CRUD operation for a resource, run the following command:
```
./fiber-crud-generator generate --name <ResourceName>
For example: Post, User
```
It will generate all CRUD endpoints

```
Test the endpoints using curl or Postman:

POST /posts → "Create Post"

GET /posts → "Get All Post"

GET /posts/1 → "Get One Post"

PUT /posts/1 → "Update Post"

DELETE /posts/1 → "Delete Post"
```

## Customizing Templates
 You cam customize your controller,service,model,router,and repogitiry file.Those files are inside templates folder
## Contributing

Contributions are welcome! If you'd like to contribute, please follow these steps:

Fork the repository.

Create a new branch for your feature or bugfix.

Make your changes and commit them.

Submit a pull request.


## Support

For support, email islamhafizul158@gmail.com.


## Acknowledgements

 - [Go Fiber](https://github.com/gofiber/fiber) for the awesome web framework.
 - [Cobra](https://github.com/spf13/cobra) or the CLI library.

## Enjoy using the Fiber CRUD Generator CLI! 🚀