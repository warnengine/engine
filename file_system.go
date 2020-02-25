package main

import (
	"io/ioutil"
	"os"
)

// InitFileSystem attach zip passed as arguments to physfs file system
func InitFileSystem() {
	// Ooops don't want to work with physfs anymore
	print("no physfs dear")
	// // Init physfs
	// physfs.Init()
	// // Add each zip passed as arguments
	// for index := 1; index <= len(os.Args)-1; index++ {
	// 	log.Print("adding")
	// 	log.Print(os.Args[index])

	// 	if _, err := os.Stat(os.Args[index]); os.IsNotExist(err) {
	// 		panic("Unable to read file at " + os.Args[index])
	// 	}
	// 	physfs.AddToSearchPath(os.Args[index], false)
	// }

}

// ReadFile gets the content of the file according the given path
func ReadFile(path string) []byte {
	// file, err := physfs.Open(path)
	// if err != nil {
	// 	panic(err)
	// }
	// length, err := file.Length()
	// if err != nil {
	// 	panic(err)
	// }
	// buffer := make([]byte, length)
	// file.Read(buffer)
	content, err := ioutil.ReadFile(os.Args[1] + path)
	if err != nil {
		panic(err)
	}

	return content
}
