package tou

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/urfave/cli/v2"
)

var (
	data string
	port string
)

var AddCmd = &cli.Command{
	Name:  "add",
	Usage: "The add command is used for adding to the running workqueue",
	Action: func(c *cli.Context) error {
		return Add(c, data) 
	},
	Flags: flags,
}

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "data",
		Usage:       "to specify port on which work queue server instance is to run",
		Aliases:     []string{"d"},
		Required:    true,
		Destination: &data,
	},
	&cli.StringFlag{
		Name:        "port",
		Usage:       "to specify port on which work queue server instance is to run",
		Aliases:     []string{"p"},
		Required:    true,
		Destination: &port,
	},
}

func Add(ctx *cli.Context, job string) error {
	if job == "" {
		return errors.New("Job is empty: Operation invalid")
	}

	body, err := json.Marshal(job)
	if err != nil {
		return errors.New("Unable to marshal result")
	}

	postBytes := bytes.NewBuffer(body)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "http://127.0.0.1:"+port+"/add", postBytes)
	c := &http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		return errors.New("Unable to make new request")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Unable to read response body")
	}

	log.Println("The response", string(respBody))
	
	return nil
}
