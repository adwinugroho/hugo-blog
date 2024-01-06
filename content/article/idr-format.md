---
title: "Golang - Formating Float to Currency (IDR) String"
date: 2024-01-07T01:07:10+08:00
draft: false

categories: ["Golang", "Tutorial"]
tags: ["golang basic",  "golang tutorial", "formatting float to currency idr string", "format mata uang dari float ke string"]
toc: false
author: "Adwin"
---
Di golang ada *library* ~~bawaan~~ yang bisa dipakai untuk *localized formatting* ataupun *translation* entah kita ingin menerjemahkan huruf yang bukan alfabet atupun hanya sekadar sebagai *translate*.

*note*: bila ada *library* dengan *package* sebagai berikut https://golang.org/pkg/#subrepo mereka adalah *package* yang ada di dalam *Go Project* tapi di luar dari *main project*

Di artikel ini kita akan membuat format mata uang rupiah daripada harus pake *third party* macam *humanize* mendingan buat sendiri pake *built in*.
```go
func IDRCurrencyFormatted(amount float64) string {
	p := message.NewPrinter(language.Indonesian)
	return p.Sprintf("%.2f\n", amount) // hanya ambil dua angka dibelakang koma
}
```



