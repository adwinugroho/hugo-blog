---
title: "A Story About PPMS"
date: 2026-09-21T12:54:55+07:00
draft: true
---
Portable Patient Monitory System a.k.a PPMS is a my lectures project for IEEE funding. The agremeent project between Institut Teknologi Sepuluh Nopember (ITS) and Hospital University of Muhammadiyah Malang (RS UMM). 

I was assigned as a web developer who integrated the hardware (raspberyPI) of the system into a web application. The system has 4 sensors : ECG, SpO2, Temperature, Blood Pressure. I build this system with Python Django and use free template dashboard for the admin & auth page. This project is my very first real project that use RaspberryPI. At that time, I only familiar with web development, but I never touch RaspberryPI before. My supervisor recommend me to try this project, because it will expand my skills. 

One of the hardest part is building TCP Socket Client to listen to the data stream from RaspberryPI. Because at that time, I never touch TCP Socket before. I need to learn it from scratch. So, PPMS used Berry P6100 to get stream vital sign data from patient such a SpO2, Heartrate, Systole, Diastole, Respiratory, and Temperature.

![PPMS - Alat P6100](/image/content/article/the-story-about-ppms/image-1.png)

The communication between raspberryPI and P6100 using bluetooth. The patient is wearing the sensors and the data is sent to the doctor's web application. The doctor can monitor the patient's vitals in real-time. The doctor can also set the patient's vitals to a certain level and if the patient's vitals go above or below the set level, the doctor will be notified.

![PPMS - Login Page](/image/content/article/the-story-about-ppms/image-2.png)

However, on this project my task is simple, I have to develop retriving data from raspberyPI into python and display that on my web application. So, user need to know the mac address bluetooth then have to login with name, mac address, port, and license number (if they are with ambulance).
On example below my python code to gather data from.

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
