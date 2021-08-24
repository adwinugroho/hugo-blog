---
title: "Golang - Calculate Date By Ages"
date: 2021-07-18T22:11:09+08:00
draft: false

categories: ["Golang", "Tutorial"]
tags: ["golang basic",  "golang tutorial",]
toc: false
author: "Adwin"
---
Sebelum kesini ada baiknya untuk liat ["artikel ini"](https://blog.cryppy.xyz/get-time-in-location), karena disini akan memakai fungsi yang sama seperti artikel tersebut untuk mendapatkan waktu berdasarkan lokasi yang diinginkan.

Ketika memasukkan umur, maka secara otomatis mengembalikan string tanggal sekarang tetapi tahunnya adalah tahun kelahiran kamu. Tambahkan 2 parameter integer di fungsi untuk bulan dan hari agar bisa mendapatkan tanggal yang sesuai dengan yang kamu mau.

```go
import (
	"time"
)

func CalcDate(ages int) string {
	if ages == 0 {
		return ""
	}
	now := TimeHostNow() // get now time based on location https://blog.cryppy.xyz/get-time-in-location
	ages = (ages * -1)
	result := now.AddDate(ages, 0, 0)
	return result.Format("20060102")
}
```