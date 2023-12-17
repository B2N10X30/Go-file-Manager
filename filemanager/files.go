package filemanager

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	FileName     string
	Extension    string
	IsExecutable bool
}

func CreateFile(fileName, fileExtension string, isExecutable bool) error {
	var filePermission uint
	filePath := fileName + "." + fileExtension

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("%v", err)
	}
	defer file.Close()

	if isExecutable {
		filePermission = 0755
	} else {
		filePermission = 0744
	}

	err = os.Chmod(filePath, fs.FileMode(filePermission))
	if err != nil {
		return err
	} else {
		fmt.Println("File set as Executable. \n", filePath)
	}
	return nil
}

func RenameFile(oldFileName, newFileName string) error {
	err := os.Rename(oldFileName, newFileName)
	if err != nil {
		log.Printf("%v", err)
	}
	fmt.Printf("File %s renamed as %s\n", oldFileName, newFileName)
	return nil
}

func IsFileExist(fileName string) (bool, error) {
	_, err := os.Stat(fileName) // get file info

	if err == nil {
		//if no file info -> means file does not exist
		// fmt.Printf("Directory '%s' exists.\n", dirName)
		return true, nil
	} else if os.IsNotExist(err) {
		// fmt.Printf("Directory '%s' does not exist.\n", dirName)
		return false, nil
	} else {
		fmt.Printf("Error checking file: %v\n", err)
		log.Printf("%v", err)
		return false, err
	}
}

func GetFileSize(filepath string) {
	fileinfo, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d bytes and was last modfied at %s", filepath, fileinfo.Size(), fileinfo.ModTime().Format(time.RFC850))
}

func RemoveFile(fileName string) error {
	ifExist, err := IsExist(fileName)
	if err != nil {
		return err
	}
	if ifExist {
		ifEmpty, err := IsDirEmpty(fileName)
		if err != nil {
			return err
		}
		if ifEmpty {
			err := os.Remove(fileName)
			if err != nil {
				return err
			}
			fmt.Printf("File %s was removed successfully!", fileName)
		} else {
			fmt.Printf("file %s is not empty\n", fileName)
		}
		return nil
	} else {
		fmt.Print("File does not exist\n")
		return err
	}

}

func Search(fileName, searchDir string) (string, error) {
	var foundFilePath string

	//use recursive function to traverse the specified directory
	err := filepath.Walk(searchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//check and skip directories
		if info.IsDir() {
			return nil
		}
		//set filepath to filename
		if info.Name() == fileName {
			foundFilePath = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		log.Printf("Error walking directory '%s': %v", searchDir, err)
		return "", err
	}
	if foundFilePath == "" {
		return "", fmt.Errorf("File '%s' not found in directory '%s'", fileName, searchDir)
	}
	fmt.Printf("%s found at %s", fileName, foundFilePath)
	return foundFilePath, nil
}

func Reader(Path string) {
	file, err := os.Open(Path)
	if err != nil {
		log.Println("Error Opening file ", err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println(err)
	}
	size := fileInfo.Size()
	myReader := make([]byte, size)
	_, err = file.Read(myReader)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(myReader))
}
