# Syndio Backend App

This is a simple Go-based REST API to save job data associated with employees stored in an SQLite database.

## Project Structure

```
syndio-backend-app/
├── main.go
├── config/
│   └── config.go                # Handles configuration (e.g., API port)
├── database/
│   └── database.go              # Database setup and connection
├── handlers/
│   └── employee_handler.go      # API endpoint handler for saving job data
├── models/
│   └── job.go                   # Job model and DB insertion logic
├── employees.db                 # SQLite database containing employee data
└── go.mod                       # Go module file
```

## Prerequisites

- **Go**: Ensure that [Go](https://golang.org/dl/) is installed on your system (version >= 1.21.3).
- **SQLite**: This project uses an SQLite database (`employees.db`) with predefined employee data.

## Installation and Setup Without Docker

1. **Clone the repository** and navigate into the project directory:

   ```bash
   git clone https://github.com/jokoyoski/Syndio-BE-Takehome_Assignment-Adeola-Oladeinde
   cd syndio-backend-app
   ```

2. **Install Go dependencies**:

   ```bash
   go mod tidy
   ```

3. **Set up the database**:  
   Ensure that `employees.db` is present in the project root. This database should contain the `employees` table with employee data as per the initial requirement.

4. **Configure the PORT** (optional):  
   By default, the API listens on port `8080`. Create a `.env` file with the following variable:
   ```
   PORT=9000
   ```

## Running the Server

To start the server, run:

```bash
go run main.go
```

You should see output indicating that the server is running on the specified port:

```
Server running on port :{PORT}
```

## API Documentation

### 1. Save Job Data

- **Endpoint**: `/save-jobs`
- **Method**: `POST`
- **Content-Type**: `application/json`

This endpoint allows you to save job data associated with employees in the database.

- **Request Payload**:
  ```json
  [
    {
      "employee_id": 1,
      "department": "Engineering",
      "job_title": "Senior Engineer"
    },
    {
      "employee_id": 2,
      "department": "Engineering",
      "job_title": "Super Senior Engineer"
    },
    { "employee_id": 3, "department": "Sales", "job_title": "Head of Sales" },
    { "employee_id": 4, "department": "Support", "job_title": "Tech Support" },
    {
      "employee_id": 5,
      "department": "Engineering",
      "job_title": "Junior Engineer"
    },
    { "employee_id": 6, "department": "Sales", "job_title": "Sales Rep" },
    {
      "employee_id": 7,
      "department": "Marketing",
      "job_title": "Senior Marketer"
    }
  ]
  ```
- **Postman Link**
  https://www.postman.com/planetary-desert-668165/workspace/syndio/collection/34269245-410dc2f5-46bc-4cdc-95da-433aae73e459?action=share&creator=34269245

- **Response**:

  - **Success**: Returns `201 Created` and a confirmation message.
  - **Failure**: Returns an error message with the appropriate HTTP status code.

- **Example**:

  ```bash
  curl -X POST http://localhost:8080/save-jobs -H "Content-Type: application/json" -d '[
    { "employee_id": 1, "department": "Engineering", "job_title": "Senior Engineer" },
    { "employee_id": 2, "department": "Engineering", "job_title": "Super Senior Engineer" },
    { "employee_id": 3, "department": "Sales", "job_title": "Head of Sales"},
    { "employee_id": 4, "department": "Support", "job_title": "Tech Support" },
    { "employee_id": 5, "department": "Engineering", "job_title": "Junior Engineer" },
    { "employee_id": 6, "department": "Sales", "job_title": "Sales Rep" },
    { "employee_id": 7, "department": "Marketing", "job_title": "Senior Marketer" }
  ]'
  ```

  **Expected Response**:

  ```json
  {
    "message": "Job data saved successfully"
  }
  ```

## Troubleshooting

- **Error Connecting to Database**: Ensure `employees.db` exists in the project directory and is accessible.
- **Invalid Payload**: Confirm that the request payload is in the expected format as shown above.

## Notes

- **Environment Variable**: If no `PORT` is provided, the server defaults to `8080`.
- **Data Persistence**: Job data is stored in the `jobs` table within the `employees.db` SQLite database.
