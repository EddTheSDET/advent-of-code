package utils

import (
	"bufio"
	"log"
	"os"
)

func OpenDataFile(filePath string) *bufio.Scanner {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path + "/" + filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	return scanner
}
