---
title: "Golang Unit Test"
date: 2021-02-23T14:10:47+08:00
draft: true

categories: []
tags: []
toc: false
author: "Adwin"
---
1. Pengertian unit test
definisi, macam-macam test dan kenapa harus test

2. Arsitektur
flow unit test

3. Keuntungan dan kelebihan
pros and cons

4. Testing package di bahasa program
- contoh library unit testing untuk bahasa lain
- kalau golang langsung disediakan

5. Golang Unit test
- golang menyediakan struct testing.T, testing.M dan testing.B
T -> digunakan untuk unit test
M -> alur eksekusi life cycle testing
B -> untuk melakukan benchmarking (mengukur kecepatan)
- Aturan di unit test golang
* Berakhiran _test.go
* Paramater t *testing.T dan tidak mengembalikan return value
* masuk ke folder/package terus run go test buat testing semua (go langsung cari func testing)
* kalau mau spesifik func go test -v -run=NamaFunction
* kalau pengen running test dari folder utama/dari luar folder go test -v ./...
* Menggagalkan test, mneggunkan Fail(), FailNow(), Error(), dan Fatal()
**Fail() menggagalkan unit test tapi tetep menjalankan unit test di package yang lain
**FailNow menggagalkan saat itu juga tanpa melanjutkan unit test yang lain
**Lebih disarankan pakai t.Error atau t.Fatal karena kita dapat menuliskan errornya karena apa
menit 33
---- Assertion -----
Pengecekan secara auto
*nama librarynya testify
----- testing main -----
dalam satu package bisa dijalankan semua test-nya dengan membuat func TestingMain(m *Testing.M)
contoh:
```go
func TestMain(m *testing.M) {
    before
    fmt.Println("before unit tes")
    m.Run()
    fmt.Println("after unit tes")
}
```
dengan begitu kita dapat membuat before dan after test (misalnya sebelum eksekusi connect ke DB)
----- testing sub test -----
kita dapat membuat func test di dalam func test
login dari luar
sama bikin user
