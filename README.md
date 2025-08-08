# ESP32 Soil Moisture Monitor

This project utilizes an ESP32 microcontroller paired with a YL-69 soil moisture sensor to continuously monitor the moisture level of your plant’s soil. The device connects to WiFi and sends the moisture readings to a backend server, which stores the data in a database. With this setup, you can easily track your plant’s soil conditions and receive timely notifications to know exactly when your plant needs watering.

---

## Features

- Reads soil moisture using the YL-69 sensor (analog pin 34)
- Connects to WiFi using credentials stored in a separate header file
- Sends moisture level notifications via Pushover API every 60 seconds
- Outputs debug information and readings to the Serial Monitor

---

## Hardware Required

- ESP32 development board  
- YL-69 soil moisture sensor  
- Jumper wires  
- USB Type-C cable to connect ESP32 to your computer

## Wiring

| YL-69 Sensor Pin    | ESP32 Pin                   |
|---------------------|-----------------------------|
| VCC                 | 3.3V or 5V                  |
| GND                 | GND                         |
| Analog Output (AO)  | GPIO 34 (Analog input pin)  |

---

🏗️ RUNNING & BUILDING

Command	Description:

    go run file.go	Runs a Go file directly (interpreted-like).

    go build	Compiles the current Go package into a binary.
    
    go build file.go	Builds a specific file.
    
    go install	Compiles and installs the package to $GOPATH/bin.
        
    go clean	Removes build files.

📦 GO MODULES (DEPENDENCY MANAGEMENT)

These commands assume you're using Go Modules (go mod) — which is standard now.
    Command	Description:

    go mod init module/name	Initializes a new Go module (like a project).
   
    go mod tidy	Adds missing and removes unused dependencies.
   
    go get pkg/path	Adds a new dependency or updates it.
   
    go list -m all	Lists all modules and their versions.
   
    go mod download	Downloads modules to local cache.



## To Do: 
- Create frontend
- Define measuring Logic
- Set a more readable value for the measurement

