// File: ./blockfreight/cmd/bftx/bftx.go
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

// Initializes BFTX app to interacts with the Blockfreight™ Network.
package main

import (
	// =======================
	// Golang Standard library
	// =======================
	"bufio"        // Implements buffered I/O.
	"encoding/hex" // Implements hexadecimal encoding and decoding.
	"errors"       // Implements functions to manipulate errors.
	"fmt"          // Implements formatted I/O with functions analogous to C's printf and scanf.
	"io"           // Provides basic interfaces to I/O primitives.
	"log"          // Implements a simple logging package.
	"os"           // Provides a platform-independent interface to operating system functionality.
	"strings"      // Implements simple functions to manipulate UTF-8 encoded strings.

	// ====================
	// Third-party packages
	// ====================
	"github.com/urfave/cli" // Provides structure and function to build command line apps in Go.

	// ===============
	// Tendermint Core
	// ===============
	"github.com/tendermint/abci/client"
	"github.com/tendermint/abci/types"

	// ======================
	// Blockfreight™ packages
	// ======================
	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/bf_tx"     // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/crypto"    // Provides useful functions to sign BF_TX.
	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/leveldb"   // Provides some useful functions to work with LevelDB.
	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/validator" // Provides functions to assure the input JSON is correct.
)

// Structure for data passed to print response.
type response struct {
	// generic abci response
	Data []byte
	Code types.CodeType
	Log  string

	Query *queryResponse
}

type queryResponse struct {
	Key    []byte
	Value  []byte
	Height uint64
	Proof  []byte
}

// client is a global variable so it can be reused by the console
var client abcicli.Client

