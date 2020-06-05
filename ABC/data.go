package abc

// TaskID :
type TaskID = int

// Tasks : Set of Task
type Tasks = map[TaskID]Task

// Key :
type Key = int

// Keys : Set of Key
type Keys = map[Key]bool

// Task : Task
type Task struct {
	TaskID TaskID // Task ID
	Keys   Keys
}

// Tree :
type Tree struct {
	Tasks Tasks
	Keys  Keys
}
