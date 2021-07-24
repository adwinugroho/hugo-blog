package main

import (
	"log"

	"github.com/adwinugroho/algolia-with-go/config"
)

func main() {
	log.SetFlags(log.Llongfile | log.Ltime)
	// readFile, err := md.ReadAllFilesInDirectory(config.PATH_ARTICLE)
	// if err != nil {
	// 	log.Println("error while read file")
	// 	return
	// }

	// fmt.Println(readFile)
	config.NewMongoDB()
}
