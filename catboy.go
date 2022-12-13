package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	var data []byte = getData()
	var url string = getUrl(data)
	CreateFile(url)
}

func CreateFile(url string) {
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

func getData() []byte {
	var req, reqErr = http.Get("https://api.catboys.com/img")
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

func getUrl(url []byte) string {
	var str []byte
	var leng int = len(url)

	for i := 0; i < leng; i++ {
		if url[i] == '\n' {
			continue
		}
		if url[i] == ',' {
			break
		}
		str = append(str, url[i])
	}
	leng = len(str)
	return string(str[8 : leng-1])
}
