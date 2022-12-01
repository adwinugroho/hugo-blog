---
title: "Golang - Download With File Url"
date: 2022-11-18T16:00:48+07:00
draft: false

categories: ["Golang", "Tutorial"]
tags: ["golang basic",  "golang tutorial", "download file with HTTP GET URL", "unduh file dari HTTP GET URL"]
toc: false
author: "adwin"
---
Cara download file di golang dan mendapatkan file-nya pada lokal komputer kamu! 

Mungkin ini kurang bermanfaat apabila kamu hanya ingin download file. Tinggal pake wget pun beres, disini sebenernya saya ingin menyelesaikan studi kasus dimana file yang besar-besar itu dikirim ke SFTP (SSH File Transfer Protocol) atau upload ke cloud storage macam AWS, google cloud, dll.

Fungsi di bawah merupakan fungsi simpel download file dengan URL menggunakan fungsi HTTP GET, file yang didapat diubah ke []byte untuk dijadikan return value. 
```go
func GetFileFromURL(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        log.Printf("[GetPicFromURL]: error message: %+v\n", err)
        return nil, err
    }
    defer resp.Body.Close()
    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("[GetPicFromURL]: error message: %+v\n", err)
        return nil, err
    }
    return bodyBytes, nil
}
```
Setelah file dalam bentuk bytes diterima, kita perlu menggunakan package github.com/h2non/filetype. Alasannya sederhana setiap kita ingin mengupload file ke AWS misalnya kita perlu mendefinisikan Content-Type, pada package tersebut kita bisa dapat ["MIME"](https://id.wikipedia.org/wiki/Ekstensi_surat_internet_multiguna) dari file dalam bentuk bytes. Selain MIME pada package tersebut kita juga bisa mendapatkan extension file-nya dimana extension penting untuk membuat file name.

```go
func main() {
	// ex file:
    // "https://gorm.io/sponsors-imgs/bytebase.png"
	// "https://bpsdm.pu.go.id/center/pelatihan/uploads/edok/2019/02/bf105_MDL_Sistem_Informasi_dan_Data_SDA.docx"
	// "https://www.stat.ipb.ac.id/agusms/wp-content/uploads/2022/02/Mater-4-Visualisasi-Data-1.pdf"
	fileBytes, err := GetFileFromURL("https://i3.ytimg.com/vi/CgjuA0EVMJo/maxresdefault.jpg")
	if err != nil {
		panic(err)
	}
	typeFile, _ := filetype.Get(fileBytes)
	// write to file system
    filename := fmt.Sprintf("./download/download.%s", typeFile.Extension)
    // 0644 -> permission
	err = os.WriteFile(filename, fileBytes, 0644)
	if err != nil {
		panic(err)
	}
}
```