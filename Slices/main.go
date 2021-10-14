package main

import "fmt"

// ################### Example of Slice ###################
// Slice create slice data structure and array data structure.
// InMemory Slice data scrturcture is created in new address (length, cap, ptr to head), but still ref to the original array structure ([]string{"Hi ", "there", "programmer!"})

func main() {
	mySlice := []stirng{"Hi", "there", "master"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(slice []string {
	slice[0] = "Bye"
})
