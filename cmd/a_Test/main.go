package main

import (
	"fmt"

	abc "gitlab.xxx.yyy./zzz/www_vvv"
)

func main() {
	fmt.Print("Nothing yet")
	var aTask Task
	aTask.TaskID = 1
	aTask.Keys = make(map[Key]bool)
	fmt.Print(aTask)
	abc.ABC()
}
