---
title: "The Story About Si Subur An IoT Journey into Fertility Tracking"
date: 2024-09-09T14:06:50+07:00
draft: false

categories: ["Random", "Story"]
tags: ["random",  "story", "cerita", "project"]
toc: false
author: "Adwin"
---
**Last edit: 9 November 2026**

Assalamualaikum, Bismillah, Hola, Moshi-moshi.
Observing bodily changes during the menstrual cycle provides critical clues about fertility status. One of the most effective and accessible methods for tracking this is monitoring Basal Body Temperature (BBT). Biologically, the fertile window is often marked by a slight dip in temperature followed by a sustained rise of approximately 0.2°C (0.4°F) lasting at least 48 hours. Since these data points are sequential and interdependent, fertility tracking is essentially a **time-series analysis** problem[^1].

### The Spark: From Neurology to Fertility

This project began with an idea from my lecturer, Ibu Diah Risqiwati. She encouraged me to build a system that could help women understand their fertility through BBT, providing me with numerous journals to study. 

I have had a keen interest in healthcare technology for a long time, partly due to my own physical disability—my right foot doesn't function normally, which affects how I walk. If you want to know more about my personal journey, feel free to check out my [About Me](https://blog.cryppy.id/page/about/) page.

Before settling on `Si Subur`, I actually pitched a different idea to Ibu Diah: an IoT device for detecting neurological disorders in children. However, after a long discussion with her and another lecturer, Ibu Lailatul Husniah, we concluded that the sensors required would be too expensive and complex for a student project. They steered me toward a more feasible yet impactful alternative: **monitoring the menstrual cycle and predicting fertility using BBT.**

They also introduced me to Bapak Wahyu, who became instrumental in helping me implement the hardware side of this project.

### What is "Si Subur"?

In Indonesian, *"subur"* means fertile, lush, or productive. We Indonesians have a habit of prefixing information systems with "SI" (Sistem Informasi)—think SIDesa, SIAkad, or SImpatika. So, naturally, this became **Si Subur**. Unlike standard period-tracking apps that rely solely on manual calendar inputs, `Si Subur` bridges the physical and digital worlds. It consists of two interconnected components:
1.  **Hardware:** A custom IoT thermometer using an ESP8266 and an MLX90614 infrared sensor.
2.  **Software:** A web application for user management, data visualization, and fertility prediction.

![Definition of Subur from KBBI](/image/content/article/the-story-about-si-subur/image-1.png)

### The Hardware: Bringing BBT to the Cloud

The core of the device is the **ESP8266** microcontroller paired with an **MLX90614** non-contact infrared temperature sensor. 

Initially, I prototyped this by creating a local Access Point where users could connect directly to the device to see their temperature. However, to make it a true IoT solution, the data needed to live in the cloud.

I chose **ThingSpeak** as the backend because it specializes in time-series data aggregation and visualization. The ESP8266 reads the object temperature (body heat) and pushes it to a ThingSpeak channel via MQTT/HTTP every 15 seconds—compliant with their free tier rate limits.

Here is a simplified look at how the ESP8266 handles the sensor reading and WiFi connection:

```cpp
#include <ESP8266WiFi.h>
#include "ThingSpeak.h" // Library for ThingSpeak integration
#include <Wire.h>
#include <Adafruit_MLX90614.h>

// Sensor & WiFi Setup
Adafruit_MLX90614 mlx = Adafruit_MLX90614();
const char* ssid = "YourWiFiSSID";
const char* password = "YourWiFiPassword";
unsigned long myChannelNumber = YOUR_CHANNEL_ID;
const char * myWriteAPIKey = "YOUR_API_KEY";

WiFiClient client;

void setup() {
  Serial.begin(115200);
  Wire.begin(4, 5); // SDA, SCL for MLX90614
  
  // Connect to WiFi (Station Mode for Cloud Access)
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  
  ThingSpeak.begin(client);
  Serial.println("Connected to WiFi & ThingSpeak");
}

void loop() {
  // Read Object Temperature (Body Temp)
  float tempC = mlx.readObjectTempC();
  
  // Send to ThingSpeak Field 1
  ThingSpeak.setField(1, tempC);
  int x = ThingSpeak.writeFields(myChannelNumber, myWriteAPIKey);
  
  if(x == 200){
    Serial.println("Channel update successful.");
  } else {
    Serial.println("Problem updating channel. HTTP error code " + String(x));
  }
  
  delay(20000); // ThingSpeak requires ~15s delay between writes
}
```

### The Software: A Complete User Experience

With the data flowing into the cloud, I built the web application using PHP 5.6 and CodeIgniter. The goal was to create a system that wasn't just a dashboard, but a supportive platform for women. 

1. The Public Portal
The main page (si-subur.com) acts as a content hub. It features articles on pre-pregnancy, post-partum care, and general health, serving as an educational resource alongside the tool itself.

![Main page of si-subur.com](/image/content/article/the-story-about-si-subur/image-2.png)

2. Authentication & Profile
Security and personalization are key. Users register with an email and password. The profile page allows them to manage their credentials and personal details, ensuring their health data remains private.

![Login page of si-subur](/image/content/article/the-story-about-si-subur/image-4.png)

![Profile page of si-subur](/image/content/article/the-story-about-si-subur/image-5.png)

3. The Dashboard: Visualizing Fertility
This is the heart of the system. Once logged in, users are greeted with a dashboard that pulls live data from ThingSpeak.
- **Cycle Statistics:** A graphical representation of BBT over time.
- **Ovulation Prediction:** The system analyzes the time-series data to highlight potential ovulation days (e.g., "Hari Ovulasi Di Siklus Berikutnya").
- **Data Entry:** While the IoT device automates collection, I also included a manual input modal for users to add historical data or correct readings.

![Login page of si-subur](/image/content/article/the-story-about-si-subur/image-3.png)

![Profile page of si-subur](/image/content/article/the-story-about-si-subur/image-6.png)

### Conclusion
Si Subur was more than just a college assignment; it was a lesson in integrating hardware constraints with software usability. By combining an affordable IR sensor with cloud analytics, we created a tool that makes fertility awareness accessible and data-driven. While the tech stack (PHP 5.6, ESP8266) might seem dated now, the core principle remains relevant: technology should simplify health monitoring, not complicate it.


References:
[^1]: https://www.researchgate.net/publication/348802419_PREDIKSI_MASA_SUBUR_PADA_WANITA_MENGGUNAKAN_PEMODELAN_STATE_SPACE_BERBASIS_SUHU_BASAL_TUBUH/fulltext/601147d792851c2d4df73653/PREDIKSI-MASA-SUBUR-PADA-WANITA-MENGGUNAKAN-PEMODELAN-STATE-SPACE-BERBASIS-SUHU-BASAL-TUBUH.pdf
