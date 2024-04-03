---
title: "What I Learned - Redis Keyspace Notification"
date: 2022-10-22T22:58:14+07:00
draft: false

categories: ["TIL", "WIL", "Tutorial", "Research"]
tags: ["golang basic",  "golang tutorial", "redis keyspace notification golang"]
toc: false
author: "adwin"
---
Kenapa bukan Today I Learned? Ya biar ga mainstream aja haha.

*First of all* apa itu *redis keyspace notification*? sederhanya, ini adalah sebuah *publish-subscribe* pada redis dimana ketika terjadi sebuah perubahan data pada redis, maka akan terdapat notif. Kalau sudah sering *develop* arsitektur yang menggunakan *microservice* tentunya tidak asing dengan istilah *publish (pub)* dan *subscribe (sub)*.

Apabila kita menggunakan *message broker* nats.io (sekarang udah jadi jetstream) tentunya paham apa arti dari pub dan sub itu tadi, atau misalnya kita pakai kafka yang mungkin sedikit berbeda penyebutannya, yaitu, *produce, consume* dan lain sebagainya. Pada intinya semua itu sama, untuk mengirim dan menerima pesan (data). Mau pake yang mana tergantung kebutuhan masing-masing.

redis sendiri, kebanyakan dipakai untuk *caching* ataupun *rate limiter* akan tetapi disini saya akan membahas implementasi *redis keyspace notification* menggunakan redis. Kenapa saya pake ini? Kasusnya adalah, ketika saya men-*develop* sebuah *service* yang tugasnya adalah untuk membagikan kumpulan data dengan *output web page* dimana *link* beserta seluruh isi *web page* tersebut akan hangus pada waktu yang telah ditentukan. *In example* ketika pasien membagikan catatan kesehatannya pada dokter.

*By default* fitur ini di redis itu ga aktif, jadi kudu di aktifin dulu nih. Oh iya rdb disitu adalah redis *client* disini saya pakai *package github.com/go-redis/redis/v9*
```go
var statusCmd = rdb.ConfigSet(ctx, "notify-keyspace-events", "KEA")
```
parameter ketiga merupakan kode untuk mengaktifkan keyspace notification karena *by default* kan *event* tersebut tidak aktif. 
```go
// example set expire
var status = rdb.SetEx(ctx, "test", "value", 1*time.Minute)
log.Println("redis command status: ", status)
```
Setelah itu kita coba untuk menginisialisasi data dengan *key* dan *value* seperti di bawah dan akan *expire* dalam waktu 1 menit. Terakhir, untuk mengeksekusi pub/sub redis keyspace notification.
```go
    pubSub := rdb.PSubscribe(ctx, "__keyevent@0__:expired")
	log.Printf("pubSub:%+v\n", pubSub)
	var isDone bool
	for {
		msgi, err := pubSub.Receive(ctx)
		if err != nil {
			log.Println("error cause: ", err)
			return
		}
		switch msg := msgi.(type) {
		case *redis.Message:
			log.Printf("Message: %s %s\n", msg.Channel, msg.Payload) // msg.Payload == id in redis
			isDone = true
		case *redis.Subscription:
			log.Printf("Subscription: %s %s %d\n", msg.Kind, msg.Channel, msg.Count)
			if msg.Count == 0 {
				return
			}
			continue
		case error:
			log.Printf("error cause: %v\n", msg)
			continue
		}
		if isDone {
			log.Println("Exit process")
			break
		}
	}
```
pada line 1 inisialiasi subscribe event, **@0** merupakan index pada DB dimana data disimpan. Jalankan *infity loop* untuk menerima pesan (data), disini sengaja diberi *variable* **isDone** agar *looping* berhenti ketika data selesai di proses.

**Take a note**
Pub/Sub di redis ini ga punya fitur persistence data, alias kalau ga ada yang nerima pesan (data) maka data tersebut akan terhapus, in example misalnya koneksi terputus.