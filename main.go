package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error = getImage()
	if err != nil {
		log.Println("error during the Process")
	}
}

func getImage() error {
	var req, reqErr = http.Get("https://pic.re/image")

	if reqErr != nil {
		return errors.New("fUCK iMG rEQ eRROR")
	}
	defer req.Body.Close()

	var data, dataErr = ioutil.ReadAll(req.Body)

	if dataErr != nil {
		return errors.New("Fuck img decode Error")
	}

	var file, fileErr = os.Create("waifu.png")

	if fileErr != nil {
		return errors.New("Fuck img decode Error")
	}
	var done, FileWriteErr = file.Write(data)
	if FileWriteErr != nil {
		return errors.New("img Write Error")
	}

	log.Println(done)
	return nil
}
