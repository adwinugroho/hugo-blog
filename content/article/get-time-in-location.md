---
title: "Golang - Get Time based on Location"
date: 2021-07-20T22:03:09+08:00
draft: false

categories: ["Golang", "Tutorial"]
tags: ["golang basic",  "golang tutorial", "time based on location", "waktu berdasarkan lokasi"]
toc: false
author: "adwin"
---
Di golang terdapat *built in package* **time** dimana kita bisa mendapatkan/memanipulasi waktu yang diperlukan, lengkapnya [time](https://pkg.go.dev/time). Untuk mendapatkan *format* waktu sesuai sama lokasi dimana kita berada, pertama panggil fungsi *time.LoadLocation("Asia/Jakarta")*.

Sebagai tambahan kita panggil fungsi *time.Now.In*(dengan parameter *variable location*) yang telah kita definisikan sebelumnya, untuk mendapatkan waktu saat ini/waktu coding ini dijalankan.  
```go
import (
	"log"
	"time"
)

func TimeHostNow() time.Time {
	// you can change Asia/Jakarta with your own location.
	// check on this https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Printf("Error get time, cause:%+v\n", err)
	}
	nowTimeInLoc := time.Now().In(location)
	return nowTimeInLoc
}
```
Dengan *return value* fungsi di atas kita mendapatkan waktu (*time.Time*) sesuai sama lokasi yang kita inginkan, agar mendapatkan hasil yang lebih maksimal silakan baca: [format time or date](https://yourbasic.org/golang/format-parse-string-time-date-example/) di golang.