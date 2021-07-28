---
title: "Golang - Generate Random String"
date: 2021-07-03T18:51:34+08:00
draft: false

categories: ["Golang", "Tutorial"]
tags: ["golang basic",  "golang tutorial", "generate random string", "membuat huruf acak"]
toc: false
author: "Adwin"
---
Di artikel sebelumnya kita sudah membahas bagaimana cara membuat [angka acak di golang](https://blog.cryppy.xyz/article/generate-random-number/), pada dasarnya sama, teorinya pun sama, kita hanya perlu mengganti _charset_ dan yang paling  penting adalah *parameter* dan *return value*-nya

```go
import (
	"math/rand"
	"time"
)

func GenerateRandomString(lengthChar int) string {
	var charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkklmnopqrstuvwxyz"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, lengthChar)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
```
