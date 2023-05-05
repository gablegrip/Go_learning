package main

import (
	"log"
)

const numPool = 1000

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
