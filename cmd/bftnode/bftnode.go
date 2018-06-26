// File: ./blockfreight/cmd/bftnode/bftnode.go
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
	// Implements command-line flag parsing.
	"fmt" // Implements formatted I/O with functions analogous to C's printf and scanf.
	"os"

	// ===============
	// Tendermint Core
	// ===============
	"github.com/tendermint/abci/client"
	tmConfig "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/node"
	tmNode "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
	"github.com/tendermint/tmlibs/log"
	// ======================
	// Blockfreight™ packages
	// ======================
	"github.com/blockfreight/go-bftx/lib/app/bft"
	// Implements the main functions to work with the Blockfreight™ Network.
)

var client abcicli.Client

var logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))

func BlockfreightAppClientCreator(addr, transport, dbDir string) proxy.ClientCreator {
	return proxy.NewLocalClientCreator(bft.NewBftApplication())
}

func main() {

	fmt.Println("Blockfreight™ Node")

	index := &tmConfig.TxIndexConfig{
		Indexer:      "kv",
		IndexTags:    "bftx.id",
		IndexAllTags: false,
	}

	config := tmConfig.DefaultConfig()

	config.P2P.Seeds = "aeabbf6b891435013f2a800fa9e22a1451ca90fd@bftx0.blockfreight.net:8888,6e9515c2cfed19464e6ce11ba2297ecdb411103b@bftx1.blockfreight.net:8888,b8b988370783bd0e58bf926d621a47160af2bdae@bftx2.blockfreight.net:8888,8c091f4e3dc4ac27db1efd38beee012d99967fd8@bftx3.blockfreight.net:8888"
	config.Consensus.CreateEmptyBlocks = false

	config.TxIndex = index
	config.DBPath = "./bft-db"
	config.Genesis = "/Users/gianfelipe/.tendermint/config/genesis.json"
	config.PrivValidator = "/Users/gianfelipe/.tendermint/config/priv_validator.json"
	config.NodeKey = "/Users/gianfelipe/.tendermint/config/node_key.json"

	logger.Info("Setting up config", "nodeInfo", config)

	node, err := tmNode.NewNode(config,
		privval.LoadOrGenFilePV(config.PrivValidatorFile()),
		BlockfreightAppClientCreator(config.ProxyApp, config.ABCI, config.DBDir()),
		tmNode.DefaultGenesisDocProviderFunc(config),
		tmNode.DefaultDBProvider,
		logger,
	)

	logger.Info("Started node", "nodeInfo", node.GenesisDoc)

	if err != nil {
		fmt.Errorf("Failed to create a node: %v", err)
	}

	if err = node.Start(); err != nil {
		fmt.Errorf("Failed to start node: %v", err)
	}

	logger.Info("Started node", "nodeInfo", node.Switch().NodeInfo())

	// Trap signal, run forever.
	node.RunForever()

}

func StartBFTXNode(config *tmConfig.Config, logger log.Logger) (*node.Node, error) {
	return tmNode.NewNode(config,
		privval.LoadOrGenFilePV(config.PrivValidatorFile()),
		BlockfreightAppClientCreator(config.ProxyApp, config.ABCI, config.DBDir()),
		tmNode.DefaultGenesisDocProviderFunc(config),
		tmNode.DefaultDBProvider,
		logger,
	)
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
