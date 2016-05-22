package task

import "time"

type status int

// status enum.
const (
	UNKNOWN status = iota
	TODO
	DONE
)

// Deadline is a struct to hold the deadline time.
type Deadline struct {
	time.Time
}

// NewDeadline returns a newly created Deadline with time t.
func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

// Task is task struct
type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}
