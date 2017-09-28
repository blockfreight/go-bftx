// Summary: Application code for Blockfreight™ | The blockchain of global freight.
// License: MIT License
// Company: Blockfreight, Inc.
// Author: Julian Nunez, Neil Tran, Julian Smith, Gian Felipe & contributors
// Site: https://blockfreight.com
// Support: <support@blockfreight.com>

// Copyright © 2017 Blockfreight, Inc. All Rights Reserved.

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// =================================================================================================================================================
// =================================================================================================================================================
//
// BBBBBBBBBBBb     lll                                kkk             ffff                         iii                  hhh            ttt
// BBBB``````BBBB   lll                                kkk            fff                           ```                  hhh            ttt
// BBBB      BBBB   lll      oooooo        ccccccc     kkk    kkkk  fffffff  rrr  rrr    eeeee      iii     gggggg ggg   hhh  hhhhh   tttttttt
// BBBBBBBBBBBB     lll    ooo    oooo    ccc    ccc   kkk   kkk    fffffff  rrrrrrrr eee    eeee   iii   gggg   ggggg   hhhh   hhhh  tttttttt
// BBBBBBBBBBBBBB   lll   ooo      ooo   ccc           kkkkkkk        fff    rrrr    eeeeeeeeeeeee  iii  gggg      ggg   hhh     hhh    ttt
// BBBB       BBB   lll   ooo      ooo   ccc           kkkk kkkk      fff    rrr     eeeeeeeeeeeee  iii   ggg      ggg   hhh     hhh    ttt
// BBBB      BBBB   lll   oooo    oooo   cccc    ccc   kkk   kkkk     fff    rrr      eee      eee  iii    ggg    gggg   hhh     hhh    tttt    ....
// BBBBBBBBBBBBB    lll     oooooooo       ccccccc     kkk     kkkk   fff    rrr       eeeeeeeee    iii     gggggg ggg   hhh     hhh     ttttt  ....
//                                                                                                        ggg      ggg
//   Blockfreight™ | The blockchain of global freight.                                                      ggggggggg
//
// =================================================================================================================================================
// =================================================================================================================================================

// Starts the Blockfreight™ Node to listen to all requests in the Blockfreight Network.
package main

import (
	// =======================
	// Golang Standard library
	// =======================
	"flag" // Implements command-line flag parsing.
	"fmt"  // Implements formatted I/O with functions analogous to C's printf and scanf.
	"log"  // Implements a simple logging package.
	"os"
	"strings"

	// ===============
	// Tendermint Core
	// ===============
	"github.com/tendermint/abci/client"
	"github.com/tendermint/abci/server"
	"github.com/tendermint/abci/types"
	"github.com/tendermint/abci/version"
	tendermint "github.com/tendermint/go-common"
	"github.com/urfave/cli"

	// ======================
	// Blockfreight™ packages
	// ======================
	"github.com/blockfreight/blockfreight-alpha/lib/app/bft" // Implements the main functions to work with the Blockfreight™ Network.
	"github.com/blockfreight/blockfreight-alpha/lib/pkg/key"
)

// Structure for data passed to print response.
type response struct {
	// generic abci response
	Data   []byte
	Code   types.CodeType
	Log    string
	Result string //Blockfreight Purposes
}

// client is a global variable so it can be reused by the console
var client abcicli.Client

func main() {

	//workaround for the cli library (https://github.com/urfave/cli/issues/565)
	cli.OsExiter = func(_ int) {}

	app := cli.NewApp()
	app.Name = "bftnode"
	app.Usage = "bftnode [command] [args...]"
	app.Version = version.Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address",
			Value: "tcp://127.0.0.1:46658",
			//Value: "tcp://172.17.0.3:46658",
			Usage: "address of application socket",
		},
		cli.StringFlag{
			Name:  "call",
			Value: "socket",
			Usage: "socket or grpc",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "print the command and results as if it were a console session",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start Blockfreight Node Application. (Parameters: none)",
			Action: func(c *cli.Context) {
				cmdStartServer()
			},
		},
		{
			Name:  "exit",
			Usage: "Leaves the program. (Parameters: none)",
			Action: func(c *cli.Context) {
				os.Exit(0)
			},
		},
		{
			Name:  "new_key",
			Usage: "Generate a new Public/Private Key",
			Action: func(c *cli.Context) error {
				return cmdGenerateKey(c)
			},
		},
	}
	app.Before = before
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func before(c *cli.Context) error {
	introduction(c)
	return nil
}

// cmdGenerateKey function creates a new public/private key
func cmdGenerateKey(c *cli.Context) error {
	result, err := key.GenerateNewKey()
	if err != nil {
		return nil
	}

	printResponse(c, response{
		Result: result,
	})
	return nil
}

func introduction(c *cli.Context) {
	fmt.Println("\n...........................................")
	fmt.Println("Blockfreight™ Node")
	fmt.Println("...........................................\n")
}

// cmdStartServer function creates the Blockfreight Node Application
func cmdStartServer() {

	// Parameters
	addrPtr := flag.String("addr", "tcp://0.0.0.0:46658", "Listen address")
	abciPtr := flag.String("bft", "socket", "socket | grpc")
	// persistencePtr := flag.String("persist", "", "directory to use for a database")
	flag.Parse()

	// Create the application - in memory or persisted to disk
	var app types.Application
	app = bft.NewBftApplication() //if *persistencePtr != "" => NewPersistentBftApplication(*persistencePtr)

	// Start the listener
	srv, err := server.NewServer(*addrPtr, *abciPtr, app)
	fmt.Println(srv)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Service created by " + *abciPtr + " server")

	// Wait forever
	tendermint.TrapSignal(func() {
		// Cleanup
		fmt.Println("Stopping service")
		srv.Stop()
	})

}

func printResponse(c *cli.Context, rsp response) {

	verbose := c.GlobalBool("verbose")

	if verbose {
		fmt.Println(">", c.Command.Name, strings.Join(c.Args(), " "))
	}

	if !rsp.Code.IsOK() {
		fmt.Printf("-> code: %s\n", rsp.Code.String())
	}
	if rsp.Result != "" {
		fmt.Printf("-> blockfreight result: %s\n", rsp.Result)
	}
	if len(rsp.Data) != 0 {
		//fmt.Printf("-> blockfreight data: %s\n", rsp.Data)
		fmt.Printf("-> data.hex: %X\n", rsp.Data)
	}
	if rsp.Log != "" {
		fmt.Printf("-> log: %s\n", rsp.Log)
	}
	if verbose {
		fmt.Println("")
	}

}

// =================================================
// Blockfreight™ | The blockchain of global freight.
// =================================================

// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBB                    BBBBBBBBBBBBBBBBBBB
// BBBBBBB                       BBBBBBBBBBBBBBBB
// BBBBBBB                        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBB         BBBBBBBBBBBBBBBB
// BBBBBBB                     BBBBBBBBBBBBBBBBBB
// BBBBBBB                        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBBB       BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBB       BBBBB
// BBBBBBB                       BBBB       BBBBB
// BBBBBBB                    BBBBBBB       BBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB

// ==================================================
// Blockfreight™ | The blockchain for global freight.
// ==================================================
