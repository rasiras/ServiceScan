// main.go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    _ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

func main() {
    var err error
    // Connect to the database (make sure to set the correct connection parameters)
    connStr := "user=youruser dbname=yourdb sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/search", searchHandler)
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// searchHandler handles the search request and is vulnerable to SQL injection.
func searchHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("query")

    // Vulnerable to SQL Injection
    sqlStatement := fmt.Sprintf("SELECT * FROM users WHERE username = '%s';", query)
    rows, err := db.Query(sqlStatement)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var username string
    for rows.Next() {
        err := rows.Scan(&username)
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "User found: %s\n", username)
    }
    if err := rows.Err(); err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}
