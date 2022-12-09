package main

import (
	"log"
	"math/rand"
)

func main() {
	var getPicsLink string = getPics()
	log.Println(getPicsLink)
}

func getPics() string {
	var catagory string = getRandom()
	log.Println(catagory)
	return ""
}

func getRandom() string {
	var collection []string = []string{"waifu", "neko", "shinobu", "megumin", "bully", "cuddle", "hug", "awoo", "kiss", "pat", "smug", "bonk", "blush", "smile", "handhold", "glomp", "happy", "wink", "dance"}
	//	leng 19
	var index int = rand.Intn(18-1) + 1
	log.Println(index)
	return collection[index]
}
