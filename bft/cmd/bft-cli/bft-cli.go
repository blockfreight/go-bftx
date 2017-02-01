package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	//cli.NewApp().Run(os.Args)
	app := cli.NewApp()
	app.Name = "bft-cli"
	app.Usage = "bft-cli app [command] [args...]"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag {
		cli.StringFlag{
      		Name: "lang",
      		Value: "english",
      		Usage: "language for the greeting",
    	},
  	}
	app.Action = func (c *cli.Context) error {
		//fmt.Println("boom!")
		name := "Blockfreight Community"
    	if c.NArg() > 0 {
    	  name = c.Args().Get(0)
    	}
    	if c.String("lang") == "ES" {	//ISO 639-1
    	  fmt.Println("Hola", name)
    	} else {
    	  fmt.Println("Hello", name)
    	}
    	return nil
	}
	app.Commands = []cli.Command{
		{
			Name:  "batch",
			Usage: "Run a batch of bft commands against an application",
			Action: func(c *cli.Context) error {
				fmt.Println("-> This is the content of batch command.")
				return nil
			},
		},
		{
			Name:  "console",
			Usage: "Start an interactive bft console for multiple commands",
			Action: func(c *cli.Context) error {
				fmt.Println("-> This is the content of console command.")
				return nil
			},
		},
		{
			Name:  "echo",
			Usage: "Print a message",
			Action: func(c *cli.Context) error {
				fmt.Println("-> This is the content of echo command.")
				return nil
			},
		},
		{
			Name: "info",
			Usage: "Get info about bft-cli",
			Action: func (c *cli.Context) error {
				fmt.Println("-> This is the content of info command.")
				return nil
			},
		},
		{
			Name:  "set_option",
			Usage: "Set an option on the application",
			Action: func(c *cli.Context) error {
				fmt.Println("-> This is the content of set_option command.")
				return nil
			},
		},
		{
			Name:  "bftx_publish",	//deliver_bftx
			Usage: "Publish/Deliver a new bftx to application",
			Action: func(c *cli.Context) error {
				fmt.Println("-> This is the content of bftx_publish command.")
				return nil
			},
		},
		{
			Name:  "bftx_validate",	//check_bftx
			Usage: "Validate a bftx",
			Action: func(c *cli.Context) error {
				fmt.Println("-> This is the content of bftx_validate command.")
				return nil
			},
		},
		{
			Name:  "bftx_construct",	//commit
			Usage: "Commit the application state and return the Merkle root hash",
			Action: func(c *cli.Context) error {
				fmt.Println("-> This is the content of bftx_construct command.")
				return nil
			},
		},
		{
			Name:  "query",
			Usage: "Query application state",
			Action: func(c *cli.Context) error {
				fmt.Println("-> This is the content of query command.")
				return nil
			},
		},
	}
	app.Run(os.Args)
}
