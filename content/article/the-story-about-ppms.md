---
title: "The Story About PPMS"
date: 2026-09-21T12:54:55+07:00
draft: false

categories: ["Random", "Story"]
tags: ["random",  "story", "cerita", "project"]
toc: false
author: "Adwin"
---
Portable Patient Monitoring System, a.k.a PPMS, was my lecturer’s project funded by IEEE. It was a collaborative agreement between Institut Teknologi Sepuluh Nopember (ITS) and Hospital University of Muhammadiyah Malang (RS UMM).

I was assigned as the web developer responsible for integrating the hardware (Raspberry Pi) into a web application. The system utilized four sensors: ECG, SpO2, Temperature, and Blood Pressure. I built this system using Python Django and used a free dashboard template for the admin and auth pages. This project was my very first real experience with Raspberry Pi. At that time, I was only familiar with web development and had never touched a Raspberry Pi before. My supervisor recommended this project to me because it would expand my skill set.

One of the hardest parts was building a TCP Socket Client to listen to the data stream from the Raspberry Pi, mostly because I had never worked with TCP Sockets before. I had to learn it from scratch. The PPMS used a Berry P6100 device to get vital sign data streams from patients, such as SpO2, Heart Rate, Systole, Diastole, Respiratory Rate, and Temperature.

![PPMS - Alat P6100](/image/content/article/the-story-about-ppms/image-1.png)

The communication between the Raspberry Pi and the P6100 used Bluetooth. The patient wears the sensors, and the data is sent to the health worker's web application. Doctors can monitor the patient's vitals in real-time. They can also set threshold levels for the patient's vitals; if the values go above or below these set levels, the emergency doctor gets notified.

![PPMS - Login Page](/image/content/article/the-story-about-ppms/image-2.png)

My task in this project was straightforward: retrieve data from the Raspberry Pi using Python and display it on the web application. Users needed to know the Bluetooth MAC address and log in with their name, MAC address, port, and license number (if they were with an ambulance).
Below is an example of my Python code for gathering the data:

```python
import bluetooth
import time
import datetime

# address & port will be inputs by staff
bd_addr = "34:81:F4:6B:F1:01"
port = 6
sock=bluetooth.BluetoothSocket( bluetooth.RFCOMM )
sock.connect((bd_addr, port))
datetime_file = str(datetime.datetime.now().strftime('%Y%m%d_%H%M'))
print('Connected to ', bd_addr)

#define variabel
temp = '0'
spo = '0'
resp = '0'
hr = '0'

with open('/Documents/PPMS/'+datetime_file+'.txt', 'a') as f:
    while True:
        a = sock.recv(1024)
        get_time_formatted = datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        c = '' 
        for i in range(len(a)):
            c = c + ' ' + str(a[i])
        f.writelines('\n')
        f.writelines(get_time_formatted + ' ' +str(len(a)) +' ' + c)
        
        #grab data
        index = len(a)
        if index is 6:
            if a[5] in range (232,238):
                resp = str(a[4])
        elif index is 8:
            temp = str(a[5]) + '.' + str(a[6])
        elif index is 16:
            spo = str(a[13])
            hr = str(a[14])
        
        #cross check data
        if float(temp) < 25 or float(temp) > 40:
            temp = '0'
        
        if float(spo) > 100:
            spo = '0'
        
        if float(hr) > 100:
            hr = '0'
        
        if float(resp) > 25:
            resp = '0'
        print(get_time_formatted, temp, spo, hr, resp)
        
sock.close()
f.close()
```

If you successfully log in, here is the view of the dashboard. As you can see on the main dashboard, there’s the RS UMM profile. Then, you can input the patient's medical record along with the static vehicle license number from the login session.

![PPMS - Dashboard Page](/image/content/article/the-story-about-ppms/image-3.png)

Finally, PPMS has an optional feature to print the data. We believed that hospitals (RS) and healthcare workers (nakes) must have physical records. It’s important for humanizing the process and improving user experience.
At the end of the day, I think systems like PPMS need to rise up, especially in the age of AI. We should focus on doing more important things in healthcare and solving many problems related to diseases. Hopefully!

Thanks to:
- Allah Azza wa Jalla
- Ibu Diah Risqiwati, ST., MT
- Bapak Hanugra Aulia Sidharta, ST., MT
- All of Civitas Academica.
