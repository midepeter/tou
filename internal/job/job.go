package job

import (
	"errors"
	"fmt"
)

type Job struct {
	id    string
	value interface{}
}

func (j Job) Validate() error {
	if j.id == " " {
		return errors.New(fmt.Sprintln("The message cannot be empty"))
	}
	return nil
}
