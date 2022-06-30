package tou

import (
	"bytes"
	"context"
	"encoding/json"
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
		//fmt.Fprintf(os.Stdout, "The id is %s", id)

		postBody, err := json.Marshal(id)
		if err != nil {
			return nil
		}

		postBytes := bytes.NewBuffer(postBody)

		req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "http://127.0.0.1:"+inspectPort+"/inspect", postBytes)
		c := &http.Client{}

		resp, err := c.Do(req)
		if err != nil {
			return errors.New("Unable to make new request")
		}

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.New("Unable to read response body")
		}

		fmt.Fprintf(os.Stdout, "%s passed a valid data \n", string(respBody))
		return nil
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
