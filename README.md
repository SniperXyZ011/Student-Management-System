# Student Management System (SMS)

A simple REST API for managing student records in a school or university, built with Go. This project serves as a learning tool to understand building CRUD backends and APIs in Go.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete student records.
- **Data Validation**: Input validation using the [`go-playground/validator`](../../D:/go_projects/SMS/internal/http/handlers/student/student.go ) library.
- **Structured Logging**: Uses Go's [`slog`](../../D:/go_projects/SMS/cmd/student-api/main.go ) for production-ready logging.
- **Database Support**: Compatible with MySQL and SQLite databases.
- **Clean Architecture**: Organized into internal packages for config, handlers, storage, types, and utilities.

## Tech Stack

- **Language**: Go 1.25.4
- **Web Framework**: Standard library [`net/http`](../../D:/go_projects/SMS/cmd/student-api/main.go ) with [`http.ServeMux`](../../D:/go_projects/SMS/cmd/student-api/main.go )
- **Database Drivers**: `github.com/go-sql-driver/mysql` and `github.com/mattn/go-sqlite3`
- **Validation**: `github.com/go-playground/validator/v10`
- **Configuration**: `github.com/ilyakaznacheev/cleanenv` for YAML config loading

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/SniperXyZ011/Student-Management-System.git
   cd Student-Management-System
   ```

2. **Install Go**:
   Ensure you have Go 1.25 or later installed. You can download it from [golang.org](https://golang.org/dl/).

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Set up the database**:
   - For MySQL: Create a database named [`students`](../../D:/go_projects/SMS/internal/http/handlers/student/student.go ) and update the connection string in [`SMS/config/local.yaml`](SMS/config/local.yaml ).
   - For SQLite: The database file will be created automatically if using SQLite (uncomment the SQLite storage in [`SMS/cmd/student-api/main.go`](SMS/cmd/student-api/main.go )).

5. **Configure the application**:
   Edit [`SMS/config/local.yaml`](SMS/config/local.yaml ) to set your database connection and server address.

## Configuration

The configuration is managed via [`SMS/config/local.yaml`](SMS/config/local.yaml ). Example:

```yaml
env: "dev"
storage_path: "root:201104@tcp(localhost:3306)/students"  # MySQL connection string
http_server:
  address: "localhost:8080"
```

- `env`: Environment (e.g., "dev", "prod")
- `storage_path`: Database connection string (MySQL format shown; for SQLite, use a file path)
- `http_server.address`: Server listen address

## Usage

1. **Run the application**:
   ```bash
   go run cmd/student-api/main.go
   ```

2. The server will start on the configured address (default: `localhost:8080`).

3. Use tools like `curl`, Postman, or any HTTP client to interact with the API.

## API Endpoints

| Method | Endpoint              | Description              |
|--------|-----------------------|--------------------------|
| POST   | [`/api/students`](../../D:/go_projects/SMS/internal/http/handlers/student/student.go )      | Create a new student    |
| GET    | [`/api/students`](../../D:/go_projects/SMS/internal/http/handlers/student/student.go )      | Get all students        |
| GET    | [`/api/students/{id}`](../../D:/go_projects/SMS/internal/http/handlers/student/student.go ) | Get student by ID       |
| PUT    | [`/api/students/{id}`](../../D:/go_projects/SMS/internal/http/handlers/student/student.go ) | Update student by ID    |
| DELETE | [`/api/students/{id}`](../../D:/go_projects/SMS/internal/http/handlers/student/student.go ) | Delete student by ID    |

### Request/Response Examples

- **Create Student** (POST [`/api/students`](../../D:/go_projects/SMS/internal/http/handlers/student/student.go )):
  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 20
  }
  ```
  Response: [`{"id": 1}`](../../D:/go_projects/SMS/internal/http/handlers/student/student.go )

- **Get All Students** (GET [`/api/students`](../../D:/go_projects/SMS/internal/http/handlers/student/student.go )):
  Response:
  ```json
  [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "age": 20
    }
  ]
  ```

- **Get Student by ID** (GET `/api/students/1`):
  Response: Same as above for the specific student.

- **Update Student** (PUT `/api/students/1`):
  Request body same as create, but only provided fields are updated.

- **Delete Student** (DELETE `/api/students/1`):
  Response: No content (204) on success.

## Project Structure

```
SMS/
├── cmd/
│   └── student-api/
│       └── main.go              # Application entry point
├── config/
│   └── local.yaml               # Configuration file
├── internal/
│   ├── config/
│   │   └── config.go            # Config loading logic
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go   # HTTP handlers for student operations
│   ├── storage/
│   │   ├── storage.go           # Storage interface
│   │   ├── sql/
│   │   │   └── sql.go           # MySQL implementation
│   │   └── sqlite/
│   │       └── sqlite.go        # SQLite implementation
│   ├── types/
│   │   └── types.go             # Data models
│   └── utils/
│       └── response/
│           └── response.go      # Response helpers
├── go.mod                       # Go module file
├── info.txt                     # Project info and learning notes
└── .gitignore                   # Git ignore rules
```

## Learning Outcomes

This project demonstrates:
1. Building REST APIs in Go using the standard library.
2. Implementing CRUD operations with database integration.
3. Input validation and error handling.
4. Structured logging for better observability.
5. Clean project structure and separation of concerns.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details. (Note: Add a LICENSE file if not present.)