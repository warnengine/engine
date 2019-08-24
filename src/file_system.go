package main

import (
	"log"
	"os"

	"github.com/DeedleFake/Go-PhysicsFS/physfs"
)

func InitFileSystem() {
	// Init physfs
	physfs.Init()
	for index := 1; index <= len(os.Args)-1; index++ {
		log.Print("adding")
		log.Print(os.Args[index])

		if _, err := os.Stat(os.Args[index]); os.IsNotExist(err) {
			panic("Unable to read file at " + os.Args[index])
		}
		physfs.AddToSearchPath(os.Args[index], false)
	}

}

func ReadFile(path string) []byte {
	file, err := physfs.Open(path)
	if err != nil {
		panic(err)
	}
	length, err := file.Length()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, length)
	file.Read(buffer)

	return buffer
}
