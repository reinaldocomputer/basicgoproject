# Go Basic Web Project (Work in Progress)

The goal of this project is to give a basic, simple and organized template for building your Web GoLang Application.

## Project Structure
This project structure was created using as basis [Package Oriented Design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html)
to organize files and folders.

## API Design
This template was designed using concepts of *Data Oriented Programming*.

## Usage
Using your terminal at root directory:
```
go run .
```

### Example
```
curl --location --request POST 'http://localhost:8080/basic' \
--header 'Content-Type: application/json' \
--data-raw '[
    {
        "id": 1,
        "name": "Guest 1",
        "age": 33
    },
    {
        "id": 2,
        "name": "Guest 2",
        "age": 24
    }
]'
```

More examples can be found [here](./go-basic-project.postman_collection.json).