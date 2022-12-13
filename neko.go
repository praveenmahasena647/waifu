package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	var data []byte = getNeko()
	var url string = createUrl(data)
	createFile(url)
}

func getNeko() []byte {
	var req, reqErr = http.Get("https://nekos.best/api/v2/neko")

	if reqErr != nil {
		log.Println(reqErr.Error())
		os.Exit(2)
	}
	defer req.Body.Close()
	var data, dataErr = ioutil.ReadAll(req.Body)
	if dataErr != nil {
		log.Println(dataErr.Error())
		os.Exit(2)
	}

	return data
}

func createUrl(data []byte) string {
	var url string
	var leng int = len(data) - 5

	for leng >= 0 {
		if data[leng] == '"' {
			break
		} else {
			url = string(data[leng]) + url
		}
		leng--
	}
	return url
}

func createFile(url string) {

	var req, reqErr = http.Get(url)

	if reqErr != nil {
		log.Println(reqErr.Error())
		os.Exit(2)
	}
	defer req.Body.Close()

	var data, dataErr = ioutil.ReadAll(req.Body)

	if dataErr != nil {
		log.Println(dataErr.Error())
		os.Exit(2)
	}

	var file, fileErr = os.Create("neko.png")

	if fileErr != nil {
		log.Println(fileErr.Error())
		os.Exit(2)
	}
	var written, err = file.Write(data)
	if err != nil {
		log.Println(err.Error())
		os.Exit(2)
	}
	log.Println(written)
}
