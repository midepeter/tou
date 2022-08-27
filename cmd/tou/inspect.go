package tou

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	id          string
	inspectPort string
)

var InspectCmd = &cli.Command{
	Name:  "inspect",
	Usage: "to start the workqueue manager on a server",
	Action: func(ctx *cli.Context) error {
		return Inspect()
	},

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "id",
			Usage:       "to inpsect a message with specific id",
			Required:    false,
			Destination: &id,
		},
		&cli.StringFlag{
			Name:        "port",
			Usage:       "to inpsect a message with specific id",
			Aliases:     []string{"p"},
			Required:    true,
			Destination: &inspectPort,
		},
	},
}

func Inspect() error {
	url := "http://127.0.0.1:" + inspectPort + "/inspect"

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to make request to url %s ", url))
	}

	c := &http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to make request to url %s", url))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(os.Stdout, "%s he data was valid\n", string(body))
	return nil
}
