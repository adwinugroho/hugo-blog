---
title: "Docker - Basic Command"
date: 2024-06-12T14:27:48+07:00
draft: true

categories: ["Docker", "DevOps"]
tags: ["docker basic",  "docker tutorial", "notes from docker", "catatan docker"]
toc: false
author: "Adwin"
---
Catatan tentang docker:

- Untuk melihat *command-command* apa saja yang tersedia:
```bash
docker help build
```

- Untuk menamai *tag* atau dengan *shorthand* -t
```bash
docker build . -t my-app-name
```

- Untuk memberi tanda pada *tag* menggunakan *command* `:` sesudah *tag name* karena kalau tidak secara *default* docker *build* akan menandai *image* kita dengan *tag latest*
```bash
docker build . -t my-app-name:v1
```
