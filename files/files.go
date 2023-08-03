package files

import (
	"fmt"
	"log"
	"os"
)

func GoFiles() {

	myGoFiles, err := os.Create("recipes.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer myGoFiles.Close()

	fmt.Println("Files created successfully")

	iterateUsers([]string{"john", "mark", "joe"}, sayGreeting)
}

func sayGreeting(greet string) {
	fmt.Printf("Good morming mr -> %v\n", greet)
}

func iterateUsers(users []string, myUser func(string)) {
	for _, user := range users {
		myUser(user)
	}
}
