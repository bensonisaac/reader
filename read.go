package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var userFile string
	fmt.Println("Enter the name of the file here: ")
	fmt.Scanln(&userFile)

	file, err := os.Open(userFile)
	if err != nil {
		fmt.Println("Error occured opening file", err)
	}

	defer file.Close()

	names := make([]Name, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		field := strings.Fields(line)

		if len(field) >= 2 {
			fname := field[0]
			lname := field[1]

			n := Name {
				fname: fname,
				lname: lname,
			}

			names = append(names, n)
		}
	}
//
//	if err := scanner.Err(); err != nil {
//		fmt.Println("Error reading the file", err)
//		return
//	}

	for _, name := range names {
		fmt.Printf("First name: %s, Last name: %s\n", name.fname, name.lname)
	}
}