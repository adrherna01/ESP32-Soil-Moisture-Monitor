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
}

void loop() {
  int moistureValue = analogRead(sensorPin);
  Serial.print("Moisture: ");
  Serial.println(moistureValue);

  if (WiFi.status() == WL_CONNECTED) {
    HTTPClient http;
    http.begin("https://api.pushover.net/1/messages.json");
    http.addHeader("Content-Type", "application/x-www-form-urlencoded");

    String payload = "token=" + String(TOKEN) +
                     "&user=" + String(USER_KEY) +
                     "&message=Moisture+level+is+" + String(moistureValue);

    Serial.println(payload);

    int httpResponseCode = http.POST(payload);

    if (httpResponseCode > 0) {
      Serial.println("Notification sent");
    } else {
      Serial.print("Error sending notification: ");
      Serial.println(httpResponseCode);
    }

    http.end();
  } else {
    Serial.println("WiFi not connected");
  }

  // sending every 60 seconds
  delay(60000); 
}