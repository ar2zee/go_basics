package main

import (
	"fmt"
)

func main() {
	// Define map
	// emails := make(map[string]string)

	// // Assgn kv
	// emails["Bob"] = "bob@mail.com"
	// emails["Art"] = "art@mail.com"
	// emails["Mike"] = "mike@mail.com"

	// Declarw map and kv
	emails := map[string]string{"Bob": "bob@mail.com", "Art": "art@mail.com"}
	emails["Mike"] = "mike@mail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	// delete(emails, "Bob")
	// fmt.Println(emails)
}
