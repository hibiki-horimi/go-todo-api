# Basic CRUD RESTful API for Todo

This is a basic CRUD (Create, Read, Update, Delete) RESTful API for managing Todo tasks. The API allows you to perform common operations on Todo items, such as creating new tasks, retrieving task details, updating task information, and deleting tasks.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)

## Features

- Create a new Todo task.
- Retrieve a list of all Todo tasks.
- Retrieve details of a specific Todo task by ID.
- Updates the name and/or the status of a Todo task.
- Delete a Todo task by ID.

## Technologies Used

The following technologies were used to create this API:

- Go: The programming language used for server-side development.
- Echo: A fast and minimalist web framework for Go.
- PostgreSQL: A powerful open-source relational database for storing Todo task data.
- Docker: A containerization platform used to package the application and its dependencies.

## Installation

1. Make sure you have Docker installed on your system.

2. Clone this repository to your local machine.

```bash
git clone https://github.com/hibiki-horimi/go-todo-api
```

3. Navigate to the project directory.

```bash
cd go-todo-api
```

4. Run the Docker container.

```bash
docker compose up
```

5. Open another shell and run database migrations to set up the necessary tables.

```bash
go run main.go migrate
```

## Usage

You can use tools like curl, Postman, or any other REST client to interact with the API. Below are the available endpoints and their functionalities.

## API Endpoints

`GET /api/todos`

Retrieves a list of all Todo tasks.

```bash
curl localhost:8080/api/todos
```

`POST /api/todos`

Creates a new Todo task. Requires a JSON payload in the request body with the following fields:

task (required): The name of the Todo task.

```bash
curl -X POST -H "Content-Type: application/json" -d '{"task":"Buy Groceries"}' localhost:8080/api/todos
```

`GET /api/todos/:id`

Retrieves details of a specific Todo task by its ID.

```bash
curl localhost:8080/api/todos/12345678-90ab-cdef-ghij-klmnopqrstuv
```

`PUT /api/todos/:id`

Updates the name and/or the status of a Todo task by its ID. Requires a JSON payload in the request body with the fields to be updated.

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"task":"Book Flight Tickets", "done": true}' localhost:8080/api/todos/12345678-90ab-cdef-ghij-klmnopqrstuv
```

`DELETE /api/todos/:id`

Deletes a Todo task by its ID.

```bash
curl -X DELETE localhost:8080/api/todos/12345678-90ab-cdef-ghij-klmnopqrstuv
```
