#include <WiFi.h>
#include <WebServer.h>
#include <HTTPClient.h>
#include "credentials.h"

const int sensorPin = 34;
const int ledPin = 2;

WebServer server(80);

bool isBackConnected() {
  String ip = WiFi.localIP().toString();
  HTTPClient http;

    String url = String("http://") + HOST_IP + ":" + String(HOST_PORT) + "/register";
    Serial.println(url);

    http.begin(url);
    http.addHeader("Content-Type", "application/json");

    String jsonData = "{\"ip\":\"" + ip + "\",\"port\":\"" + String(PORT) + "\"}";

    int httpResponseCode = http.POST(jsonData);

    if (httpResponseCode != -1) {
      Serial.print("Http response code: ");
      Serial.println(String(httpResponseCode));
    }

    if (httpResponseCode == 200) {
      http.end();
      return true;
    }
    http.end();
    return false;
}

void handleMeasure() {
  int moistureValue = analogRead(sensorPin);

  if (WiFi.status() == WL_CONNECTED) {
    HTTPClient http;
    String url = String("http://") + HOST_IP + ":" + String(HOST_PORT) + "/measurements";
    
    // Serial.println(url);

    http.begin(url);
    http.addHeader("Content-Type", "application/json");
    String jsonData = "{\"humidity\":" + String(moistureValue) + "}";
    int httpResponseCode = http.POST(jsonData);
    http.end();
  }

  // Respond to the requester
  String json = "{\"humidity\": " + String(moistureValue) + "}";
  server.send(200, "application/json", json);
}

void setup() {
  Serial.begin(115200);
  WiFi.begin(SSID, PASS);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
  }

  Serial.print("WiFi connected \n");
  // Serial.print("IP address: ");
  // Serial.println(WiFi.localIP());

  while (!isBackConnected())
    delay(5000);

  server.on("/measure", HTTP_GET, handleMeasure);
  server.begin();
}

void loop() {
  server.handleClient();
}

// Logear ip al backend asi el backend luego puede hacer un call all api
// de el esp32.