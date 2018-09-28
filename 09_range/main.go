package main

import "fmt"

func main() {
	// ids := []int{33, 76, 45, 34, 12, 54, 23, 26}

	// Loop through ids
	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }

	// Not using index
	// for _, id := range ids {
	// 	fmt.Printf("ID: %d\n", id)
	// }

	// Add ids together
	// summ := 0
	// for _, id := range ids {
	// 	summ += id
	// }
	// fmt.Println("Sum", summ)

	// Range with map
	emails := map[string]string{"Bob": "bob@mail.com", "Art": "art@mail.com"}

	// for k, v := range emails {
	// 	fmt.Printf("%s: %s\n", k, v)
	// }

	for k := range emails {
		fmt.Println("Name: ", k)
	}

}
