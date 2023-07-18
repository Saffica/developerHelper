package main

import (
	"fmt"
	"log"
	"main/internal/client"
)

func main() {
	vcs, err := client.GetVCS()
	if err != nil {
		log.Panic(err.Error())
	}

	vcsForRequest, err := client.PrepareVCS(vcs)
	if err != nil {
		log.Panic(err.Error())
	}

	files, err := client.GetRecordsByVCS(vcsForRequest)
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(files)
	fileIsExist, err := client.Exists("D:\\Developer")
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(fileIsExist)
}
