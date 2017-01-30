package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	v1Tag, err := ReadV1Tag(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v1Tag)
}

type V1Tag struct {
	title   string
	artist  string
	album   string
	year    int
	comment string
	track   int
	genre   int
}

func ReadV1Tag(file *os.File) (V1Tag, error) {
	// http://id3.org/ID3v1
	v1Tag := V1Tag{}
	file.Seek(-128, 2)
	tagBytes := make([]byte, 3)
	file.Read(tagBytes)
	// First 3 bytes must = ASCII "TAG" from ID3V1 spec.
	if tagBytes[0] != 84 || tagBytes[1] != 65 || tagBytes[2] != 71 {
		errString := "File: \"" + file.Name() + "\" has no ID3 V1 tag."
		return v1Tag, errors.New(errString)
	}
	// Parse elements in order.
	v1Tag.title = getNextV1Field(file, 30)
	v1Tag.artist = getNextV1Field(file, 30)
	v1Tag.album = getNextV1Field(file, 30)
	v1Tag.year, _ = strconv.Atoi(getNextV1Field(file, 4))
	v1Tag.comment = getNextV1Field(file, 28)
	file.Seek(1, 0)
	v1Tag.track, _ = strconv.Atoi(getNextV1Field(file, 1))
	v1Tag.genre, _ = strconv.Atoi(getNextV1Field(file, 28))
	return v1Tag, nil
}

func getNextV1Field(file *os.File, size int) string {
	outputBytes := make([]byte, size)
	file.Read(outputBytes)
	output := string(outputBytes)
	return output
}
