package job

import (
	"errors"
	"fmt"
	"time"
)

type Job struct {
	Id      string    `json:"id"`
	Payload string    `json:"payload"`
	TTL     time.Time //TTL determines the duration a particular job can live for before it expires
}

func (j Job) Validate() error {
	if j.Id == " " {
		return errors.New(fmt.Sprintln("The message cannot be empty"))
	}
	return nil
}
