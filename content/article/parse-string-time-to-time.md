---
title: "Golang - Parse String Time to Time"
date: 2022-11-18T15:27:24+07:00
draft: false

categories: ["Golang", "Tutorial"]
tags: ["golang basic",  "golang tutorial", "convert time string to time golang", "konversi waktu string ke time golang"]
toc: false
author: "adwin"
---
Langsung *to the point* aja, ketika menemukan format tanggal dalam bentuk *string* dan ingin mengolahnya tentunya kita perlu *parsing* tanggal tersebut agar menjadi bentuk *time.Time* di golang.
```go
func parsingTimeStringToTime(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, errors.New("invalid date")
	}

	// fit the layout
	// see here default layout from go time.Time: https://pkg.go.dev/time#Constants
	// i.e date = "2023-01-25 15:04:05"
	layout := "2006-01-02 15:04:05"
	convDateStr, err := time.Parse(layout, date)
	if err != nil {
		log.Printf("[parsingTimeStringToTime] Error while parse time%+v\n:", err)
		return time.Time{}, err
	}
	// manipulation date as you like
	convDateStr = convDateStr.AddDate(+1, 0, +12)
	return convDateStr, nil
}
```
