---
title: "The Story About Si Subur"
date: 2024-09-09T14:06:50+07:00
draft: false

categories: ["Random", "Story"]
tags: ["random",  "story", "cerita", "project"]
toc: false
author: "Adwin"
---
Assalamualaikum, Bismillah, hola, moshi-moshi.
Observing changes experiences by the body during menstruation will provide important clues about fertility status. One of the most effective and easy to implement methods is to control the body's basal temperature. Basically, basal body temperature (BBT) is a condition where the temperature depends on the lowest level. The future influx is marked by a slight decrease in temperature, followed by a sudden upward rise of 0.2 °C (0.4 °F) for the next 48 hours. Of course the fertile cycle is included in the time series data because the data are interconnected with one another in sequence, arranged from one object consisting of several time periods [^1].

I have an idea from my lecture (Ibu Diah Risqiwati) that we should make a system or tool to close how to know fertility from BBT. She is gave a bunch of journals to me to read. Basically, I have interest to on healthcare technology since I was have a disablity too. My right foot doesn't normal and can't walk properly like a normal human. I have another story about this and if you wanna know more of me, just click [about me!](https://blog.cryppy.id/page/about/). Let's continue into this article, before my lecture gave me this idea, I already tell her my idea and I want to build something technology Internet of Things (IoT) about Neurological Disorder in children but Mrs. Diah thought it will be harder and more pricey because the censor.

Finally after a long discussion with another my lecture (Ibu Lailatul Husniah), they tell me to build a system to observing and predicted fertility based on BBT. The implementation involved two device, first is microcontroller with ESP8266 WiFi and infrared censor (to know the temperature), and second is Web to user login and monitoring their BBT. This is a system called `Si Subur`. What does `Si Subur` mean? In Indonesia, "subur" means fertile, lush, or productive (synonyms include gemuk, gendut, etc.). The "fertility" I want to talk about in this article refers to human fertility.
![Subur based on KBBI](/image/content/article/the-story-about-si-subur/image-1.png)

In English, fertile can mean (of a seed or egg) capable of becoming a new individual, or (of a person, animal, or plant) able to conceive young or produce seed. We Indonesians love naming systems that start with "SI" when they are related to information systems, like SIDesa, SIAkad, SImpatika, or whatever name you can think of. Now, back when I was in college, I often joined competitions like PKM, Gemastik, etc. Initially, I had my own ideas, such as building an IoT device to prevent taxi drivers from getting tired and falling asleep or creating a hat for kids to detect brain disorders. Unfortunately, my lecturer thought those ideas were too difficult to implement, so she gave me a new one the idea of monitoring the menstrual cycle and women's fertility.

What makes this different from regular menstrual monitoring apps? The difference is in the device. I built a device similar to a thermometer, but the body temperature readings are automatically displayed on the system I developed. So, `Si Subur` essentially consists of two interconnected devices.
