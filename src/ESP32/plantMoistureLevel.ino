#include <WiFi.h>
#include <HTTPClient.h>
#include "credentials.h"

const int sensorPin = 34;
const int ledPin = 2;

void setup() {
  Serial.begin(115200);
  WiFi.begin(SSID, PASS);

  Serial.print("Connecting to WiFi");
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("\nWiFi connected");

  pinMode(ledPin, OUTPUT);
}

void loop() {
  int moistureValue = analogRead(sensorPin);
  digitalWrite(ledPin, HIGH);
  delay(1000);
  digitalWrite(ledPin, LOW);
  Serial.print("Moisture: ");
  Serial.println(moistureValue);

  if (WiFi.status() == WL_CONNECTED) {
    HTTPClient http;

    String url = String("http://") + HOST_IP + ":" + String(PORT) + "/measurements";
    http.begin(url);
    http.addHeader("Content-Type", "application/json");

    String jsonData = "{\"humidity\":" + String(moistureValue) + "}";

    Serial.println(jsonData);

    int httpResponseCode = http.POST(jsonData);

    if (httpResponseCode > 0) {
        String response = http.getString();
        Serial.println("Server response: " + response);
      } else {
          Serial.print("Error sending notification: ");
          Serial.println(httpResponseCode);
        }
    http.end();
  } else {
    Serial.println("WiFi not connected");
  }

  delay(10000); // send every 10 seconds
}