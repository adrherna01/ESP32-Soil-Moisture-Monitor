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

type Device struct {
	Ip string `json:"ip"`
	Port string `json:"port"`
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

    go pollSensor()

	http.HandleFunc("/measurements", measurementHandler)
	http.HandleFunc("/register", registerHandler)
	
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

	query := `INSERT INTO measurements (humidity, timestamp) VALUES ($1, NOW())`
    _, err = db.Exec(query, m.Humidity)
    if err != nil {
        log.Println("DB insert error:", err)
        http.Error(w, "Failed to insert", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Measurement logged"))
}

func registerHandler(w http.ResponseWriter, r *http.Request) { 
	if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	var d Device
    err := json.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

	fmt.Printf("From backend Ip address: %s Port: %s\n", d.Ip, d.Port)

	query := `INSERT INTO devices (ip, port) VALUES ($1, $2)`
	_, err = db.Exec(query, d.Ip, d.Port)
	if err != nil {
		log.Println("DB insert error:", err)
		http.Error(w, "Failed to insert", http.StatusInternalServerError)
		return
	}
	
    w.WriteHeader(http.StatusOK)
    fmt.Println("Device logged")
}

func pollSensor() {
	interval := 10 * time.Second // poll every 10s

	for {
		var ip, port string
		err := db.QueryRow(`SELECT ip, port FROM devices ORDER BY id ASC LIMIT 1`).Scan(&ip, &port)
		if err != nil {
			log.Println("No device found or DB error:", err)
			time.Sleep(interval)
			continue
		}

		url := fmt.Sprintf("http://%s:%s/measure", ip, port)

		resp, err := http.Get(url)
		if err != nil {
			log.Println("Failed to contact ESP32:", err)
			time.Sleep(interval)
			continue
		}

		var m Measurement
		err = json.NewDecoder(resp.Body).Decode(&m)
		resp.Body.Close()
		if err != nil {
			log.Println("Failed to decode response:", err)
			time.Sleep(interval)
			continue
		}

		log.Printf("Polled ESP32: %+v", m)

		query := `INSERT INTO measurements (humidity) VALUES ($1)`
		_, err = db.Exec(query, m.Humidity)
		if err != nil {
			log.Println("DB insert error:", err)
		}

		time.Sleep(interval)
	}
}
