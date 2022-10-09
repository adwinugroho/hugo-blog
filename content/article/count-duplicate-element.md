---
title: "Golang - Count Duplicate Element"
date: 2021-07-20T22:04:09+08:00
draft: false

categories: ["Golang", "Tutorial"]
tags: ["golang basic",  "golang tutorial", "count duplicate element", "menghitung elemen duplikat"]
toc: false
author: "Adwin"
---

Di golang ada sebuah tipe data bernama *map* sama seperti *dictionary* bila di python atau bila di javascript tentunya kita sudah sangat akrab, yakni objects. Which is ketiganya mirip yaitu sama-sama bisa menyimpan banyak value dan sama-sama terdapat 2 buah value data, *key* dan *value*.

Di artikel ini kita akan membahas tentang berapa *value* yang sama, yang ada pada tiap *element* di *array*. Ini kalau tidak salah saya disuruh untuk mengembalikan *response* harga obat beserta nama dan jumlahnya dimana data yang kita dapat berupa *array* 

Eh sebentar, di golang kita sebutnya *array* atau *slice* ya, atau *list* hehehe 😂😂
```go
// CountDuplicateElement count how much duplicate value in array
// in example we use array of string
// ex we have an array like this [hello, how, are, you, today, you, fine]
// result will gonna be hello: 1, how: 1, are: 1 today: 1 and you: 2
func CountDuplicateElement(list []string) map[string]int {
	duplicateElement := make(map[string]int)
	// looping list
	for _, item := range list {
		// check if keys found in map
		_, exist := duplicateElement[item]
		if exist {
			duplicateElement[item] += 1 // iteration element
		} else {
			duplicateElement[item] = 1
		}
	}
	return duplicateElement
}
```