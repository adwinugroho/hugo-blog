---
title: "Golang - Julian Date"
date: 2023-05-24T19:25:42+08:00
draft: false

categories: ["Golang", "Tutorial"]
tags: ["golang basic",  "golang tutorial", "date format", "julian-date"]
toc: false
author: "Adwin"
---
Julian Date? Adalah sebuah tanggal dimana hitungan hari terus menerus yang dimulai dari periodenya, disini saya mencontoh sebuah situs web yang meng-*generate* julian date, ["klik disini"](https://www.longpelaexpertise.com.au/toolsJulian.php). Di situs web tersebut menggunakan format yy-dddd dan dengan hitungan pertahun jadai agak sedikit berbeda jika kita merujuk pada julian date di wiki.

```go
func GetJulianDays(nowTime string) string {
	var layout = "2006-01-02"
	getDate, _ := time.Parse(layout, nowTime)
	var yearString = nowTime[0:4]
	var isLeapYear bool
	var numberOfDays = 366
	
    getIntYear, _ := strconv.Atoi(yearString)
	if (getIntYear%4 == 0 && getIntYear%100 != 0) || getIntYear%400 == 0 {
		isLeapYear = true
	}

	startedTime, _ := time.Parse(layout, fmt.Sprintf("%s-01-01", yearString))
	plusOneYear := startedTime.AddDate(1, 0, 0)
	days := plusOneYear.Sub(getDate).Hours() / 24
	if isLeapYear {
		numberOfDays = 367
	}
	
    var daysToString = strconv.Itoa(int(days))
	if len(daysToString) > 2 {
		return fmt.Sprintf("%s%s", yearString[2:], daysToString)
	}
	return fmt.Sprintf("%s0%s", yearString[2:], daysToString)
}
```