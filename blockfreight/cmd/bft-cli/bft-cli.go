// File: ./blockfreight/cmd/bft-cli/bft-cli.go
// Summary: Application code for Blockfreight™ | The blockchain of global freight.
// License: MIT License
// Company: Blockfreight, Inc.
// Author: Julian Nunez, Neil Tran, Julian Smith & contributors
// Site: https://blockfreight.com
// Support: <support@blockfreight.com>

// Copyright 2017 Blockfreight, Inc.

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

package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/bf_tx"
	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/validator"
	"github.com/blockfreight/blockfreight-alpha/blockfreight/crypto"
	"github.com/tendermint/abci/client"
	"github.com/tendermint/abci/types"
	. "github.com/tendermint/go-common"
	"github.com/urfave/cli"
)

//structure for data passed to print response
// variables must be exposed for JSON to read
type response struct {
	Res       types.Result
	Data      string
	PrintCode bool
	Code      string
}

func newResponse(res types.Result, data string, printCode bool) *response {
	rsp := &response{
		Res:       res,
		Data:      data,
		PrintCode: printCode,
		Code:      "",
	}

	if printCode {
		rsp.Code = res.Code.String()
	}

	return rsp
}

// client is a global variable so it can be reused by the console
var client abcicli.Client

func main() {

	//workaround for the cli library (https://github.com/urfave/cli/issues/565)
	cli.OsExiter = func(_ int) {}

	app := cli.NewApp()
	app.Name = "bft-cli"
	app.Usage = "bft-cli [command] [args...]"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
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
		/*cli.StringFlag{
		  		Name: "lang",
		  		Value: "english",
		  		Usage: "language for the greeting",
			},*/
		cli.StringFlag{
			Name:  "json_path",
			Value: "./files/bf_tx_example.json",
			Usage: "define the source path where the json is",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "batch",
			Usage: "Run a batch of bft commands against an application",
			Action: func(c *cli.Context) error {
				return cmdBatch(app, c)
			},
		},
		{
			Name:  "console",
			Usage: "Start an interactive bft console for multiple commands",
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
			Name:  "publish_bf_tx", //"deliver_tx",
			Usage: "Deliver a new bf_tx to application",
			Action: func(c *cli.Context) error {
				return cmdDeliverTx(c)
			},
		},
		{
			Name:  "check_bf_tx",
			Usage: "Validate a bf_tx",
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
		},
		{
			Name:  "validate_bf_tx",
			Usage: "Verify the structure of the bill of lading",
			Action: func(c *cli.Context) error {
				return cmdValidateBf_Tx(c)
			},
		},
	}
	app.Before = before
	err := app.Run(os.Args)
	if err != nil {
		Exit(err.Error())
	}

}

func before(c *cli.Context) error {
	introduction(c)
	if client == nil {
		var err error
		client, err = abcicli.NewClient(c.GlobalString("address"), c.GlobalString("bft"), false)
		if err != nil {
			Exit(err.Error())
		}
	}
	return nil
}

// badCmd is called when we invoke with an invalid first argument (just for console for now)
func badCmd(c *cli.Context, cmd string) {
	fmt.Println("Unknown command:", cmd)
	fmt.Println("Please try one of the following:")
	fmt.Println("")
	cli.DefaultAppComplete(c)
}

//Generates new Args array based off of previous call args to maintain flag persistence
func persistentArgs(line []byte) []string {

	//generate the arguments to run from orginal os.Args
	// to maintain flag arguments
	args := os.Args
	args = args[:len(args)-1] // remove the previous command argument

	if len(line) > 0 { //prevents introduction of extra space leading to argument parse errors
		args = append(args, strings.Split(string(line), " ")...)
	}
	return args
}

//--------------------------------------------------------------------------------

func cmdBatch(app *cli.App, c *cli.Context) error {
	bufReader := bufio.NewReader(os.Stdin)
	for {
		line, more, err := bufReader.ReadLine()
		if more {
			return errors.New("Input line is too long")
		} else if err == io.EOF {
			break
		} else if len(line) == 0 {
			continue
		} else if err != nil {
			return err
		}

		args := persistentArgs(line)
		app.Run(args) //cli prints error within its func call
	}
	return nil
}

func cmdConsole(app *cli.App, c *cli.Context) error {
	// don't hard exit on mistyped commands (eg. check vs check_tx)
	app.CommandNotFound = badCmd

	for {
		fmt.Printf("\n> ")
		bufReader := bufio.NewReader(os.Stdin)
		line, more, err := bufReader.ReadLine()
		if more {
			return errors.New("Input is too long")
		} else if err != nil {
			return err
		}

		args := persistentArgs(line)
		app.Run(args) //cli prints error within its func call
	}
}

// Have the application echo a message
func cmdEcho(c *cli.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return errors.New("Command echo takes 1 argument")
	}
	res := client.EchoSync(args[0])
	rsp := newResponse(res, string(res.Data), false)
	printResponse(c, rsp)
	return nil
}

// Get some info from the application
func cmdInfo(c *cli.Context) error {
	resInfo, err := client.InfoSync()
	if err != nil {
		return err
	}
	rsp := newResponse(types.Result{}, string(resInfo.Data), false)
	printResponse(c, rsp)
	return nil
}

