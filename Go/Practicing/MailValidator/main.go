package main

import (
	"fmt"
	"regexp"
)

func main() {
	emails := []string{"test@example.com", "invalid-email", "user@domain.com", "another@domain", "@missingusername.com"}

	for _, email := range emails {
		if isValidEmail(email) {
			fmt.Printf("%s is valid\n", email)
		} else {
			fmt.Printf("%s is invalid\n", email)
		}
	}
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
