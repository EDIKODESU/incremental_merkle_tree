package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/EDIKODESU/incremental_merkle_tree/tree_methods"
)

func main() {
	tree := tree_methods.IncrementalMerkleTree{}

	// var choice int
	var data string
	var id int

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add")
		fmt.Println("2. Delete")
		fmt.Println("3. Update")
		fmt.Println("4. Check")
		fmt.Println("5. Show Tree")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		if !scanner.Scan() {
			fmt.Println("Failed to read input")
			return
		}

		input := scanner.Text()

		if input == "" {
			fmt.Println("Invalid choice!")
			continue
		}

		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 6 {
			fmt.Println("Invalid choice!")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter data to add: ")
			if !scanner.Scan() {
				fmt.Println("Failed to read input")
				return
			}
			data = scanner.Text()

			fmt.Print("Enter new ID for block: ")
			if !scanner.Scan() {
				fmt.Println("Failed to read input")
				return
			}
			id, err = strconv.Atoi(scanner.Text())
			if err != nil || id < 0 {
				fmt.Println("Invalid block ID. Please enter a positive integer.")
				continue
			}
			err = tree.AddData(id, data)
			if err == nil {
				fmt.Println("Data added!")
			}
		case 2:
			fmt.Print("Enter block ID to delete:")
			if !scanner.Scan() {
				fmt.Println("Failed to read input")
				return
			}
			id, err = strconv.Atoi(scanner.Text())
			if err != nil || id < 0 {
				fmt.Println("Invalid block ID. Please enter a positive integer.")
				continue
			}
			err = tree.DeleteData(id)
			if err == nil {
				fmt.Println("Data deleted successfully!")
			}
		case 3:
			fmt.Print("Enter block ID to update: ")
			if !scanner.Scan() {
				fmt.Println("Failed to read input")
				return
			}
			id, err = strconv.Atoi(scanner.Text())
			if err != nil || id < 0 {
				fmt.Println("Invalid block ID. Please enter a positive integer.")
				continue
			}
			fmt.Print("Enter new data: ")
			if !scanner.Scan() {
				fmt.Println("Failed to read input")
				return
			}
			data = scanner.Text()
			err = tree.UpdateData(id, data)
			if err == nil {
				fmt.Println("Data updated!")
			}
		case 4:
			fmt.Print("Enter data to check: ")
			if !scanner.Scan() {
				fmt.Println("Failed to read input")
				return
			}
			data = scanner.Text()
			tree.CheckData(data)
		case 5:
			tree.Display()
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice!")
			continue
		}
	}
}
