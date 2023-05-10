package main

import (
	"fmt"
	"sort-files/sort"
)

func main() {
	files, _ := sort.GetFilesInDir()
	num, _ := sort.MoveFilesToDirs(files)
	fmt.Printf("Script moved %d files to directories.", num)
}
