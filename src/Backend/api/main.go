package main

import (
    "log"
    "net/http"
	"database/sql"
	"os"
	"fmt"
	"time"
	"encoding/json"

    _ "github.com/lib/pq"
)

type Measurement struct {
	Humidity float64 `json:"humidity"`
}

var db *sql.DB

func main() {
    dbUser := os.Getenv("POSTGRES_USER")
    dbPassword := os.Getenv("POSTGRES_PASSWORD")
    dbName := os.Getenv("POSTGRES_DB")
    dbHost := "db"
    dbPort := "5432"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    fmt.Println(connStr)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	for {
        if err = db.Ping(); err == nil {
            break
        }
        log.Println("Waiting for database...")
        time.Sleep(2 * time.Second)
    }

	http.HandleFunc("/measurements", measurementHandler)
	
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func measurementHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	var m Measurement
    err := json.NewDecoder(r.Body).Decode(&m)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

	query := `INSERT INTO measurements (humidity) VALUES ($1)`
    _, err = db.Exec(query, m.Humidity)
    if err != nil {
        log.Println("DB insert error:", err)
        http.Error(w, "Failed to insert", http.StatusInternalServerError)
        return
    }
	
	log.Printf("Inserted measurement: %+v\n", m)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Measurement logged"))
}