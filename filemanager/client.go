package filemanager

/* authour: sameul
 */

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func Client() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var filepath string
	fmt.Print("Please provide the file path (e.g., path/to/file): ")

	// Read the file path from the user input
	reader := bufio.NewReader(os.Stdin)
	filepath, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim any leading or trailing spaces or newline characters
	filepath = strings.TrimSpace(filepath)

	// Extract the filename from the path
	filename := filepath[strings.LastIndex(filepath, "/")+1:]

	// Send the filename first
	conn.Write([]byte(filename))

	// Open the file for reading
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Send the file data
	_, err = io.Copy(conn, file)
	if err != nil {
		panic(err)
	}
}
