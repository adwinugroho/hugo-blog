---
title: "The Story About Si Subur"
date: 2024-09-09T14:06:50+07:00
draft: true

categories: ["Random", "Story"]
tags: ["random",  "story", "cerita", "project"]
toc: false
author: "Adwin"
---
**Last edit: 8 November 2026**

Assalamualaikum, Bismillah, hola, moshi-moshi.
Observing changes experiences by the body during menstruation will provide important clues about fertility status. One of the most effective and easy to implement methods is to control the body's basal temperature. Basically, basal body temperature (BBT) is a condition where the temperature depends on the lowest level. The future influx is marked by a slight decrease in temperature, followed by a sudden upward rise of 0.2 °C (0.4 °F) for the next 48 hours. Of course the fertile cycle is included in the time series data because the data are interconnected with one another in sequence, arranged from one object consisting of several time periods [^1].

I have an idea from my lecture (Ibu Diah Risqiwati) that we should make a system or tool to close how to know fertility from BBT. She is gave a bunch of journals to me to read. Basically, I have interest to on healthcare technology since I was have a disablity too. My right foot doesn't normal and can't walk properly like a normal human. I have another story about this and if you wanna know more of me, just click [about me!](https://blog.cryppy.id/page/about/). Let's continue into this article, before my lecture gave me this idea, I already tell her my idea and I want to build something technology Internet of Things (IoT) about Neurological Disorder in children but Mrs. Diah thought it will be harder and more pricey because the censor.

Finally after a long discussion with another my lecture (Ibu Lailatul Husniah), they tell me to build a system to observing and predicted fertility based on BBT. The implementation involved two device, first is microcontroller with ESP8266 WiFi and infrared censor (to know the temperature), and second is Web to user login and monitoring their BBT. This is a system called `Si Subur`. What does `Si Subur` mean? In Indonesia, "subur" means fertile, lush, or productive (synonyms include gemuk, gendut, etc.). The "fertility" I want to talk about in this article refers to human fertility.
![Subur based on KBBI](/image/content/article/the-story-about-si-subur/image-1.png)

In English, fertile can mean (of a seed or egg) capable of becoming a new individual, or (of a person, animal, or plant) able to conceive young or produce seed. We Indonesians love naming systems that start with "SI" when they are related to information systems, like SIDesa, SIAkad, SImpatika, or whatever name you can think of. Now, back when I was in college, I often joined competitions like PKM, Gemastik, etc. Initially, I had my own ideas, such as building an IoT device to prevent taxi drivers from getting tired and falling asleep or creating a hat for kids to detect brain disorders. Unfortunately, my lecturer thought those ideas were too difficult to implement, so she gave me a new one the idea of monitoring the menstrual cycle and women's fertility.

What makes this different from regular menstrual monitoring apps? The difference is in the device. I built a device similar to a thermometer, but the body temperature readings are automatically displayed on the system I developed. So, `Si Subur` essentially consists of two interconnected devices. Ok here is the implementation on microcontroller which is ESP8266 included WiFi communicate with thingspeak channel MQTT to get data

```c
// Create a WiFi access point and provide a web server on it.
// Go to http://192.168.4.1 in a web browser connected to this access point to see it.
// Control LED connected to  GPIO2

#include <ESP8266WiFi.h>
#include <WiFiClient.h> 
#include <ESP8266WebServer.h>
#include <Wire.h>
#include <Adafruit_MLX90614.h>
Adafruit_MLX90614 mlx = Adafruit_MLX90614();

int val;
int i;
// Set these to your desired AP credentials.
const char *ssid = "masasubur";                // setting ssid name of our ESP
const char *password = "masasubur";           // setting password

ESP8266WebServer server(80);

void setup() {
  Wire.begin(4, 5);                            // 4 = sda, 5 = scl setting input i2c komunikasi dari sensor MLX (sda scl)
	delay(1000);
	Serial.begin(115200);                        // setting baudrate
       
       
	Serial.println();
	Serial.print("Configuring access point...");
	WiFi.softAP(ssid);                           // Remove password parameter if you want the AP to be open

	IPAddress myIP = WiFi.softAPIP();
	Serial.print("AP IP address: ");
	Serial.println(myIP);
 
	server.begin();
	Serial.println("HTTP server started");
}

void loop() {
 
 server.handleClient();
        
// send data to serial monitor------- 

 Serial.print("Ambient = "); Serial.print(mlx.readAmbientTempC()); 
 Serial.print("*C\tObject = "); Serial.print(mlx.readObjectTempC()); Serial.println("*C");
 Serial.print("Ambient = "); Serial.print(mlx.readAmbientTempF()); 
 Serial.print("*F\tObject = "); Serial.print(mlx.readObjectTempF()); Serial.println("*F");
 
 val=mlx.readObjectTempC(); // GET the calculate of temperature
 Serial.println();
 
// data server

 server.on("/", [](){  // if you add this subdirectory to your webserver call, you get text below :)
   String webString="";     // String to display
   //webString="<FONT SIZE=+4> ESP8266 Wi-Fi webserver suhu dalam Celcius <BR> ";
   webString +=(val);
   
   server.send(200, "text/html", webString);            // send to someones browser when asked
   
 });
delay(500);


}
```
### Bring It to the Cloud: ThingSpeak Integration

To make `Si Subur` a true Internet of Things (IoT) device, the temperature data captured by the ESP8266 needs to be sent to a cloud platform for remote monitoring and data logging. For this, I utilized **ThingSpeak**, an IoT analytics platform service that allows us to aggregate, visualize, and analyze live data streams.

While the early prototype relied on a local web server via an Access Point (`WiFi.softAP`), sending data to the cloud requires an active internet connection. Therefore, the ESP8266 configuration must be switched to Station mode (`WiFi.begin`) to connect to a local Wi-Fi router.

The integration is simplified using the official `ThingSpeak.h` library. The logic is straightforward: the ESP8266 reads the Basal Body Temperature (BBT) from the MLX90614 sensor, serves the real-time value on a local web page, and securely publishes the data to a specific ThingSpeak Channel Field using a unique `Write API Key` every 15 seconds. This interval complies with ThingSpeak's rate limits for free accounts while providing sequential time-series data sufficient for tracking the user's fertility cycle. Thanks to Bapak Wahya who is my another lecture to helping me a lot to implment this device (`Si Subur`)

### Implementation to Web Application

The web application began with a page like a news application or blog, something like that. I built that with PHP 5.6 with codeigniter framework. It serve a content management server where admin can write a news or user experiences about the tools (`Si Subur`). 
![si-subur.com main page](/image/content/article/the-story-about-si-subur/image-2.png)

Since we are successfully get data from the ESP8266 to Thingspeak, we are going to use Thingspeak API to implement that data to web application. The users is a woman who use `Si Subur` which already have an email (username) and password to login.
![Subur based on KBBI](/image/content/article/the-story-about-si-subur/image-1.png)
