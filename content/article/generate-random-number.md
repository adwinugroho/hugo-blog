---
title: "Golang - Generate Random Number"
date: 2021-07-01T18:44:58+08:00
draft: false

categories: ["Golang", "Tutorial"]
tags: ["golang basic",  "golang tutorial", "generate random number", "membuat angka acak"]
toc: false
author: "Adwin"
---
Biasanya _generate_ angka random dipakai untuk membuat _passcode_, ID (biasanya campuran angka dan huruf), dll. Di artikel ini kita hanya membuat angka acak dengan panjang angka sesuai sama yang kita inginkan. Di golang ada _built in_ _package_ **math/rand** untuk membuat data acak. 

Basis dari _package_ tersebut ialah PRNG atau _pseudo-random number generator_. Deret angka _random_ yang dihasilkan sangat tergantung dengan angka _seed_ yang digunakan. Oleh karena itu disini kita menggunkan _package_ **time** untuk memanggil fungsi waktu sekarang dengan format sampai _nano-second_ untuk mendapatkan _random_ angka yang tak mudah ditebak/berulang.

```go
import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber(lengthNumber int) int {
	var charset = "0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, lengthNumber)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	byteToInt, _ := strconv.Atoi(string(b))
	return byteToInt
}
```
Catatan: untuk nomor acak yang sensitif terhadap keamanan, gunakan _packge_ **crypto/rand** sebagai gantinya.