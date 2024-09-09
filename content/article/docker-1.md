---
title: "Docker - Basic Command"
date: 2024-06-12T14:27:48+07:00
draft: false

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

- Perintah `docker run` itu sebenarnya menjalankan dua perintah yaitu membuat *container* baru (`docker container create`) lalu menjalankan *container* tersebut (`docker container start`)

- Untuk memberi *environment variable*:
```bash
docker container create -e PORT=8083 -p 8083:8083
```

- Untuk running secara *detached* atau di dalam *background* yaitu dengan perintah `-d`
```bash
docker run -it -d -e PORT=8073 -p 8083:8073 my-app-image
```

- *container* yang sedang jalan tidak dapat dihapus, perintah hapus
```bash
docker container rm container_id
```

- Untuk mencari *ip address* yang ada di dalam docker network
```bash
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' container_name
```