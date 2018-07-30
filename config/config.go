// File: ./blockfreight/config/config.go
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

// Blockfreight™ App Configuration

// Package config is a package that handles with the application configutarions.
package config

import (
	"fmt"
	"os"

	// Implements common functions for Blockfreight™
	"github.com/BurntSushi/toml"
	tmConfig "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"
)

var homeDir = os.Getenv("HOME")
var GenesisJSONURL = "https://raw.githubusercontent.com/blockfreight/tools/master/blockfreightnet-kubernetes/examples/blockfreight/genesis.json"
var ConfigDir = homeDir + "/.blockfreight/config"
var Logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
var config = tmConfig.DefaultConfig()
var index = &tmConfig.TxIndexConfig{
	Indexer:      "kv",
	IndexTags:    "bftx.id",
	IndexAllTags: false,
}

func GetBlockfreightConfig(verbose bool) *tmConfig.Config {

	var blockfreightConfig Config

	if _, err := toml.DecodeFile("config.toml", &blockfreightConfig); err != nil {
		fmt.Println(err)
		return config
	}

	config.P2P.Seeds = blockfreightConfig.getValidatorsSeedString()
	config.Consensus.CreateEmptyBlocks = blockfreightConfig.CreateEmptyBlocks
	config.RPC.ListenAddress = blockfreightConfig.RPC_ListenAddress
	config.TxIndex = index
	config.DBPath = ConfigDir + "/data"
	config.Genesis = ConfigDir + "/genesis.json"
	config.PrivValidator = ConfigDir + "/priv_validator.json"
	config.NodeKey = ConfigDir + "/node_key.json"
	config.P2P.ListenAddress = blockfreightConfig.P2P_ListenAddress

	if !verbose {
		config.LogLevel = fmt.Sprintf("*:%s", tmConfig.DefaultLogLevel())
	}

	fmt.Printf("%+v\n", config)
	fmt.Printf("%+v\n", config.P2P)
	fmt.Printf("%+v\n", config.RPC)

	return config
}

type Config struct {
	GenesisJSON_URL   string
	RPC_ListenAddress string
	P2P_ListenAddress string
	Validator_Domain  string
	P2P_PORT          string
	CreateEmptyBlocks bool
	Validators        map[string]Validator
}

func (config Config) getValidatorsSeedString() string {
	bftx0 := config.Validators["bftx0"]
	bftx1 := config.Validators["bftx1"]
	bftx2 := config.Validators["bftx2"]
	bftx3 := config.Validators["bftx3"]

	seedsString := fmt.Sprintf("%s,%s,%s,%s", bftx0.getValidatorSeedString(config), bftx1.getValidatorSeedString(config), bftx2.getValidatorSeedString(config), bftx3.getValidatorSeedString(config))
	return seedsString
}

type Validator struct {
	NodeID        string
	ValidatorName string
}

func (validator Validator) getValidatorSeedString(config Config) string {
	seedString := fmt.Sprintf("%s@%s.%s:%s", validator.NodeID, validator.ValidatorName, config.Validator_Domain, config.P2P_PORT)
	return seedString
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
