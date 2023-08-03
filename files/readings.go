package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFiles() {

	readNewFiles()

	myFiles, err := os.Create("workspace.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer myFiles.Close()

	fmt.Println("File created")
}

func readNewFiles() {
	f, err := os.Open("workspace.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	myScanner := bufio.NewScanner(f)

	for myScanner.Scan() {
		fmt.Println(myScanner.Text())
	}

	if err := myScanner.Err(); err != nil {
		log.Fatal(err)
	}
}
