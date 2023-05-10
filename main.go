package main

import (
	"log"
	"sort-files/sort"
)

func main() {
	files, err := sort.GetFilesInDir()
	if err != nil {
		log.Fatalln(err)
	}
	num, err := sort.MoveFilesToDirs(files)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Script moved %d files to directories.", num)
}
