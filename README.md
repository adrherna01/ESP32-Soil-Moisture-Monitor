# ESP32 Soil Moisture Monitor with Pushover Notifications

This project uses an ESP32 microcontroller connected to a YL-69 soil moisture sensor to monitor the moisture level of your plant's soil. It connects to WiFi and sends notifications via [Pushover](https://pushover.net/) when it reads the current moisture level, helping you know when your plant needs watering.

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

