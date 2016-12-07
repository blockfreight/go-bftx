package main

import (
	"fmt"
	"github.com/urfave/cli"
)

// client is a global variable so it can be reused by the console
var client tmspcli.cli

func main(){
	bftApp := cli.NewApp()
	/*app.Name = "bft-cli"
	app.Usage = "bft-cli [command] [args...]"*/
	app.Version = "0.0.1"
	/*app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address",
			Value: "tcp://127.0.0.1:46658",
			Usage: "address of application socket",
		},
		cli.StringFlag{
			Name:  "tmsp",
			Value: "socket",
			Usage: "socket or grpc",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "print the command and results as if it were a console session",
		},
	}*/
	app.Commands = []cli.Command{
		/*{
			Name:  "batch",
			Usage: "Run a batch of tmsp commands against an application",
			Action: func(c *cli.Context) error {
				return cmdBatch(app, c)
			},
		},
		{
			Name:  "console",
			Usage: "Start an interactive tmsp console for multiple commands",
			Action: func(c *cli.Context) error {
				return cmdConsole(app, c)
			},
		},
		{
			Name:  "echo",
			Usage: "Have the application echo a message",
			Action: func(c *cli.Context) error {
				return cmdEcho(c)
			},
		},
		{
			Name:  "info",
			Usage: "Get some info about the application",
			Action: func(c *cli.Context) error {
				return cmdInfo(c)
			},
		},
		{
			Name:  "set_option",
			Usage: "Set an option on the application",
			Action: func(c *cli.Context) error {
				return cmdSetOption(c)
			},
		},
		{
			Name:  "append_tx",
			Usage: "Append a new tx to application",
			Action: func(c *cli.Context) error {
				return cmdAppendTx(c)
			},
		},
		{
			Name:  "check_tx",
			Usage: "Validate a tx",
			Action: func(c *cli.Context) error {
				return cmdCheckTx(c)
			},
		},
		{
			Name:  "commit",
			Usage: "Commit the application state and return the Merkle root hash",
			Action: func(c *cli.Context) error {
				return cmdCommit(c)
			},
		},
		{
			Name:  "query",
			Usage: "Query application state",
			Action: func(c *cli.Context) error {
				return cmdQuery(c)
			},
		},*/
		{
			Name:  "hello",
			Usage: "Test Hello World",
			Action: func(c *cli.Context, app.Version) error {
				return bftHello(c)
			},
		},
	}
	/*app.Before = before
	err := app.Run(os.Args)
	if err != nil {
		Exit(err.Error())
	}*/
}

// Test bftHello
func bftHello(c *cli.Context, version string) error {
	args := c.Args()
	//res := client.SetOptionSync(args[0], args[1])
	//printResponse(c, res, Fmt("%s=%s", args[0], args[1]), false)
	fmt.Println("Hello Blockfreight™ world!")
	fmt.Println("Version: "+version)
	return nil
}