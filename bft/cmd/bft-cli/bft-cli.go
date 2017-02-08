package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/blockfreight/blockfreight-alpha/bft/client"
)

// client is a global variable, it can be reused by the console
var client bftcli.Client

func main() {
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
    	cli.StringFlag{
			Name:  "address",
			Value: "tcp://127.0.0.1:46658",
			Usage: "address of application socket",
		},
		cli.StringFlag{
			Name:  "bft",
			Value: "socket",
			Usage: "socket or grpc",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "print the command and results as if it were a console session",
		},
  	}
	//app.Action = introduction()
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
	app.Before = InitializeClient
	//app.Run(os.Args)
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func InitializeClient(c *cli.Context) error {
	fmt.Println("**********************************")
	introduction(c)
	if client == nil {
		var err error
		fmt.Println("Address "+c.GlobalString("address"))
		fmt.Println("BFT Implementation:  "+c.GlobalString("bft"))
		client, err = bftcli.NewClient(c.GlobalString("address"), c.GlobalString("bft"), false)
		if err != nil {
			fmt.Println("Error: "+err.Error())
			//Exit(err.Error())
		}
	}
	fmt.Println("**********************************\n")
	return nil
}

func introduction (c *cli.Context) {
	name := "Blockfreight Community"
    if c.NArg() > 0 {
      name = c.Args().Get(0)
    }
    if c.String("lang") == "ES" {	//ISO 639-1
      fmt.Println("Hola", name)
    } else {
      fmt.Println("Hello", name)
    }
}