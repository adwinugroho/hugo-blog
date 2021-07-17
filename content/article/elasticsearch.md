---
title: "About Elasticsearch"
date: 2021-03-17T11:27:18+08:00
draft: true

categories: []
tags: []
toc: false
author: ""
---
## Sejarah elasticsearch
dari library java dll

## Kapan butuh elasticsearch
- saat program membutuhkan pencarian yang kompleks
- saat data yang dicari sudah banyak
- saat kita butuh aggregate data

## Bagian 1
- Hanya digunakan untuk pendamping DB utama
- Elasticsearch node -> yang ada di aplikasi
- Node harus ganjil kalau di prod
- masterless, bisa akses dari node manapun (insert delete dll) asal tetep satu cluster

## Elasticsearch index
- Index
   - Ketika buat index, jadinya lucene document (library java)
   - Tambahin prefix tiap aplikasi, karena di elasticsearch gaada istilah DB
- Type (ini sesat)

## Index Sharding
- Memotong index menjadi beberapa bagian
- Secara default index dipotong menjadi 5
- Agar mudah di distribusikan di Node

## Replication
- Menduplikasi data (index)
- Secara default elasticsearch membuat 1 replication/duplikasi index
- Ga pernah dibuat di satu node yang sama
- Index master dan index replikasi

## Document Routing
- Menentukan dimana lokasi document harus disimpan
- Tekniknya hash(id) terus dimodulo jumlah dari shard-nya

## Distributed Document
- Creating
- Indexing
- Deleting

## Immutable document
- Tidak akan berubah selamanya / dihapus
- Tidak perlu locking data
- Tiap delete terbuat file baru dengan tanda .del
- Tiap update delete (menandai file dengan tanda .del) baru buat dokumen bary
- Ada scheduler buat clean up buat data-datanya

## Distributed Search
- Tahapan dibagi 2, Query dan Fetch
- Hasil searching hanya berupa ID Document + Score dari hasil searchin
- Setelah itu baru di fetch, dibagi-bagi sesuai query dan apa yang dibutuhkan

## Problem search elastic
- Terjadi deep pagination
- Biaya besar untuk memory, bandwith, dan CPU
- Disarankan tidak melakukan paging terlalu dalam