package fileManager

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func FileOrganizer() {
	listDirectory, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range listDirectory {
		extension := filepath.Ext(file.Name())
		if len(extension) > 1 {
			dirName := extension[1:]

			if _, err := os.Stat(dirName); os.IsNotExist(err) {
				// Directory does not exist, so create it
				err := os.Mkdir(dirName, 0755) // 0755 is the permission mode
				if err != nil {
					log.Fatal(err)
				}
			}
			newPath := filepath.Join(dirName, file.Name())
			err = os.Rename(file.Name(), newPath)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("No extension found")
		}
	}
}
