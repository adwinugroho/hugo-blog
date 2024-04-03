---
title: "What I Learned - Sekilas Tentang Banking System"
date: 2023-06-06T00:04:20+08:00
draft: false

categories: ["TIL", "WIL", "Research"]
tags: ["bank",  "banking system"]
toc: false
author: "adwin"
---
Akhirnya jumpa lagi di edisi artikel **What I Learned**.

Berhubung di *company* tempat saya bekerja sedang migrasi dan pindah fokus industri (awalnya kesehatan), sekarang jadi lebih *general* fokusnya dikarenakan bertambah 1 cabang industri yang akan dijelajahi yaitu, fintech dan *banking*.

Tidak perlu dijelaskan panjang lebar, tetapi pada intinya para dev serta tim-tim yang berurusan dengan tech sekarang dihadapkan dengan istilah-istilah baru, yang biasanya mendengar kata lab, health record, BMI, ICD-10 dan lain sebagainya itu sekarang ada hal yang asing di telinga.

Pada artikel ini mungkin lebih banyak pembahasan secara umum saja seperti apa tujuannya sistem di bank, apa saja aktivitas yang ada pada sistem bank serta istilah-istilah umum yang sering digunakan. Mungkin kedepannya apabila telah banyak istilah yang lebih *technical* akan saya buatkan lagi satu artikel khusus tentang itu atau mungkin saja digabung pada artikel ini.

Aktivitas inti dari sebuah *banking system* adalah:
- Mengumpulkan dana pihak ketiga (DPK) 
- Menyalurkan kembali ke pihak lain dengan imbalan
- Mengelola uang yang dititipkan nasabah untuk menghasilkan keuntungan (dari biaya dan bunga/margin) lalu membagikan keuntungan kepada pihak yang menitipkan dana tersebut
- Memberikan “rasa aman” dan “kepercayaan” sebagai jaminan kepada pihak ketiga

Istilah umum yang ada pada sistem perbankan yaitu:
- **CIF (Customer Information File)** berisi tentang data pribadi nasabah
- **KYC (Know Your Customer)** merupakan proses verifikasi identitas dan memonitor transaksi nasabah 
- **Alternate ID/Number** merupakan ID Customer/Nomor Rekening yang berada pada sistem lain.
- **Cek/BG (Bilyet Giro)** fasilitas yang diberikan bank kepada nasabah tertentu untuk memerintah bank melakukan pemindahan buku
- **Saving** rekening simpanan yang diberikan kepda nasabah perseorangan
- **QQ** merupakan rekening atas nama (misalnya rekening anak tapi anak belum cukup syarat untuk pembukaan rekening)
- **RDN (Rekening Dana Nasabah)** biasanya ada pada perusahaan efek atau diwajibkan OJK (Otoritas Jasa Keuangan), nasabah hanya bisa meletakkan dana pada rekening itu dan tidak bisa mengeluarkan dana, kalo kalian tau dan untuk apa aplikasi **bibit** tentunya tidak asing dengan kata ini
- **Pencairan** rekening yang biasanya dibuka untuk mewakili pinjaman yang diambil nasabah, misalnya ingin kredit rumah
- **Escrow/Rekber (Rekening Bersama)** fasilitas yang diberikan pihak ketiga untuk proses bisnis 2 arah. Misal kamu *checkout* di aplikasi e-commerce kesayangan kamu, penjual tidak otomatis menerima uang yang kamu transfer
- **Virtual Account** rekening yang tidak ada di bank, tidak rill yang biasanya terdiri dari nomor BIN (Business Identification Number) dan angka penanda pemilik. Uang yang ke rekening ini biasanya diletakkan pada rekening escrow/rekber
- **Internal Account** rekening internal bank yang dibuat dengan tujuan tertentu
- **Saldo Minimum** nominal saldo minimal yang harus ada pada rekening
- **Locking Amount** mekanisme untuk mengunci saldo nasabah dengan nominal tertentu
- **Dorman** kondisi dimana rekening tidak ada aktivitas dalam kurun waktu tertentu
- **Overdraft** sebuah kondisi dimana saldo kita tidak mencukupi untuk melakukan aktivitas (saldo minus) 
- **Category PnL** bukan berbentuk rekening tapi bisa ditransisikan layaknya rekening, tujuannya untuk mencatat keuntungan/kerugian (neraca) bank 
- **Collateral/Jaminan** merupakan aset *tangible* (dapat dilihat fisik dan punya nilai tertentu) maupun *intangible* (seperti paten, copyright, dan yang terbaru kanal youtube) yang dijadikan syarat untuk melakukan pinjaman
- **Limit** dapat berupa batas transaksi harian, nominal dalam 1x pindah buku, banyaknya melakukan pindah buku dalam 1 hari
- **Kolek (Kolektibilitas Kredit) Nasabah** biasanya dikategorikan berdasarkan
    1. Lancar
    2. DPK (Dalam Perhatian Khusus), tunggakkan 1-90 hari
    3. Tidak Lancar, tunggakkan 91-120 hari
    4. Diragukan, tunggakkan 121-180 hari
    5. Macet, tunggakkan lebih dari 180 hari
- **Bunga/Margin** untuk bank konvensional biasa ditampakan dalam bentuk % (*fixed* / *floating* mengikuti suku bunga Bank Indonesia). *Margin* lumrah dalam bank syariah biasa berupa nominal yang harus dibayarkan dan sudah jelas saat akad totalan *margin* sampai akhir masa pinjaman.
- **Bagi Hasil/Profit Sharing** merupakan jenis akad pinjaman pada bank syariah dimana *margin* dijadikan sebagai bagi hasil
