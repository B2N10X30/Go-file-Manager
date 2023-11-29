package fileManager

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// create Directories
// Rename Directories
// Delete Directories
// organize Files into Directories
func CreateDirectory(dirName string) error {
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		log.Printf("%v", err)
	}
	fmt.Println("Directory  created succesfully!", dirName)
	return nil
}

func RenameDirectory(oldDirName, newDirName string) error {
	err := os.Rename(oldDirName, newDirName)
	if err != nil {
		log.Printf("%v", err)
	}
	fmt.Printf("Directory %s renamed as %s\n", oldDirName, newDirName)
	return nil
}

func RemoveDirectory(dirName string) error {
	ifExist, err := IsExist(dirName)
	if ifExist {
		ifEmpty, _ := IsDirEmpty(dirName)
		if ifEmpty {
			os.Remove(dirName)
			fmt.Printf("Directory %s was removed successfully!", dirName)
		} else {
			fmt.Printf("directory %s is not empty\n", dirName)
		}
		return nil
	} else {
		fmt.Print("Directory does not exist\n")
		return err
	}

}

func IsDirEmpty(dirPath string) (bool, error) {
	// Open the directory
	dir, err := os.Open(dirPath)
	if err != nil {
		return false, err
	}
	defer dir.Close()

	// Read the directory entries
	_, err = dir.Readdirnames(1) //checks for atleast one file or subdirectory
	if err == nil {
		// No error means there is at least one file or subdirectory
		fmt.Println("Directory contains at least one file")
		return false, nil
	}

	// If the error is "no more files(End of file, EOF)", the directory is empty
	if err == io.EOF {
		fmt.Print("Directory is empty\n")
		return true, nil
	}
	return false, err
}

func GetDirSize(dirPath string) {
	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d bytes", dirPath, dirInfo.Size())
}

func IsExist(dirName string) (bool, error) {
	_, err := os.Stat(dirName) // get directory info

	if err == nil {
		//if no directory info -> means directory does not exist
		// fmt.Printf("Directory '%s' exists.\n", dirName)
		return true, nil
	} else if os.IsNotExist(err) {
		// fmt.Printf("Directory '%s' does not exist.\n", dirName)
		return false, nil
	} else {
		fmt.Printf("Error checking Directory: %v\n", err)
		log.Printf("%v", err)
		return false, err
	}
}

func FileOrganizer(dirName string) {
	listDirectory, err := os.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range listDirectory {
		extension := filepath.Ext(file.Name())
		if len(extension) > 1 {
			directoryName := extension[1:]
			//new directory will be dirName/directoryName
			newDirPath := filepath.Join(dirName, directoryName)

			if _, err := os.Stat(newDirPath); os.IsNotExist(err) {
				// Directory does not exist, so create it
				err := os.Mkdir(newDirPath, 0755) // 0755 is the permission mode
				if err != nil {
					log.Fatal(err)
				}
			}
			newFilePath := filepath.Join(newDirPath, file.Name())
			err = os.Rename(filepath.Join(dirName, file.Name()), newFilePath)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("No extension found")
		}
	}
}
