package main

import (
	"fmt"

	abc "gitlab.xxx.yyy/zzz/www_vvv/ABC"
)

func main() {
	fmt.Print("Nothing yet")
	var aTask abc.Task
	aTask.TaskID = 1
	aTask.Keys = make(map[abc.Key]bool)
	fmt.Print(aTask)
}
