package main

import (
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	path := "/home/blanka/Downloads/"

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		uuid := uuid.NewV4().String()

		err := os.Rename(path+file.Name(), path+uuid+".mp3")

		if err != nil {
			log.Fatal(err)
		}
	}
}
