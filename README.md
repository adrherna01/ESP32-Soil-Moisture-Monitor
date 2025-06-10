# ESP32 Soil Moisture Monitor with Pushover Notifications

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