// Set an option on the application
func cmdSetOption(c *cli.Context) error {
	args := c.Args()
	if len(args) != 2 {
		return errors.New("Command set_option takes 2 arguments (key, value)")
	}
	res := client.SetOptionSync(args[0], args[1])
	rsp := newResponse(res, Fmt("%s=%s", args[0], args[1]), false)
	printResponse(c, rsp)
	return nil
}

// Append a new tx to application
func cmdDeliverTx(c *cli.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return errors.New("Command deliver_tx takes 1 argument")
	}
	/*txBytes, err := stringOrHexToBytes(c.Args()[0])
	if err != nil {
		return err
	}*/

	//Sign BF_TX
	bft_tx := bf_tx.SetBF_TX(args[0])
	bft_tx = ecdsa.Sign_BF_TX(bft_tx)
	//fmt.Println(bf_tx.Signhash)
	content := bf_tx.BF_TXContent(bft_tx)

	//Save on DB
	//fmt.Println(bf_tx.Signature)
	if validator.RecordOnDB( /*bf_tx.Signature, */ content) { //TODO: Check the id
		fmt.Println("Stored on DB!")
	}

	//TODO: Validate if that JSON was already stored
	txBytes := []byte(content)
	//txBytes := []byte(bf_tx.Signature+"="+content)	//TODO: Check the id
	res := client.DeliverTxSync(txBytes)
	rsp := newResponse(res, string(res.Data), true)
	printResponse(c, rsp)
	return nil
}

// Validate a tx
func cmdCheckTx(c *cli.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return errors.New("Command check_tx takes 1 argument")
	}
	txBytes, err := stringOrHexToBytes(c.Args()[0])
	if err != nil {
		return err
	}
	res := client.CheckTxSync(txBytes)
	rsp := newResponse(res, string(res.Data), true)
	printResponse(c, rsp)
	return nil
}

// Get application Merkle root hash
func cmdCommit(c *cli.Context) error {
	res := client.CommitSync()
	rsp := newResponse(res, Fmt("0x%X", res.Data), false)
	printResponse(c, rsp)
	return nil
}

// Query application state
func cmdQuery(c *cli.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return errors.New("Command query takes 1 argument")
	}
	/*queryBytes, err := stringOrHexToBytes(c.Args()[0])
	if err != nil {
		return err
	}*/
	//queryBytes := common.ReadJSON(args[0])

	//TODO: Check the query because when the bf_tx is added to the blockchain, it is signed. But, in here is not signed. Them, doesn't find match
	bft_tx := bf_tx.SetBF_TX(args[0])
	queryBytes := []byte(bf_tx.BF_TXContent(bft_tx))
	//queryBytes := []byte(args[0])
	res := client.QuerySync(queryBytes)
	rsp := newResponse(res, string(res.Data), true)
	printResponse(c, rsp)
	return nil
}

//Verify the structure of the bill of lading
func cmdValidateBf_Tx(c *cli.Context) error {
	args := c.Args()
	if len(args) > 1 {
		return errors.New("Command validate_bf_tx takes 1 argument")
	}
	bf_tx := bf_tx.SetBF_TX(args[0])
	res := client.EchoSync(validator.ValidateBf_Tx(bf_tx))
	rsp := newResponse(res, string(res.Data), false)
	printResponse(c, rsp)
	return nil
}

//--------------------------------------------------------------------------------

func printResponse(c *cli.Context, rsp *response) {

	verbose := c.GlobalBool("verbose")

	if verbose {
		fmt.Println(">", c.Command.Name, strings.Join(c.Args(), " "))
	}

	if rsp.PrintCode {
		fmt.Printf("-> code: %s\n", rsp.Code)
	}

	//if pr.res.Error != "" {
	//	fmt.Printf("-> error: %s\n", pr.res.Error)
	//}

	if rsp.Data != "" {
		fmt.Printf("-> blockfreight data: %s\n", rsp.Data)
	}
	if rsp.Res.Log != "" {
		fmt.Printf("-> log: %s\n", rsp.Res.Log)
	}

	if verbose {
		fmt.Println("")
	}

}

// NOTE: s is interpreted as a string unless prefixed with 0x
func stringOrHexToBytes(s string) ([]byte, error) {
	if len(s) > 2 && strings.ToLower(s[:2]) == "0x" {
		b, err := hex.DecodeString(s[2:])
		if err != nil {
			err = fmt.Errorf("Error decoding hex argument: %s", err.Error())
			return nil, err
		}
		return b, nil
	}

	if !strings.HasPrefix(s, "\"") || !strings.HasSuffix(s, "\"") {
		err := fmt.Errorf("Invalid string arg: \"%s\". Must be quoted or a \"0x\"-prefixed hex string", s)
		return nil, err
	}
	return []byte(s[1 : len(s)-1]), nil
}

func introduction(c *cli.Context) {
	fmt.Println("\n...........................................")
	fmt.Println("Blockfreight™ Go App")
	fmt.Println("Address " + c.GlobalString("address"))
	fmt.Println("BFT Implementation:  " + c.GlobalString("bft"))
	fmt.Println("...........................................\n")
	/*name := "Blockfreight Community"
	  if c.NArg() > 0 {
	    name = c.Args().Get(0)
	  }
	  if c.String("lang") == "ES" {	//ISO 639-1
	    fmt.Println("Hola", name)
	  } else {
	    fmt.Println("Hello", name)
	  }*/
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
