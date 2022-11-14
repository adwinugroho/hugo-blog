---
title: "Golang - Add New Element Without Redudancy"
date: 2022-10-15T22:24:18+07:00
draft: false

categories: ["Golang", "Tutorial"]
tags: ["golang basic",  "golang tutorial", "adding element without redudancy", "menambahkan elemen tanpa duplikat"]
toc: false
author: "Adwin"
---
Ga kayak di javascript (JS) kita bisa memapulasi data banyak dengan *library* bawaan. Entah itu buat *searching element* di *array, remove element* di *array* atau bahkan sekadar me-*reverse* kita bisa dengan hanya memanggil *method*-nya saja. 

note: di golang ada yang menyebut *array* itu *slice* atau mungkin keduanya dibedakan, hanya saja saya tidak terbiasa dengan hal itu. Bila kamu ingin mendalami lebih lanjut sebenernya emang *array* dan *slice* di golang itu berbeda.

Di golang juga banyak kok *library* bawaan *unfortunately* di golang tuh salah satu yang paling penting dan pasti dipakai apabila kita banyak bermain dengan data yaitu *remove array element* ga ada. Ada banyak cara sebenernya nge-*remove array element* di golang salah satunya berikut ini:
```go
// AddNewElementWithoutRedudancy always return newData in last index
func AddNewElementWithoutRedudancy(datas []string, newData string) []string {
	var isNew bool
	if len(datas) == 0 {
		isNew = true
	} else {
		for i := range datas {
			if datas[i] == newData {
				datas = append(datas[:i], datas[i+1:]...)
				datas = append(datas, newData)
				isNew = false
				break
			} else {
				isNew = true
			}
		}
	}
	if isNew {
		datas = append(datas, newData)
	}
	return datas
}
```
*remove element*-nya sebenernya hanya pada line 22, hanya saja pada contoh fungsi di atas, kita bisa memasukkan *element* baru tanpa adanya *duplicate/redudancy* data.