func main() {

	//workaround for the cli library (https://github.com/urfave/cli/issues/565)
	cli.OsExiter = func(_ int) {}

	app := cli.NewApp()
	app.Name = "bftx"
	app.Usage = "bftx [command] [args...]"
	app.Version = "0.1.0" //version.Version   //TODO JCNM: Check the way for doing in version library mode.
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address",
			Value: "tcp://127.0.0.1:46658",
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
		/*cli.StringFlag{
		    Name: "lang",
		    Value: "english",
		    Usage: "language for the greeting",
		},*/
		cli.StringFlag{
			Name:  "json_path, jp",
			Value: "./files/bf_tx_example.json",
			Usage: "define the source path where the json is",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "batch",
			Usage: "Run a batch of Blockfreight™ commands against an application",
			Action: func(c *cli.Context) error {
				return cmdBatch(app, c)
			},
		},
		{
			Name:  "console",
			Usage: "Start an interactive Blockfreight™ console for multiple commands",
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
			Name:  "broadcast",
			Usage: "Deliver a new BF_TX to application",
			Action: func(c *cli.Context) error {
				return cmdDeliverBfTx(c)
			},
		},
		{
			Name:  "validate",
			Usage: "Validate a bf_tx",
			Action: func(c *cli.Context) error {
				return cmdCheckBfTx(c)
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
			Name:  "verify",
			Usage: "Verify the structure of the Blockfreight™ Transaction [BF_TX]",
			Action: func(c *cli.Context) error {
				return cmdValidateBf_Tx(c)
			},
		},
		{
			Name:  "get",
			Usage: "Retrieve a [BF_TX] by its ID",
			Action: func(c *cli.Context) error {
				return cmdReturnBf_Tx(c)
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
	if client == nil {
		var err error
		client, err = abcicli.NewClient(c.GlobalString("address"), c.GlobalString("bft"), false)
		if err != nil {
			log.Fatal(err.Error())
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

// Generates new Args array based off of previous call args to maintain flag persistence
func persistentArgs(line []byte) []string {

	// generate the arguments to run from orginal os.Args
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
	resEcho := client.EchoSync(args[0])
	printResponse(c, response{
		Data: resEcho.Data,
	})
	return nil
}

// Get some info from the application
func cmdInfo(c *cli.Context) error {
	resInfo, err := client.InfoSync()
	if err != nil {
		return err
	}
	printResponse(c, response{
		Data: []byte(resInfo.Data),
	})
	return nil
}

// Set an option on the application
func cmdSetOption(c *cli.Context) error {
	args := c.Args()
	if len(args) != 2 {
		return errors.New("Command set_option takes 2 arguments (key, value)")
	}
	resSetOption := client.SetOptionSync(args[0], args[1])
	printResponse(c, response{
		Log: resSetOption.Log,
	})
	return nil
}

// Append a new bf_tx to application
func cmdDeliverBfTx(c *cli.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return errors.New("Command deliver_bf_tx takes 1 argument")
	}

	// Sign BF_TX
	bft_tx := bf_tx.SetBF_TX(args[0])
	bft_tx = crypto.Sign_BF_TX(bft_tx)

	// Get the BF_TX content in string format
	content := bf_tx.BF_TXContent(bft_tx)

	// Save on DB
	if leveldb.RecordOnDB( /*bf_tx.Signature, */ content) { //TODO JCNM: Check the id
		fmt.Println("Stored on DB!")
	}

	// TODO JCNM: Validate if that JSON was already stored
	txBytes := []byte(content)
	// txBytes := []byte(bf_tx.Signature+"="+content)    //TODO JCNM: Check the id

	res := client.DeliverTxSync(txBytes)
	printResponse(c, response{
		Code: res.Code,
		Data: res.Data,
		Log:  res.Log,
	})
	return nil
}

// Validate a bf_tx
func cmdCheckBfTx(c *cli.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return errors.New("Command check_bf_tx takes 1 argument")
	}
	txBytes, err := stringOrHexToBytes(c.Args()[0])
	if err != nil {
		return err
	}
	res := client.CheckTxSync(txBytes)
	printResponse(c, response{
		Code: res.Code,
		Data: res.Data,
		Log:  res.Log,
	})
	return nil
}

// Get application Merkle root hash
func cmdCommit(c *cli.Context) error {
	res := client.CommitSync()
	printResponse(c, response{
		Code: res.Code,
		Data: res.Data,
		Log:  res.Log,
	})
	return nil
}

// Query application state
// TODO: Make request and response support all fields.
func cmdQuery(c *cli.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return errors.New("Command query takes 1 argument")
	}

	// TODO JCNM: Check the query because when the bf_tx is added to the blockchain, it is signed. But, in here is not signed. Them, doesn't find match
	// TODO JCNM: Query from blockchain
	bft_tx := bf_tx.SetBF_TX(args[0])
	queryBytes := []byte(bf_tx.BF_TXContent(bft_tx))

	resQuery, err := client.QuerySync(types.RequestQuery{
		Data:   queryBytes,
		Path:   "/store", // TOOD expose
		Height: 0,        // TODO expose
		//Prove:  true,     // TODO expose
	})
	if err != nil {
		return err
	}
	printResponse(c, response{
		Code: resQuery.Code,
		Log:  resQuery.Log,
		Query: &queryResponse{
			Key:    resQuery.Key,
			Value:  resQuery.Value,
			Height: resQuery.Height,
			Proof:  resQuery.Proof,
		},
	})
	return nil
}

// Verify the structure of the input Bill of Lading
func cmdValidateBf_Tx(c *cli.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return errors.New("Command validate_bf_tx takes 1 argument")
	}
	bf_tx := bf_tx.SetBF_TX(args[0])
	resEcho := client.EchoSync(validator.ValidateBf_Tx(bf_tx))
	printResponse(c, response{
		Data: resEcho.Data,
	})

	return nil
}

// how the output JSON
func cmdReturnBf_Tx(c *cli.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return errors.New("Command get_bf_tx takes 1 argument")
	}
	resEcho := client.EchoSync(leveldb.GetBfTx(args[0]))
	printResponse(c, response{
		Data: resEcho.Data,
	})

	return nil
}

//--------------------------------------------------------------------------------

func printResponse(c *cli.Context, rsp response) {

	verbose := c.GlobalBool("verbose")

	if verbose {
		fmt.Println(">", c.Command.Name, strings.Join(c.Args(), " "))
	}

	if !rsp.Code.IsOK() {
		fmt.Printf("-> code: %s\n", rsp.Code.String())
	}
	if len(rsp.Data) != 0 {
		fmt.Printf("-> blockfreight data: %s\n", rsp.Data)
		fmt.Printf("-> data.hex: %X\n", rsp.Data)
	}
	if rsp.Log != "" {
		fmt.Printf("-> log: %s\n", rsp.Log)
	}

	if rsp.Query != nil {
		fmt.Printf("-> height: %d\n", rsp.Query.Height)
		if rsp.Query.Key != nil {
			fmt.Printf("-> key: %s\n", rsp.Query.Key)
			fmt.Printf("-> key.hex: %X\n", rsp.Query.Key)
		}
		if rsp.Query.Value != nil {
			fmt.Printf("-> value: %s\n", rsp.Query.Value)
			fmt.Printf("-> value.hex: %X\n", rsp.Query.Value)
		}
		if rsp.Query.Proof != nil {
			fmt.Printf("-> proof: %X\n", rsp.Query.Proof)
		}
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
	  if c.String("lang") == "ES" { //ISO 639-1
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
