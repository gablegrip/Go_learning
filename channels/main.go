package main

import (
	"log"

	"github.com/gablegrip/myniceprogram/helpers"
)

const numPool = 10

func CalulateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalulateValue(intChan)

	num := <-intChan
	log.Println(num)
}
