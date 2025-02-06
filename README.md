# GoLang Library REST API

## Overview
This project is a simple REST API built with Go that manages books and authors in a library system. The API supports CRUD operations and follows best Go practices.

## Features
- **Add a book** (`POST /books`)
- **Retrieve all books** (`GET /books`)
- **Retrieve a book by ID** (`GET /books/{id}`)
- **Delete a book by ID** (`DELETE /books/{id}`)
- **Basic error handling**
- **Unit tests for API endpoints**
- **Bonus:** Search books by author (`GET /books?author=Author Name`)

## Technologies Used
- **Go** (Golang)
- **Gorilla Mux** (Router Library)
- **Built-in Go Testing Package**

## Installation & Running the API

1. Clone the repository:
   ```sh
   git clone https://github.com/faigy2549/library-api.git
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Run the application:
   ```sh
   go run main.go
   ```

4. The API will be available at `http://localhost:8080`.

## API Endpoints

### Add a Book
- **Endpoint:** `POST /books`
- **Request Body:**
  ```json
  {
    "title": "Book Title",
    "author": "Author Name",
    "publishedYear": 2022
  }
  ```
- **Response:** `201 Created`

---

### Retrieve All Books
- **Endpoint:** `GET /books`
- **Response:** `200 OK`
  ```json
  [
    {
      "id": 1,
      "title": "Book Title",
      "author": "Author Name",
      "publishedYear": 2022
    }
  ]
  ```

---

### Retrieve a Book by ID
- **Endpoint:** `GET /books/{id}`
- **Response:** `200 OK`
  ```json
  {
    "id": 1,
    "title": "Book Title",
    "author": "Author Name",
    "publishedYear": 2022
  }
  ```

---

### Delete a Book
- **Endpoint:** `DELETE /books/{id}`
- **Response:** `204 No Content`

---

### Search Books by Author (Bonus)
- **Endpoint:** `GET /books?author=Author Name`
- **Response:** `200 OK`
  ```json
  [
    {
      "id": 1,
      "title": "Book Title",
      "author": "Author Name",
      "publishedYear": 2022
    }
  ]
  ```

## Running Tests
To run unit tests, use:
```sh
go test ./...
```

## Project Structure
```
/library-api
│── main.go            # Entry point
│── routes.go          # API route definitions
│── handlers.go        # Handlers for API requests
│── models.go          # Data models
│── storage.go         # In-memory data storage
│── tests/             # Unit tests
│── go.mod             # Module dependencies
│── README.md          # Documentation
```

## Error Handling
- `404 Not Found` if a book does not exist.
- `400 Bad Request` for invalid input.
- Proper HTTP status codes for all operations.

## License
This project is licensed under the MIT License.
