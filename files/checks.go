package files

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func CheckFiles() {
	_, err := os.Stat("shedrack.txt")

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file does not exist")
	} else {
		fmt.Println("Files exists")
	}

	newFile, err := os.Create("shedrack.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer newFile.Close()
	fmt.Println("shedrack file is created")

	deleteFiles()
	getFilesInfo("shedrack.txt")

}

func deleteFiles() {
	err := os.Remove("recipes.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File deleted")
}

func getFilesInfo(fileName string) string {
	fileInfos, err := os.Stat(fileName)

	if err != nil {
		log.Fatal(err)
	}

	myTime := fileInfos.ModTime()
	fmt.Printf("The last modified time is -> %v\n", myTime)

	fileSize := fileInfos.Size()
	fmt.Printf("File size is -> %v\n", fileSize)

	return fileName
}
