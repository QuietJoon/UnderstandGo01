package main

import (
	"fmt"
)

func main() {
	fmt.Print("Nothing yet")
	var aTask Task
	aTask.TaskID = 1
	aTask.Keys = make(map[Key]bool)
	fmt.Print(aTask)
	ABC()
}
