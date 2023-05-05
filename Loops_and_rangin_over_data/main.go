package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"Jared", "Derby", "jared.derby@hotmail.com", 48})
	users = append(users, User{"Mary", "Jone", "mary@jones.com", 22})
	users = append(users, User{"Sally", "Brown", "sally@smith", 45})
	users = append(users, User{"Alex", "Anderson", "alex@smith", 17})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)

	}

}
