package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
)

func main() {
    dbURL := os.Getenv("DATABASE_URL")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        db, err := sql.Open("postgres", dbURL)
        status := "SUCCESS"
        if err != nil || db.Ping() != nil {
            status = "FAILED"
        }
        defer db.Close()
        fmt.Fprintf(w, "<h1>DevOps Project Status</h1><p>Application: RUNNING</p><p>Database Connection: %s</p>", status)
    })

    log.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
