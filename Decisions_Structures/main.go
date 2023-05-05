package main

import "log"

func main() {

	myVar := "horse"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")

	case "dog":
		log.Println("cat is set to dog")

	case "fish":
		log.Println("fish is set to fish")

	default:
		log.Println("cat is something else")

	}

}
